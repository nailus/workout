package main
import (
	"github.com/BurntSushi/toml"
	"log"
	"github.com/nailus/workout/internal/app/apiserver"
	//"fmt"
)

func main() {
	config := apiserver.NewConfig()
	
	_, err := toml.DecodeFile("config/apiserver.toml", config)

	//fmt.Println(config.Database)


	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}