package tegra

import (
	"log"
	"os"
	"regexp"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type HandlerFunc func(upd *Update)

type handlerData struct {
	Pattern *regexp.Regexp
	Funcs   []HandlerFunc
}

type Handler struct {
	Actions  *Actions
	Storage  *Storage
	Replicas *Replicas

	addedHandlers  []handlerData
	addedCallbacks []handlerData
	addedActions   map[string][]HandlerFunc
	cmdNotFound    []HandlerFunc

	logInfo *log.Logger
}

func NewHandler(replicas map[string]string) *Handler {
	return &Handler{
		Actions:  newActions(),
		Storage:  newStorage(),
		Replicas: newReplicas(replicas),

		addedActions: make(map[string][]HandlerFunc),
		logInfo:      log.New(os.Stdout, "[tegra:info] ", 0),
	}
}

func (h *Handler) HandleMessage(update *tgbotapi.Update) {
	var upd = newUpdate(update)
	h.logInfo.Printf("MESSAGE: @%s | `%s`", upd.SentFrom().UserName, upd.Message.Text)

	if upd.Message.IsCommand() {
		h.Actions.Clear(upd.SentID)
		h.HandleMessageCommand(upd)
	} else {
		h.HandleMessageAction(upd)
	}
}

func (h *Handler) AddCommand(regexpPattern string, handlers ...HandlerFunc) {
	pattern, err := regexp.Compile(regexpPattern)
	if err != nil {
		panic(err)
	}

	h.addedHandlers = append(h.addedHandlers, handlerData{Pattern: pattern, Funcs: handlers})
	h.logInfo.Printf("Command `%s` has been added", regexpPattern)
}

func (h *Handler) HandleMessageCommand(upd *Update) {
	for _, hand := range h.addedHandlers {
		if ok := hand.Pattern.MatchString(upd.Message.Text); ok {
			for _, f := range hand.Funcs {
				if upd.isStopped() {
					h.useCmdNotFound(upd)
					break
				}
				upd.stop()
				f(upd)
			}
			return
		}
	}
	h.useCmdNotFound(upd)
}

func (h *Handler) CommandNotFound(handlers ...HandlerFunc) {
	h.cmdNotFound = handlers
}

func (h *Handler) useCmdNotFound(upd *Update) {
	for _, f := range h.cmdNotFound {
		f(upd)
	}
}

func (h *Handler) AddAction(actionName string, handlers ...HandlerFunc) {
	h.addedActions[actionName] = handlers
}

func (h *Handler) HandleMessageAction(upd *Update) {
	for _, fn := range h.addedActions[h.Actions.Get(upd.SentID)] {
		if upd.isStopped() {
			break
		}
		upd.stop()
		fn(upd)
	}
}

func (h *Handler) AddCallback(cbRegexpPattern string, handlers ...HandlerFunc) {
	pattern, err := regexp.Compile(cbRegexpPattern)
	if err != nil {
		panic(err)
	}

	h.addedCallbacks = append(h.addedCallbacks, handlerData{Pattern: pattern, Funcs: handlers})
	h.logInfo.Printf("Command `%s` has been added", cbRegexpPattern)
}

func (h *Handler) HandleCallback(update *tgbotapi.Update) {
	var upd = newUpdate(update)
	for _, callback := range h.addedCallbacks {
		if ok := callback.Pattern.MatchString(upd.CallbackData()); ok {
			for _, fn := range callback.Funcs {
				if upd.isStopped() {
					break
				}
				upd.stop()
				fn(upd)
			}
			break
		}
	}
}
