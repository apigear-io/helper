package helper

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// Join joins any number of path elements into a single path.
// If any element is an absolute path, Join returns the absolute path of that element and ignores previous elements.
func Join(elem ...string) string {
	// if last elem isAbs then return it
	// if any elem is abs path, join from there
	for i := len(elem) - 1; i >= 0; i-- {
		if filepath.IsAbs(elem[i]) {
			return filepath.Join(elem[i:]...)
		}
	}
	return filepath.Join(elem...)
}

func RmDir(src string) error {
	return os.RemoveAll(src)
}

// IsFile returns true if the path is a file.
func IsFile(src string) bool {
	fi, err := os.Stat(src)
	if err != nil {
		return false
	}
	return !fi.IsDir()
}

// IsDir returns true if the path is a directory.
func IsDir(src string) bool {
	fi, err := os.Stat(src)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

func FileExists(src string) bool {
	_, err := os.Stat(src)
	return err == nil || os.IsExist(err)
}

func CopyDir(src, dst string) error {
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		s := Join(src, entry.Name())
		t := Join(dst, entry.Name())
		if entry.IsDir() {
			if err := CopyDir(s, t); err != nil {
				return err
			}
		} else {
			if err := CopyFile(s, t); err != nil {
				return err
			}
		}
	}
	return nil
}

func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	// ensure target directory exists
	err = os.MkdirAll(filepath.Dir(dst), 0755)
	if err != nil {
		return err
	}
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	return err
}

func MkDir(src string) error {
	return os.MkdirAll(src, 0755)
}
func WriteFile(dst string, data []byte) error {
	return os.WriteFile(dst, data, 0644)
}

func ReadFile(src string) ([]byte, error) {
	return os.ReadFile(src)
}

func HasExt(src string, ext string) bool {
	return filepath.Ext(src) == ext
}

func Ext(file string) string {
	return filepath.Ext(file)
}

func ListDir(src string) ([]fs.DirEntry, error) {
	return os.ReadDir(src)
}

func FindFile(src string, name string) (fs.FileInfo, error) {
	// walk recursively
	var found fs.FileInfo
	err := filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Name() == name {
			found = info
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if found == nil {
		return nil, fmt.Errorf("file not found: %s", name)
	}
	return found, nil
}

// FilterDir returns a list of files in a directory that match filter
func FilterDir(src string, filter func(fs.DirEntry) bool) ([]fs.DirEntry, error) {
	entries, err := os.ReadDir(src)
	if err != nil {
		return nil, err
	}
	var filtered []fs.DirEntry
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if filter(entry) {
			filtered = append(filtered, entry)
		}
	}
	return filtered, nil
}

// RmFile removes a file.
func RmFile(src string) error {
	return os.Remove(src)
}

// Rename renames a file or directory.
func Rename(src, dst string) error {
	return os.Rename(src, dst)
}
