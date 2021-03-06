package main

import (
	"flag"

	"github.com/nytimes/gizmo/config"
	"github.com/nytimes/gizmo/pubsub"
	"github.com/nytimes/gizmo/server"

	"github.com/nytimes/gizmo/examples/pubsub/api-kafka-websocket-pubsub/service"
)

func main() {
	cfg := config.NewConfig("./config.json")

	// set the pubsub's Log to be the same as server's
	pubsub.Log = server.Log

	// in case we want to override the port or log location via CLI
	flag.Parse()
	config.SetServerOverrides(cfg.Server)

	server.Init("gamestream-example", cfg.Server)

	err := server.Register(service.NewStreamService(cfg.Server.HTTPPort, cfg.Kafka))
	if err != nil {
		server.Log.Fatal(err)
	}

	if err = server.Run(); err != nil {
		server.Log.Fatal(err)
	}
}
