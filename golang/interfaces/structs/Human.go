package structs

type Human struct {
	Element *Element `json:"element"`
}

func NewHuman(name string) *Human {

	var human Human

	human.Element = NewElement(name)

	return &human

}

func (human *Human) Render() *Element {
	return human.Element
}

func (human *Human) String() string {
	return human.Element.String()
}
