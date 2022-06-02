package tegra

import (
	"fmt"
	"time"
)

type Action struct {
	Name      string
	CreatedAt int64
}

type ActionStorage map[int64]Action

//lifetime - action lifetime (in milliseconds)
func NewActions(lifetime int64) ActionStorage {
	a := make(ActionStorage)

	if lifetime == 0 {
		lifetime = 24 * 60 * 60 * 1000 //1 day
	}

	a.Garbage(lifetime)
	return a
}

func (a ActionStorage) Set(tgID int64, name string) {
	a[tgID] = Action{
		Name:      name,
		CreatedAt: time.Now().UnixMilli(),
	}
}

func (a ActionStorage) Get(tgID int64) string {
	return a[tgID].Name
}

func (a ActionStorage) Count() int {
	return len(a)
}

func (a ActionStorage) Clear(tgID int64) {
	delete(a, tgID)
}

func (a ActionStorage) Garbage(lifetime int64) {
	go func() {
		for t := range time.Tick(time.Millisecond * time.Duration(lifetime)) {
			n := lifetime - t.UnixMilli()
			for tgID, action := range a {
				if action.CreatedAt-n <= 0 {
					a.Clear(tgID)
					fmt.Printf("clear action with tgID `%d`", tgID)
				}
			}
		}
	}()
}
