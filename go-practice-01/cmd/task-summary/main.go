package main

import (
	"fmt"
)

func statusLabel(status string) string {
	switch status {
	case "new":
		return "Новая"
	case "in_progress":
		return "В работе"
	case "done":
		return "Завершена"
	default:
		return "Неизвестный статус"
	}
}

func countStatus(statuses []string, target string) int {
	count := 0
	for _, status := range statuses {
		if status == target {
			count += 1
		}
	}
	return count
}

func main() {
	statuses := []string{"new", "in_progress", "done", "done"}
	statuses = append(statuses, "new")

	counts := make(map[string]int)

	for _, status := range statuses {
		counts[status]++
	}

	fmt.Printf("Задачь всего: %d\n", len(statuses))

	for i, status := range statuses {
		fmt.Printf("%d. %s\n", i+1, status)
	}

	fmt.Printf("Завершено: %d\n", counts["done"])

	keys := make([]string, 0, len(counts))
	for k := range counts {
		keys = append(keys, k)
	}

	fmt.Print("Статистика: ")
	for i, status := range keys {
		fmt.Printf("%s=%d", status, counts[status])
		if i == len(keys)-1 {
			fmt.Println()
			continue
		}
		fmt.Print(", ")
	}
}

// Задач всего: 5
// 1. Новая
// 2. В работе
// 3. Завершена
// 4. Завершена
// 5. Новая
// Завершено: 2
// Статистика: new=2, in_progress=1, done=2
