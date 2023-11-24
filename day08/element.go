package main

import (
	"errors"
	"unsafe"
)

func getElem(arr []int, idx int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("empty slice")
	}

	if idx <= 0 || idx > len(arr) {
		return 0, errors.New("invalid idx")
	}

	step := unsafe.Sizeof(arr[0])
	value := unsafe.Pointer(uintptr(unsafe.Pointer(&arr[0])) + uintptr(idx)*step)

	return *(*int)(value), nil
}
