package shutdown

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type Shutdown struct {
	funcs map[string]func() error
	mutex sync.Mutex
}

var Global = New()

func New() Shutdown {
	return Shutdown{
		funcs: make(map[string]func() error),
	}
}

func (s *Shutdown) Add(name string, fn func() error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.funcs == nil {
		s.funcs = make(map[string]func() error)
	}

	s.funcs[name] = fn
}

func (s *Shutdown) AddAnonymous(fn func() error) string {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	name := uuid.NewString()

	s.funcs[name] = fn

	return name
}

func (s *Shutdown) Delete(name string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.funcs, name)
}

func (s *Shutdown) Run() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, fn := range s.funcs {
		if err := fn(); err != nil {
			log.Err(err).Msg("shutdown error")
		}
	}
}

func (s *Shutdown) WatchCtx(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	<-ctx.Done()

	s.Run()
}
