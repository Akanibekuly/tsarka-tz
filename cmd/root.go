package cmd

import (
	"context"
	"github.com/Akanibekuly/tsarka-tz/internal/adapters/db"
	"github.com/jackc/pgx/v4"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Akanibekuly/tsarka-tz/internal/adapters/cache/redis"
	"github.com/Akanibekuly/tsarka-tz/internal/adapters/httpapi/httpc"
	"github.com/Akanibekuly/tsarka-tz/internal/adapters/logger/zap"
	"github.com/Akanibekuly/tsarka-tz/internal/domain/services"
	"github.com/Akanibekuly/tsarka-tz/internal/interfaces"
	"github.com/spf13/viper"
)

func Execute() {
	var err error

	loadConf()

	app := struct {
		lg       interfaces.Logger
		cache    interfaces.Cache
		reps     *db.Repository
		services *services.Services
		restApi  *httpc.St
	}{}

	debug := viper.GetBool("debug")
	app.lg, err = zap.New(viper.GetString("log_level"), debug, false)
	if err != nil {
		log.Fatal(err)
	}

	app.cache = redis.New(
		app.lg,
		viper.GetString("REDIS_URL"),
		viper.GetString("REDIS_PSW"),
		viper.GetInt("REDIS_DB"),
	)

	conn, err := pgx.Connect(context.Background(), viper.GetString("PG_DSN"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := conn.Close(context.Background())
		app.lg.Errorw("database close", err)
	}()

	app.reps = db.New(app.lg, conn)
	if err != nil {
		log.Fatal(err)
	}

	app.services = services.New(app.lg, app.cache, app.reps)

	restApiEChan := make(chan error, 1)
	app.restApi = httpc.New(
		app.lg,
		viper.GetString("HTTP_LISTEN"),
		restApiEChan,
		app.services,
	)

	app.restApi.Start()

	stopSignalChan := make(chan os.Signal, 1)
	signal.Notify(stopSignalChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	var exitCode int

	select {
	case <-stopSignalChan:
	case <-restApiEChan:
		exitCode = 1
	}

	app.lg.Infow("Shutting down...")

	err = app.restApi.Shutdown(20 * time.Second)
	if err != nil {
		app.lg.Errorw("Fail to shutdown http-api", err)
		exitCode = 1
	}

	os.Exit(exitCode)

}

func loadConf() {
	viper.SetDefault("debug", "false")
	viper.SetDefault("http_listen", ":80")
	viper.SetDefault("log_level", "info")

	confFilePath := os.Getenv("CONF_PATH")
	if confFilePath == "" {
		confFilePath = "conf.yml"
	}

	viper.SetConfigFile(confFilePath)
	_ = viper.ReadInConfig()

	viper.AutomaticEnv()
}
