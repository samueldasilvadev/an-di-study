package main

import (
	"fmt"
	di "godistudy/cmd/personaldi/lib"
)

type Dummy struct {
}

func (Dummy) Log(msg string) {
	fmt.Println(msg)
}

func main() {
	container := di.NewDIContainer()
	container.Provide(di.Dependency{
		Key: "test",
		Setup: func() di.Injectable {
			return Dummy{}
		},
		Lazy: true,
	})
	test := container.Inject("test").(Dummy)
	test.Log("Lazy inject!")
}
