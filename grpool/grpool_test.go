package grpool

import (
	"bytes"
	"context"
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"testing"
	"time"
)

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func TestNew(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()
	wg := &sync.WaitGroup{}
	pool, err := New(ctx, 4, 1000)
	if err != nil {
		t.Error(err)
		return
	}

	f := func() interface{} {
		defer wg.Done()
		time.Sleep(time.Second)
		gid := getGID()
		gidStr := fmt.Sprintf("GID: %d", gid)
		t.Log(gidStr)
		return gidStr
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		pool.Exec(f)
	}

	wg.Wait()
}
