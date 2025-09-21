package structs

type Animal struct {
	Element *Element `json:"element"`
}

func NewAnimal(name string) *Animal {

	var animal Animal

	animal.Element = NewElement(name)

	return &animal

}

func (animal *Animal) Render() *Element {
	return animal.Element
}

func (animal *Animal) String() string {
	return animal.Element.String()
}
