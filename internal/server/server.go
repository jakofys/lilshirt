package server

import (
	"errors"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"sync"
	"sync/atomic"
	"syscall"

	"github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
)

// Server is http.Server wrapper that extend feature and
// adapt to the StrictServerInterface
// can manage soft shutdown with recovery function
type Server struct {
	srv    *http.Server
	quit   chan struct{}
	locked *atomic.Bool
	funcs  map[string]OnShudownFunc
	err    chan error
}

// OnShudownFunc is func that will be executed when server stop whatever the reason
type OnShudownFunc func() error

func New(handler StrictServerInterface, middlewares []nethttp.StrictHTTPMiddlewareFunc, opts ...Option) (*Server, error) {
	srv := &http.Server{
		Handler: Handler(NewStrictHandler(handler, middlewares)),
	}
	for _, opt := range opts {
		err := opt(srv)
		if err != nil {
			return nil, err
		}
	}
	return &Server{
		srv:    srv,
		quit:   make(chan struct{}),
		funcs:  make(map[string]OnShudownFunc),
		locked: new(atomic.Bool),
	}, nil
}

// OnShutdown declared all functions that will be run when server stop
// timeout is set to 3 seconds, then return error
func (s *Server) OnShudown(funcs ...OnShudownFunc) error {
	if s.locked.Load() {
		return ErrRunningServer
	}
	for _, f := range funcs {
		name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		_, ok := s.funcs[name]
		if ok {
			return ErrAlreadyDeclareOnShutdownFunction(name)
		}
		s.funcs[name] = f
	}
	return nil
}

func (s *Server) Run(addr string) error {
	s.locked.Store(true)
	defer s.locked.Store(false)
	go s.onShutdown()
	s.srv.Addr = addr
	err := s.srv.ListenAndServe()
	<-s.quit
	errs := <-s.err
	errs = errors.Join(errs, err)
	return errs
}

func (s *Server) onShutdown() {
	notifier := make(chan os.Signal, 1)
	signal.Notify(notifier, syscall.SIGABRT, syscall.SIGINT, syscall.SIGTERM)
	<-notifier
	var errs error
	var wg sync.WaitGroup
	for _, f := range s.funcs {
		wg.Add(1)
		go func(f OnShudownFunc) {
			defer wg.Done()
			err := f()
			if err != nil {
				errs = errors.Join(errs, err)
			}
		}(f)
	}
	wg.Wait()
	s.err <- errs
	s.quit <- struct{}{}
}
