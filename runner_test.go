package opencl

import (
	"slices"
	"testing"
	"unsafe"
)

func TestRunner(t *testing.T) {
	info, _ := Info()

	if len(info.Platforms) < 1 {
		t.Fatal("No OpenCL Devices")
	}
	if len(info.Platforms[0].Devices) < 1 {
		t.Fatal("No OpenCL Devices")
	}

	code := `__kernel void helloworld(__global int* in, __global int* out)
		 {
			 int num = get_global_id(0);
			 out[num] = in[num] * in[num];
		 }`

	// InitRunner
	device := info.Platforms[0].Devices[0]
	runner, err := device.InitRunner()
	if err != nil {
		t.Fatal("InitRunner err:", err)
	}
	defer runner.Free()

	// CompileKernels
	codes := []string{code}
	kernelNameList := []string{"helloworld"}
	err = runner.CompileKernels(codes, kernelNameList, "")
	if err != nil {
		t.Fatal("CompileKernels err:", err)
	}

	// create kernel params
	/* buffer 1 param */
	input := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	itemSize := int(unsafe.Sizeof(input[0]))
	itemCount := len(input)

	input_buf, err := CreateBuffer(runner, READ_ONLY|COPY_HOST_PTR, input)
	if err != nil {
		t.Fatal("CreateBuffer err:", err)
	}
	/* buffer 2 param */
	output_buf, err := runner.CreateEmptyBuffer(WRITE_ONLY, itemCount*itemSize)
	if err != nil {
		t.Fatal("CreateEmptyBuffer err:", err)
	}

	// RunKernel
	err = runner.RunKernel("helloworld", 1, nil, []uint64{uint64(itemCount)}, nil, []KernelParam{
		BufferParam(input_buf),
		BufferParam(output_buf),
	}, true)
	if err != nil {
		t.Fatal("RunKernel err:", err)
	}

	// ReadBuffer
	result := make([]int32, itemCount)
	err = ReadBuffer(runner, 0, output_buf, result)
	if err != nil {
		t.Fatal("ReadBuffer err:", err)
	}

	// check Result
	expected_result := make([]int32, itemCount, itemCount)
	for i, v := range input {
		expected_result[i] = v * v
	}
	if !slices.Equal(result, expected_result) {
		t.Fatal("result error:", result)
	}

}
