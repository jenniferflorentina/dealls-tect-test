package models_test

import (
	"testing"

	"jennifer/dealls-tech-test/internal/server/config"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/postgres"
	"github.com/spf13/viper"
)

func TestModels(t *testing.T) {
	var container *gnomock.Container

	BeforeSuite(func() {
		var err error

		dbName := "testdb"
		dbUser := "testuser"
		dbPassword := "testpassword"

		p := postgres.Preset(postgres.WithDatabase(dbName), postgres.WithUser(dbUser, dbPassword))
		container, err = gnomock.Start(p)
		Expect(err).NotTo(HaveOccurred())

		viper.Set(config.DatabaseUser, dbUser)
		viper.Set(config.DatabasePassword, dbPassword)
		viper.Set(config.DatabaseHost, container.Host)
		viper.Set(config.DatabasePort, container.DefaultPort())
		viper.Set(config.DatabaseName, dbName)
	})

	AfterSuite(func() {
		err := gnomock.Stop(container)
		if err != nil {
			return
		}
	})

	RegisterFailHandler(Fail)
	RunSpecs(t, "Models Suite")
}
