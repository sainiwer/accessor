// Code generated by go generate ; DO NOT EDIT.

package example

import "accessor/src/ss"

func (p *Station) GetVV() chan int64 {
	return p.VV
}

func (p *Station) GetBB() chan *ss.HHH {
	return p.BB
}
