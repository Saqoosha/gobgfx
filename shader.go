package bgfx

// #include <bgfx/c99/bgfx.h>
import "C"
import "unsafe"

// Shader type
type Shader struct {
	h C.bgfx_shader_handle_t
}

// NewShader creates shader
func NewShader(data []byte) *Shader {
	return &Shader{
		h: C.bgfx_create_shader(
			// to keep things simple for now, we'll just copy
			C.bgfx_copy(unsafe.Pointer(&data[0]), C.uint32_t(len(data))),
		),
	}
}

// Destroy destroys shader
func (s *Shader) Destroy() {
	C.bgfx_destroy_shader(s.h)
}

// Program type
type Program struct {
	h C.bgfx_program_handle_t
}

// NewProgram creates program
func NewProgram(vsh, fsh *Shader, destroyShaders bool) *Program {
	return &Program{
		h: C.bgfx_create_program(vsh.h, fsh.h, C.bool(destroyShaders)),
	}
}

// Destroy destroys program
func (p *Program) Destroy() {
	C.bgfx_destroy_program(p.h)
}

// Uniform type
type Uniform struct {
	h C.bgfx_uniform_handle_t
}

// NewUniform creates instance
func NewUniform(name string, typ UniformType, num int) *Uniform {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	h := C.bgfx_create_uniform(cname, C.bgfx_uniform_type_t(typ), C.uint16_t(num))
	return &Uniform{h: h}
}

// Destroy destroys uniform
func (u *Uniform) Destroy() {
	C.bgfx_destroy_uniform(u.h)
}

// SetVec4 sets uniform with float32 array
func (u *Uniform) SetVec4(value [4]float32, num int) {
	C.bgfx_set_uniform(u.h, unsafe.Pointer(&value[0]), C.uint16_t(num))
}

// SetTexture sets texture
func (u *Uniform) SetTexture(stage uint8, texture *Texture) {
	C.bgfx_set_texture(C.uint8_t(stage), u.h, texture.h, C.UINT32_MAX)
}
