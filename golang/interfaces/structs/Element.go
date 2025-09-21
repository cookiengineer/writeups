package structs

type Element struct {
	Name string `json:"name"`
}

func NewElement(name string) *Element {

	var element Element

	element.Name = name

	return &element

}

func (element *Element) String() string {
	return element.Name
}
