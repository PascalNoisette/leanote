package db

import (
	"fmt"
	"reflect"

	"github.com/leanote/leanote/app/info"
	"github.com/leanote/leanote/app/lea"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type FolderNotebooks struct {
	CollectionLike
	Mkdocs        *Mkdocs
	Name          string // "collection"
	CurrentFilter bson.M
	WriteHistory  *WriteHistory
}

func (c *FolderNotebooks) FindId(id interface{}) CollectionLike {
	//TODO
	fmt.Println("FindId " + c.Name)
	return &FolderNotebooks{Name: "notebooks", Mkdocs: c.Mkdocs, CurrentFilter: bson.M{"_id": id}, WriteHistory: c.WriteHistory}
}

// Count returns the total number of documents in the collection.
func (c *FolderNotebooks) Count() (n int, err error) {
	//TODO
	fmt.Println("Count " + c.Name)
	notebooks := []info.Notebook{}
	c.All(&notebooks)
	return len(notebooks), err
}

func (c *FolderNotebooks) Find(query interface{}) CollectionLike {
	//TODO
	fmt.Println("Find " + c.Name)
	fmt.Println(query)
	return &FolderNotebooks{Name: "notebooks", Mkdocs: c.Mkdocs, CurrentFilter: query.(bson.M), WriteHistory: c.WriteHistory}
}
func (c *FolderNotebooks) Skip(n int) CollectionLike {
	//TODO
	fmt.Println("Skip " + c.Name)
	return c
}
func (c *FolderNotebooks) DropIndex(key ...string) error {
	//TODO
	fmt.Println("DropIndex " + c.Name)
	return nil
}
func (c *FolderNotebooks) Sort(fields ...string) CollectionLike {
	//TODO
	fmt.Println("Sort " + c.Name)
	return c
}
func (c *FolderNotebooks) Select(selector interface{}) CollectionLike {
	//TODO
	fmt.Println("Select " + c.Name)
	return c
}

func (c *FolderNotebooks) Limit(n int) CollectionLike {
	//TODO
	fmt.Println("Limit " + c.Name)
	return c
}

func (c *FolderNotebooks) Distinct(key string, result interface{}) error {
	//TODO
	fmt.Println("Distinct " + c.Name)
	return nil
}

func (c *FolderNotebooks) One(result interface{}) (err error) {
	//TODO
	fmt.Println("One " + c.Name)
	notebooks := []info.Notebook{}
	fmt.Println(c.CurrentFilter)
	c.All(&notebooks)
	for _, notebook := range notebooks {
		valuePtr := reflect.ValueOf(result)
		valuePtr.Elem().Set(reflect.ValueOf(notebook))
		break
	}
	return nil
}

func (c *FolderNotebooks) All(result interface{}) error {
	fmt.Println("All" + c.Name)
	valuePtr := reflect.ValueOf(result)
	nodelist := valuePtr.Elem()

	_, filterBlog := c.CurrentFilter["IsBlog"]
	if filterBlog {
		return nil
	}

	files := c.Mkdocs.WalkDirectory()
	for i, file := range files {

		_, filterIsSet := c.CurrentFilter["UrlTitle"]
		if filterIsSet && c.CurrentFilter["UrlTitle"] != file.Name {
			continue
		}
		_, filterIdIsSet := c.CurrentFilter["_id"]

		if filterIdIsSet {
			_, filterByObjectId := c.CurrentFilter["_id"].(bson.ObjectId)
			if filterByObjectId && c.WriteHistory.GetRealId(c.CurrentFilter["_id"]) != bson.ObjectId(lea.Md5(file.Name)[:12]).Hex() {
				continue
			}
		}

		_, filterByParent := c.CurrentFilter["ParentNotebookId"]
		if filterByParent {
			// no parent implemented
			continue
		}

		notebook := info.Notebook{}
		notebook.NotebookId = bson.ObjectId(lea.Md5(file.Name)[:12])
		notebook.Title = file.Name
		notebook.UrlTitle = file.Name
		notebook.UserId = GlobalUserId
		notebook.Usn = i
		notebook.IsDeleted = false
		notebook.IsTrash = false
		notebook.NumberNotes = len(file.Mardowns)
		x := reflect.ValueOf(notebook)
		nodelist.Set(reflect.Append(nodelist, x))
	}
	return nil
}

func (c *FolderNotebooks) UpdateAll(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	info = &mgo.ChangeInfo{Updated: 1}
	fmt.Println("UpdateAll" + c.Name)
	fmt.Println(selector)
	fmt.Println(update)
	filterId, filterUidIsSet := selector.(bson.M)["_id"]
	_, isSetOperation := update.(bson.M)["$set"]
	if filterUidIsSet && isSetOperation {
		notebooks := c.Mkdocs.WalkDirectory()
		for _, notebook := range notebooks {
			notebookId := bson.ObjectId(lea.Md5(notebook.Name)[:12])
			if c.WriteHistory.GetRealId(filterId) != notebookId.Hex() {
				continue
			}
			_, delete := update.(bson.M)["$set"].(bson.M)["IsDeleted"]
			if delete {
				lea.DeleteFile(notebook.FilePath)
			}
			newTitle, rename := update.(bson.M)["$set"].(bson.M)["Title"]
			if rename {
				err := c.Mkdocs.RenameInPath(notebook.FilePath, newTitle.(string))
				if err == nil {
					lea.DeleteFile(notebook.FilePath)
					c.WriteHistory.RenameObjectId(notebookId, bson.ObjectId(lea.Md5(newTitle.(string))[:12]).Hex())
				}
			}
		}
	}
	return info, err
}

func (c *FolderNotebooks) Insert(docs ...interface{}) error {
	//TODO
	fmt.Println("Insert " + c.Name)
	for _, doc := range docs {
		c.Mkdocs.createDirectory(doc.(info.Notebook).Title)
		c.WriteHistory.RenameObjectId(doc.(info.Notebook).NotebookId, bson.ObjectId(lea.Md5(doc.(info.Notebook).Title)[:12]).Hex())
	}
	return nil
}

func (c *FolderNotebooks) Update(selector interface{}, update interface{}) error {
	//TODO
	_, err := c.UpdateAll(selector, update)
	return err
}

func (c *FolderNotebooks) Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	fmt.Println("Upsert " + c.Name)
	info = &mgo.ChangeInfo{Updated: 1}
	return info, err
}

func (c *FolderNotebooks) Remove(selector interface{}) error {
	//TODO
	fmt.Println("Remove " + c.Name)
	return nil
}

func (c *FolderNotebooks) RemoveAll(selector interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	fmt.Println("RemoveAll " + c.Name)
	info = &mgo.ChangeInfo{Updated: 1}
	return info, nil
}
