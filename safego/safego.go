package safego

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

const (
	infinite = -999
)

// Run runs goroutine safely and attempts to recover safely in the event of an exception
func Run(f func()) {
	RunWithFailOver(f, infinite)
}

// RunGoWithFailOver runs goroutine safely and attempts to recover safely in the event of an exception
func RunWithFailOver(f func(), failOverCount int) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				stack := make([]byte, 1024*8)
				stack = stack[:runtime.Stack(stack, false)]
				stackStr := string(stack)
				info := fmt.Sprintf("%v", err)
				log.Println(info, stackStr, "fail over remain count:", failOverCount)
				time.Sleep(time.Second)

				if failOverCount == infinite {
					RunWithFailOver(f, failOverCount)
				} else if failOverCount > 0 {
					failOverCount--
					RunWithFailOver(f, failOverCount)
				}
			}
		}()

		f()
	}()
}
