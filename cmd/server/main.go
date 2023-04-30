package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/suntoryota/go-restAPI-postgres/internal/comment"
	"github.com/suntoryota/go-restAPI-postgres/internal/database"
	transportHttp "github.com/suntoryota/go-restAPI-postgres/internal/transport/http"
)

// Run - sets up our application
func Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Setting Up Our APP")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		log.Error("failed to setup connection to the database")
		return err
	}
	if err = db.MigrateDB(); err != nil {
		fmt.Println("failed to setup database")
		return err
	}

	commentService := comment.NewService(db)

	httpHandler := transportHttp.NewHandler(commentService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Error(err)
		log.Fatal("Error starting up our REST API")
	}
}
