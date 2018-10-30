package util

import (
	"image"
	"image/draw"

	// for JPEG/PNG format
	_ "image/jpeg"
	_ "image/png"

	"io"
	"io/ioutil"
	"os"

	"github.com/Saqoosha/gobgfx"
)

// LoadProgram loads vertex and fragment shader and compiles it to actual shader program.
func LoadProgram(vspath, fspath string) bgfx.Program {
	data, err := ioutil.ReadFile(vspath)
	if err != nil {
		return bgfx.Program{}
	}
	vs := bgfx.CreateShader(data)

	data, err = ioutil.ReadFile(fspath)
	if err != nil {
		return bgfx.Program{}
	}
	fs := bgfx.CreateShader(data)

	return bgfx.CreateProgram(vs, fs, true)
}

// LoadTexture loads texture from a image file
func LoadTexture(path string) bgfx.Texture {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	config, _, err := image.DecodeConfig(file)
	if err != nil {
		panic(err)
	}

	file.Seek(0, io.SeekStart)
	data, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	rgba := image.NewRGBA(data.Bounds())
	draw.Draw(rgba, rgba.Bounds(), data, image.Pt(0, 0), draw.Src)
	return bgfx.CreateTexture2D(config.Width, config.Height, false, 1, bgfx.TextureFormatRGBA8, bgfx.TextureNone|bgfx.SamplerNone, rgba.Pix)
}
