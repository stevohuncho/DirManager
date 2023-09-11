package dirman

import (
	"fmt"
	"io/fs"
	"path/filepath"

	gofile "github.com/stevohuncho/GoFile"
)

type DirMan struct {
	root string
}

// creates manager and saves the path as the root dir
func NewDirManager(path string) (*DirMan, error) {
	if !gofile.Exists(path) {
		return nil, fmt.Errorf("failed to find path: %s", path)
	}
	dm := DirMan{
		root: path,
	}
	return &dm, nil
}

// makes new dir
func (dm *DirMan) Mkdir(path string) error {
	newPath := filepath.Join(dm.root, path)
	return gofile.SafeDir(newPath)
}

// removes file/dir
func (dm *DirMan) Rm(path string) error {
	newPath := filepath.Join(dm.root, path)
	return gofile.Rm(newPath)
}

// removes file/dir
func (dm *DirMan) RmAll(path string) error {
	newPath := filepath.Join(dm.root, path)
	return gofile.RmAll(newPath)
}

// writes data to path
func (dm *DirMan) WriteJson(path string, d any) error {
	newPath := filepath.Join(dm.root, path)
	return gofile.WriteJson(newPath, d)
}

func (dm *DirMan) WriteIndentedJson(path string, d any) error {
	newPath := filepath.Join(dm.root, path)
	return gofile.WriteIndentedJson(newPath, d)
}

// retrieves all files/dirs filtered by fastfile filters
func (dm *DirMan) Ls(path string, opts ...gofile.ReadDirOptFunc) ([]fs.DirEntry, error) {
	newPath := filepath.Join(dm.root, path)
	return gofile.ReadDir(newPath, opts...)
}

// gets the string data of the file
func (dm *DirMan) String(path string) (string, error) {
	newPath := filepath.Join(dm.root, path)
	return gofile.String(newPath)
}

// gets json data of the file and saves it to the pointer
func (dm *DirMan) Json(path string, p interface{}) error {
	newPath := filepath.Join(dm.root, path)
	return gofile.Json(newPath, p)
}

// gets the csv data of the file
func (dm *DirMan) Csv(path string) ([][]string, error) {
	newPath := filepath.Join(dm.root, path)
	return gofile.Csv(newPath)
}

// creates a new manager relative to the root of the current
func (dm *DirMan) New(path string) (*DirMan, error) {
	newPath := filepath.Join(dm.root, path)
	if !gofile.Exists(newPath) {
		return nil, fmt.Errorf("failed to find path: %s", newPath)
	}
	newDm := DirMan{
		root: newPath,
	}
	return &newDm, nil
}
