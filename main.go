package main

import "fmt"

func main() {
	var (
		user_position int
	)
	var dict = map[string][]string{
		"English": {
			"apple", "book", "pen", "pencil", "student", "walk", "hello", "chair", "go", "run", "school", "air", "sky",
		},
		"Uzbek": {
			"salom", "kitob", "ruchka", "qalam", "dunyo", "anor", "telefon", "do'st", "daryo", "gul", "qora", "ko'k",
		},
		"Russian": {
			"ручка", "книга", "вода", "дерево", "стол", "красный", "бытсро",
		},
	}

	
	for {
		showPosition()
		fmt.Scan(&user_position)
		if user_position == 0 {
			fmt.Println("Thank you for your attention")
			break

		} else if user_position == 1 {
			showLanguages(dict, func (s string)  {
				fmt.Println(s)
			})

		} else if user_position == 2{
			showWords(dict, func(s string, i int) {
				fmt.Printf("%d - %s\n", i+1, s)
			})

		} else if user_position == 3 {
			definingLanguage(dict, func(s string) {
				fmt.Printf("This word is in: %s", s)
			})
		}
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
func showLanguages(dict map[string][]string, f func(string)) {
	for k, _ := range dict {
		f(k)
	}
}

func showWords(dict map[string][]string, f func(string, int)) {
	var user_lan string
	i := 0

	fmt.Print("Enter language: ")
	fmt.Scan(&user_lan)
	for k, _ := range dict {
		if k == user_lan {
			for _, v := range dict[k] {
				f(v, i)
				i++
			}
		}
	}
}

func definingLanguage(dict map[string][]string, f func(string)) {
	var user_words string
	
	size := 0
	fmt.Println("Enter your words: ")
	fmt.Scan(&user_words)
	for k, _ := range dict {
		for j := range dict[k] {
			if user_words == dict[k][j] {
				f(k)
				size++
				break
			}
		}
		if size == 1 {
			break
		}
	}
}