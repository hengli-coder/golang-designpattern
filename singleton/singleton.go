package singleton

import "sync"

type single struct {
}

var once sync.Once
var singleInstance *single

func getSingleInstance() *single {
	if singleInstance == nil {
		once.Do(func() {
			singleInstance = &single{}
		})
	}

	return singleInstance
}
