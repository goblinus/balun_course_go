package data_types

import "unsafe"

func IsLittleEndian() bool {
	var check int16 = 0x0001
	pointer := (*int8)(unsafe.Pointer(&check))
	return *pointer == 1
}

func IsBigEndian() bool {
	return !IsLittleEndian()
}

func ToLittleEndian(number uint32) uint32 {
	var result uint32

	numberPointer := unsafe.Pointer(&number)
	resultPointer := unsafe.Pointer(&result)

	length := unsafe.Sizeof(number)
	resultIdx := 3
	for numberIdx := 0; numberIdx < int(length); numberIdx++ {
		*(*uint8)(unsafe.Add(resultPointer, resultIdx)) = *(*uint8)(unsafe.Add(numberPointer, uint32(numberIdx)))
		resultIdx--
	}

	return result
}
