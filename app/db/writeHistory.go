package db

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type WriteHistory struct {
	OldIds map[string]string
}

func (c *WriteHistory) RenameObjectId(old bson.ObjectId, new string) {
	fmt.Println(c)
	c.OldIds[old.Hex()] = new
}

func (c *WriteHistory) GetRealId(oldObjectId interface{}) string {
	old := oldObjectId.(bson.ObjectId)
	if newId, hasHistory := c.OldIds[old.Hex()]; hasHistory {
		return newId
	}
	return old.Hex()
}
