// *Задание 1. Работа с файлами**
// Что нужно сделать:
// -Напишите программу, которая на вход получала бы строку, введённую пользователем, а в файл писала № строки, дату и сообщение в формате:
// 2020-02-10 15:00:00 продам гараж.
// -При вводе слова exit программа завершает работу.

package main

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

func main() {

	tm := time.Now().Format("2012-01-13 15:27:13")
	var text string
	file, err := os.Create("information.txt")
	defer file.Close()

	for {

		fmt.Println("Введите вашу запись")
		fmt.Scan(&text)

		if err != nil {
			fmt.Println("Не удалось создать файл", err)
			return
		}
		if text == "exit" {
			return
		}
		var b bytes.Buffer
		b.WriteString(fmt.Sprintf("\t%v\t", tm))
		b.WriteString(fmt.Sprintf("%s \n", text))
		_, err = file.Write(b.Bytes())
		if err != nil {
			fmt.Println("Не удалось сделать запись в файл", err)
			return
		}
	}
}
