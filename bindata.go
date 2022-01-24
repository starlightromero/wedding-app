// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// public/views/home.html
package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _publicViewsHomeHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x94\x31\x6f\xdb\x30\x10\x85\x67\xf9\x57\x1c\xb4\x64\x33\x13\x77\x33\x68\x4d\x45\x33\x35\x0d\x62\x20\x99\x29\xf1\x6c\x12\xa5\x48\x96\x3c\x59\x51\x7f\x7d\x41\x4a\x72\xe4\xb8\x09\x10\x0f\xc6\xf1\xa8\x7b\xdf\x13\xf0\x4e\x5c\x51\x6b\xaa\x55\xc1\x6b\x27\x87\x6a\x55\x14\x5c\xdd\x55\x2f\x68\x1a\xd7\x22\x90\x83\x3d\x89\x60\xf4\x51\x11\x08\x2b\x61\xaf\xe5\x4d\x84\x17\x94\x52\xdb\x23\x67\xea\x2e\x4f\xf8\xea\x87\x36\x06\x5c\x47\x40\x0a\xe1\xe0\x42\x0b\x35\x1a\xd7\x27\x81\xa7\xfd\xf3\xe3\x1a\x1e\x0d\x8a\x88\xf9\x00\xf5\x00\x3f\x45\x68\x14\x7c\xbb\x25\xb5\xe6\xcc\x67\x91\x34\x95\x8a\x82\x1b\x51\xa3\x49\x2a\xbb\xd2\x8a\x16\xcb\xea\x41\xb4\xb8\x05\xce\xf2\x45\xb5\x02\x48\x4f\x69\xeb\x13\x6f\xf0\xb8\x2b\x09\x5f\xa9\x04\x2d\xa7\x01\x48\xff\xf3\x30\xaf\xc3\x28\xeb\xab\xef\x1a\x49\x84\x01\x9e\x30\x52\xd0\x0d\x69\x67\xe3\x76\xe2\x5f\x0a\x36\x0a\x9b\xdf\xb5\x7b\x9d\x44\x9d\x7d\x13\xcd\xf5\x49\x98\x6e\x3e\x5c\x9b\xce\xdd\x07\x67\x71\xb6\x7c\xf6\xf0\x31\xe3\x84\xc7\x64\x4e\x0b\x3b\x93\x96\x9d\x89\xb7\x68\x5d\x51\x97\x77\xcf\xe7\xfa\x6b\x0e\x2e\xe0\x97\xdc\x0f\x90\x13\xed\x4b\x20\x47\x0a\x43\x6e\xcd\xb4\x65\x67\x42\x2e\x5a\x57\xdc\xe5\xdd\xaf\x54\x9f\xf3\x15\x3d\x36\xfa\x30\x8c\xe9\x5b\x7f\xea\xe9\x2d\x33\x59\xee\xc2\xca\x98\x9a\xff\x8e\xc5\xae\x6e\x35\x9d\x6d\xee\xc7\x63\x4e\x30\x9b\x23\xcc\x7d\x75\xaf\x0f\x14\x41\x04\xb4\x37\x04\x01\xff\x74\x3a\xa0\x84\xba\xa3\xd4\x83\x13\x86\x01\xda\xae\x51\x20\xbc\x0f\xd8\x68\x41\x28\xdf\x6d\x02\xcc\xbf\xa2\xe0\x75\x47\xe4\xec\x3b\x07\xe9\x31\x91\x43\xbc\x2b\x15\x91\x8f\x5b\xc6\xfa\xbe\x5f\x8b\x56\xfc\x75\x76\xdd\xb8\x96\xf5\xe3\xa2\xb2\xa8\x44\x40\x16\xb5\x14\x56\xc6\x79\xa1\x37\xb7\x9b\x4d\x59\xdd\xbb\xbc\xa4\x78\xd4\x91\xc2\xc0\xd9\x88\x9a\xf1\x8b\xb7\xe2\x6c\xfc\x42\x70\x36\x7e\x31\xfe\x05\x00\x00\xff\xff\xea\xd6\x3d\x00\x39\x04\x00\x00")

func publicViewsHomeHtmlBytes() ([]byte, error) {
	return bindataRead(
		_publicViewsHomeHtml,
		"public/views/home.html",
	)
}

func publicViewsHomeHtml() (*asset, error) {
	bytes, err := publicViewsHomeHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "public/views/home.html", size: 1081, mode: os.FileMode(436), modTime: time.Unix(1642897957, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"public/views/home.html": publicViewsHomeHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"public": &bintree{nil, map[string]*bintree{
		"views": &bintree{nil, map[string]*bintree{
			"home.html": &bintree{publicViewsHomeHtml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
