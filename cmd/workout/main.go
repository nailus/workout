package main
import (
	//"github.com/BurntSushi/toml"
	"log"
	"github.com/nailus/workout/pkg/httpserver"
	"github.com/nailus/workout/internal/handler"
	//"fmt"
)

func main() {
	s := new(httpserver.Server)
	if err := s.Start("1000", handler.InitRouters()); err != nil {
		log.Fatal(err)
	}
}