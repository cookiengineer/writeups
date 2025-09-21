package main

import "example/caches"
import "example/structs"
import "fmt"

func main() {

	zoo := caches.NewZoo()

	animal1 := structs.NewAnimal("Dog")
	animal2 := structs.NewAnimal("Cat")

	zoo.Add(animal1)
	zoo.Add(animal2)

	human1  := structs.NewHuman("John")
	human2  := structs.NewHuman("Jane")

	zoo.Add(human1)
	zoo.Add(human2)


	dog_as_interface := zoo.Get("Dog")

	if dog_as_interface != nil {

		dog_as_animal, ok := dog_as_interface.(*structs.Animal)

		if ok == true {
			fmt.Println("Dog as Animal:", dog_as_animal.String())
		}

	}

	john_as_interface := zoo.Get("John")

	if john_as_interface != nil {

		john_as_human, ok := john_as_interface.(*structs.Human)

		if ok == true {
			fmt.Println("John as Human:", john_as_human.String())
		}

	}

}
