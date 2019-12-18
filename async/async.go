package async

import (
	"errors"
	"time"
	"xbg/safego"
)

type (
	// Function is the type of the function to run asynchronously.
	Function func() interface{}

	// Result is an object that can receive the result of asynchronous execution.
	Result interface {
		Get(timeout time.Duration) (interface{}, error)
	}

	asyncResultImpl struct {
		Ch chan interface{}
	}
)

func (ari *asyncResultImpl) Get(timeout time.Duration) (interface{}, error) {
	if timeout <= 0 {
		return <-ari.Ch, nil
	} else {
		select {
		case res := <-ari.Ch:
			return res, nil
		case <-time.After(timeout):
			return nil, errors.New("timeout")
		}
	}
}

func newAsync() *asyncResultImpl {
	return &asyncResultImpl{Ch: make(chan interface{}, 2)}
}

func async(af Function) Result {
	ari := newAsync()
	safego.Run(func() {
		ari.Ch <- af()
	})
	return ari
}

// Run The Async function executes the parameter function asynchronously
// and returns a result
func Run(af Function) Result {
	return async(af)
}
