package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
)

const defaultUser = "World"

func main() {
	http.HandleFunc("/", defaultHandler(getUsername()))
	http.ListenAndServe(":8080", nil)
}

func defaultHandler(user string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %s!", user)
	}
}

func getUsername() string {
	user := os.Getenv("USER")
	if user == "" {
		user = readFromConfig()
	}
	return user
}

type Config struct {
	User string `toml:"username"`
}

func readFromConfig() string {
	cfgFile := "/data/properties.ini"
	if _, err := os.Stat(cfgFile); err != nil {
		return defaultUser
	}

	var config Config
	if _, err := toml.DecodeFile(cfgFile, &config); err != nil {
		return defaultUser
	}

	return config.User
}
