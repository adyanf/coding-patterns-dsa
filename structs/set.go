package structs

type Set struct {
	hashMap map[interface{}]bool
}

// NewSet will initialize and return a new object of Set.
func NewSet() *Set {
	s := new(Set)
	s.hashMap = make(map[interface{}]bool)
	return s
}

// Add will add the value in the Set.
func (s *Set) Add(value interface{}) {
	s.hashMap[value] = true
}

// Delete will delete the value from the set.
func (s *Set) Delete(value interface{}) {
	delete(s.hashMap, value)
}

// Exists will check if the value exists in the set or not.
func (s *Set) Exists(value interface{}) bool {
	_, ok := s.hashMap[value]
	return ok
}
