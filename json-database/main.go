package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/jcelliott/lumber"
)

const (
	DIR     = "./"
	VERSION = "1.0.0"
)

type (
	Logger interface {
		Fatal(string, ...interface{})
		Error(string, ...interface{})
		Warn(string, ...interface{})
		Info(string, ...interface{})
		Debug(string, ...interface{})
		Trace(string, ...interface{})
	}

	Driver struct {
		mutex   sync.Mutex
		mutexes map[string]*sync.Mutex
		dir     string
		log     Logger
	}
)

type Options struct {
	Logger
}

func NewDatabase(dir string, options *Options) (*Driver, error) {
	dir = filepath.Clean(dir)
	opts := Options{}

	if options != nil {
		opts = *options
	}
	if opts.Logger == nil {
		opts.Logger = lumber.NewConsoleLogger((lumber.INFO))
	}

	driver := Driver{
		dir:     dir,
		mutexes: make(map[string]*sync.Mutex),
		log:     opts.Logger,
	}

	if _, err := os.Stat(dir); err == nil {
		opts.Logger.Debug("Using '%s' (Database Already Exists)\n", dir)
		return &driver, nil
	}
	opts.Logger.Debug("Creating Database in directory '%s", dir)
	return &driver, os.MkdirAll(dir, 0755)
}

func (d *Driver) Write(collection, resource string, v interface{}) error {
	if collection == "" {
		return fmt.Errorf("Missing Collection - No place to save the record!")
	}
	if resource == "" {
		return fmt.Errorf("Missing Resource - Unable to save Record (No Name)")
	}
	mutex := d.GetOrCreateMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(d.dir, collection)
	finalPath := filepath.Join(dir, resource+".json")
	tempPath := filepath.Join(finalPath + ".tmp")

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	b, err := json.MarshalIndent(v, "", "\t")
	CheckError(err)

	b = append(b, byte('\n'))
	if err := ioutil.WriteFile(tempPath, b, 0644); err != nil {
		return err
	}

	return os.Rename(tempPath, finalPath)
}

func (d *Driver) Read(collection, resource string, v interface{}) error {
	if collection == "" {
		return fmt.Errorf("Missing Collcetion - Unable to Read")
	}
	if resource == "" {
		return fmt.Errorf("Missing Resource - Unable to read record ( No Name )")
	}

	record := filepath.Join(d.dir, collection, resource)

	if _, err := stat(record); err != nil {
		return err
	}

	b, err := ioutil.ReadFile(record + ",json")
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &v)
}

func (d *Driver) ReadAll(collection string) ([]string, error) {
	if collection == "" {
		return nil, fmt.Errorf("Missing Collcetion - Unable to Read")
	}
	dir := filepath.Join(d.dir, collection)

	if _, err := stat(dir); err != nil {
		return nil, err
	}

	files, _ := ioutil.ReadDir(dir)
	var records []string

	for _, file := range files {
		b, err := ioutil.ReadFile(filepath.Join(dir, file.Name()))
		if err != nil {
			return nil, err
		}
		records = append(records, string(b))
	}
	return records, nil
}
func (d *Driver) Delete(collection string, resource string) error {

	path := filepath.Join(collection, resource)
	mutex := d.GetOrCreateMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(d.dir, path)
	switch fi, err := stat(dir); {
	case fi == nil, err != nil:
		return fmt.Errorf("Unable to find the file '%v'\n", path)
	case fi.Mode().IsDir():
		return os.RemoveAll(dir)
	case fi.Mode().IsRegular():
		return os.RemoveAll(dir + ".json")
	}
	return nil

}

func (d *Driver) GetOrCreateMutex(collection string) *sync.Mutex {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	m, ok := d.mutexes[collection]
	if !ok {
		m = &sync.Mutex{}
		d.mutexes[collection] = m
	}
	return m
}

func stat(path string) (fi os.FileInfo, err error) {
	if fi, err = os.Stat(path); os.IsNotExist(err) {
		fi, err = os.Stat(path + ".json")
	}
	return
}

type Address struct {
	City    string
	State   string
	Country string
	Pincode json.Number
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {

	db, err := NewDatabase(DIR, nil)
	CheckError(err)

	employees := []User{
		{"Tanmay", "26", "976256272", "Apple", Address{"SJ", "California", "USA", "6273663"}},
		{"Rahul", "23", "974256272", "Apple", Address{"NJC", "NJ", "USA", "6273663"}},
		{"Adi", "24", "976226272", "Apple", Address{"NY", "NY", "USA", "6273663"}},
		{"Shanker", "25", "576256272", "Apple", Address{"Seattle", "Wash", "USA", "6273663"}},
		{"Manav", "29", "977256272", "Apple", Address{"Austin", "Texas", "USA", "6273663"}},
	}

	for _, value := range employees {
		db.Write("users", value.Name, User{
			Name:    value.Name,
			Age:     value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: value.Address,
		})
	}

	records, err := db.ReadAll("users")
	CheckError(err)

	fmt.Println(records)

	allusers := []User{}
	for _, r := range records {
		employeeFound := User{}
		err := json.Unmarshal([]byte(r), &employeeFound)
		CheckError(err)

		allusers = append(allusers, employeeFound)

	}
	fmt.Println(allusers)

	// if err := db.Delete("users", "Rahul"); err != nil {
	// 	fmt.Println("Error: ", err)
	// }

	// if err := db.Delete("users", ""); err != nil {
	// 	fmt.Println("Error: ", err)
	// }
}
