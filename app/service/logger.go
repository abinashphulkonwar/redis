package service

import (
	"bytes"
	"encoding/gob"
	"fmt"
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
	if l.file != nil {
		return
	}
	file, err := os.OpenFile(l.file_path, os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	}
	l.file = file

}

func (l *Logger) Read() {
	l.open()
	defer l.close()
	buf := make([]byte, 10000)

	_, err := l.file.Read(buf)
	if err != nil {
		panic(err)
	}

	dec := gob.NewDecoder(bytes.NewBuffer(buf))
	var log Log
	err = dec.Decode(&log)
	if err != nil {
		panic(err)
	}
	println(log.Key, log.Value)

}

func (l *Logger) close() {
	if l.file != nil {
		l.file.Close()
	}
	l.file = nil
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

		current := node.(*Log)
		buf := new(bytes.Buffer)
		enc := gob.NewEncoder(buf)
		if err := enc.Encode(*current); err != nil {
			fmt.Println(err)
		}

		println(current.Key)
		size, err := l.file.Write(buf.Bytes())
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
