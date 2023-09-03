package db

import (
	"fmt"
	"reflect"

	"github.com/leanote/leanote/app/info"
	"github.com/leanote/leanote/app/lea"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type FolderImages struct {
	CollectionLike
	Mkdocs        *Mkdocs
	Name          string // "collection"
	CurrentFilter bson.M
}

func (c *FolderImages) FindId(id interface{}) CollectionLike {
	//TODO
	fmt.Println("FindId " + c.Name)
	return &FolderImages{Name: c.Name, Mkdocs: c.Mkdocs, CurrentFilter: bson.M{"_id": id}}
}

// Count returns the total number of documents in the collection.
func (c *FolderImages) Count() (n int, err error) {
	//TODO
	fmt.Println("Count " + c.Name)
	files := []info.Album{}
	c.All(&files)
	return len(files), err
}

func (c *FolderImages) Find(query interface{}) CollectionLike {
	//TODO
	fmt.Println("Find " + c.Name)
	fmt.Println(query)
	return &FolderImages{Name: c.Name, Mkdocs: c.Mkdocs, CurrentFilter: query.(bson.M)}
}
func (c *FolderImages) Skip(n int) CollectionLike {
	//TODO
	fmt.Println("Skip " + c.Name)
	return c
}
func (c *FolderImages) DropIndex(key ...string) error {
	//TODO
	fmt.Println("DropIndex " + c.Name)
	return nil
}
func (c *FolderImages) Sort(fields ...string) CollectionLike {
	//TODO
	fmt.Println("Sort " + c.Name)
	return c
}
func (c *FolderImages) Select(selector interface{}) CollectionLike {
	//TODO
	fmt.Println("Select " + c.Name)
	return c
}

func (c *FolderImages) Limit(n int) CollectionLike {
	//TODO
	fmt.Println("Limit " + c.Name)
	return c
}

func (c *FolderImages) Distinct(key string, result interface{}) error {
	//TODO
	fmt.Println("Distinct " + c.Name)
	return nil
}

func (c *FolderImages) One(result interface{}) (err error) {
	//TODO
	fmt.Println("One " + c.Name)
	files := []info.Album{}
	fmt.Println(c.CurrentFilter)
	c.All(&files)
	for _, file := range files {
		valuePtr := reflect.ValueOf(result)
		valuePtr.Elem().Set(reflect.ValueOf(file))
		break
	}
	return nil
}

func (c *FolderImages) All(result interface{}) error {
	fmt.Println("All" + c.Name)
	valuePtr := reflect.ValueOf(result)
	nodelist := valuePtr.Elem()

	files := c.Mkdocs.listDirectory()
	for _, name := range files {

		f := info.Album{}
		f.AlbumId = bson.ObjectId(lea.Md5(name)[:12]) /*service.DEFAULT_ALBUM_ID 52d3e8ac99c37b7f0d000001*/
		f.UserId = GlobalUserId
		f.Name = name

		x := reflect.ValueOf(f)
		nodelist.Set(reflect.Append(nodelist, x))
	}
	return nil
}

func (c *FolderImages) UpdateAll(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	info = &mgo.ChangeInfo{Updated: 1}
	fmt.Println("UpdateAll" + c.Name)
	return info, err
}

func (c *FolderImages) Insert(docs ...interface{}) error {
	//TODO
	fmt.Println("Insert " + c.Name)
	return nil
}

func (c *FolderImages) Update(selector interface{}, update interface{}) error {
	//TODO
	fmt.Println("Update " + c.Name)
	return nil
}

func (c *FolderImages) Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	fmt.Println("Upsert " + c.Name)
	info = &mgo.ChangeInfo{Updated: 1}
	return info, err
}

func (c *FolderImages) Remove(selector interface{}) error {
	//TODO
	fmt.Println("Remove " + c.Name)
	return nil
}

func (c *FolderImages) RemoveAll(selector interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	fmt.Println("RemoveAll " + c.Name)
	info = &mgo.ChangeInfo{Updated: 1}
	return info, nil
}
