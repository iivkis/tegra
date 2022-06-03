package tegra

import (
	"fmt"
)

type Replicas struct {
	m map[string]string
}

func NewReplicas(m map[string]string) *Replicas {
	return &Replicas{m}
}

func (r Replicas) Get(title string, v ...interface{}) string {
	s, ok := r.m[title]
	if !ok {
		fmt.Printf("undefined title `%s`", title)
		return ""
	}
	return fmt.Sprintf(s, v...)
}
