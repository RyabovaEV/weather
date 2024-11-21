package main

import (
	"flag"
	"fmt"
	"io"
	"strings"
)

func main() {
	//вызов - go run . --"name"="Vasiliy" --age=20 --isAdmin
	fmt.Println("новый проект")
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода погоды")

	flag.Parse()
	fmt.Println(*city)
	fmt.Println(*format)

	r := strings.NewReader("Привет! я поток данных!")
	block := make([]byte, 4)
	for {
		_, err := r.Read(block)
		fmt.Printf("%q\n", block)
		if err == io.EOF {
			break
		}
	}
}
