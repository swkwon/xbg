package grpool

import (
	"context"
	"errors"
	"fmt"
	"log"
	"runtime"
	"time"
)

type (
	WorkerFunction func() interface{}

	Pool interface {
		Exec(wf WorkerFunction) Result
	}

	Result interface {
		Get(timeout time.Duration) (interface{}, error)
	}

	resultImpl struct {
		Ch chan interface{}
	}

	pool struct {
		// channel
		jobQueue chan func()
		ctx      context.Context
	}
)

func (r *resultImpl) Get(timeout time.Duration) (interface{}, error) {
	if timeout <= 0 {
		return <-r.Ch, nil
	} else {
		select {
		case res := <-r.Ch:
			return res, nil
		case <-time.After(timeout):
			return nil, errors.New("timeout")
		}
	}
}

func (p *pool) Exec(wf WorkerFunction) Result {
	result := &resultImpl{
		Ch: make(chan interface{}, 2),
	}
	p.jobQueue <- func() {
		result.Ch <- wf()
	}

	return result
}

func (p *pool) start() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				stack := make([]byte, 1024*8)
				stack = stack[:runtime.Stack(stack, false)]
				stackStr := string(stack)
				info := fmt.Sprintf("%v", err)
				log.Println(info, stackStr)
				time.Sleep(time.Second)
				p.start()
			}
		}()

		for {
			select {
			case job := <-p.jobQueue:
				job()
			case <-p.ctx.Done():
				return
			}
		}
	}()
}

// New makes pool object
func New(ctx context.Context, poolSize, queueSize int) (Pool, error) {
	if poolSize < 1 {
		return nil, fmt.Errorf("invalid pool size: %d", poolSize)
	}

	if queueSize < 1 {
		return nil, fmt.Errorf("invalid queue size: %d", queueSize)
	}

	p := &pool{
		jobQueue: make(chan func(), queueSize),
		ctx:      ctx,
	}

	for i := 0; i < poolSize; i++ {
		p.start()
	}

	return p, nil
}
