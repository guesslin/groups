package groups

import (
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

var defaultHandler func(interface{}) = func(raw interface{}) {
	if raw == nil {
		fmt.Println("no panic invoked")
	}
	switch raw.(type) {
	case error:
		fmt.Println("panic group runs into error:", raw.(error))
	}
}

type panicGroup struct {
	eg     *errgroup.Group
	lock   *sync.Mutex
	queue  []Callback
	handle func(interface{})
}

func NewPanic() *panicGroup {
	return &panicGroup{
		eg:     &errgroup.Group{},
		lock:   &sync.Mutex{},
		queue:  make([]Callback, 0),
		handle: defaultHandler,
	}
}

func (p *panicGroup) Go(c Callback) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.queue = append(p.queue, c)
}

func (p *panicGroup) Wait() error {
	defer func() {
		p.handle(recover())
	}()
	for i := range p.queue {
		p.eg.Go(p.queue[i])
	}
	if err := p.eg.Wait(); err != nil {
		panic(err)
	}
	return nil
}
