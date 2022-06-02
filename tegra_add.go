package tegra

import (
	"log"
	"regexp"
)

func (h *TegraHandler) AddCommand(cmdRegexpPattern string, handlers ...HandlerFunc) {
	pattern, err := regexp.Compile(cmdRegexpPattern)
	if err != nil {
		panic(err)
	}

	log.Printf("Command `%s` has been added", cmdRegexpPattern)
	h.allHandlers = append(h.allHandlers, &Handler{Pattern: pattern, Funcs: handlers})
}

func (h *TegraHandler) AddCallback(cbRegexpPattern string, handlers ...HandlerFunc) {
	pattern, err := regexp.Compile(cbRegexpPattern)
	if err != nil {
		panic(err)
	}

	log.Printf("Command `%s` has been added", cbRegexpPattern)
	h.allCallbacks = append(h.allCallbacks, &Handler{Pattern: pattern, Funcs: handlers})
}

func (h *TegraHandler) AddAction(name string, handlers ...HandlerFunc) {
	h.allActions[name] = handlers
}

func (h *TegraHandler) AddNotFound(handlers ...HandlerFunc) {
	h.notFoundFuncs = handlers
}
