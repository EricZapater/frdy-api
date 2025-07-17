package main

import (
	"frdy-api/config"
	"frdy-api/server"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := config.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer db.Close()

	srv := server.NewServer(cfg, db)
    
    // Configurar middlewares i rutes
    if err := srv.Setup(); err != nil {
        log.Fatalf("failed to set up middlewares: %v", err)
    }
    
    // Iniciar el servidor
    if err := srv.Run(); err != nil {
        log.Fatalf("failed to start server: %v", err)
    }
}