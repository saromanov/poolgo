package poolgo

import
(
	"sync"
)

type
(
	workfunc func(interface{}) interface{}
)

type Pool struct{
	Mutex sync.RWMutex
	//map of workers. Need when  to call worker by name
	Workers map[string] workfunc
	Poolnums int
}

//Create basic pool
func Create (poolnum int) (*Pool){
	pool := new(Pool)
	pool.Poolnums = poolnum
	pool.Workers = make(map[string]workfunc)
	return pool
}

//Append new func for this pool
func (pool*Pool) AppendFunc(name string, value func(interface{}) interface{}){
	pool.Workers[name] = value
}
//Append new job
func (pool*Pool) Append(value interface{}){
	pool.Mutex.Lock()
	defer pool.Mutex.Unlock()
}

func (pool*Pool) Run(name string, param interface{}){

}

//Run jobs in first case - set list of params
func(pool*Pool) RunWithValues(name string, params[] interface{}){
	for i := 0; i < pool.Poolnums; i++ {
		go func(param int){
			result := pool.Workers[name](params[param])
		}(i)
	}
}

func (pool*Pool) Close(){
	pool.Mutex.Lock()
}