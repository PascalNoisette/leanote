package db

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BsonReader struct {
	CollectionLike
	Dir  *os.File
	Name string // "collection"
}

func (c *BsonReader) FindId(id interface{}) CollectionLike {
	//TODO
	fmt.Println("FindId " + c.Name)
	fmt.Println(id)
	return c
}

// Count returns the total number of documents in the collection.
func (c *BsonReader) Count() (n int, err error) {
	//TODO
	fmt.Println("Count " + c.Name)
	return 1, err
}

func (c *BsonReader) Find(query interface{}) CollectionLike {
	//TODO
	fmt.Println("Find " + c.Name)
	fmt.Println(query)
	return c
}
func (c *BsonReader) Skip(n int) CollectionLike {
	//TODO
	fmt.Println("Skip " + c.Name)
	fmt.Println(n)
	return c
}
func (c *BsonReader) DropIndex(key ...string) error {
	//TODO
	fmt.Println("DropIndex " + c.Name)
	fmt.Println(key)
	return nil
}
func (c *BsonReader) Sort(fields ...string) CollectionLike {
	//TODO
	fmt.Println("Sort " + c.Name)
	fmt.Println(fields)
	return c
}
func (c *BsonReader) Select(selector interface{}) CollectionLike {
	//TODO
	fmt.Println("Select " + c.Name)
	fmt.Println(selector)
	return c
}

func (c *BsonReader) Limit(n int) CollectionLike {
	//TODO
	fmt.Println("Limit " + c.Name)
	fmt.Println(n)
	return c
}

func (c *BsonReader) Distinct(key string, result interface{}) error {
	//TODO
	fmt.Println("Distinct " + c.Name)
	fmt.Println(key)
	return nil
}

func (c *BsonReader) One(result interface{}) error {
	var nodelist map[string]interface{}
	var err = c.All(nodelist)
	if err != nil {
		return err
	}
	for _, v := range nodelist {
		var data, err = bson.Marshal(v)
		if err != nil {
			return err
		}
		err = bson.Unmarshal(data, result)
		fmt.Println("Found one " + c.Name)
		fmt.Println(result)
		if err != nil {
			return err
		}
		break
	}
	fmt.Println("Found none " + c.Name)
	return nil
}

func (c *BsonReader) All(result interface{}) error {
	filePath := filepath.Join("/leanote/mongodb_backup/leanote_install_data", c.Name+".bson")
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	err = bson.Unmarshal(data, &result)
	fmt.Println(result)
	return err
}

func (c *BsonReader) UpdateAll(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	fmt.Println("UpdateAll " + c.Name)
	info = &mgo.ChangeInfo{Updated: 1}
	return info, err
}

func (c *BsonReader) Insert(docs ...interface{}) error {
	//TODO
	fmt.Println("Insert " + c.Name)
	return nil
}

func (c *BsonReader) Update(selector interface{}, update interface{}) error {
	//TODO
	fmt.Println("Update " + c.Name)
	return nil
}

func (c *BsonReader) Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	fmt.Println("Upsert " + c.Name)
	info = &mgo.ChangeInfo{Updated: 1}
	return info, err
}

func (c *BsonReader) Remove(selector interface{}) error {
	//TODO
	fmt.Println("Remove " + c.Name)
	return nil
}

func (c *BsonReader) RemoveAll(selector interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	fmt.Println("RemoveAll " + c.Name)
	fmt.Println(selector)
	info = &mgo.ChangeInfo{Updated: 1}
	return info, nil
}
