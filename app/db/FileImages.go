package db

import (
	"fmt"
	"path/filepath"
	"reflect"

	"github.com/leanote/leanote/app/info"
	"github.com/leanote/leanote/app/lea"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type FileImages struct {
	CollectionLike
	Mkdocs        *Mkdocs
	Name          string // "collection"
	CurrentFilter bson.M
	WriteHistory  *WriteHistory
}

func (c *FileImages) FindId(id interface{}) CollectionLike {
	//TODO
	fmt.Println("FindId " + c.Name)
	return &FileImages{Name: c.Name, Mkdocs: c.Mkdocs, CurrentFilter: bson.M{"_id": id}, WriteHistory: c.WriteHistory}
}

// Count returns the total number of documents in the collection.
func (c *FileImages) Count() (n int, err error) {
	//TODO
	fmt.Println("Count " + c.Name)
	files := []info.File{}
	c.All(&files)
	return len(files), err
}

func (c *FileImages) Find(query interface{}) CollectionLike {
	//TODO
	fmt.Println("Find " + c.Name)
	fmt.Println(query)
	return &FileImages{Name: c.Name, Mkdocs: c.Mkdocs, CurrentFilter: query.(bson.M), WriteHistory: c.WriteHistory}
}
func (c *FileImages) Skip(n int) CollectionLike {
	//TODO
	fmt.Println("Skip " + c.Name)
	c.CurrentFilter["Skip"] = n
	return c
}
func (c *FileImages) DropIndex(key ...string) error {
	//TODO
	fmt.Println("DropIndex " + c.Name)
	return nil
}
func (c *FileImages) Sort(fields ...string) CollectionLike {
	//TODO
	fmt.Println("Sort " + c.Name)
	return c
}
func (c *FileImages) Select(selector interface{}) CollectionLike {
	//TODO
	fmt.Println("Select " + c.Name)
	return c
}

func (c *FileImages) Limit(n int) CollectionLike {
	//TODO
	fmt.Println("Limit " + c.Name)
	c.CurrentFilter["Limit"] = n
	return c
}

func (c *FileImages) Distinct(key string, result interface{}) error {
	//TODO
	fmt.Println("Distinct " + c.Name)
	return nil
}

func (c *FileImages) One(result interface{}) (err error) {
	//TODO
	fmt.Println("One " + c.Name)
	files := []info.File{}
	fmt.Println(c.CurrentFilter)
	c.All(&files)
	for _, file := range files {
		valuePtr := reflect.ValueOf(result)
		valuePtr.Elem().Set(reflect.ValueOf(file))
		break
	}
	return nil
}

func (c *FileImages) All(result interface{}) error {
	fmt.Println("All" + c.Name)
	valuePtr := reflect.ValueOf(result)
	nodelist := valuePtr.Elem()
	i := 0

	dirs := c.Mkdocs.listDirectory()
	for _, album := range dirs {
		albumId := bson.ObjectId(lea.Md5(album)[:12])
		_, filterAlbumIdIsSet := c.CurrentFilter["AlbumId"]
		if filterAlbumIdIsSet && c.CurrentFilter["AlbumId"] != albumId {
			continue
		}
		files := c.Mkdocs.listAllFiles(album)
		for _, name := range files {

			i = i + 1

			_, pagerIsSet := c.CurrentFilter["Skip"]
			if pagerIsSet && c.CurrentFilter["Skip"].(int)+1 > i {
				continue
			}
			_, limitIsSet := c.CurrentFilter["Limit"]
			if limitIsSet && nodelist.Len() > c.CurrentFilter["Limit"].(int) {
				break
			}
			fileId := bson.ObjectId(lea.Md5(name)[:12])
			_, fileIdIsSet := c.CurrentFilter["_id"]
			if fileIdIsSet && c.WriteHistory.GetRealId(c.CurrentFilter["_id"]) != fileId.Hex() {
				continue
			}

			f := info.File{}
			f.FileId = fileId
			f.UserId = GlobalUserId
			f.AlbumId = albumId // bson.ObjectId("52d3e8ac99c37b7f0d000001" /*service.DEFAULT_ALBUM_ID*/)
			f.Name = filepath.Base(name)
			f.Title = filepath.Base(name)
			f.Path = name
			f.IsDefaultAlbum = true

			x := reflect.ValueOf(f)
			nodelist.Set(reflect.Append(nodelist, x))
		}
	}
	return nil
}

func (c *FileImages) UpdateAll(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	info = &mgo.ChangeInfo{Updated: 1}
	fmt.Println("UpdateAll" + c.Name)
	return info, err
}

func (c *FileImages) Insert(docs ...interface{}) error {
	//TODO
	fmt.Println("Insert " + c.Name)
	for _, doc := range docs {
		name := c.Mkdocs.WriteImage(doc.(info.File).Title, doc.(info.File).Path)
		c.WriteHistory.RenameObjectId(doc.(info.File).FileId, bson.ObjectId(lea.Md5(name)[:12]).Hex())
	}
	return nil
}

func (c *FileImages) Update(selector interface{}, update interface{}) error {
	//TODO
	fmt.Println("Update " + c.Name)
	return nil
}

func (c *FileImages) Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	fmt.Println("Upsert " + c.Name)
	info = &mgo.ChangeInfo{Updated: 1}
	return info, err
}

func (c *FileImages) Remove(selector interface{}) error {
	//TODO
	fmt.Println("Remove " + c.Name)
	return nil
}

func (c *FileImages) RemoveAll(selector interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	fmt.Println("RemoveAll " + c.Name)
	info = &mgo.ChangeInfo{Updated: 1}
	return info, nil
}
