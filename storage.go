package tegra

type Storage struct {
	m map[string]interface{}
}

func NewStorage() *Storage {
	return &Storage{make(map[string]interface{})}
}

func (s Storage) Set(name string, v interface{}) {
	s.m[name] = v
}

func (s Storage) Get(name string) (interface{}, bool) {
	v, ex := s.m[name]
	return v, ex
}

func (s Storage) MustGet(name string) interface{} {
	return s.m[name]
}
