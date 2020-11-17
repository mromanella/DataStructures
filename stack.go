package datastructures

// Stack A last in first out stack
type Stack struct {
	ll LinkedList
}

// Push Pushes value to the top of the stack
func (s *Stack) Push(value interface{}) {
	s.ll.Add(value)
}

// Pop Pops the value from the top of the stack or nil if there is nothing in the stack
func (s *Stack) Pop() interface{} {
	n, err := s.ll.Remove(s.ll.Length() - 1)
	if err != nil {
		return n.Value
	}
	return nil
}
