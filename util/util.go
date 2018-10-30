package util

import (
	"io/ioutil"

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
