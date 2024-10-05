package example

import "accessor/src/ss"

//go:generate go run ../accessor_engine.go Station get

type Station struct {
	VV chan int64
	BB chan *ss.HHH
}
