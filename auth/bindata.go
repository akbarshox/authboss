package auth

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _views_login_tpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x51\x4d\x4f\xc4\x20\x10\xbd\x9b\xf8\x1f\xc8\x78\xde\x34\xde\x69\x8f\x9e\xf6\x60\x4c\xfc\x01\xb4\xcc\x16\x12\x5a\x70\x80\xd5\xfd\xf7\x4e\x2b\xa0\x8d\x31\x91\xcb\xbc\xf9\xe0\xf1\x78\x23\x4d\x5a\xdc\x70\x7f\x27\x47\xaf\x6f\x5b\xd4\xf6\x2a\x26\xa7\x62\xec\xc1\xf9\xd9\xae\xa7\x49\x91\x06\xee\x08\x3e\xd2\x3c\x0e\x67\x3f\x9f\xec\x2a\x3b\x86\x72\xa4\xda\xb8\x78\x5a\x0a\xde\x73\xbb\x86\x9c\x44\xba\x05\xec\x21\xe1\x47\x02\xb1\xaa\x85\x71\x8e\x48\x1b\x02\x11\x9c\x9a\xd0\x78\xa7\x91\x7a\x78\x6d\x65\xc2\xb7\x6c\x09\x35\x5f\xa3\x8c\xf0\x17\x67\x60\x85\xef\x9e\x95\x15\xde\xef\xfc\xc0\xfb\xdc\xca\xff\xe4\x8d\x79\x5c\x6c\x53\xbb\x3b\x00\x07\x3f\xc4\x97\x2b\x75\xee\xaa\x5c\xe6\xc1\xf3\x3e\x58\xcd\xe8\x8a\x1b\x25\xfd\x65\xa9\x41\x17\x0e\x0a\x94\x30\x84\x97\x1e\x1e\x60\x78\xc1\xd9\xc6\x84\x24\x3b\x35\xfc\xac\x3f\x79\x9a\x7d\x12\xf5\x43\x5b\xbb\xbe\xc6\xfc\xdb\xe6\x5a\xac\xab\xec\xca\x6e\x3f\x03\x00\x00\xff\xff\x05\x40\xcd\x07\xe4\x01\x00\x00")

func views_login_tpl_bytes() ([]byte, error) {
	return bindata_read(
		_views_login_tpl,
		"views/login.tpl",
	)
}

func views_login_tpl() (*asset, error) {
	bytes, err := views_login_tpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "views/login.tpl", size: 484, mode: os.FileMode(438), modTime: time.Unix(1421030024, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	"views/login.tpl": views_login_tpl,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"views/login.tpl": &_bintree_t{views_login_tpl, map[string]*_bintree_t{
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
