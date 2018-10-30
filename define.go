package bgfx

// #include <bgfx/c99/bgfx.h>
import "C"

// From "bgfx/defines.h"

// State type
type State uint64

// State constants
const (
	StateWriteR    State = C.BGFX_STATE_WRITE_R
	StateWriteG          = C.BGFX_STATE_WRITE_G
	StateWriteB          = C.BGFX_STATE_WRITE_B
	StateWriteA          = C.BGFX_STATE_WRITE_A
	StateWriteZ          = C.BGFX_STATE_WRITE_Z
	StateWriteRGB        = StateWriteR | StateWriteG | StateWriteB
	StateWriteMask       = StateWriteRGB | StateWriteA | StateWriteZ

	/// Depth test state. When `BGFX_STATE_DEPTH_` is not specified depth test will be disabled.
	StateDepthTestLess     = C.BGFX_STATE_DEPTH_TEST_LESS
	StateDepthTestLequal   = C.BGFX_STATE_DEPTH_TEST_LEQUAL
	StateDepthTestEqual    = C.BGFX_STATE_DEPTH_TEST_EQUAL
	StateDepthTestGequal   = C.BGFX_STATE_DEPTH_TEST_GEQUAL
	StateDepthTestGreater  = C.BGFX_STATE_DEPTH_TEST_GREATER
	StateDepthTestNotequal = C.BGFX_STATE_DEPTH_TEST_NOTEQUAL
	StateDepthTestNever    = C.BGFX_STATE_DEPTH_TEST_NEVER
	StateDepthTestAlways   = C.BGFX_STATE_DEPTH_TEST_ALWAYS
	StateDepthTestShift    = C.BGFX_STATE_DEPTH_TEST_SHIFT
	StateDepthTestMask     = C.BGFX_STATE_DEPTH_TEST_MASK

	/// Use MakeBlendFuncState(_src, _dst) or MakeBlendFuncSeparateState(_srcRGB, _dstRGB, _srcA, _dstA) helper method.
	StateBlendZero        = C.BGFX_STATE_BLEND_ZERO
	StateBlendOne         = C.BGFX_STATE_BLEND_ONE
	StateBlendSrcColor    = C.BGFX_STATE_BLEND_SRC_COLOR
	StateBlendInvSrcColor = C.BGFX_STATE_BLEND_INV_SRC_COLOR
	StateBlendSrcAlpha    = C.BGFX_STATE_BLEND_SRC_ALPHA
	StateBlendInvSrcAlpha = C.BGFX_STATE_BLEND_INV_SRC_ALPHA
	StateBlendDstAlpha    = C.BGFX_STATE_BLEND_DST_ALPHA
	StateBlendInvDstAlpha = C.BGFX_STATE_BLEND_INV_DST_ALPHA
	StateBlendDstColor    = C.BGFX_STATE_BLEND_DST_COLOR
	StateBlendInvDstColor = C.BGFX_STATE_BLEND_INV_DST_COLOR
	StateBlendSrcAlphaSat = C.BGFX_STATE_BLEND_SRC_ALPHA_SAT
	StateBlendFactor      = C.BGFX_STATE_BLEND_FACTOR
	StateBlendInvFactor   = C.BGFX_STATE_BLEND_INV_FACTOR
	StateBlendShift       = C.BGFX_STATE_BLEND_SHIFT
	StateBlendMask        = C.BGFX_STATE_BLEND_MASK

	/// Use BGFX_STATE_BLEND_EQUATION(_equation) or BGFX_STATE_BLEND_EQUATION_SEPARATE(_equationRGB, _equationA)
	/// helper macros.
	StateBlendEquationAdd    = C.BGFX_STATE_BLEND_EQUATION_ADD
	StateBlendEquationSub    = C.BGFX_STATE_BLEND_EQUATION_SUB
	StateBlendEquationRevsub = C.BGFX_STATE_BLEND_EQUATION_REVSUB
	StateBlendEquationMin    = C.BGFX_STATE_BLEND_EQUATION_MIN
	StateBlendEquationMax    = C.BGFX_STATE_BLEND_EQUATION_MAX
	StateBlendEquationShift  = C.BGFX_STATE_BLEND_EQUATION_SHIFT
	StateBlendEquationMask   = C.BGFX_STATE_BLEND_EQUATION_MASK

	StateBlendIndependent     = C.BGFX_STATE_BLEND_INDEPENDENT
	StateBlendAlphaToCoverage = C.BGFX_STATE_BLEND_ALPHA_TO_COVERAGE

	/// Cull state. When `BGFX_STATE_CULL_*` is not specified culling will be disabled.
	StateCullCW    = C.BGFX_STATE_CULL_CW
	StateCullCCW   = C.BGFX_STATE_CULL_CCW
	StateCullShift = C.BGFX_STATE_CULL_SHIFT
	StateCullMask  = C.BGFX_STATE_CULL_MASK

	/// See BGFX_STATE_ALPHA_REF(_ref) helper macro.
	StateAlphaRefShift = C.BGFX_STATE_ALPHA_REF_SHIFT
	StateAlphaRefMask  = C.BGFX_STATE_ALPHA_REF_MASK

	StatePtTristrip  = C.BGFX_STATE_PT_TRISTRIP
	StatePtLines     = C.BGFX_STATE_PT_LINES
	StatePtLinestrip = C.BGFX_STATE_PT_LINESTRIP
	StatePtPoints    = C.BGFX_STATE_PT_POINTS
	StatePtShift     = C.BGFX_STATE_PT_SHIFT
	StatePtMask      = C.BGFX_STATE_PT_MASK

	/// See BGFX_STATE_POINT_SIZE(_size) helper macro.
	StatePointSizeShift = C.BGFX_STATE_POINT_SIZE_SHIFT
	StatePointSizeMask  = C.BGFX_STATE_POINT_SIZE_MASK

	/// Enable MSAA write when writing into MSAA frame buffer.
	/// This flag is ignored when not writing into MSAA frame buffer.
	StateMsaa               = C.BGFX_STATE_MSAA
	StateLineaa             = C.BGFX_STATE_LINEAA
	StateConservativeRaster = C.BGFX_STATE_CONSERVATIVE_RASTER

	/// Do not use!
	StateReservedShift = C.BGFX_STATE_RESERVED_SHIFT
	StateReservedMask  = C.BGFX_STATE_RESERVED_MASK

	///
	StateNone = C.BGFX_STATE_NONE
	StateMask = C.BGFX_STATE_MASK
)

// MakeBlendFuncState makes state value for blend func
func MakeBlendFuncState(src, dst State) State {
	return MakeBlendFuncSeparateState(src, dst, src, dst)
}

// MakeBlendFuncSeparateState makes state value for blend func color and alpha separately
func MakeBlendFuncSeparateState(srcRGB, dstRGB, srcA, dstA State) State {
	return State((uint64(srcRGB) | (uint64(dstRGB) << 4) | ((uint64(srcA) | (uint64(dstA) << 4)) << 8)))
}

// ClearFlag type
type ClearFlag uint16

// ClearFlag constants
const (
	ClearNone             ClearFlag = C.BGFX_CLEAR_NONE
	ClearColor                      = C.BGFX_CLEAR_COLOR
	ClearDepth                      = C.BGFX_CLEAR_DEPTH
	ClearStencil                    = C.BGFX_CLEAR_STENCIL
	ClearDiscardColor0              = C.BGFX_CLEAR_DISCARD_COLOR_0
	ClearDiscardColor1              = C.BGFX_CLEAR_DISCARD_COLOR_1
	ClearDiscardColor2              = C.BGFX_CLEAR_DISCARD_COLOR_2
	ClearDiscardColor3              = C.BGFX_CLEAR_DISCARD_COLOR_3
	ClearDiscardColor4              = C.BGFX_CLEAR_DISCARD_COLOR_4
	ClearDiscardColor5              = C.BGFX_CLEAR_DISCARD_COLOR_5
	ClearDiscardColor6              = C.BGFX_CLEAR_DISCARD_COLOR_6
	ClearDiscardColor7              = C.BGFX_CLEAR_DISCARD_COLOR_7
	ClearDiscardDepth               = C.BGFX_CLEAR_DISCARD_DEPTH
	ClearDiscardStencil             = C.BGFX_CLEAR_DISCARD_STENCIL
	ClearDiscardColorMask           = ClearDiscardColor0 | ClearDiscardColor1 | ClearDiscardColor2 | ClearDiscardColor3 | ClearDiscardColor4 | ClearDiscardColor5 | ClearDiscardColor6 | ClearDiscardColor7
	ClearDiscard                    = ClearDiscardColorMask | ClearDiscardDepth | ClearDiscardStencil
)

// DebugFlag type
type DebugFlag uint32

// DebugFlag constants
const (
	DebugNone      DebugFlag = C.BGFX_DEBUG_NONE
	DebugWireframe           = C.BGFX_DEBUG_WIREFRAME
	DebugIFH                 = C.BGFX_DEBUG_IFH
	DebugStats               = C.BGFX_DEBUG_STATS
	DebugText                = C.BGFX_DEBUG_TEXT
	DebugProfiler            = C.BGFX_DEBUG_PROFILER
)

// BufferFlag type
type BufferFlag uint16

// BufferFlag constants
const (
	BufferNone BufferFlag = C.BGFX_BUFFER_NONE

	BufferComputeFormat8x1   = C.BGFX_BUFFER_COMPUTE_FORMAT_8x1
	BufferComputeFormat8x2   = C.BGFX_BUFFER_COMPUTE_FORMAT_8x2
	BufferComputeFormat8x4   = C.BGFX_BUFFER_COMPUTE_FORMAT_8x4
	BufferComputeFormat16x1  = C.BGFX_BUFFER_COMPUTE_FORMAT_16x1
	BufferComputeFormat16x2  = C.BGFX_BUFFER_COMPUTE_FORMAT_16x2
	BufferComputeFormat16x4  = C.BGFX_BUFFER_COMPUTE_FORMAT_16x4
	BufferComputeFormat32x1  = C.BGFX_BUFFER_COMPUTE_FORMAT_32x1
	BufferComputeFormat32x2  = C.BGFX_BUFFER_COMPUTE_FORMAT_32x2
	BufferComputeFormat32x4  = C.BGFX_BUFFER_COMPUTE_FORMAT_32x4
	BufferComputeFormatShift = C.BGFX_BUFFER_COMPUTE_FORMAT_SHIFT
	BufferComputeFormatMask  = C.BGFX_BUFFER_COMPUTE_FORMAT_MASK

	BufferComputeTypeInt   = C.BGFX_BUFFER_COMPUTE_TYPE_INT
	BufferComputeTypeUint  = C.BGFX_BUFFER_COMPUTE_TYPE_UINT
	BufferComputeTypeFloat = C.BGFX_BUFFER_COMPUTE_TYPE_FLOAT
	BufferComputeTypeShift = C.BGFX_BUFFER_COMPUTE_TYPE_SHIFT
	BufferComputeTypeMask  = C.BGFX_BUFFER_COMPUTE_TYPE_MASK

	BufferComputeRead  = C.BGFX_BUFFER_COMPUTE_READ
	BufferComputeWrite = C.BGFX_BUFFER_COMPUTE_WRITE
	BufferDrawIndirect = C.BGFX_BUFFER_DRAW_INDIRECT
	BufferAllowResize  = C.BGFX_BUFFER_ALLOW_RESIZE
)

// TextureFlag type
type TextureFlag uint64

// TextureFlag constants
const (
	TextureNone         TextureFlag = C.BGFX_TEXTURE_NONE
	TextureMSAASample               = C.BGFX_TEXTURE_MSAA_SAMPLE
	TextureRT                       = C.BGFX_TEXTURE_RT
	TextureRTMSAAX2                 = C.BGFX_TEXTURE_RT_MSAA_X2
	TextureRTMSAAX4                 = C.BGFX_TEXTURE_RT_MSAA_X4
	TextureRTMSAAX8                 = C.BGFX_TEXTURE_RT_MSAA_X8
	TextureRTMSAAX16                = C.BGFX_TEXTURE_RT_MSAA_X16
	TextureRTMSAAShift              = C.BGFX_TEXTURE_RT_MSAA_SHIFT
	TextureRTMSAAMask               = C.BGFX_TEXTURE_RT_MSAA_MASK
	TextureRTWriteOnly              = C.BGFX_TEXTURE_RT_WRITE_ONLY
	TextureRTMask                   = C.BGFX_TEXTURE_RT_MASK
	TextureComputeWrite             = C.BGFX_TEXTURE_COMPUTE_WRITE
	TextureSRGB                     = C.BGFX_TEXTURE_SRGB
	TextureBlitDst                  = C.BGFX_TEXTURE_BLIT_DST
	TextureReadBack                 = C.BGFX_TEXTURE_READ_BACK

	SamplerNone             = C.BGFX_SAMPLER_NONE
	SamplerUMirror          = C.BGFX_SAMPLER_U_MIRROR
	SamplerUClamp           = C.BGFX_SAMPLER_U_CLAMP
	SamplerUBorder          = C.BGFX_SAMPLER_U_BORDER
	SamplerUShift           = C.BGFX_SAMPLER_U_SHIFT
	SamplerUMask            = C.BGFX_SAMPLER_U_MASK
	SamplerVMirror          = C.BGFX_SAMPLER_V_MIRROR
	SamplerVClamp           = C.BGFX_SAMPLER_V_CLAMP
	SamplerVBorder          = C.BGFX_SAMPLER_V_BORDER
	SamplerVShift           = C.BGFX_SAMPLER_V_SHIFT
	SamplerVMask            = C.BGFX_SAMPLER_V_MASK
	SamplerWMirror          = C.BGFX_SAMPLER_W_MIRROR
	SamplerWClamp           = C.BGFX_SAMPLER_W_CLAMP
	SamplerWBorder          = C.BGFX_SAMPLER_W_BORDER
	SamplerWShift           = C.BGFX_SAMPLER_W_SHIFT
	SamplerWMask            = C.BGFX_SAMPLER_W_MASK
	SamplerMinPoint         = C.BGFX_SAMPLER_MIN_POINT
	SamplerMinAnisotropic   = C.BGFX_SAMPLER_MIN_ANISOTROPIC
	SamplerMinShift         = C.BGFX_SAMPLER_MIN_SHIFT
	SamplerMinMask          = C.BGFX_SAMPLER_MIN_MASK
	SamplerMagPoint         = C.BGFX_SAMPLER_MAG_POINT
	SamplerMagAnisotropic   = C.BGFX_SAMPLER_MAG_ANISOTROPIC
	SamplerMagShift         = C.BGFX_SAMPLER_MAG_SHIFT
	SamplerMagMask          = C.BGFX_SAMPLER_MAG_MASK
	SamplerMipPoint         = C.BGFX_SAMPLER_MIP_POINT
	SamplerMipShift         = C.BGFX_SAMPLER_MIP_SHIFT
	SamplerMipMask          = C.BGFX_SAMPLER_MIP_MASK
	SamplerCompareLess      = C.BGFX_SAMPLER_COMPARE_LESS
	SamplerCompareLequal    = C.BGFX_SAMPLER_COMPARE_LEQUAL
	SamplerCompareEqual     = C.BGFX_SAMPLER_COMPARE_EQUAL
	SamplerCompareGequal    = C.BGFX_SAMPLER_COMPARE_GEQUAL
	SamplerCompareGreater   = C.BGFX_SAMPLER_COMPARE_GREATER
	SamplerCompareNotequal  = C.BGFX_SAMPLER_COMPARE_NOTEQUAL
	SamplerCompareNever     = C.BGFX_SAMPLER_COMPARE_NEVER
	SamplerCompareAlways    = C.BGFX_SAMPLER_COMPARE_ALWAYS
	SamplerCompareShift     = C.BGFX_SAMPLER_COMPARE_SHIFT
	SamplerCompareMask      = C.BGFX_SAMPLER_COMPARE_MASK
	SamplerSampleStencil    = C.BGFX_SAMPLER_SAMPLE_STENCIL
	SamplerBorderColorShift = C.BGFX_SAMPLER_BORDER_COLOR_SHIFT
	SamplerBorderColorMask  = C.BGFX_SAMPLER_BORDER_COLOR_MASK
	SamplerReservedShift    = C.BGFX_SAMPLER_RESERVED_SHIFT
	SamplerReservedMask     = C.BGFX_SAMPLER_RESERVED_MASK
)

// ResetFlag type
type ResetFlag uint32

// ResetFlag constants
const (
	ResetNone             ResetFlag = C.BGFX_RESET_NONE
	ResetFullscreen                 = C.BGFX_RESET_FULLSCREEN
	ResetFullscreenShift            = C.BGFX_RESET_FULLSCREEN_SHIFT
	ResetFullscreenMask             = C.BGFX_RESET_FULLSCREEN_MASK
	ResetMSAAX2                     = C.BGFX_RESET_MSAA_X2
	ResetMSAAX4                     = C.BGFX_RESET_MSAA_X4
	ResetMSAAX8                     = C.BGFX_RESET_MSAA_X8
	ResetMSAAX16                    = C.BGFX_RESET_MSAA_X16
	ResetMSAAShift                  = C.BGFX_RESET_MSAA_SHIFT
	ResetMSAAMask                   = C.BGFX_RESET_MSAA_MASK
	ResetVsync                      = C.BGFX_RESET_VSYNC
	ResetMaxanisotropy              = C.BGFX_RESET_MAXANISOTROPY
	ResetCapture                    = C.BGFX_RESET_CAPTURE
	ResetFlushAfterRender           = C.BGFX_RESET_FLUSH_AFTER_RENDER
	ResetFlipAfterRender            = C.BGFX_RESET_FLIP_AFTER_RENDER
	ResetSRGBBackbuffer             = C.BGFX_RESET_SRGB_BACKBUFFER
	ResetHDR10                      = C.BGFX_RESET_HDR10
	ResetHiDPI                      = C.BGFX_RESET_HIDPI
	ResetDepthClamp                 = C.BGFX_RESET_DEPTH_CLAMP
	ResetSuspend                    = C.BGFX_RESET_SUSPEND
	ResetReservedShift              = C.BGFX_RESET_RESERVED_SHIFT
	ResetReservedMask               = C.BGFX_RESET_RESERVED_MASK
)
