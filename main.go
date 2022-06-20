package main

import "fmt"

func main() {
	var (
		user_position int
	)
	var dict = map[string][]string{
		"english": {
			"apple", "book", "pen", "pencil", "student", "walk", "hello", "chair", "go", "run", "school", "air", "sky",
		},
		"uzbek": {
			"salom", "kitob", "ruchka", "qalam", "dunyo", "anor", "telefon", "do'st", "daryo", "gul", "qora", "ko'k",
		},
		"russian": {
			"ручка", "книга", "вода", "дерево", "стол", "красный", "бытсро",
		},
	}
	fmt.Println(dict)
	fmt.Scan(&user_position)
	
	showPosition()

	if user_position == 1 {
		
	}


}

func showPosition() {
	fmt.Print(`
	  1 ->   List of language
	  2 ->   Words of to given language
	  3 ->   Defining language
	  0 ->   Exit
	  Enter one these parties: `)
}