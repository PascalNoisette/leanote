package db

import (
	"fmt"
	"os"
	"strings"

	. "github.com/leanote/leanote/app/lea"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
)

var dir *os.File

// 各个表的Collection对象
var Notebooks CollectionLike
var Notes CollectionLike
var NoteContents CollectionLike
var NoteContentHistories CollectionLike

var ShareNotes CollectionLike
var ShareNotebooks CollectionLike
var HasShareNotes CollectionLike
var Blogs CollectionLike
var Users CollectionLike
var Groups CollectionLike
var GroupUsers CollectionLike

var Tags CollectionLike
var NoteTags CollectionLike
var TagCounts CollectionLike

var UserBlogs CollectionLike

var Tokens CollectionLike

var Suggestions CollectionLike

// Album & file(image)
var Albums CollectionLike
var Files CollectionLike
var Attachs CollectionLike

var NoteImages CollectionLike
var Configs CollectionLike
var EmailLogs CollectionLike

// blog
var BlogLikes CollectionLike
var BlogComments CollectionLike
var Reports CollectionLike
var BlogSingles CollectionLike
var Themes CollectionLike

var Sessions CollectionLike

// 初始化时连接数据库
func Init(url, dbname string) {
	ok := true
	config := revel.Config
	if url == "" {
		url, ok = config.String("db.url")
		if !ok {
			url, ok = config.String("db.urlEnv")
			if ok {
				Log("get db conf from urlEnv: " + url)
			}
		} else {
			Log("get db conf from db.url: " + url)
		}

		if ok {
			// get dbname from urlEnv
			urls := strings.Split(url, "/")
			dbname = urls[len(urls)-1]
		}
	}
	if dbname == "" {
		dbname, _ = config.String("db.dbname")
	}

	// get db config from host, port, username, password
	if !ok {
		host, _ := revel.Config.String("db.host")
		port, _ := revel.Config.String("db.port")
		username, _ := revel.Config.String("db.username")
		password, _ := revel.Config.String("db.password")
		usernameAndPassword := username + ":" + password + "@"
		if username == "" || password == "" {
			usernameAndPassword = ""
		}
		url = "mongodb://" + usernameAndPassword + host + ":" + port + "/" + dbname
	}
	fmt.Println(url)

	// Use url as folder
	dir, err := os.Open(url)
	if err != nil {
		panic(err)
	}

	mkdocs := &Mkdocs{Dir: dir}
	writeHistory := &WriteHistory{OldIds: make(map[string]string)}

	// notebook
	//Notebooks = &FolderNotebooks{Name: "notebooks", Dir: dir, Fallback: &BsonReader{Name: "notebooks", Dir: dir}}
	//Notebooks = &BsonReader{Name: "notebooks", Dir: dir}
	Notebooks = &FolderNotebooks{Name: "notebooks", Mkdocs: mkdocs, WriteHistory: writeHistory}

	// notes
	Notes = &FileNotes{Name: "notes", Mkdocs: mkdocs, FolderNotebooks: Notebooks, WriteHistory: writeHistory}
	//Notes = &BsonReader{Name: "notes", Dir: dir}

	// noteContents
	NoteContents = &FileNoteContents{Name: "note_contents", Mkdocs: mkdocs, FolderNotebooks: Notebooks, FileNotes: Notes, WriteHistory: writeHistory}
	NoteContentHistories = &BsonReader{Name: "note_content_histories", Dir: dir}

	// share
	ShareNotes = &BsonReader{Name: "share_notes", Dir: dir}
	ShareNotebooks = &BsonReader{Name: "share_notebooks", Dir: dir}
	HasShareNotes = &BsonReader{Name: "has_share_notes", Dir: dir}

	// user
	Users = &FileUsers{Name: "users", Mkdocs: mkdocs, Fallback: &BsonReader{Name: "users", Dir: dir}}
	// group
	Groups = &BsonReader{Name: "groups", Dir: dir}
	GroupUsers = &BsonReader{Name: "group_users", Dir: dir}

	// blog
	Blogs = &BsonReader{Name: "blogs", Dir: dir}

	// tag
	Tags = &BsonReader{Name: "tags", Dir: dir}
	NoteTags = &ParsedTags{Name: "note_tags", Mkdocs: mkdocs, FileNoteContents: NoteContents}
	TagCounts = &BsonReader{Name: "tag_count", Dir: dir}

	// blog
	UserBlogs = &BsonReader{Name: "user_blogs", Dir: dir}
	BlogSingles = &BsonReader{Name: "blog_singles", Dir: dir}
	Themes = &BsonReader{Name: "themes", Dir: dir}

	// find password
	Tokens = &BsonReader{Name: "tokens", Dir: dir}

	// Suggestion
	Suggestions = &BsonReader{Name: "suggestions", Dir: dir}

	// Album & file
	Albums = &FolderImages{Name: "albums", Mkdocs: mkdocs}
	Files = &FileImages{Name: "files", Mkdocs: mkdocs, WriteHistory: writeHistory}
	Attachs = &ParsedAttachs{Name: "attachs", Mkdocs: mkdocs, FileNoteContents: NoteContents}

	NoteImages = &BsonReader{Name: "note_images", Dir: dir}

	Configs = &BsonReader{Name: "configs", Dir: dir}
	EmailLogs = &BsonReader{Name: "email_logs", Dir: dir}

	// 社交
	BlogLikes = &BsonReader{Name: "blog_likes", Dir: dir}
	BlogComments = &BsonReader{Name: "blog_comments", Dir: dir}

	// 举报
	Reports = &BsonReader{Name: "reports", Dir: dir}

	// mgo.Session
	Sessions = &BsonReader{Name: "sessions", Dir: dir}
}

func close() {
	defer dir.Close()
}

// common DAO
// 公用方法

//----------------------

func Insert(collection CollectionLike, i interface{}) bool {
	err := collection.Insert(i)
	return Err(err)
}

//----------------------

// 适合一条记录全部更新
func Update(collection CollectionLike, query interface{}, i interface{}) bool {
	err := collection.Update(query, i)
	return Err(err)
}
func Upsert(collection CollectionLike, query interface{}, i interface{}) bool {
	_, err := collection.Upsert(query, i)
	return Err(err)
}
func UpdateAll(collection CollectionLike, query interface{}, i interface{}) bool {
	_, err := collection.UpdateAll(query, i)
	return Err(err)
}
func UpdateByIdAndUserId(collection CollectionLike, id, userId string, i interface{}) bool {
	err := collection.Update(GetIdAndUserIdQ(id, userId), i)
	return Err(err)
}

func UpdateByIdAndUserId2(collection CollectionLike, id, userId bson.ObjectId, i interface{}) bool {
	err := collection.Update(GetIdAndUserIdBsonQ(id, userId), i)
	return Err(err)
}
func UpdateByIdAndUserIdField(collection CollectionLike, id, userId, field string, value interface{}) bool {
	return UpdateByIdAndUserId(collection, id, userId, bson.M{"$set": bson.M{field: value}})
}
func UpdateByIdAndUserIdMap(collection CollectionLike, id, userId string, v bson.M) bool {
	return UpdateByIdAndUserId(collection, id, userId, bson.M{"$set": v})
}

func UpdateByIdAndUserIdField2(collection CollectionLike, id, userId bson.ObjectId, field string, value interface{}) bool {
	return UpdateByIdAndUserId2(collection, id, userId, bson.M{"$set": bson.M{field: value}})
}
func UpdateByIdAndUserIdMap2(collection CollectionLike, id, userId bson.ObjectId, v bson.M) bool {
	return UpdateByIdAndUserId2(collection, id, userId, bson.M{"$set": v})
}

func UpdateByQField(collection CollectionLike, q interface{}, field string, value interface{}) bool {
	_, err := collection.UpdateAll(q, bson.M{"$set": bson.M{field: value}})
	return Err(err)
}
func UpdateByQI(collection CollectionLike, q interface{}, v interface{}) bool {
	_, err := collection.UpdateAll(q, bson.M{"$set": v})
	return Err(err)
}

// 查询条件和值
func UpdateByQMap(collection CollectionLike, q interface{}, v interface{}) bool {
	_, err := collection.UpdateAll(q, bson.M{"$set": v})
	return Err(err)
}

//------------------------

// 删除一条
func Delete(collection CollectionLike, q interface{}) bool {
	err := collection.Remove(q)
	return Err(err)
}
func DeleteByIdAndUserId(collection CollectionLike, id, userId string) bool {
	err := collection.Remove(GetIdAndUserIdQ(id, userId))
	return Err(err)
}
func DeleteByIdAndUserId2(collection CollectionLike, id, userId bson.ObjectId) bool {
	err := collection.Remove(GetIdAndUserIdBsonQ(id, userId))
	return Err(err)
}

// 删除所有
func DeleteAllByIdAndUserId(collection CollectionLike, id, userId string) bool {
	_, err := collection.RemoveAll(GetIdAndUserIdQ(id, userId))
	return Err(err)
}
func DeleteAllByIdAndUserId2(collection CollectionLike, id, userId bson.ObjectId) bool {
	_, err := collection.RemoveAll(GetIdAndUserIdBsonQ(id, userId))
	return Err(err)
}

func DeleteAll(collection CollectionLike, q interface{}) bool {
	_, err := collection.RemoveAll(q)
	return Err(err)
}

//-------------------------

func Get(collection CollectionLike, id string, i interface{}) {
	collection.FindId(bson.ObjectIdHex(id)).One(i)
}
func Get2(collection CollectionLike, id bson.ObjectId, i interface{}) {
	collection.FindId(id).One(i)
}

func GetByQ(collection CollectionLike, q interface{}, i interface{}) {
	collection.Find(q).One(i)
}
func ListByQ(collection CollectionLike, q interface{}, i interface{}) {
	collection.Find(q).All(i)
}

func ListByQLimit(collection CollectionLike, q interface{}, i interface{}, limit int) {
	collection.Find(q).Limit(limit).All(i)
}

// 查询某些字段, q是查询条件, fields是字段名列表
func GetByQWithFields(collection CollectionLike, q bson.M, fields []string, i interface{}) {
	selector := make(bson.M, len(fields))
	for _, field := range fields {
		selector[field] = true
	}
	collection.Find(q).Select(selector).One(i)
}

// 查询某些字段, q是查询条件, fields是字段名列表
func ListByQWithFields(collection CollectionLike, q bson.M, fields []string, i interface{}) {
	selector := make(bson.M, len(fields))
	for _, field := range fields {
		selector[field] = true
	}
	collection.Find(q).Select(selector).All(i)
}
func GetByIdAndUserId(collection CollectionLike, id, userId string, i interface{}) {
	collection.Find(GetIdAndUserIdQ(id, userId)).One(i)
}
func GetByIdAndUserId2(collection CollectionLike, id, userId bson.ObjectId, i interface{}) {
	collection.Find(GetIdAndUserIdBsonQ(id, userId)).One(i)
}

// 按field去重
func Distinct(collection CollectionLike, q bson.M, field string, i interface{}) {
	collection.Find(q).Distinct(field, i)
}

//----------------------

func Count(collection CollectionLike, q interface{}) int {
	cnt, err := collection.Find(q).Count()
	if err != nil {
		Err(err)
	}
	return cnt
}

func Has(collection CollectionLike, q interface{}) bool {
	if Count(collection, q) > 0 {
		return true
	}
	return false
}

//-----------------

// 得到主键和userId的复合查询条件
func GetIdAndUserIdQ(id, userId string) bson.M {
	return bson.M{"_id": bson.ObjectIdHex(id), "UserId": bson.ObjectIdHex(userId)}
}
func GetIdAndUserIdBsonQ(id, userId bson.ObjectId) bson.M {
	return bson.M{"_id": id, "UserId": userId}
}

// DB处理错误
func Err(err error) bool {
	if err != nil {
		fmt.Println(err)
		// 删除时, 查找
		if err.Error() == "not found" {
			return true
		}
		return false
	}
	return true
}

// 检查mognodb是否lost connection
// 每个请求之前都要检查!!
func CheckMongoSessionLost() {
	// nothing to do
}
