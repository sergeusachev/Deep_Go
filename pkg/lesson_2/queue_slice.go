package lesson_2

import(
	"fmt"
	"errors"
)

type QueueS struct {
	size int
	front int
	rear int
	buffer []int
}

func NewQueueS(size int) *QueueS {
	return &QueueS{
		size: size,
		front: -1,
		rear: -1,
		buffer: make([]int, size),
	}
}

func (q *Queue) Enqueue(element int) error {
	if q.IsFull() {
		return errors.New("Can't enqueue. Queue is full!")
	}
	
	if q.IsEmpty() {
		q.front = 0
	}
	
	q.rear++
	q.buffer[q.rear] = element
	
	return nil
}

func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Can't dequeue. Queue is empty!")
	}
	
	dequeuedElement := q.buffer[q.front]
	if q.front == q.rear { //to empty state
		q.front = -1
		q.rear = -1
	} else {
		q.front++
	}
	
	return dequeuedElement, nil
}

func (q *Queue) IsEmpty() bool {
	return q.front == -1
}

func (q *Queue) IsFull() bool {
	return q.rear == q.size - 1
}

func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Cant't peek. Queue is empty!")
	}
	return q.buffer[q.front], nil
}

func (q *Queue) PrintQueueStruct() {
	fmt.Println("queue.size: ", q.size)
	fmt.Println("queue.rear: ", q.rear)
	fmt.Println("queue.front: ", q.front)
	fmt.Println("queue.buffer: ", q.buffer)
	fmt.Println()
}
