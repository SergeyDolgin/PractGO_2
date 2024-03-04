// Задание 2. Интерфейс io.Reader
// Что нужно сделать:
// -Напишите программу, которая читает и выводит в консоль строки из файла, созданного в предыдущей практике, без использования ioutil. Если файл отсутствует или пуст, выведите в консоль соответствующее сообщение.

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	r := NewReader()

	r.Read("information.txt")
}

type Reader interface {
	Read(path string) error
	Display() error
}

func NewReader() Reader {
	return &reader{}
}

type reader struct {
	r io.Reader
}

func (r *reader) Read(path string) error {
	stat, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("Файл не найден")
	}

	if stat.IsDir() {
		return fmt.Errorf("Файл не найден: %w", err)
	}

	b, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("Не удалось прочесть файл: %w", err)
	}

	fmt.Printf("Размер файла: %d изменен: %v\n\n", stat.Size(), stat.ModTime())
	fmt.Println(string(b))

	return nil
}

func (r *reader) Display() error {
	return nil
}
