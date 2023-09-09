package db

import (
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/leanote/leanote/app/info"
	"github.com/leanote/leanote/app/lea"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ParsedTags struct {
	CollectionLike
	Mkdocs           *Mkdocs
	Name             string // "collection"
	CurrentFilter    bson.M
	FileNoteContents CollectionLike
}

func (c *ParsedTags) FindId(id interface{}) CollectionLike {
	//TODO
	fmt.Println("FindId " + c.Name)
	return &ParsedTags{Name: c.Name, Mkdocs: c.Mkdocs, CurrentFilter: bson.M{"_id": id}}
}

// Count returns the total number of documents in the collection.
func (c *ParsedTags) Count() (n int, err error) {
	//TODO
	fmt.Println("Count " + c.Name)
	files := []info.NoteTag{}
	c.All(&files)
	return len(files), err
}

func (c *ParsedTags) Find(query interface{}) CollectionLike {
	//TODO
	fmt.Println("Find " + c.Name)
	fmt.Println(query)
	return &ParsedTags{Name: c.Name, Mkdocs: c.Mkdocs, CurrentFilter: query.(bson.M), FileNoteContents: c.FileNoteContents}
}
func (c *ParsedTags) Skip(n int) CollectionLike {
	//TODO
	fmt.Println("Skip " + c.Name)
	return c
}
func (c *ParsedTags) DropIndex(key ...string) error {
	//TODO
	fmt.Println("DropIndex " + c.Name)
	return nil
}
func (c *ParsedTags) Sort(fields ...string) CollectionLike {
	//TODO
	fmt.Println("Sort " + c.Name)
	return c
}
func (c *ParsedTags) Select(selector interface{}) CollectionLike {
	//TODO
	fmt.Println("Select " + c.Name)
	return c
}

func (c *ParsedTags) Limit(n int) CollectionLike {
	//TODO
	fmt.Println("Limit " + c.Name)
	return c
}

func (c *ParsedTags) Distinct(key string, result interface{}) error {
	//TODO
	fmt.Println("Distinct " + c.Name)
	return nil
}

func (c *ParsedTags) One(result interface{}) (err error) {
	//TODO
	fmt.Println("One " + c.Name)
	files := []info.NoteTag{}
	fmt.Println(c.CurrentFilter)
	c.All(&files)
	for _, file := range files {
		valuePtr := reflect.ValueOf(result)
		valuePtr.Elem().Set(reflect.ValueOf(file))
		break
	}
	return nil
}

func (c *ParsedTags) All(result interface{}) error {
	fmt.Println("All" + c.Name)
	valuePtr := reflect.ValueOf(result)
	nodelist := valuePtr.Elem()
	tags := make(map[string]int)

	notebooks := c.Mkdocs.WalkDirectory()
	for _, notebook := range notebooks {
		for _, file := range notebook.Mardowns {
			content, _ := ioutil.ReadFile(file.FilePath)
			for _, tag := range c.Mkdocs.GetTags(content) {
				_, issetTag := tags[tag]
				if issetTag {
					tags[tag] += 1
				} else {
					tags[tag] = 1
				}
			}
		}
	}

	for tag, cnt := range tags {
		f := info.NoteTag{}
		f.UserId = GlobalUserId
		f.TagId = bson.ObjectId(lea.Md5(tag)[:12])
		f.Tag = tag
		f.Count = cnt

		x := reflect.ValueOf(f)
		nodelist.Set(reflect.Append(nodelist, x))
	}

	return nil
}

func (c *ParsedTags) UpdateAll(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	info = &mgo.ChangeInfo{Updated: 1}
	fmt.Println("UpdateAll" + c.Name)
	return info, err
}

func (c *ParsedTags) Insert(docs ...interface{}) error {
	//TODO
	fmt.Println("Insert " + c.Name)
	return nil
}

func (c *ParsedTags) Update(selector interface{}, update interface{}) error {
	//TODO
	fmt.Println("Update " + c.Name)
	return nil
}

func (c *ParsedTags) Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	fmt.Println("Upsert " + c.Name)
	info = &mgo.ChangeInfo{Updated: 1}
	return info, err
}

func (c *ParsedTags) Remove(selector interface{}) error {
	//TODO
	fmt.Println("Remove " + c.Name)
	return nil
}

func (c *ParsedTags) RemoveAll(selector interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	fmt.Println("RemoveAll " + c.Name)
	info = &mgo.ChangeInfo{Updated: 1}
	return info, nil
}
