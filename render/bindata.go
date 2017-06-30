// Code generated by go-bindata.
// sources:
// render/html/paragraph.html
// render/html/section.html
// render/html/sequence.html
// render/html/string.html
// DO NOT EDIT!

package render

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

var _renderHtmlParagraphHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x29\xb0\xab\xae\xd6\x73\xce\xcf\x2b\x49\xcd\x2b\x51\xa8\x51\x28\x4a\xcd\x4b\x49\x2d\xaa\xad\xb5\xd1\x2f\xb0\xe3\x02\x04\x00\x00\xff\xff\x7e\xf4\x37\x0c\x1d\x00\x00\x00")

func renderHtmlParagraphHtmlBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlParagraphHtml,
		"render/html/paragraph.html",
	)
}

func renderHtmlParagraphHtml() (*asset, error) {
	bytes, err := renderHtmlParagraphHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/paragraph.html", size: 29, mode: os.FileMode(420), modTime: time.Unix(1498794853, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlSectionHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\xc9\x30\xb4\xab\xae\xd6\x0b\xc9\x2c\xc9\x49\x55\xa8\x51\x28\x4a\xcd\x4b\x49\x2d\xaa\xad\xb5\xd1\xcf\x30\xb4\xe3\xe2\xaa\xae\xd6\x73\xca\x4f\xa9\x44\x92\xe0\x02\x04\x00\x00\xff\xff\xdb\x05\x16\x5a\x31\x00\x00\x00")

func renderHtmlSectionHtmlBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlSectionHtml,
		"render/html/section.html",
	)
}

func renderHtmlSectionHtml() (*asset, error) {
	bytes, err := renderHtmlSectionHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/section.html", size: 49, mode: os.FileMode(420), modTime: time.Unix(1498788510, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlSequenceHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x2e\x4a\xcc\x4b\x4f\x55\xd0\x73\xce\xcf\x2b\x49\xcd\x2b\x29\xae\xad\xad\xae\xd6\x53\xa8\x51\x28\x4a\xcd\x4b\x49\x2d\x02\xf1\x52\xf3\x52\x6a\x6b\xb9\x00\x01\x00\x00\xff\xff\xfc\xbe\x50\x06\x29\x00\x00\x00")

func renderHtmlSequenceHtmlBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlSequenceHtml,
		"render/html/sequence.html",
	)
}

func renderHtmlSequenceHtml() (*asset, error) {
	bytes, err := renderHtmlSequenceHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/sequence.html", size: 41, mode: os.FileMode(420), modTime: time.Unix(1498700481, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlStringHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x56\xd0\x0b\x2e\x29\xca\xcc\x4b\x57\xd0\xad\xad\xe5\x02\x04\x00\x00\xff\xff\x11\x3f\xd1\x9c\x0f\x00\x00\x00")

func renderHtmlStringHtmlBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlStringHtml,
		"render/html/string.html",
	)
}

func renderHtmlStringHtml() (*asset, error) {
	bytes, err := renderHtmlStringHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/string.html", size: 15, mode: os.FileMode(420), modTime: time.Unix(1498796001, 0)}
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
	"render/html/paragraph.html": renderHtmlParagraphHtml,
	"render/html/section.html": renderHtmlSectionHtml,
	"render/html/sequence.html": renderHtmlSequenceHtml,
	"render/html/string.html": renderHtmlStringHtml,
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
	"render": &bintree{nil, map[string]*bintree{
		"html": &bintree{nil, map[string]*bintree{
			"paragraph.html": &bintree{renderHtmlParagraphHtml, map[string]*bintree{}},
			"section.html": &bintree{renderHtmlSectionHtml, map[string]*bintree{}},
			"sequence.html": &bintree{renderHtmlSequenceHtml, map[string]*bintree{}},
			"string.html": &bintree{renderHtmlStringHtml, map[string]*bintree{}},
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

