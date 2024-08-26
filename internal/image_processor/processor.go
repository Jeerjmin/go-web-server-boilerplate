package image_processor

import "golang.org/x/exp/rand"

func ProcessTask(task Task) {
	// Выполняем ресурсоемкую операцию
	result := heavyMemoryOperation()

	// Отправляем результат обратно в канал
	task.Response <- result
	close(task.Response)
}

func heavyMemoryOperation() int64 {
	// Создаем большой массив
	size := 10000000
	data := make([]int, size)

	// Заполняем массив случайными значениями
	for i := 0; i < size; i++ {
		data[i] = rand.Intn(100)
	}

	// Выполняем какую-либо операцию (например, считаем сумму)
	var sum int64
	for _, value := range data {
		sum += int64(value)
	}

	// Освобождаем память, обнуляя массив
	data = nil

	// Принудительно вызываем сборщик мусора (для тестирования)
	// runtime.GC()

	return sum
}
