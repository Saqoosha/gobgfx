package bgfx

/*
#cgo CFLAGS: -I./include/
#cgo linux LDFLAGS: -L./lib/linux64_gcc/
#cgo darwin LDFLAGS: -L./lib/osx64_clang/
#cgo darwin LDFLAGS: -framework Cocoa -framework QuartzCore -framework Metal
#cgo LDFLAGS: -lstdc++ -lbgfxDebug -lbxDebug -lbimgDebug

#include <bgfx/c99/bgfx.h>

void bgfx_dbg_text_print(uint16_t _x, uint16_t _y, uint8_t _attr, const char* text) {
	bgfx_dbg_text_printf(_x, _y, _attr, text);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// Config is for initializing bgfx
type Config struct {
	RenderType         RenderType
	VendorID, DeviceID uint16
	Debug, Profile     bool

	NativeDisplayType, NativeWindowHandle, Context, BackBuffer, BackBufferDS uintptr

	Format                          TextureFormat
	Width, Height, Reset            uint32
	NumBackBuffers, MaxFrameLatency uint8
}

// Init bgfx
func Init(params Config) {
	var init C.bgfx_init_t
	C.bgfx_init_ctor(&init)
	init._type = C.bgfx_renderer_type_t(params.RenderType)
	init.vendorId = C.uint16_t(params.VendorID)
	init.deviceId = C.uint16_t(params.DeviceID)
	init.debug = C._Bool(params.Debug)
	init.profile = C._Bool(params.Profile)
	init.platformData.ndt = unsafe.Pointer(params.NativeDisplayType)
	init.platformData.nwh = unsafe.Pointer(params.NativeWindowHandle)
	init.platformData.context = unsafe.Pointer(params.Context)
	init.platformData.backBuffer = unsafe.Pointer(params.BackBuffer)
	init.platformData.backBufferDS = unsafe.Pointer(params.BackBufferDS)
	init.resolution.format = C.bgfx_texture_format_t(params.Format)
	init.resolution.width = C.uint32_t(params.Width)
	init.resolution.height = C.uint32_t(params.Height)
	init.resolution.reset = C.uint32_t(params.Reset)
	init.resolution.numBackBuffers = C.uint8_t(params.NumBackBuffers)
	init.resolution.maxFrameLatency = C.uint8_t(params.MaxFrameLatency)
	C.bgfx_init(&init)
}

// Shutdown bgfx
func Shutdown() {
	C.bgfx_shutdown()
}

// Reset bgfx
func Reset(width uint, height uint, flags uint, textureFormat TextureFormat) {
	C.bgfx_reset(C.uint(width), C.uint(height), C.uint(flags), C.bgfx_texture_format_t(textureFormat))
}

// Frame advances to the next frame. Returns the current frame number.
func Frame(capture bool) uint32 {
	return uint32(C.bgfx_frame(C._Bool(capture)))
}

// SetDebug bgfx
func SetDebug(flags DebugFlag) {
	C.bgfx_set_debug(C.uint(flags))

}

// DebugTextClear clears debug text
func DebugTextClear(backgroundColor uint8, smallFont bool) {
	C.bgfx_dbg_text_clear(C.uchar(backgroundColor), C._Bool(smallFont))
}

// DebugTextPrintf prints debug text with format
func DebugTextPrintf(x, y int, attr uint8, format string, args ...interface{}) {
	text := []byte(fmt.Sprintf(format+"\x00", args...))
	C.bgfx_dbg_text_print(
		C.ushort(x),
		C.ushort(y),
		C.uchar(attr),
		(*C.char)(unsafe.Pointer(&text[0])),
	)
}

// ViewID bgfx_view_id_t
type ViewID uint16

// SetViewRect sets view rect
func SetViewRect(view ViewID, x, y, w, h int) {
	C.bgfx_set_view_rect(
		C.ushort(view),
		C.ushort(x),
		C.ushort(y),
		C.ushort(w),
		C.ushort(h),
	)
}

// SetViewTransform sets view transform
func SetViewTransform(viewID ViewID, view, proj [16]float32) {
	C.bgfx_set_view_transform(
		C.ushort(viewID),
		unsafe.Pointer(&view[0]),
		unsafe.Pointer(&proj[0]),
	)
}

// SetViewClear bgfx
func SetViewClear(view ViewID, flags ClearFlag, rgba uint32, depth float32, stencil uint8) {
	C.bgfx_set_view_clear(C.ushort(view), C.ushort(flags), C.uint(rgba), C.float(depth), C.uchar(stencil))
}

// Touch bgfx
func Touch(id ViewID) {
	C.bgfx_touch(C.ushort(id))
}

// SetTransform sets transform
func SetTransform(mtx [16]float32) {
	C.bgfx_set_transform(unsafe.Pointer(&mtx[0]), 1)
}

// SetState set render states for draw primitive.
func SetState(state State) {
	C.bgfx_set_state(C.uint64_t(state), 0)
}

// Submit bgfx
func Submit(id ViewID, program Program, depth int32, preserveState bool) {
	C.bgfx_submit(C.ushort(id), program.h, C.uint32_t(depth), C._Bool(preserveState))
}
