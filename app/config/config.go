package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var (
	JWT            string
	ADMINPASSWORD  string
	GCP_CREDENTIAL string
	GCP_PROJECTID  string
	GCP_BUCKETNAME string
	GCP_PATH       string
)

type AppConfig struct {
	DBUSER     string
	DBPASSWORD string
	DBHOST     string
	DBPORT     string
	DBNAME     string
}

func InitConfig() *AppConfig {
	return readEnv()
}

func readEnv() *AppConfig {
	app := AppConfig{}
	isRead := true

	if val, found := os.LookupEnv("DBUSER"); found {
		app.DBUSER = val
		isRead = false
	}

	if val, found := os.LookupEnv("DBPASSWORD"); found {
		app.DBPASSWORD = val
		isRead = false
	}

	if val, found := os.LookupEnv("DBHOST"); found {
		app.DBHOST = val
		isRead = false
	}

	if val, found := os.LookupEnv("DBPORT"); found {
		app.DBPORT = val
		isRead = false
	}

	if val, found := os.LookupEnv("DBNAME"); found {
		app.DBNAME = val
		isRead = false
	}

	if val, found := os.LookupEnv("JWT"); found {
		JWT = val
		isRead = false
	}

	if val, found := os.LookupEnv("ADMINPASSWORD"); found {
		ADMINPASSWORD = val
		isRead = false
	}

	if val, found := os.LookupEnv("GCP_CREDENTIAL"); found {
		GCP_CREDENTIAL = val
		isRead = false
	}

	if val, found := os.LookupEnv("GCP_PROJECTID"); found {
		GCP_PROJECTID = val
		isRead = false
	}

	if val, found := os.LookupEnv("GCP_BUCKETNAME"); found {
		GCP_BUCKETNAME = val
		isRead = false
	}

	if val, found := os.LookupEnv("GCP_PATH"); found {
		GCP_PATH = val
		isRead = false
	}

	if isRead {
		viper.AddConfigPath(".")
		viper.SetConfigName("local")
		viper.SetConfigType("env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Println("error read config : ", err.Error())
			return nil
		}

		app.DBUSER = viper.GetString("DBUSER")
		app.DBPASSWORD = viper.GetString("DBPASSWORD")
		app.DBHOST = viper.GetString("DBHOST")
		app.DBPORT = viper.GetString("DBPORT")
		app.DBNAME = viper.GetString("DBNAME")
		JWT = viper.GetString("JWT")
		ADMINPASSWORD = viper.GetString("ADMINPASSWORD")
		GCP_CREDENTIAL = viper.GetString("GCP_CREDENTIAL")
		GCP_PROJECTID = viper.GetString("GCP_PROJECTID")
		GCP_BUCKETNAME = viper.GetString("GCP_BUCKETNAME")
		GCP_PATH = viper.GetString("GCP_PATH")
	}

	return &app
}
