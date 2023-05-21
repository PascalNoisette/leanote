package db

import (
	"fmt"
	"os"
	"reflect"

	"github.com/leanote/leanote/app/info"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type FolderNotebooks struct {
	CollectionLike
	Dir      *os.File
	Name     string // "collection"
	Fallback CollectionLike
}

func (c *FolderNotebooks) FindId(id interface{}) CollectionLike {
	//TODO
	return c
}

// Count returns the total number of documents in the collection.
func (c *FolderNotebooks) Count() (n int, err error) {
	//TODO
	fmt.Println("Count " + c.Name)
	return 1, err
}

func (c *FolderNotebooks) Find(query interface{}) CollectionLike {
	//TODO
	return c
}
func (c *FolderNotebooks) Skip(n int) CollectionLike {
	//TODO
	return c
}
func (c *FolderNotebooks) DropIndex(key ...string) error {
	//TODO
	return nil
}
func (c *FolderNotebooks) Sort(fields ...string) CollectionLike {
	//TODO
	return c
}
func (c *FolderNotebooks) Select(selector interface{}) CollectionLike {
	//TODO
	return c
}

func (c *FolderNotebooks) Limit(n int) CollectionLike {
	//TODO
	return c
}

func (c *FolderNotebooks) Distinct(key string, result interface{}) error {
	//TODO
	return nil
}

func (c *FolderNotebooks) One(result interface{}) (err error) {
	//TODO
	return nil
}

func (c *FolderNotebooks) All(result interface{}) error {

	valuePtr := reflect.ValueOf(result)
	nodelist := valuePtr.Elem()
	files, err := c.Dir.Readdir(0)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.IsDir() {
			notebook := info.Notebook{}
			notebook.NotebookId = bson.NewObjectId()
			notebook.Title = file.Name()
			notebook.UrlTitle = file.Name()
			x := reflect.ValueOf(notebook)
			nodelist.Set(reflect.Append(nodelist, x))
		}
	}
	return nil
}

func (c *FolderNotebooks) UpdateAll(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	info = &mgo.ChangeInfo{Updated: 1}
	return info, err
}

func (c *FolderNotebooks) Insert(docs ...interface{}) error {
	//TODO
	return nil
}

func (c *FolderNotebooks) Update(selector interface{}, update interface{}) error {
	//TODO
	return nil
}

func (c *FolderNotebooks) Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	info = &mgo.ChangeInfo{Updated: 1}
	return info, err
}

func (c *FolderNotebooks) Remove(selector interface{}) error {
	//TODO
	return nil
}

func (c *FolderNotebooks) RemoveAll(selector interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	info = &mgo.ChangeInfo{Updated: 1}
	return info, nil
}
