package sync

import "log"

//Go 保护携程panic
func Go(x func()) {
	go func() {
		defer func() {
			if e := recover(); e != nil {
				log.Printf("goroutine panic:%v\n", e)
			}
		}()
		x()
	}()
}
