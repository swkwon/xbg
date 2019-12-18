package async

import (
	"testing"
	"time"
)

func TestAsync(t *testing.T) {
	f := func() interface{} {
		time.Sleep(time.Second * 3)
		return "hello world"
	}
	result1 := Run(f)
	if ret, e := result1.Get(0); e != nil {
		t.Error(e)
		return
	} else {
		t.Log(ret.(string))
	}

	result2 := Run(f)
	if _, e := result2.Get(time.Second * 2); e == nil {
		t.Error("e must not nil")
	} else {
		t.Log(e)
	}
}
