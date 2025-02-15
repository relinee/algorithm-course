package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"priority-queue/internal"
)

func main() {
	pq := internal.NewPriorityQueue()

	pq.Enqueue(internal.NewTask(internal.Low, 3*time.Second))
	pq.Enqueue(internal.NewTask(internal.Medium, 4*time.Second))
	pq.Enqueue(internal.NewTask(internal.High, 5*time.Second))

	pq.StartWorker()

	go func() {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("[CONSOLE] Введите новое задание в формате ('номер приоритета [0, 1, 2] : время в секундах') или 'stop' для выхода:\n")
		for {
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			if input == "stop" {
				pq.StopWorker()
				fmt.Println("[CONSOLE] Обработка остановлена.")
				return
			}

			var priority int
			var duration int
			_, err := fmt.Sscanf(input, "%d : %d", &priority, &duration)
			if err != nil || (priority < 0 || priority > 2) || duration < 0 {
				fmt.Println("[CONSOLE] Невалидный ввод. Повторите снова. Пример ввода: '1 : 2' - запуск задачи с приоритетом Medium на 2 секунды.")
				continue
			}

			task := internal.NewTask(
				internal.Priority(priority),
				time.Duration(duration)*time.Second,
			)
			pq.Enqueue(task)
		}
	}()

	pq.Wait()
}
