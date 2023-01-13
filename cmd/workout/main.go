package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/nailus/workout/internal/handler"
	"github.com/nailus/workout/internal/repository"
	"github.com/nailus/workout/internal/service"
	"github.com/nailus/workout/pkg/httpserver"
)

func main() {
	db, err := sqlx.Open("postgres", "host=db port=5432 user=postgres dbname=workout_dev sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	repository := repository.New(db)
	service := service.New(repository)
	handler := handler.New(service)
	server := new(httpserver.Server)
	if err := server.Start("1000", handler.InitRouters()); err != nil {
		log.Fatal(err)
	}
}