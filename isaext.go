// Copyright 2019 John Papandriopoulos.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.md file.

package zydis

// ISAExt is an enum of Instruction Set Architecture Extensions.
//go:generate stringer -type=ISAExt -linecomment
type ISAExt int

// IsaExt enum values.
const (
	ISAExtInvalid     ISAExt = iota // INVALID
	ISAExtADOX_ADCX                 // ADOX_ADCX
	ISAExtAES                       // AES
	ISAExtAMD3DNOW                  // AMD3DNOW
	ISAExtAVX                       // AVX
	ISAExtAVX2                      // AVX2
	ISAExtAVX2GATHER                // AVX2GATHER
	ISAExtAVX512EVEX                // AVX512EVEX
	ISAExtAVX512VEX                 // AVX512VEX
	ISAExtAVXAES                    // AVXAES
	ISAExtBase                      // BASE
	ISAExtBMI1                      // BMI1
	ISAExtBMI2                      // BMI2
	ISAExtCET                       // CET
	ISAExtCLDEMOTE                  // CLDEMOTE
	ISAExtCLFLUSHOPT                // CLFLUSHOPT
	ISAExtCLFSH                     // CLFSH
	ISAExtCLWB                      // CLWB
	ISAExtCLZero                    // CLZERO
	ISAExtF16C                      // F16C
	ISAExtFMA                       // FMA
	ISAExtFMA4                      // FMA4
	ISAExtGFNI                      // GFNI
	ISAExtINVPCID                   // INVPCID
	ISAExtKNC                       // KNC
	ISAExtKNCE                      // KNCE
	ISAExtKNCV                      // KNCV
	ISAExtLongMode                  // LONGMODE
	ISAExtLZCNT                     // LZCNT
	ISAExtMMX                       // MMX
	ISAExtMonitor                   // MONITOR
	ISAExtMonitorX                  // MONITORX
	ISAExtMOVBE                     // MOVBE
	ISAExtMOVDIR                    // MOVDIR
	ISAExtMPX                       // MPX
	ISAExtPadlock                   // PADLOCK
	ISAExtPause                     // PAUSE
	ISAExtPCLMULQDQ                 // PCLMULQDQ
	ISAExtPConfig                   // PCONFIG
	ISAExtPKU                       // PKU
	ISAExtPrefetchWT1               // PREFETCHWT1
	ISAExtPT                        // PT
	ISAExtRDPID                     // RDPID
	ISAExtRDRAND                    // RDRAND
	ISAExtRDSEED                    // RDSEED
	ISAExtRDTSCP                    // RDTSCP
	ISAExtRDWRFSGS                  // RDWRFSGS
	ISAExtRTM                       // RTM
	ISAExtSGX                       // SGX
	ISAExtSGX_ENCLV                 // SGX_ENCLV
	ISAExtSHA                       // SHA
	ISAExtSMAP                      // SMAP
	ISAExtSMX                       // SMX
	ISAExtSSE                       // SSE
	ISAExtSSE2                      // SSE2
	ISAExtSSE3                      // SSE3
	ISAExtSSE4                      // SSE4
	ISAExtSSE4A                     // SSE4A
	ISAExtSSSE3                     // SSSE3
	ISAExtSVM                       // SVM
	ISAExtTBM                       // TBM
	ISAExtVAES                      // VAES
	ISAExtVMFUNC                    // VMFUNC
	ISAExtVPCLMULQDQ                // VPCLMULQDQ
	ISAExtVTX                       // VTX
	ISAExtWAITPKG                   // WAITPKG
	ISAExtX87                       // X87
	ISAExtXOP                       // XOP
	ISAExtXSave                     // XSAVE
	ISAExtXSaveC                    // XSAVEC
	ISAExtXSaveOPT                  // XSAVEOPT
	ISAExtXSaveS                    // XSAVE
)
