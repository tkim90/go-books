package main

import (

)

type server struct {
	apiHandler http.Handler
	listenPort int
	serv       *http.Server
	config     *serverConfig
}

type serverConfig struct {
	ListenPort int
}

func newServer(config *serverConfig) (*server, error) {
  // TODO db migration

  apiHandler := api.New(&api.Config{

  })
} 