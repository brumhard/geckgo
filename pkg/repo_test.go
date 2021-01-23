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
		connectStr := fmt.Sprintf(
			"host=localhost port=%s user=%s password=%s dbname=%s sslmode=disable",
			resource.GetPort(fmt.Sprintf("%d/tcp", port)), user, password, database,
		)

		err = pool.Retry(func() error {
			db, err = sqlx.Connect(driver, connectStr)

			return err
		})
		Expect(err).ToNot(HaveOccurred())

		return []byte(connectStr) // return connect string here?
	}, func(data []byte) {
		var err error
		db, err = sqlx.Connect(driver, string(data))
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
					Expect(repo.UpdateList(ctx, List{ID: 1, Name: "lol"})).To(Succeed())

					lists, err := repo.GetListByID(ctx, id)
					Expect(err).ToNot(HaveOccurred())
					Expect(lists).To(Equal(&List{Name: "lol", ID: 1, Settings: nil}))
				})
				Context("add settings", func() {
					It("works", func() {
						dailyTime := 8 * time.Hour
						updated := List{Name: "miau", ID: id, Settings: &ListSettings{DailyTime: &Duration{dailyTime}}}
						Expect(repo.UpdateList(ctx, updated)).To(Succeed())

						lists, err := repo.GetListByID(ctx, id)
						Expect(err).ToNot(HaveOccurred())
						Expect(lists).To(Equal(&List{Name: "miau", ID: id, Settings: &ListSettings{DailyTime: &Duration{dailyTime}}}))
					})
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
					Expect(repo.UpdateList(ctx, updated)).To(Succeed())

					lists, err := repo.GetListByID(ctx, id)
					Expect(err).ToNot(HaveOccurred())
					Expect(lists).To(Equal(&List{Name: "miau", ID: id, Settings: &ListSettings{DailyTime: &Duration{updatedDailyTime}}}))
				})
			})
		})
	})
})
