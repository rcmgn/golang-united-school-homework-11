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

type sem struct{}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	var wg sync.WaitGroup
	ch := make(chan sem, pool)
	var r userCount
	for i := int64(0); i < n; i++ {
		wg.Add(1)
		ch <- struct{}{}
		go func(i int64) {
			user := getOne(i)
			r.Mutex.Lock()
			defer r.Mutex.Unlock()
			r.u = append(r.u, user)
			<-ch
			wg.Done()
		}(i)
	}
	wg.Wait()
	close(ch)
	return r.u
}
