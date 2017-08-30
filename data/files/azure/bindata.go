// Code generated by go-bindata.
// sources:
// cloud.json
// DO NOT EDIT!

package azure

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

var _cloudJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x9b\x5d\x73\x9b\x3e\x16\xc6\xef\xf3\x29\x34\xbe\x6e\xb2\x06\xbf\xfe\x73\xe7\xfa\x85\xcd\xa6\x93\x74\x83\xd3\xce\xee\x4e\xc7\xa3\x60\xc5\xa5\x26\x92\x57\x88\xa4\x69\x27\xdf\x7d\x07\xec\x60\x1b\x21\x78\x20\x64\x7b\xd3\x74\x2c\xa1\xf3\x93\xce\x79\xce\x39\xc6\xf0\xfb\x84\x90\x16\xa7\x0f\xac\x75\x4e\x5a\xf4\x57\x24\x59\xeb\x43\xfc\x11\xe3\x8f\xad\x73\xf2\x9f\x13\x42\x08\x69\x2d\xd9\x63\xf2\x29\x21\xad\xff\xd2\xd6\x09\x21\xdf\x92\x39\x92\xad\x7c\xc1\xc3\x74\xde\xef\xe4\x5f\x42\x5a\x81\xf0\xa8\xf2\x05\x8f\xd7\xfc\xe2\xcb\x95\xcf\x7d\xba\x5b\x20\xbd\x2c\x1e\x9b\xd2\x50\x91\x5b\x77\x3f\xf4\x4b\x70\xb6\x5f\x2f\xf9\x88\xd1\x50\x45\x61\x6b\xf7\xc1\xb7\xe4\xef\xcb\x87\x37\xd9\x23\x36\x60\xd1\xc6\x4d\x5e\x88\xa7\x5c\x73\x63\xc6\x95\xa4\x41\xd9\x0e\xbd\xed\xb4\x2a\x9b\xbc\x08\x02\x9f\x0b\x3f\xcc\xb3\x7a\x25\xa4\xfa\x4e\x40\xdb\x3c\x9e\x5c\x03\x60\xce\x7e\xd2\x5c\xeb\xae\x88\x70\xeb\x61\x3c\xb9\x86\xf5\xaf\x2c\x54\xb9\x36\xf6\x1c\xc6\x29\x39\x18\x4f\x2c\x54\x35\x28\xc6\x34\xf0\xef\x85\x34\xc4\x5a\x02\x00\x18\xae\xbc\xef\xe3\xf8\xd5\x0c\x96\x04\xf7\xd6\x64\x85\xe0\xfe\x67\xc4\xee\x98\x47\xc6\xbe\x7a\xce\x8d\x71\xca\xe9\x92\x92\x58\x59\xc5\x41\x9e\xcc\x8b\xa5\x55\x21\xc8\x84\x14\x5c\x89\x02\xb3\x3b\x07\x03\x96\x77\xfe\xc5\x8d\xbb\x54\x90\xcf\x34\x0a\x04\x71\x15\x55\x2c\x0f\xe2\xa3\xa4\xbf\xfc\x80\x24\x21\x5f\x88\x70\x97\x4c\x4c\xa2\xbd\x82\xc6\x25\x0b\x28\x5f\x9a\x25\x3e\x8d\xa4\xd8\xb0\x72\x79\xb3\xed\x3c\xd8\xf0\x15\x53\xdf\x99\x8c\x6d\xe7\x2a\x3c\x89\x33\xc0\x76\x1c\x69\x55\x4d\x8f\xa9\x5c\xfa\xf7\xf7\x79\x66\x6f\x2f\x49\x6c\xb9\xd0\x64\xb4\x8e\x8d\xe2\xe6\x3e\x09\xbe\x14\xdc\x60\xad\xdc\xad\xd1\xba\xa2\x4b\x5d\x9f\xaf\xe8\x46\xc8\xdc\x68\x4a\xec\xc5\x02\x21\xa3\xf0\x30\xa5\x98\xf2\x66\x3c\x95\xc6\x33\x61\xf3\x7f\x17\x7c\x45\x2e\x05\x5f\x19\x6b\x63\xa9\xe5\xea\x46\xaf\xd8\xd3\xf6\x28\xc9\x57\x1a\xb0\xdc\x88\x1a\x45\x61\xac\x4d\x1f\x48\x23\xf4\x75\x6a\xb5\x4c\xf2\xc5\xf7\x94\x90\xf9\x89\x7a\x6f\x3d\xf5\x00\x86\x90\x7a\xa1\x4a\x46\x5b\x3f\x8b\x0f\xc4\xa5\xbe\xa2\x0f\xb9\x34\xff\xa0\x1b\xca\xcb\xcf\xe1\x47\x3c\xad\x9a\xed\xeb\x90\xae\x0b\x4c\x96\x8a\x2b\x31\x99\xa3\xaf\xb4\x1d\xf4\x79\xa8\x28\xf7\xd8\xfc\x79\xc3\x72\x9a\xc2\x70\x1d\x25\x61\xae\x28\x5f\x52\xb9\x5c\x8c\xda\x7b\x73\x4b\x16\x7a\xd2\xdf\xa4\x3a\xc9\x9b\xe3\x51\xc5\x56\x42\x3e\x27\x3e\x3b\x0d\x99\xf4\x0f\xa3\xc9\xdb\xc4\xab\x5b\xfb\xed\xd1\x87\xd6\x39\x69\x9f\x0d\xfa\xc3\xbd\x15\x3f\x5c\xb7\xce\x89\xdd\xce\x3d\x2b\x8d\xcf\x02\xf8\xac\xb7\xf1\x59\x67\x83\x5e\x16\x6f\x00\xe2\xd9\x00\x9e\x5d\x0d\xcf\xce\xe0\x75\xce\x34\x3a\xab\xd3\xc3\xf0\x3a\x00\x5e\xa7\x1a\x5e\x37\x83\x37\xd0\x3c\x3b\x04\xe1\xba\x00\x5c\xb7\x1a\xdc\x30\xeb\xda\x6e\x96\xae\xdf\x06\xe9\x7a\x00\x5d\xef\x6d\x9e\xd5\xe9\x60\xc7\xf6\x01\xba\xfe\xdb\x1c\x6b\xeb\x9a\x45\x3d\x3b\x00\xe8\x06\x6f\xf3\x6c\xaf\x5f\xdb\xb3\x43\x80\x6e\x58\x4c\x47\x4e\x89\x27\x1e\x36\x91\x62\xa7\x3e\x57\x8c\x87\xfe\x23\x23\xaf\xa9\xb7\x06\x7b\x67\x68\x63\xec\x7f\x01\xec\x7f\x35\xc9\x6e\xf5\xb3\x41\x6b\xd9\xb5\xe9\x2d\xa4\xda\x58\x25\xe5\xe6\x8f\x9d\xbd\x05\xd5\xa2\x92\x62\xf4\xe7\x4e\x7f\x02\xe0\x4f\x0c\xf4\x13\xb0\x94\xe6\xd4\xaa\x1e\x56\x49\x27\x40\x25\x9d\x18\x2a\xa9\x89\x2e\x9b\x6f\xb5\x52\x65\xb5\x41\x38\xa0\x8e\x4e\x0c\x75\xd4\x04\x97\x4d\xb7\x7a\x31\xb0\x51\x3a\xa0\x90\x4e\x0c\x85\xd4\x44\x97\x95\x8d\x5e\x0c\xba\x28\x1d\x22\x9b\x89\x49\x36\xa8\x6b\x73\x4a\x29\xcc\x87\x44\x9e\x55\x31\xf4\x80\x62\x0a\xf3\x21\xc1\x67\x55\x8c\xbe\xf2\xb4\x88\xfb\x17\x09\x3f\xab\x62\xfc\x21\x89\x6f\x08\x13\x2e\x1e\x21\x1f\x1f\x4d\x3b\xa2\x7c\xb4\xff\x0f\x09\x10\xa3\xb4\x6b\x50\x36\x98\x08\x31\xc8\x4e\x0d\xc8\x26\x13\x22\x46\xd9\xad\x41\xd9\x60\x62\xec\x61\x94\xbd\x3a\x61\x99\x95\x8f\xae\x6f\x5c\x3d\xa8\x7c\xea\xe8\xa7\xd1\x3c\x0e\x72\xd6\x51\x50\xa3\xf9\x1c\xe4\xac\x23\xa2\x46\xf3\x3a\xc8\x59\x47\x46\xcd\xe6\x77\x50\x49\x56\x1d\x29\xd9\x6d\x2d\x44\xdb\x39\x31\x0a\xa2\xba\x48\x33\xe4\x9a\x9a\x21\xb7\x7e\x15\x1a\x80\x7c\xc8\x39\xba\xa6\x43\x34\xf1\x95\xd7\x9f\x2e\x88\x87\xf4\x42\xae\xa9\x17\x32\xe1\x01\x95\x67\x08\xf2\x21\xbd\x90\x6b\xea\x85\x4c\x7c\xe5\x35\xa7\xd7\x47\xc3\x0f\x8b\xbf\xaa\x01\x58\x9e\xc6\xe1\x13\x84\xda\x71\xd7\xd8\x8f\xa3\x3e\x7e\xcb\x19\x42\x41\x68\xec\xc8\x51\x2f\xeb\xa9\xdb\xb2\xc0\x6f\xfa\x2e\xd4\x93\xbb\xc6\xa6\xdc\x98\x67\x80\xac\x6d\xdb\xa8\x96\xc1\xbe\xc2\x2d\xe8\x2b\xdc\xb7\x34\xe6\x70\x4a\x04\x39\x0b\xfa\x0a\x33\x67\x83\xa9\x11\xc4\x2c\x68\x2b\xcc\x98\x4d\xa6\x48\x90\xb3\xa0\xad\x30\x73\x36\x98\x2a\xc1\xa6\xc2\x2d\x68\x2a\x0a\xc2\xb3\xbc\x41\xaf\xa0\x76\x58\x49\xb5\xa4\xd4\x68\x72\x47\x49\x6b\x89\xa9\xd1\x24\x8f\x92\xd6\xd2\x53\xb3\xc9\x1e\x45\xad\x25\xa9\x86\x93\x3e\x2a\xab\xa2\x66\xbd\x20\x56\x81\x6e\xdd\x1e\x62\xcd\xfa\x0c\xe8\x95\x66\x86\x4e\x69\x06\x96\x25\xed\x24\x2d\x2c\x40\x67\xc0\x19\xce\x0c\xc7\x67\x62\xcb\xea\x5c\x93\x79\x07\x8b\xc8\x19\xd0\x7c\xcc\x0c\xad\x87\x89\x2d\xab\x6c\x4d\xd8\x7d\x2c\x00\x67\xc0\x8f\x7f\x33\xc3\x8f\x7f\x26\x36\xed\x47\x67\x5d\xcb\x60\x82\x9c\x59\xc0\xef\xba\x47\x93\xa0\x98\xcb\x0a\xb8\xa3\xeb\x17\xcc\x8b\x33\x2b\x44\x00\x43\x03\x60\x58\x57\x15\xa0\x73\x6d\x84\xce\xae\x4a\x57\xaa\x0b\xd0\xb9\x5d\x84\xae\x5b\x95\xae\x54\x19\x68\x46\x19\x22\x78\xc3\xaa\x78\xe5\xe2\x40\xb3\x8a\xd5\x87\x62\xaf\x5f\x39\xf8\xca\xe5\x01\x26\x17\x07\x28\x18\x8e\xa1\x60\x38\x60\xf0\xe9\x2d\x4d\x67\x08\xd2\x01\x25\xc3\x31\x94\x0c\x13\x5d\x36\xf8\xf4\x36\x66\xd0\xc7\xd4\xe1\x00\x5f\xaa\x1d\xc3\x57\x6a\x13\x9d\x16\x7b\x7a\xe3\x62\xf5\x3a\x98\x3c\x1c\xa0\xa8\x39\x86\xa2\x66\xe2\xd3\x22\xcf\xb6\xf5\x9a\xdb\x1e\x60\xfa\x70\x80\x07\x96\x1c\xc3\x03\x4b\x26\xc0\x8e\x96\xfa\xba\x7a\xe1\xb5\xba\x60\xfc\x21\xf7\x3e\x1d\xd3\xbd\x4f\x07\xbd\xf5\x54\xbb\xe7\x77\x90\x7b\x9f\x8e\xe9\xde\xa7\x91\xaf\x5c\x22\x3d\xb0\x3b\x70\x90\xbb\x9f\x8e\xe9\xee\xa7\x11\x10\x51\x49\x1b\x3e\x43\x44\x26\xa6\x1b\xa0\x46\x44\x44\x28\x36\xfa\x85\xc9\x71\x11\xa5\xb8\x26\xa9\x98\x18\x11\xad\x74\xd3\xef\x49\xe9\xf3\xba\x9e\x64\x4b\xc6\x95\x4f\x83\x9c\xa7\x75\x37\x52\x3c\xfa\x4b\x26\x63\xc3\xa3\xf4\xb5\xb0\xd7\x15\x37\x01\x7d\x9e\x09\xf9\x40\x55\x3c\x7e\xef\xb3\xe0\xe0\x3d\x05\xca\xb9\x50\xc9\x63\xc7\xf1\xba\xaf\x2b\xc6\x6b\x7e\xa7\xf2\x81\xc9\x33\xba\xd9\x84\x9e\x58\xb2\x33\x4f\x3c\xfc\xcd\x0b\xa2\x50\x31\x79\xba\xa7\x89\x97\x4c\x57\x33\x5d\xb6\xe4\x61\xf6\x92\xdd\x15\x2f\x29\x48\xc2\x75\xfc\x14\xf3\x9e\x66\xfb\x8e\x9b\x27\xf8\xbd\xbf\x4a\x36\xf9\xef\xdb\x9b\xe9\x62\x3e\xbd\x1a\x5d\xcd\x17\x17\x93\x03\x80\x78\x25\x21\xe3\x73\xdd\xbe\x20\xb7\x50\x8c\x53\xae\x16\xfe\xf2\x78\xd2\x8f\x70\xeb\xc9\xed\x70\x76\x89\x80\xde\xb1\x84\x73\x9e\x0c\x93\x8b\xcc\xd5\x3e\xdf\x44\x6a\x7b\xf9\xcf\xf4\x11\xeb\x83\xdd\x94\xb3\xbb\xb7\x1f\xdd\xf1\xcd\xc5\xe7\xf9\xc5\xf5\x55\xc9\x0e\xc2\xe8\x2e\x0d\x3e\xe3\x3e\x0e\x27\x19\x77\xe3\x1e\x4c\x7a\x87\x3d\x8d\x3f\x5d\x4c\x4b\xfd\xe1\x05\x3e\x2b\xf0\xc7\x76\xd8\xb8\x83\x71\x32\xfc\x7e\xec\xee\x74\x7c\x33\x9d\x03\xfc\x21\xf3\x24\x53\x45\x7b\x70\x73\x66\x64\xf7\x91\x37\x27\xdd\xcb\x86\x86\xe1\x93\x90\xcb\x83\xfd\xec\xfe\x97\xff\xfe\x80\x96\x05\x5c\x25\x24\x5d\xbd\x5b\x32\x08\xb7\xcb\xbf\x87\xb2\xdd\xf9\xf5\xcd\xc8\x99\x2e\x46\xe3\xf1\xf5\xed\x55\xa1\x3f\x76\x14\x0b\xea\x79\x22\xe2\x06\x8f\xe4\x0e\xa6\xce\x48\x0e\x8b\xec\x4e\x8b\x8c\xf2\xe6\x36\x20\xf8\xdd\x96\x2e\xa7\xff\x42\xb6\xb3\x66\xcf\xf9\x5b\xd1\x06\x8a\xb7\x41\x2e\xb3\xf3\xf1\xf8\x4a\x8b\xcf\x3a\xba\x63\x92\x33\xc5\xf6\x61\xd1\x7a\x64\x32\x3c\x7a\xa1\xf8\xf0\x08\x5e\x47\x63\x33\xd6\xd9\xe0\xac\x7b\x58\x22\x32\xb5\x54\x1b\xdf\xbe\xcd\x7c\x74\x9e\x4b\x16\x7f\xa4\x64\xc4\x34\xd6\xed\xdf\x98\xf8\xe5\xe4\xe5\x7f\x01\x00\x00\xff\xff\x0a\xe3\x4a\xda\x21\x3d\x00\x00")

func cloudJsonBytes() ([]byte, error) {
	return bindataRead(
		_cloudJson,
		"cloud.json",
	)
}

func cloudJson() (*asset, error) {
	bytes, err := cloudJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cloud.json", size: 15649, mode: os.FileMode(420), modTime: time.Unix(1453795200, 0)}
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
	"cloud.json": cloudJson,
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
	"cloud.json": {cloudJson, map[string]*bintree{}},
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