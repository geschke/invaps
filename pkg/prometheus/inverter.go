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

var devLocAc = struct {
	CosPhi        prometheus.Gauge `valtype:"avg" convert:"float"`
	Frequency     prometheus.Gauge `valtype:"avg" convert:"float"`
	InvIn_P       prometheus.Gauge `valtype:"avg" convert:"float"`
	InvOut_P      prometheus.Gauge `valtype:"avg" convert:"float"`
	L1_I          prometheus.Gauge `valtype:"avg" convert:"float"`
	L1_P          prometheus.Gauge `valtype:"avg" convert:"float"`
	L1_U          prometheus.Gauge `valtype:"avg" convert:"float"`
	L2_I          prometheus.Gauge `valtype:"avg" convert:"float"`
	L2_P          prometheus.Gauge `valtype:"avg" convert:"float"`
	L2_U          prometheus.Gauge `valtype:"avg" convert:"float"`
	L3_I          prometheus.Gauge `valtype:"avg" convert:"float"`
	L3_P          prometheus.Gauge `valtype:"avg" convert:"float"`
	L3_U          prometheus.Gauge `valtype:"avg" convert:"float"`
	P             prometheus.Gauge `valtype:"avg" convert:"float"`
	Q             prometheus.Gauge `valtype:"avg" convert:"float"`
	ResidualCDc_I prometheus.Gauge `valtype:"avg" convert:"float"`
	S             prometheus.Gauge `valtype:"avg" convert:"float"`
}{
	CosPhi: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_ac_cos_phi",
		Help: "Local AC CosPhi",
	}),
	Frequency: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_ac_frequency",
		Help: "Local AC Frequency",
	}),
	InvIn_P: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_ac_inv_in_p",
		Help: "Local AC InvIn_P",
	}),
	InvOut_P: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_ac_inv_out_p",
		Help: "Local AC InvOut_P",
	}),
	L1_I: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_ac_l1_i",
		Help: "Local AC L1_I",
	}),
	L1_P: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_ac_l1_p",
		Help: "Local AC L1_P",
	}),
	L1_U: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_ac_l1_u",
		Help: "Local AC L1_U",
	}),
	L2_I: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_ac_l2_i",
		Help: "Local AC L2_I",
	}),
	L2_P: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_ac_l2_p",
		Help: "Local AC L2_P",
	}),
	L2_U: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_ac_l2_u",
		Help: "Local AC L2_U",
	}),
	L3_I: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_ac_l3_i",
		Help: "Local AC L3_I",
	}),
	L3_P: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_ac_l3_p",
		Help: "Local AC L3_P",
	}),
	L3_U: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_ac_l3_u",
		Help: "Local AC L3_U",
	}),
	P: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_ac_p",
		Help: "Local AC P",
	}),
	Q: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_ac_q",
		Help: "Local AC Q",
	}),
	ResidualCDc_I: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_ac_residual_cdc_i",
		Help: "Local Residual_CDc_I",
	}),
	S: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_ac_s",
		Help: "Local AC S",
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

var devLocPowermeter = struct {
	CosPhi    prometheus.Gauge `valtype:"avg" convert:"float"`
	Frequency prometheus.Gauge `valtype:"avg" convert:"float"`
	P         prometheus.Gauge `valtype:"avg" convert:"float"`
	Q         prometheus.Gauge `valtype:"avg" convert:"float"`
	S         prometheus.Gauge `valtype:"avg" convert:"float"`
}{
	CosPhi: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_powermeter_cos_phi",
		Help: "Powermeter CosPhi",
	}),
	Frequency: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_powermeter_frequency",
		Help: "Powermeter Frequency",
	}),
	P: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_powermeter_p",
		Help: "Powermeter P",
	}),
	Q: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_powermeter_q",
		Help: "Powermeter Q",
	}),
	S: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_powermeter_s",
		Help: "Powermeter S",
	}),
}

var devLocPv1 = struct {
	I prometheus.Gauge `valtype:"avg" convert:"float"`
	P prometheus.Gauge `valtype:"avg" convert:"float"`
	U prometheus.Gauge `valtype:"avg" convert:"float"`
}{
	I: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_pv1_i",
		Help: "PV1 I",
	}),
	P: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_pv1_p",
		Help: "PV1 P",
	}),
	U: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_pv1_u",
		Help: "PV1 U",
	}),
}

var devLocPv2 = struct {
	I prometheus.Gauge `valtype:"avg" convert:"float"`
	P prometheus.Gauge `valtype:"avg" convert:"float"`
	U prometheus.Gauge `valtype:"avg" convert:"float"`
}{
	I: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_pv2_i",
		Help: "PV2 I",
	}),
	P: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_pv2_p",
		Help: "PV2 P",
	}),
	U: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "devices_local_pv2_u",
		Help: "PV2 U",
	}),
}

var scbStatisticEnergyFlow = struct {
	StatisticAutarkyDay               prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticAutarkyMonth             prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticAutarkyTotal             prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticAutarkyYear              prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticCO2SavingDay             prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticCO2SavingMonth           prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticCO2SavingTotal           prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticCO2SavingYear            prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyChargeGridDay      prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyChargeGridMonth    prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyChargeGridTotal    prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyChargeGridYear     prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyChargeInvInDay     prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyChargeInvInMonth   prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyChargeInvInTotal   prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyChargeInvInYear    prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyChargePvDay        prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyChargePvMonth      prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyChargePvTotal      prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyChargePvYear       prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyDischargeDay       prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyDischargeMonth     prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyDischargeTotal     prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyDischargeYear      prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyDischargeGridDay   prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyDischargeGridMonth prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyDischargeGridTotal prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyDischargeGridYear  prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyHomeDay            prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyHomeMonth          prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyHomeTotal          prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyHomeYear           prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyHomeBatDay         prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyHomeBatMonth       prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyHomeBatTotal       prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyHomeBatYear        prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyHomeGridDay        prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyHomeGridMonth      prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyHomeGridTotal      prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyHomeGridYear       prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyHomeOwnTotal       prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyHomePvDay          prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyHomePvMonth        prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyHomePvTotal        prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyHomePvYear         prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyPv1Day             prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyPv1Month           prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyPv1Total           prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyPv1Year            prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyPv2Day             prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyPv2Month           prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyPv2Total           prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyPv2Year            prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyPv3Day             prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyPv3Month           prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyPv3Total           prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticEnergyPv3Year            prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticOwnConsumptionRateDay    prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticOwnConsumptionRateMonth  prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticOwnConsumptionRateTotal  prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticOwnConsumptionRateYear   prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticYieldDay                 prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticYieldMonth               prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticYieldTotal               prometheus.Gauge `valtype:"last" convert:"float"`
	StatisticYieldYear                prometheus.Gauge `valtype:"last" convert:"float"`
}{
	StatisticAutarkyDay: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_autarky_day",
		Help: "Statistic Autarky Day",
	}),
	StatisticAutarkyMonth: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_autarky_month",
		Help: "Statistic Autarky Month",
	}),
	StatisticAutarkyTotal: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_autarky_total",
		Help: "Statistic Autarky Total",
	}),
	StatisticAutarkyYear: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_autarky_year",
		Help: "Statistic Autarky Year",
	}),
	StatisticCO2SavingDay: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_co2saving_day",
		Help: "Statistic CO2Saving Day",
	}),
	StatisticCO2SavingMonth: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_co2saving_month",
		Help: "Statistic CO2Saving Month",
	}),
	StatisticCO2SavingTotal: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_co2saving_total",
		Help: "Statistic CO2Saving Total",
	}),
	StatisticCO2SavingYear: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_co2saving_year",
		Help: "Statistic CO2Saving Year",
	}),
	StatisticEnergyChargeGridDay: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_charge_grid_day",
		Help: "Statistic EnergyChargeGrid Day",
	}),
	StatisticEnergyChargeGridMonth: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_charge_grid_month",
		Help: "Statistic EnergyChargeGrid Month",
	}),
	StatisticEnergyChargeGridTotal: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_charge_grid_total",
		Help: "Statistic EnergyChargeGrid Total",
	}),
	StatisticEnergyChargeGridYear: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_charge_grid_year",
		Help: "Statistic EnergyChargeGrid Year",
	}),
	StatisticEnergyChargeInvInDay: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_charge_inv_in_day",
		Help: "Statistic EnergyChargeInvIn Day",
	}),
	StatisticEnergyChargeInvInMonth: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_charge_inv_in_month",
		Help: "Statistic EnergyChargeInvIn Month",
	}),
	StatisticEnergyChargeInvInTotal: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_charge_inv_in_total",
		Help: "Statistic EnergyChargeInvIn Total",
	}),
	StatisticEnergyChargeInvInYear: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_charge_inv_in_year",
		Help: "Statistic EnergyChargeInvIn Year",
	}),
	StatisticEnergyChargePvDay: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_charge_pv_day",
		Help: "Statistic EnergyChargePv Day",
	}),
	StatisticEnergyChargePvMonth: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_charge_pv_month",
		Help: "Statistic EnergyChargePv Month",
	}),
	StatisticEnergyChargePvTotal: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_charge_pv_total",
		Help: "Statistic EnergyChargePv Total",
	}),
	StatisticEnergyChargePvYear: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_charge_pv_year",
		Help: "Statistic EnergyChargePv Year",
	}),
	StatisticEnergyDischargeDay: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_discharge_day",
		Help: "Statistic EnergyDischarge Day",
	}),
	StatisticEnergyDischargeMonth: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_discharge_month",
		Help: "Statistic EnergyDischarge Month",
	}),
	StatisticEnergyDischargeTotal: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_discharge_total",
		Help: "Statistic EnergyDischarge Total",
	}),
	StatisticEnergyDischargeYear: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_discharge_year",
		Help: "Statistic EnergyDischarge Year",
	}),
	StatisticEnergyDischargeGridDay: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_discharge_grid_day",
		Help: "Statistic EnergyDischargeGrid Day",
	}),
	StatisticEnergyDischargeGridMonth: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_discharge_grid_month",
		Help: "Statistic EnergyDischargeGrid Month",
	}),
	StatisticEnergyDischargeGridTotal: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_discharge_grid_total",
		Help: "Statistic EnergyDischargeGrid Total",
	}),
	StatisticEnergyDischargeGridYear: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_discharge_grid_year",
		Help: "Statistic EnergyDischargeGrid Year",
	}),
	StatisticEnergyHomeDay: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_home_day",
		Help: "Statistic EnergyHome Day",
	}),
	StatisticEnergyHomeMonth: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_home_month",
		Help: "Statistic EnergyHome Month",
	}),
	StatisticEnergyHomeTotal: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_home_total",
		Help: "Statistic EnergyHome Total",
	}),
	StatisticEnergyHomeYear: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_home_year",
		Help: "Statistic EnergyHome Year",
	}),
	StatisticEnergyHomeBatDay: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_home_bat_day",
		Help: "Statistic EnergyHomeBat Day",
	}),
	StatisticEnergyHomeBatMonth: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_home_bat_month",
		Help: "Statistic EnergyHomeBat Month",
	}),
	StatisticEnergyHomeBatTotal: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_home_bat_total",
		Help: "Statistic EnergyHomeBat Total",
	}),
	StatisticEnergyHomeBatYear: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_home_bat_year",
		Help: "Statistic EnergyHomeBat Year",
	}),
	StatisticEnergyHomeGridDay: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_home_grid_day",
		Help: "Statistic EnergyHomeGrid Day",
	}),
	StatisticEnergyHomeGridMonth: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_home_grid_month",
		Help: "Statistic EnergyHomeGrid Month",
	}),
	StatisticEnergyHomeGridTotal: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_home_grid_total",
		Help: "Statistic EnergyHomeGrid Total",
	}),
	StatisticEnergyHomeGridYear: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_home_grid_year",
		Help: "Statistic EnergyHomeGrid Year",
	}),
	StatisticEnergyHomeOwnTotal: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_home_own_total",
		Help: "Statistic EnergyHomeOwn Total",
	}),
	StatisticEnergyHomePvDay: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_home_pv_day",
		Help: "Statistic EnergyHomePv Day",
	}),
	StatisticEnergyHomePvMonth: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_home_pv_month",
		Help: "Statistic EnergyHomePv Month",
	}),
	StatisticEnergyHomePvTotal: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_home_pv_total",
		Help: "Statistic EnergyHomePv Total",
	}),
	StatisticEnergyHomePvYear: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_home_pv_year",
		Help: "Statistic EnergyHomePv Year",
	}),
	StatisticEnergyPv1Day: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_pv1_day",
		Help: "Statistic EnergyPv1 Day",
	}),
	StatisticEnergyPv1Month: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_pv1_month",
		Help: "Statistic EnergyPv1 Month",
	}),
	StatisticEnergyPv1Total: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_pv1_total",
		Help: "Statistic EnergyPv1 Total",
	}),
	StatisticEnergyPv1Year: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_pv1_year",
		Help: "Statistic EnergyPv1 Year",
	}),
	StatisticEnergyPv2Day: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_pv2_day",
		Help: "Statistic EnergyPv2 Day",
	}),
	StatisticEnergyPv2Month: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_pv2_month",
		Help: "Statistic EnergyPv2 Month",
	}),
	StatisticEnergyPv2Total: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_pv2_total",
		Help: "Statistic EnergyPv2 Total",
	}),
	StatisticEnergyPv2Year: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_pv2_year",
		Help: "Statistic EnergyPv2 Year",
	}),
	StatisticEnergyPv3Day: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_pv3_day",
		Help: "Statistic EnergyPv3 Day",
	}),
	StatisticEnergyPv3Month: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_pv3_month",
		Help: "Statistic EnergyPv3 Month",
	}),
	StatisticEnergyPv3Total: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_pv3_total",
		Help: "Statistic EnergyPv3 Total",
	}),
	StatisticEnergyPv3Year: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_energy_pv3_year",
		Help: "Statistic EnergyPv3 Year",
	}),

	StatisticOwnConsumptionRateDay: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_own_consumption_rate_day",
		Help: "Statistic OwnConsumptionRate Day",
	}),
	StatisticOwnConsumptionRateMonth: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_own_consumption_rate_month",
		Help: "Statistic OwnConsumptionRate Month",
	}),
	StatisticOwnConsumptionRateTotal: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_own_consumption_rate_total",
		Help: "Statistic OwnConsumptionRate Total",
	}),
	StatisticOwnConsumptionRateYear: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_own_consumption_rate_year",
		Help: "Statistic OwnConsumptionRate Year",
	}),

	StatisticYieldDay: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_yield_day",
		Help: "Statistic Yield Day",
	}),
	StatisticYieldMonth: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_yield_month",
		Help: "Statistic Yield Month",
	}),
	StatisticYieldTotal: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_yield_total",
		Help: "Statistic Yield Total",
	}),
	StatisticYieldYear: prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "statistic_yield_year",
		Help: "Statistic Yield Year",
	}),
}

func init() {

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

	prometheus.MustRegister(devLocAc.CosPhi)
	prometheus.MustRegister(devLocAc.Frequency)
	prometheus.MustRegister(devLocAc.InvIn_P)
	prometheus.MustRegister(devLocAc.InvOut_P)
	prometheus.MustRegister(devLocAc.L1_I)
	prometheus.MustRegister(devLocAc.L1_P)
	prometheus.MustRegister(devLocAc.L1_U)
	prometheus.MustRegister(devLocAc.L2_I)
	prometheus.MustRegister(devLocAc.L2_P)
	prometheus.MustRegister(devLocAc.L2_U)
	prometheus.MustRegister(devLocAc.L3_I)
	prometheus.MustRegister(devLocAc.L3_P)
	prometheus.MustRegister(devLocAc.L3_U)
	prometheus.MustRegister(devLocAc.P)
	prometheus.MustRegister(devLocAc.Q)
	prometheus.MustRegister(devLocAc.ResidualCDc_I)
	prometheus.MustRegister(devLocAc.S)

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

	prometheus.MustRegister(devLocPowermeter.CosPhi)
	prometheus.MustRegister(devLocPowermeter.Frequency)
	prometheus.MustRegister(devLocPowermeter.P)
	prometheus.MustRegister(devLocPowermeter.Q)
	prometheus.MustRegister(devLocPowermeter.S)

	prometheus.MustRegister(devLocPv1.I)
	prometheus.MustRegister(devLocPv1.P)
	prometheus.MustRegister(devLocPv1.U)

	prometheus.MustRegister(devLocPv2.I)
	prometheus.MustRegister(devLocPv2.P)
	prometheus.MustRegister(devLocPv2.U)

	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticAutarkyDay)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticAutarkyMonth)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticAutarkyTotal)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticAutarkyYear)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticCO2SavingDay)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticCO2SavingMonth)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticCO2SavingTotal)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticCO2SavingYear)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyChargeGridDay)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyChargeGridMonth)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyChargeGridTotal)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyChargeGridYear)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyChargeInvInDay)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyChargeInvInMonth)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyChargeInvInTotal)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyChargeInvInYear)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyChargePvDay)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyChargePvMonth)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyChargePvTotal)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyChargePvYear)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyDischargeDay)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyDischargeMonth)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyDischargeTotal)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyDischargeYear)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyDischargeGridDay)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyDischargeGridMonth)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyDischargeGridTotal)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyDischargeGridYear)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyHomeDay)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyHomeMonth)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyHomeTotal)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyHomeYear)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyHomeBatDay)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyHomeBatMonth)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyHomeBatTotal)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyHomeBatYear)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyHomeGridDay)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyHomeGridMonth)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyHomeGridTotal)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyHomeGridYear)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyHomeOwnTotal)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyHomePvDay)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyHomePvMonth)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyHomePvTotal)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyHomePvYear)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyPv1Day)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyPv1Month)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyPv1Total)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyPv1Year)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyPv2Day)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyPv2Month)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyPv2Total)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyPv2Year)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyPv3Day)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyPv3Month)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyPv3Total)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticEnergyPv3Year)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticOwnConsumptionRateDay)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticOwnConsumptionRateMonth)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticOwnConsumptionRateTotal)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticOwnConsumptionRateYear)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticYieldDay)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticYieldMonth)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticYieldTotal)
	prometheus.MustRegister(scbStatisticEnergyFlow.StatisticYieldYear)

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

	devLocAcValuesDB := db.GetDevicesLocalAc()
	fillPromValues("avg", devLocAcValuesDB, devLocAc)

	devLocPowermeterDB := db.GetDevicesLocalPowermeter()
	fillPromValues("avg", devLocPowermeterDB, devLocPowermeter)

	devLocPv1DB := db.GetDevicesLocalPv1()
	fillPromValues("avg", devLocPv1DB, devLocPv1)

	devLocPv2DB := db.GetDevicesLocalPv2()
	fillPromValues("avg", devLocPv2DB, devLocPv2)

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

	statisticsEnergyFlow := db.GetStatisticEnergyFlowLast()
	fmt.Println(statisticsEnergyFlow)
	fillPromValues("last", statisticsEnergyFlow, scbStatisticEnergyFlow)

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
	go func() {
		for {
			log.Println("in recordcurrentValues again Remove Data!!!")
			//FillValuesTest()
			db.RemoveData(2)
			time.Sleep(24 * time.Hour)
		}
	}()
}

func CheckSomething(db *invdb.Repository) interface{} {
	//fillCurrentFromDB(db)
	db.RemoveData(2)

	return devLoc
}
