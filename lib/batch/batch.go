package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

type sem struct{}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	var wg sync.WaitGroup
	var mx sync.Mutex
	ch := make(chan sem, pool)
	var r []user
	for i := int64(0); i < n; i++ {
		wg.Add(1)
		ch <- struct{}{}
		go func(i int64) {
			defer wg.Done()
			user := getOne(i)
			mx.Lock()
			defer mx.Unlock()
			r = append(r, user)
			<-ch
		}(i)
	}
	wg.Wait()
	close(ch)
	return r
}
