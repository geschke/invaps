package invaprom

import (
	"fmt"
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

var home = struct {
	OwnP  prometheus.Gauge
	PvP   prometheus.Gauge
	P     prometheus.Gauge
	BatP  prometheus.Gauge
	GridP prometheus.Gauge
}{
	OwnP: prometheus.NewGauge(prometheus.GaugeOpts{
		//Namespace: "Home Consumption",
		Name: "home_own_p_watt",
		Help: "Home Consumption in Watt.",
	}),
	PvP: prometheus.NewGauge(prometheus.GaugeOpts{
		//Namespace: "Home Consumption",
		Name: "home_pv_p_watt",
		Help: "Current home consumption is covered from PV",
	}),
	P: prometheus.NewGauge(prometheus.GaugeOpts{
		//Namespace: "Home Consumption",
		Name: "home_p_watt",
		Help: "Current home consumption",
	}),
	BatP: prometheus.NewGauge(prometheus.GaugeOpts{
		//Namespace: "Home Consumption",
		Name: "home_bat_p_watt",
		Help: "Current home consumption is covered from Battery",
	}),
	GridP: prometheus.NewGauge(prometheus.GaugeOpts{
		//Namespace: "Home Consumption",
		Name: "home_grid_p_watt",
		Help: "Current home consumption is covered from Grid",
	}),
}

var devLocBat = struct {
	BatManufacturer prometheus.Gauge
	BatModel        prometheus.Gauge
	BatSerialNo     prometheus.Gauge
	BatVersionFW    prometheus.Gauge
	Cycles          prometheus.Gauge
	FullChargeCap_E prometheus.Gauge
	I               prometheus.Gauge
	P               prometheus.Gauge
	SoC             prometheus.Gauge
	U               prometheus.Gauge
	WorkCapacity    prometheus.Gauge
}{
	BatManufacturer: prometheus.NewGauge(prometheus.GaugeOpts{
		//Namespace: "Home Consumption",
		Name: "devices_local_battery_bat_manufacturer",
		Help: "Battery Manufacturer",
	}),
	BatModel: prometheus.NewGauge(prometheus.GaugeOpts{
		//Namespace: "Home Consumption",
		Name: "devices_local_battery_bat_model",
		Help: "Battery Model",
	}),
	BatSerialNo: prometheus.NewGauge(prometheus.GaugeOpts{
		//Namespace: "Home Consumption",
		Name: "devices_local_battery_bat_serial_no",
		Help: "Battery Serial Number",
	}),
	BatVersionFW: prometheus.NewGauge(prometheus.GaugeOpts{
		//Namespace: "Home Consumption",
		Name: "devices_local_battery_bat_version_fw",
		Help: "Battery Firmware Version",
	}),
	Cycles: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_battery_cycles",
		Help: "Battery Cycles",
	}),
	FullChargeCap_E: prometheus.NewGauge(prometheus.GaugeOpts{
		//Namespace: "Home Consumption",
		Name: "devices_local_battery_full_charge_cap_e",
		Help: "Battery FullChargeCap_E",
	}),
	I: prometheus.NewGauge(prometheus.GaugeOpts{
		//Namespace: "Home Consumption",
		Name: "devices_local_battery_i",
		Help: "Battery I",
	}),
	P: prometheus.NewGauge(prometheus.GaugeOpts{

		Name: "devices_local_battery_p",
		Help: "Battery P",
	}),
	SoC: prometheus.NewGauge(prometheus.GaugeOpts{

		Name: "devices_local_battery_soc",
		Help: "Battery SoC",
	}),
	U: prometheus.NewGauge(prometheus.GaugeOpts{

		Name: "devices_local_battery_u",
		Help: "Battery U",
	}),
	WorkCapacity: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_battery_work_capacity",
		Help: "Battery Work Capacity",
	}),
}

/*var (

	homeOwnP = prometheus.NewGauge(prometheus.GaugeOpts{
		//Namespace: "Home Consumption",
		Name: "home_own_p_watt",
		Help: "Home Consumption in Watt.",
	})
	homePvP = prometheus.NewGauge(prometheus.GaugeOpts{
		//Namespace: "Home Consumption",
		Name: "home_pv_p_watt",
		Help: "Current home consumption is covered from PV",
	})
	homeP = prometheus.NewGauge(prometheus.GaugeOpts{
		//Namespace: "Home Consumption",
		Name: "home_p_watt",
		Help: "Current home consumption",
	})
	homeBatP = prometheus.NewGauge(prometheus.GaugeOpts{
		//Namespace: "Home Consumption",
		Name: "home_bat_p_watt",
		Help: "Current home consumption is covered from Battery",
	})
	homeGridP = prometheus.NewGauge(prometheus.GaugeOpts{
		//Namespace: "Home Consumption",
		Name: "home_grid_p_watt",
		Help: "Current home consumption is covered from Grid",
	})
)
*/
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
	prometheus.MustRegister(home.OwnP)
	prometheus.MustRegister(home.P)
	prometheus.MustRegister(home.PvP)
	prometheus.MustRegister(home.BatP)
	prometheus.MustRegister(home.GridP)
	prometheus.MustRegister(devLocBat.BatManufacturer)
	prometheus.MustRegister(devLocBat.BatModel)
	prometheus.MustRegister(devLocBat.BatSerialNo)
	prometheus.MustRegister(devLocBat.BatVersionFW)
	prometheus.MustRegister(devLocBat.Cycles)
	prometheus.MustRegister(devLocBat.FullChargeCap_E)
	prometheus.MustRegister(devLocBat.I)
	prometheus.MustRegister(devLocBat.P)
	prometheus.MustRegister(devLocBat.SoC)
	prometheus.MustRegister(devLocBat.U)
	prometheus.MustRegister(devLocBat.WorkCapacity)

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
	fmt.Println(homeConsumption)

	homeOwnPValue, err := strconv.ParseFloat(homeConsumption.HomeOwnP, 64)
	if err != nil {
		return err
	}
	home.OwnP.Set(homeOwnPValue)
	homePvPValue, err := strconv.ParseFloat(homeConsumption.HomePvP, 64)
	if err != nil {
		return err
	}
	home.PvP.Set(homePvPValue)
	homePValue, err := strconv.ParseFloat(homeConsumption.HomeP, 64)
	if err != nil {
		return err
	}
	home.P.Set(homePValue)
	homeBatPValue, err := strconv.ParseFloat(homeConsumption.HomeBatP, 64)
	if err != nil {
		return err
	}
	home.BatP.Set(homeBatPValue)
	homeGridPValue, err := strconv.ParseFloat(homeConsumption.HomeGridP, 64)
	if err != nil {
		return err
	}
	home.GridP.Set(homeGridPValue)

	values := db.GetDevicesLocalBattery()
	FullChargeCap_E, err := strconv.ParseFloat(values.FullChargeCap_E, 64)
	if err != nil {
		return err
	}
	devLocBat.FullChargeCap_E.Set(FullChargeCap_E)

	I, err := strconv.ParseFloat(values.I, 64)
	if err != nil {
		return err
	}
	devLocBat.I.Set(I)

	P, err := strconv.ParseFloat(values.P, 64)
	if err != nil {
		return err
	}
	devLocBat.P.Set(P)

	SoC, err := strconv.ParseFloat(values.SoC, 64)
	if err != nil {
		return err
	}
	devLocBat.SoC.Set(SoC)

	U, err := strconv.ParseFloat(values.U, 64)
	if err != nil {
		return err
	}
	devLocBat.U.Set(U)

	WorkCapacity, err := strconv.ParseFloat(values.WorkCapacity, 64)
	if err != nil {
		return err
	}
	devLocBat.WorkCapacity.Set(WorkCapacity)

	return nil
}

func fillLastFromDB(db *invdb.Repository) error {

	batteryLast := db.GetDevicesLocalBatteryLast()
	fmt.Println(batteryLast)
	BatManufacturer, err := strconv.ParseFloat(batteryLast.BatManufacturer, 64)
	if err != nil {
		return err
	}
	devLocBat.BatManufacturer.Set(BatManufacturer)

	BatModel, err := strconv.ParseFloat(batteryLast.BatModel, 64)
	if err != nil {
		return err
	}
	devLocBat.BatModel.Set(BatModel)

	BatSerialNo, err := strconv.ParseFloat(batteryLast.BatSerialNo, 64)
	if err != nil {
		return err
	}
	devLocBat.BatSerialNo.Set(BatSerialNo)

	BatVersionFW, err := strconv.ParseFloat(batteryLast.BatVersionFW, 64)
	if err != nil {
		return err
	}
	devLocBat.BatVersionFW.Set(BatVersionFW)

	Cycles, err := strconv.ParseFloat(batteryLast.Cycles, 64)
	if err != nil {
		return err
	}
	devLocBat.Cycles.Set(Cycles)

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
	go func() {
		for {
			log.Println("in recordcurrentValues again with last values!!!")
			//FillValuesTest()
			fillLastFromDB(db)
			time.Sleep(60 * time.Second)
		}
	}()
}
