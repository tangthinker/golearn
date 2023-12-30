package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

type config struct {
	Website struct {
		Title         string
		Comment       string
		UpdateCounter int64
	}
	mutex       sync.RWMutex
	filePointer *os.File
}

func (conf *config) init() {
	conf.filePointer, _ = os.OpenFile("/Users/shanliao/code/GoProject/golearn/concurrency/rwmutex/application/config.txt", os.O_RDWR, 0777)
}

func (conf *config) readFile() {
	buff := make([]byte, 1024)
	n, _ := conf.filePointer.Read(buff)
	_ = json.Unmarshal(buff[:n], conf)
}

func (conf *config) getConfig() config {
	conf.mutex.RLock()
	defer conf.mutex.RUnlock()
	conf.readFile()
	result := config{
		Website: struct {
			Title         string
			Comment       string
			UpdateCounter int64
		}{Title: conf.Website.Title, Comment: conf.Website.Comment, UpdateCounter: conf.Website.UpdateCounter},
	}
	return result
}

func (conf *config) updateConfig(newConf *config) {
	conf.mutex.Lock()
	defer conf.mutex.Unlock()
	b, _ := json.Marshal(newConf)
	_, _ = conf.filePointer.Seek(0, 0)
	_, _ = conf.filePointer.Write(b)
	_, _ = conf.filePointer.Seek(0, 0)
}

func main() {
	var conf config
	conf.init()
	newConf := &config{
		Website: struct {
			Title         string
			Comment       string
			UpdateCounter int64
		}{Title: "the World of Shanliao", Comment: "Welcome!", UpdateCounter: 0},
	}
	conf.updateConfig(newConf)
	fmt.Println(conf.getConfig())

	go func() {
		add := int64(1)
		for {
			time.Sleep(5 * time.Second)
			newConf.Website.UpdateCounter = add
			conf.updateConfig(newConf)
			add++
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println(conf.getConfig())
		}
	}()

	time.Sleep(16 * time.Second)

}
