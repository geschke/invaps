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

func infoDb() invdb.Repository {
	fmt.Println("Test")
	config := dbconn.ConnectDB(GetDbConfig())

	repository := invdb.NewRepository(config)
	//repository.GetProcessdata()
	return *repository
}

func main() {
	/*r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	*/

	err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	invDbRepository := infoDb()

	router := gin.Default()

	router.GET("/metrics", prom.PromHandler())

	prom.RecordCurrentValues(&invDbRepository)

	/*router.GET("/metrics", miakprom.PromHandler())

	miakprom.InitAssetGauges(stockDbRepository)
	miakprom.RecordCurrentValues(stockDbRepository)*/

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	router.GET("/val", func(c *gin.Context) {

		config := dbconn.ConnectDB(GetDbConfig())

		repository := invdb.NewRepository(config)
		values := repository.GetHomeConsumption()

		c.JSON(200, gin.H{
			"message": "val test",
			"values":  values,
		})
	})

	router.GET("/listtest", func(c *gin.Context) {
		//depotID := getDepotFromSession(c)

		//Depots := stockDbRepository.GetDepots()

		//miakprom.PrintAssetGauges()
		//miakprom.FillValuesTest()
		/*Items, _ := stockDbRepository.GetLastAssetValues()
		for _, item := range Items {
			log.Println(item)
		}*/
		c.JSON(200, gin.H{
			"message": "listtest",
		})
	})

	router.Run(":8080")
}
