package avatar

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

type Result struct {
	FilePath string
	HexColor string
	Email    string
}

type Generator struct {
	Config Config
}

func New(config Config) *Generator {
	return &Generator{
		Config: config,
	}
}

func (g *Generator) Generate(
	email string,
	filename string,
) (*Result, error) {

	colors, err :=
		loadColors(g.Config.ColorsPath)

	if err != nil {
		return nil, err
	}

	selected :=
		colors[rand.Intn(len(colors))]

	bgColor, err :=
		hexToRGBA(selected.Hex)

	if err != nil {
		return nil, err
	}

	img := image.NewRGBA(
		image.Rect(
			0,
			0,
			g.Config.Width,
			g.Config.Height,
		),
	)

	draw.Draw(
		img,
		img.Bounds(),
		&image.Uniform{bgColor},
		image.Point{},
		draw.Src,
	)

	letter :=
		strings.ToUpper(string(email[0]))

	fontColor :=
		contrastColor(bgColor)

	err = g.drawLetter(
		img,
		letter,
		fontColor,
	)

	if err != nil {
		return nil, err
	}

	err = os.MkdirAll(
		g.Config.OutputDir,
		os.ModePerm,
	)

	if err != nil {
		return nil, err
	}

	fullPath :=
		filepath.Join(
			g.Config.OutputDir,
			filename,
		)

	file, err := os.Create(fullPath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	err = png.Encode(file, img)

	if err != nil {
		return nil, err
	}

	return &Result{
		FilePath: fullPath,
		HexColor: selected.Hex,
		Email: truncate(
			email,
			g.Config.TruncateLength,
		),
	}, nil
}

func (g *Generator) drawLetter(
	img *image.RGBA,
	text string,
	clr color.Color,
) error {

	fontBytes, err :=
		os.ReadFile(g.Config.FontPath)

	if err != nil {
		return err
	}

	parsedFont, err :=
		opentype.Parse(fontBytes)

	if err != nil {
		return err
	}

	face, err :=
		opentype.NewFace(
			parsedFont,
			&opentype.FaceOptions{
				Size: g.Config.FontSize,
				DPI:  72,
			},
		)

	if err != nil {
		return err
	}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(clr),
		Face: face,
	}

	bounds, _ := d.BoundString(text)

	width :=
		(bounds.Max.X - bounds.Min.X).Ceil()

	height :=
		(bounds.Max.Y - bounds.Min.Y).Ceil()

	x := (g.Config.Width - width) / 2
	y := (g.Config.Height + height) / 2

	d.Dot = fixed.P(x, y)

	d.DrawString(text)

	return nil
}