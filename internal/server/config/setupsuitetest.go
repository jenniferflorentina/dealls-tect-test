package config

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/postgres"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func SetupSuiteTest(container *gnomock.Container) {
	BeforeSuite(func() {
		var err error

		dbName := "testdb"
		dbUser := "testuser"
		dbPassword := "testpassword"

		p := postgres.Preset(postgres.WithDatabase(dbName), postgres.WithUser(dbUser, dbPassword))
		container, err = gnomock.Start(p)
		Expect(err).NotTo(HaveOccurred())

		viper.Set(DatabaseUser, dbUser)
		viper.Set(DatabasePassword, dbPassword)
		viper.Set(DatabaseHost, container.Host)
		viper.Set(DatabasePort, container.DefaultPort())
		viper.Set(DatabaseName, dbName)
	})

	AfterSuite(func() {
		w := httptest.NewRecorder()
		err := gnomock.Stop(container)
		if err != nil {
			log.Error().Err(err).Msg("Custom Redis 2 failed")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	})
}
