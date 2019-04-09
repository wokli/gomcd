package storage

import (
	"errors"
	"log"
	"strings"
)

const (
	emptyStr = ""
	okStr    = "OK"
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
func (st *Storage) Add(key string, value string) error {
	st.data[key] = value
	return nil
}

// Del deletes
func (st *Storage) Del(key string) error {
	if _, ok := st.data[key]; ok {
		delete(st.data, key)
		return nil
	}
	return errors.New("not found")
}

// Get gets
func (st *Storage) Get(key string) (string, error) {
	res, ok := st.data[key]
	if ok {
		return res, nil
	}
	return "", errors.New("not found")
}

// ProcessCommand processes command %)
func ProcessCommand(st *Storage, cmd string) (string, error) {
	var command, key, val string
	var err error

	words := strings.Fields(cmd)

	if len(words) != 2 && len(words) != 3 {
		return emptyStr, errors.New("Syntax err")
	}

	command = words[0]
	if _, ok := supportedCommands[command]; !ok {
		log.Println("command not supported:", command)
		return emptyStr, errors.New("Not supported")
	}

	key = words[1]

	if len(words) == 3 {
		val = words[2]
	}

	switch command {
	case "SET":
		err = st.Add(key, val)
		if err != nil {
			// log.Println(err)
			return emptyStr, err
		}
		return okStr, nil
	case "DEL":
		err = st.Del(key)
		if err != nil {
			// log.Println(err)
			return emptyStr, err
		}
		return okStr, nil
	case "GET":
		res, err := st.Get(key)
		if err != nil {
			// log.Println(err)
			return emptyStr, err
		}
		return res, nil
	}

	log.Println("not implemented:", command)
	return emptyStr, errors.New("Not implemented")
}
