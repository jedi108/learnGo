package recover

import (
	"log"
	"testing"
)

func TestRecover(t *testing.T) {
	defer func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("%#v", r)
			}
		}()

		if r := recover(); r != nil {
			log.Printf("%#v", r)
		}

		panic(2)
	}()

	panic(1)
}
