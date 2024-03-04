// Задание 3. Уровни доступа
// Что нужно сделать:
// Напишите программу, создающую текстовый файл только для чтения, и проверьте, что в него нельзя записать данные.

package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("file.txt")
	if err := os.Chmod("file.txt", 0444); err != nil {
		fmt.Println(err)
	}
	file.WriteString("Осталось мало времени")
	if err != nil {
		fmt.Println(err)
	}
	file.Close()

	file, err = os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

}
