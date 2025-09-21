package caches

import "example/interfaces"

type Zoo struct {
	Animals map[string]interfaces.Animal
}

func NewZoo() *Zoo {

	var zoo Zoo

	zoo.Animals = make(map[string]interfaces.Animal)

	return &zoo

}

func (zoo *Zoo) Add(animal interfaces.Animal) bool {

	var result bool

	name := animal.String()

	if name != "" {
		zoo.Animals[name] = animal
	}

	return result

}

func (zoo *Zoo) Get(name string) interfaces.Animal {
	return zoo.Animals[name]
}
