package miakprom

import (
	"log"
	"strconv"
	"time"

	"github.com/geschke/invafetch/pkg/invdb"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// AssetValueTest is a definition of Prometheus Gauge
/*var AssetValueGlobal = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "asset_value_euro",
	Help: "Asset values",
})*/

var (
	homeOwnP = prometheus.NewGauge(prometheus.GaugeOpts{
		//Namespace: "Home Consumption",
		Name: "home_own_p_watt",
		Help: "Home Consumption in Watt.",
	})
)

/*
// AssetGaugeGlobal is the Gauge vector to be filled with asset values
var AssetGaugeGlobal = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		//Namespace: "our_company",
		//Subsystem: "blob_storage",
		Name: "Home Consumasset_value_euro",
		Help: "Asset Values",
	},
	[]string{
		// Our depot
		"depot",
		"isin",
		// Asset type (share, fund, etf)
		"asset_type",
		"asset_name",
	},
)

type assetGaugeOptsType struct {
	Name      string
	Help      string
	AssetType string
	//DepotName string
	//AssetGauge prometheus.Gauge
	AssetGauge *prometheus.GaugeVec
}
*/
/*
var assetGauge []prometheus.Gauge
var assetMap = make(map[string]assetGaugeOptsType)
*/
func init() {
	prometheus.MustRegister(homeOwnP)
}

// PromHandler is the main Prometheus http handler
func PromHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func fillCurrentFromDB(db *invdb.Repository) error {
	homeConsumption := db.GetHomeConsumption()

	homeOwnPValue, err := strconv.ParseFloat(homeConsumption.HomeOwnP, 64)
	if err != nil {
		return err
	}
	homeOwnP.Set(homeOwnPValue)
	return nil
}

// RecordCurrentValues fills Prometheus data structure with new test values
func RecordCurrentValues(db *invdb.Repository) {
	go func() {
		for {
			log.Println("in recordCurrentValues again!!!")
			//FillValuesTest()
			fillCurrentFromDB(db)
			time.Sleep(30 * time.Second)
		}
	}()
}
