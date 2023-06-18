package main

import "encoding/binary"
import "os"
import "path/filepath"
import "fmt"

var BYTE_ORDER binary.ByteOrder

func init() {

	_, err1 := os.Readlink("/sbin/init")

	if err1 == nil {

		resolved, err2 := filepath.EvalSymlinks("/sbin/init")

		if err2 == nil {

			buffer := make([]byte, 6)
			handle, err3 := os.Open(resolved)

			if err3 == nil {

				handle.Read(buffer)

				if buffer[5] == 0x01 {
					BYTE_ORDER = binary.LittleEndian
				} else if buffer[5] == 0x02 {
					BYTE_ORDER = binary.BigEndian
				}

			} else {
				fmt.Println(err2)
			}

		} else {
			fmt.Println(err2)
		}

	} else {
		fmt.Println(err1)
	}

}

func main() {

	if BYTE_ORDER == binary.LittleEndian {
		fmt.Println("Little Endian")
	} else if BYTE_ORDER == binary.BigEndian {
		fmt.Println("Big Endian")
	}

}

