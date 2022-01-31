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

