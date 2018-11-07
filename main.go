package main

import (
	"net"
	"net/http"
	"time"

	"github.com/phassans/exville/clients/rocket"
	"github.com/phassans/exville/common"
	"github.com/phassans/exville/engines"
	"github.com/phassans/exville/route"
)

func main() {
	// configs
	config()

	// setup logger
	logger := common.GetLogger()
	logger.Info().Msg("successfully configured logger")

	// initialize rocket client
	rClient := rocket.NewRocketClient(rocketURL, logger)
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
	err = userEngine.CreateOrCheckUserChannels(nil)

	// start the server
	server = http.Server{Addr: net.JoinHostPort("", serverPort), Handler: route.APIServerHandler()}
	go func() { serverErrChannel <- server.ListenAndServe() }()

	// log server start time
	logger.Info().Msgf("ExVille server started at %s. time:%s", server.Addr, serverStartTime)

	// wait for any server error
	select {
	case err := <-serverErrChannel:
		logger.Fatal().Msgf("service stopped due to error %v with uptime %v", err, time.Since(serverStartTime))
	}
}
