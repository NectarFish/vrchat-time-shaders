package main

import (
	"io"
	"log"
	"strconv"
	"time"

	"image"
	"image/color"
	"image/draw"
	"image/png"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.GET("/time", func(ctx *gin.Context) {
		_timezoneStr := ctx.Query("timezone")
		timezone := func() int {
			_timezone, err := strconv.Atoi(_timezoneStr)
			if err == nil {
				return _timezone
			}
			return 9
		}()
		timeBuilder(ctx.Writer, timezone)
	})

	log.Fatal(r.Run(":4305"))
}

func timeBuilder(w io.Writer, timezone int) {
	rectImage := image.NewRGBA(image.Rect(0, 0, 64, 64))
	draw.Draw(rectImage, rectImage.Bounds(), &image.Uniform{
		color.RGBA{0, 0, 0, 255},
	}, image.ZP, draw.Src)

	now := time.Now().In(time.FixedZone("offsetTime", int((time.Duration(timezone) * time.Hour).Seconds())))
	year := now.Year() - 1900
	month := int(now.Month()) - 1
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	ms := int(now.UnixNano() / 1e6)
	weekday := now.Weekday()
	moonAge := (((now.Year()-2009)%19)*11 + (int(now.Month()) + 1) + (now.Day() + 1)) % 30

	drawCall(rectImage, 0, 0, hour&0b111)
	drawCall(rectImage, 1, 0, hour>>3)
	drawCall(rectImage, 2, 0, minute&0b111)
	drawCall(rectImage, 3, 0, minute>>3)
	drawCall(rectImage, 4, 0, second&0b111)
	drawCall(rectImage, 5, 0, second>>3)
	drawCall(rectImage, 6, 0, ms&0b111)
	drawCall(rectImage, 7, 0, ms>>3)

	drawCall(rectImage, 0, 1, year&0b111)
	drawCall(rectImage, 1, 1, (year>>3)&0b111)
	drawCall(rectImage, 2, 1, (year>>6)&0b111)
	drawCall(rectImage, 3, 1, month&0b111)
	drawCall(rectImage, 4, 1, month>>3)
	drawCall(rectImage, 5, 1, day&0b111)
	drawCall(rectImage, 6, 1, day>>3)
	drawCall(rectImage, 7, 1, int(weekday))

	drawCall(rectImage, 0, 2, moonAge&0b111)
	drawCall(rectImage, 1, 2, moonAge>>3)

	_ = png.Encode(w, rectImage)
}

func drawCall(rectImage *image.RGBA, x, y, val int) {
	x0 := x * 8
	y0 := y * 8
	x1 := (x + 1) * 8
	y1 := (y + 1) * 8
	f := func(o uint) int {
		if val&(1<<o) != 0 {
			return 255
		}
		return 0
	}
	r := f(0)
	g := f(1)
	b := f(2)

	draw.Draw(rectImage,
		image.Rect(x0, y0, x1, y1),
		&image.Uniform{
			color.RGBA{uint8(r), uint8(g), uint8(b), 255},
		}, image.ZP, draw.Src,
	)
}
