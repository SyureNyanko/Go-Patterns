package main

import (
	"fmt"
)

type Object struct {
	s int
}
type Pool chan *Object

func (obj *Object) Do() {
	fmt.Printf("Hello World! %d \n", obj.s)
}

/* generate -> queueing -> waiting */
func New(total int) chan *Object {
	p := make(chan *Object, total)

	for i := 0; i < total; i++ {
		/* ex. Heavy initialization/Generate */
		p <- &Object{s: i}
		//p <- new(Object{s: i})
	}
	return p
}

func main() {
	p := New(5)
	for i := 0; i < 10; i++ {
		select {
		case obj := <-p:
			obj.Do()

			/* return if Done */
			p <- obj
		default:
			return
		}
	}

}
