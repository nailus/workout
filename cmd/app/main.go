package main
import (
	"github.com/BurntSushi/toml"
	"flag"
	"log"
	"github.com/nailus/workout/internal/app/apiserver"
	//"fmt"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/apiserver.toml", "config path")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}