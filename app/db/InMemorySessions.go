package db

import (
	"fmt"
	"reflect"

	"github.com/leanote/leanote/app/info"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type InMemorySessions struct {
	CollectionLike
	database      map[string]info.Session
	CurrentFilter bson.M
	Name          string // "collection"
}

func (c *InMemorySessions) FindId(id interface{}) CollectionLike {
	//TODO
	fmt.Println("FindId " + c.Name)
	fmt.Println(id)
	return &InMemorySessions{Name: c.Name, CurrentFilter: bson.M{"_id": id}, database: c.database}
}

// Count returns the total number of documents in the collection.
func (c *InMemorySessions) Count() (n int, err error) {
	//TODO
	fmt.Println("Count " + c.Name)
	sessions := []info.Session{}
	c.All(&sessions)
	return len(sessions), err
}

func (c *InMemorySessions) Find(query interface{}) CollectionLike {
	//TODO
	fmt.Println("Find " + c.Name)
	fmt.Println(query)
	return &InMemorySessions{Name: c.Name, CurrentFilter: query.(bson.M), database: c.database}
}
func (c *InMemorySessions) Skip(n int) CollectionLike {
	//TODO
	fmt.Println("Skip " + c.Name)
	fmt.Println(n)
	return c
}
func (c *InMemorySessions) DropIndex(key ...string) error {
	//TODO
	fmt.Println("DropIndex " + c.Name)
	fmt.Println(key)
	return nil
}
func (c *InMemorySessions) Sort(fields ...string) CollectionLike {
	//TODO
	fmt.Println("Sort " + c.Name)
	fmt.Println(fields)
	return c
}
func (c *InMemorySessions) Select(selector interface{}) CollectionLike {
	//TODO
	fmt.Println("Select " + c.Name)
	fmt.Println(selector)
	return c
}

func (c *InMemorySessions) Limit(n int) CollectionLike {
	//TODO
	fmt.Println("Limit " + c.Name)
	fmt.Println(n)
	return c
}

func (c *InMemorySessions) Distinct(key string, result interface{}) error {
	//TODO
	fmt.Println("Distinct " + c.Name)
	fmt.Println(key)
	return nil
}

func (c *InMemorySessions) One(result interface{}) error {
	fmt.Println("One " + c.Name)

	_, filterIdIsSet := c.CurrentFilter["SessionId"]
	if filterIdIsSet {
		filterByObjectId := c.CurrentFilter["SessionId"].(string)
		session, existInDatabase := c.database[filterByObjectId]
		if existInDatabase {
			valuePtr := reflect.ValueOf(result)
			valuePtr.Elem().Set(reflect.ValueOf(session))
		}
	}

	return nil
}

func (c *InMemorySessions) All(result interface{}) error {
	fmt.Println("All" + c.Name)
	return nil
}

func (c *InMemorySessions) UpdateAll(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	fmt.Println("UpdateAll " + c.Name)
	fmt.Println(selector)
	fmt.Println(update)
	_, filterIdIsSet := selector.(bson.M)["SessionId"]
	_, isSetOperation := update.(bson.M)["$set"]
	if filterIdIsSet && isSetOperation {
		filterByObjectId := selector.(bson.M)["SessionId"].(string)
		session, existInDatabase := c.database[filterByObjectId]
		_, setLoginTime := update.(bson.M)["$set"].(bson.M)["LoginTimes"]
		if existInDatabase && setLoginTime {
			session.LoginTimes = update.(bson.M)["$set"].(bson.M)["LoginTimes"].(int)
		}
		_, setCaptcha := update.(bson.M)["$set"].(bson.M)["Captcha"]
		if existInDatabase && setCaptcha {
			session.Captcha = update.(bson.M)["$set"].(bson.M)["Captcha"].(string)
		}
		_, setUserId := update.(bson.M)["$set"].(bson.M)["UserId"]
		if existInDatabase && setUserId {
			session.UserId = update.(bson.M)["$set"].(bson.M)["UserId"].(string)
		}
		c.database[filterByObjectId] = session
	}
	info = &mgo.ChangeInfo{Updated: 1}
	return info, err
}

func (c *InMemorySessions) Insert(docs ...interface{}) error {
	//TODO
	fmt.Println("Insert " + c.Name)
	for _, doc := range docs {
		fmt.Println(doc)
		c.database[doc.(info.Session).SessionId] = doc.(info.Session)
	}
	return nil
}

func (c *InMemorySessions) Update(selector interface{}, update interface{}) error {
	//TODO
	fmt.Println("Update " + c.Name)
	fmt.Println(selector)
	fmt.Println(update)
	return nil
}

func (c *InMemorySessions) Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	fmt.Println("Upsert " + c.Name)
	info = &mgo.ChangeInfo{Updated: 1}
	return info, err
}

func (c *InMemorySessions) Remove(selector interface{}) error {
	//TODO
	fmt.Println("Remove " + c.Name)
	return nil
}

func (c *InMemorySessions) RemoveAll(selector interface{}) (info *mgo.ChangeInfo, err error) {
	//TODO
	fmt.Println("RemoveAll " + c.Name)
	fmt.Println(selector)
	info = &mgo.ChangeInfo{Updated: 1}
	return info, nil
}
