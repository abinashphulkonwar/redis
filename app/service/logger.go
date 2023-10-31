package service

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
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

func (l *Logger) ReadLogs() {
	l.open()
	defer l.close()
	start_buf := make([]byte, 18)
	n, err := l.file.Read(start_buf)
	if err != nil {
		panic(err)
	}
	position := n
	current_buf := make([]byte, 10000)
	copy(current_buf[0:18], start_buf)
	for {
		buf := make([]byte, 200)
		n, err := l.file.Read(buf)
		if err != nil {
			println(err.Error())
			break
		}
		index_buf := bytes.Index(buf, start_buf)
		if index_buf != -1 {
			copy(current_buf[position:position+index_buf], buf)
			l.read(&current_buf)

			current_buf = make([]byte, 10000)
			copy(current_buf[0:], buf[index_buf:])

			position = n - index_buf
		} else {
			copy(current_buf[position:position+n], buf)
			position = position + n
		}
	}
	l.read(&current_buf)

}

func (l *Logger) read(current_buf *[]byte) {

	dec := gob.NewDecoder(bytes.NewBuffer(*current_buf))

	var log Log
	err := dec.Decode(&log)
	if err == io.EOF {
		println(err.Error())
	} else if err != nil {
		panic(err)
	}
	fmt.Printf("Log: Time=%v Path=%s Status=%s Method=%s Command=%s Key=%s Value=%s\n", log.Time, log.Path, log.Status, log.Method, log.Command, log.Key, log.Value)

}

func (l *Logger) ReadLacacy() {
	l.open()
	defer l.close()
	buf := make([]byte, 100)

	n, err := l.file.Read(buf)
	if err != nil {
		panic(err)
	}
	println(n)
	start := string(buf[0:18])

	itr := -1
	log_buf := make([]byte, 4000)
	index := 0
	isEnd := false
	isStart := true
	bytes_read := n
	isNot_Ending_Log := false
	currentLenght := 0
	for {
		if itr == 3 {
			break
		}
		if n <= 0 || isEnd {
			break
		}
		isDone := false
		itr++

		for {

			length := index + 18
			if isNot_Ending_Log {
				max := currentLenght
				if currentLenght < n {
					max = n
				}
				if index >= max {
					break
				}
			} else {
				if index >= n {
					break
				}
			}
			if length > n {
				log_buf[index] = buf[index]
				index++
				continue
			}
			if itr == 1 {
				println(string(buf[index:length]), start)
			}

			if string(buf[index:length]) == start {
				println("done", isDone, n, index)
			}
			if string(buf[index:length]) == start && !isStart {
				isDone = true
				isNot_Ending_Log = false
				break
			}
			isStart = false
			isNot_Ending_Log = true
			log_buf[index] = buf[index]
			index++
			currentLenght++
		}
		buf = buf[index:]
		n = n - index
		if !isDone {
			var a [100]byte
			buf = a[0:]

			n_size, err := l.file.ReadAt(buf, int64(bytes_read))
			if err != nil {
				if err == io.EOF {
					isEnd = true
				} else {
					panic(err)
				}

			} else {
				println("size", n_size, string(buf))
				n = n_size
				bytes_read = bytes_read + n
				continue
			}
		}

		println("data")
		dec := gob.NewDecoder(bytes.NewBuffer(log_buf))
		var log Log
		err := dec.Decode(&log)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Printf("Log: Time=%v Path=%s Status=%s Method=%s Command=%s Key=%s Value=%s\n", log.Time, log.Path, log.Status, log.Method, log.Command, log.Key, log.Value)
		println(n)
		index = 0
		var a [4000]byte
		log_buf = a[0:]
		isStart = true
		currentLenght = 0

	}

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
