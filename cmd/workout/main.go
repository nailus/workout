package main
import (
	_ "github.com/lib/pq"
	"log"
	"github.com/nailus/workout/pkg/httpserver"
	"github.com/nailus/workout/internal/handler"
	"github.com/nailus/workout/internal/service"
	"github.com/jmoiron/sqlx"
	"github.com/nailus/workout/internal/repository"
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