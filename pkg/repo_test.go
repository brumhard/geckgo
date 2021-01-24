package pkg

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"

	"github.com/go-kit/kit/log"

	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"

	"github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var _ = Describe("repo", func() {
	const (
		port     = 5432
		user     = "postgres"
		password = "secret-af"
		database = "postgres"
		driver   = "pgx"
	)
	var (
		pool     *dockertest.Pool
		resource *dockertest.Resource
		db       *sqlx.DB
	)
	SynchronizedBeforeSuite(func() []byte {
		// uses a sensible default on windows (tcp/http) and linux/osx (socket)
		var err error
		pool, err = dockertest.NewPool("")
		Expect(err).ToNot(HaveOccurred())

		// pulls an image, creates a container based on it and runs it
		resource, err = pool.RunWithOptions(&dockertest.RunOptions{
			Repository: "postgres",
			Tag:        "13.1",
			Env: []string{
				fmt.Sprintf("POSTGRES_USER=%s", user),
				fmt.Sprintf("POSTGRES_PASSWORD=%s", password),
				"listen_addresses = '*'", // listen to all interfaces
			},
		}, func(config *docker.HostConfig) {
			// set AutoRemove to true so that stopped container goes away by itself
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{
				Name: "no",
			}
		})
		Expect(err).ToNot(HaveOccurred())

		// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
		connectString := fmt.Sprintf(
			"host=localhost port=%s user=%s password=%s dbname=%s sslmode=disable",
			resource.GetPort(fmt.Sprintf("%d/tcp", port)), user, password, database,
		)

		err = pool.Retry(func() error {
			db, err = sqlx.Connect(driver, connectString)

			return err
		})
		Expect(err).ToNot(HaveOccurred())

		return []byte(connectString) // return connect string here?
	}, func(connectString []byte) {
		var err error
		db, err = sqlx.Connect(driver, string(connectString))
		Expect(err).ToNot(HaveOccurred())
	})
	SynchronizedAfterSuite(func() {}, func() {
		Expect(pool.Purge(resource))
	})

	var (
		repo       Repository
		migrations *migrate.Migrate
	)
	BeforeEach(func() {
		repo = NewRepository(db, log.NewNopLogger())

		// move initialization to synchronized before suite
		driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
		Expect(err).ToNot(HaveOccurred())
		migrations, err = migrate.NewWithDatabaseInstance(
			"file://../db/migrations",
			database, driver)
		Expect(err).ToNot(HaveOccurred())
		Expect(migrations.Up()).To(Succeed())
	})
	AfterEach(func() {
		Expect(migrations.Down()).To(Succeed())
	})
	Describe("Integration", func() {
		var ctx = context.Background()

		Context("List", func() {
			var id int

			Context("no settings", func() {
				BeforeEach(func() {
					// AddList
					testList := List{Name: "miau"}

					var err error
					id, err = repo.AddList(ctx, testList)
					Expect(err).ToNot(HaveOccurred())

					Expect(id).To(Equal(1))
				})
				Describe("AddList -> AddList", func() {
					It("works because of id", func() {
						testList := List{Name: "miau"}

						otherID, err := repo.AddList(ctx, testList)
						Expect(err).ToNot(HaveOccurred())

						Expect(otherID).To(Equal(2))
					})
				})
				Describe("AddList -> GetListByID", func() {
					It("works", func() {
						list, err := repo.GetListByID(ctx, id)
						Expect(err).ToNot(HaveOccurred())
						Expect(list).To(Equal(&List{Name: "miau", ID: id, Settings: nil}))
					})
				})
				Describe("AddList -> GetLists", func() {
					It("works", func() {
						lists, err := repo.GetLists(ctx)
						Expect(err).ToNot(HaveOccurred())
						Expect(lists).To(Equal([]List{{Name: "miau", ID: id, Settings: nil}}))
					})
				})
				Describe("AddList -> UpdateList -> GetListById", func() {
					It("works", func() {
						err := repo.UpdateList(ctx, List{ID: 1, Name: "lol"})
						Expect(err).ToNot(HaveOccurred())

						lists, err := repo.GetListByID(ctx, id)
						Expect(err).ToNot(HaveOccurred())
						Expect(lists).To(Equal(&List{Name: "lol", ID: 1, Settings: nil}))
					})
					Context("add settings", func() {
						It("works", func() {
							dailyTime := 8 * time.Hour
							updated := List{Name: "miau", ID: id, Settings: &ListSettings{DailyTime: &Duration{dailyTime}}}
							err := repo.UpdateList(ctx, updated)
							Expect(err).To(Succeed())

							lists, err := repo.GetListByID(ctx, id)
							Expect(err).ToNot(HaveOccurred())
							Expect(lists).To(Equal(&List{Name: "miau", ID: id, Settings: &ListSettings{DailyTime: &Duration{dailyTime}}}))
						})
					})
				})
				Describe("DeleteListByID non existing", func() {
					It("does not return err", func() {
						Expect(repo.DeleteListByID(ctx, 420)).To(Succeed())
					})
				})
				Describe("UpdateList non existing", func() {
					It("returns err", func() {
						Expect(repo.UpdateList(ctx, List{ID: 420, Name: "lel"})).ToNot(Succeed())
					})
				})
			})
			Context("settings set", func() {
				var dailyTime time.Duration
				BeforeEach(func() {
					// AddList with settings
					dailyTime = 8 * time.Hour
					testList := List{Name: "miau", Settings: &ListSettings{DailyTime: &Duration{dailyTime}}}

					var err error
					id, err = repo.AddList(ctx, testList)
					Expect(err).ToNot(HaveOccurred())

					Expect(id).To(Equal(1))
				})
				Describe("AddList -> GetListByID", func() {
					It("works", func() {
						list, err := repo.GetListByID(ctx, id)
						Expect(err).ToNot(HaveOccurred())
						Expect(list).To(Equal(&List{Name: "miau", ID: id, Settings: &ListSettings{DailyTime: &Duration{dailyTime}}}))
					})
				})
				Describe("AddList -> GetLists", func() {
					It("works", func() {
						lists, err := repo.GetLists(ctx)
						Expect(err).ToNot(HaveOccurred())
						Expect(lists).To(Equal([]List{{Name: "miau", ID: id, Settings: &ListSettings{DailyTime: &Duration{dailyTime}}}}))
					})
				})
				Describe("AddList -> UpdateList -> GetListById", func() {
					It("updates settings", func() {
						updatedDailyTime := 9 * time.Hour
						updated := List{Name: "miau", ID: id, Settings: &ListSettings{DailyTime: &Duration{updatedDailyTime}}}
						err := repo.UpdateList(ctx, updated)
						Expect(err).To(Succeed())

						lists, err := repo.GetListByID(ctx, id)
						Expect(err).ToNot(HaveOccurred())
						Expect(lists).To(Equal(&List{Name: "miau", ID: id, Settings: &ListSettings{DailyTime: &Duration{updatedDailyTime}}}))
					})
				})
				Describe("AddList -> DeleteListByID -> GetListById", func() {
					It("deletes the list and all related settings", func() {
						Expect(repo.DeleteListByID(ctx, id)).To(Succeed())

						_, err := repo.GetListByID(ctx, id)
						Expect(err).To(HaveOccurred())

						rows, err := db.QueryContext(ctx, "SELECT * FROM list_settings")
						Expect(err).ToNot(HaveOccurred())

						Expect(rows.Next()).To(BeFalse())
					})
				})
			})
		})

		Context("Day", func() {
			Context("list not available", func() {
				Describe("foreign key constraint", func() {
					It("fails", func() {
						now := time.Now()
						toAdd := Day{Date: now, Moments: []Moment{{Type: MomentTypeStart, Time: now}}}
						err := repo.AddDay(ctx, 1, toAdd)
						Expect(err).To(HaveOccurred())
					})
				})
			})

			Context("list available", func() {
				var (
					id  int
					now time.Time
					day Day
				)

				BeforeEach(func() {
					var err error
					id, err = repo.AddList(ctx, List{Name: "lul"})
					Expect(err).ToNot(HaveOccurred())

					now = time.Now()
					day = Day{Date: now, Moments: []Moment{{Type: MomentTypeEnd, Time: now}}}
					err = repo.AddDay(ctx, id, day)
					Expect(err).ToNot(HaveOccurred())
				})
				Describe("AddDay -> AddDay", func() {
					It("fails", func() {
						err := repo.AddDay(ctx, id, day)
						Expect(err).To(HaveOccurred())
					})
				})
				Describe("AddDay -> GetDayByDate", func() {
					It("works", func() {
						got, err := repo.GetDayByDate(ctx, id, now)
						Expect(err).ToNot(HaveOccurred())

						expected := &Day{
							Date:    now.UTC().Truncate(24 * time.Hour),               // now does not contain any time information
							Moments: []Moment{{Type: MomentTypeEnd, Time: now.UTC()}}, // everything in UTC
						}
						Expect(got).To(Equal(expected))
					})
				})
				Describe("AddDay -> GetDays", func() {
					It("works", func() {
						got, err := repo.GetDays(ctx, id)
						Expect(err).ToNot(HaveOccurred())

						expected := []Day{{
							Date:    now.UTC().Truncate(24 * time.Hour),               // now does not contain any time information
							Moments: []Moment{{Type: MomentTypeEnd, Time: now.UTC()}}, // everything in UTC
						}}
						Expect(got).To(Equal(expected))
					})
				})
				Describe("AddDay -> UpdateDay -> GetDayByDate", func() {
					It("works", func() {
						addedMoment := Moment{Type: MomentTypeStart, Time: now.UTC()}
						day.Moments = append(day.Moments, addedMoment)
						err := repo.UpdateDay(ctx, id, day)
						Expect(err).ToNot(HaveOccurred())

						got, err := repo.GetDayByDate(ctx, id, now)
						Expect(err).ToNot(HaveOccurred())

						expected := &Day{
							Date:    now.UTC().Truncate(24 * time.Hour),                            // now does not contain any time information
							Moments: []Moment{{Type: MomentTypeEnd, Time: now.UTC()}, addedMoment}, // everything in UTC
						}
						Expect(got).To(Equal(expected))
					})
				})
				Describe("AddDay -> UpdateDay -> GetDays", func() {
					It("works", func() {
						addedMoment := Moment{Type: MomentTypeStart, Time: now.UTC()}
						day.Moments = append(day.Moments, addedMoment)

						err := repo.UpdateDay(ctx, id, day)
						Expect(err).ToNot(HaveOccurred())

						got, err := repo.GetDays(ctx, id)
						Expect(err).ToNot(HaveOccurred())

						expected := []Day{{
							Date:    now.UTC().Truncate(24 * time.Hour),                            // now does not contain any time information
							Moments: []Moment{{Type: MomentTypeEnd, Time: now.UTC()}, addedMoment}, // everything in UTC
						}}
						Expect(got).To(Equal(expected))
					})
				})
				Describe("AddDay -> DeleteDayByDate -> GetDays", func() {
					It("works", func() {
						err := repo.DeleteDayByDate(ctx, id, now)
						Expect(err).ToNot(HaveOccurred())

						got, err := repo.GetDays(ctx, id)
						Expect(err).ToNot(HaveOccurred())

						Expect(got).To(BeNil())
					})
				})
				Describe("DeleteDayByDate non existing", func() {
					It("does not return err", func() {
						Expect(repo.DeleteDayByDate(ctx, id, time.Now().AddDate(0, 0, 1))).To(Succeed())
					})
				})
			})
		})
	})
})
