package logger

import (
	log "github.com/inconshreveable/log15"
)

type (
	Logger interface {
		log.Logger
	}

	Options struct {
		Verbose bool
	}
)

func New(opt *Options) Logger {
	l := log.New()
	lvl := log.LvlDebug
	handlers := []log.Handler{}
	verboseHandler := log.LvlFilterHandler(lvl, log.StdoutHandler)
	handlers = append(handlers, verboseHandler)
	l.SetHandler(log.MultiHandler(handlers...))
	return l
}
