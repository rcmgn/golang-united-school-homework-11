package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

type userCount struct {
	sync.Mutex
	u []user
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	//var lock sync.Mutex
	//var wg sync.WaitGroup
	ch := make(chan user, pool)
	//var r userCount
	var r []user
	for i := int64(0); i < n; i++ {
		//wg.Add(1)
		go func(i int64) {
			user := getOne(i)
			//r.Mutex.Lock()
			//defer r.Mutex.Unlock()
			//r.u = append(r.u, user)
			ch <- user
			//wg.Done()
		}(i)
	}
	//wg.Wait()
	for i := int64(0); i < n; i++ {
		var m user
		m = <-ch
		//fmt.Println(m)
		r = append(r, m)

	}
	close(ch)
	return r
}
