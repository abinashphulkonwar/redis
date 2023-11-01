package internalstorage

import "sync"

type QNode struct {
	Next *QNode

	Value interface{}
}

type Queue struct {
	head   *QNode
	tail   *QNode
	mutex  sync.Mutex
	Length int
}

func (q *Queue) Insert(value interface{}) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	node := &QNode{
		Value: value,
		Next:  nil,
	}

	if q.Length == 0 {
		q.Length++
		q.head = node
		q.tail = node
		return
	}
	q.Length++

	q.tail.Next = node
	q.tail = node

}

func (q *Queue) Get() (interface{}, bool) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.Length == 0 {
		return nil, false
	}

	return q.head.Value, true
}

func (q *Queue) Remove() bool {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.Length == 0 {
		return false
	}

	if q.Length == 1 {
		q.Length--
		q.head = nil
		q.tail = nil
		return true
	}

	q.Length--

	q.head = q.head.Next
	return true

}

func InitQueue() *Queue {
	return &Queue{
		head:   nil,
		Length: 0,
		tail:   nil,
		mutex:  sync.Mutex{},
	}

}
