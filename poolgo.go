package poolgo

import (
	"fmt"
	"sync"
)

type (
	workfunc func(interface{}) interface{}
)

type Pool struct {
	Mutex sync.RWMutex
	//map of workers. Need when to call worker by name
	Workers map[string]workfunc
	//just slice of workers
	UnnamedWorkers []workfunc
	Poolnums       int
	Data           []interface{}
}

//Create basic pool
func Create(poolnum int) *Pool {
	pool := new(Pool)
	pool.Poolnums = poolnum
	pool.Workers = make(map[string]workfunc)
	pool.UnnamedWorkers = make([]workfunc, poolnum)
	return pool
}

//Append new func for this pool
func (pool *Pool) AppendFunc(name string, value func(interface{}) interface{}) {
	pool.Workers[name] = value
}

//Append new work to unnamed slice
//value is function
func (pool *Pool) AppendFuncs(value func(interface{}) interface{}) {
	for i := 0; i < pool.Poolnums; i++ {
		pool.UnnamedWorkers = append(pool.UnnamedWorkers, value)
	}
}

//Add data to workers
func (pool *Pool) AddData(data []interface{}) {
	pool.Data = data
}

func (pool *Pool) Run(name string, param interface{}) {
	for i := 0; i < pool.Poolnums; i++ {
		go func(param interface{}) {
			result := pool.UnnamedWorkers[i](param)
			fmt.Println(result)
		}(pool.Data[i])
	}
}

//Run jobs in first case - set list of params
func (pool *Pool) RunWithValues(name string, params []interface{}) {
	for i := 0; i < pool.Poolnums; i++ {
		go func(param int) {
			result := pool.Workers[name](params[param])
			fmt.Println(result)
		}(i)
	}
}

//Remove all workers in the case if worker are stopped
func (pool *Pool) RemoveAll() {

}

func (pool *Pool) Close() {
	pool.Mutex.Lock()
}
