package main

import (
	"fmt"
	"log"

	"github.com/geschke/invafetch/pkg/dbconn"
	"github.com/geschke/invafetch/pkg/invdb"
	prom "github.com/geschke/invaps/pkg/prometheus"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// LoadConfig uses the viper library to load and extract database configuration from .env file or environment variables
func LoadConfig(path string) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	//err = viper.Unmarshal(&config)
	return
}

func GetDbConfig() dbconn.DatabaseConfiguration {
	// could something go wrong here?
	fmt.Println(viper.Get("dbName"))
	fmt.Println(viper.Get("dbHost"))
	fmt.Println(viper.Get("dbUser"))
	fmt.Println(viper.Get("dbPassword"))
	fmt.Println(viper.Get("dbPort"))
	var config dbconn.DatabaseConfiguration
	config.DBHost = viper.GetString("dbHost")
	config.DBName = viper.GetString("dbName")
	config.DBPassword = viper.GetString("dbPassword")
	config.DBUser = viper.GetString("dbUser")
	config.DBPort = viper.GetString("dbPort")
	return config
}

func getDbRepository() invdb.Repository {

	config := dbconn.ConnectDB(GetDbConfig())

	repository := invdb.NewRepository(config)

	return *repository
}

func main() {

	err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	invDbRepository := getDbRepository()

	router := gin.Default()

	router.GET("/metrics", prom.PromHandler())

	prom.RecordCurrentValues(&invDbRepository)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Nothing here! Please use the /metrics endpoint to get data for Prometheus.",
		})
	})

	router.Run(":8080")
}
