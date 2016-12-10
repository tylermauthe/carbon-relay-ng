// Code generated by go-bindata.
// sources:
// admin_http_assets/app.css
// admin_http_assets/app.js
// admin_http_assets/index.html
// DO NOT EDIT!

package web

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

var _admin_http_assetsAppCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8a\xce\x4b\x8f\xb1\x4a\xce\xc9\x4f\xcc\x8e\xd5\x51\x00\x72\x74\xe1\xec\x94\xc4\x92\x44\x5d\x64\x81\x0a\x64\x9e\x1e\x8c\x0d\x64\x22\x24\x20\xe2\x19\x99\x29\xa9\x0a\xd5\x5c\x0a\x40\x90\x92\x59\x5c\x90\x93\x58\x69\xa5\x90\x97\x9f\x97\xaa\xa0\x98\x99\x5b\x90\x5f\x54\x92\x98\x57\x62\xcd\x55\x0b\x08\x00\x00\xff\xff\x9c\x9e\x21\xb0\x7a\x00\x00\x00")

func admin_http_assetsAppCssBytes() ([]byte, error) {
	return bindataRead(
		_admin_http_assetsAppCss,
		"admin_http_assets/app.css",
	)
}

func admin_http_assetsAppCss() (*asset, error) {
	bytes, err := admin_http_assetsAppCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "admin_http_assets/app.css", size: 122, mode: os.FileMode(420), modTime: time.Unix(1481021499, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _admin_http_assetsAppJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x58\xdf\x6f\xa4\xb6\x13\x7f\xcf\x5f\x61\xa1\x48\x18\x85\xb0\xf9\x7e\x1f\x2a\x95\x28\xba\xa6\xd7\x9e\x7a\x0f\x91\xaa\xf4\xda\x3e\xa4\x7b\x92\x03\x0e\x41\xf1\xda\xc8\x36\x9b\xac\xf6\xf8\xdf\x3b\x36\x06\x8c\x77\x2f\x97\xdb\x5c\xdb\x95\x22\xc0\xcc\x8f\xcf\x7c\x3c\x33\x1e\xb2\x26\x12\x91\xa6\x41\x17\x88\xd3\x47\x44\x78\xd5\x32\x22\xb3\x95\x28\x5b\x46\x71\x54\x10\x79\x2b\xf8\xa9\xa4\x8c\x6c\x4e\x79\x15\xa5\xe8\x26\xe2\xd5\x35\x55\xa2\x95\x05\x85\xc7\xa8\xad\xb3\x5b\x21\xb4\xd2\x92\x34\xd1\x32\x39\x3f\x3a\x02\x6b\x59\x21\xb8\x96\x82\x31\x2a\x71\x74\x45\x6a\xfe\x56\x33\xab\x7b\xac\x0a\xd1\x58\xbd\x63\xe9\x19\x39\x06\x7f\xc4\x48\xdc\xb5\xbc\xd0\xb5\xe0\xb8\x17\x4c\xd1\x28\x06\xb7\x56\x28\xd9\x1e\x21\xd4\xbf\xcd\x08\xd8\xd7\x0a\xa0\xdf\x2c\xcf\x61\xd5\x84\xf2\x56\xf0\xbb\xba\x82\xa5\x51\x11\x47\x8b\xc2\x2e\x2e\xa2\x64\x90\xfa\x40\x6e\x19\x0d\x84\xb4\x59\xf3\x64\xae\xe9\xa3\xac\x35\x95\x81\x98\x74\xcb\x6a\x91\xd7\xbc\xa4\x4f\x93\xc2\x8f\x8c\x14\x0f\xac\x56\x3a\xd0\xb8\x1d\xd6\x77\x55\x2e\xab\x4a\xd2\x8a\x68\x11\x7a\x21\xe3\x8b\x5d\xa5\x6b\xd1\xea\x10\xbc\x34\x6b\x20\xfa\x40\x37\xc0\xe2\x16\x2e\x39\x8a\x7f\x80\x4b\xdc\xc1\x63\x37\xea\xfe\x44\x95\xae\x39\x31\x14\x7f\xde\xc2\xa2\x9c\xa4\x7c\xf7\x47\x13\xf3\x6b\xc2\xea\xf2\xb2\x2c\xc1\x80\xe1\x7f\xf1\xf1\xe6\x63\xbe\x3c\xf9\x2b\xbf\x39\x3b\xfd\x7e\x79\x82\x73\xfb\x98\xbc\x39\x5e\x9c\x07\x3a\x16\xfc\x87\x4d\x43\xad\x16\x56\x94\x97\xf8\x92\xb1\x4f\xef\x6a\xa9\x74\x72\x45\x74\x71\x9f\x7c\xc2\xb0\x61\x0a\x08\xa3\x5c\xff\x42\xd4\x7d\xcd\xab\x64\x11\xda\xa1\x15\x7d\x02\x1b\x78\xcc\x98\x04\x99\xcc\x30\x3f\x49\x75\x2b\xf9\xf8\x68\x7e\x10\x9a\xce\xa7\xec\x02\x13\x2d\x4d\x66\x12\xe6\x67\x18\xaa\xd5\x1f\xc6\x3e\x98\xd6\xb2\xa5\xe7\x81\x84\x96\x9b\x1d\x2d\x64\x2b\x07\x00\xfd\xfc\xd4\x38\xcb\xa1\x5a\x87\x0a\x13\x19\xde\xe3\x13\x79\x1e\xef\x08\x53\x3b\x2e\xbb\xe0\xd9\x45\xe7\xb4\x7c\xe9\x41\xb2\x33\x8b\x5d\x82\xcd\x96\x05\x3b\x56\x55\xef\x80\x83\x6f\xcd\x5b\x7d\x87\xfa\x57\xe8\xe2\x02\x45\xaa\x5d\x45\xfb\xe2\x1c\x1d\xec\x23\x36\x8c\x72\x6e\x92\xac\xab\xd7\x9b\x74\xa2\x3b\x24\x3f\x4b\x9b\xab\xe7\x91\x81\xba\x7c\x4a\x7a\x1c\xb6\x87\x64\x15\xd5\x13\x95\x25\xd1\x24\x19\x50\x3a\x03\xda\xb5\x1a\xf3\xae\xf7\xda\x57\x63\x17\x7a\xc1\x76\xb9\xef\x5f\x73\xb3\xc5\x5d\x35\x04\xef\xe4\x8b\xa1\xcb\xc1\xab\x1e\xb5\x6f\x0d\xd2\xd1\x6b\x5e\x7d\x72\xf6\x8f\xbd\x8f\xa1\x7b\x96\xa5\x27\xb6\x93\x0f\xfb\x7b\xec\x3e\x27\xd9\xb1\x22\x6b\x8a\x93\x4c\xdf\x53\x3e\xc1\x86\xde\xd0\x4c\x9b\xf6\x42\x6c\x9e\xe8\x44\x0a\x04\x98\x8e\x66\xa9\x94\x3b\x56\x27\x90\xdb\x95\xaa\x72\x04\x32\x99\x21\x3c\x83\x1b\x21\xbb\x65\x40\xfc\xa8\x27\xe9\x4a\xac\xe9\x3e\x16\xa6\x7d\x36\x99\x68\xf9\x96\x2b\x1c\x5f\x4a\x8a\x36\xa2\x45\xaa\x75\x37\x8f\x84\x6b\xa4\x05\x2a\x29\xa3\xd0\x95\x87\xd3\x01\x41\xeb\x82\x56\xc1\x45\x86\x62\x74\x82\x8c\xb5\x09\xf4\xc8\x5b\xaf\x84\xb7\xb1\x6d\xb1\x71\x0e\x62\xdd\xb3\x2c\x84\x79\x03\xfc\x41\x49\x3b\x22\xa7\xc3\x04\x6f\x4d\x87\xcd\xa1\x6e\x2a\x73\x64\xbf\xe7\xe0\x0c\x8a\x29\xff\xee\x2c\x45\x7f\x92\x1a\x2a\xfa\x7f\xff\x3f\xeb\x82\x5c\x98\x9d\x45\x07\x64\x03\xe8\x7f\x65\x22\xc0\xf9\xf1\x1a\xe8\xff\x45\xaa\xec\x27\xe9\xc0\x64\x99\x0e\xf9\x67\xd3\x65\xf2\xf9\x4d\x12\x66\x18\x1f\x6c\xed\x99\x7b\xbc\xfd\xb5\x86\x01\x05\x48\xb7\xad\x31\x45\xbf\x35\x42\xb0\xdc\xf6\xd4\x14\x05\xef\xdc\xee\x98\x73\x1b\x8e\x6d\x7b\x5e\x47\x61\x26\x0d\x1e\x0e\x69\x29\x46\xf5\x6b\xfb\xc9\x3f\x14\xd0\xbf\x92\x5f\x13\x71\x26\x66\x9f\x34\x3b\x90\x7d\x89\xb9\x9e\x2f\x4b\xd7\x16\x26\x3d\xab\x93\xce\x88\x9f\x47\x00\xe8\x7b\xb8\x41\x0c\x2f\x46\xbf\x0f\x79\x5f\x1a\xfe\xf8\xfb\xda\xca\x18\x47\xe6\x67\x0b\x63\xf4\xf8\xfa\xba\x70\x07\x41\x98\xb8\x30\x0c\x7f\x09\xfe\x9b\xd8\xef\xec\x76\x3b\x46\x34\x66\xfe\x36\x03\xf5\x21\x58\xe6\x93\xba\x8f\x28\x45\x2f\x60\x75\x06\xcb\x1f\xfb\x77\xd1\xa5\xe8\x60\xda\xe0\x8f\x87\xa4\x21\xbb\xe9\xce\xb7\x99\xa8\xed\x67\xdb\x7b\xae\x34\xe1\x85\xfd\x6e\xb1\x0b\x56\x17\x0f\x08\x35\x5d\x35\x8c\x68\xfa\xbb\x84\x3a\x8d\xdb\x06\x72\xae\xdf\x8d\x2b\x2b\x7b\xaf\x57\x2c\x76\x99\x8b\x00\xf2\xad\x20\xb2\x1c\x4a\xd8\x2d\x4f\x9f\x9c\xb9\x87\x65\xfc\x8e\x9c\xa1\x70\x95\xe2\x8f\x94\x03\xf9\x2e\x1a\x7b\x3d\x0f\xdf\x8a\x87\x59\xa0\xf3\x91\x74\xee\x22\x2b\x98\x50\x14\xfb\x66\xbd\xcf\x82\x6e\xc7\x74\x61\x94\xd8\xcb\xcd\x97\xb5\x5a\xd5\x4a\xe1\xb8\x57\x8c\xf7\x19\x1f\x8a\x1d\x59\x17\xb9\x73\x95\x8e\xe3\xbe\x12\x6c\x0d\xcb\x93\x17\x0b\x33\xff\x2c\x04\x37\x41\x0f\xff\x24\x00\x63\x9b\x59\x80\xea\x06\xb6\x7e\xe9\x23\x39\xf2\xaf\x6e\x4a\x45\xf3\x8c\x80\x7c\x57\x2d\xd3\xf3\x66\x8f\x70\xb0\x41\x5e\x93\xc4\x3e\x99\x53\x3b\xea\x8c\xe3\xbf\x03\x00\x00\xff\xff\x0c\x69\x41\x3c\xce\x10\x00\x00")

func admin_http_assetsAppJsBytes() ([]byte, error) {
	return bindataRead(
		_admin_http_assetsAppJs,
		"admin_http_assets/app.js",
	)
}

func admin_http_assetsAppJs() (*asset, error) {
	bytes, err := admin_http_assetsAppJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "admin_http_assets/app.js", size: 4302, mode: os.FileMode(420), modTime: time.Unix(1481021499, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _admin_http_assetsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x3b\x7b\x6f\xdc\x36\xf2\xff\xe7\x53\x30\x42\x81\x24\xf8\x45\x92\xed\xb4\x68\x6b\xec\x1a\x48\xd3\xe4\xd7\xde\x35\x4d\xe0\xb4\xf7\x40\xd1\x2b\xb8\x22\x77\xc5\x98\x12\x55\x92\x5a\xdb\x17\xf4\xbb\xdf\x0c\x25\xad\x1e\xab\x95\xb4\xf1\x63\x93\xd6\xb1\x45\x6a\x38\x9c\xf7\x0c\x49\x71\xf6\xf0\xfb\x37\x2f\x7e\xf9\xf7\xdb\x97\x24\xb6\x89\x3c\x7b\x30\xc3\x3f\x44\xd2\x74\x35\xf7\x78\xea\x91\x74\xe5\xd3\x2c\x9b\x7b\x11\xd5\x0b\x95\xfa\x9a\x4b\x7a\xed\xa7\x2b\x0f\x21\x39\x65\x67\x0f\x08\x99\x25\xdc\x52\x12\xc5\x54\x1b\x6e\xe7\x5e\x6e\x97\xfe\x37\x5e\xfd\x22\xb6\x36\xf3\xf9\x9f\xb9\x58\xcf\xbd\x7f\xf9\xbf\x3e\xf7\x5f\xa8\x24\xa3\x56\x2c\x24\xf7\x48\xa4\x52\xcb\x53\x18\xf5\xe3\xcb\x39\x67\x2b\xde\x18\x97\xd2\x84\xcf\xbd\xb5\xe0\x97\x99\xd2\xb6\x01\x7a\x29\x98\x8d\xe7\x8c\xaf\x45\xc4\x7d\xd7\x78\x4a\x44\x2a\xac\xa0\xd2\x37\x11\x95\x7c\x7e\xbc\x85\x86\x71\x13\x69\x91\x59\xa1\xd2\x06\xa6\x2d\x30\x9a\xdb\x58\xe9\x2d\x08\x29\xd2\x0b\x02\xac\xcf\x3d\x11\x21\x82\x58\xf3\xe5\xdc\x0b\x82\x10\x7e\x96\x74\x8d\x9d\x01\xfc\x02\x60\x84\xb6\xc2\x4a\x7e\xf6\xa2\x10\xd8\xb9\x13\xd8\xcf\xff\x4f\x28\x4b\x44\x3a\x0b\x8b\x97\x0e\xee\xa1\xef\x93\x9f\xa8\xe5\xc6\xc2\x7c\x49\x26\x24\x67\x84\xa6\x8c\x00\x9c\x58\x0a\x68\xbc\x78\xf7\x8e\xf8\x7e\x87\x02\x63\xaf\x25\x37\x31\xe7\xb6\xa2\x23\x0c\x13\x7a\x15\xb1\x34\x58\x28\x65\x8d\xd5\x34\xc3\x06\xa0\x0c\x37\x1d\xe1\xb3\xe0\x24\x38\x0a\x23\x63\xea\xbe\x00\xe6\x09\xa0\xc7\xab\xa9\x79\xe3\x04\x44\x25\xb1\x31\x4f\xf8\x1d\xce\xed\xbb\x09\x3a\x14\x0c\xce\x03\x46\xd8\x21\xf6\x87\x5f\x5e\xff\xf4\x15\x31\xb1\x48\x9c\xd4\xce\xb9\xc9\x54\xca\x82\xf7\x86\xfc\xf8\xf2\x1b\x62\xf2\x0c\xcd\x86\xa8\x65\x09\xc8\x25\xcc\x98\x5a\x53\x88\x98\x33\x41\xc9\x9f\x39\xd7\x82\x9b\x8a\x4f\x40\xfa\x9b\x58\x12\x69\x01\x01\xf9\xf6\x77\xd7\x57\x58\x0d\x31\x3a\x9a\x7b\x68\xc8\xe6\x34\x0c\x95\x31\x41\xc9\x35\x32\x8a\x0e\xf3\x15\x90\xb1\x06\x46\xbf\x0e\x4e\xea\xb6\x63\xef\x3d\x90\x3c\x0b\x0b\x34\x53\x31\xea\x82\x95\xf0\x38\xf8\x12\xf0\x95\xad\x7e\x6c\x0f\x7f\xe3\x29\x13\xcb\xdf\x91\x85\x59\x58\x78\xe4\x83\xd9\x42\xb1\x6b\xa2\x95\x44\xc3\x57\x51\x8e\x7c\x7b\xa4\x96\xdc\x2b\x71\x05\xd6\x95\xd2\xf5\x82\xea\x8a\x79\x26\xd6\x24\x92\xd4\x98\xb9\x57\xbe\x28\xfe\xf8\x22\x5d\x73\x70\x6c\xaf\xc4\x07\xbd\x62\x45\x9d\x1f\xe1\xb8\xf6\x48\x74\x1b\x2a\x52\xae\xcb\x77\x7d\x78\x7d\x24\xb2\x01\x01\x30\xb4\x03\xb1\xd0\xa0\xa3\x8d\xe6\xbd\xae\x2b\xcd\x42\xba\x41\x1f\x02\xfe\x81\xb9\x22\x25\x25\xcd\x0c\x27\xd5\x43\x73\xda\x5c\x36\xa0\x2b\x76\xe1\x4f\x03\xc6\x59\x65\x05\x45\x23\x2b\xd6\xbc\xf5\xd6\x11\xbf\xa1\xf3\x07\x95\xf0\x06\x71\x05\x81\x52\x74\xd0\xed\x1a\xbf\xa0\xec\x35\xb7\x5a\x44\x26\x3c\xf9\x32\x06\x55\xa3\x88\xbf\xa3\x68\xac\xae\x77\x10\xf3\x2c\xcc\x65\xbf\x50\x40\xe1\x61\x00\x5c\xd5\xb2\x00\x8d\xd7\x30\xe5\xc3\x83\x5d\x8a\x2c\xd5\x9e\x40\xcb\x25\x03\x7c\x03\x5d\x92\xeb\xb9\xf7\x1a\x3a\x5f\x58\xd9\x30\x84\x3e\x55\x68\x75\x59\x8c\x94\x8a\x5e\x34\xb5\x0e\x38\x2c\xbe\xd0\x3c\xe3\x14\x62\x6d\xd1\x21\x52\xe2\x1e\xc0\xd2\x3f\x7c\x70\x4f\x41\x62\x56\x7f\xfd\x05\xec\x63\xa3\x81\xa0\x45\xaf\xf4\x13\xe6\x1f\x9f\xb4\x75\x17\x9f\x9c\xfd\x83\x4a\xc1\x9c\xbd\x82\x7b\x9c\x74\x64\x6f\x29\x24\xa1\x0a\x47\xd1\x70\xbf\x91\x4b\xc6\x53\xc3\x59\x47\xdb\x38\xa6\x4a\x7b\xdd\x7e\xbd\xdd\xe9\xc0\x4b\xf3\x85\xd0\x1f\xef\x82\x28\xf5\x4e\x20\x4e\x0e\x81\x95\xbc\x70\xa2\x34\x78\x50\x3f\x24\xf4\x6e\x11\x82\x90\x3d\x44\xcf\x2c\xc6\x89\x3d\x58\x61\xa0\x11\x90\xcc\x52\xac\x82\x5a\xac\x7f\x48\xbe\xe6\x12\x7e\xaf\x68\x74\x8d\x6a\xb2\x3d\xd2\x19\x1b\x9d\x9c\x1c\xed\x37\x94\xff\xe1\x64\xb0\x6b\xd0\x0e\x29\x74\xf9\x85\x2e\x54\x77\xd3\x8d\x1a\x46\x3c\xc5\xc4\x0a\x23\xfb\x4e\xd2\xe8\x42\x0a\x63\x0f\x66\x63\xaf\xa9\x8d\x62\x92\x41\x24\x11\x57\x83\x96\xe6\xe0\x4c\xbe\x80\x34\x2c\xd2\xd5\x38\xa8\xe6\x2b\x3e\x88\xf1\x79\x84\x7a\x34\xb7\x64\x8d\xcd\x78\xb0\xc0\x58\xe0\x64\x15\x2c\x2a\x01\x6f\xc9\x6a\xa7\x54\x2a\xbb\x59\x04\x85\x54\x76\xdb\x57\x0d\xb9\x91\xcb\x14\x60\x27\x99\x11\xc0\x8d\xde\xf9\x95\xf5\x23\x48\xc3\x98\xf8\x20\xe4\xbb\x70\x28\xa2\x0b\x88\x8e\x3c\x51\x6b\xbe\x31\xa0\xc7\x5f\x08\xb0\x8a\xab\x27\x00\xb5\xc9\x3b\x2b\x79\x9d\xc5\x58\x63\x92\xcd\x93\x5f\x0c\xf3\x23\xa1\x23\x28\xa3\xc3\x33\x4c\x0e\x37\x72\x85\x1e\x67\x70\x9d\x60\xcf\xe7\xfc\x52\x0b\xa0\xdc\xf4\x58\xf7\x52\xe9\xa4\xac\x9d\x75\x09\xf6\x0a\xba\xbc\x8a\x76\x7c\xef\x53\xc6\x88\x7b\x58\x69\x95\x67\x24\xa6\xc6\x5f\x72\xce\x16\xc0\x73\x95\x5f\x96\x6e\x10\x88\x05\x54\x90\x08\x4c\x06\x8c\x55\xf3\x3e\x7e\x02\x6f\xd4\xba\x74\xfc\x83\xb8\xd7\x1b\xc9\x86\x7c\xe0\x67\x7e\x39\xec\x49\x07\xf2\x20\x5d\x7b\x50\xa5\x1d\xb3\xbf\x07\xe9\x40\x49\x36\xee\x11\x3a\x48\xf9\xe5\x14\x30\xa8\x74\x6f\xc9\x6d\x36\x16\x72\x38\xaf\xa9\x05\xbf\x9f\x5c\x5b\x0e\xd2\xe3\x17\xfd\x03\xdd\x60\x91\x66\xb9\x2b\x9c\x12\xc5\x70\xb5\x04\x62\xaf\x04\x11\x80\xa1\x7a\xa5\x3f\x2a\x7c\x6c\xce\x52\x96\x6e\x1e\xc9\x20\xda\xf0\x18\xde\x63\x11\xe7\xc0\x34\xae\xcf\x35\xdf\xa1\x11\x37\x2b\xa6\x42\xf4\xcf\x58\x5d\xb6\x7d\x1d\xad\x23\x00\x05\x38\x07\x1d\xa0\xdb\x61\x31\x19\x4d\x87\xd0\x70\xad\x95\x0e\x32\x6a\xa1\x17\xca\xdf\x97\x57\x19\x8f\x2c\x2e\x56\x40\x7b\x3c\xc9\xec\x35\xa9\x12\x17\x62\x1a\x20\xb7\x9d\xc8\xdb\xaf\x26\x58\xde\x2d\xab\x05\x02\x44\xa5\x96\x14\x1f\xc7\xd5\x82\x60\x87\xa0\xdf\x5e\x67\x48\x64\x9e\x2c\xb0\xf8\xef\xe7\x06\xe2\x59\xc5\x4d\x82\x8f\xbd\xdc\xdc\xc8\xa4\x00\xed\x6d\x98\x94\x43\xb3\xcb\xa4\x20\x9c\x38\x2e\x89\x55\x04\x32\x85\x82\xba\x1a\xe0\x81\x6e\xa7\x87\x62\xaf\xe0\x31\x25\x99\x32\x02\xd7\x7d\x25\xf4\x53\x28\xbe\x89\x7f\x8c\xf9\x0c\x8c\x92\x48\x01\xd9\xea\xc9\x5d\x59\xe3\x8e\x17\x8b\xdc\x5a\x88\x6b\xa5\xd8\x17\x36\x25\xf0\xcf\x37\x89\xfb\x93\x69\x91\x50\x7d\xed\x9e\x17\x52\x61\x8a\x2d\x74\x5a\x64\x56\xa7\x53\x26\x0c\x26\x05\xd6\x11\x57\x2d\xf1\xe7\x0c\xf2\x5d\x31\xcd\x3e\x64\xf7\x45\xcd\xbd\xaa\x8d\x10\x6d\x68\xbb\x02\x79\xbe\x5a\x41\xb5\x45\xad\x1a\xab\x41\xe8\x6a\x75\x6b\xe5\x47\x3d\xe9\x27\x50\x80\xbc\xca\xd3\xa8\x58\xc4\xee\xae\x23\xce\xc7\x4a\xf5\x37\xb9\x45\x17\x47\x66\xa9\x1d\x02\xfc\x11\x33\x2e\xf0\x4b\x1e\x1b\x1e\x3d\x19\x82\xfc\x27\x15\x76\x1c\xea\xee\x2a\x1c\x5a\x57\x38\xb4\x36\x92\xfd\x6b\x1c\x1a\x2c\xf3\x74\xbc\x78\xa1\x93\xaa\x7e\x07\x08\xc2\x7e\x95\xd8\x29\x90\x95\xb4\xa7\xc0\xa2\xbc\x6f\xa9\x78\x6a\xd8\xf7\x8d\xcb\xa7\x9b\x44\x82\x82\xec\x43\x57\x50\x20\x8d\x00\x9c\xac\xca\x6b\x4b\x7c\x1c\xcf\xd2\x95\x5b\xd6\xe9\x0e\xb1\x96\xb9\x66\xee\xb9\x90\x01\x98\x11\x6c\x88\x9e\x56\x22\x2c\x63\x18\x1a\xe4\x47\xe6\xc0\x16\x86\x9d\xe9\xaf\x54\x3f\x50\x4f\x96\x25\x1b\xa7\xc4\xe4\x09\x26\x38\xba\xfe\x4c\x2a\x2c\x54\x9b\x8b\x7b\xde\x66\x15\xea\x1a\xe3\xaa\x2b\x01\xb7\xd4\x55\x60\xdb\x57\x59\x0e\xdb\x0d\xd5\x55\xe2\xd8\xa5\x30\x87\x19\xf7\x63\x72\x49\x35\xe1\x57\x99\xe6\xc6\xb8\x94\xf0\xb9\x28\xaa\x08\x89\x9b\xf5\x49\x6e\x97\xd8\x9a\xb0\x44\xc9\xed\x58\x3d\x79\xb0\xa2\x18\xf9\xaa\x02\x78\xc5\x99\xd8\xb4\x7b\x79\x6b\x9a\x5c\xf8\x9f\xdf\x8e\xfc\x6f\x7f\xff\xbf\x2f\xc2\xbd\x4d\xae\x9a\xe5\x86\x56\x57\xa3\x19\x2d\x94\x1f\x43\xae\x85\x54\x0f\x75\x8d\x79\xd2\xa8\x9a\xff\xcc\x69\x6a\xc5\x7f\x61\x55\x46\x2a\x64\x9f\xac\x4d\x8e\x68\x12\xd3\x6b\xa5\xc5\x4b\xf7\x7c\x97\x1a\xc4\x19\x6e\xa8\xbd\x02\xc5\xc7\x6a\x0e\xd7\x3b\x88\xe1\xf3\x5d\xbf\x54\x72\xf8\x64\x97\x2e\xe7\x10\xbc\xdc\x7e\xc5\xe0\xd6\x29\x00\xf1\xdb\xdb\x37\x45\x6c\x07\x5e\xb3\x14\x75\xbf\x2b\xf8\x8b\x33\xe7\x4b\xfe\x48\x4a\x82\x2e\x88\xa5\xa4\x21\x31\xd7\xd5\x07\x06\x7d\x23\x1d\x0f\xe4\x82\x5f\x0f\x2e\x78\x1c\x10\x5a\xc7\xf8\x39\xc6\x41\x8e\x46\x18\xc3\x1c\x3d\x04\xf2\x2e\x53\x4a\x0e\x01\xbc\x85\x8a\x5d\x0e\xf1\x87\x47\xdb\xe8\xbd\xf3\x13\x72\x4f\x3b\xc9\x28\xf5\x3d\x96\x58\x18\x07\x8a\xe8\x35\xb0\xb4\x80\x7c\x7f\xed\x11\xaa\x05\xf5\x63\xc1\xc0\x06\xc1\x2e\x75\xce\xdd\xd7\x0e\x18\x9a\x06\xce\x06\x2b\xb4\x22\x5d\x2a\xcf\x6d\x30\x83\xd9\x0c\x9e\x26\x6e\x8f\x40\x1b\xda\x73\x48\x82\x36\xc0\xf5\xe8\xe9\xd2\xe0\xe0\x09\x07\x4e\x83\xe3\x47\x56\xa3\xdd\xb1\x1b\x5b\xf1\x9e\x79\x93\x45\x5a\x0f\x3a\xe9\xdf\x84\x77\xe1\xc6\x49\xfd\x8e\x37\xe0\x9d\x89\x35\xad\x92\xa1\x55\xea\x80\x71\x03\x31\xb6\xf9\x51\x4a\x0f\x4f\x67\x18\x86\x46\xed\x10\xa4\xba\xd6\x48\xa8\x58\xc5\x76\xc8\x20\x21\x74\x0d\xa7\xbe\x1b\xbc\x2d\x44\xec\x68\xfc\xf0\x88\xd1\x74\xc5\xf5\xa3\x53\xf2\x90\x05\x2a\x95\x22\xe5\x4f\xc9\x23\x54\x0c\x74\x55\x3d\x7f\xa1\x55\xb0\xc9\x26\x79\x2b\x93\x4c\x3d\x2b\xbd\xe9\x3c\x93\x8e\x59\x3f\x72\x0e\x5a\x04\xe8\xdb\xc7\xde\xbb\x07\xd3\x8f\xdf\x19\xe0\xa6\xae\x63\x81\xc1\x8c\xe0\x0d\xd9\x68\xcc\x18\xb8\x4e\x9f\xe7\x0e\x78\x74\x83\x07\x52\x31\x41\xee\x8e\x8b\xcc\xa5\xad\x41\x36\xd4\x85\x6f\xc4\x2a\xbd\x2f\x56\x86\x83\xd8\xf7\x75\x08\x29\x42\xd9\xd3\x43\x1e\x29\x7e\xc4\x97\x39\xf7\x91\x61\xf7\x5b\x75\xf5\x1e\x88\x61\xae\x08\xfe\xce\xaf\xab\x45\xd6\x2f\x90\x78\x47\x0e\x90\x3a\xdb\x36\x55\x61\xd8\x3b\xeb\x3d\x32\x51\x50\x3e\xca\x45\x8b\x78\xc3\x53\xf6\x5c\x4a\x57\x3e\x0e\x6d\x19\xba\x19\x1c\xd2\x1d\x54\xb5\x4f\xce\xaa\xd5\x83\xa3\x69\x7c\x41\xd9\x3d\x31\xeb\x0c\xdf\xb5\x98\x6c\x12\xff\xd4\xb5\x5e\x09\x6d\x6c\xd9\x56\x1a\x3f\x9a\x36\xc2\xe0\x77\xd3\x3f\x50\x13\x8f\x1c\xd2\xee\x5c\x52\xde\xa3\x0a\xdf\xba\x3c\x59\x29\xb1\x6a\x8d\xab\xb1\x84\x3c\x30\xf5\xef\xaa\x04\x5c\x31\xd0\xe8\x18\xe7\xa1\x06\x3e\x30\x1b\xad\xdd\xdb\xf3\xa9\xbb\xb7\xe7\x7b\xef\xde\xde\x23\x4b\xe5\xba\xaf\x62\x6a\xd3\xdc\x23\xcc\xd5\x28\xb6\x8e\x13\xca\x37\xf7\xc9\x63\xb1\x01\x03\xe5\x58\x74\xb1\x50\x57\xdd\xcf\x02\x0a\x6b\x2c\xca\x96\xd2\x12\x5b\x35\x4c\x8b\xdf\x4f\x8c\xec\xb7\x65\x9d\x52\x86\x80\x76\xd5\x32\x91\xf0\x1d\x19\xf9\xce\xf7\xbe\xea\xc0\x3d\x79\xf7\xeb\xf6\x3f\x11\xec\xee\x7b\xb5\x3f\x3e\xaf\x1b\x8d\xef\xcc\x37\x9f\x9e\xe3\xbe\x50\xb8\xf9\xd4\xdc\xed\x06\x55\xdd\xdf\x55\x37\x45\x20\xab\x68\x4e\xfe\x46\xd7\xf4\x9d\xbb\xf5\xe0\x90\xcd\xf7\xfe\xaf\x71\xc7\x83\xbc\x45\x47\x63\x84\x5a\xbc\xe6\x42\x20\x8f\xe1\x25\x11\x7c\xac\xee\x4a\x10\xa3\x5c\x3b\xa3\x2b\x6e\x88\x54\x94\x91\x25\x85\xc4\xb6\xb9\x2c\xd1\x77\x87\x83\xbe\xa7\x57\xc1\x4a\xa9\x95\xe4\x34\x13\xc6\x5d\xe4\xc0\xbe\x50\x8a\x85\x09\xdf\xe3\x5d\x93\xeb\xf0\x38\x38\x86\x9f\xb2\x35\x7e\x3f\x64\xfa\xed\x9a\xf7\xdd\x8b\x3d\x6d\xbc\xbb\x88\x8e\xc0\x1d\x02\xa8\xa0\xf1\x48\xe9\xbd\x09\x94\x5e\x01\x89\x27\xc1\xf1\x51\x58\x76\x4e\xbf\xc3\x32\x8a\x0a\x4a\x65\xa3\x72\x1d\xf1\x29\x7c\x03\x9f\x80\x24\x92\x2a\x67\x4b\x18\xcb\x3b\xe2\xac\x50\xe6\xc2\xaf\x25\x71\x84\xc2\x3d\x0a\x9b\x7d\xbe\xcd\xa4\x19\x91\x85\xbb\x5a\xb4\x8b\x9c\xc2\xfd\x70\x01\x14\x82\xef\x59\x9e\x40\x94\xb6\x10\x21\x04\xf8\x5f\x9e\xe1\xd6\xaa\x8b\x23\xaf\x15\xa3\x32\xc0\x9b\x3f\x3d\xb7\x62\x12\x7c\xd9\xbd\xf6\x32\x8b\x9f\xb5\xdf\xbb\x8b\x61\xde\xd9\xaf\x0e\x29\x71\xbe\x7d\x4a\x3e\x7c\x70\x0f\xd5\x8e\x56\xfc\xac\xe5\x4c\xee\xb1\xb1\x8d\xbc\xbc\xc1\x0e\xf2\xd6\x56\xf1\x36\x07\x18\x0f\x9a\xf7\x67\x1a\x00\xf5\x1c\x9d\xcb\x33\x74\xc1\x25\x52\x30\xf7\x30\x8b\x79\x67\x6f\x8b\x5c\x36\x0b\xdd\x9b\x16\x6c\x37\x9d\x16\x9c\xe3\x80\x2a\x3c\x3b\x14\x53\xea\xb3\xb2\x76\x9d\x52\x1e\xb4\x4b\x69\x44\xea\x4a\xdf\x5d\x55\x74\xa7\x76\x6e\xc0\x0f\x1f\xdb\x0e\x1d\xd7\x76\x2f\x1d\xec\xbe\x83\x30\x41\xcc\xb8\xa5\xe1\xd5\x3b\xcf\x53\xc5\x8c\x03\x2a\x31\x3b\x14\x13\xc4\x7c\x7c\xf2\x75\x70\x04\xff\x1f\x9f\x9e\x1c\x1d\x7d\x39\xf8\x09\x44\x4f\xcd\xd2\x23\x78\x9c\x78\x1f\xc1\x17\xf0\x3b\x04\x7f\x4a\x1e\xc5\xca\xd8\x53\xbc\x05\xf8\x68\x3f\xa1\xef\xbc\x55\x56\xf8\xc1\x12\xa2\x4b\xfb\x02\x5b\x4f\x52\xc7\x14\xce\xf8\x92\xe6\xd2\x7a\x8d\xdd\x86\x88\xa6\x11\x97\x8f\x9f\x78\x67\x2f\xa4\x32\x7c\x3b\x57\xef\x28\x10\xca\xca\xa0\x53\x01\x2c\x5b\xc9\xbf\x31\x8d\xba\xc0\x29\x8a\x58\xd2\x9d\xa3\x95\x94\xab\xe4\x5d\xc7\x3e\x00\x77\x79\x7f\x16\x16\x17\x82\xff\x17\x00\x00\xff\xff\xa3\x20\xf7\x4a\x21\x3c\x00\x00")

func admin_http_assetsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_admin_http_assetsIndexHtml,
		"admin_http_assets/index.html",
	)
}

func admin_http_assetsIndexHtml() (*asset, error) {
	bytes, err := admin_http_assetsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "admin_http_assets/index.html", size: 15393, mode: os.FileMode(420), modTime: time.Unix(1481021499, 0)}
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
	"admin_http_assets/app.css": admin_http_assetsAppCss,
	"admin_http_assets/app.js": admin_http_assetsAppJs,
	"admin_http_assets/index.html": admin_http_assetsIndexHtml,
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
	"admin_http_assets": &bintree{nil, map[string]*bintree{
		"app.css": &bintree{admin_http_assetsAppCss, map[string]*bintree{}},
		"app.js": &bintree{admin_http_assetsAppJs, map[string]*bintree{}},
		"index.html": &bintree{admin_http_assetsIndexHtml, map[string]*bintree{}},
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
