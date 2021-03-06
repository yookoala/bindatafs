// Code generated by go-bindata.
// sources:
// assets/hello/bar.txt
// assets/hello/world.txt
// assets/hello.txt
// assets/index.html
// DO NOT EDIT!

package example1

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _helloBarTxt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xf2\x48\xcd\xc9\xc9\x57\x70\x4a\x2c\xe2\x02\x04\x00\x00\xff\xff\xec\x29\xf6\x0e\x0a\x00\x00\x00")

func helloBarTxtBytes() ([]byte, error) {
	return bindataRead(
		_helloBarTxt,
		"hello/bar.txt",
	)
}

func helloBarTxt() (*asset, error) {
	bytes, err := helloBarTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "hello/bar.txt", size: 10, mode: os.FileMode(436), modTime: time.Unix(1479662730, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _helloWorldTxt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x0a\xcf\x2f\xca\x49\x51\x54\x54\xe4\x02\x04\x00\x00\xff\xff\xd9\x4d\xc9\xdc\x09\x00\x00\x00")

func helloWorldTxtBytes() ([]byte, error) {
	return bindataRead(
		_helloWorldTxt,
		"hello/world.txt",
	)
}

func helloWorldTxt() (*asset, error) {
	bytes, err := helloWorldTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "hello/world.txt", size: 9, mode: os.FileMode(436), modTime: time.Unix(1479662730, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _helloTxt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xf2\x48\xcd\xc9\xc9\x57\x08\xcf\x2f\xca\x49\xe1\x02\x04\x00\x00\xff\xff\xe3\xe5\x95\xb0\x0c\x00\x00\x00")

func helloTxtBytes() ([]byte, error) {
	return bindataRead(
		_helloTxt,
		"hello.txt",
	)
}

func helloTxt() (*asset, error) {
	bytes, err := helloTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "hello.txt", size: 12, mode: os.FileMode(436), modTime: time.Unix(1479662730, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x8f\x4b\x6e\xc5\x20\x0c\x45\xe7\x59\x85\x9b\x79\x6b\x65\x5a\xb9\xcc\xba\x83\x6e\xc0\x09\x4e\x40\x82\x80\xc0\xfd\x64\xf7\x25\xe1\x29\x7a\x12\x12\x57\xc0\xe1\x1e\xd3\x8b\x4d\x8b\x1e\x59\xcc\x40\x4e\x63\x38\x37\x61\x6b\x06\x00\x52\xaf\x41\xcc\xe7\x1f\xc7\x1c\x04\xb8\x56\x51\x58\x53\x81\xd9\xef\x96\x95\xd7\x4a\xd8\x9f\x0c\x84\x1d\xa2\x39\xd9\xe3\x62\x23\xfb\xfd\x0c\x2d\xba\xc9\xdc\x04\x48\xff\xad\x01\xd3\xe3\x3a\x9b\x2f\xe7\x2b\xb4\x95\xf6\x70\x00\xb7\x8a\x60\xa5\x80\x26\x50\xa9\x0a\xbf\x5e\x1d\x10\x83\x2b\xb2\x7e\x8c\x4e\x35\xd7\x77\xc4\xad\x9d\x7e\xcf\x6f\x4b\x8a\xb8\xa5\xd7\x2a\xe5\x47\xf0\x6e\x19\xcd\x93\x22\x9b\x56\x71\x2a\x61\x77\x22\xec\x92\x4d\xe1\x9a\xf7\x3f\x00\x00\xff\xff\x40\x9b\x81\x7b\x02\x01\x00\x00")

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 258, mode: os.FileMode(436), modTime: time.Unix(1479662730, 0)}
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
	"hello/bar.txt": helloBarTxt,
	"hello/world.txt": helloWorldTxt,
	"hello.txt": helloTxt,
	"index.html": indexHtml,
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
	"hello": &bintree{nil, map[string]*bintree{
		"bar.txt": &bintree{helloBarTxt, map[string]*bintree{}},
		"world.txt": &bintree{helloWorldTxt, map[string]*bintree{}},
	}},
	"hello.txt": &bintree{helloTxt, map[string]*bintree{}},
	"index.html": &bintree{indexHtml, map[string]*bintree{}},
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

