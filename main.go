package main

import (
	"fmt"
	"log"
	"os"

	"github.com/geschke/invafetch/pkg/dbconn"
	"github.com/geschke/invafetch/pkg/invdb"
	prom "github.com/geschke/invaps/pkg/prometheus"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// LoadConfig uses the viper library to load and extract database configuration from .env file or environment variables
func LoadConfig() (dbconn.DatabaseConfiguration, string, error) {
	var config dbconn.DatabaseConfiguration
	var port string
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("/config") // for use in a Docker container

	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.SetDefault("dbPort", "3306")
	viper.SetDefault("port", "8080")

	viper.BindEnv("DBHOST")
	viper.BindEnv("DBNAME")
	viper.BindEnv("DBUSER")
	viper.BindEnv("DBPASSWORD")
	viper.BindEnv("DBPORT")
	viper.BindEnv("PORT")

	//viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			//return err
		} else {

			return config, port, fmt.Errorf("config file was found but another error ocurred: %v", err)
		}
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		return config, port, err
	}

	port = viper.Get("port").(string)

	return config, port, nil
}

func getDbRepository(dbConfig dbconn.DatabaseConfiguration) (*invdb.Repository, error) {

	var repository *invdb.Repository
	conn, err := dbconn.ConnectDB(dbConfig, 15)
	if err != nil {
		return repository, err
	}

	repository = invdb.NewRepository(conn)

	return repository, nil
}

func main() {

	dbConfig, port, err := LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	log.Printf("invaps starting on port %s...\n", port)

	invDbRepository, err := getDbRepository(dbConfig)
	if err != nil {
		log.Println("an error occurred:", err.Error())
		os.Exit(1)
	}

	prom.RecordCurrentValues(invDbRepository)

	router := gin.Default()

	router.GET("/metrics", prom.PromHandler())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Nothing here! Please use the /metrics endpoint to get data for Prometheus.",
		})
	})

	err = router.Run(":" + port)
	if err != nil {
		log.Println("an error occurred:", err.Error())
	}
}
