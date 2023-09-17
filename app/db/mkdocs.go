package db

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/leanote/leanote/app/lea"
	"github.com/revel/revel"
	"gopkg.in/yaml.v3"
)

type Mkdocs struct {
	Dir *os.File
}

type Author struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Avatar      string `yaml:"avatar"`
	Pwd         string `yaml:"password_hash"`
}

type Category struct {
	FilePath string
	Name     string
	Usn      int
	Mardowns []Mardown
}

type Mardown struct {
	FilePath string
	Name     string
	Usn      int
	//Content  []byte
	ModTime time.Time
}

func (c *Mkdocs) ReadAuthorFile() map[string]Author {
	filePath := filepath.Join(c.Dir.Name(), ".authors.yml")

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, ".authors.yml not found\n")
		return map[string]Author{"admin": {Name: "admin", Pwd: "e99a18c428cb38d5f260853678922e03"}}
	}

	var m map[string]Author
	err = yaml.Unmarshal(data, &m)
	if err != nil {
		panic(err)
	}

	return m
}

func (c *Mkdocs) WalkDirectory() []Category {

	dirs, err := ioutil.ReadDir(c.Dir.Name())
	result := make([]Category, 0, len(dirs))
	if err != nil {
		panic(err)
	}
	for i, dir := range dirs {
		if dir.IsDir() && unicode.IsUpper([]rune(dir.Name())[0]) {
			path := filepath.Join(c.Dir.Name(), dir.Name())
			notes, err := ioutil.ReadDir(path)
			mardowns := make([]Mardown, 0, len(notes))
			if err != nil {
				panic(err)
			}
			for j, file := range notes {
				subpath := filepath.Join(c.Dir.Name(), dir.Name(), file.Name())
				//data, _ := ioutil.ReadFile(subpath)
				markdown := Mardown{
					Name:     file.Name(),
					Usn:      j,
					FilePath: subpath,
					//Content:  data,
					ModTime: file.ModTime(),
				}
				mardowns = append(mardowns, markdown)
			}
			category := Category{
				Name:     dir.Name(),
				Usn:      i,
				Mardowns: mardowns,
				FilePath: path,
			}
			result = append(result, category)
		}
	}
	return result
}

func (c *Mkdocs) listDirectory() []string {

	dirs, err := ioutil.ReadDir(filepath.Join(c.Dir.Name(), "images"))
	result := make([]string, 0, len(dirs))
	if err != nil {
		panic(err)
	}
	for _, dir := range dirs {
		if dir.IsDir() {
			result = append(result, dir.Name())
		}
	}
	return result
}

func (c *Mkdocs) listAllFiles(album string) []string {
	root := filepath.Join(c.Dir.Name(), "images", album)
	result := make([]string, 0)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			relpath, _ := filepath.Rel(c.Dir.Name(), path)
			result = append(result, filepath.Join("docs", relpath))
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return result
}

func (c *Mkdocs) createDirectory(name string) {

	err := os.Mkdir(filepath.Join(c.Dir.Name(), strings.Title(name)), 0755)
	if err != nil {
		panic(err)
	}
}

func (c *Mkdocs) RenameInPath(fullpath string, newTitle string) error {
	parent := filepath.Dir(fullpath)
	dest := filepath.Join(parent, newTitle)
	err := os.Rename(fullpath, dest)
	return err
}

func (c *Mkdocs) WriteImage(basename string, path string) string {
	base, _ := lea.SplitFilename(basename)
	if _, err := strconv.ParseFloat(base, 64); err == nil {
		// replace to avoid colision (every image from api are called "0.jpeg"...)
		basename = filepath.Base(path)
	}
	md5path := lea.Md5(basename)
	imageDir := filepath.Join(c.Dir.Name(), "images", md5path[:1], md5path[:2])
	lea.MkdirAll(imageDir)
	destFile := filepath.Join(imageDir, basename)
	srcFile := filepath.Join(revel.BasePath, path)
	fmt.Println(srcFile + " -> " + destFile)
	_, err := lea.CopyFile(srcFile, destFile)
	if err != nil {
		panic(err)
	}
	return strings.Trim(strings.Replace(destFile, revel.BasePath, "", 1), "/")
}

func (c *Mkdocs) WriteFile(notebook string, filename string, Content string) {
	fullpath := filepath.Join(c.Dir.Name(), strings.Title(notebook), filename)
	fmt.Println("write " + fullpath)
	err := ioutil.WriteFile(fullpath, []byte(Content), 0755)
	if err != nil {
		panic(err)
	}
}

func (c *Mkdocs) GetTags(data []byte) []string {
	tags := make([]string, 0)
	re := regexp.MustCompile(`tags: \[(.+)\]`)

	match := re.FindStringSubmatch(string(data))
	if len(match) < 1 {
		return tags
	}
	parts := strings.Split(match[1], ",")
	for _, part := range parts {
		tags = append(tags, strings.Trim(part, " "))
	}
	return tags
}

func (c *Mkdocs) Contains(tags []string, search string) bool {
	found := false
	for _, tag := range tags {
		if tag == search {
			found = true
		}
	}
	return found
}
