// Code generated by go-bindata.
// sources:
// data/.config_schema_v3.json.swo
// data/.config_schema_v3.json.swp
// data/config_schema_v3.json
// DO NOT EDIT!

package schema

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

var _dataConfig_schema_v3JsonSwo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x9a\x7b\x8c\x1d\x55\xfd\xc0\x87\xee\xb6\x3c\x0a\x14\x7e\xe4\x67\x14\x25\x69\xa7\x98\x18\xdb\x7d\x60\x37\x42\x41\x22\x8f\x16\xc5\x84\x47\x08\x02\x52\x97\x61\xf6\xce\xb9\xbb\x43\xe7\xce\x4c\x67\xe6\xee\x76\xbb\x7d\x06\xa5\x04\x53\x13\x95\x84\xa4\xd8\x28\x46\x94\x6a\x03\x2a\x50\xb1\x46\x21\x46\x24\x21\x96\x44\x40\x13\x8c\x11\x93\x56\x6b\xff\x51\x82\x60\x2b\x12\xf4\x33\x8f\x7b\xef\xcc\xec\xdc\x3b\xb3\x77\x6f\x0b\x89\xf3\x4d\x3e\x7b\xb3\x33\xe7\x9c\x39\x8f\xef\x39\xe7\xfb\xfd\x9e\x33\x36\x7c\xcb\xb5\xd7\x2d\xbd\x78\x70\x44\x42\xce\x91\xa4\x65\x8f\x2f\xb9\xed\xde\x3b\xf7\x9f\xb2\x61\xb1\x24\x69\xaa\xa9\x0b\x43\xca\x15\x4b\xb3\xd5\x02\xc9\xa4\xad\x61\x81\x43\xb6\x70\xdc\xa1\x8a\xa5\x09\xfe\xd4\x6c\xcb\x15\x03\x55\xdd\x10\x43\x6e\x65\x42\xd4\xd4\x21\x4d\xf5\x54\x5e\x98\x55\x7d\x5c\x09\x1f\x29\x93\xab\x06\xef\x72\x2d\xb3\xc0\x27\x4a\x29\xa5\x94\x2c\xa9\x7b\xd5\x81\x4b\xce\x92\x56\x7d\xec\xa2\x61\xff\xdf\xe5\xf2\xb2\xa5\xe7\xfd\xdf\x67\xdf\xed\x5a\x95\x52\x4a\x29\xa5\x94\x52\x4a\x29\xa5\x94\x52\x4a\x29\x27\x51\x3c\xbb\x4f\xda\xc6\xef\x82\xe8\xff\x4d\xd1\xef\x29\xa9\xdf\xfe\xe8\xb7\xe1\xe7\xcf\xa4\xde\xf7\x45\xbf\x9b\xa3\xdf\x23\xa9\xf7\xa5\x94\x52\x4a\x29\xa5\x94\x52\x4a\x29\xa5\x94\x52\x4a\x29\xa5\x94\x52\xca\xbb\x27\xaa\x26\x49\xa7\xf1\x7b\xa6\x1f\x00\x38\xa7\xe5\xff\xbf\xb3\x44\x92\x7e\x05\x0f\xc3\x57\x60\x12\xea\x50\x81\x9b\xe0\x32\x18\x81\x55\xb0\x0c\xde\x3e\x5b\x92\x0e\xc3\x73\xf0\x08\x6c\x05\x0f\xd6\xc3\xa7\xe0\x0a\xf8\x24\x0c\xc0\x69\xf0\xda\x59\x92\xf4\x12\x3c\x01\xf7\xc1\x76\xd8\x08\x1b\xc0\x06\x05\x2e\x83\xe5\xb0\x08\x8e\x9c\x29\x49\x3f\x87\x27\xe0\xfb\xb0\x13\x76\xc0\x66\x98\x81\xdb\xe1\x73\x70\x0d\xac\x85\x11\x58\x02\xc7\x16\x4b\xd2\x3f\xe1\x4f\xf0\x08\xdc\x08\xc7\xce\x90\xa4\xe7\xe1\x49\xf8\x26\x6c\x87\xb5\xb0\x1a\x06\x41\x82\xa7\x4f\x97\xa4\x71\x18\x84\xbf\xd1\x41\xbf\x84\x7d\xf0\x00\x4c\xc1\x27\x60\x18\x3e\x0c\x47\x4e\xa5\x6c\xf8\x2e\x7c\x0d\x36\x42\x05\x56\xc3\xf9\x70\x16\xfc\x7d\x11\xfd\x03\x87\xe0\x0b\xb0\x16\x2e\x80\xe3\x0b\x25\xe9\xb7\xf0\x20\xac\x85\x35\x30\x04\xe7\xc2\x39\xf0\x46\x3f\xe3\x00\x7b\xe1\x8b\xf0\x79\xf8\x00\x2c\x84\x7f\xf4\x51\x26\x1c\x80\xc7\xe0\x51\xb8\x1b\xae\x87\xff\x87\xa3\x8c\xe9\x5f\xe1\x45\xf8\x19\xdc\x03\x1b\xe1\x36\x58\xb5\x20\x1c\xf3\xc5\x0b\x5a\xba\xb0\x34\x25\xf2\x1d\xeb\xd4\x81\x4d\x57\x0e\xdc\x3e\x3c\xb0\x7a\x50\x19\x18\x5d\x71\xa1\x7c\xe9\xd2\x99\x54\x32\xd9\x56\x3d\x4f\x38\xe6\x8d\x8e\x65\x0b\xc7\xd3\x85\x9b\x95\xc8\x9b\xb6\x05\xcf\x65\x6b\xec\x2e\x51\xf1\xe4\x95\x89\xf7\xc9\xd4\x33\xf2\x85\x8e\xa8\xfa\x89\x97\x0f\x69\xa2\xaa\x9b\xba\xa7\x5b\xa6\x3b\x64\xe8\xae\xa7\x58\x55\xc5\xf5\x1c\xdd\x1c\x77\xe5\x2d\xf1\x42\x64\xcb\x14\x37\xf8\x99\xd6\x35\x1f\xca\xa6\xf0\xa6\x2c\x67\x7d\x58\x9f\x59\x8f\x95\x9a\xa5\xf9\x55\x9a\x69\xd6\x2d\x2c\x38\x5e\xae\x5c\x13\x35\x77\x4a\xb5\x15\x43\xaf\xe9\x5e\x3c\xf1\x3a\xd9\xac\xd7\xc6\x84\x23\xaf\x6c\xe6\x1b\x4d\x65\xec\x22\x93\x5a\x51\x54\x4d\x73\x84\xeb\xb6\xa9\x58\x33\xed\x96\x64\x0f\xca\x64\x0b\xba\x49\x35\x12\xe3\x50\x55\x0d\x57\x24\x52\xa6\x32\x06\x8f\x66\x3d\xc9\x7e\xe6\x2b\xc4\x60\xa8\x03\xad\x16\x45\x75\xa3\x45\xad\xb6\x99\x75\xc3\xa0\x65\x19\x45\x14\x52\x96\x20\x61\x47\x85\x09\x52\x58\x76\xa0\x17\x99\x05\xc8\x9a\xa3\x4f\x52\x99\xce\xc3\x1b\xd6\x28\x59\x15\xa9\x58\x2d\x64\xc3\x1a\x1f\xf7\x8b\x4b\xea\x96\xa1\x9b\xeb\x93\x63\xa7\x3a\x8e\x3a\xed\x77\x8a\xee\xa1\x4b\xd9\x15\x5a\x2a\xd7\x4d\x7d\x43\x5d\x5c\x1b\x25\xf1\x9c\xba\x88\xeb\x85\xa1\x8e\x09\x23\xcc\xdb\x69\x6a\x38\x8a\xa6\x53\xc9\x78\x4e\xdd\xae\xe4\xa9\xb8\x5e\x53\xc7\x73\xe7\xc1\x84\xe5\x7a\xa6\x5a\xcb\x4d\x27\x36\x7a\x8e\xaa\xf8\xa9\xbb\xaa\x2e\xd9\x51\x0e\xd5\x50\x7a\xda\x8f\xcd\xf2\x93\x2b\xc6\xec\xb4\xb1\xb7\x69\x2d\xa9\x5a\x4e\x4d\xf5\x27\x33\x75\xf4\x6f\xac\xc9\x99\x7a\x92\x31\x1d\x46\x13\xdf\x6c\x56\x5e\xca\x50\xb2\xa8\x89\xf1\xee\x08\x3e\x95\x54\x31\x61\x4e\xea\x8e\x65\xd6\x84\xe9\x75\xd5\xc5\xe6\xa4\xe2\xdf\xb6\xeb\x90\x37\x6c\x83\x9f\xdb\x2f\x25\x9e\x3b\xd1\x2b\xa3\x89\x2e\x98\xd3\x48\x6d\x69\x93\x35\x6b\x8a\x66\xad\xec\xb4\xdd\x99\xb6\x2d\x3d\xec\x82\xd6\x73\xcd\xaa\xa9\xba\x59\x44\x51\x35\xd3\x55\x5c\xa1\x3a\x95\x89\xae\x3a\xc2\xcf\xdf\x65\x46\x31\xa9\x57\x44\xef\xd7\x08\x4d\xd8\xc2\xd4\x5c\xc5\x32\x73\xb5\x22\x73\x0b\x95\x2b\x96\xe9\xd1\x79\xc2\x51\x0a\xf4\xdf\x7b\x45\x11\x2a\x56\xad\xa6\x9a\x5a\x52\x0b\x2a\xe3\x8e\x55\xb7\x15\x5b\x75\x1a\x73\xa4\x83\x22\x54\xd8\xdd\x35\xb6\x80\x9e\x0f\x88\x5f\x30\xbb\xf2\x89\x18\x68\xc3\x9a\xee\x30\xc8\x61\x82\x60\x7d\x68\xe6\x6b\xb3\xcb\xb5\xdb\xdf\x64\x5d\xcb\x50\x6a\xe1\xf8\xaa\x1b\x25\x92\x1b\xff\x46\xc5\xc9\xb1\xa4\xcd\x67\x9d\xac\x92\x60\x5d\x8e\x2a\xd8\x50\x83\x22\x66\x4c\x5c\x0b\xe3\xda\x91\xdd\x19\x93\x96\x51\xaf\xc5\x56\xeb\x8e\x26\x6d\x47\xeb\x24\xb7\xaf\x5a\x5d\x1c\x7d\xd5\x6d\x74\x55\xe3\xdf\x66\xcf\x37\xdb\xb0\x65\x2e\x6d\x89\x0c\xd7\x93\xdd\x98\xa6\x19\x1d\xb5\x66\xb6\x59\xdd\x52\xb3\x9e\x8f\x5f\x43\xcb\x4e\x72\x9b\xa3\xcf\xba\x29\x65\xcf\x6c\x73\x6a\x06\x47\x43\x2e\x1c\x57\x0f\xd7\x62\x29\x73\xfa\xc9\x8e\xd8\x50\xd7\x1d\xa1\x05\x66\x43\x23\x79\x60\x2d\x64\x55\x34\xaa\x64\xe6\xb5\xf8\x30\xc1\x85\xe1\x43\x3f\xd5\x84\xe7\xd9\x97\x0e\x0d\xf9\xef\x06\xc2\xa7\x83\x96\x33\x3e\xa4\x39\x6a\xd5\x1b\x18\x1e\x89\x2e\xda\x2f\x27\xdf\x4c\xe0\xff\xdf\x8a\x0f\x58\xe9\x0f\xfd\xff\xc6\xf9\xfd\x7e\x7c\xfa\xbd\xb0\x03\xa6\xc0\x84\xcf\xc0\xd5\x70\x15\x0c\xc3\x19\xf0\x3a\xbe\xfc\x1f\xe0\x29\xb8\x1f\xbe\x04\x13\xa0\xc2\x9d\x70\x03\x7c\x34\xf2\xf9\x5f\xc7\x8f\xff\x35\x3c\x0e\x0f\xc0\x24\x5c\x0d\x17\xc3\x4a\x38\x1f\xde\xc1\x67\x3f\x0c\x07\xe0\x5b\x70\x37\xdc\x05\x9f\x86\x8f\xc0\xbf\xf0\xdf\xff\x02\x4f\x45\xfe\xfc\x6e\xd8\x05\x02\x6e\x87\x35\xf0\x41\x38\x05\x7e\x8f\x1f\x7f\x10\x1e\x86\xdd\xa0\xc2\x6d\x30\x02\x03\xb0\x12\x3e\x04\xc7\xf1\xf1\x7f\x07\xcf\xc0\x6e\xd8\x01\x02\xd6\xc0\xb9\xb0\x10\xde\xc2\xcf\x3f\x0c\x3f\x85\x87\x60\x13\x98\xa0\xc1\x3a\x58\x0e\xff\xc1\xcf\x3f\x0e\xc7\xe0\x55\x78\x01\x7e\x04\xdb\xa3\x98\xc0\x06\xb0\xe1\x56\xb8\x02\x2e\x80\x3e\x78\x7b\x91\x24\xfd\x1b\xfe\x0c\x2f\xc3\xf3\xf0\x28\xec\x82\x69\xb8\x06\x2e\x81\x15\xd0\x17\xc5\x0c\xf6\xc1\xb7\x61\x0f\x7c\x1d\x76\x82\x07\x57\xc2\x20\x2c\x86\xa3\x8c\xe9\x1f\xa3\xf8\xc1\x0f\x61\x17\xec\x80\x29\xb0\x40\x07\xd1\x1f\x8e\x7d\xa5\xff\x64\x46\x9d\x4a\xf9\x1f\x12\x7f\x99\x0f\x97\xfa\x2d\x1d\xb7\xbd\xa4\x31\x98\x5c\x20\x7d\xeb\x14\xef\x50\x8e\xdb\x9a\xf2\x58\x5d\x37\x52\xb6\x9f\x9d\x1d\x58\x48\x78\x61\x33\xc9\xa2\x43\x07\x38\x11\xae\x48\xa5\x08\xbf\x93\x08\xd5\xa8\xe6\x74\xda\x16\xcd\xb2\x99\xa8\x36\xdb\x02\x66\xb5\x97\xb6\x9f\x52\x16\x54\xf0\x20\x96\x3a\x63\xa3\x69\x35\xbd\x07\x06\x65\x6b\x00\x12\x26\x76\xc1\x30\x52\xca\x3b\x4e\x47\x7a\x8a\x44\x87\x12\x59\xe4\xc1\x15\x69\xa7\x38\x2f\x50\xd4\x3e\x34\xd3\xdc\xef\x93\xfe\x82\x9c\x70\x89\x33\xf6\xf1\xb6\xf1\x80\x4e\x9d\x2b\xa5\x6a\x13\xf7\xe1\xe5\xb4\xbb\xd5\x79\x4c\x0b\xba\x6c\xb1\x1c\x6d\x6b\x93\x6a\x79\xca\x1d\xed\xce\x6e\xeb\x36\x0a\xd5\xc6\xde\x4b\x4d\xf6\xb6\x3e\x67\x7c\xc8\xdb\x4c\xee\xb8\xc2\x8d\x59\x96\x21\x54\xd3\xd7\xb3\x48\x31\x46\x33\xe2\x4b\x89\xec\x09\x65\x4e\x55\xab\x40\xbc\x73\x34\x59\xc5\xae\xd5\x36\x0a\x59\x2a\x96\xed\xb9\xa9\xa8\x46\x7e\x2c\x33\xbb\x6b\x5a\xb5\x6e\x7c\xad\x31\xfb\x3a\x3a\x7a\x91\xcf\x94\x70\x5e\xe6\x6f\xec\xcf\x55\x7f\xb2\xc7\xa7\xcb\x15\xea\x3d\xa3\x5f\x3d\x5a\x61\x67\xc7\xd6\x8b\x9f\x00\x64\x14\x1f\x15\xe1\xd6\xc7\x70\xee\xf2\xba\x29\xaf\xb3\xa2\xf7\x39\x01\xfc\xcc\x50\x68\x22\x63\x3a\x1c\x1a\xbc\x0c\x1d\xa0\x59\xd9\x8a\xc6\xfb\xf3\x87\x38\x63\x62\xea\xb6\x5a\x2b\x57\x8c\x46\x8a\xac\x15\xa3\x11\x99\x48\x46\x08\x3a\xae\x19\x85\x96\x8c\xb1\x69\x4f\x64\x6f\xbc\xad\x34\x15\xbb\x9e\xb3\x37\x77\x6c\x69\xa1\xd0\x97\x23\x5c\xab\xee\xb4\x62\x5f\xcd\xff\x7b\x1b\x02\xe9\x72\x6d\x4b\x99\x8d\x73\x8b\xbf\xfa\xfe\xff\x65\x18\xe9\xa7\x46\xe7\xff\x8d\xfb\xfd\x87\xf0\xed\x0f\x46\xe7\xff\xdb\xe0\x22\x78\x13\x1f\xfe\x00\x3c\x04\x35\xb8\x05\x2e\x82\x0b\x60\x31\xbc\x85\xff\xfe\x02\x3c\x13\x9d\xeb\xef\x85\xcd\x30\x03\x13\x70\x1d\x5c\x02\x67\xc3\x6b\xf8\xf2\x87\xe0\xd9\xc8\xd7\xff\x09\x7c\x19\x34\xb8\xfc\xcc\xf0\xbc\x7e\x1f\xec\x80\x5b\x60\x09\xbc\x8c\xaf\xfe\x03\x78\x10\xb6\xc1\x1d\xb0\x06\x56\xc0\x22\x78\x03\x3f\xfd\x37\xf0\x18\x6c\x85\xeb\x61\x35\x9c\x0e\xaf\xe0\x9f\x3f\x07\x3f\x86\xbd\xb0\x07\xbe\x1a\xf9\xf0\xef\x83\xa3\xf8\xe1\x4f\xc3\x7e\x78\x12\xf6\xc0\x24\xd4\xe0\x5a\xb8\x0a\xae\x84\x01\x58\x78\x6a\x78\x8e\x7f\x10\xf6\xc2\xf6\xc8\x47\xf7\xfd\xf3\x25\xf0\x26\xbe\xf7\xb3\x70\x0f\xdc\x0c\xe7\xc1\xab\xf8\xd7\x2f\xc2\xf7\xe0\x1b\x70\x2f\x18\x70\x13\xbc\x1f\x5e\xe9\xa3\x3c\xf8\x05\x7c\x07\xee\x03\x1b\xae\x86\x7e\x78\x89\xf1\xb9\x1f\x76\x82\x02\x37\xc3\xe5\xb0\x6c\x41\x38\x76\xa7\x2d\x98\x97\x8f\x16\x53\xa7\x6e\x96\x48\xdb\x50\x2b\xa2\x71\x24\x25\x65\x2a\x6a\xb7\xda\x3d\xa5\x9b\x9a\x35\x95\xa5\xbd\x2b\xe3\xc7\x72\x5a\xdd\x51\xfd\xd2\x53\x5b\x95\x5c\x53\x37\x2a\xfe\x0a\x5b\xb3\x53\xb3\x83\xe9\x22\xc6\x59\x24\xd3\xdf\xd3\x84\xa1\x4e\xcf\xe1\x73\xb3\x27\x63\xd8\xcc\x13\xb3\x17\xb1\xf2\x78\xaa\xe3\x29\xb6\x65\xe8\x95\xe9\xc2\xbb\x12\xd9\x70\x36\xd5\x66\x74\xbe\x8d\x05\xd6\x5c\xd7\x52\xd9\x83\x9b\x0c\xc5\x32\xf6\xa2\x85\x41\x51\x6e\xef\x55\xc9\xd7\x85\xaa\xaa\x1b\x75\x47\x28\xc1\xf0\x25\xc6\x28\xda\x9d\xd3\xfa\x63\xd1\x46\x2b\x73\x2f\x2d\xa6\x10\x8d\x0f\xaa\x95\x42\x5a\x31\x7f\x0d\xb4\x55\x47\x35\x0c\x81\x35\x5d\xcb\x57\xf8\xae\x86\xa8\x6e\x6b\xaa\x27\x94\x0c\x5b\x6c\x1e\x57\x15\x1c\x61\xa3\xd2\x6a\xee\x24\x95\x73\xef\xed\xf4\xd4\xb4\x89\x9d\xa6\x45\x9b\x7f\xec\xc9\xfc\x9d\xa2\xf4\x05\xa3\x75\xb1\x9b\x43\x8d\xe0\x40\x74\xb6\x2b\xcc\x4a\xea\xf8\x2e\x5e\x90\x6f\x70\xf9\x4e\xbe\xa6\x67\xdb\x7d\xad\x94\xa1\x43\xa7\x14\xb2\x98\xe3\x47\x57\x3d\x3d\xc5\xac\xbb\xb9\xdf\xee\xb0\x9e\xa5\xdd\x91\xd1\xd4\xff\xd9\xf7\x97\xe6\xe6\x22\xa5\x23\x90\xae\x55\x0d\x74\x66\x42\x75\xb4\xe4\xad\x8e\xcc\x0a\xc7\x8a\x09\x72\x66\x6a\x75\x76\xfa\xe0\x0b\x05\xb6\xaa\x28\x79\x8e\x23\xd6\xd4\xfc\xf6\x57\xa9\x66\x67\x2a\xf0\xed\x8c\xe3\xf8\xe0\x71\x70\x3e\x37\xfb\xaa\xe0\x7c\xfc\x9c\x7a\x6b\xff\x69\x3d\xf4\xbc\xe4\x1a\xd9\xf0\xcc\xe3\x5a\xe6\xd5\xec\x6a\x97\xf7\x34\x5c\xcf\xb2\x95\x71\x07\xb3\x46\xa1\xbe\xba\xa5\x75\xb5\x20\x87\xc5\xb8\xfa\x78\x14\x14\xe8\x34\xd5\x5c\x4f\xd3\x4d\x9c\x39\x61\xe6\xb6\xcb\x9d\xa8\x51\xe8\x26\x91\x74\x34\x3b\x5f\x2e\x74\x45\xa5\xee\xe8\xde\xb4\xef\x2e\xf6\x7c\x3a\x47\x66\x49\x5e\x13\x1d\xa1\x6a\x8a\x65\x1a\xf9\x23\x67\xb3\x3a\xe9\x06\x9a\xa7\xb5\x4b\xda\x4c\xdb\x93\x5b\x5d\xb6\xe5\xa0\x60\xc9\xb7\x27\xe6\x52\x57\xf8\xa5\xe4\x9d\x2e\x5b\xd7\xe4\xb6\x31\x83\xe0\x2e\x65\xa1\xab\x37\xb3\xe2\xf0\x73\x58\xf0\x0a\xde\x0b\x9d\xbd\xd2\xfa\x12\x37\xa2\xa8\x6e\xf6\xba\xd6\x66\x79\x9c\xfb\xba\xdc\xb1\xb0\x20\x5e\x33\xf9\xf1\x8e\xf7\x68\x3b\x66\x1d\xc9\xb9\x82\xdb\x36\xaf\x6a\xe8\xaa\x2b\xf2\x6d\x9f\x36\x37\x98\x13\x65\xe5\xae\xe9\x41\xaa\xdc\x3b\xb2\xbe\x64\xee\x08\xad\xb5\xfb\xbf\x01\x00\x00\xff\xff\x16\x6b\x6e\x5e\x00\x50\x00\x00")

func dataConfig_schema_v3JsonSwoBytes() ([]byte, error) {
	return bindataRead(
		_dataConfig_schema_v3JsonSwo,
		"data/.config_schema_v3.json.swo",
	)
}

func dataConfig_schema_v3JsonSwo() (*asset, error) {
	bytes, err := dataConfig_schema_v3JsonSwoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/.config_schema_v3.json.swo", size: 20480, mode: os.FileMode(420), modTime: time.Unix(1477423928, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataConfig_schema_v3JsonSwp = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x9a\x6d\x6c\x14\xd5\x1a\xc7\x0f\xa5\xf4\xf2\x7a\xe1\x72\x73\x3f\xdc\x17\x6e\xca\xb4\xf7\xc3\x0d\xec\x6e\xef\x6d\x8d\x16\x4c\x14\xa4\x12\x4c\x54\x62\xa2\x01\x4a\x1d\x67\x77\xce\x6e\x87\xce\x5b\xcf\xcc\xf6\x85\xa5\x08\x88\x2f\x31\x86\xc4\x18\xa3\xa6\xd1\x18\x0d\x0a\x11\x25\x51\x24\xc1\xd7\x98\x18\x49\x34\x6a\xa2\x21\x7e\x30\x46\x12\xc5\xaa\x1f\xd4\x04\x23\xa0\x28\xfe\x67\x67\x76\x77\x66\x3a\x3b\xbb\x5d\x6b\x21\xf1\x3c\xc9\x8f\x65\x67\x9e\x73\xe6\x39\xe7\x3c\x73\xce\xf3\x3c\xdd\x74\xc7\x2d\x1b\xae\x6f\xbd\x3c\xd9\x45\x20\x4b\x08\x99\x7f\x68\xf1\xa6\xbd\xb7\x1d\x9d\xf5\x08\x23\x44\x96\x74\x85\xaa\xa4\xa6\x18\xb2\x29\xd5\xa1\x46\x76\xba\x1d\xa6\x4c\xca\xac\x54\xc6\x90\x29\xfe\xd1\x4c\xc3\xa2\x89\xac\xa2\xd2\x94\x95\xe9\xa7\x9a\x94\x92\x25\x5b\xc2\x0d\x3d\xab\xe4\x44\xf7\x92\x38\xd4\x99\xdc\x66\x19\x7a\x1d\x8f\xe0\xc2\x85\x4b\x55\x59\x44\x3a\xff\xff\xbf\x0e\xe7\x7f\x6d\xc2\xf2\xd6\xbf\x2e\xbd\xf9\x62\x1b\xc4\x85\x0b\x17\x2e\x5c\xb8\x70\xe1\xc2\x85\x0b\x97\x19\x14\xdb\x9c\x4d\x6e\xc7\x67\x93\xf7\x7d\xbb\xf7\x39\x2b\xf4\x79\xc1\x93\x52\x9e\x5f\x08\xdd\x9f\xed\x7d\xee\xf0\x3e\x27\x42\xf7\xb9\x70\xe1\xc2\x85\x0b\x17\x2e\x5c\xb8\x70\xe1\xc2\x85\x0b\x17\x2e\x17\x4f\x24\x99\x90\xb9\xf8\x5c\xe8\x14\x00\x96\x54\xf2\xff\x5f\x16\x13\xf2\x16\xd8\x0f\xee\x07\x43\x20\x0f\x32\xe0\x26\xb0\x1a\x74\x81\x4e\xb0\x1c\x9c\xff\x33\x21\x9f\x83\xe3\xe0\x00\xd8\x09\x6c\x30\x00\xd6\x83\xab\xc1\x55\x20\x01\xe6\x82\xef\x16\x11\xf2\x21\x38\x02\xee\x05\xbb\xc0\x08\x18\x04\x26\x10\xc1\x6a\xd0\x06\x5a\xc0\xc4\x42\x42\x5e\x03\x47\xc0\x33\xe0\x6e\xb0\x1b\xec\x00\x05\xb0\x05\x6c\x06\xd7\x82\x1e\xd0\x05\x16\x83\x33\x0b\x08\xf9\x01\x9c\x04\x07\xc0\x46\x70\x66\x3e\x21\x6f\x83\x17\xc1\xe3\x60\x17\xe8\x01\xdd\x20\x09\x08\x78\x7d\x1e\x21\x39\x90\x04\xdf\x60\x82\xde\x04\x87\xc0\x43\x60\x18\x5c\x09\x3a\xc0\x7f\xc0\xc4\x9f\xd0\x37\x78\x1a\x3c\x00\x46\x40\x06\x74\x83\x7f\x80\x45\xe0\xdb\x16\xcc\x0f\xf8\x0c\xec\x05\x3d\x60\x19\x38\x3b\x87\x90\x13\x60\x1c\xf4\x80\x75\x20\x05\xfe\x02\x96\x80\xef\x9b\xb1\x0e\xe0\x20\xb8\x13\x6c\x05\x7f\x07\x73\xc0\xe9\xd9\xe8\x13\xbc\x04\x0e\x83\xe7\xc0\x1d\xe0\x06\xf0\x37\xf0\x15\xd6\xf4\x4b\xf0\x01\x78\x15\xdc\x05\x46\xc0\x26\xd0\xd9\xe4\xae\xf9\x82\xa6\x8a\x2f\xb4\x86\x44\xb8\xb5\x57\x4a\x6c\x5f\x93\xd8\xd2\x91\xe8\x4e\x8a\x89\xbe\x15\xed\xc2\xaa\xd6\x42\x48\x4d\x30\x25\xdb\xa6\x4c\xdf\xc8\x0c\x93\x32\x5b\xa1\x56\x94\x92\x3d\x6a\x52\x5c\x17\x8c\xf4\x36\x9a\xb1\x85\x95\x81\xfb\x41\xed\x82\xd0\xce\x68\xd6\x51\x6e\x4b\xc9\x34\xab\xe8\x8a\xad\x18\xba\x95\x52\x15\xcb\x16\x8d\xac\x68\xd9\x4c\xd1\x73\x96\x30\xe6\xef\x44\x30\x74\x7a\xa3\xd3\xa8\xb7\x7c\x51\xd0\xa9\x3d\x6c\xb0\x01\xd7\x9e\x49\x97\x45\xcd\x90\x1d\x93\x0a\x65\xdb\xdc\x8e\xfd\xfd\x0a\x1a\xd5\xac\x61\xc9\x14\x55\x45\x53\x6c\xbf\x72\xaf\xa0\xe7\xb5\x34\x65\xc2\xca\x72\xbb\xbe\x50\xc3\x06\x1a\x49\x19\x51\x92\x65\x46\x2d\xab\x8a\x61\x65\xdd\xb1\xe0\x0c\x0a\x68\x56\x9c\x26\x49\x0d\xac\x43\x56\x52\x2d\x1a\xd0\x0c\x35\x2c\x5e\x9a\x74\x25\xfa\x9a\xe3\x10\x49\xd7\x07\x2a\x23\xf2\x6c\xc3\x88\x2a\x63\xd3\xf3\xaa\x8a\x91\x45\x74\x51\x97\xb3\x14\x15\x63\x1d\xa6\xa8\x61\x98\x45\xbf\x88\xec\x40\x90\x99\x32\x04\x63\xe2\x97\xd7\xb5\x28\x68\x0a\xa9\xcf\x0a\x41\x35\x72\x39\xa7\xbb\xa0\x6f\xa9\x8a\x3e\x10\x5c\x3b\x89\x31\x69\xd4\x99\x14\xc5\x86\x2f\x45\x1b\xd4\x2a\xe4\x75\x65\x30\x4f\x37\x78\x2a\x36\xcb\x53\xbf\x5f\xa8\x52\x9a\xaa\x6e\xdb\xb8\x57\x83\x89\xb2\x02\x23\xfd\x2d\x15\x33\x53\xcb\xc5\x15\x4d\xca\xd5\x7c\x0f\xfa\x0d\xcb\xd6\x25\xad\xa6\x1e\x1d\xb1\x99\x24\x3a\xda\x0d\x99\x8b\xe6\x70\x0e\x49\x15\xa7\x75\x1e\xcb\xfd\x07\x77\x8c\xc9\xba\xbe\xbb\x61\x2f\xc9\x1a\x4c\x93\x9c\x97\x19\x36\x3a\xbf\x58\x13\x22\xfd\x24\xe2\x75\xe8\x0b\x3c\xb3\x6c\x3c\x89\x70\x32\x6f\x88\xfe\xe9\x28\x3e\x2a\xe8\x62\x54\x1f\x52\x98\xa1\x6b\x54\xb7\x1b\x9a\x62\x7d\x48\x74\x7e\x6d\x17\xd3\xd6\x1d\x83\xd3\xda\xe9\xc5\xdf\x3a\x30\x2b\x7d\x81\x29\x98\xd2\x4a\x8d\x55\x69\x1a\xf5\x8a\x46\xed\xec\x18\x3b\x1b\x35\x0d\xc5\x9d\x82\xca\x75\xd9\xd0\x24\x45\xaf\xc7\x51\x65\xdd\x12\x2d\x2a\xb1\x4c\x7f\x43\x13\xe1\xb4\x6f\xb0\x21\x1d\x52\x32\x74\xfa\xf7\x08\x99\x9a\x54\x97\x2d\xd1\xd0\x6b\x7a\x45\xe4\x11\x2a\x64\x0c\xdd\xc6\xe4\x51\x26\xd6\x31\x7f\x97\x8a\x23\x64\x0c\x4d\x93\x74\x39\xe8\x05\x99\x1c\x33\xf2\xa6\x68\x4a\xac\xf4\x8e\xc4\x38\x42\x06\xa7\xbb\x8c\x23\x60\xda\x17\xc4\xe9\x18\xa7\xf2\xef\xb1\xd0\xaa\x31\x1a\xb3\xc8\xae\x42\x71\x7f\x28\xb7\xab\x72\xca\x55\x3b\xdf\x04\x45\x8e\x70\x6a\xca\x1c\xd7\xf5\x94\x84\xd2\x57\xaf\x3b\xc1\xa7\x5a\xbe\x16\x17\x95\x14\xf7\x65\xcf\xc0\x92\x1b\xd4\x13\xc6\xf8\xbd\xd0\xef\x1d\xd1\x93\x31\x64\xa8\x79\xcd\xb7\x5b\xc7\x86\xb4\xb1\xd1\x49\xcd\xb9\xaa\x4c\xb1\xf7\x54\xab\x34\x55\xa5\xaf\xe5\x99\x2f\x8f\x61\x6c\x2a\x63\xf1\x02\xd7\x99\x1e\x4c\x39\x8c\xf6\x46\x33\x39\xac\xae\xb8\xd9\xb4\xaf\x5f\xc9\xcb\x66\x78\xcc\xde\x63\xad\x90\xb3\x47\x8e\x39\xf4\x06\x7b\x4b\x4e\x99\xa5\xb8\x7b\x31\x89\x7c\xfd\x04\x46\x07\xf3\x0a\xa3\x72\x31\x6c\x28\xa9\x17\xa3\x85\x28\x43\x3d\x23\x23\x7f\x16\xef\x2a\xb4\xbb\x17\x1d\xad\x7e\xdb\x36\x57\xa5\x52\xce\xbd\x84\x7b\x35\x69\xb0\x5c\x4a\x66\x52\xd6\x4e\x74\x74\x79\x3f\xb4\x6f\x43\xbb\x42\x31\xff\x77\xf2\x41\xb9\xd9\xcd\xff\x4b\x7f\xbf\x3f\x8a\x9c\xfe\x20\xd8\x0d\x86\x81\x0e\xae\x03\xd7\x80\xb5\xa0\x03\x2c\x00\xa7\x91\xcb\x7f\x02\x8e\x81\x07\xc1\x7d\x40\x01\x69\x20\x81\x8d\x60\x05\x98\x07\x4e\x23\x8f\x7f\xd7\xcb\xf9\x1f\x06\xc3\x60\x1d\xb8\x02\x24\xc0\x3f\xc1\x05\xe4\xec\xa7\xc0\xcb\xe0\x49\xb0\x17\x0c\x80\x0d\xe0\xbf\xe0\x47\xe4\xef\x13\xe0\x18\x38\x08\xc6\xc1\x3e\x90\x05\xbd\xa0\x07\xfc\x0b\x34\x81\x8f\x91\xc7\xbf\x07\x9e\x02\xe3\x20\x0d\x36\x83\xcb\xbc\x3c\x3f\x01\x96\x81\x73\xc8\xf1\x3f\x02\x6f\x80\x71\xb0\x07\x64\x41\x0f\x58\x0a\x5a\xc0\x4f\xc8\xf3\x4f\x81\x57\xc0\x13\xa0\x00\x0c\x40\xc1\x56\xd0\xee\x14\x51\xc0\x39\xe4\xfa\x67\xc1\x49\xf0\x3e\x78\x01\xec\x06\xa3\x80\x81\x41\xb0\x09\xac\x01\xff\x06\xcd\xe0\xe7\x16\x42\xce\x83\x2f\xc0\x09\xf0\x0e\x38\x0c\xf6\x81\xed\x60\x3d\xe8\x06\x2b\x41\x33\x38\x37\x87\x90\x67\xc1\x7e\xf0\x18\x78\x14\xdc\x03\xf2\x60\xad\x57\x3f\x58\x08\xbe\xc6\x9a\x7e\x0a\x8e\x83\xe7\xc1\x3e\xb0\x07\x8c\x00\x13\x6c\x03\xd9\x66\x77\xed\x33\xcd\x33\x58\x74\xe2\xf2\x87\x12\x67\x9f\x77\xf7\xfa\xb1\xd8\x73\x2f\x18\x0d\x06\x77\x48\x27\x3c\x45\x7a\x28\xf8\x83\x4d\x21\x9d\x57\xd4\x50\xf0\x67\x46\x57\x16\x02\x69\x58\x21\xd8\xb5\x9b\x01\x07\xea\x15\x21\x0d\xf7\x39\x81\x5a\x8d\xa4\x8f\x86\x83\xd1\xa8\xa0\x09\x66\xe3\x5c\x40\x5c\x6d\x87\x03\xa8\x50\x08\x55\xbc\xe0\xd3\x8e\x38\x69\x2a\x43\x9f\x86\x88\xb2\xb2\x00\x81\x18\xbb\xce\x3a\x52\x28\x3d\x0e\x97\x7a\xea\x29\x0f\x05\x9a\x08\xc9\x15\xe1\xac\xb8\x56\xa5\xa8\x7a\x6d\xa6\x7c\xe0\x07\x13\x06\x21\x90\x13\x47\x1c\xe4\x55\x0b\x02\x71\x93\x4b\x42\xd6\xf8\x93\x78\x21\x9c\x6f\xc5\xaf\x69\x9d\x39\x9b\xaf\x45\x55\x6b\x42\x23\x0f\xe5\xa3\x8d\x05\x6e\x8d\x96\xa1\xaa\x04\x7c\xa1\x97\xbd\x6a\xd2\xe9\x5f\xf2\x2a\x2f\xb7\xdf\xe1\xd2\x86\xa1\x52\x49\x77\xfc\xcc\x73\x8c\xbe\x88\x02\x53\xa0\x79\xc0\x99\x43\x66\xd5\x51\xf0\xec\x0b\x9a\xd8\xb0\xdb\x7a\x35\x4b\xd1\x30\x6d\x2b\x54\xd6\xa8\x5d\xcc\x8c\x9e\x9a\x8a\xd5\xa5\xa7\x95\xde\xbe\xd8\x4c\xcf\x4b\x9a\x02\xd9\xcb\x6f\x8f\xf6\xa7\xea\x3f\xd1\xeb\xd3\xe0\x0e\x75\xc9\xf8\xd7\x34\xed\xb0\x93\x8b\xeb\xf5\xff\x09\x20\xa2\x7b\xaf\x0b\x2b\x9f\x46\x76\x57\x6b\x9a\x6a\x4d\x96\x77\xbf\x46\x05\x3f\xb2\x16\x1a\x68\x18\xae\x87\x16\x6f\xba\x19\xd0\xa4\x66\xf5\x16\xfc\x6b\x2f\x71\xc4\x8b\xa9\x98\x92\xc6\x77\x8c\x92\x46\xd4\x8e\x51\x2a\x4d\x04\x4b\x04\xb1\x7b\x46\x5d\x5b\x46\x7a\xd4\xa6\xd1\x07\x6f\x45\x27\x63\xe6\x6b\x9c\xcd\xb1\x23\xad\x56\x0e\xf0\x0f\x8f\x51\xcb\xc8\xb3\x4a\xed\xab\xfc\x7d\x7a\x4b\x20\x0d\x6e\x6d\xa1\xa8\x71\x6a\xf5\xd7\x5f\x03\x00\x00\xff\xff\x5d\xec\x62\x72\x00\x40\x00\x00")

func dataConfig_schema_v3JsonSwpBytes() ([]byte, error) {
	return bindataRead(
		_dataConfig_schema_v3JsonSwp,
		"data/.config_schema_v3.json.swp",
	)
}

func dataConfig_schema_v3JsonSwp() (*asset, error) {
	bytes, err := dataConfig_schema_v3JsonSwpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/.config_schema_v3.json.swp", size: 16384, mode: os.FileMode(420), modTime: time.Unix(1477422692, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataConfig_schema_v3Json = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5a\xcf\x92\xdb\x28\x13\xbf\xfb\x29\xa6\x94\xdc\x62\x8f\x53\xf5\xa5\xbe\xaa\x9d\xdb\x1e\xf7\xb4\x7b\x5e\x97\x42\x61\x09\xdb\x24\x12\x10\x40\x9e\x38\x29\xbf\xfb\x82\x90\x64\x90\x40\x48\xb6\x76\x67\x0f\x9b\x43\x6a\x0a\xba\x9b\xfe\xf3\xa3\xbb\x69\xf9\xe7\xea\xe9\x29\x79\x2f\xb2\x13\x2a\x61\xf2\xf2\x94\x9c\xa4\x64\x2f\xdb\xed\x17\x41\xc9\xc6\xac\x3e\x53\x7e\xdc\xe6\x1c\x1e\xe4\xe6\xe3\xa7\xad\x59\x7b\x97\xac\x35\x1f\xce\x35\x4b\x46\xc9\x01\x1f\x81\xd9\x01\xe7\xff\x3d\x6b\x66\x43\x20\x2f\x0c\x69\x12\xba\xff\x82\x32\x69\xd6\x38\xfa\x56\x61\x8e\x34\xeb\x2e\x39\x23\x2e\xb0\xa2\x4e\xd7\x2b\xbd\xc7\x38\x65\x88\x4b\x8c\x84\xda\xd5\xaa\xa9\xb5\x96\xa4\x5d\xb0\xc4\x0a\xc9\x31\x39\x26\xf5\xf2\xb5\x96\xa0\x36\x05\xe2\x67\x9c\x59\x12\x3a\x45\xdf\x6d\x6f\xf2\xb7\x1d\xd9\xba\x2f\xd5\x52\xb6\x5e\x67\x50\x4a\xc4\xc9\x1f\x43\xdd\xea\xed\xcf\x3b\xb8\xf9\xf1\xeb\xe6\xcf\x8f\x9b\x5f\x9e\xc1\x26\xfd\xf0\xde\xd9\xd6\xde\xe5\xe8\x60\x8e\xcf\xd1\x01\x13\x2c\x95\x35\xdd\xf9\x49\x47\x79\x6d\xfe\xba\x76\x07\xc3\x3c\xaf\x89\x61\xe1\x9c\x7d\x80\x85\x40\xae\xcd\x04\xc9\x57\xca\xbf\xc6\x6c\xee\xc8\xde\xc8\xe6\xe6\x7c\x8f\xcd\xae\x39\x67\x5a\x54\x65\x34\x82\x2d\xd5\x1b\x19\x63\x8e\x7f\x2c\x7e\xab\xd6\xe8\x51\x5a\x43\x61\x9d\x5d\x2b\xe8\xa0\xdd\xe7\x2a\x1f\xda\xc2\xbe\xea\x9c\x15\xf0\x52\x8e\x58\x41\x2f\x7a\x2d\xe0\x0f\x43\x50\x22\x22\x93\xce\x05\x8a\x2f\x83\x0c\x28\xe3\x6a\xc6\xf6\x50\xc8\x39\xbc\x24\x6b\xa5\xaa\x44\xa5\x70\xb6\x9a\x2b\x7d\x55\x9b\x15\xc1\xdf\x2a\xf4\x5b\x43\x22\x79\x85\xfa\x72\x73\xa5\xeb\xf2\x82\x8f\x9c\x56\x0c\x30\xc8\xb5\x29\x5e\x11\x16\x31\x2d\x4b\x48\xf2\x3e\x62\x28\x41\xbf\x6b\x17\xed\xac\xc5\xa7\x51\x49\xee\xfe\x04\x3b\xae\x16\x6b\x7a\x03\xa1\xa3\x1b\x91\x10\x13\xc4\x01\x81\x25\x8a\x59\xa2\xc2\x87\x48\x2e\x80\xc9\xb4\x81\x18\x17\x58\x48\x40\x0f\xc0\xf0\x8b\x9e\x80\x2e\xed\x2e\x1a\x8f\xdc\xe0\x3d\x94\x46\x6b\x31\x80\x72\xa0\x75\x4b\x7a\x8c\x40\x20\xc8\xb3\xd3\x9d\xfc\xb4\x54\xee\x9b\xe2\x3b\x05\x14\x7e\x61\x14\x1b\xbc\xfc\xeb\x80\x80\xc8\x19\x1c\x70\x81\xee\x73\x83\xe2\xc6\x9c\x92\xb2\xbd\x0d\x63\xc8\xe0\x20\xc7\x59\xcd\x6e\xf1\x7f\x67\x54\xa0\xbe\x63\x7a\x06\xda\x5b\x9d\xa9\x8e\x4f\x5a\x8e\x5d\x6b\xb8\x72\x0a\xa9\xca\x3d\xe2\xba\x79\x70\x28\x0f\x94\x97\x50\x2b\xdb\x9e\x6d\x6d\x3b\x9e\xf6\x20\xcf\x76\xa0\x6d\x83\x2e\x20\xb0\x50\xde\x21\x5f\x97\x87\xb8\x12\xcf\x21\x38\x51\x21\xc7\xa0\xde\x77\x71\xc7\xae\x19\xa7\xc0\x14\x97\xf0\x18\x27\x62\x59\x8c\xa4\x80\x7b\x54\xdc\xa5\xe9\xa2\xee\xb3\xc4\xd2\xe3\x51\x93\x86\x30\x33\xa8\x72\xcd\x76\xa0\xd6\x75\xfb\x39\xc7\xaa\xfb\x8c\xb8\xa3\xa3\xa6\xec\x56\x9c\xfb\x9b\xe1\xee\xc4\x55\x68\xb4\x53\x71\x48\x3f\x3f\x9b\x46\x65\xe4\x5e\xd4\x7f\x15\x45\x92\x5e\x3d\x22\x86\x6b\xee\x4a\xcf\xc2\x78\x33\xd3\x30\xdd\xa2\x52\xc2\x4c\x57\x7e\x8e\x44\x20\xae\x16\x29\x2a\xd5\xd5\x2a\xb1\x74\x2d\xba\xd9\xd1\xf0\xa4\x3d\x26\xf1\xaa\xba\x80\xd9\x8c\x4d\x1b\x0a\x4a\x9a\x87\xae\xc3\x80\xb8\x1f\x89\x60\x66\x9f\x5d\x38\x6b\xb6\x5e\x30\x62\x60\x99\x08\x94\x68\x6b\x1b\xb1\x26\xa4\xde\x54\x35\x6f\xea\xc6\x01\x5d\xd3\xc1\x02\x43\x81\xe2\xa9\x25\xe8\x48\x47\x1a\x66\xe7\x4f\x13\x11\xe8\xe3\xfd\xff\x28\x6f\x80\x35\x28\x73\xda\x05\x9a\x20\xea\xa6\x4a\x7d\xb9\x7d\x8a\xa4\x91\xbb\xed\x91\x3d\x5d\xbd\x78\xbf\xc1\x70\x1e\xce\x4c\x75\x3e\xb2\x2f\x18\xa3\x5c\x0e\x6e\xd7\x3f\xd3\x1e\x98\xa3\x1f\xee\x0e\x98\x2a\x13\xaa\xbd\x3a\x22\xf7\x95\xb3\xa7\xb4\x40\x90\x38\x89\x8e\x23\x98\xab\x16\xbb\xb8\x4c\xa0\x14\x12\xf2\xe8\x03\x44\xa0\xac\xe2\x58\x5e\x80\xaa\x3e\x8b\xf7\x25\xe2\x54\x02\x81\x7f\xa0\x39\xc9\x55\xc8\x1c\x13\xa5\x0d\x22\x51\x13\x85\xa4\x4c\xc9\x3f\x2a\xcc\x45\xcd\xd4\xa4\x47\x0e\x33\x04\x14\x36\x31\xcd\x7d\x0c\x6b\x3b\xb6\x79\xc5\xa1\xc6\xb3\x23\x46\x96\xec\x70\xe7\x6b\x42\xca\x78\xcc\xaa\xba\x14\x05\xc1\xec\xc9\x92\x13\x12\xb9\x49\xe2\xfe\xdc\x3d\x92\xb7\x6f\x9a\xaa\x67\x89\xc2\x26\xf7\xa5\xbb\x91\x46\x65\xbc\x4f\x99\xd0\xa0\x9c\x20\x77\xa3\x34\xa2\x47\xcd\x20\xe8\x41\xfa\x19\x7c\xed\x8b\x57\x2f\x67\xb6\x58\xcb\x5b\x37\x8a\xa4\x5e\xfa\x59\x39\xb9\xaf\x46\x1a\x4c\x8b\x57\x6f\x5a\xac\x44\xb4\x97\xb4\x27\x5f\x8b\xde\x64\x23\x17\x4c\x6a\x68\x13\xdd\xef\xe8\x6b\x90\x63\x3f\xe9\xaa\xb3\xad\x1d\x1a\x99\xe1\x01\x22\xd9\x60\x6c\xd4\x6f\xd2\x76\x56\xaf\x97\xae\x7a\x4e\x9a\x31\xf1\xb4\xa6\x4d\x91\xd1\x97\x45\xd9\x9f\x7e\xed\x3a\x90\xb7\x95\xe9\x36\x33\x0c\x8c\xc1\x46\x7a\xc6\x8e\x86\xab\x13\x71\x06\x45\x0c\xfd\x0f\x3c\xa4\x2a\x96\x43\x89\x80\x19\xbd\xcf\xca\x37\x23\x89\x86\x41\x0e\x8b\x02\xa9\x43\xcb\x29\x17\x57\xc5\xa0\x80\x97\xbb\x12\x71\xcd\x7e\x80\xb8\xa8\x38\x02\x30\x93\xcd\x7c\x3f\xd2\xa0\x29\xe7\x2b\xc7\x50\x2f\x28\xa7\x1d\x59\xc2\xef\xa0\x3d\xb6\x26\x71\x44\x35\x85\xed\x1a\x6c\x08\xa6\xbe\x81\x2c\x24\x08\x5a\xf1\x6c\xe0\xec\xbb\x43\x74\x2b\x30\x01\xc4\xb4\x27\x0e\x4c\x57\x1b\x88\x9f\x61\xf7\x44\x8d\xf2\x47\xf3\x59\xd3\xa1\x00\x46\x15\xda\x2f\x4b\x59\xa8\x20\x6d\x9c\x3c\x05\x10\x0f\x22\x50\xc3\x41\xd7\xdf\x92\xc9\xe8\x65\xad\x19\x5e\x31\xc9\xe9\xeb\x8c\x03\x97\x83\x12\x2b\x54\xf3\xd3\xcb\x77\x8f\x3a\x5a\xe9\x0e\x95\xa9\xb3\xcb\xcd\xa3\x66\x3d\x90\xf7\x3b\x7c\x46\xb2\x7e\x47\x17\xff\x3a\x14\xc8\xf4\x19\xab\xa2\x73\x8b\xfd\x45\xa2\xd1\xe7\xa1\xc7\xc4\x89\xdf\xf2\x62\x06\xb6\x64\x0b\xd4\xb4\x69\x5d\x81\xa1\xd2\xef\x8c\xc5\x1b\xdc\xf8\x28\x2b\x8d\xa7\x23\xcc\x60\xb9\xd4\xdd\x98\x3c\xf8\x4b\xbc\x15\xd8\x39\x7b\xf8\x84\x35\xea\x7a\x9f\xb1\x31\xad\xe3\xba\x37\x14\xa2\xda\x2b\x84\x8c\x41\xf3\xf6\xcf\x37\xcc\x9c\xd1\x19\x5f\xc3\x7d\xf0\x63\x29\xaf\x9d\xba\x07\xa2\xba\xeb\xde\x60\xeb\xce\x57\xe9\xe4\x10\x07\x07\xe6\xcb\xe9\x3f\xb3\xbd\x7b\x20\x2b\x36\xdf\xa2\x23\x29\xa3\xa1\xfa\x2f\x63\x34\x52\xde\x1e\x5f\x23\x25\xf1\xce\xb7\xc1\x0c\xd0\xf4\x86\x1d\x16\x78\x86\x53\x85\xb1\x38\x4f\x1e\xd5\x36\x1c\xa9\xab\x46\x9f\xec\x65\xf8\x3b\x1f\x37\x85\x8e\xbd\x83\x5b\x92\xc0\xe8\xae\x77\x68\xe3\xbc\x71\xcb\x17\x84\xed\xf3\x87\x91\x42\x31\xf6\x01\xe7\x6f\xca\xb0\x0b\xcc\x18\xfc\x31\xed\xf5\x96\xad\x77\x87\x3f\x56\x09\x64\x2a\x8b\x7f\xf0\xd3\x15\x6d\x27\xb9\x0c\xa6\x5e\x3f\xdd\xe1\xcf\xbe\xc2\x45\xee\x4c\x26\x07\x24\xe6\x73\xa8\x95\x27\x52\xbb\xdd\x0e\x85\xb1\x91\x3c\x88\xad\x23\x5a\xff\xf4\x42\x65\x17\x27\x70\x9e\xcb\xbe\x6a\xff\xbf\xae\xae\xab\xbf\x02\x00\x00\xff\xff\x4e\x83\x1c\xd5\x8a\x27\x00\x00")

func dataConfig_schema_v3JsonBytes() ([]byte, error) {
	return bindataRead(
		_dataConfig_schema_v3Json,
		"data/config_schema_v3.json",
	)
}

func dataConfig_schema_v3Json() (*asset, error) {
	bytes, err := dataConfig_schema_v3JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/config_schema_v3.json", size: 10122, mode: os.FileMode(420), modTime: time.Unix(1477423905, 0)}
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
	"data/.config_schema_v3.json.swo": dataConfig_schema_v3JsonSwo,
	"data/.config_schema_v3.json.swp": dataConfig_schema_v3JsonSwp,
	"data/config_schema_v3.json":      dataConfig_schema_v3Json,
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
	"data": {nil, map[string]*bintree{
		".config_schema_v3.json.swo": {dataConfig_schema_v3JsonSwo, map[string]*bintree{}},
		".config_schema_v3.json.swp": {dataConfig_schema_v3JsonSwp, map[string]*bintree{}},
		"config_schema_v3.json":      {dataConfig_schema_v3Json, map[string]*bintree{}},
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
