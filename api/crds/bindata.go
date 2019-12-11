// Package crds Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// meta.appscode.com_resourceclasses.yaml
// meta.appscode.com_resourcedescriptors.yaml
package crds

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

// ModTime return file modify time
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

var _metaAppscodeCom_resourceclassesYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x59\xdf\x6f\xdb\xbe\x11\x7f\xf7\x5f\x71\xe8\x1e\xda\x02\xb1\x8c\xb4\x28\xb6\xf9\xad\x70\xdb\x21\x5b\xdb\x15\x49\xdb\xf7\x33\x79\x96\xb8\x50\x24\xc7\x1f\x6e\xd2\x61\xff\xfb\x70\x94\x64\x5b\xb1\x24\xa7\xce\x8a\x2f\xdf\x7c\x47\x1e\xef\xf7\x7d\x44\xcf\xe6\xf3\xf9\x0c\x9d\xfa\x4e\x3e\x28\x6b\x96\x80\x4e\xd1\x5d\x24\xc3\xbf\x42\x71\xfb\x97\x50\x28\xbb\xd8\x5e\xae\x29\xe2\xe5\xec\x56\x19\xb9\x84\x55\x0a\xd1\xd6\xd7\x14\x6c\xf2\x82\xde\xd1\x46\x19\x15\x95\x35\xb3\x9a\x22\x4a\x8c\xb8\x9c\x01\x08\x4f\xc8\xc4\xaf\xaa\xa6\x10\xb1\x76\x4b\x30\x49\xeb\x19\x80\xc1\x9a\x96\xe0\xdb\xe3\x42\x63\x08\x14\x0a\x3e\x5b\xa0\x73\x41\x58\x49\x85\xb0\xf5\x2c\x38\x12\x2c\x09\xa5\xcc\xe2\x51\x7f\xf1\xca\x44\xf2\x2b\xab\x53\x6d\x02\xf3\xe6\xf0\xf7\x9b\x7f\x7e\xfe\x82\xb1\x5a\x42\x11\x22\xc6\x14\x0a\x57\x61\xa0\x19\x40\x77\xd3\x4d\x26\x67\x42\xbc\x77\xb4\x84\x10\xbd\x32\xe5\xc3\xd3\x9d\xf2\xc5\x91\xe6\x07\xb2\xde\x96\x74\x20\x48\x62\xe4\x9f\xa5\xb7\xc9\x2d\xe1\xd8\x82\xe6\x54\x56\x14\xa0\xf1\x5d\xe7\xb5\x15\x9b\x9d\xe9\x5a\x85\xf8\x8f\x63\xde\x47\x15\x62\xe6\x3b\x9d\x3c\xea\x23\x87\x65\x5e\x50\xa6\x4c\x1a\xfd\x03\xee\x0c\xc0\x79\x0a\xe4\xb7\xf4\xcd\xdc\x1a\xfb\xc3\x7c\x50\xa4\x65\x58\xc2\x06\x75\xf6\x4d\x10\x96\x0d\xf8\xcc\xea\x39\x14\x24\x99\x96\xd6\x9d\x94\x56\xe5\xc6\xa1\x4b\xf8\xcf\x7f\x67\x00\x5b\xd4\x4a\x66\xc7\x34\x4c\xeb\xc8\xbc\xfd\x72\xf5\xfd\xf5\x8d\xa8\xa8\xc6\x86\xc8\x17\x5b\x47\x3e\xaa\x4e\x06\xaf\x83\xfc\xda\xd1\x00\x24\x05\xe1\x95\xcb\x12\xe1\x39\x8b\x6a\xf6\x80\xe4\x8c\xa2\x00\xb1\x22\xd8\x36\x34\x92\x10\xf2\x35\x60\x37\x10\x2b\x15\xc0\x53\x36\xd1\xc4\xac\xd2\x81\x58\xe0\x2d\x68\xc0\xae\xff\x45\x22\x16\x70\xc3\x6e\xf0\x01\x42\x65\x93\x96\x20\xac\xd9\x92\x8f\xe0\x49\xd8\xd2\xa8\x9f\x3b\xc9\x01\xa2\xcd\x57\x6a\x8c\xd4\xfa\xbe\x5b\x39\xed\x0c\x6a\x76\x42\xa2\x0b\x40\x23\xa1\xc6\x7b\xf0\xc4\x77\x40\x32\x07\xd2\xf2\x96\x50\xc0\x27\xeb\x09\x94\xd9\xd8\x25\x54\x31\xba\xb0\x5c\x2c\x4a\x15\xbb\x8a\x12\xb6\xae\x93\x51\xf1\x7e\x21\xac\x89\x5e\xad\x53\xb4\x3e\x2c\x24\x6d\x49\x2f\x82\x2a\xe7\xe8\x45\xa5\x22\x89\x98\x3c\x2d\xd0\xa9\x79\x56\xdc\xc4\x5c\x96\xb5\xfc\xd3\x2e\x54\xcf\x0f\x34\x7d\x90\xdf\xcd\xca\x89\x37\xea\x77\x4e\x3d\x50\x01\xb0\x3d\xd6\xe8\xbf\x77\x2f\x93\xd8\x2b\xd7\xef\x6f\xbe\xee\xb2\x2c\x87\xa0\xef\xf3\xec\xed\xfd\xb1\xb0\x77\x3c\x3b\x4a\x99\x0d\xf9\x26\x70\x1b\x6f\xeb\x2c\x91\x8c\x74\x56\x99\x98\x7f\x08\xad\xc8\xf4\x9d\x1e\xd2\xba\x56\x91\x23\xfd\xef\x44\x21\x72\x7c\x0a\x58\xa1\x31\x36\xc2\x9a\x20\x39\x2e\x3f\x59\xc0\x95\x81\x15\xd6\xa4\x57\x18\xe8\xb7\xbb\x9d\x3d\x1c\xe6\xec\xd2\xd3\x8e\x3f\x6c\x87\xfd\x8d\x8d\xb7\x76\xe4\xae\xd3\x75\x6b\xa8\x86\xda\x3a\xfa\x5b\xee\x35\x3d\xea\xc8\xed\xbc\x88\x6d\x7c\x28\x05\x40\x45\xaa\x8f\x88\xe3\xd7\xb6\x87\x84\x35\x83\x8c\x51\x79\xcd\xea\xa5\xdb\x55\x8d\x25\xdd\x38\x12\x5c\x87\x11\x95\x09\x39\x56\xbe\xce\x55\x0c\xb8\xb6\x29\x72\xf1\x2a\xde\x07\x29\xe4\xc6\x34\xbc\x30\xe4\x8d\xc2\x9a\x62\x64\xcf\xb4\x3d\xcd\x0a\xea\x27\x8d\x73\x1f\x68\xff\xc2\xba\x66\x10\xbd\x84\xaf\x15\xe5\xb3\x4d\x3f\xa2\x56\x61\x65\xc0\xa9\x3b\xd2\x61\x42\x22\xc0\x0b\x2a\xca\xe2\x02\x5e\xbd\xb9\x7b\xf5\xe6\xe5\x98\xf2\x30\x15\xd7\x9e\x05\x5e\x3c\xda\x80\xac\x75\x53\xc1\x1b\xeb\x5b\xa5\x77\x35\x4b\x92\x9d\x4a\x2a\x56\xe4\x27\x0d\x40\x8e\x54\xb0\x3a\x45\x82\x6f\xd7\x1f\xbb\xc6\xd9\x88\xb3\x1e\x10\xde\x61\xc4\xcc\x6a\xc3\x3c\x65\x40\xb6\xb4\x3b\x5e\xec\x8e\x06\x40\x4f\xed\x2c\x90\xec\xd9\xeb\x0f\x2b\x78\xf5\xfa\xaf\x7f\x7e\xb2\xc7\xf2\xa6\x33\x63\x5e\x2b\x43\x59\x40\x3f\xf0\x4d\x48\x27\x4d\x7c\x96\x77\x2e\x9c\x29\x9f\x3d\x35\xe8\xdc\x11\x95\x27\x39\x66\xc4\x9c\x93\x62\x84\x37\xd8\x80\x8e\x37\xa0\xf7\x78\x3f\xc0\xcf\x30\x68\xe2\xdc\xa8\xde\x9d\x1b\xa7\x0e\xaf\xad\xd5\x84\x66\x60\x87\x63\x8c\x76\xce\xb5\xe3\xb1\xee\x45\x39\xb7\xd5\x16\x7f\x74\x30\x0c\x92\xc1\x7a\xad\xca\x64\x53\xd0\xf7\xa0\x24\x8f\x81\x8d\x22\x9e\x95\xdd\x20\x2c\x00\xae\x86\xdd\x08\x20\x2d\x05\xf3\x9c\xdb\x99\x35\xf7\x75\x2b\xc4\x08\x9d\x24\xf5\xee\xe3\xf2\xc1\xad\x55\x12\x30\x45\xcb\x9d\x70\x2c\x76\xc2\x92\x17\xca\x9a\x7c\xeb\x4e\x7e\x0a\x04\x38\x22\x51\x64\xa8\x0e\x35\xfa\x50\xa1\xd6\x63\x29\x75\xba\x53\x96\x43\x73\xe7\xd1\x51\x80\x9c\xb4\x8d\xcb\x9e\x24\x64\x7b\x0c\x23\x7f\x51\xc6\x74\xed\xcc\x1b\x43\x47\x78\x9d\x09\x23\xec\x56\xb9\x89\x34\x1d\x2d\xbb\x1f\xa4\xca\x2a\x4e\x65\x38\x43\xd0\x72\xa0\x2f\x8f\xdb\x33\xcf\xc5\x3a\x40\xee\x4a\x71\x80\xc5\x77\x0d\x90\x1b\xf5\x8e\x18\x13\x46\x8d\x75\x91\x41\x14\xd1\x87\x08\xbc\x23\x23\x52\xb3\xd3\x34\x7f\x24\x71\xc7\xcd\xc7\xf3\xec\xe2\x19\xe4\x9c\x56\x22\x23\x87\xe3\x86\xca\x62\x7a\xd8\xa2\x2d\xbd\xe6\xb3\xa2\x89\xe3\x45\x1e\xe1\x2d\x9c\x57\x75\xd3\xd8\x1f\x8a\x1a\x01\x37\x4f\x84\x35\x80\x43\x18\x61\x14\xd0\x4c\x17\xe8\x38\x88\x39\x13\xbe\x4c\xcf\xb3\x29\xe0\x72\xa2\xfe\x46\xc1\xca\x99\x30\xe5\x10\x8a\x8c\x68\xfb\x68\x80\x32\x08\x42\x46\x84\xfe\x12\x34\xf9\xff\x8c\xa8\x5f\x01\x22\x8f\x00\x1b\x93\x4a\x4d\x35\x95\x21\x68\x71\x46\x23\xd0\xca\xdc\x4e\x37\x82\x8f\xbc\x23\x7b\x1a\x77\xf5\xbf\xdb\xb0\xa5\x06\x23\x72\x5b\x34\x92\x24\x87\x78\x3d\xf2\xc5\x10\x2d\x84\xe4\x37\x28\xe8\xe0\x2d\x09\xa4\x15\xa9\xee\xde\x0f\x2e\x40\x62\xa8\xd6\x16\xbd\x0c\x17\x40\x51\x9c\xd3\x05\x58\xe1\x13\x0d\xa0\x4d\xc6\x4e\x9d\xbe\x0e\x03\x71\x3a\xa1\xd5\xa9\xc6\x70\xa8\xde\xe9\x0c\x7b\xb7\xff\xc1\xfd\xb7\x4a\x35\x1a\xf0\x84\x12\xd7\x9a\xb2\x61\x64\x22\xd0\x9d\xd3\x93\x88\x9e\x93\xd1\x25\xef\x6c\xd8\xe5\x26\x07\xfb\xac\xda\x48\x7e\x04\x30\xf6\x14\xff\xe6\x35\x0b\x52\x02\xb5\xbe\x87\xfc\xa8\x10\x00\x23\x20\xfc\xa0\x75\x50\x31\xc7\xdd\x53\x08\x67\xe8\x70\x46\x6a\xd7\xa8\x72\x0e\x90\x9f\x4e\xf0\x4f\xfb\x7d\x63\xf3\xee\x40\x54\xe7\xcb\xc9\xa9\x97\x3b\xc3\xa1\x58\xd3\x3c\xba\x64\x79\x1d\x23\x8b\xd9\x4f\x41\x10\x56\xd2\x45\xd3\x3f\xc2\x71\x12\xf2\x70\x74\x28\x6e\xb9\xbb\x70\x37\x3e\xa1\xc3\x63\x0a\x65\xc5\x35\x22\x62\x6e\xb2\x27\x06\xa6\x91\x6a\xab\x64\x1a\x00\x2b\xc0\x8d\xdc\xfa\x12\x8d\xfa\x39\xe2\x8d\xe9\xe2\xa0\x1a\xd5\x63\xb2\xeb\x3d\xef\xe3\x00\xe5\x47\xab\xfc\xe3\xfc\x74\x9a\xfa\x82\xea\xdd\xfa\x19\x6b\xea\x2e\x3d\xec\x7a\x7c\xfc\x37\x97\x92\xc8\x0f\xa5\xfb\x82\x5a\xd3\x1f\x55\x49\xc3\xb0\x78\x0c\x10\x0f\x4d\xad\xf9\xee\xd1\xac\x47\x6c\x1f\xc6\x7a\xb4\x23\x90\x3b\xa0\xf0\x03\x52\xf7\x39\x02\xdb\x4b\xd4\xae\xc2\xcb\x3d\xad\xfd\x7b\xa2\xf9\xdb\xe0\x80\x0d\x90\xdf\xe5\xe5\x12\xa2\x4f\x0d\x66\x09\xd1\x7a\x2c\xa9\xa5\xfc\x2f\x00\x00\xff\xff\x33\x95\xf7\x69\x92\x19\x00\x00")

func metaAppscodeCom_resourceclassesYamlBytes() ([]byte, error) {
	return bindataRead(
		_metaAppscodeCom_resourceclassesYaml,
		"meta.appscode.com_resourceclasses.yaml",
	)
}

func metaAppscodeCom_resourceclassesYaml() (*asset, error) {
	bytes, err := metaAppscodeCom_resourceclassesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "meta.appscode.com_resourceclasses.yaml", size: 6546, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metaAppscodeCom_resourcedescriptorsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x5c\x5f\x73\xdb\x36\x12\x7f\xf7\xa7\xd8\xf1\x3d\xa4\x9d\x91\xe8\x8b\x3b\x9d\xbb\xd3\x9b\xc7\x49\x6f\x7c\x97\xa6\x99\xd8\xcd\xcb\xb5\x0f\x10\xb9\x12\x51\x93\x00\x0b\x80\xb2\xd5\x4e\xbf\xfb\xcd\x2e\x08\x8a\x92\x08\x92\x52\x92\x49\x8b\x27\x0b\x04\x16\x8b\xc5\xfe\xf9\x01\x0b\xf8\x62\x3e\x9f\x5f\x88\x4a\x7e\x40\x63\xa5\x56\x0b\x10\x95\xc4\x67\x87\x8a\x7e\xd9\xe4\xf1\x9f\x36\x91\xfa\x6a\xf3\x72\x89\x4e\xbc\xbc\x78\x94\x2a\x5b\xc0\x6d\x6d\x9d\x2e\xdf\xa3\xd5\xb5\x49\xf1\x15\xae\xa4\x92\x4e\x6a\x75\x51\xa2\x13\x99\x70\x62\x71\x01\x90\x1a\x14\x54\xf9\x20\x4b\xb4\x4e\x94\xd5\x02\x54\x5d\x14\x17\x00\x4a\x94\xb8\x00\xd3\x74\xcf\xd0\xa6\x46\x56\x4e\x1b\x9b\x50\xff\x44\x54\x95\x4d\x75\x86\x49\xaa\xcb\x0b\x5b\x61\x4a\xd4\x44\x96\xf1\x10\xa2\x78\x67\xa4\x72\x68\x6e\x75\x51\x97\xca\xd2\xb7\x39\xfc\xe7\xfe\x87\xb7\xef\x84\xcb\x17\x90\x58\x27\x5c\x6d\x93\x2a\x17\x16\x2f\x00\xc2\x68\xf7\x5c\xcd\x15\x6e\x5b\xe1\x02\xac\x33\x52\xad\x0f\x7b\x87\x09\x24\x47\xdc\x77\x68\xdd\xac\xb1\x43\x28\x13\x8e\x7e\xae\x8d\xae\xab\x05\x1c\xcf\xc0\xf7\x62\x46\x01\xbc\xfc\x76\x92\x0b\x53\xe7\x8f\x85\xb4\xee\xbf\x91\x06\x6f\xa4\x75\xdc\xa8\x2a\x6a\x23\x8a\x5e\xf1\xf1\x77\x9b\x6b\xe3\xde\xee\x46\x9c\x83\xc9\xfc\x07\xa9\xd6\x75\x21\x4c\x5f\xd7\x0b\x80\xca\xa0\x45\xb3\xc1\x1f\xd5\xa3\xd2\x4f\xea\x3b\x89\x45\x66\x17\xb0\x12\x05\xcb\xd1\xa6\x9a\x26\xcb\x84\x2b\x91\x22\xd1\xb4\xf5\x32\x90\x6a\x06\xf3\xc2\x5f\xc0\xef\x7f\x5c\x00\x6c\x44\x21\x33\x16\xa2\xff\xa8\x2b\x54\x37\xef\xee\x3e\x7c\x73\x9f\xe6\x58\x0a\x5f\x49\x03\xeb\x0a\x8d\x93\x81\x06\x95\x8e\x3e\xb6\x75\x00\x81\x5f\x56\xd3\x17\x44\xca\xb7\x81\x8c\x34\x10\x2d\xb8\x1c\x61\xe3\xeb\x30\x03\xcb\xc3\x80\x5e\x81\xcb\xa5\x05\x83\x3c\x45\xe5\x98\xa5\x0e\x59\xa0\x26\x42\x81\x5e\xfe\x82\xa9\x4b\xe0\x9e\xc4\x60\x2c\x49\xb2\x2e\x32\x48\xb5\xda\xa0\x71\x60\x30\xd5\x6b\x25\x7f\x6b\x29\x5b\x70\x9a\x87\x2c\x84\xc3\x66\x75\x42\x61\x15\x55\xa2\x20\x21\xd4\x38\x03\xa1\x32\x28\xc5\x16\x0c\xd2\x18\x50\xab\x0e\x35\x6e\x62\x13\xf8\x5e\x1b\x04\xa9\x56\x7a\x01\xb9\x73\x95\x5d\x5c\x5d\xad\xa5\x0b\x16\x98\xea\xb2\xac\x95\x74\xdb\xab\x54\x2b\x67\xe4\xb2\xa6\x25\xbf\xca\x70\x83\xc5\x95\x95\xeb\xb9\x30\x69\x2e\x1d\xa6\xae\x36\x78\x25\x2a\x39\x67\xc6\x95\x63\x33\x2e\xb3\xbf\xb5\x4b\xf5\xa2\xc3\xe9\x81\x2d\xf8\xc2\x4a\x1a\x95\x3b\x69\x28\x48\x0b\xa2\xe9\xe6\xf9\xdf\x89\x97\xaa\x48\x2a\xef\x5f\xdf\x3f\xb4\xaa\xc6\x4b\xb0\x2f\x73\x96\xf6\xae\x9b\xdd\x09\x9e\x04\x25\xd5\x0a\x8d\x5f\xb8\x95\xd1\x25\x53\x44\x95\x55\x5a\x2a\xc7\x3f\xd2\x42\xa2\xda\x17\xba\xad\x97\xa5\x74\xb4\xd2\xbf\xd6\x68\x1d\xad\x4f\x02\xb7\x42\x29\xed\x60\x89\x50\x57\x64\xaa\x59\x02\x77\x0a\x6e\x45\x89\xc5\xad\xb0\xf8\xd9\xc5\x4e\x12\xb6\x73\x12\xe9\xb8\xe0\xbb\xee\x73\xbf\xa1\x97\x56\x5b\x1d\xbc\x62\x28\x7d\x36\x44\x25\xdd\x39\xc9\x3d\xed\x74\x58\x1e\x55\x1e\xac\x73\x70\x41\xde\xcf\xee\x7c\x3c\x8f\x2d\x57\x12\x49\x03\x3c\x7d\x58\x69\x03\xec\x3c\xcc\x11\x4d\xf2\x3b\x19\x42\x45\x4e\x5b\xaa\x75\x72\xd4\x20\xc6\xfa\x31\x4b\x3d\x9f\x0f\x78\xee\xfc\xf0\x1a\x9a\xd7\xa5\x50\x60\x50\x64\x62\x59\xe0\xde\xf7\xc6\x2d\xf4\x12\x0d\x82\x3b\xe6\x16\x62\x4b\xb7\x2b\x2b\x6d\x4a\xe1\x26\x70\xeb\x1b\x32\xa3\x0a\x74\xe5\xa3\x1b\xfc\xe0\xbd\x24\x8f\xe2\x1d\x9b\x17\xfb\x4a\xf7\x09\x97\xf9\x21\x23\x69\x18\x86\x87\x1c\xe1\x05\x05\x9c\x17\x5d\xfa\x55\x55\x48\xcc\x82\xbf\xaa\x8c\x2c\x85\xd9\x82\xcc\x48\x4f\x57\xb2\x77\xd9\x76\x52\xa0\x6e\xc2\x5a\x69\x1d\x48\xd5\xd8\x9d\x0d\x9d\xb7\x64\xee\x4d\x3b\xe9\x3d\x70\xb0\xf9\x08\x51\x62\x8e\x6c\x1d\xbb\xe6\x96\xd7\x4b\x0a\x96\x57\x3f\xdc\xdc\x5d\x35\x12\x98\xdf\x7b\x45\x4b\xd9\x5f\x5f\x2d\x0b\xbd\xbc\x2a\x85\x75\x68\xae\x1a\x1f\x6f\xaf\xae\x93\xbf\x93\x8d\x91\xcd\xcc\xd9\xd0\x22\x63\x92\x82\x96\xda\xe0\x59\x0b\xfa\x8b\xd5\x8a\xe1\xc1\xf8\x92\x06\x24\xd1\xf8\x47\x59\x56\x05\x72\x25\x54\xc2\xe5\x33\x90\x09\x26\xf0\x24\x5d\x0e\xc2\x18\xb1\x05\xa5\x7d\x34\x3a\x8b\x2f\x06\x23\xe3\x3c\x51\xb3\x5e\x6b\xe0\x0f\x24\x19\x76\xa7\xe7\x2b\x7c\x65\xa4\x36\xd2\x6d\x27\xf0\x12\x9a\x36\x4a\x4f\x51\x72\x8d\xa6\x51\xf3\x26\x6e\x18\x2c\x84\x93\x1b\x04\x59\x56\xda\x38\xa1\xa2\x9a\x14\x22\x7b\xa3\x7e\xa9\x2e\x2b\x61\xbc\x9a\x6b\x97\xa3\xb1\x09\xbc\xd1\x4f\x68\x40\xd5\xe5\x92\x02\x8b\x30\x34\x51\x45\x1e\xc9\x60\x16\xa1\x9a\xcb\x75\x8e\xa6\x65\x35\x81\x06\x69\x82\xcb\x85\xe3\xd8\xb4\x44\xd0\xa5\x74\x0e\x33\x32\x88\x42\x96\x92\xfe\x64\x58\x14\xa1\x69\x53\x54\xc2\x48\xdd\x82\x8a\x25\xc2\x5a\x6e\x50\xd1\xaa\x1c\x8c\xd7\x4b\xa2\x71\x2a\x24\xb1\x6f\xae\x07\x16\xa9\x91\x68\x4f\x0b\xfe\x3e\xbe\x42\xec\x76\xfc\xea\x0c\x78\xa2\x3d\x9f\x13\x99\xf3\x5f\xc6\xc2\x09\x2e\x48\x83\xd9\xb1\x74\xe6\xad\xed\xf7\x7c\x22\xf3\xe9\xa9\x0e\xeb\xd8\xf3\x89\xb8\x38\xaa\xee\x8d\xef\xdd\x4f\xec\x29\x0e\xa2\xba\x52\x98\x32\xc2\x98\x18\xd9\x87\x83\x6c\x41\x60\xa6\x5f\x37\x26\xb8\xa0\x07\x2c\x2b\x42\xc0\x67\x13\x60\xc3\x89\x3b\xd8\x11\x0a\x06\x57\x68\x50\xa5\xfd\x53\x3b\xd0\xee\xcb\xf7\x6d\x6b\x76\x06\x82\x17\x98\x7c\xb3\xb7\x6e\x83\xae\x36\x8a\x7c\xe5\xed\xfd\x87\xc6\xe8\x5c\xd4\x53\xec\x86\xe6\xc8\x2a\xcc\x1a\x5d\x1b\xfd\x2c\xfc\xa4\xe0\x6e\x05\x28\xd2\x1c\x8c\x7e\x82\x5c\xf8\x98\xa0\xd6\x45\x70\xb7\xb3\x08\x61\xc9\x51\xbb\x21\xe8\xe3\xe5\xc3\x3e\x75\xfa\xae\xb4\x9a\xb7\xf2\xcb\x40\x1b\xa8\x6d\xd4\x3e\xc8\xb3\x5a\x72\xf8\x6d\x0f\x10\x3e\x5c\x37\x04\x03\xe5\x04\x5e\x3f\x0b\x8a\x5b\x0b\x50\x2f\x41\x5d\xc3\x4f\x2a\x42\xf2\x70\x72\xee\x49\x07\xc0\x39\xdb\x9f\xc2\xff\x68\xd0\x59\x3b\xf2\xcf\x31\x97\xd1\x19\x79\xa6\x2c\x8d\x3e\x53\xf6\xba\x4f\x90\x2e\x37\x88\xdd\xd1\x62\x92\x8c\xf1\x30\x23\x44\xfe\x73\x72\x38\xe4\xec\x31\x8c\x3a\x7b\xec\xf7\xb3\xd0\xc7\xce\x4a\xd7\x66\xda\xdc\x79\xdc\x99\xa8\xe4\xbf\x8d\xae\xab\xa9\x92\x98\x3d\xbe\x6c\xfb\x74\x18\x6c\xeb\xae\x2f\x7b\xe9\x44\xbc\x81\x2f\x23\x76\x15\x77\x3f\xbe\x58\x2c\x30\x75\xda\x4c\xb0\xba\x1b\x28\xc4\x12\x8b\xb6\x8b\x47\x23\xbe\xee\xd7\x1a\xcd\x16\xf4\x06\x0d\x19\x07\x3a\x0a\xea\xad\x09\xc5\xa4\xf3\xe0\x51\x66\x5d\x70\xf3\x52\xb8\x34\x7f\x43\xd4\x6c\xb3\xcf\x76\x69\xfe\xfa\x99\xf6\x95\x1c\x43\xd8\xd2\x6f\xde\xbe\xa2\xad\xdf\x4d\x4c\x99\xb1\xac\xdc\xf6\x90\x4f\xa6\x44\xae\xa2\x28\x1a\x0f\x6d\x13\xb8\xe1\x63\xac\x83\xa6\x11\xaa\x81\x80\xd2\x6d\xff\xde\x96\xc3\xfe\xb9\xa5\xd4\x99\x54\xac\xdd\x81\xe8\x8f\x64\xe1\x45\x4f\x58\x5e\xaf\xa6\xcd\x01\x76\x21\xb2\xf4\xfb\x74\x2f\xfe\x5d\x4d\x47\xc0\x51\x1a\x83\xaa\x78\xc4\xf6\x91\xc6\x74\x86\x6b\xb0\xf5\x38\xd3\xe0\x7d\x3a\xed\xdb\x85\x54\xb6\x39\x67\x99\x81\x80\x47\xdc\xfa\x23\x19\xde\x78\xa1\x11\x8e\x21\x0d\x07\x00\x3e\xcc\x19\xa1\x8a\x44\x81\x09\x34\x67\x37\x03\xed\xc7\x97\xd6\x97\x47\x8c\x00\xe8\x88\x88\x88\x83\x66\xc3\xe5\x65\x45\x15\x3c\x07\x76\xea\x53\xc4\x03\x7c\xd2\x46\x5b\x43\x3e\x2a\x19\x69\x3b\xea\x2f\x42\x09\x12\x3d\x69\x3a\xed\x32\xec\x0e\x84\xfc\x42\xbd\xb0\xcd\x8e\x40\x2b\x9b\xcb\x6a\x74\x42\xb4\x59\x0d\x8e\x24\x9c\xac\x7d\x10\x85\xcc\xda\x21\xbc\xbe\xde\xa9\x19\xbc\xd5\xee\x2e\x1a\x84\x77\xe5\xf5\xb3\xb4\xce\xfb\x96\x57\x1a\xed\x5b\xed\xb8\xe6\x93\x09\xcc\xb3\x79\x92\xb8\x7c\x97\x06\xa8\xfb\xbd\xa4\x5e\xed\x1d\xc8\xd9\x04\xee\x56\xe3\xd2\xca\x71\x27\x7a\x69\xe1\x4e\x11\x8e\xf0\x72\xf1\xc7\xa9\x7e\x20\x3f\x44\x59\xdb\x58\xa0\xdd\x95\x25\x32\x32\x61\x87\x4a\x3c\x1c\x8d\xd1\x88\x53\x9b\x3d\x69\x8e\x2f\x43\x2f\x3b\x34\x5c\x33\xd4\x03\xed\x49\xfc\x17\x7f\xdc\x5b\x34\xe7\xd4\x23\x72\xad\x59\x68\x7c\x9c\x29\x1c\xae\x65\x0a\x25\x9a\x35\xd2\x96\x3d\xcd\xc7\x16\x79\xd4\xaf\x35\xbc\x4f\xd5\x85\xb1\xb0\x1b\x4a\x7c\xd3\xb2\x2b\x73\xb2\x9f\xc1\xef\x61\x59\x06\x1a\x0d\xec\x4f\x4e\xe1\xb9\x13\xa4\xe3\x2c\x77\x93\x3c\x53\xbc\xe6\x24\xa9\x1e\xc7\xc3\x06\x2b\x70\x1c\x29\x45\x45\x96\xf3\x3b\x85\x04\x56\xae\x3f\xa0\x12\xd2\x50\x9c\x1f\x18\xb8\x41\xf1\xdd\x5e\x52\xb1\x82\x76\x07\x20\xda\xd2\x02\xad\xd4\x46\x14\x87\xa7\xd5\x07\x53\xd1\x64\xc9\x58\xf8\x10\x17\x50\x4d\x27\x72\xcf\xe0\x29\xd7\xd6\x47\x9e\x95\xc4\x82\xcf\xe0\x2f\x1f\x71\x7b\x39\x64\x39\x87\xb6\x77\x79\xa7\x2e\x7d\xe8\x3b\xb2\xa6\x36\x4e\x6a\x55\x0c\x69\xcd\x25\xf7\xba\x3c\x0f\x06\x8c\x6a\xd3\x48\x83\x10\xd7\xce\xde\x2e\x7a\x54\x3e\x01\xb4\x3e\x6c\x2b\xfc\x1e\x9d\x68\x6a\x97\xd8\x9c\x58\x65\x72\x23\xb3\x5a\x04\x40\x48\xeb\x2e\x14\xdc\xbc\xbb\x8b\x6e\x12\x6d\xa5\x95\x45\x68\x50\x0c\x5a\xe7\x4f\x00\x3d\x8b\xf6\x38\x71\xc2\xe7\x2d\x7c\xb6\x15\x3d\xf8\xf2\x43\xd3\x32\x4a\x67\x69\xf0\x90\xec\x6a\xce\x4d\x12\xb8\x77\xa6\xe6\x9c\x44\x73\x6e\x45\x6b\xd3\xa6\xc6\x62\x64\x0d\x54\xd4\xc4\xf2\x71\x96\x3f\xa8\x92\xaa\x90\x0a\x5b\x69\x9c\x0b\x5f\xfb\x13\x7a\x03\xe2\x3f\x35\xbd\x37\xe8\x2a\xbb\x89\xbf\x53\x93\x7d\x43\x5e\xa0\x2f\x0d\x78\x4a\xea\x6f\x80\xf6\x97\x4c\x0a\xee\x97\x09\x4e\xf6\x30\x61\xb8\x5f\xce\x4f\x1f\x0e\x2e\x6a\x27\xb1\x38\x3d\x99\x38\x40\x71\x28\xcd\x38\x94\x5a\x1c\x20\xf9\x67\x4b\x3a\xee\x97\x89\x07\x00\x51\x5f\xec\x5d\x29\xc7\xba\x89\xe9\x91\x17\x19\xae\x44\x5d\xb8\x45\x9b\xe9\x4c\x78\xef\x12\x61\x72\xcc\x95\x47\x4f\xb4\xcf\x3e\xee\xf5\x53\xfa\x8c\x47\xb7\x32\xed\xd9\xbb\xef\xc9\xe8\x2e\x0d\xdb\xf4\x4e\x52\x30\x6c\xd8\xb9\x3b\x9f\x75\x13\xec\xa7\xad\x5b\x1a\xc9\x1f\x11\x19\x56\x30\x53\x7a\xb7\x27\x55\x5a\xd4\x19\x76\xcf\xfc\x66\x60\xe5\x6f\xc1\x43\xc9\xd2\x47\x9e\x43\x52\x53\x72\xc5\x77\xa5\x58\xe3\x7d\x85\xe9\x0e\x44\x74\x87\x16\x4b\x5d\x3b\x8e\x9e\xd4\x0e\x6a\x8b\x19\x88\x3e\x53\xa4\x26\x69\xdf\x6c\x86\x03\x0c\xcd\x62\x82\xf6\x7d\x15\xc4\xf9\x35\x43\x17\xea\xd5\x84\xd9\x86\x31\xa9\xa0\x92\xcf\x84\xdc\xbe\xc2\x64\x9d\xc4\x70\xd5\xf5\xb7\xcf\xd7\xdf\x7e\x7d\x56\xaa\xcc\x9a\x74\x0a\xf4\xd8\x9d\xca\xd2\x52\x7b\xe6\x5a\xdf\xc6\xc2\x03\x94\x2e\x47\xaf\x06\x4b\xab\x8b\xda\xc5\xb0\xc2\x8f\xef\xdf\x84\x00\xe5\x09\x91\xf2\xc0\x2b\xe1\x04\x7f\x6a\x16\x2c\x78\x5d\x6e\x92\xb4\x9f\x63\xfe\x92\xd0\x84\x8f\xca\x9c\x02\x7b\xff\xdd\x2d\x5c\x7f\xf3\xaf\x7f\x9c\x25\x93\x89\x89\xa9\x83\xd5\x2b\x09\x95\x74\x90\x52\x33\x39\xbf\x70\x70\xc9\xbf\xae\x2a\xb5\xbe\x3c\x67\xa1\x86\x5c\x84\x35\xe9\xa7\x70\x04\x8f\xb8\xf5\x47\xfa\x67\x5d\xce\x38\x19\x9c\xb6\x30\xb4\x47\x18\xe7\x00\xd3\x0e\x04\xed\xa1\x78\x16\x28\xdd\x83\x9f\x7d\xb6\x33\x11\x90\x0e\xfb\x8a\x31\x20\xfa\x39\x40\xe8\x67\x01\xa0\x9f\x0b\x7c\xfe\x39\x80\xe7\x88\xcf\x88\x03\xce\x4f\x70\x57\x2d\x9a\x41\x3b\x05\x68\x1e\xc3\xc9\x81\xc4\xdc\x38\xc8\x8c\x43\xc9\x08\xd9\x2f\x0d\x30\x07\x17\xf0\x0c\x7f\x59\x48\xf5\x38\x0c\x9c\xde\x50\x8b\x26\xab\x1b\xf0\x52\xdb\x60\x83\x1c\xce\xd8\x3e\x54\xe6\xef\x88\x2c\x3d\x14\x39\xe6\x4e\x83\xad\xcd\x8a\x53\xa4\xed\x81\x14\x64\x3a\xad\xcb\x60\xc3\x33\xc8\x84\xcd\x97\x5a\x98\xcc\xce\x00\x5d\x7a\x0e\x6a\x22\x86\x47\x00\x53\x13\xbc\x03\x3b\xfb\x3c\xf4\xc8\x7d\x84\xab\x4f\x7b\xdb\xee\xd5\xfe\x6d\xbb\x83\xdb\x45\x34\x31\x52\x6a\x7c\xae\x0a\x0f\x30\x06\x0c\xa0\xaa\x4d\xa5\x6d\x1b\x5c\x68\xb1\xcf\xc2\x12\xb5\x89\xdc\x63\xd8\x63\xfc\x47\x53\x10\x21\x99\x8a\xa2\xd8\x02\x1b\xa7\x05\x0a\x47\xf0\x84\x4b\x2b\x1d\xaf\xbb\x41\x1b\xc9\xeb\x7c\x6a\xd5\x2e\x85\x64\x1d\x40\x33\xac\xe0\xdf\xef\xda\xc5\xf6\x07\x1d\x52\x41\x96\x83\xbb\x04\x46\x52\x5d\xb2\xca\x3b\x2f\xa6\x17\x3e\xf8\xd8\xbf\xc3\xa4\xa9\xce\x70\xe6\xf1\x96\x3d\x56\x42\x8a\x38\x95\x48\x1f\x09\x8d\x85\xcb\x65\x03\x3c\x4c\x31\x94\x5b\xb2\x91\xd4\x31\x28\x1d\xd9\x60\xb4\x08\xa8\x1f\xe7\x68\xb3\x16\x4a\xfe\x16\x91\xc6\xb0\x71\x60\x29\xe4\x14\xed\x7a\x4d\xed\x42\x6a\x8e\x3b\x7d\x84\x3a\x4d\xbe\xe5\xf7\xb6\xb9\xe5\x47\x83\x76\xbd\x1e\xdf\x1e\xf9\xbc\xa6\x94\x32\x70\xd9\x19\xd4\x12\xbf\x94\x25\x85\x08\x3e\x68\x46\xe1\x62\xf3\xdd\xab\xdd\x05\x58\x42\x08\x91\xf0\x3f\xa4\x14\xfe\xcd\x47\x6c\x4f\x13\x9b\x58\x0c\xbc\xec\x71\x19\x90\x8b\x4f\xe6\x1a\x29\x0a\x7e\x2b\x40\x7d\x83\x65\xef\x6e\xec\xc0\x9d\x83\x08\xfa\x54\x64\x24\xb4\x2a\x2d\x62\x60\x0b\x0d\x8f\x41\xfa\xd6\x65\x90\xf9\x98\x3a\xee\xe3\xae\xae\x36\xfa\x07\x2b\xfe\xb6\xe9\x01\xe7\x1c\xdb\x08\x3e\xf1\x14\x7a\xf9\xe7\x1c\x1f\xe7\x42\xb8\x67\x97\x4a\xec\x15\xd2\xdc\xe0\x5a\x72\x32\xef\xf0\x9d\x47\x3b\x43\xad\x17\x0d\x5f\x09\xaf\x61\x73\x94\xde\x66\x14\x45\x51\x40\xa1\x9f\xd0\xa4\x04\xb1\xfa\x10\xce\xa0\x90\xfc\x63\x99\x31\x29\x05\xce\xef\xa9\x75\xe3\xd1\x51\xd5\xe5\xfe\x75\xd8\x4c\xae\xf8\x96\x99\xf3\x54\xfb\x57\x59\x6c\x84\x2c\x38\xec\x72\xfa\x3b\x65\xc9\x0c\x01\xda\x41\xf6\x37\xf1\x0d\xd2\x40\xbf\xd8\xa6\x79\xee\xcd\xe4\xa8\x96\x54\xf9\xa8\xb2\xe7\x52\xe5\xdc\x4f\xfc\xa8\xb6\xe1\xb2\xd7\x35\xf4\x78\x0d\x5b\x2f\x1f\x48\x42\x9f\xe6\xca\x64\xe4\x61\xc5\x20\x45\x5f\x3e\xe6\x91\x45\xf4\x28\xd7\x3f\xbe\x18\x7b\x6a\x31\x65\x62\xc7\x4c\x0e\x9c\x6f\x9f\xfb\xf8\x62\x80\xe4\xe1\xad\xee\xa1\xfc\xfb\xa4\xf4\xef\xd0\x73\x8c\x9e\x59\x9c\xf4\x28\x63\x70\x1e\x87\xd7\xa4\x4f\x79\x9a\x31\x48\x78\xf7\x6c\xe3\xc4\x07\x1a\xc3\x54\xf7\x1f\x6f\x7c\x91\x67\x1a\x3b\xc1\xc5\xaf\x72\xfb\x32\x69\xe9\x87\x1f\x6e\xf8\xf2\x31\xcf\x37\x06\x67\x31\xfc\xb4\xe3\x84\x59\xc4\x01\x60\xcf\x0c\x3e\xfe\xb1\xc7\x09\x9c\x0d\x3f\xfc\xe8\xe1\xee\x94\xe7\x1f\xc3\xca\xda\x3e\x0d\x39\xf9\x11\xc8\x20\xdd\xfd\x07\x22\x27\x3d\x05\x19\xe6\xf7\xe0\x99\xc8\xc7\x3f\x08\xf1\x65\xfc\x59\x88\x2f\x63\x8f\x43\x3a\xad\xa6\xae\xe5\x09\x0f\x45\x06\x45\x13\xbc\xe3\x5f\xd8\xd3\x8c\xdd\xc5\x1a\x78\x46\x12\x1a\xf4\x3e\x26\x09\x1f\xa3\x4f\x4a\x42\x83\xde\xec\x64\x97\xfd\xd1\x0b\x37\x31\x7f\xc6\x37\x8d\xce\xbe\x6c\x13\x77\x5d\x67\xa7\x5f\x7a\x05\x75\xf2\x56\xb1\x6f\x84\x79\x1f\x5e\xee\xa1\x7c\x50\x15\xc0\x32\x6c\x5e\x8a\xa2\xca\xc5\xcb\x5d\x5d\xf3\xcf\x09\xfc\x3f\x0d\xe8\x7c\x6e\xf0\x5a\xb6\x00\x67\x6a\x3f\x9a\x75\xda\x88\x35\x36\x35\xff\x0f\x00\x00\xff\xff\xc2\x0b\x21\x16\x94\x41\x00\x00")

func metaAppscodeCom_resourcedescriptorsYamlBytes() ([]byte, error) {
	return bindataRead(
		_metaAppscodeCom_resourcedescriptorsYaml,
		"meta.appscode.com_resourcedescriptors.yaml",
	)
}

func metaAppscodeCom_resourcedescriptorsYaml() (*asset, error) {
	bytes, err := metaAppscodeCom_resourcedescriptorsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "meta.appscode.com_resourcedescriptors.yaml", size: 16788, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
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
	"meta.appscode.com_resourceclasses.yaml":     metaAppscodeCom_resourceclassesYaml,
	"meta.appscode.com_resourcedescriptors.yaml": metaAppscodeCom_resourcedescriptorsYaml,
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
	"meta.appscode.com_resourceclasses.yaml":     &bintree{metaAppscodeCom_resourceclassesYaml, map[string]*bintree{}},
	"meta.appscode.com_resourcedescriptors.yaml": &bintree{metaAppscodeCom_resourcedescriptorsYaml, map[string]*bintree{}},
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
