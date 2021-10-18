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

var tempreature_lapse_rate float64 = 6.5
var reference_temperature float64 = 293.0

func Pressure(params SimParam) float64 {
	noise := perlin.NewPerlin(2, 2, 3, int64(params.Noise_seed))

	noiseFactor := noise.Noise1D(float64(params.TimeStamp.UnixMilli()) / 1000)

	var reference_height float64 = 0
	var univ_gas_const float64 = 8.314
	var gravitational_acc float64 = 9.807
	var molar_mass_earth_athmos float64 = 0.02896
	var reference_pressure float64 = noiseFactor * 101300
	var height float64 = Altitude(params)

	pressure := reference_pressure * math.Pow((reference_temperature+tempreature_lapse_rate*(height-reference_height)/reference_temperature), -gravitational_acc*molar_mass_earth_athmos/(univ_gas_const*tempreature_lapse_rate))

	return pressure
}

func Altitude(params SimParam) float64 {
	noise := perlin.NewPerlin(2, 2, 3, int64(params.Noise_seed))

	noiseFactor := noise.Noise1D(float64(params.TimeStamp.UnixMilli()) / 1000)

	altitude := params.Origin_altitude * noiseFactor

	return altitude
}

func Temperature(params SimParam) float64 {
	noise := perlin.NewPerlin(2, 2, 3, int64(params.Noise_seed))

	noiseFactor := noise.Noise1D(float64(params.TimeStamp.UnixMilli()) / 1000)

	temperature := Altitude(params) / 1000 * -tempreature_lapse_rate * reference_temperature * noiseFactor

	return temperature
}

func Latitude(params SimParam) float64 {
	noise := perlin.NewPerlin(2, 2, 3, int64(params.Noise_seed))

	noiseFactor := noise.Noise1D(float64(params.TimeStamp.UnixMilli()) / 1000)

	latitude := params.Origin_latitude * noiseFactor

	return latitude
}

func Longitude(params SimParam) float64 {
	noise := perlin.NewPerlin(2, 2, 3, int64(params.Noise_seed))

	noiseFactor := noise.Noise1D(float64(params.TimeStamp.UnixMilli()) / 1000)

	longitude := params.Origin_longitude * noiseFactor

	return longitude
}

func AirPoll() {

}

func WindSpeed(params SimParam) float64 {
	noise := perlin.NewPerlin(2, 2, 3, int64(params.Noise_seed))

	noiseFactor := noise.Noise1D(float64(params.TimeStamp.UnixMilli()) / 1000)

	var windspeed_base_factor float64 = 2
	var windspeed_increase_gradient float64 = 2

	wind_speed := windspeed_base_factor * windspeed_increase_gradient * noiseFactor

	return wind_speed
}

func Humidity() {

}

func WindDirX(params SimParam) float64 {
	noise := perlin.NewPerlin(2, 2, 3, int64(params.Noise_seed+1))

	noiseFactor := noise.Noise1D(float64(params.TimeStamp.UnixMilli()) / 1000)

	wind_speed_X := 1 * noiseFactor

	return wind_speed_X
}

func WindDirY(params SimParam) float64 {
	noise := perlin.NewPerlin(2, 2, 3, int64(params.Noise_seed+2))

	noiseFactor := noise.Noise1D(float64(params.TimeStamp.UnixMilli()) / 1000)

	wind_speed_Y := 1 * noiseFactor

	return wind_speed_Y
}
