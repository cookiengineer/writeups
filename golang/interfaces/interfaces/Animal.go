package interfaces

import "example/structs"

type Animal interface {
	Render() *structs.Element
	String() string
}
