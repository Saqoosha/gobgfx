package bgfx

// #include <bgfx/c99/bgfx.h>
import "C"

// From "bgfx/c99/bgfx.h"

// RenderType type
type RenderType uint32

// RenderType constants
const (
	RendererTypeNoop       RenderType = C.BGFX_RENDERER_TYPE_NOOP
	RendererTypeDirect3D9             = C.BGFX_RENDERER_TYPE_DIRECT3D9
	RendererTypeDirect3D11            = C.BGFX_RENDERER_TYPE_DIRECT3D11
	RendererTypeDirect3D12            = C.BGFX_RENDERER_TYPE_DIRECT3D12
	RendererTypeGNM                   = C.BGFX_RENDERER_TYPE_GNM
	RendererTypeMetal                 = C.BGFX_RENDERER_TYPE_METAL
	RendererTypeNVN                   = C.BGFX_RENDERER_TYPE_NVN
	RendererTypeOpenGLES              = C.BGFX_RENDERER_TYPE_OPENGLES
	RendererTypeOpenGL                = C.BGFX_RENDERER_TYPE_OPENGL
	RendererTypeVulkan                = C.BGFX_RENDERER_TYPE_VULKAN

	RendererTypeCount = C.BGFX_RENDERER_TYPE_COUNT
)

// Attrib type
type Attrib uint32

// Attrib constants (bgfx_attrib_t)
const (
	AttribPosition  Attrib = C.BGFX_ATTRIB_POSITION
	AttribNormal           = C.BGFX_ATTRIB_NORMAL
	AttribTangent          = C.BGFX_ATTRIB_TANGENT
	AttribBitangent        = C.BGFX_ATTRIB_BITANGENT
	AttribColor0           = C.BGFX_ATTRIB_COLOR0
	AttribColor1           = C.BGFX_ATTRIB_COLOR1
	AttribColor2           = C.BGFX_ATTRIB_COLOR2
	AttribColor3           = C.BGFX_ATTRIB_COLOR3
	AttribIndices          = C.BGFX_ATTRIB_INDICES
	AttribWeight           = C.BGFX_ATTRIB_WEIGHT
	AttribTexcoord0        = C.BGFX_ATTRIB_TEXCOORD0
	AttribTexcoord1        = C.BGFX_ATTRIB_TEXCOORD1
	AttribTexcoord2        = C.BGFX_ATTRIB_TEXCOORD2
	AttribTexcoord3        = C.BGFX_ATTRIB_TEXCOORD3
	AttribTexcoord4        = C.BGFX_ATTRIB_TEXCOORD4
	AttribTexcoord5        = C.BGFX_ATTRIB_TEXCOORD5
	AttribTexcoord6        = C.BGFX_ATTRIB_TEXCOORD6
	AttribTexcoord7        = C.BGFX_ATTRIB_TEXCOORD7

	AttribCount = C.BGFX_ATTRIB_COUNT
)

// AttribType type
type AttribType uint32

// AttribType constants (bgfx_attrib_type_t)
const (
	AttribTypeUint8  AttribType = C.BGFX_ATTRIB_TYPE_UINT8
	AttribTypeUint10            = C.BGFX_ATTRIB_TYPE_UINT10
	AttribTypeInt16             = C.BGFX_ATTRIB_TYPE_INT16
	AttribTypeHalf              = C.BGFX_ATTRIB_TYPE_HALF
	AttribTypeFloat             = C.BGFX_ATTRIB_TYPE_FLOAT

	AttribTypeCount = C.BGFX_ATTRIB_TYPE_COUNT
)

// TextureFormat type
type TextureFormat uint32

// TextureFormat constants
const (
	TextureFormatBC1      TextureFormat = C.BGFX_TEXTURE_FORMAT_BC1
	TextureFormatBC2                    = C.BGFX_TEXTURE_FORMAT_BC2
	TextureFormatBC3                    = C.BGFX_TEXTURE_FORMAT_BC3
	TextureFormatBC4                    = C.BGFX_TEXTURE_FORMAT_BC4
	TextureFormatBC5                    = C.BGFX_TEXTURE_FORMAT_BC5
	TextureFormatBC6H                   = C.BGFX_TEXTURE_FORMAT_BC6H
	TextureFormatBC7                    = C.BGFX_TEXTURE_FORMAT_BC7
	TextureFormatETC1                   = C.BGFX_TEXTURE_FORMAT_ETC1
	TextureFormatETC2                   = C.BGFX_TEXTURE_FORMAT_ETC2
	TextureFormatETC2A                  = C.BGFX_TEXTURE_FORMAT_ETC2A
	TextureFormatETC2A1                 = C.BGFX_TEXTURE_FORMAT_ETC2A1
	TextureFormatPTC12                  = C.BGFX_TEXTURE_FORMAT_PTC12
	TextureFormatPTC14                  = C.BGFX_TEXTURE_FORMAT_PTC14
	TextureFormatPTC12A                 = C.BGFX_TEXTURE_FORMAT_PTC12A
	TextureFormatPTC14A                 = C.BGFX_TEXTURE_FORMAT_PTC14A
	TextureFormatPTC22                  = C.BGFX_TEXTURE_FORMAT_PTC22
	TextureFormatPTC24                  = C.BGFX_TEXTURE_FORMAT_PTC24
	TextureFormatATC                    = C.BGFX_TEXTURE_FORMAT_ATC
	TextureFormatATCE                   = C.BGFX_TEXTURE_FORMAT_ATCE
	TextureFormatATCI                   = C.BGFX_TEXTURE_FORMAT_ATCI
	TextureFormatASTC4x4                = C.BGFX_TEXTURE_FORMAT_ASTC4X4
	TextureFormatASTC5x5                = C.BGFX_TEXTURE_FORMAT_ASTC5X5
	TextureFormatASTC6x6                = C.BGFX_TEXTURE_FORMAT_ASTC6X6
	TextureFormatASTC8x5                = C.BGFX_TEXTURE_FORMAT_ASTC8X5
	TextureFormatASTC8x6                = C.BGFX_TEXTURE_FORMAT_ASTC8X6
	TextureFormatASTC10x5               = C.BGFX_TEXTURE_FORMAT_ASTC10X5

	TextureFormatUnknown = C.BGFX_TEXTURE_FORMAT_UNKNOWN

	TextureFormatR1       = C.BGFX_TEXTURE_FORMAT_R1
	TextureFormatA8       = C.BGFX_TEXTURE_FORMAT_A8
	TextureFormatR8       = C.BGFX_TEXTURE_FORMAT_R8
	TextureFormatR8I      = C.BGFX_TEXTURE_FORMAT_R8I
	TextureFormatR8U      = C.BGFX_TEXTURE_FORMAT_R8U
	TextureFormatR8S      = C.BGFX_TEXTURE_FORMAT_R8S
	TextureFormatR16      = C.BGFX_TEXTURE_FORMAT_R16
	TextureFormatR16I     = C.BGFX_TEXTURE_FORMAT_R16I
	TextureFormatR16U     = C.BGFX_TEXTURE_FORMAT_R16U
	TextureFormatR16F     = C.BGFX_TEXTURE_FORMAT_R16F
	TextureFormatR16S     = C.BGFX_TEXTURE_FORMAT_R16S
	TextureFormatR32I     = C.BGFX_TEXTURE_FORMAT_R32I
	TextureFormatR32U     = C.BGFX_TEXTURE_FORMAT_R32U
	TextureFormatR32F     = C.BGFX_TEXTURE_FORMAT_R32F
	TextureFormatRG8      = C.BGFX_TEXTURE_FORMAT_RG8
	TextureFormatRG8I     = C.BGFX_TEXTURE_FORMAT_RG8I
	TextureFormatRG8U     = C.BGFX_TEXTURE_FORMAT_RG8U
	TextureFormatRG8S     = C.BGFX_TEXTURE_FORMAT_RG8S
	TextureFormatRG16     = C.BGFX_TEXTURE_FORMAT_RG16
	TextureFormatRG16I    = C.BGFX_TEXTURE_FORMAT_RG16I
	TextureFormatRG16U    = C.BGFX_TEXTURE_FORMAT_RG16U
	TextureFormatRG16F    = C.BGFX_TEXTURE_FORMAT_RG16F
	TextureFormatRG16S    = C.BGFX_TEXTURE_FORMAT_RG16S
	TextureFormatRG32I    = C.BGFX_TEXTURE_FORMAT_RG32I
	TextureFormatRG32U    = C.BGFX_TEXTURE_FORMAT_RG32U
	TextureFormatRG32F    = C.BGFX_TEXTURE_FORMAT_RG32F
	TextureFormatRGB8     = C.BGFX_TEXTURE_FORMAT_RGB8
	TextureFormatRGB8I    = C.BGFX_TEXTURE_FORMAT_RGB8I
	TextureFormatRGB8U    = C.BGFX_TEXTURE_FORMAT_RGB8U
	TextureFormatRGB8S    = C.BGFX_TEXTURE_FORMAT_RGB8S
	TextureFormatRGB9E5F  = C.BGFX_TEXTURE_FORMAT_RGB9E5F
	TextureFormatBGRA8    = C.BGFX_TEXTURE_FORMAT_BGRA8
	TextureFormatRGBA8    = C.BGFX_TEXTURE_FORMAT_RGBA8
	TextureFormatRGBA8I   = C.BGFX_TEXTURE_FORMAT_RGBA8I
	TextureFormatRGBA8U   = C.BGFX_TEXTURE_FORMAT_RGBA8U
	TextureFormatRGBA8S   = C.BGFX_TEXTURE_FORMAT_RGBA8S
	TextureFormatRGBA16   = C.BGFX_TEXTURE_FORMAT_RGBA16
	TextureFormatRGBA16I  = C.BGFX_TEXTURE_FORMAT_RGBA16I
	TextureFormatRGBA16U  = C.BGFX_TEXTURE_FORMAT_RGBA16U
	TextureFormatRGBA16F  = C.BGFX_TEXTURE_FORMAT_RGBA16F
	TextureFormatRGBA16S  = C.BGFX_TEXTURE_FORMAT_RGBA16S
	TextureFormatRGBA32I  = C.BGFX_TEXTURE_FORMAT_RGBA32I
	TextureFormatRGBA32U  = C.BGFX_TEXTURE_FORMAT_RGBA32U
	TextureFormatRGBA32F  = C.BGFX_TEXTURE_FORMAT_RGBA32F
	TextureFormatR5G6B5   = C.BGFX_TEXTURE_FORMAT_R5G6B5
	TextureFormatRGBA4    = C.BGFX_TEXTURE_FORMAT_RGBA4
	TextureFormatRGB5A1   = C.BGFX_TEXTURE_FORMAT_RGB5A1
	TextureFormatRGB10A2  = C.BGFX_TEXTURE_FORMAT_RGB10A2
	TextureFormatRG11B10F = C.BGFX_TEXTURE_FORMAT_RG11B10F

	TextureFormatUnknownDepth = C.BGFX_TEXTURE_FORMAT_UNKNOWNDEPTH

	TextureFormatD16   = C.BGFX_TEXTURE_FORMAT_D16
	TextureFormatD24   = C.BGFX_TEXTURE_FORMAT_D24
	TextureFormatD24S8 = C.BGFX_TEXTURE_FORMAT_D24S8
	TextureFormatD32   = C.BGFX_TEXTURE_FORMAT_D32
	TextureFormatD16F  = C.BGFX_TEXTURE_FORMAT_D16F
	TextureFormatD24F  = C.BGFX_TEXTURE_FORMAT_D24F
	TextureFormatD32F  = C.BGFX_TEXTURE_FORMAT_D32F
	TextureFormatD0S8  = C.BGFX_TEXTURE_FORMAT_D0S8

	TextureFormatCount = C.BGFX_TEXTURE_FORMAT_COUNT
)

// UniformType type
type UniformType uint32

// UniformType constants
const (
	UniformTypeSampler UniformType = C.BGFX_UNIFORM_TYPE_SAMPLER
	UniformTypeEnd                 = C.BGFX_UNIFORM_TYPE_END
	UniformTypeVec4                = C.BGFX_UNIFORM_TYPE_VEC4
	UniformTypeMat3                = C.BGFX_UNIFORM_TYPE_MAT3
	UniformTypeMat4                = C.BGFX_UNIFORM_TYPE_MAT4

	UniformTypeCount = C.BGFX_UNIFORM_TYPE_COUNT
)
