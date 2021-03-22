package cache

type SimpleCache struct {
	message string
}

func New() *SimpleCache {
	return &SimpleCache{}
}

func (s *SimpleCache) Message() string {
	if s.message == "" {
		s.message = "Hello from cache"
		return ""
	}
	return s.message
}
