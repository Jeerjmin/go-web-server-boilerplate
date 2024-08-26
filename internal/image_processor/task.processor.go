package image_processor

import (
	"github.com/gin-gonic/gin"

	"sync"
)

type Task struct {
	Response chan<- int64
	Context  *gin.Context
}

type TaskProcessor struct {
	tasks chan Task
	wg    sync.WaitGroup
}

func NewTaskProcessor(bufferSize int) *TaskProcessor {
	tp := &TaskProcessor{
		tasks: make(chan Task, bufferSize),
	}
	go tp.run()
	return tp
}

func (tp *TaskProcessor) run() {
	for task := range tp.tasks {
		tp.wg.Add(1)
		go func(t Task) {
			defer tp.wg.Done()
			ProcessTask(t)
		}(task)
	}
}

func (tp *TaskProcessor) AddTask(task Task) {
	tp.tasks <- task
}

func (tp *TaskProcessor) Wait() {
	tp.wg.Wait()
}
