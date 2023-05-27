package db

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/leanote/leanote/app/info"
	"github.com/leanote/leanote/app/lea"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type FileNoteContents struct {
	CollectionLike
	Mkdocs          *Mkdocs
	Name            string // "collection"
	CurrentFilter   bson.M
	FolderNotebooks CollectionLike
	FileNotes       CollectionLike
}

func (c *FileNoteContents) FindId(id interface{}) CollectionLike {
	//TODO
	fmt.Println("FindId" + c.Name)
	fmt.Println(id)
	return &FileNoteContents{Name: "note_contents", Mkdocs: c.Mkdocs, CurrentFilter: bson.M{"_id": id}}
}

// Count returns the total number of documents in the collection.
func (c *FileNoteContents) Count() (n int, err error) {
	//TODO
	fmt.Println("Count " + c.Name)
	return 1, err
}

func (c *FileNoteContents) Find(query interface{}) CollectionLike {
	//TODO
	fmt.Println("Find" + c.Name)
	return &FileNoteContents{Name: "note_contents", Mkdocs: c.Mkdocs, CurrentFilter: query.(bson.M)}
}
func (c *FileNoteContents) Skip(n int) CollectionLike {
	//TODO
	fmt.Println("Skip" + c.Name)
	fmt.Println(n)
	return c
}
func (c *FileNoteContents) DropIndex(key ...string) error {
	//TODO
	fmt.Println("DropIndex" + c.Name)
	fmt.Println(key)
	return nil
}
func (c *FileNoteContents) Sort(fields ...string) CollectionLike {
	//TODO
	fmt.Println("Sort" + c.Name)
	fmt.Println(fields)
	return c
}
func (c *FileNoteContents) Select(selector interface{}) CollectionLike {
	//TODO
	fmt.Println("Select" + c.Name)
	fmt.Println(selector)
	return c
}

func (c *FileNoteContents) Limit(n int) CollectionLike {
	//TODO
	fmt.Println("Limit" + c.Name)
	fmt.Println(n)
	return c
}

func (c *FileNoteContents) Distinct(key string, result interface{}) error {
	//TODO
	fmt.Println("Distinct" + c.Name)
	fmt.Println(key)
	fmt.Println(result)
	return nil
}

func (c *FileNoteContents) One(result interface{}) (err error) {
	//TODO
	fmt.Println("One " + c.Name)
	fmt.Println("current filter  " + c.CurrentFilter["_id"].(bson.ObjectId).Hex())
	notebooks := c.Mkdocs.WalkDirectory()
	for _, notebook := range notebooks {
		for _, file := range notebook.Mardowns {
			if c.CurrentFilter["_id"] == bson.ObjectId(lea.Md5(notebook.Name + file.Name)[:12]) {
				data, err := ioutil.ReadFile(file.FilePath)
				if err != nil {
					fmt.Fprintf(os.Stderr, "file "+file.FilePath+" not found\n")
					return nil
				}
				result.(*info.NoteContent).Content = string(data)
				result.(*info.NoteContent).NoteId = c.CurrentFilter["_id"].(bson.ObjectId)
				result.(*info.NoteContent).UserId = GlobalUserId
			}
		}
	}

	return nil
}

func (c *FileNoteContents) All(result interface{}) error {
	fmt.Println("All" + c.Name)
	return nil
}

func (c *FileNoteContents) UpdateAll(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	fmt.Println("UpdateAll" + c.Name)
	fmt.Println(selector)
	fmt.Println(update)
	info = &mgo.ChangeInfo{Updated: 1}
	return info, err
}

func (c *FileNoteContents) Insert(docs ...interface{}) error {
	//TODO
	fmt.Println("Insert" + c.Name)
	notebook := info.Notebook{}
	note := info.Note{}
	for _, doc := range docs {
		c.FileNotes.FindId(doc.(info.NoteContent).NoteId).One(&note)
		c.FolderNotebooks.FindId(note.NotebookId).One(&notebook)
		c.Mkdocs.WriteFile(notebook.Title, note.Title, doc.(info.NoteContent).Content)
	}
	return nil
}

func (c *FileNoteContents) Update(selector interface{}, update interface{}) error {
	//TODO
	fmt.Println("Update" + c.Name)
	fmt.Println(selector)
	fmt.Println(update)
	return nil
}

func (c *FileNoteContents) Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	fmt.Println("Upsert" + c.Name)
	fmt.Println(selector)
	fmt.Println(update)
	info = &mgo.ChangeInfo{Updated: 1}
	return info, err
}

func (c *FileNoteContents) Remove(selector interface{}) error {
	//TODO
	fmt.Println("Remove" + c.Name)
	fmt.Println(selector)
	return nil
}

func (c *FileNoteContents) RemoveAll(selector interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	fmt.Println("RemoveAll" + c.Name)
	fmt.Println(selector)
	info = &mgo.ChangeInfo{Updated: 1}
	return info, nil
}
