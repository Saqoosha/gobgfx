package bgfx

// #include <bgfx/c99/bgfx.h>
import "C"
import "unsafe"

// Texture type
type Texture struct {
	h C.bgfx_texture_handle_t
}

// CreateTexture2D creates Texture from raw 2D bytes
func CreateTexture2D(width, height int, hasMips bool, numLayers int, format TextureFormat, flags TextureFlag, data []byte) Texture {
	var mem *C.bgfx_memory_t
	if data != nil {
		mem = C.bgfx_copy(unsafe.Pointer(&data[0]), C.uint32_t(len(data)))
	}
	h := C.bgfx_create_texture_2d(
		C.ushort(width),
		C.ushort(height),
		C._Bool(hasMips),
		C.ushort(numLayers),
		C.bgfx_texture_format_t(format),
		C.ulonglong(flags),
		mem,
	)
	return Texture{h: h}
}

