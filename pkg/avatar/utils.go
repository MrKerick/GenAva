package avatar

import (
	"encoding/json"
	"fmt"
	"image/color"
	"os"
)

type ColorItem struct {
	Name string `json:"name"`
	Hex  string `json:"hex"`
}

func loadColors(path string) ([]ColorItem, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var colors []ColorItem

	err = json.Unmarshal(data, &colors)

	return colors, err
}

func hexToRGBA(hex string) (color.RGBA, error) {

	var r, g, b uint8

	_, err := fmt.Sscanf(
		hex,
		"#%02x%02x%02x",
		&r,
		&g,
		&b,
	)

	if err != nil {
		return color.RGBA{}, err
	}

	return color.RGBA{
		R: r,
		G: g,
		B: b,
		A: 255,
	}, nil
}

func contrastColor(c color.RGBA) color.RGBA {

	brightness :=
		(299*int(c.R) +
			587*int(c.G) +
			114*int(c.B)) / 1000

	if brightness > 128 {
		return color.RGBA{0, 0, 0, 255}
	}

	return color.RGBA{255, 255, 255, 255}
}

func truncate(text string, limit int) string {

	if len(text) <= limit {
		return text
	}

	return text[:limit] + "..."
}