package db

import (
	"gopkg.in/mgo.v2"
)

type CollectionLike interface {
	UpdateAll(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error)
	Find(query interface{}) CollectionLike
	Insert(docs ...interface{}) error
	Update(selector interface{}, update interface{}) error
	Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error)
	Remove(selector interface{}) error
	RemoveAll(selector interface{}) (info *mgo.ChangeInfo, err error)
	FindId(id interface{}) CollectionLike
	Count() (n int, err error)
	One(result interface{}) (err error)
	All(result interface{}) error
	Select(selector interface{}) CollectionLike
	Limit(n int) CollectionLike
	Distinct(key string, result interface{}) error
	Sort(fields ...string) CollectionLike
	Skip(n int) CollectionLike
	DropIndex(key ...string) error
}
