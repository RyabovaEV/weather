package main

import (
	"flag"
	"fmt"
)

func main() {
	//вызов - go run . --"name"="Vasiliy" --age=20 --isAdmin
	fmt.Println("новый проект")
	name := flag.String("name", "Elena", "Имя пользователя")
	age := flag.Int("age", 18, "Возраст пользователя")
	isAdmin := flag.Bool("isAdmin", false, "Является админом")

	flag.Parse()
	fmt.Println(*name)
	fmt.Println(*age)
	fmt.Println(*isAdmin)
}
