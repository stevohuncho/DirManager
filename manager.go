package dirman

import (
	"io/fs"

	"gitlab.com/tesla-scripts/go-utils/fastfile"
)

type DirMan struct {
	root string
}

// creates manager and saves the path as the root dir
func NewDirManager(path string) (*DirMan, error) {
	return nil, nil
}

// makes new dir
func (dm *DirMan) Mkdir(path string) error {
	return nil
}

// removes file/dir
func (dm *DirMan) Rm(path string) error {
	return nil
}

// writes data to path
func (dm *DirMan) WriteJson(path string, data any) error {
	return nil
}

// retrieves all files/dirs filtered by fastfile filters
func (dm *DirMan) Ls(path string, opts ...fastfile.ReadDirOptFunc) ([]fs.DirEntry, error) {
	return []fs.DirEntry{}, nil
}

// gets the string data of the file
func (dm *DirMan) String(path string) (string, error) {
	return "", nil
}

// gets json data of the file and saves it to the pointer
func (dm *DirMan) Json(path string, pointer interface{}) error {
	return nil
}

// gets the csv data of the file
func (dm *DirMan) Csv(path string) ([][]string, error) {
	return [][]string{}, nil
}

// creates a new manager relative to the root of the current
func (dm *DirMan) New(path string) (*DirMan, error) {
	return nil, nil
}
