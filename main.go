package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/phassans/exville/clients/rocket"
	"github.com/phassans/exville/common"
	"github.com/phassans/exville/db"
	"github.com/phassans/exville/engines"
	"github.com/phassans/exville/route"
	"github.com/rs/zerolog"
)

var (
	roach        db.Roach
	logger       zerolog.Logger
	rocketClient rocket.Client

	dbEngine      engines.DatabaseEngine
	userEngine    engines.UserEngine
	genericEngine engines.Engine

	dbHost     string
	dbPort     string
	dbUser     string
	dbPassword string
	dbDatabase string
)

func main() {
	// load ENVs
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	initEnvs()

	// configs
	config()

	// setup logger
	logger = common.GetLogger()
	logger.Info().Msg("successfully configured logger")

	// init DB
	initDB()

	// init dependencies
	initDependencies()

	// initialize engines
	dbEngine = engines.NewDatabaseEngine(roach.Db, logger)
	userEngine, err = engines.NewUserEngine(rocketClient, nil, dbEngine, logger)
	if err != nil {
		logger = logger.With().Str("error", err.Error()).Logger()
		logger.Fatal().Msgf("could not initialize userEngine")
	}
	logger.Info().Msg("engines initialized")
	genericEngine = engines.NewGenericEngine(userEngine)

	// start the server
	server = http.Server{Addr: net.JoinHostPort("", serverPort), Handler: route.APIServerHandler(genericEngine)}
	go func() { serverErrChannel <- server.ListenAndServe() }()

	// log server start time
	logger.Info().Msgf("ExVille server started at %s. time:%s", server.Addr, serverStartTime)

	// wait for any server error
	select {
	case err := <-serverErrChannel:
		logger.Fatal().Msgf("service stopped due to error %v with uptime %v", err, time.Since(serverStartTime))
	}
}

func initEnvs() {
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_DATABASE")
}

func initDB() {
	// set up DB
	var err error
	roach, err = db.New(db.Config{Host: dbHost, Port: dbPort, User: dbUser, Password: dbPassword, Database: dbDatabase})
	if err != nil {
		logger.Fatal().Msgf("could not connect to db. error %s", err)
	}
	logger.Info().Msg("successfully connected to db")
}

func initDependencies() {
	// initialize rocket client
	rocketClient = rocket.NewRocketClient(rocketURL, logger)
	logger.Info().Msg("init rocket client")
}
