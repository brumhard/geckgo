package pkg

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var _ = Describe("repo", func() {
	const (
		port     = 5432
		user     = "postgres"
		password = "secret-af"
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
		resource, err := pool.RunWithOptions(&dockertest.RunOptions{
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
			"host=localhost port=%d user=%s password=%s dbname=postgres sslmode=disable",
			resource.GetPort(fmt.Sprintf("%s/tcp", port)), user, password,
		)

		err = pool.Retry(func() error {
			db, err = sqlx.Connect("postgres", connectStr)

			return err
		})
		Expect(err).ToNot(HaveOccurred())

		return []byte(connectStr) // return connect string here?
	}, func(data []byte) {
		var err error
		db, err = sqlx.Connect("postgres", string(data))
		Expect(err).ToNot(HaveOccurred())
	})
	BeforeEach(func() {

	})
	SynchronizedAfterSuite(func() {}, func() {
		Expect(pool.Purge(resource))
	})
})
