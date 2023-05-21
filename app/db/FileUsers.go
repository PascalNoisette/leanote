package db

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"

	"github.com/leanote/leanote/app/info"
	"gopkg.in/mgo.v2"
	"gopkg.in/yaml.v3"
)

type FileUsers struct {
	CollectionLike
	Dir      *os.File
	Name     string // "collection"
	Fallback CollectionLike
}
type Author struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Avatar      string `yaml:"avatar"`
	Pwd         string `yaml:"password_hash"`
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
	c.Fallback.One(result)
	var m = c.ReadFile()
	for k, v := range m {
		c.DeepCopy(k, v, result.(*info.User))
		break
	}
	return nil
}

func (c *FileUsers) All(result interface{}) error {
	fallbackData := info.User{}
	c.Fallback.One(&fallbackData)

	var m = c.ReadFile()

	valuePtr := reflect.ValueOf(result)
	nodelist := valuePtr.Elem()
	for k, v := range m {
		c.DeepCopy(k, v, &fallbackData)
		x := reflect.ValueOf(fallbackData)
		nodelist.Set(reflect.Append(nodelist, x))
	}
	return nil
}
func (c *FileUsers) DeepCopy(username string, in Author, out *info.User) {
	out.Username = username
	out.UsernameRaw = username
	if len(in.Pwd) == 32 {
		out.Pwd = in.Pwd
	} else {
		fmt.Fprintf(os.Stderr, "Invalid md5sum hash for password_hash found in .authors.yml\n")
	}
}
func (c *FileUsers) ReadFile() map[string]Author {
	filePath := filepath.Join(c.Dir.Name(), ".authors.yml")

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, ".authors.yml not found\n")
		return nil
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
