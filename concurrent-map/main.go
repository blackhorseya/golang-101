package main

// AccessCounter is a struct that holds the access counter.
type AccessCounter struct {
	visits map[string]int
}

// NewAccessCounter is a constructor that returns a new AccessCounter.
func NewAccessCounter() *AccessCounter {
	return &AccessCounter{
		visits: make(map[string]int),
	}
}

// Visit is a method that increments the access counter for the given username.
func (c *AccessCounter) Visit(username string) {
	// todo: 2024/3/30|sean|Implement the Visit method
	panic("implement me")
}

// GetVisit is a method that returns the access counter for the given username.
func (c *AccessCounter) GetVisit(username string) int {
	// todo: 2024/3/30|sean|Implement the GetVisit method
	panic("implement me")
}

func main() {
}
