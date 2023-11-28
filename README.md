# go-opencl

go-opencl provides a high-level interface to the OpenCL API, allowing you to run OpenCL programs from Go.

## Development Status

**WARNING**: This project is currently under development and has not been fully tested. Use it at your own risk. We welcome any feedback and contributions.

## Requirements

**linux**
```bash
sudo apt install ocl-icd-opencl-dev opencl-headers
```

**windows**

This project incorporates OpenCL-Headers and OpenCL-ICD-Loader, which are included in the `include-3.0.13` and `lib-windows-3.0.13-x64` directories respectively for Windows.

The sources for these components are as follows:

- OpenCL-Headers: [KhronosGroup/OpenCL-Headers v2023.02.06](https://github.com/KhronosGroup/OpenCL-Headers/releases/tag/v2023.02.06)

- OpenCL-ICD-Loader: [KhronosGroup/OpenCL-SDK v2023.02.06](https://github.com/KhronosGroup/OpenCL-SDK/releases/tag/v2023.02.06)


## cl-info command

The cl-info command provides information about the OpenCL platforms and devices on your system.

To install cl-info, run the following command:

```bash
go install github.com/nexis-dev/go-opencl/cmd/cl-info
```

## OpenCL runner

```go
import cl "github.com/nexis-dev/go-opencl"
```

Refer to the [runner_test.go](./runner_test.go) file for usage examples of the OpenCL runner.

## Other resources

[go-cuda](https://github.com/nexis-dev/go-cuda): High-level Go interface to the CUDA API.

[CCXminer](https://github.com/cyberchain/ccxminer): Cryptonight-GPU-light miner using go-opencl and go-cuda.

OPENCL 3.0 Reference: https://registry.khronos.org/OpenCL/sdk/3.0/docs/man/html/
