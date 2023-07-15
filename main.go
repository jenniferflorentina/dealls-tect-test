package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"jennifer/dealls-tech-test/internal/domain/models"

	"jennifer/dealls-tech-test/internal/server"
	"jennifer/dealls-tech-test/internal/server/config"
	"jennifer/dealls-tech-test/internal/server/middlewares"

	"github.com/spf13/viper"
)

var buildTime string = "now"

func init() {
	config.Setup()
	viper.Set("buildTime", buildTime)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	db, err := server.Conn(ctx)
	if err != nil {
		log.Fatal("Unable to connect to database: ", err)
		cancel()
		return
	}

	// Migration will be initiated using database migration script

	err = models.AutoMigrate(db)
	if err != nil {
		log.Fatal("Unable to migrate database: ", err)
		cancel()
		return
	}

	middlewares.SetupLogger()
	routes := middlewares.LogHandler(server.Routes(db))
	routes = middlewares.CORSHandler(routes)
	server := server.InitServer(ctx, routes)

	// Accepts graceful shutdowns when quitting via SIGINT (Ctrl + C)
	// SIGKILL, SIGQUIT or SIGTERM will not be caught and will forcefully shuts the application down.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Blocks until we receive graceful shutdown signal
	<-c

	err = server.Shutdown(ctx)
	if err != nil {
		return
	}

	cancel()
	<-ctx.Done()
}
