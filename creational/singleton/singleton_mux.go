package main

import (
	"fmt"
	"sync"

	"github.com/gorilla/mux"
)

var muxSingleton *mux.Router

func createMux() *mux.Router {
	var lock sync.Mutex

	if muxSingleton == nil {

		lock.Lock()
		muxSingleton = mux.NewRouter()
		lock.Unlock()

		fmt.Println("Creating single instance now.")
		return muxSingleton
	}

	fmt.Println("Single instance already created.")
	return muxSingleton
}
