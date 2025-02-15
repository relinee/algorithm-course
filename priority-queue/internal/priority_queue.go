package internal

import (
	"fmt"
	"sync"
	"time"
)

type Collection[T any] interface {
	Enqueue(T)
	Dequeue() (*T, bool)
}

type PriorityCollection struct {
	stopChan    chan struct{}
	wg          sync.WaitGroup
	collections map[Priority]Collection[*Task]
}

func NewPriorityQueue() *PriorityCollection {
	return &PriorityCollection{
		stopChan: make(chan struct{}),
		collections: map[Priority]Collection[*Task]{
			Low:    NewQueue[*Task](),
			Medium: NewQueue[*Task](),
			High:   NewQueue[*Task](),
		},
	}
}

func NewPriorityStack() *PriorityCollection {
	return &PriorityCollection{
		stopChan: make(chan struct{}),
		collections: map[Priority]Collection[*Task]{
			Low:    NewStack[*Task](),
			Medium: NewStack[*Task](),
			High:   NewStack[*Task](),
		},
	}
}

func (pq *PriorityCollection) Enqueue(task *Task) {
	switch task.priority {
	case Low:
		pq.collections[Low].Enqueue(task)
	case Medium:
		pq.collections[Medium].Enqueue(task)
	case High:
		pq.collections[High].Enqueue(task)
	}
}

func (pq *PriorityCollection) Dequeue() (*Task, bool) {
	if task, ok := pq.collections[High].Dequeue(); ok {
		return *task, true
	}

	if task, ok := pq.collections[Medium].Dequeue(); ok {
		return *task, true
	}

	if task, ok := pq.collections[Low].Dequeue(); ok {
		return *task, true
	}

	return nil, false
}

func (pq *PriorityCollection) StartWorker() {
	pq.wg.Add(1)
	go pq.worker()
}

func (pq *PriorityCollection) StopWorker() {
	close(pq.stopChan)
	pq.wg.Wait()
}

func (pq *PriorityCollection) Wait() {
	pq.wg.Wait()
}

func (pq *PriorityCollection) worker() {
	defer pq.wg.Done()
	for {
		select {
		case <-pq.stopChan:
			return
		default:
			task, ok := pq.Dequeue()
			if ok {
				fmt.Printf("[WORKER] Начало обработки задачи %v. Приоритет: %v. Время обработки: %v\n",
					task.id, task.priority, task.duration)
				time.Sleep(task.duration)
				fmt.Printf("[WORKER] Обработка задачи %v завершена.\n", task.id)
			} else {
				time.Sleep(time.Second)
			}
		}
	}
}
