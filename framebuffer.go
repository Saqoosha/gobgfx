package bgfx

/*
#include <bgfx/c99/bgfx.h>
bgfx_frame_buffer_handle_t hoge(bgfx_texture_handle_t a, bgfx_texture_handle_t b, bool destroyTexture) {
  bgfx_texture_handle_t handles[] = {a, b};
  return bgfx_create_frame_buffer_from_handles(2, handles, destroyTexture);
}
*/
import "C"

// FrameBuffer type
type FrameBuffer struct {
	h C.bgfx_frame_buffer_handle_t
}

// NewFrameBuffer creates frame buffer (simple).
func NewFrameBuffer(width, height int, format TextureFormat, flags TextureFlag) *FrameBuffer {
	h := C.bgfx_create_frame_buffer(
		C.ushort(width),
		C.ushort(height),
		C.bgfx_texture_format_t(format),
		C.uint64_t(flags))
	return &FrameBuffer{h: h}
}

// NewFrameBufferFromTextures creates MRT frame buffer from texture handles (simple).
func NewFrameBufferFromTextures(targets []*Texture, destroyTexture bool) *FrameBuffer {
	h := C.hoge(targets[0].h, targets[1].h, C._Bool(destroyTexture))
	return &FrameBuffer{h: h}
}

// SetView sets view frame buffer.
func (f *FrameBuffer) SetView(id ViewID) {
	C.bgfx_set_view_frame_buffer(C.ushort(id), f.h)
}

// GetTexture obtains texture handle of frame buffer attachment.
func (f *FrameBuffer) GetTexture() *Texture {
	h := C.bgfx_get_texture(f.h, 0)
	return &Texture{h: h}
}

// Destroy frame buffer.
func (f *FrameBuffer) Destroy() {
	C.bgfx_destroy_frame_buffer(f.h)
}
