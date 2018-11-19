package bgfx

// #include <bgfx/c99/bgfx.h>
import "C"
import (
	"errors"
	"math"
	"reflect"
	"unsafe"
)

// VertexDecl bgfx_vertex_decl_t
type VertexDecl struct {
	decl C.bgfx_vertex_decl_t
}

// Begin vertex declaration
func (v *VertexDecl) Begin() *VertexDecl {
	C.bgfx_vertex_decl_begin(&v.decl, C.BGFX_RENDERER_TYPE_NOOP)
	return v
}

// Add vertex declaration
func (v *VertexDecl) Add(attrib Attrib, num uint8, typ AttribType, normalized bool, asint bool) *VertexDecl {
	C.bgfx_vertex_decl_add(
		&v.decl,
		C.bgfx_attrib_t(attrib),
		C.uint8_t(num),
		C.bgfx_attrib_type_t(typ),
		C._Bool(normalized),
		C._Bool(asint),
	)
	return v
}

// SetOffset of vertex declaration
func (v *VertexDecl) SetOffset(attrib Attrib, offset uint) *VertexDecl {
	v.decl.offset[attrib] = C.uint16_t(offset)
	return v
}

// Skip vertex declaration
func (v *VertexDecl) Skip(num uint8) *VertexDecl {
	C.bgfx_vertex_decl_skip(&v.decl, C.uint8_t(num))
	return v
}

// End vertex declaration
func (v *VertexDecl) End() *VertexDecl {
	C.bgfx_vertex_decl_end(&v.decl)
	return v
}

// Stride of vertex declaration
func (v *VertexDecl) Stride() int {
	return int(v.decl.stride)
}

// VertexBuffer type
type VertexBuffer struct {
	h C.bgfx_vertex_buffer_handle_t
}

// NewVertexBuffer creates vertex buffer
func NewVertexBuffer(slice interface{}, decl VertexDecl) VertexBuffer {
	val := reflect.ValueOf(slice)
	if val.Kind() != reflect.Slice {
		panic(errors.New("bgfx: expected slice"))
	}
	size := uintptr(val.Len()) * val.Type().Elem().Size()
	return VertexBuffer{
		h: C.bgfx_create_vertex_buffer(
			// to keep things simple for now, we'll just copy
			C.bgfx_copy(unsafe.Pointer(val.Pointer()), C.uint32_t(size)),
			&decl.decl,
			C.ushort(BufferNone),
		),
	}
}

// Destroy vertex buffer
func (vb *VertexBuffer) Destroy() {
	C.bgfx_destroy_vertex_buffer(vb.h)
}

// Set vertex buffer
func (vb *VertexBuffer) Set(stream int8) {
	C.bgfx_set_vertex_buffer(C.uchar(stream), vb.h, 0, C.UINT32_MAX)
}

// DynamicVertexBuffer type
type DynamicVertexBuffer struct {
	h C.bgfx_dynamic_vertex_buffer_handle_t
}

// NewDynamicVertexBuffer creates emmpty dynamic vertex buffer
func NewDynamicVertexBuffer(num int, decl VertexDecl) DynamicVertexBuffer {
	return DynamicVertexBuffer{
		h: C.bgfx_create_dynamic_vertex_buffer(
			C.uint32_t(num),
			&decl.decl,
			C.ushort(BufferNone),
		),
	}
}

// NewDynamicVertexBufferMem creates dynamic vertex buffer and initialize it.
func NewDynamicVertexBufferMem(mem interface{}, decl VertexDecl) DynamicVertexBuffer {
	val := reflect.ValueOf(mem)
	if val.Kind() != reflect.Slice {
		panic(errors.New("bgfx: expected slice"))
	}
	size := uintptr(val.Len()) * val.Type().Elem().Size()
	return DynamicVertexBuffer{
		h: C.bgfx_create_dynamic_vertex_buffer_mem(
			// to keep things simple for now, we'll just copy
			C.bgfx_copy(unsafe.Pointer(val.Pointer()), C.uint32_t(size)),
			&decl.decl,
			C.ushort(BufferNone),
		),
	}
}

// Update dynamic vertex buffer.
func (b *DynamicVertexBuffer) Update(mem interface{}) {
	val := reflect.ValueOf(mem)
	if val.Kind() != reflect.Slice {
		panic(errors.New("bgfx: expected slice"))
	}
	size := uintptr(val.Len()) * val.Type().Elem().Size()
	C.bgfx_update_dynamic_vertex_buffer(b.h, 0, C.bgfx_copy(unsafe.Pointer(val.Pointer()), C.uint32_t(size)))
}

// Set vertex buffer for draw primitive
func (b *DynamicVertexBuffer) Set(stream int8) {
	C.bgfx_set_dynamic_vertex_buffer(C.uint8_t(stream), b.h, 0, C.uint32_t(math.MaxUint32))
}

// Destroy buffer
func (b *DynamicVertexBuffer) Destroy() {
	C.bgfx_destroy_dynamic_vertex_buffer(b.h)
}

// IndexBuffer type
type IndexBuffer struct {
	h C.bgfx_index_buffer_handle_t
}

// NewIndexBuffer creates index buffer
func NewIndexBuffer(data []uint16) IndexBuffer {
	return IndexBuffer{
		h: C.bgfx_create_index_buffer(
			// to keep things simple for now, we'll just copy
			C.bgfx_copy(unsafe.Pointer(&data[0]), C.uint32_t(len(data)*2)),
			C.ushort(BufferNone),
		),
	}
}

// Destroy destroys index buffer
func (ib *IndexBuffer) Destroy() {
	C.bgfx_destroy_index_buffer(ib.h)
}

// Set index buffer
func (ib *IndexBuffer) Set() {
	C.bgfx_set_index_buffer(ib.h, 0, C.UINT32_MAX)
}
