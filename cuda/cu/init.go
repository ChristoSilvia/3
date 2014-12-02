package cu

// This file implements CUDA driver initialization

//#include <cuda.h>
import "C"
import "log"

// Initialize the CUDA driver API.
// Currently, flags must be 0.
// If Init() has not been called, any function from the driver API will panic with ERROR_NOT_INITIALIZED.
func Init(flags int) {
	log.Println("DEBUG: Attempting to initialize")
	err := Result(C.cuInit(C.uint(flags)))
	log.Println("DEBUG: Created Error object")
	if err == SUCCESS {
		log.Println("DEBUG: CUDA_SUCCESS")
	}
	if err == ERROR_INVALID_DEVICE {
		log.Println("DEBUG: CUDA_ERROR_INVALID_DEVICE")
	}
	if err != SUCCESS {
		panic(err)
	}
}
