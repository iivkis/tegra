package tegra

import (
	"regexp"
)

type HandlerFunc func(upd *Update)

type Handler struct {
	Pattern *regexp.Regexp
	Funcs   []HandlerFunc
}

type HandlerActions map[string][]HandlerFunc

type TegraHandlerConfig struct {
	Replicas       map[string]string
	ActionLifetime int64 //in milliseconds
}

type TegraHandler struct {
	Local         *Storage
	Replicas      *Replicas
	ActionStorage *ActionStorage

	allHandlers  []*Handler
	allCallbacks []*Handler
	allActions   HandlerActions

	notFoundFuncs []HandlerFunc
}

func NewTegraHandler(config *TegraHandlerConfig) *TegraHandler {
	return &TegraHandler{
		Local:         NewStorage(),
		Replicas:      NewReplicas(config.Replicas),
		ActionStorage: NewActions(config.ActionLifetime),
		allActions:    make(HandlerActions),
	}
}

func (h *TegraHandler) ExecFuncs(upd *Update, funcs []HandlerFunc) {
	for _, fn := range funcs {
		if upd.Stopped() {
			h.UseNotFound(upd)
			return
		}
		upd.stop()
		fn(upd)
	}
}

func (h *TegraHandler) UseNotFound(upd *Update) {
	h.ExecFuncs(upd, h.notFoundFuncs)
}
