package conf

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func init() { loadEnv() }

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file is not found. Will use valirables on OS environment instead of it.")
	}
}

func ReLoad() { loadEnv() }

type Val struct {
	MQTTPort     string
	MQTTSSLPort  string
	Debug        bool
	AuthProvider string
	DefaultUser  string
	DefaultPass  string
}

func Get() *Val {
	return &Val{
		MQTTPort:     getEnv("MQTT_PORT", true),
		MQTTSSLPort:  getEnv("MQTT_SSL_PORT", true),
		Debug:        toBool(getEnv("DEBUG", true)),
		AuthProvider: getEnv("AUTH_PROVIDER", true),
		DefaultUser:  getEnv("MQTT_DEFAULT_USER", true),
		DefaultPass:  getEnv("MQTT_DEFAULT_PASS", true),
	}
}

func getEnv(key string, required bool) string {
	v := os.Getenv(key)
	if len(v) == 0 {
		panic(fmt.Sprintf("%s key is defined, it is reqired!", key))
	}
	return v
}

func toBool(value string) bool {
	b, err := strconv.ParseBool(value)
	if err != nil {
		panic(err)
	}
	return b
}
