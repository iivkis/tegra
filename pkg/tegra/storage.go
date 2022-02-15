package tegra

type Storage struct {
	s map[string]interface{}
}

func newStorage() *Storage {
	return &Storage{
		s: make(map[string]interface{}),
	}
}

func (s *Storage) Set(name string, v interface{}) {
	s.s[name] = v
}

func (s *Storage) Get(name string) interface{} {
	return s.s[name]
}
