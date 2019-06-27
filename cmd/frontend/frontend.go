package main

import (
	"flag"
	"os"

	"github.com/emwalker/digraph/cmd/frontend/server"
	// Load the PQ drivers
	_ "github.com/lib/pq"
)

const defaultPort = "8080"

func getPlaygroundPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	return port
}

func main() {
	devMode := flag.Bool("dev", false, "Development mode")
	logLevel := flag.Int("log", 1, "Print debugging information to the console")

	flag.Parse()

	connectionString := os.Getenv("DIGRAPH_POSTGRES_CONNECTION")
	if connectionString == "" {
		panic("DIGRAPH_POSTGRES_CONNECTION not set")
	}

	redisHost := os.Getenv("DIGRAPH_REDIS_HOST")
	if redisHost == "" {
		redisHost = "localhost:6379"
	}

	s := server.New(
		getPlaygroundPort(),
		*devMode,
		os.Getenv("DIGRAPH_BASIC_AUTH_USERNAME"),
		os.Getenv("DIGRAPH_BASIC_AUTH_PASSWORD"),
		redisHost,
		os.Getenv("DIGRAPH_REDIS_PASSWORD"),
		*logLevel,
		connectionString,
	)

	s.Routes()
	s.Run()
}
