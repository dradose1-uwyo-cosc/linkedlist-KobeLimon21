package ds

type Stack struct {
	list LinkedList
}

func (s *Stack) Push(value string) {
	_ = s.list.InsertAt(0, value) // adds head at the front 
}

func (s *Stack) Pop() (string, bool) {
	if s.list.IsEmpty() {
		return "", false // stack is empty 
	}
	val := s.list.Head.data
	_ = s.list.RemoveAt(0) // removes element
	return val, true
}
