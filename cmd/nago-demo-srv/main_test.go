package main

import (
	"testing"
	"time"
)

func Test_create(t *testing.T) {
	app := create()
	go func() {
		time.Sleep(10 * time.Second)
		app.Stop()
	}()

	app.Run()
}
