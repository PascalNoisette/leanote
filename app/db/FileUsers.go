package db

import (
	"fmt"
	"os"
	"reflect"

	"github.com/leanote/leanote/app/info"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var GlobalUserId = bson.NewObjectId()

type FileUsers struct {
	CollectionLike
	Name   string // "collection"
	Mkdocs *Mkdocs
}

func (c *FileUsers) FindId(id interface{}) CollectionLike {
	//TODO
	return c
}

// Count returns the total number of documents in the collection.
func (c *FileUsers) Count() (n int, err error) {
	var m = c.Mkdocs.ReadAuthorFile()
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
	users := []info.User{}
	c.All(&users)
	for _, u := range users {
		valuePtr := reflect.ValueOf(result)
		valuePtr.Elem().Set(reflect.ValueOf(u))
		break
	}
	return nil
}

func (c *FileUsers) All(result interface{}) error {
	var m = c.Mkdocs.ReadAuthorFile()

	valuePtr := reflect.ValueOf(result)
	nodelist := valuePtr.Elem()
	for k, v := range m {
		user := info.User{}
		c.DeepCopy(k, v, &user)
		x := reflect.ValueOf(user)
		nodelist.Set(reflect.Append(nodelist, x))
	}
	return nil
}
func (c *FileUsers) DeepCopy(username string, in Author, out *info.User) {
	out.Username = username
	out.UsernameRaw = username
	out.UserId = GlobalUserId
	out.Email = username
	if len(in.Pwd) == 32 {
		out.Pwd = in.Pwd
	} else {
		fmt.Fprintf(os.Stderr, "Invalid md5sum hash for password_hash found in .authors.yml\n")
	}
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
