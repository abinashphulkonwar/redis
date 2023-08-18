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
	go func() {

		for {

			data, isFound := queue.Get()

			if isFound {
				queue.Remove()
				switch v := data.(type) {
				case int:
					println(v)
					queue_value = append(queue_value, v)
					if v == 1000 {

						asysc.Done()

					}

				}
			}

		}

	}()
	mx := sync.Mutex{}
	go func() {
		for {

			if len(queue_value) == 1000 {
				mx.Lock()

				println("length :", len(queue_value))

				for i := 1; i <= 1000; i++ {
					if i != queue_value[i-1] {
						println(i, queue_value[i-1])
						mx.Unlock()
						asysc.Done()
						T.Errorf("un ordered")
					}
				}
				println("done")
				asysc.Done()
				mx.Unlock()

			}

		}
	}()
	for i := 1; i <= 1000; i++ {
		queue.Insert(i)
	}
	asysc.Wait()

}
