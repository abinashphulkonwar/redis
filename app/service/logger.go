package service

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/abinashphulkonwar/redis/storage"
)

type Log struct {
	Time    string
	Status  string
	Path    string
	Method  string
	Command string
	Key     string
	Value   string
}

type Logger struct {
	file      *os.File
	file_path string
	wg        *sync.WaitGroup
	queue     *storage.Queue
}

func InitLogger(path string) *Logger {
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
	file, err := os.OpenFile(l.file_path, os.O_APPEND, 0666)

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

	mode := os.Getenv("mode")
	val := 0
	for {
		if mode == "Test" {
			if val == 1000 {
				return
			}
			val++
		}
		node, isFound := l.queue.Get()
		if !isFound {
			continue
		}
		val = 0
		curruent := node.(*Log)
		data, err := json.Marshal(&curruent)
		if err != nil {
			println(err.Error())
		}
		println(curruent.Key)
		size, err := l.file.Write(data)
		if err != nil {
			println(err.Error())
		}
		println(size)
		l.queue.Remove()
	}

}

func (l *Logger) Add(log *Log) {
	l.queue.Insert(log)
}

func (l *Logger) New() {

	l.open()
	l.wg.Add(1)
	go l.start()
	l.wg.Wait()
	l.close()

}
