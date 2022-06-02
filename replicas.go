package tegra

import (
	"fmt"
)

type Replicas map[string]string

func NewReplicas(r Replicas) Replicas {
	return r
}

func (r Replicas) Get(title string, v ...interface{}) string {
	s, ok := r[title]
	if !ok {
		fmt.Printf("undefined title `%s`", title)
		return ""
	}
	return fmt.Sprintf(s, v...)
}
