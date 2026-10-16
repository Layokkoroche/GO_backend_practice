package main

import "fmt"

func main() {
	var name string = "Сергей"
	var group string = "ЭФБО-07_26"
	course := "Бэкенд-разработка на языке Go"
	var completed_count uint64 = 18_446_744_073_709_551_615
	var interest bool = true

	fmt.Printf("Участник: %s\n", name)
	fmt.Printf("Группа: %s\n", group)
	fmt.Printf("Курс: %s\n", course)
	fmt.Printf("Завершено работ: %d\n", completed_count)
	fmt.Printf("Интерес к бэкенду %t\n", interest)
}
