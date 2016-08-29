package gokml_gx

import (
	"fmt"

	colourful "github.com/lucasb-eyer/go-colorful"
)

type Color struct {
	rgb colourful.Color
	a   float64
}

func NewColorRGBA(r, g, b, a float64) *Color {
	var color Color

	color.rgb = colourful.Color{r, g, b}
	color.a = a

	return &color
}

func NewColorHSLA(h, s, l, a float64) *Color {
	var color Color
	color.rgb = colourful.Hsl(h, s, l)
	color.a = a
	return &color
}

func (c *Color) Marshal() ([]byte, error) {
	r, g, b := c.rgb.RGB255()
	a := int(c.a * 255)
	return []byte(fmt.Sprintf("#%02x%02x%02x%02x", a, b, g, r)), nil
}

func (c *Color) Code() string {
	r, g, b := c.rgb.RGB255()
	a := int(c.a * 255)
	return fmt.Sprintf("#%02x%02x%02x%02x", a, b, g, r)

}
