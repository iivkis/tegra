package tegra

import (
	"log"
	"os"
	"time"
)

type Actions struct {
	actions map[int64]action
	log     *log.Logger
}

type action struct {
	Name       string
	ExpiresdIn int64
}

func newActions() *Actions {
	actions := &Actions{
		actions: make(map[int64]action),
		log:     log.New(os.Stdout, "[tegra:actions] ", 0),
	}

	actions.garbage()
	return actions
}

func (a *Actions) Set(telegramID int64, actionName string) {
	a.actions[telegramID] = action{
		Name:       actionName,
		ExpiresdIn: time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
}

func (a *Actions) Get(telegramID int64) string {
	return a.actions[telegramID].Name
}

func (a *Actions) Count() int {
	return len(a.actions)
}

func (a *Actions) Clear(telegramID int64) {
	delete(a.actions, telegramID)
}

func (a *Actions) garbage() {
	go func() {
		for t := range time.Tick(time.Hour * 24) {
			for telegramID, action := range a.actions {
				if t.Unix() >= action.ExpiresdIn {
					a.Clear(telegramID)
					a.log.Printf("autoclear %d/n", telegramID)
				}
			}
		}
	}()
}
