package cachepool

import (
	"reflect"
	"unsafe"
)

//go:linkname mallocgc runtime.mallocgc
func mallocgc(size uintptr, typ uintptr, needzero bool) unsafe.Pointer

// RawByteSlice allocates a new byte slice without calling the garbage collector.
func RawByteSlice(size int) []byte {
	// Allocate memory for the slice header.
	header := &reflect.SliceHeader{
		Data: 0, // Will be set later.
		Len:  size,
		Cap:  size,
	}

	// Allocate the memory for the slice itself.
	p := mallocgc(uintptr(size), 0, false)

	// Set the Data field of the header to the allocated memory.
	header.Data = uintptr(p)

	// Convert the header to a byte slice.
	b := *(*[]byte)(unsafe.Pointer(header))
	return b
}

// RawMalloc allocates a new slice. The slice is not zeroed.
func RawMalloc(size int) unsafe.Pointer {
	return mallocgc(uintptr(size), 0, false)
}

// RawMallocByteSlice allocates a new byte slice. The slice is not zeroed.
func RawMallocByteSlice(size int) []byte {
	p := RawMalloc(size)
	b := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(p),
		Len:  size,
		Cap:  size,
	}))
	return b
}
