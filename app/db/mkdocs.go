package db

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode"

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
		return nil
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

func (c *Mkdocs) createDirectory(name string) {

	err := os.Mkdir(filepath.Join(c.Dir.Name(), strings.Title(name)), 0755)
	if err != nil {
		panic(err)
	}
}

func (c *Mkdocs) WriteFile(notebook string, filename string, Content string) {
	if !strings.HasSuffix(filename, ".md") {
		filename = filename + ".md"
	}
	fullpath := filepath.Join(c.Dir.Name(), strings.Title(notebook), filename)
	fmt.Println("write " + fullpath)
	err := ioutil.WriteFile(fullpath, []byte(Content), 0755)
	if err != nil {
		panic(err)
	}
}
