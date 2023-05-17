package db

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/leanote/leanote/app/info"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/yaml.v3"
)

type FileUsers struct {
	CollectionLike
	Dir  *os.File
	Name string // "collection"
}
type Author struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Avatar      string `yaml:"avatar"`
}

func (c *FileUsers) FindId(id interface{}) CollectionLike {
	//TODO
	return c
}

// Count returns the total number of documents in the collection.
func (c *FileUsers) Count() (n int, err error) {
	var m = c.ReadFile()
	return len(m), err
}

func (c *FileUsers) Find(query interface{}) CollectionLike {
	//TODO
	return c
}
func (c *FileUsers) Skip(n int) CollectionLike {
	//TODO
	return c
}
func (c *FileUsers) DropIndex(key ...string) error {
	//TODO
	return nil
}
func (c *FileUsers) Sort(fields ...string) CollectionLike {
	//TODO
	return c
}
func (c *FileUsers) Select(selector interface{}) CollectionLike {
	//TODO
	return c
}

func (c *FileUsers) Limit(n int) CollectionLike {
	//TODO
	return c
}

func (c *FileUsers) Distinct(key string, result interface{}) error {
	//TODO
	return nil
}

func (c *FileUsers) One(result interface{}) (err error) {
	var m = c.ReadFile()

	for k, v := range m {
		buf := &bytes.Buffer{}
		gob.NewEncoder(buf).Encode(k)
		var res = info.User{UserId: bson.ObjectId(buf.Bytes()), Username: v.Name}
		result = res
	}

	return nil
}

func (c *FileUsers) All(result interface{}) error {
	var m = c.ReadFile()

	var i = len(m)
	var res = make([]info.User, i)
	for k, v := range m {
		i = i - 1
		buf := &bytes.Buffer{}
		gob.NewEncoder(buf).Encode(k)
		res[i] = info.User{UserId: bson.ObjectId(buf.Bytes()), Username: v.Name}
	}
	result = res
	return nil
}

func (c *FileUsers) ReadFile() map[string]Author {
	filePath := filepath.Join(c.Dir.Name(), ".authors.yml")

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var m map[string]Author
	//err := json.Unmarshal([]byte(data), &m)
	err = yaml.Unmarshal(data, &m)
	if err != nil {
		panic(err)
	}

	return m
}

func (c *FileUsers) UpdateAll(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	info = &mgo.ChangeInfo{Updated: 1}
	return info, err
}

func (c *FileUsers) Insert(docs ...interface{}) error {
	//TODO
	return nil
}

func (c *FileUsers) Update(selector interface{}, update interface{}) error {
	//TODO
	return nil
}

func (c *FileUsers) Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	info = &mgo.ChangeInfo{Updated: 1}
	return info, err
}

func (c *FileUsers) Remove(selector interface{}) error {
	//TODO
	return nil
}

func (c *FileUsers) RemoveAll(selector interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	info = &mgo.ChangeInfo{Updated: 1}
	return info, nil
}
