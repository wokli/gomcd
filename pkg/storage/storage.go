package storage

import (
	"log"
	"strings"
)

var supportedCommands = map[string]bool{
	"GET": true,
	"SET": true,
	"DEL": true,
}

// Storage stores
type Storage struct {
	data map[string]string
}

// CreateStorage creates storage
func CreateStorage() (st *Storage) {
	data := make(map[string]string)
	return &Storage{data}
}

// Add adds
func (st *Storage) Add(key string, value string) (err error) {
	st.data[key] = value
	return nil
}

// Del deletes
func (st *Storage) Del(key string) (err error) {
	delete(st.data, key)
	return nil
}

// Get gets
func (st *Storage) Get(key string) (res string, ok bool) {
	log.Println(st.data)
	res, ok = st.data[key]
	log.Println(res, ok)
	return
}

// ProcessCommand processes command %)
func ProcessCommand(st *Storage, cmd string) (*string, bool) {
	var command, key, val string
	var err error

	words := strings.Fields(cmd)
	log.Printf("words: %#+v\n", words)

	command = words[0]
	if _, ok := supportedCommands[command]; !ok {
		log.Fatal("command not supported:", command)
		return nil, false
	}

	key = words[1]
	if len(words) == 3 {
		val = words[2]
	}

	switch command {
	case "SET":
		err = st.Add(key, val)
		if err != nil {
			log.Fatal(err)
			return nil, false
		}
		return nil, true
	case "DEL":
		err = st.Del(key)
		if err != nil {
			log.Fatal(err)
			return nil, false
		}
		return nil, true
	case "GET":
		res, ok := st.Get(key)
		if !ok {
			return nil, ok
		}
		return &res, ok
	}
	log.Fatal("not implemented:", command)
	return nil, false
}
