package main

import (
	"net"
	"net/http"
	"time"

	"github.com/phassans/exville/engines/database"

	"github.com/phassans/exville/common"
	"github.com/phassans/exville/db"
	"github.com/phassans/exville/engines"
	"github.com/phassans/exville/route"
)

func main() {
	// configs
	config()

	// setup logger
	logger := common.GetLogger()
	logger.Info().Msg("successfully configured logger")

	// set up DB
	roachDb, err := db.New(db.Config{Host: "localhost", Port: "5432", User: "pshashidhara", Password: "banana123", Database: "banana"})
	if err != nil {
		logger.Fatal().Msgf("could not connect to db. errpr %s", err)
	}

	logger.Info().Msg("successfully connected to db")

	// initialize rocket client
	/*rClient := rocket.NewRocketClient(rocketURL, logger)
	logger.Info().Msg("init rocket client")

	// initialize user engine
	userEngine, err := engines.NewUserEngine(rClient, rocketAdminUser, rocketAdminPassword, logger)
	if err != nil {
		logger = logger.With().Str("error", err.Error()).Logger()
		logger.Error().Msgf("could not initialize userEngine")
		panic("could not initialize userEngine")
	}
	logger.Info().Msg("init userEngine")

	// trying a end-point
	err = userEngine.CreateOrCheckUserChannels(nil)*/

	dbEngine := database.NewDatabaseEngine(roachDb.Db, logger)
	engines := engines.NewGenericEngine(dbEngine)

	// start the server
	server = http.Server{Addr: net.JoinHostPort("", serverPort), Handler: route.APIServerHandler(engines)}
	go func() { serverErrChannel <- server.ListenAndServe() }()

	// log server start time
	logger.Info().Msgf("ExVille server started at %s. time:%s", server.Addr, serverStartTime)

	// wait for any server error
	select {
	case err := <-serverErrChannel:
		logger.Fatal().Msgf("service stopped due to error %v with uptime %v", err, time.Since(serverStartTime))
	}
}
