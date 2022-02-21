package tegra

import (
	"fmt"
	"log"
	"os"
)

type Replicas struct {
	r      map[string]string
	logErr *log.Logger
}

func newReplicas(r map[string]string) *Replicas {
	return &Replicas{
		r:      r,
		logErr: log.New(os.Stdout, "[replicas:err] ", 0),
	}
}

func (r *Replicas) Get(titile string, v ...interface{}) string {
	s, ok := r.r[titile]
	if !ok {
		r.logErr.Printf("undefined title `%s`", titile)
	}
	return fmt.Sprintf(s, v...)
}
