package db

import (
	"fmt"
	"path/filepath"
	"reflect"
	"regexp"

	"github.com/leanote/leanote/app/info"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ParsedAttachs struct {
	CollectionLike
	Mkdocs           *Mkdocs
	Name             string // "collection"
	CurrentFilter    bson.M
	FileNoteContents CollectionLike
}

func (c *ParsedAttachs) FindId(id interface{}) CollectionLike {
	//TODO
	fmt.Println("FindId " + c.Name)
	return &ParsedAttachs{Name: c.Name, Mkdocs: c.Mkdocs, CurrentFilter: bson.M{"_id": id}}
}

// Count returns the total number of documents in the collection.
func (c *ParsedAttachs) Count() (n int, err error) {
	//TODO
	fmt.Println("Count " + c.Name)
	files := []info.Attach{}
	c.All(&files)
	return len(files), err
}

func (c *ParsedAttachs) Find(query interface{}) CollectionLike {
	//TODO
	fmt.Println("Find " + c.Name)
	fmt.Println(query)
	return &ParsedAttachs{Name: c.Name, Mkdocs: c.Mkdocs, CurrentFilter: query.(bson.M), FileNoteContents: c.FileNoteContents}
}
func (c *ParsedAttachs) Skip(n int) CollectionLike {
	//TODO
	fmt.Println("Skip " + c.Name)
	return c
}
func (c *ParsedAttachs) DropIndex(key ...string) error {
	//TODO
	fmt.Println("DropIndex " + c.Name)
	return nil
}
func (c *ParsedAttachs) Sort(fields ...string) CollectionLike {
	//TODO
	fmt.Println("Sort " + c.Name)
	return c
}
func (c *ParsedAttachs) Select(selector interface{}) CollectionLike {
	//TODO
	fmt.Println("Select " + c.Name)
	return c
}

func (c *ParsedAttachs) Limit(n int) CollectionLike {
	//TODO
	fmt.Println("Limit " + c.Name)
	return c
}

func (c *ParsedAttachs) Distinct(key string, result interface{}) error {
	//TODO
	fmt.Println("Distinct " + c.Name)
	return nil
}

func (c *ParsedAttachs) One(result interface{}) (err error) {
	//TODO
	fmt.Println("One " + c.Name)
	files := []info.Attach{}
	fmt.Println(c.CurrentFilter)
	c.All(&files)
	for _, file := range files {
		valuePtr := reflect.ValueOf(result)
		valuePtr.Elem().Set(reflect.ValueOf(file))
		break
	}
	return nil
}

func (c *ParsedAttachs) All(result interface{}) error {
	fmt.Println("All" + c.Name)
	valuePtr := reflect.ValueOf(result)
	nodelist := valuePtr.Elem()

	_, filterNoteIdIsSet := c.CurrentFilter["NoteId"]
	if filterNoteIdIsSet {
		noteContent := info.NoteContent{}
		c.FileNoteContents.FindId(c.CurrentFilter["NoteId"]).One(&noteContent)

		re := regexp.MustCompile(`!\[.*?\]\((.*?)(?:\s".*?")?\)`)
		matches := re.FindAllStringSubmatch(noteContent.Content, -1)
		for _, match := range matches {
			fmt.Println(match[1])

			f := info.Attach{}
			f.AttachId = bson.ObjectId([]byte(match[1]))
			f.NoteId = c.CurrentFilter["NoteId"].(bson.ObjectId)
			f.UploadUserId = GlobalUserId
			f.Name = filepath.Base(match[1])
			f.Title = filepath.Base(match[1])
			ext := filepath.Ext(match[1])
			if len(ext) > 0 {
				f.Type = filepath.Ext(match[1])[1:]
			}
			f.Path = match[1]

			x := reflect.ValueOf(f)
			nodelist.Set(reflect.Append(nodelist, x))
		}
	}

	return nil
}

func (c *ParsedAttachs) UpdateAll(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	info = &mgo.ChangeInfo{Updated: 1}
	fmt.Println("UpdateAll" + c.Name)
	return info, err
}

func (c *ParsedAttachs) Insert(docs ...interface{}) error {
	//TODO
	fmt.Println("Insert " + c.Name)
	for _, doc := range docs {
		c.Mkdocs.WriteImage(doc.(info.Attach).Title, doc.(info.Attach).Path)
	}
	return nil
}

func (c *ParsedAttachs) Update(selector interface{}, update interface{}) error {
	//TODO
	fmt.Println("Update " + c.Name)
	return nil
}

func (c *ParsedAttachs) Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	fmt.Println("Upsert " + c.Name)
	info = &mgo.ChangeInfo{Updated: 1}
	return info, err
}

func (c *ParsedAttachs) Remove(selector interface{}) error {
	//TODO
	fmt.Println("Remove " + c.Name)
	return nil
}

func (c *ParsedAttachs) RemoveAll(selector interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	fmt.Println("RemoveAll " + c.Name)
	info = &mgo.ChangeInfo{Updated: 1}
	return info, nil
}
