package bgfx

// #include <bgfx/c99/bgfx.h>
import "C"
import "unsafe"

// Shader type
type Shader struct {
	h C.bgfx_shader_handle_t
}

// CreateShader creates shader
func CreateShader(data []byte) Shader {
	return Shader{
		h: C.bgfx_create_shader(
			// to keep things simple for now, we'll just copy
			C.bgfx_copy(unsafe.Pointer(&data[0]), C.uint32_t(len(data))),
		),
	}
}

// DestroyShader destroys shader
func DestroyShader(s Shader) {
	C.bgfx_destroy_shader(s.h)
}

// Program type
type Program struct {
	h C.bgfx_program_handle_t
}

// CreateProgram creates program
func CreateProgram(vsh, fsh Shader, destroyShaders bool) Program {
	return Program{
		h: C.bgfx_create_program(vsh.h, fsh.h, C.bool(destroyShaders)),
	}
}

// DestroyProgram destroys program
func DestroyProgram(p Program) {
	C.bgfx_destroy_program(p.h)
}

// Uniform type
type Uniform struct {
	h C.bgfx_uniform_handle_t
}

// CreateUniform creates instance
func CreateUniform(name string, typ UniformType, num int) Uniform {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	h := C.bgfx_create_uniform(cname, C.bgfx_uniform_type_t(typ), C.uint16_t(num))
	return Uniform{h: h}
}

// DestroyUniform destroys uniform
func DestroyUniform(u Uniform) {
	C.bgfx_destroy_uniform(u.h)
}

// SetTexture sets texture
func SetTexture(stage uint8, sampler Uniform, texture Texture) {
	C.bgfx_set_texture(C.uint8_t(stage), sampler.h, texture.h, C.UINT32_MAX)
}
