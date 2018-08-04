package conf

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file is not found. Will use valirables on OS environment instead of it.")
	}
}

type Val struct {
	MQTTPort     string
	MQTTSSLPort  string
	Debug        bool
	AuthProvider string
	DefaultUser  string
	DefaultPass  string
}

func Get() *Val {
	return &Val{}
}

func getEnv(key string, required bool) string {

}
