package main

import "fmt"
import "os"
import "slices"
import "sort"
import "strconv"
import "strings"

type Data map[string]any

func main() {

	dataset := []Data{
		map[string]any{
			"name": "Jean-Luc Picard",
			"rank": "Captain",
			"age":  59,
		},
		map[string]any{
			"name": "William Riker",
			"rank": "Commander",
			"age":  29,
		},
		map[string]any{
			"name": "Beverly Crusher",
			"rank": "Commander",
			"age":  40,
		},
		map[string]any{
			"name": "Deanna Troi",
			"rank": "Lieutenant Commander",
			"age":  28,
		},
		map[string]any{
			"name": "Data",
			"rank": "Lieutenant Commander",
			"age":  26,
		},
		map[string]any{
			"name": "Geordi La Forge",
			"rank": "Lieutenant Commander",
			"age":  29,
		},
		map[string]any{
			"name": "Worf",
			"rank": "Lieutenant",
			"age":  24,
		},
		map[string]any{
			"name": "Wesley Crusher",
			"rank": "Ensign",
			"age":  16,
		},
	}

	property := "rank"

	if len(os.Args) == 2 {

		tmp := strings.TrimSpace(os.Args[1])

		if tmp == "name" || tmp == "rank" || tmp == "age" {
			property = tmp
		}

	}

	fmt.Println("Sorting []map[string]any by \"" + property + "\"")

	sort.Slice(dataset, func(a int, b int) bool {

		value_a, ok_a := dataset[a][property]
		value_b, ok_b := dataset[b][property]

		if ok_a == true && ok_b == true {

			switch value_a.(type) {
			case bool:

				tmp_a := value_a.(bool)
				tmp_b := value_b.(bool)

				if tmp_a == true && tmp_b == false {
					return true
				} else {
					return false
				}

			case int:

				tmp_a := value_a.(int)
				tmp_b := value_b.(int)

				return tmp_a < tmp_b

			case string:

				tmp_a := value_a.(string)
				tmp_b := value_b.(string)

				if property == "rank" {

					ranks := []string{
						"Captain",
						"Commander",
						"Lieutenant Commander",
						"Lieutenant",
						"Ensign",
					}

					index_a := slices.Index(ranks, tmp_a)
					index_b := slices.Index(ranks, tmp_b)

					return index_a < index_b

				} else {
					return tmp_a < tmp_b
				}

			}

		}

		return false

	})

	if property == "name" {

		for _, data := range dataset {

			name, _ := data["name"].(string)
			rank, _ := data["rank"].(string)
			age,  _ := data["age"].(int)

			fmt.Println(name + " | " + rank + " | " + strconv.Itoa(age))

		}

	} else if property == "age" {

		for _, data := range dataset {

			name, _ := data["name"].(string)
			rank, _ := data["rank"].(string)
			age,  _ := data["age"].(int)

			fmt.Println(strconv.Itoa(age) + " | " + name + " | " + rank)

		}

	} else if property == "rank" {

		for _, data := range dataset {

			name, _ := data["name"].(string)
			rank, _ := data["rank"].(string)
			age,  _ := data["age"].(int)

			fmt.Println(rank + " | " + name + " | " + strconv.Itoa(age))

		}

	}

}
