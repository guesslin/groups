package main

import (
	"fmt"

	"github.com/guesslin/groups"
)

func main() {
	fmt.Println("Start")
	p := groups.NewPanic()
	defer fmt.Println("End")
	p.Go(func() error { return fmt.Errorf("fake") })
	p.Go(func() error { return nil })

	p.Wait()
	/*
		execute output:
		Start
		panic group runs into error: fake
		End
	*/
}
