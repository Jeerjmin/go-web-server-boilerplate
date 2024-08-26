package image

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadHandler(c *gin.Context) {
	resultChan := make(chan int64)

	go func() {
		result := heavyMemoryOperation()
		resultChan <- result
		close(resultChan)
	}()

	result := <-resultChan

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func UploadHandlerSync(c *gin.Context) {
	result := heavyMemoryOperation()

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func heavyMemoryOperation() int64 {
	// Создаем большой массив
	size := 10000000
	data := make([]int, size)

	// Заполняем массив случайными значениями
	for i := 0; i < size; i++ {
		data[i] = i * 2
	}

	// Выполняем какую-либо операцию (например, считаем сумму)
	var sum int64
	for _, value := range data {
		sum += int64(value)
	}

	// Освобождаем память, обнуляя массив
	data = nil

	return sum
}
