package db

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"regexp"
	"strings"

	"github.com/leanote/leanote/app/info"
	"github.com/leanote/leanote/app/lea"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type FileNotes struct {
	CollectionLike
	Mkdocs          *Mkdocs
	Name            string // "collection"
	CurrentFilter   bson.M
	FolderNotebooks CollectionLike
	WriteHistory    *WriteHistory
}

func (c *FileNotes) FindId(id interface{}) CollectionLike {
	//TODO
	fmt.Println("FindId" + c.Name)
	fmt.Println(id)
	return &FileNotes{Name: "notes", Mkdocs: c.Mkdocs, CurrentFilter: bson.M{"_id": id}, FolderNotebooks: c.FolderNotebooks, WriteHistory: c.WriteHistory}
}

// Count returns the total number of documents in the collection.
func (c *FileNotes) Count() (n int, err error) {
	//TODO
	fmt.Println("Count " + c.Name)
	notes := []info.Note{}
	c.All(&notes)
	return len(notes), err
}

func (c *FileNotes) Find(query interface{}) CollectionLike {
	//TODO
	fmt.Println("Find" + c.Name)
	fmt.Println(query)
	return &FileNotes{Name: "notes", Mkdocs: c.Mkdocs, CurrentFilter: query.(bson.M), FolderNotebooks: c.FolderNotebooks, WriteHistory: c.WriteHistory}
}
func (c *FileNotes) Skip(n int) CollectionLike {
	//TODO
	fmt.Println("Skip" + c.Name)
	fmt.Println(n)
	return c
}
func (c *FileNotes) DropIndex(key ...string) error {
	//TODO
	fmt.Println("DropIndex" + c.Name)
	fmt.Println(key)
	return nil
}
func (c *FileNotes) Sort(fields ...string) CollectionLike {
	//TODO
	fmt.Println("Sort" + c.Name)
	fmt.Println(fields)
	return c
}
func (c *FileNotes) Select(selector interface{}) CollectionLike {
	//TODO
	fmt.Println("Select" + c.Name)
	fmt.Println(selector)
	return c
}

func (c *FileNotes) Limit(n int) CollectionLike {
	//TODO
	fmt.Println("Limit" + c.Name)
	fmt.Println(n)
	return c
}

func (c *FileNotes) Distinct(key string, result interface{}) error {
	//TODO
	fmt.Println("Distinct" + c.Name)
	fmt.Println(key)
	fmt.Println(result)
	return nil
}

func (c *FileNotes) One(result interface{}) (err error) {
	//TODO
	fmt.Println("One " + c.Name)
	notes := []info.Note{}
	c.All(&notes)
	for _, note := range notes {
		valuePtr := reflect.ValueOf(result)
		valuePtr.Elem().Set(reflect.ValueOf(note))
		break
	}
	fmt.Println(result)

	return nil
}

func (c *FileNotes) All(result interface{}) error {
	fmt.Println("All" + c.Name)
	valuePtr := reflect.ValueOf(result)
	nodelist := valuePtr.Elem()
	notebooks := c.Mkdocs.WalkDirectory()
	fmt.Println("filter " + c.Name)
	fmt.Println(c.CurrentFilter)
	_, filterBlog := c.CurrentFilter["IsBlog"]
	if filterBlog {
		return nil
	}
	for _, notebook := range notebooks {
		notebookId := bson.ObjectId(lea.Md5(notebook.Name)[:12])
		_, filterNidIsSet := c.CurrentFilter["NotebookId"]
		if filterNidIsSet && c.CurrentFilter["NotebookId"].(bson.ObjectId).Hex() != notebookId.Hex() {
			continue
		}
		for i, file := range notebook.Mardowns {
			noteId := bson.ObjectId(lea.Md5(notebook.Name + file.Name)[:12])
			_, filterUidIsSet := c.CurrentFilter["_id"]

			if filterUidIsSet && c.WriteHistory.GetRealId(c.CurrentFilter["_id"]) != noteId.Hex() {
				continue
			}
			_, filterUrlIsSet := c.CurrentFilter["UrlTitle"]
			if filterUrlIsSet && c.CurrentFilter["UrlTitle"] != file.Name {
				continue
			}

			content, _ := ioutil.ReadFile(file.FilePath)

			tags := c.Mkdocs.GetTags(content)
			_, filterTagsIsSet := c.CurrentFilter["Tags"]
			if filterTagsIsSet && !c.Mkdocs.Contains(tags, c.CurrentFilter["Tags"].(bson.M)["$all"].([]string)[0]) {
				continue
			}

			_, filterOrSet := c.CurrentFilter["$or"]
			if filterOrSet {
				// text search NoteService.SearchNote
				_, searchText := c.CurrentFilter["$or"].([]bson.M)[0]["Title"]
				if searchText {
					re := regexp.MustCompile(c.CurrentFilter["$or"].([]bson.M)[0]["Title"].(bson.M)["$regex"].(bson.RegEx).Pattern)
					if re.FindSubmatchIndex(content) == nil && re.FindSubmatchIndex([]byte(file.Name)) == nil {
						continue
					}

				}
			}

			ImgSrc := ""
			re := regexp.MustCompile(`!\[.*?\]\((.*?)(?:\s".*?")?\)`)
			match := re.FindStringSubmatch(string(content))
			if len(match) > 0 {
				ImgSrc = match[1]
			}
			var Desc []string
			re = regexp.MustCompile(`(?m)^#+\s+(.+)$`)
			matches := re.FindAllStringSubmatch(string(content), -1)
			for _, match := range matches {
				Desc = append(Desc, match[1])
			}

			note := info.Note{}
			note.NoteId = noteId
			note.NotebookId = notebookId
			note.Title = file.Name
			note.UrlTitle = file.Name
			note.IsMarkdown = true
			note.UserId = GlobalUserId
			note.Usn = i
			note.IsTrash = false
			note.CreatedTime = file.ModTime
			note.UpdatedTime = file.ModTime
			note.Tags = tags
			note.ImgSrc = ImgSrc
			note.Desc = strings.Join(Desc, ", ")
			x := reflect.ValueOf(note)
			nodelist.Set(reflect.Append(nodelist, x))
		}
	}
	return nil
}

func (c *FileNotes) UpdateAll(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	fmt.Println("UpdateAll" + c.Name)
	fmt.Println(selector)
	fmt.Println(update)
	filterId, filterUidIsSet := selector.(bson.M)["_id"]
	_, isSetOperation := update.(bson.M)["$set"]
	if filterUidIsSet && isSetOperation {

		notebooks := c.Mkdocs.WalkDirectory()
		for _, notebook := range notebooks {
			for _, file := range notebook.Mardowns {
				noteId := bson.ObjectId(lea.Md5(notebook.Name + file.Name)[:12])

				if c.WriteHistory.GetRealId(filterId) != noteId.Hex() {
					continue
				}
				_, trash := update.(bson.M)["$set"].(bson.M)["IsTrash"]
				if trash {
					lea.DeleteFile(file.FilePath)
				}
				newTitle, rename := update.(bson.M)["$set"].(bson.M)["Title"]
				if rename {
					err := c.Mkdocs.RenameInPath(file.FilePath, newTitle.(string))
					if err == nil {
						lea.DeleteFile(file.FilePath)
						c.WriteHistory.RenameObjectId(noteId, bson.ObjectId(lea.Md5(notebook.Name + newTitle.(string))[:12]).Hex())
					}
				}
			}
		}
	}
	info = &mgo.ChangeInfo{Updated: 1}
	return info, err
}

func (c *FileNotes) Insert(docs ...interface{}) error {
	//TODO
	fmt.Println("Insert" + c.Name)
	fmt.Println(docs)
	notebook := info.Notebook{}
	for _, doc := range docs { //TODO
		c.FolderNotebooks.FindId(doc.(info.Note).NotebookId).One(&notebook)
		fmt.Println("wihthin notebook")
		fmt.Println(notebook)
		c.Mkdocs.WriteFile(notebook.Title, doc.(info.Note).Title, "")
		c.WriteHistory.RenameObjectId(doc.(info.Note).NoteId, bson.ObjectId(lea.Md5(notebook.Title + doc.(info.Note).Title)[:12]).Hex())
	}
	return nil
}

func (c *FileNotes) Update(selector interface{}, update interface{}) error {
	_, err := c.UpdateAll(selector, update)
	return err
}

func (c *FileNotes) Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	fmt.Println("Upsert" + c.Name)
	fmt.Println(selector)
	fmt.Println(update)
	info = &mgo.ChangeInfo{Updated: 1}
	return info, err
}

func (c *FileNotes) Remove(selector interface{}) error {
	//TODO
	fmt.Println("Remove" + c.Name)
	fmt.Println(selector)
	return nil
}

func (c *FileNotes) RemoveAll(selector interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	fmt.Println("RemoveAll" + c.Name)
	fmt.Println(selector)
	info = &mgo.ChangeInfo{Updated: 1}
	return info, nil
}
