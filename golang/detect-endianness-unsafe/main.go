package main

import "encoding/binary"
import "unsafe"
import "fmt"

var BYTE_ORDER binary.ByteOrder

func init() {

	buffer := [2]byte{}
	*(*uint16)(unsafe.Pointer(&buffer[0])) = uint16(0xABCD)

	if buffer == [2]byte{0xCD, 0xAB} {
		BYTE_ORDER = binary.LittleEndian
	} else if buffer == [2]byte{0xAB, 0xCD} {
		BYTE_ORDER = binary.BigEndian
	}

}

func main() {

	if BYTE_ORDER == binary.LittleEndian {
		fmt.Println("Little Endian")
	} else if BYTE_ORDER == binary.BigEndian {
		fmt.Println("Big Endian")
	}

}
