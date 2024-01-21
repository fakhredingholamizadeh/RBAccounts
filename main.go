package main

import (
	"database/sql"
	"log"

	"github.com/flfreecode/rbaccounts/api"
	db "github.com/flfreecode/rbaccounts/db/sqlc"
	"github.com/flfreecode/rbaccounts/util"
	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("Cannot load config: %v", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("connot connect to db:", err)
	}
	log.Print("connect to db:", err)

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatalf("could not start server: %v", err)
	}

	// config, err := util.LoadConfig(".")
	// if err != nil {
	// 	log.Fatal().Err(err).Msg("cannot load config")
	// }

	// if config.Environment == "development" {
	// 	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	// }

	// connPool, err := pgxpool.New(context.Background(), config.DBSource)
	// if err != nil {
	// 	log.Fatal().Err(err).Msg("cannot connect to db")
	// }

	// runDBMigration(config.MigrationURL, config.DBSource)

	// store := db.NewStore(connPool)

	// redisOpt := asynq.RedisClientOpt{
	// 	Addr: config.RedisAddress,
	// }

	// taskDistributor := worker.NewRedisTaskDistributor(redisOpt)
	// go runTaskProcessor(config, redisOpt, store)
	// go runGatewayServer(config, store, taskDistributor)
	runGrpcServer( /*config, store, taskDistributor*/ )
}

func runGrpcServer( /*config util.Config, store db.Store, taskDistributor worker.TaskDistributor*/ ) {
	// server, err := gapi.NewServer( /*config, store, taskDistributor*/ )
	// if err != nil {
	// 	log.Fatal().Err(err).Msg("cannot create server")
	// }

	// gprcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)
	// grpcServer := grpc.NewServer(gprcLogger)
	// pb.RegisterSimpleBankServer(grpcServer, server)
	// reflection.Register(grpcServer)

	// listener, err := net.Listen("tcp", config.GRPCServerAddress)
	// if err != nil {
	// 	log.Fatal().Err(err).Msg("cannot create listener")
	// }

	// log.Info().Msgf("start gRPC server at %s", listener.Addr().String())
	// err = grpcServer.Serve(listener)
	// if err != nil {
	// 	log.Fatal().Err(err).Msg("cannot start gRPC server")
	// }
}
