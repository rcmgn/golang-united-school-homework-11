package batch

import (
	"fmt"
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	var lock sync.Mutex
	var wg sync.WaitGroup
	ch := make(chan struct{}, pool)
	var r []user
	for i := int64(0); i < n; i++ {
		ch <- struct{}{}
		wg.Add(1)
		go func(r1 *[]user) []user {
			user := getOne(i)
			fmt.Println(user)
			lock.Lock()
			r = append(r, user)
			lock.Unlock()
			<-ch
			wg.Done()
			return *r1
		}(&r)
	}
	wg.Wait()
	close(ch)
	return r
}
