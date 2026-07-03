package models 

import (
	"io/fs"
	"strings"
	"path/filepath"
)


type EntryInfo struct {
	BaseName string
	Path	 string 
	AbsolutePath string
	Type	 string 
	Ext 	 string
}

func NewEntryInfo(baseName, path, absolutePath, entryType, ext string) EntryInfo {
	if (entryType == "directory" && !strings.HasSuffix(path, "/")) {
		path = path + "/"
	}

	return EntryInfo{
		BaseName:     baseName,
		Path:         path,
		AbsolutePath: absolutePath,
		Type:         entryType,
		Ext:          ext,
	}
}

func NewEntryInfoFromDirEntry(rootPath string, path string, d fs.DirEntry) EntryInfo {
	absolutePath := filepath.Join(rootPath, path)
	entryType := "file"

	if d.IsDir() {
		entryType = "directory"
	}

	return NewEntryInfo(d.Name(), path, absolutePath, entryType, filepath.Ext(d.Name()))
}

func NewEntryInfoFromFileInfo(rootPath string, path string, d fs.FileInfo) EntryInfo {
	absolutePath := filepath.Join(rootPath, path)
	entryType := "file"

	if d.IsDir() {
		entryType = "directory"
	}

	return NewEntryInfo(d.Name(), path, absolutePath, entryType, filepath.Ext(d.Name()))
}

