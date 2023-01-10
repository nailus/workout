package main
import (
	//"github.com/BurntSushi/toml"
	//"database/sql"
	_ "github.com/lib/pq"
	"log"
	"github.com/nailus/workout/pkg/httpserver"
	"github.com/nailus/workout/internal/handler"
	"github.com/jmoiron/sqlx"
	"github.com/nailus/workout/internal/repository"
	//"fmt"
)

func main() {
	db, err := sqlx.Open("postgres", "host=db port=5432 user=postgres dbname=workout_dev sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rep := new(repository.Repository)
	rep.Db = db

	s := new(httpserver.Server)
	if err := s.Start("1000", handler.InitRouters(rep)); err != nil {
		log.Fatal(err)
	}
}