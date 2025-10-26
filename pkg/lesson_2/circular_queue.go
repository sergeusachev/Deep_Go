// Предположим, что эта очередь будет оперировать только положительными
// числами (отрицательные числа ей никогда не поступят на вход)

/*
type CircularQueue struct { ... }

func NewCircularQueue(size int)              // создать очередь с определенным размером буффера
func (q *CircularQueue) Push(value int) bool // добавить значение в конец очереди (false, если очередь заполнена)
func (q *CircularQueue) Pop() bool           // удалить значение из начала очереди (false, если очередь пустая)
func (q *CircularQueue) Front() int          // получить значение из начала очереди (-1, если очередь пустая)
func (q *CircularQueue) Back() int           // получить значение из конца очереди (-1, если очередь пустая)
func (q *CircularQueue) Empty() bool         // проверить пустая ли очередь
func (q *CircularQueue) Full() bool          // проверить заполнена ли очередь
*/

package lesson_2

type CircularQueue struct {
    size int
	front int
	rear int
	buffer []int
}

func NewCircularQueue(size int) *CircularQueue {
    return &CircularQueue{
    	size: size,
    	front: -1,
    	rear: -1,
    	buffer: make([]int, size)
    }
}

              
func (q *CircularQueue) Push(value int) bool {
	if Full() {
		return false
	}
	
	if Empty() {
		q.front = 0
	}
	
	q.rear = (q.rear+1) % q.size
	q.buffer[q.rear] = value
	
	return true
}

func (q *CircularQueue) Pop() bool {
	if Empty() {
		return false
	}
	
	if q.front == q.rear {
		q.front = -1
		q.rear = -1
	} else {
		q.front = (q.front+1) % q.size
	}
	
	return true	
}
          
func (q *CircularQueue) Front() int {
	if Empty() {
		return -1
	}
	
	return q.buffer[q.front]
}
          
func (q *CircularQueue) Back() int {
	if Empty() {
		return -1
	}
	
	return q.buffer[q.rear]
}
          
func (q *CircularQueue) Empty() bool {
	return q.front == -1
}
        
func (q *CircularQueue) Full() bool {
	return (q.front == 0 && q.rear == q.size-1) || (q.rear+1 == q.front)
}          