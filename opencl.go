package opencl

// #cgo CFLAGS: -DCL_TARGET_OPENCL_VERSION=120
// #cgo !darwin LDFLAGS: -lOpenCL
// #cgo darwin LDFLAGS: -framework OpenCL
import "C"
