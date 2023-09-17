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
	old := oldObjectId.(bson.ObjectId).Hex()
	_, hasHistory := c.OldIds[old]
	for hasHistory {
		if _, hasHistory := c.OldIds[old]; !hasHistory {
			break
		}
		old = c.OldIds[old]
	}
	return old
}
