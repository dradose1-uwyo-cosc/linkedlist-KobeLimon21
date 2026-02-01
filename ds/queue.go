package ds

import "errors"

type Queue struct {
	list LinkedList
}

func (q *Queue) Push(value string) {
	q.list.Insert(value) // adds element at end 
}

func (q *Queue) Pop() (string, error) {
	if q.list.IsEmpty() {
		return "", errors.New("queue now empty") // queue is empty 
	}
	val := q.list.Head.data
	_ = q.list.RemoveAt(0) // removes element
	return val, nil
}
