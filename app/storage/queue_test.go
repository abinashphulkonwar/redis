package storage_test

import (
	"sync"
	"testing"

	"github.com/abinashphulkonwar/redis/storage"
)

func TestQueue(T *testing.T) {
	queue := storage.InitQueue()
	asysc := sync.WaitGroup{}
	asysc.Add(2)
	var queue_value []int
	mx := sync.Mutex{}

	go func() {
		defer asysc.Done()

		for {
			mx.Lock()

			data, isFound := queue.Get()
			if isFound {
				queue.Remove()
				switch v := data.(type) {
				case int:
					queue_value = append(queue_value, v)
					if v == 1000 {
						mx.Unlock()

						return
					}

				}
			}
			mx.Unlock()
		}

	}()
	for i := 1; i <= 1000; i++ {
		queue.Insert(i)
	}
	go func() {
		defer asysc.Done()
		for {
			mx.Lock()
			if len(queue_value) == 1000 {

				println("length :", len(queue_value))

				for i := 1; i <= 1000; i++ {
					if i != queue_value[i-1] {
						println(i, queue_value[i-1])
						mx.Unlock()
						T.Errorf("un ordered")
						return
					}
				}
				println("done")
				return
			}
			mx.Unlock()

		}
	}()

	asysc.Wait()

}
