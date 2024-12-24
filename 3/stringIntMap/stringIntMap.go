package stringIntMap

import "fmt"

type StringIntMap struct {
	x map[string]int
}

func New() *StringIntMap {
	return &StringIntMap{x: make(map[string]int)}
}

func (s *StringIntMap) Add(key string, value int) {
	s.x[key] = value
}

func (s *StringIntMap) Remove(key string) {
	delete(s.x, key)
}

func (s *StringIntMap) Copy() map[string]int {
	result := make(map[string]int)

	for k, v := range s.x {
		result[k] = v
	}
	return result
}

func (s *StringIntMap) Exists(key string) bool {
	_, ok := s.x[key]
	return ok
}

func (s *StringIntMap) Get(key string) (int, bool) {
	v, ok := s.x[key]
	return v, ok
}

func (s *StringIntMap) Print() {
	fmt.Println(s.x)
}
