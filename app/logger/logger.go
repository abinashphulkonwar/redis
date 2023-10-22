package logger

import (
	"os"
	"sync"

	"github.com/abinashphulkonwar/redis/storage"
)

type Logger struct {
	file      *os.File
	file_path string
	wg        *sync.WaitGroup
	queue     *storage.Queue
}

func Init(path string) *Logger {
	var wg sync.WaitGroup
	queue := storage.InitQueue()
	return &Logger{
		file:      nil,
		file_path: path,
		wg:        &wg,
		queue:     queue,
	}
}

func (l *Logger) open() {
	file, err := os.Open(l.file_path)
	if err != nil {
		panic(err)
	}
	l.file = file

}

func (l *Logger) close() {
	if l.file != nil {
		l.file.Close()
	}
}

func (l *Logger) start() {

	defer l.wg.Done()
	for {

		node, isFound := l.queue.Get()
		if !isFound {
			continue
		}
		curruent := node.(int)
		println("🚀", curruent)

	}

}

func (l *Logger) New() {

	l.open()
	l.wg.Add(1)
	go l.start()
	l.wg.Wait()
	l.close()

}
