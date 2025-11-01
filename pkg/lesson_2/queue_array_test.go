package lesson_2

import(
	//"fmt"
	"testing"
)

func TestNewQueue(t *testing.T) {
	result := NewQueue()
	
	if !(result.size == 5 && result.rear == -1 && result.front == -1) {
		t.Error("Creation is not correct.")
	}
}

func TestEnqueue(t *testing.T) {
	q := NewQueue()
	
	q.Enqueue(13)
	q.Enqueue(2)
	err := q.Enqueue(3)
	if err != nil {
		t.Error("err:", err)
	}
	frontValue, err := q.Peek()
	if frontValue != 13 || err != nil {
		t.Error("Peek is not correct")
	}
}