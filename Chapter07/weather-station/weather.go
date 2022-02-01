package weatherstation

import (
	"image/color"
	"strconv"
	"time"

	"tinygo.org/x/drivers/bme280"
	"tinygo.org/x/drivers/st7735"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
)

var (
	white = color.RGBA{255, 255, 255, 255}
	black = color.RGBA{0, 0, 0, 255}
)

type Service interface {
	CheckSensorConnectivity()
	ReadData() (temperature, pressure, humidity int32, err error)
	DisplayData(temperature, pressure, humidity int32)
	GetFormattedReadings(temperature, pressure, humidity int32) (temp, press, hum string)
	SavePressureReading(pressure float64)
	CheckAlert(alertThreshold float64, timeSpan int8) (bool, float64)
}

type service struct {
	sensor            *bme280.Device
	display           *st7735.Device
	readings          [6]float64
	readingsIndex     int8
	firstReadingSaved bool
}

func New(sensor *bme280.Device, display *st7735.Device) Service {
	return &service{
		sensor:            sensor,
		display:           display,
		readingsIndex:     int8(0),
		readings:          [6]float64{},
		firstReadingSaved: false,
	}
}

func (service *service) ReadData() (temp, press, hum int32, err error) {
	temp, err = service.sensor.ReadTemperature()
	if err != nil {
		return
	}

	press, err = service.sensor.ReadPressure()
	if err != nil {
		return
	}

	hum, err = service.sensor.ReadHumidity()
	if err != nil {
		return
	}

	return
}

