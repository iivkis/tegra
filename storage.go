package tegra

type Storage map[string]interface{}

func NewStorage() Storage {
	return make(map[string]interface{})
}

func (s Storage) Set(name string, v interface{}) {
	s[name] = v
}

func (s Storage) Get(name string) (interface{}, bool) {
	v, ex := s[name]
	return v, ex
}

func (s Storage) MustGet(name string) interface{} {
	return s[name]
}
