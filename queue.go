package datastructures

// Queue A simple queue with enqueue and dequeue ability
type Queue struct {
	ll LinkedList
}

// Enqueue Add the value to the back of the queue
func (q *Queue) Enqueue(value interface{}) {
	q.ll.Add(value)
}

// Dequeue Remove and return the value or nil at the front of the queue
func (q *Queue) Dequeue() interface{} {
	n, err := q.ll.Remove(0)
	if err != nil {
		return n.Value
	}
	return nil
}
