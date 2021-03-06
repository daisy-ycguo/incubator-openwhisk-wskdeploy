// Code generated by go-bindata.
// sources:
// wski18n/resources/de_DE.all.json
// wski18n/resources/en_US.all.json
// wski18n/resources/es_ES.all.json
// wski18n/resources/fr_FR.all.json
// wski18n/resources/it_IT.all.json
// wski18n/resources/ja_JA.all.json
// wski18n/resources/ko_KR.all.json
// wski18n/resources/pt_BR.all.json
// wski18n/resources/zh_Hans.all.json
// wski18n/resources/zh_Hant.all.json
// DO NOT EDIT!

package wski18n

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

var _wski18nResourcesDe_deAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesDe_deAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesDe_deAllJson,
		"wski18n/resources/de_DE.all.json",
	)
}

func wski18nResourcesDe_deAllJson() (*asset, error) {
	bytes, err := wski18nResourcesDe_deAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/de_DE.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1520374115, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesEn_usAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x5b\x6d\x8f\xdb\x36\xf2\x7f\x9f\x4f\x31\x08\xfe\x40\x5a\xc0\x51\xd2\xfe\x71\xc0\x21\xc0\xbe\xc8\x75\xd3\x76\xaf\x4d\x36\xd8\xcd\x5e\x50\xe4\x16\x0a\x2d\x8d\x6d\xd6\x12\x29\x90\x94\x1d\x77\xe1\xef\x7e\x98\x21\xf5\x60\xef\x52\xd2\x3a\xed\x5d\xdf\xd4\x09\x87\x33\xbf\x19\x0e\xe7\x89\xca\xa7\x27\x00\x77\x4f\x00\x00\x9e\xca\xfc\xe9\x2b\x78\x5a\xda\x65\x5a\x19\x5c\xc8\x2f\x29\x1a\xa3\xcd\xd3\x99\x5f\x75\x46\x28\x5b\x08\x27\xb5\x22\xb2\x37\xbc\xf6\x04\x60\x3f\x1b\xe0\x20\xd5\x42\x47\x18\x5c\xd0\xd2\xd8\x7e\x5b\x67\x19\x5a\x1b\x61\x71\x1d\x56\xc7\xb8\x6c\x85\x51\x52\x2d\x23\x5c\x3e\x86\xd5\x28\x97\xac\xcc\xd3\x1c\x6d\x96\x16\x5a\x2d\x53\x83\x95\x36\x2e\xc2\xeb\x8a\x17\x2d\x68\x05\x39\x56\x85\xde\x61\x0e\xa8\x9c\x74\x12\x2d\x7c\x23\x13\x4c\x66\xf0\x5e\x64\x6b\xb1\x44\x3b\x83\xd7\x19\xed\xb3\x33\xf8\x60\xe4\x72\x89\xc6\xce\xe0\xaa\x2e\x68\x05\x5d\x96\x7c\x0b\xc2\xc2\x16\x8b\x82\xfe\x6f\x30\x43\xe5\x78\xc7\x86\xa5\x59\x90\x0a\xdc\x0a\xc1\x56\x98\xc9\x85\xc4\x1c\x94\x28\xd1\x56\x22\xc3\x64\xb2\x2e\x5a\xc7\x34\x79\x0d\x4e\xeb\x02\x9c\x0e\x8a\xcc\xa0\x56\xfe\x17\x08\x95\x83\xdd\xa9\x0c\x74\x85\x6a\xbb\x92\x76\x0d\x55\xd0\x09\x6a\x2b\xd5\x12\x04\x94\x42\xc9\x05\x5a\xc7\xc4\xba\x22\xae\xa2\x08\xac\x4a\xd2\x64\x21\x8b\x96\xfc\xb7\xd7\x6f\x7f\x9d\x82\xd9\xae\xb4\x71\xc3\x07\xf0\xde\xe8\x8d\xcc\xd1\x82\x00\x5b\x97\xa5\x30\x3b\xf0\xf4\xa0\x17\xb0\x5d\x09\xf7\xcc\xc2\x1c\xb1\x77\x3c\x5f\x67\xc6\x00\x69\xd4\x8e\x16\x1d\xd9\x72\x85\x45\x15\x44\xc3\x4e\xd7\x66\x92\x09\xc9\x54\xd3\xb1\x6c\xd0\x58\x92\x1d\xb3\x8f\x54\x8e\x15\x0e\x74\xa0\xea\x72\x8e\x86\xcd\x63\xd7\x1e\xda\x64\x59\xe4\x05\xa3\xfe\xc3\xae\xc2\xca\x5e\x56\xa8\x3e\x1e\x2a\x3b\x47\xb7\xa5\xe3\xc8\x0a\x49\x5e\xc1\xae\x85\x66\x83\x66\xb2\x0f\x4f\xc7\xd0\xf3\x3e\x92\xd3\xf8\x33\xff\x85\x5e\xfc\x37\xbd\x79\x51\x88\x65\x2a\x2a\x99\xae\xb4\x8d\x39\x8e\x87\xf2\xfa\xfd\x05\x7c\xfe\xf9\xf2\xfa\xc3\xe7\x89\x1c\x87\x8f\xbf\xc7\xf4\x5f\x6f\xae\xae\x2f\x2e\xdf\x4d\xe2\x5b\xbb\x55\xba\xc6\x5d\x84\x29\x2d\x6b\x23\xff\xe0\xbf\x80\xcf\xbf\xbc\xf9\x6d\x0a\xd3\x0c\x8d\x4b\xc9\x6e\x11\xae\x95\x70\x2b\x3a\x16\xf2\xd5\x84\x88\xd9\xc8\x53\x18\x6b\xb5\x90\xb1\x60\xef\x17\x99\x15\x7c\x93\xe3\x42\xd4\x85\x03\x69\xe1\xff\x7e\xbe\x7c\xfb\xe6\x45\xb2\xb5\xeb\xca\xe8\xca\x7e\x3b\xc5\x2a\x45\xa1\xb7\x69\xe0\x11\x4b\x51\x4c\x04\x2d\xd1\x38\xd7\xce\xa9\x86\xec\xd2\x86\xe5\xd6\xfb\x26\xb0\xae\x0c\x6e\x24\x6e\x23\x7c\xed\x8a\x81\x36\x4c\x5f\x1c\x5c\x8f\xaa\x10\x6a\x82\x84\x35\xee\x26\x1f\xe9\x1a\x77\x53\x81\x7b\x4b\x97\x42\x89\x25\xe6\x83\x86\xae\x8c\xfe\x1d\x33\xd7\xe5\x5c\xa7\x61\x8e\x50\x0a\xb3\xc6\x1c\x1a\x0e\x53\x4c\xc5\x7c\x52\xca\x05\x31\x65\x82\x28\x26\x19\xe7\xd8\x84\x90\x91\x53\x3d\x08\xfa\x13\xd8\xb6\xc9\x2a\xc2\xb7\x5b\x9f\xac\xf4\x08\x42\x1f\x9e\x0b\xb4\xb6\xb1\xf6\x04\xd6\xd6\x19\x19\xe5\xec\x8f\xae\xb6\x68\xe8\xa2\x48\x85\x39\x98\x5a\x39\x59\xb6\x49\x6a\x82\x04\x67\xe2\x46\xe0\x35\xd0\xb5\xab\xea\x29\x60\xbd\xbb\x6d\xd0\xcc\xb5\x8d\xb1\x0c\xab\xe3\x4c\x39\xde\xa4\xa5\xb4\x94\x1b\x38\x92\xc6\x03\xe9\x87\x15\x02\x51\x90\xf7\x66\x3e\x9a\xd2\x2d\x91\x16\x94\x76\xe0\x59\xd5\x06\xf3\xe4\xdf\x43\x16\x39\x92\x58\xc9\x81\x24\x43\x12\x29\x1b\x10\xc9\xd7\xc9\x19\x73\x44\x92\xd4\xd2\x9c\x26\x2a\xa8\x32\xd4\x54\x1c\xeb\xf3\xe9\xee\x2e\xa1\xdf\xfb\xfd\xed\x0c\x16\x46\x97\x70\x77\x97\x58\x5d\x9b\x0c\xf7\xfb\x49\x32\xfd\x81\x8d\xc9\x24\xb2\xe6\xac\x2c\xba\xd3\x64\xb5\xe6\x19\x93\x76\x60\x47\x52\xb1\xfd\x8b\xd3\xf5\xac\xe4\x72\x9b\x0a\xee\xa7\x52\xa7\xd7\xa8\x46\x55\xa6\x1d\xe0\x77\x00\xef\x38\x4d\xf9\x5a\x95\xc2\xd8\x95\x28\xd2\x42\x67\xa2\x88\x48\xbc\x69\xa8\x7a\x45\x64\x08\x12\xd6\xcb\xe3\xdd\xb0\x11\x45\x8d\x76\xa2\x40\x85\x6e\xab\xcd\xfa\x64\x91\x52\x39\x34\x0a\x1d\x08\x47\xea\xd6\xa6\x18\xd1\xb5\xcb\xa8\x69\x26\x54\x86\x45\x11\xcd\x67\x97\xbf\x24\xf0\x83\xa7\xa1\xa2\xb2\xdb\x39\x55\xc0\x42\xc8\x38\xf7\xf3\x2e\xb5\xe7\x32\x0f\x77\xb1\xac\x0a\x74\x08\xb6\xa6\x23\x5d\xd4\x45\xb1\x4b\xe0\xaa\x56\xf0\xb9\x6d\x0c\xda\x9a\xf9\x33\x65\x02\x83\xa5\xde\x20\x54\xc2\x38\x29\x8a\x62\xd7\xf5\x54\xc2\x5a\x74\xc3\xa7\xd0\x43\xea\x1b\xb4\xd4\x3a\xe1\xea\x58\x1d\xf5\xfc\xf9\xf3\xe7\x67\x67\x67\x67\xbd\xb3\xe8\xe9\x70\xcd\x5b\x81\x08\x88\x70\x92\x54\x1e\x2d\x60\x3e\xc5\x44\x8d\x69\x72\x08\xf3\x08\x6f\x9c\x61\x27\x3b\xfd\xac\xfb\x7b\xa7\x0b\x19\x3c\xef\x9b\x7e\x31\x37\x78\xe2\x93\xe5\x8d\xd9\xef\x40\xe4\x09\x16\xcc\x74\x59\x0a\x95\xa7\xdc\x54\x71\x55\x49\x51\x2e\x15\x2e\xa5\x4a\x24\x22\xf4\xee\x2e\xc9\xca\x7c\xbf\x0f\xad\xd8\xdd\x5d\x42\x1b\xdd\xae\xc2\xfd\x9e\x23\x25\xed\xdd\xef\x6f\x93\x64\x50\x36\x97\x8f\xbb\xb4\xf1\xe7\x91\x31\xd4\xdd\x1d\x15\xb3\x41\x00\x81\xdc\xef\x6f\x61\x25\xc2\xa0\xa1\xaf\x70\x7b\x43\xa6\x4b\x8f\xcf\xad\xce\x9b\x75\x78\x10\x40\x92\x0c\x34\xa1\x41\x44\x73\xa0\x7f\xa6\x8a\x1d\xcf\x29\x4a\x36\xd4\x71\x35\x6f\x3a\x8a\x07\x15\x1d\xd4\x33\xc7\x0a\x55\x8e\x2a\x7b\x8c\x39\xbb\x4d\xa7\xcb\xe9\xae\x48\xd4\xa6\xe7\x0f\x8a\xf9\x1a\xc7\x79\x18\x05\x05\x86\xda\xc4\xea\xb2\xf3\x83\x19\xc8\xc3\xaa\xff\x0f\x73\x44\xa3\xcf\xe3\xfc\xe4\xeb\x4e\xf0\x7e\x98\xfb\x73\xce\x70\xe2\xcd\x88\x21\x19\x3e\xc7\x9b\xa3\x69\xd6\x29\x27\x39\x84\x2a\xf4\xce\xa7\xe6\x1c\x46\xe4\x33\x40\xdb\x9b\x0f\x61\x81\xbc\x36\x74\x92\x41\x6c\xbf\xfe\xf9\xeb\xfc\xad\xd1\x71\xa1\x6b\x95\xa7\x01\x6f\x88\x54\x51\x07\x28\xd0\x45\x63\xf0\x76\x25\xb3\x15\x6c\x79\x7e\x4f\xb8\x72\x5f\x37\xba\x15\x42\x56\x1b\x43\x86\x69\x14\x6c\xc6\x09\x9c\xa4\xfc\x6f\xe2\x20\x2c\xeb\x42\xf6\x9b\x5c\x16\x84\x69\x53\x1a\xc6\x98\xb1\x49\xb0\x5f\xe5\x66\x02\x7a\x93\x30\x83\xdc\xe1\xe7\x33\x10\x45\xbf\xf4\x6d\x8f\x8d\x70\x98\x76\x47\x10\x02\xc2\x60\x6b\xeb\x17\x9d\xa7\x43\x2e\x0d\x66\x2e\x78\xbf\xf1\x73\xe0\xb1\x09\xfb\x9b\xab\xab\xcb\xab\xeb\x08\xee\xb3\xe3\xff\xc0\x93\xc3\xbd\x85\xb3\xb3\x81\xf4\x63\xcc\xe1\x45\x5b\x2b\xbd\x55\x29\x55\x0a\xe3\x57\x9d\xa8\xc8\x54\x61\x57\x02\xdd\xe8\x1c\xb4\x2a\x76\x60\xeb\xca\xbf\x03\xbd\xe0\x81\x6b\x62\x77\xd6\x61\x09\x73\xa9\x72\xa9\x96\x16\xb4\x81\xa5\x74\xab\x7a\x9e\x64\xba\x6c\xc7\xcd\xc3\xf9\xd2\x98\x26\x67\x66\x06\x85\x8b\xc1\xe4\x77\x39\x60\x92\x03\xb7\xdc\x4a\xb7\x02\x7e\xd0\x83\x12\xad\x15\x4b\x7c\x45\x8b\x68\xcc\x7e\xcf\x63\x6d\xbf\x96\xe9\xdc\x2f\xd0\x8f\x91\x6e\xa6\x07\xc9\xdf\x95\x41\x48\xf9\xbd\x9b\xf2\x17\x41\x5a\x20\xe6\xa9\x54\x1b\xbd\x8e\x01\xfa\x91\xc3\x16\x85\x0b\x4f\xc6\x17\x92\xb6\xc1\x76\xc5\x4f\x43\x01\xa9\xf3\xcf\x72\x61\xe9\xaf\x41\xbb\xc6\x5d\x3b\x43\xa1\x7a\x57\x38\x6d\x86\xe6\x43\x2d\x0d\x8f\x1b\x3e\x35\xc6\xbc\x25\x7f\x0c\x7c\x46\x65\x36\x43\xc6\x54\x69\xe7\x83\x5d\x44\xe0\xdb\xfe\x34\x92\x63\x35\x53\x53\xbf\xcb\xe3\xc0\x7e\x45\x3d\x26\x94\xab\xf7\x52\xda\x52\xb8\x2c\x56\xbe\x93\x82\xad\x7b\xd0\x86\x9c\x45\xe4\x4d\x3c\x95\xea\x78\xec\xed\xd7\x03\x06\xc8\x35\xfa\xc1\x12\x0b\xe1\x63\xe5\xf0\x46\x44\x65\x8f\xc9\xc1\x94\xd5\xaf\x36\x6a\x0c\x2b\x11\xfa\x7f\x72\x2f\x51\xc8\x98\xd9\x2e\xfc\x2a\x5d\xf3\x70\x24\xed\x40\x93\x64\x85\xdf\x84\xa5\x7b\x75\x3c\x40\xa5\x0d\x63\x17\xfc\x3e\xcc\x7b\xfc\xcf\x29\x76\x6e\x20\x8e\x98\xfa\xea\x31\x80\x8e\xec\xca\x57\xc1\x23\x7a\x66\xc1\x4f\x79\xbc\x29\xf1\x8b\x43\x65\x1b\xd0\xf8\x85\x73\x18\xa9\xf3\x35\xaa\xd8\x74\x89\xb1\x01\x66\x77\x95\x97\xe8\xdf\x35\x43\xec\xed\x86\xc8\x61\x58\xd3\x65\x32\xca\x6f\x32\xeb\x5d\xdf\xc9\x36\xf5\xd0\x53\xaf\x31\xdf\x9e\x56\x5a\x04\xdf\x81\xc2\x5c\x17\x92\x19\x3b\x2b\x0b\xb5\x6b\x7d\x83\x82\x48\xef\xd8\x47\xed\x1a\x86\xa8\x2d\x84\x51\x35\x6a\x53\x3c\xde\x73\xfd\x60\x2b\xb4\xd0\x37\x57\xbf\x32\x02\x1e\x75\xf1\x55\xfa\x74\xd0\x63\xdf\xfa\xc7\xea\x29\x40\x4a\x51\x2c\xb4\x29\xa3\x96\x7b\xdb\xac\x0f\x21\x48\xe0\x83\xd9\x81\x58\x0a\xa9\xc6\x5a\x7a\x63\xd2\xdf\xad\x56\x6d\xb0\xcd\xca\x7c\xe0\x4d\xf3\x9f\xd7\x97\xef\x40\xaa\xaa\x76\x90\x0b\x27\xe0\x6d\xb0\xc6\xb3\xac\xcc\x9f\x51\xe8\x1d\x96\x24\x2a\xd9\x4d\xe0\xf9\x38\x63\x4f\xd4\x0f\xdc\x8b\x5e\x28\xf7\xae\x7b\x10\xae\x66\x3c\xe0\xe6\x0d\x95\x24\xea\x4c\x28\x5f\x75\xcc\xd1\xe7\x7d\xcc\x61\x2e\x2c\xe6\xa0\x55\xdf\x9f\xee\xb3\x4a\xe0\x7d\x81\xc2\x22\xd4\x55\x2e\x1c\x1e\x85\x45\x4e\x8f\x59\x51\xe7\x78\x04\x4f\x58\x10\xb0\xc5\x79\x60\x3c\x6a\xf6\x70\x6b\x86\x3d\x2f\x66\x88\xb0\x2b\x81\x0b\xe7\xdb\x2a\xed\x56\x9c\x64\xf9\xba\x2c\x6a\x15\x2e\x4b\x73\xa3\x66\xde\x16\x5a\x61\x78\x69\x2c\x89\x0b\x7e\xa9\x30\x9b\x72\x45\x02\xd6\xe6\xec\x9a\x8b\x4f\x11\x2f\x25\xa9\x5f\x89\x9e\x81\x77\xb7\x9f\xd8\xea\xda\xf5\xa3\x40\x02\x1f\xbb\xe8\xda\xc4\x00\xda\x36\x6b\xe3\x04\xb9\x47\x53\x05\x8c\xe4\xab\xa0\x4e\x63\xa6\x94\xda\x10\x87\x69\x2e\xcd\xa4\xe8\xf5\xa0\x5a\xa4\x47\x6b\xf7\x4a\x4b\xe5\x6b\x25\xdf\x7b\x39\x0c\x15\x3f\x55\x28\xdd\x3d\x9d\x51\x6f\xd7\x68\x65\xb9\x59\x38\x0c\x5d\xc3\x6a\x64\x82\x3a\x71\xb1\xc1\x34\xd7\xd9\x1a\x63\xdf\xa4\xfd\x20\x14\x73\x15\x1b\x84\x73\x26\x04\x59\x72\x65\x3d\x52\x31\xca\x02\x53\x51\x18\x14\xf9\x2e\xc5\x2f\xd2\x46\x9f\xf3\x7f\xa4\x8b\x11\x28\xc1\x53\x46\x78\x7f\x7c\x7d\xf5\xee\xe2\xdd\x4f\xd3\xbb\x98\x66\xc3\xe3\xfa\x98\xad\x30\xaa\x1d\x95\x1a\x74\xd1\xda\xf1\x8a\xd6\xe8\xa0\x3e\x35\x33\xd2\x5b\x10\x0b\x87\xc6\xd7\xad\xaf\x7c\x62\xa1\x6c\x78\x3b\x74\x47\x82\x3c\x7e\x33\x7a\x74\x2a\xe9\x7f\x7c\xd1\x2b\x1d\x21\x47\x37\x7e\x3b\x59\x32\x55\xcb\x39\x56\x06\x33\x0a\x73\xa9\xc1\xaa\x10\x59\xd4\x7d\xa9\x9c\x24\x39\xba\xc8\x43\x91\xcc\x4f\x74\x3e\x2a\x1e\xce\x86\xb7\xb2\x28\xc0\x6a\xad\x28\x8a\x76\x12\x66\x50\x85\x08\x69\x7d\x97\xc0\xdd\x3d\x6e\x0f\xd8\x59\x87\x62\x22\xf6\x60\x89\x53\xea\x7b\xbb\xd2\x75\x91\x13\x3c\x8b\x2e\x81\x1b\xeb\x07\x5d\xbe\x0b\xe7\x67\x2f\xa6\xe6\x5f\xe3\x13\xee\x16\x11\xd3\x8f\x1c\x25\xe1\xf2\x12\xa8\x94\xbb\xdf\x77\xd0\xa5\xf3\x71\xee\x11\x22\x39\xfe\x88\xcd\xe0\xe1\x8d\x09\xe5\xfd\xcd\x81\x36\x13\x95\xe6\xc3\xb6\xfe\x17\x6d\xe3\xc0\x0a\x59\x4a\x97\xca\xa5\xd2\x26\x0a\xa9\x71\xe9\x10\x9c\x79\x0b\xa3\xe2\x5f\xc7\xbd\x05\x85\x7f\xcf\x6e\xaa\xf4\x6c\x25\xd4\x12\xc5\x3c\xfa\x01\xcd\xaf\xad\xc4\xb6\x99\xb1\x8d\xde\xc5\xce\x0f\xd3\x5a\x1e\x09\x5c\x90\x78\x6a\x08\x27\xf8\x02\x23\xb0\x69\xa1\x97\xa9\x95\x7f\xc4\x00\x14\x7a\x79\x2d\xff\xe0\xd4\xea\x37\x1c\x68\xdc\xb9\xa8\x50\xfc\x38\x4a\xcd\x73\xf3\x85\xdf\x4b\x4e\xd8\xdf\xbd\x9c\x0c\xa5\xc4\x52\x9b\xdd\x10\x1a\x4f\x71\x2a\xa0\xef\xbe\xff\x3b\x43\xfa\xdb\x77\xdf\x4f\xc6\x44\xf9\x57\xd7\xb1\x66\x24\xac\x9e\x04\xe6\xa5\xb7\xcf\xff\xbf\xa4\xff\xc6\xf1\xf0\x5c\x29\xad\x8c\xae\xd0\x38\x89\xb1\x7c\xd5\x44\xc0\x5e\xbc\xf2\xd3\x48\x67\x24\xb6\xf3\x48\x3f\xa4\xea\x98\x35\x73\xcb\x87\x63\x62\x13\x12\x73\xcd\x0e\x47\x91\x51\x3a\xd0\xb5\xb3\x32\xe7\x83\xf8\x60\xc4\x46\x5a\x98\xd7\xb2\xc8\x87\x87\x5a\xac\x8a\x0f\x07\x86\xdc\x76\x52\x28\x68\xbd\xff\x20\x20\xa8\xa3\x80\x1e\xac\xcd\xa3\xba\xbb\xbb\x24\xfc\x6d\x63\x6e\xea\xf8\xa5\x0a\x83\x1b\xfa\x83\xc8\x46\xda\x40\x86\xda\x94\x83\xfe\x92\xc5\xc2\x44\xd3\x5a\x07\x2a\x2a\x8d\x8e\xba\xec\x07\xca\xf3\x68\x23\x7d\x52\xf7\xcc\x68\xc3\x6c\x8e\x07\x2f\x83\x55\xcd\xbd\xb1\xcb\x41\x88\x39\x2a\x77\x9a\x9e\xc1\x62\x81\x19\x95\xae\xda\xad\xd0\xcf\x74\xc7\x21\x35\xa3\xce\xd1\xc9\x53\x48\x85\x47\x53\x9a\xa6\x60\xc8\xb4\x72\x82\x3f\x14\x53\x7a\xda\xf8\x94\xa5\xf7\x5e\x2e\xd8\x28\x53\x40\x3c\x38\xd7\x0f\x19\xe7\x78\x16\xb5\x0d\xed\xbd\x1f\x92\x05\xa2\xc3\x8e\x6b\xdc\x42\xbd\x2f\x0f\x53\xbd\x41\x63\x64\x9e\x63\xac\x7d\x24\x84\xfd\x0f\x11\xbb\x97\xa7\x6e\x6b\x53\x2b\xf4\x1f\x16\xa6\x1e\x54\x2a\x6d\x5a\xd5\xf3\x42\xc6\x3e\xb1\xf6\xa7\xc2\xb4\x4d\xe7\xea\xbf\xb5\x14\x16\xfc\xc6\x7b\xf3\xb6\x19\x85\x0b\x8e\x2d\x73\x84\x8d\xb4\x72\x5e\xf8\x66\x8e\x1a\x59\x8a\x8e\xfc\x64\x46\x4d\xec\x8e\x1a\x23\xad\x06\xbe\x5d\x64\xac\xa1\xc9\xd9\xe2\x3c\xc5\x2f\xfc\x29\xc9\x70\x1a\xbf\xdf\xce\x70\x53\xc9\x1d\xad\xca\xe9\xff\xcf\x3d\x9f\x7b\x5d\x25\x5d\x04\x32\xe5\x16\xe7\x33\x9f\xdc\xc3\x9f\xc2\x86\x81\x3e\xc3\x23\xed\x8d\x05\x08\xee\x49\xa3\x01\xf2\xb0\x7e\xff\x3d\x69\x20\xe0\xbf\x87\xe8\x35\xed\xf0\x83\x56\x1b\x0a\xf7\xa1\x25\xe8\x44\x38\x3d\xad\xbd\x3f\x7f\xf3\x8f\x9b\x9f\x26\xb7\x37\x4c\xfd\xb8\xde\x26\x9f\x2f\x53\x8b\xc2\x64\x2b\xb2\x57\x73\x31\xda\xf6\x32\xf6\x8f\x82\x9a\x1d\xed\xc5\x38\x6c\x48\x9b\x18\x42\x71\xb5\x4b\x20\x23\x25\x12\x41\x39\x8e\x1e\x7f\x76\xe4\x38\x31\x6a\x10\xb4\x36\xac\xfa\x97\xcb\x81\x7f\x8e\x72\xfe\xc0\xf8\x3c\x58\xe4\x15\xfc\xc8\x08\xba\x7f\xfd\xc0\x2f\x76\xc4\xec\xb1\x00\x86\xbf\x24\x7e\x3c\x86\xfe\xe3\x68\xf3\x98\x1f\x20\x3d\xb9\x7d\xf2\x9f\x00\x00\x00\xff\xff\xd1\x4e\x2d\xf4\xc3\x36\x00\x00")

func wski18nResourcesEn_usAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesEn_usAllJson,
		"wski18n/resources/en_US.all.json",
	)
}

func wski18nResourcesEn_usAllJson() (*asset, error) {
	bytes, err := wski18nResourcesEn_usAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/en_US.all.json", size: 14019, mode: os.FileMode(420), modTime: time.Unix(1522283010, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesEs_esAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesEs_esAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesEs_esAllJson,
		"wski18n/resources/es_ES.all.json",
	)
}

func wski18nResourcesEs_esAllJson() (*asset, error) {
	bytes, err := wski18nResourcesEs_esAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/es_ES.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1520374115, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesFr_frAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\xe6\x52\x50\xa8\xe6\x52\x50\x50\x50\x50\xca\x4c\x51\xb2\x52\x50\x4a\xaa\x2c\x48\x2c\x2e\x56\x48\x4e\x2d\x2a\xc9\x4c\xcb\x4c\x4e\x2c\x49\x55\x48\xce\x48\x4d\xce\xce\xcc\x4b\x57\xd2\x81\x28\x2c\x29\x4a\xcc\x2b\xce\x49\x2c\xc9\xcc\xcf\x03\xe9\x08\xce\xcf\x4d\x55\x40\x12\x53\xc8\xcc\x53\x70\x2b\x4a\xcd\x4b\xce\x50\xe2\x52\x50\xa8\xe5\x8a\xe5\x02\x04\x00\x00\xff\xff\x45\xa4\xe9\x62\x65\x00\x00\x00")

func wski18nResourcesFr_frAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesFr_frAllJson,
		"wski18n/resources/fr_FR.all.json",
	)
}

func wski18nResourcesFr_frAllJson() (*asset, error) {
	bytes, err := wski18nResourcesFr_frAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/fr_FR.all.json", size: 101, mode: os.FileMode(420), modTime: time.Unix(1520374115, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesIt_itAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesIt_itAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesIt_itAllJson,
		"wski18n/resources/it_IT.all.json",
	)
}

func wski18nResourcesIt_itAllJson() (*asset, error) {
	bytes, err := wski18nResourcesIt_itAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/it_IT.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1520374115, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesJa_jaAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesJa_jaAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesJa_jaAllJson,
		"wski18n/resources/ja_JA.all.json",
	)
}

func wski18nResourcesJa_jaAllJson() (*asset, error) {
	bytes, err := wski18nResourcesJa_jaAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/ja_JA.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1520374115, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesKo_krAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesKo_krAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesKo_krAllJson,
		"wski18n/resources/ko_KR.all.json",
	)
}

func wski18nResourcesKo_krAllJson() (*asset, error) {
	bytes, err := wski18nResourcesKo_krAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/ko_KR.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1520374115, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesPt_brAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesPt_brAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesPt_brAllJson,
		"wski18n/resources/pt_BR.all.json",
	)
}

func wski18nResourcesPt_brAllJson() (*asset, error) {
	bytes, err := wski18nResourcesPt_brAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/pt_BR.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1520374115, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesZh_hansAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesZh_hansAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesZh_hansAllJson,
		"wski18n/resources/zh_Hans.all.json",
	)
}

func wski18nResourcesZh_hansAllJson() (*asset, error) {
	bytes, err := wski18nResourcesZh_hansAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/zh_Hans.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1520374115, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesZh_hantAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesZh_hantAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesZh_hantAllJson,
		"wski18n/resources/zh_Hant.all.json",
	)
}

func wski18nResourcesZh_hantAllJson() (*asset, error) {
	bytes, err := wski18nResourcesZh_hantAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/zh_Hant.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1520374115, 0)}
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
	"wski18n/resources/de_DE.all.json":   wski18nResourcesDe_deAllJson,
	"wski18n/resources/en_US.all.json":   wski18nResourcesEn_usAllJson,
	"wski18n/resources/es_ES.all.json":   wski18nResourcesEs_esAllJson,
	"wski18n/resources/fr_FR.all.json":   wski18nResourcesFr_frAllJson,
	"wski18n/resources/it_IT.all.json":   wski18nResourcesIt_itAllJson,
	"wski18n/resources/ja_JA.all.json":   wski18nResourcesJa_jaAllJson,
	"wski18n/resources/ko_KR.all.json":   wski18nResourcesKo_krAllJson,
	"wski18n/resources/pt_BR.all.json":   wski18nResourcesPt_brAllJson,
	"wski18n/resources/zh_Hans.all.json": wski18nResourcesZh_hansAllJson,
	"wski18n/resources/zh_Hant.all.json": wski18nResourcesZh_hantAllJson,
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
	"wski18n": &bintree{nil, map[string]*bintree{
		"resources": &bintree{nil, map[string]*bintree{
			"de_DE.all.json":   &bintree{wski18nResourcesDe_deAllJson, map[string]*bintree{}},
			"en_US.all.json":   &bintree{wski18nResourcesEn_usAllJson, map[string]*bintree{}},
			"es_ES.all.json":   &bintree{wski18nResourcesEs_esAllJson, map[string]*bintree{}},
			"fr_FR.all.json":   &bintree{wski18nResourcesFr_frAllJson, map[string]*bintree{}},
			"it_IT.all.json":   &bintree{wski18nResourcesIt_itAllJson, map[string]*bintree{}},
			"ja_JA.all.json":   &bintree{wski18nResourcesJa_jaAllJson, map[string]*bintree{}},
			"ko_KR.all.json":   &bintree{wski18nResourcesKo_krAllJson, map[string]*bintree{}},
			"pt_BR.all.json":   &bintree{wski18nResourcesPt_brAllJson, map[string]*bintree{}},
			"zh_Hans.all.json": &bintree{wski18nResourcesZh_hansAllJson, map[string]*bintree{}},
			"zh_Hant.all.json": &bintree{wski18nResourcesZh_hantAllJson, map[string]*bintree{}},
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
