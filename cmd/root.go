package cmd

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Akanibekuly/tsarka-tz/internal/adapters/cache/redis"
	"github.com/Akanibekuly/tsarka-tz/internal/adapters/httpapi/httpc"
	"github.com/Akanibekuly/tsarka-tz/internal/adapters/logger/zap"
	"github.com/Akanibekuly/tsarka-tz/internal/domain/core"
	"github.com/Akanibekuly/tsarka-tz/internal/interfaces"
	"github.com/spf13/viper"
)

func Execute() {
	var err error

	loadConf()

	app := struct {
		lg      interfaces.Logger
		cache   interfaces.Cache
		core    *core.St
		restApi *httpc.St
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

	app.core = core.New(app.lg, app.cache)

	restApiEChan := make(chan error, 1)
	app.restApi = httpc.New(
		app.lg,
		viper.GetString("HTTP_LISTEN"),
		restApiEChan,
		app.core,
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
