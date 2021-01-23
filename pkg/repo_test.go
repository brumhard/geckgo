package pkg

import (
	"context"
	"fmt"

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
			// TODO: is this the right URL?
			"file://../db/migrations",
			database, driver)
		Expect(err).ToNot(HaveOccurred())
		Expect(migrations.Up()).To(Succeed())
	})
	AfterEach(func() {
		Expect(migrations.Down()).To(Succeed())
	})
	Describe("Integration", func() {
		Context("AddList -> GetListByID", func() {
			It("works", func() {
				ctx := context.Background()
				testList := List{Name: "miau"}

				id, err := repo.AddList(ctx, testList)
				Expect(err).ToNot(HaveOccurred())

				Expect(id).To(Equal(1))

				list, err := repo.GetListByID(ctx, id)
				Expect(err).ToNot(HaveOccurred())
				Expect(list).To(Equal(&List{Name: "miau", ID: 1, Settings: nil}))
			})
		})

	})
})
