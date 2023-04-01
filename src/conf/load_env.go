package conf

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Env = &env{}

type env struct {
	AppMode string
	GinMode string
	Host    string
	Port    string
}

var IsLoaded = false

func LoadEnv() {
	// httpclients and Persistence layer needs the env vars
	// but we don't want to load them twice
	if IsLoaded {
		return
	}

	// if not "prod"
	if Env.AppMode != "prod" {
		curDir, err := os.Getwd()
		if err != nil {
			log.Fatal(err, "conf", "LoadEnv", "error loading os.Getwd()")
		}
		// load the /.env file
		loadErr := godotenv.Load(curDir + "/.env")
		if loadErr != nil {
			log.Fatal(loadErr, "conf", "LoadEnv", "can't load env file from current directory: "+curDir)
		}
		Env.GinMode = "debug"
	} else {
		Env.GinMode = "release"
	}

	// load the env vars
	Env.AppMode = os.Getenv("APP_MODE")
	Env.Host = os.Getenv("HOST")
	Env.Port = os.Getenv("PORT")

	IsLoaded = true
}
