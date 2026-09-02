package shutdown

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
)

type Shutdown struct {
	funcs  map[string]func() error
	nextID uint64
	mutex  sync.Mutex
}

var Global = New()

func New() Shutdown {
	return Shutdown{funcs: make(map[string]func() error)}
}

func (s *Shutdown) Add(name string, fn func() error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.ensureMap()
	s.funcs[name] = fn
}

func (s *Shutdown) AddAnonymous(fn func() error) string {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.ensureMap()

	for {
		s.nextID++
		name := fmt.Sprintf("anonymous-%d", s.nextID)
		if _, exists := s.funcs[name]; !exists {
			s.funcs[name] = fn
			return name
		}
	}
}

func (s *Shutdown) Delete(name string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.funcs, name)
}

func (s *Shutdown) Run() {
	s.mutex.Lock()
	funcs := make([]func() error, 0, len(s.funcs))
	for _, fn := range s.funcs {
		funcs = append(funcs, fn)
	}
	s.mutex.Unlock()

	for _, fn := range funcs {
		if err := fn(); err != nil {
			slog.Error("shutdown error", slog.String("error", err.Error()))
		}
	}
}

func (s *Shutdown) WatchCtx(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	<-ctx.Done()
	s.Run()
}

func (s *Shutdown) ensureMap() {
	if s.funcs == nil {
		s.funcs = make(map[string]func() error)
	}
}
