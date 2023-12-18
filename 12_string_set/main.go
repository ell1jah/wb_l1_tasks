package main

import "fmt"

func StringSet(stringArr []string) map[string]bool {
	res := make(map[string]bool, len(stringArr)/2)

	// добавляем строки в множество
	for _, str := range stringArr {
		res[str] = true
	}

	return res
}

func main() {
	stringArr := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println("string arr: ", stringArr)
	fmt.Println("string set: ", StringSet(stringArr))
}
