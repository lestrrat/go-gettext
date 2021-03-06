package gettext

import (
	"io/ioutil"
	"path/filepath"
)

func (f SourceFunc) ReadFile(s string) ([]byte, error) {
	return f(s)
}

func NewFileSystemSource(dir string) *FileSystemSource {
	return &FileSystemSource{root: dir}
}

func (f FileSystemSource) ReadFile(s string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(f.root, s))
}
