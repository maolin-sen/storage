// Code generated by go-bindata. DO NOT EDIT.
// sources:
// pair.tmpl (657B)

package main

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _pairTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\x41\x8f\xda\x30\x14\x84\xef\xfe\x15\xa3\x08\x55\x50\xd1\xf8\x4e\xc5\xa1\x82\x1e\xaa\x4a\xa5\x07\xd4\xee\xf5\x91\xbc\x35\x56\x1c\xdb\x72\x9c\xa0\x28\xeb\xff\xbe\x4a\xd8\x65\xf7\xc0\x2e\xf8\xe6\x99\x79\xe3\xf7\x59\x4a\x6c\x5c\xc9\x50\x6c\x39\x50\xe4\x12\x87\x1e\xca\x5d\xee\xd0\x36\x72\xb0\x64\x64\x51\x97\xd2\x93\x0e\xcd\x77\x6c\x77\xf8\xb3\xdb\xe3\xe7\xf6\xd7\x3e\x17\x9e\x8a\x8a\x14\x63\xf2\x84\xd0\xb5\x77\x21\x62\x2e\x00\x20\x53\x3a\x1e\xdb\x43\x5e\xb8\x5a\x3e\xb4\x64\x4f\x4e\x36\xd1\x05\x52\x9c\xdd\xf0\xa5\xaf\x94\x6c\x58\xd5\x6c\xe3\x5d\x59\xb6\xa5\x77\xfa\xce\x70\x11\xb8\x64\x1b\x35\x99\x9b\xf1\xd8\x7b\x6e\x32\xb1\x10\x42\x4a\xfc\x30\x06\xd4\x91\x36\x74\x30\x2f\xc4\xb9\x28\x9c\x6d\x46\xe0\x61\xf8\x86\x40\x56\x31\x66\xd5\x12\xb3\x0e\xab\x35\xf2\x2d\x45\x42\x4a\xd3\x2b\xc3\x80\x59\x85\x27\x14\x54\xb3\xd9\x50\xc3\x48\x09\x6b\x64\x67\x3d\xa5\x6c\xaa\x60\x5b\x8e\x03\x0b\xf1\x79\xa1\x94\xf8\xaf\xe3\xf1\x6a\xe7\x49\x8f\x7b\x7a\x6f\x7a\xbc\x76\xa3\x23\xd3\x32\xa2\xc3\xce\x47\xed\x6c\x23\x1e\x5b\x5b\x7c\x58\x31\xef\xa6\xc9\x0e\x29\x2d\xf0\x75\xfa\x83\xfc\x2f\xe9\x80\x61\x22\x09\x1c\xdb\x60\xf1\xe5\xcd\x38\xeb\xe3\xf9\xcd\xfd\xea\x2a\xea\xf2\x12\xf9\x37\xee\xb2\x42\x77\x56\x92\x48\xef\xc0\x9f\x03\x00\x00\xff\xff\x66\x1e\xf8\x79\x91\x02\x00\x00")

func pairTmplBytes() ([]byte, error) {
	return bindataRead(
		_pairTmpl,
		"pair.tmpl",
	)
}

func pairTmpl() (*asset, error) {
	bytes, err := pairTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pair.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x68, 0xbd, 0x2b, 0x15, 0x76, 0xb2, 0x2d, 0x4a, 0x66, 0x3e, 0xf0, 0xe1, 0x1e, 0xb9, 0xf0, 0xa5, 0xac, 0xa5, 0x24, 0x8c, 0xce, 0xce, 0x3, 0xc2, 0x97, 0x36, 0xe0, 0x56, 0x27, 0x2e, 0x75, 0xb5}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"pair.tmpl": pairTmpl,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"pair.tmpl": &bintree{pairTmpl, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
