package main

import (
	"fmt"
	"unsafe"
)

func IsLittleEndian() bool {
	var check int16 = 0x0001
	pointer := (*int8)(unsafe.Pointer(&check))
	return *pointer == 1
}

func IsBigEndian() bool {
	return !IsLittleEndian()
}

func main() {
	var result uint32
	var data uint32 = 0x00FF01FF

	if IsBigEndian() {
		fmt.Println("big endian")
	} else {
		fmt.Println("little endian")
	}

	// resultIdx := 0
	dataPointer := unsafe.Pointer(&data)
	resultPointer := unsafe.Pointer(&result)
	fmt.Printf("===== data memory: %X =====\n", data)
	fmt.Printf("[%p]: %X\n", dataPointer, *(*uint8)(dataPointer))
	fmt.Printf("[%p]: %X\n", (*uint8)(unsafe.Add(dataPointer, 1)), *(*uint8)(unsafe.Add(dataPointer, 1)))
	fmt.Printf("[%p]: %X\n", (*uint8)(unsafe.Add(dataPointer, 2)), *(*uint8)(unsafe.Add(dataPointer, 2)))
	fmt.Printf("[%p]: %X\n", (*uint8)(unsafe.Add(dataPointer, 3)), *(*uint8)(unsafe.Add(dataPointer, 3)))

	fmt.Println()

	*(*uint8)(unsafe.Add(resultPointer, 3)) = *(*uint8)(dataPointer)
	*(*uint8)(unsafe.Add(resultPointer, 2)) = *(*uint8)(unsafe.Add(dataPointer, 1))
	*(*uint8)(unsafe.Add(resultPointer, 1)) = *(*uint8)(unsafe.Add(dataPointer, 2))
	*(*uint8)(resultPointer) = *(*uint8)(unsafe.Add(dataPointer, 3))

	fmt.Printf("===== result memory: %X =====\n", result)
	fmt.Printf("[%p]: %X\n", dataPointer, *(*uint8)(dataPointer))
	fmt.Printf("[%p]: %X\n", (*uint8)(unsafe.Add(resultPointer, 1)), *(*uint8)(unsafe.Add(dataPointer, 1)))
	fmt.Printf("[%p]: %X\n", (*uint8)(unsafe.Add(resultPointer, 2)), *(*uint8)(unsafe.Add(dataPointer, 2)))
	fmt.Printf("[%p]: %X\n", (*uint8)(unsafe.Add(resultPointer, 3)), *(*uint8)(unsafe.Add(dataPointer, 3)))

	fmt.Println()
	fmt.Printf("%p : %p\n", dataPointer, resultPointer)
	// for dataIdx := 3; dataIdx >= 0; dataIdx-- {
	// 	value := *(*uint8)(unsafe.Add(dataPointer, dataIdx))
	// 	*(*uint8)(unsafe.Add(resultPointer, resultIdx)) = value
	// 	fmt.Printf(
	// 		"input data : 0x%2o [%b]: %b\n",
	// 		*(*uint8)(unsafe.Add(resultPointer, resultIdx)),
	// 		*(*uint8)(unsafe.Add(resultPointer, value)),
	// 		value,
	// 	)
	// 	resultIdx++
	// }

	// fmt.Printf("input data : 0x%4o\n", data)
	// fmt.Printf("result data: 0x%4o\n", result)
	// var number int32 = 0x12345678
	// pointer := unsafe.Pointer(&number)

	// fmt.Print("0x")
	// if IsLittleEndian() {
	// 	for i := 3; i >= 0; i-- {
	// 		byteValue := *(*int8)(unsafe.Add(pointer, i))
	// 		fmt.Printf("%x", byteValue)
	// 	}
	// }
}
