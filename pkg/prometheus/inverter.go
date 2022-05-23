package invaprom

import (
	"fmt"
	"log"
	"reflect"
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

/*var home = struct {
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
*/
var devLoc = struct {
	Bat2Grid_P    prometheus.Gauge `valtype:"avg" convert:"float"`
	Dc_P          prometheus.Gauge `valtype:"avg" convert:"float"`
	DigitalIn     prometheus.Gauge `valtype:"avg" convert:"float"`
	EM_State      prometheus.Gauge `valtype:"avg" convert:"float"`
	Grid2Bat_P    prometheus.Gauge `valtype:"avg" convert:"float"`
	Grid_L1_I     prometheus.Gauge `valtype:"avg" convert:"float"`
	Grid_L1_P     prometheus.Gauge `valtype:"avg" convert:"float"`
	Grid_L2_I     prometheus.Gauge `valtype:"avg" convert:"float"`
	Grid_L2_P     prometheus.Gauge `valtype:"avg" convert:"float"`
	Grid_L3_I     prometheus.Gauge `valtype:"avg" convert:"float"`
	Grid_L3_P     prometheus.Gauge `valtype:"avg" convert:"float"`
	Grid_P        prometheus.Gauge `valtype:"avg" convert:"float"`
	Grid_Q        prometheus.Gauge `valtype:"avg" convert:"float"`
	Grid_S        prometheus.Gauge `valtype:"avg" convert:"float"`
	HomeBat_P     prometheus.Gauge `valtype:"avg" convert:"float"`
	HomeGrid_P    prometheus.Gauge `valtype:"avg" convert:"float"`
	HomeOwn_P     prometheus.Gauge `valtype:"avg" convert:"float"`
	HomePv_P      prometheus.Gauge `valtype:"avg" convert:"float"`
	Home_P        prometheus.Gauge `valtype:"avg" convert:"float"`
	InverterState prometheus.Gauge `valtype:"last" convert:"int"`
	Iso_R         prometheus.Gauge `valtype:"avg" convert:"float"`
	LimitEvuRel   prometheus.Gauge `valtype:"avg" convert:"float"`
	PV2Bat_P      prometheus.Gauge `valtype:"avg" convert:"float"`
	SinkMax_P     prometheus.Gauge `valtype:"last" convert:"float"`
	SourceMax_P   prometheus.Gauge `valtype:"last" convert:"float"`
	WorkTime      prometheus.Gauge `valtype:"last" convert:"float"`
}{
	Bat2Grid_P: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_bat2grid_p",
		Help: "Local Bat2Grid_P",
	}),
	Dc_P: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_dc_p",
		Help: "Local Dc_P",
	}),
	DigitalIn: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_digital_in",
		Help: "Local DigitalIn",
	}),
	EM_State: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_em_state",
		Help: "Local EM_State",
	}),
	Grid2Bat_P: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_grid2bat_p",
		Help: "Local Grid2Bat_P",
	}),
	Grid_L1_I: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_grid_l1_i",
		Help: "Local Grid_L1_I",
	}),
	Grid_L1_P: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_grid_l1_p",
		Help: "Local Grid_L1_P",
	}),
	Grid_L2_I: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_grid_l2_i",
		Help: "Local Grid_L2_I",
	}),
	Grid_L2_P: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_grid_l2_p",
		Help: "Local Grid_L2_P",
	}),
	Grid_L3_I: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_grid_l3_i",
		Help: "Local Grid_L3_I",
	}),
	Grid_L3_P: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_grid_l3_p",
		Help: "Local Grid_L3_P",
	}),
	Grid_P: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_grid_p",
		Help: "Local Grid_P",
	}),
	Grid_Q: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_grid_q",
		Help: "Local Grid_Q",
	}),
	Grid_S: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_grid_s",
		Help: "Local Grid_S",
	}),
	HomeBat_P: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_home_bat_p",
		Help: "Current home consumption is covered from Battery",
	}),
	HomeGrid_P: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_home_grid_p",
		Help: "Current home consumption is covered from Grid",
	}),
	HomeOwn_P: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_home_own_p",
		Help: "Home Consumption in Watt.",
	}),
	HomePv_P: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_home_pv_p",
		Help: "Current home consumption is covered from PV",
	}),
	Home_P: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_home_p",
		Help: "Current home consumption",
	}),
	InverterState: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_inverter_state",
		Help: "Local InverterState",
	}),
	Iso_R: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_iso_r",
		Help: "Local Iso_R",
	}),
	LimitEvuRel: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_limit_evu_rel",
		Help: "Local LimitEvuRel",
	}),
	PV2Bat_P: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_pv2bat_p",
		Help: "Local PV2Bat_P",
	}),
	SinkMax_P: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_sink_max_p",
		Help: "Local SinkMax_P",
	}),
	SourceMax_P: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_source_max_p",
		Help: "Local SourceMax_P",
	}),
	WorkTime: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_work_time",
		Help: "Local WorkTime",
	}),
}

var devLocBat = struct {
	BatManufacturer prometheus.Gauge `valtype:"last" convert:"float"`
	BatModel        prometheus.Gauge `valtype:"last" convert:"float"`
	BatSerialNo     prometheus.Gauge `valtype:"last" convert:"float"`
	BatVersionFW    prometheus.Gauge `valtype:"last" convert:"float"`
	Cycles          prometheus.Gauge `valtype:"last" convert:"float"`
	FullChargeCap_E prometheus.Gauge `valtype:"avg" convert:"float"`
	I               prometheus.Gauge `valtype:"avg" convert:"float"`
	P               prometheus.Gauge `valtype:"avg" convert:"float"`
	SoC             prometheus.Gauge `valtype:"avg" convert:"float"`
	U               prometheus.Gauge `valtype:"avg" convert:"float"`
	WorkCapacity    prometheus.Gauge `valtype:"avg" convert:"float"`
}{
	BatManufacturer: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_battery_bat_manufacturer",
		Help: "Battery Manufacturer",
	}),
	BatModel: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_battery_bat_model",
		Help: "Battery Model",
	}),
	BatSerialNo: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_battery_bat_serial_no",
		Help: "Battery Serial Number",
	}),
	BatVersionFW: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_battery_bat_version_fw",
		Help: "Battery Firmware Version",
	}),
	Cycles: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_battery_cycles",
		Help: "Battery Cycles",
	}),
	FullChargeCap_E: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_battery_full_charge_cap_e",
		Help: "Battery FullChargeCap_E",
	}),
	I: prometheus.NewGauge(prometheus.GaugeOpts{
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
	prometheus.MustRegister(devLoc.Bat2Grid_P)
	prometheus.MustRegister(devLoc.Dc_P)
	prometheus.MustRegister(devLoc.DigitalIn)
	prometheus.MustRegister(devLoc.EM_State)
	prometheus.MustRegister(devLoc.Grid2Bat_P)
	prometheus.MustRegister(devLoc.Grid_L1_I)
	prometheus.MustRegister(devLoc.Grid_L1_P)
	prometheus.MustRegister(devLoc.Grid_L2_I)
	prometheus.MustRegister(devLoc.Grid_L2_P)
	prometheus.MustRegister(devLoc.Grid_L3_I)
	prometheus.MustRegister(devLoc.Grid_L3_P)
	prometheus.MustRegister(devLoc.Grid_P)
	prometheus.MustRegister(devLoc.Grid_Q)
	prometheus.MustRegister(devLoc.Grid_S)
	prometheus.MustRegister(devLoc.HomeBat_P)
	prometheus.MustRegister(devLoc.HomeGrid_P)
	prometheus.MustRegister(devLoc.HomeOwn_P)
	prometheus.MustRegister(devLoc.HomePv_P)
	prometheus.MustRegister(devLoc.Home_P)
	prometheus.MustRegister(devLoc.InverterState)
	prometheus.MustRegister(devLoc.Iso_R)
	prometheus.MustRegister(devLoc.LimitEvuRel)
	prometheus.MustRegister(devLoc.PV2Bat_P)
	prometheus.MustRegister(devLoc.SinkMax_P)
	prometheus.MustRegister(devLoc.SourceMax_P)
	prometheus.MustRegister(devLoc.WorkTime)

}

// PromHandler is the main Prometheus http handler
func PromHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func fillPromValues(valueType string, valueSource, promLocation interface{}) {

	typeSrc := reflect.TypeOf(promLocation)
	valSrc := reflect.ValueOf(promLocation)
	fmt.Println("typeSrc:", typeSrc)
	fmt.Println("valSrc:", valSrc)

	valueSourceElem := reflect.ValueOf(&valueSource).Elem().Elem()

	for i := 0; i < valSrc.NumField(); i++ {
		typeSrcField := typeSrc.Field(i)
		fmt.Println("typeSrcField:", typeSrcField.Name)
		valField := valueSourceElem.FieldByName(typeSrcField.Name)
		if !valField.IsValid() {
			continue
		}

		fmt.Println("value field value:", valField)

		srcTag := typeSrc.Field(i).Tag
		fmt.Println("Tag valtype", srcTag.Get("valtype"))
		if srcTag.Get("valtype") != valueType {
			continue
		}
		fmt.Println("do convert! with ", typeSrcField.Name)
		fmt.Println("Tag convert", srcTag.Get("convert"))
		if srcTag.Get("convert") == "float" {
			fmt.Println("convert to float")

			convertedValue, err := strconv.ParseFloat(valField.String(), 64)
			if err != nil {
				continue
				//return err
			}
			fmt.Println("converted value:", convertedValue)
			//fmt.Println("canfloat?", valField.CanFloat())
			//fmt.Println("string:", valField.String())
			argv := make([]reflect.Value, 1)
			argv[0] = reflect.ValueOf(convertedValue)
			fmt.Println("reflectConverted", argv)
			valSrc.Field(i).MethodByName("Set").Call(argv)

		} else if srcTag.Get("convert") == "int" {
			fmt.Println("convert to int")
			convertedValue, err := strconv.ParseInt(valField.String(), 10, 64)
			if err != nil {
				continue
				//return err
			}
			fmt.Println("converted value:", convertedValue)
			//valSrc.Field(i).MethodByName("Set").Call(convertedValue)
			argv := make([]reflect.Value, 1)
			argv[0] = reflect.ValueOf(float64(convertedValue))
			fmt.Println("reflectConverted", argv)
			valSrc.Field(i).MethodByName("Set").Call(argv)

		} else {
			fmt.Println("don't convert")
			continue
		}
		/*		if srcTag.Get("structfield") == "nocopy" {
				continue
			}*/
		//dstField.Set(valSrc.Field(i))
	}
}

func fillCurrentFromDB(db *invdb.Repository) error {
	devLocValuesDB := db.GetDevicesLocal()
	fmt.Println(devLocValuesDB)

	fillPromValues("avg", devLocValuesDB, devLoc)

	values := db.GetDevicesLocalBattery()

	fillPromValues("avg", values, devLocBat)

	return nil
}

func fillLastFromDB(db *invdb.Repository) error {

	batteryLast := db.GetDevicesLocalBatteryLast()
	fmt.Println(batteryLast)

	fillPromValues("last", batteryLast, devLocBat)

	devLocLast := db.GetDevicesLocalLast()
	fmt.Println("devLocLast:")
	fmt.Println(devLocLast)
	fillPromValues("last", devLocLast, devLoc)

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

func CheckSomething(db *invdb.Repository) interface{} {
	fillCurrentFromDB(db)

	return devLoc
}
