package main

import "fmt"

func dereference_update(themap *map[int]int) {

	// dereference pointer, update original map
	(*themap)[3] = 0

}

func dereference_replace(themap *map[int]int) {

	// replace pointer address with new map
	*themap = make(map[int]int)

	(*themap)[1] = 33
	(*themap)[3] = 37

}

func main() {

	themap := map[int]int{1: 3, 3: 7}

	fmt.Println("before update", themap)

	dereference_update(&themap)

	fmt.Println("after update", themap)

	dereference_replace(&themap)

	fmt.Println("after replace", themap)

}
