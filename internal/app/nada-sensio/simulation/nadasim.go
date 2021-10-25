package nadasensio

import (
	"math"
	"time"

	perlin "github.com/aquilax/go-perlin"
)

type SimParam struct {
	Noise_seed       int
	Origin_latitude  float64
	Origin_longitude float64
	Origin_altitude  float64
	TimeStamp        time.Time
}

var temperature_lapse_rate float64 = 9.8
var reference_temperature float64 = 293.0

func getNoiseFactor(params SimParam) float64 {
	noise := perlin.NewPerlin(1, 10000, 3, int64(params.Noise_seed))
	timerange := 100 * 365 * 24 * 3600
	var timeseconds float64 = float64(params.TimeStamp.UnixMilli()) / 1000.0
	x1 := float64(timeseconds)
	x := x1 / float64(timerange)
	noiseFactor := noise.Noise1D(x)

	return noiseFactor
}

func Pressure(params SimParam) float64 {
	noiseFactor := getNoiseFactor(params)

	var univ_gas_const float64 = 8.314
	var gravitational_acc float64 = 9.807
	var molar_mass_earth_athmos float64 = 0.02896
	var reference_pressure float64 = noiseFactor*10 + 101300
	var height float64 = Altitude(params)
	var reference_temperature float64 = 288

	pressure := reference_pressure * math.Exp(-(gravitational_acc*height*molar_mass_earth_athmos)/(reference_temperature*univ_gas_const))

	return pressure
}

func Altitude(params SimParam) float64 {
	noiseFactor := getNoiseFactor(params)

	noiseScale := 50

	altitude := params.Origin_altitude + noiseFactor*float64(noiseScale)

	return altitude
}

func Temperature(params SimParam) float64 {
	noiseFactor := getNoiseFactor(params)

	var temperature float64 = reference_temperature - 273 - (Altitude(params)/1000)*temperature_lapse_rate + 5*noiseFactor
	return temperature
}

func Latitude(params SimParam) float64 {
	noiseFactor := getNoiseFactor(params)

	latitude := params.Origin_latitude + 20*noiseFactor

	return latitude
}

func Longitude(params SimParam) float64 {
	noiseFactor := getNoiseFactor(params)

	longitude := params.Origin_longitude + 20*noiseFactor

	return longitude
}

func AirPoll() {

}

func WindSpeed(params SimParam) float64 {
	noiseFactor := getNoiseFactor(params)

	var windspeed_base_factor float64 = 2
	var windspeed_increase_gradient float64 = 2
	var altitude float64 = Altitude(params)

	wind_speed := 20 + altitude/1000*windspeed_base_factor*windspeed_increase_gradient + noiseFactor*20

	return wind_speed
}

func Humidity() {

}

func WindDirX(params SimParam) float64 {
	noiseFactor := getNoiseFactor(params)

	wind_speed_X := 1 * noiseFactor

	return wind_speed_X
}

func WindDirY(params SimParam) float64 {
	noiseFactor := getNoiseFactor(params)

	wind_speed_Y := 1 * noiseFactor

	return wind_speed_Y
}
