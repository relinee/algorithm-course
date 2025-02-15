package internal

import (
	"fmt"
	"sync"
	"time"
)

type PriorityQueue struct {
	stopChan chan struct{}
	wg       sync.WaitGroup
	queues   map[Priority]*Queue[*Task]
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		stopChan: make(chan struct{}),
		queues: map[Priority]*Queue[*Task]{
			Low:    NewQueue[*Task](),
			Medium: NewQueue[*Task](),
			High:   NewQueue[*Task](),
		},
	}
}

func (pq *PriorityQueue) Enqueue(task *Task) {
	switch task.priority {
	case Low:
		pq.queues[Low].Enqueue(task)
	case Medium:
		pq.queues[Medium].Enqueue(task)
	case High:
		pq.queues[High].Enqueue(task)
	}
}

func (pq *PriorityQueue) Dequeue() (*Task, bool) {
	if task, ok := pq.queues[High].Dequeue(); ok {
		return *task, true
	}

	if task, ok := pq.queues[Medium].Dequeue(); ok {
		return *task, true
	}

	if task, ok := pq.queues[Low].Dequeue(); ok {
		return *task, true
	}

	return nil, false
}

func (pq *PriorityQueue) StartWorker() {
	pq.wg.Add(1)
	go pq.worker()
}

func (pq *PriorityQueue) StopWorker() {
	close(pq.stopChan)
	pq.wg.Wait()
}

func (pq *PriorityQueue) Wait() {
	pq.wg.Wait()
}

func (pq *PriorityQueue) worker() {
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
