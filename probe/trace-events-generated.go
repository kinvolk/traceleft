// Code generated by go-bindata.
// sources:
// ../dist/trace_events.bpf
// DO NOT EDIT!

package probe

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

var _trace_eventsBpf = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x9c\x7b\x6c\x54\xc7\x15\xc6\x67\xfd\x5c\x9b\x97\x09\x2f\xf3\x70\x30\x90\x80\xcd\xc3\x36\x78\x0d\x76\x08\x60\x02\xc6\x4e\x4a\x52\x07\x68\xe3\xa6\xb4\xcb\x7a\xb9\x8e\x8d\x1f\xbb\xda\x5d\x61\x93\xd2\xc6\xa9\x92\x14\x45\x51\xe5\x2a\x49\x05\x52\x14\x45\x6d\x95\xf2\x47\xa9\xac\x34\x6a\xa8\xd2\x16\xaa\x4a\x2d\x4a\x1f\xb2\x50\x95\x44\x8d\xd4\xd2\x8a\x4a\xa8\xaa\x54\xb7\x4d\x15\xfe\x40\xa5\x9a\xd9\xbb\xec\xf5\x39\xdf\x59\x8f\x54\x55\xad\x94\x19\x69\x31\xf3\x7d\x73\xe6\x77\xef\xb9\xf3\xb8\x77\xaf\xe1\xa9\x8e\x83\x07\x8a\x42\x21\x95\x2b\x21\xf5\x91\xca\xd7\xf2\xa5\xf5\x48\xfe\xef\xed\xfe\x9f\x09\x15\x52\x97\x96\x66\xb5\x67\x95\x52\xf3\x95\x52\xa3\x4a\xa9\x5a\xa5\xd4\xa5\x58\x56\xaf\x2e\x9a\xd9\xcf\xa5\xe2\x7c\xfb\xb9\x05\xda\xbd\x45\xda\xbd\xe5\xeb\x2f\xe7\xfa\x71\x5c\xc7\x75\x5c\xc7\x75\x5c\xc7\x75\x5c\xc7\x75\x5c\xc7\x75\x5c\xc7\x75\x5c\xc7\x75\xdc\xff\x1a\xb7\x24\xf0\xc9\xfe\xa1\x94\xee\x22\x93\x8a\xc5\xbd\x21\xaf\x2f\xa3\x3e\x9e\xa5\x38\x98\x97\xf1\xff\xf5\xd1\xfc\xff\x14\x97\x17\x5c\x5c\x5e\x70\x71\x79\xc1\xc5\xe5\x05\x17\x97\x17\x5c\x5c\x5e\x70\x71\x79\xc1\xc5\xe5\x05\x17\x97\x17\x5c\x5c\x5e\x70\x71\x79\xc1\xc5\xe5\x05\x17\x97\x17\x5c\x5c\x5e\x70\x71\x79\xc1\xc5\xe5\x05\x17\x97\x17\x5c\x5c\x5e\x70\x71\x79\xc1\xc5\xe5\x05\x17\x97\x17\x5c\x5c\x5e\x70\x71\x79\xc1\xc5\xe5\x05\x17\x97\x17\x5c\x5c\x5e\x70\x71\x79\xc1\xc5\xe5\x05\x17\x97\x17\x5c\x5c\x5e\x70\x71\x79\xc1\xc5\xe5\x05\x17\x97\x17\x5c\x3a\xbb\x0f\xaa\x7f\xdd\xbe\x7d\xbb\xca\xaf\x87\x9e\x3c\xa4\xc2\xa7\xe7\x84\x56\x2a\xa5\xb4\x56\xed\xeb\x63\xfe\x4f\xad\x6d\x06\x5a\x3b\xd0\x8e\x00\xad\x1f\x68\xa7\x81\x36\x01\xb4\x6f\x01\xed\x22\xd0\x7e\x09\xb4\x6b\x40\xfb\x10\x68\xe1\x10\xd7\x56\x01\xad\x09\x68\xfb\x81\xd6\x03\xb4\x21\xa0\x8d\x03\xed\x25\xa0\x9d\x07\xda\x0f\x81\x36\x05\xb4\xeb\x40\xbb\x09\xb4\xb9\x45\x5c\xab\x05\x5a\x04\x68\x5d\x40\x3b\x0a\xb4\x24\xd0\x9e\x01\xda\x59\xa2\x2d\x56\x4a\x7d\x87\x68\xb3\x95\x62\x55\x2c\xe8\xa5\x82\x5e\x2e\xe8\x15\x82\x3e\x47\xd0\xe7\x09\xfa\x02\x41\x5f\x28\xe8\x8b\x04\x7d\x89\xa0\x2f\x13\xf4\xe5\x82\xbe\x52\xd0\x6b\x04\x7d\xb5\xa0\xaf\x11\xf4\x75\x82\x7e\xaf\xa0\x6f\x10\xf4\x7a\x41\xdf\x24\xe8\x5b\x04\xbd\x51\xd0\xb7\x0a\x7a\xb3\xa0\xb7\x08\xfa\x0e\x41\x6f\x13\xf4\x9d\x82\xbe\x4b\xd0\xf7\x08\xfa\x5e\x41\xdf\x27\xe8\x1d\x4c\x1b\xad\xd0\xf3\xee\x04\xd3\x5f\x2c\xd7\xfa\x20\xd3\x9f\x36\xf3\xb4\x93\xe9\x8f\x95\x69\xfd\x28\xd3\x63\x45\x5a\xff\x1c\xd3\xbf\x59\xaa\xf5\xcf\x33\xfd\x6b\x21\xad\x47\x99\xbe\xde\xf4\xff\x10\xd3\xeb\x4c\xff\x9f\x60\xfa\x6e\xd3\xbe\x87\xe9\xfb\x4d\xfb\xcf\xf0\xe3\x37\xc7\xf3\x38\xd3\xdb\xcc\xf1\x7c\x96\xe9\xcf\x9b\xf6\x47\x98\x9e\x31\xed\x3f\xc5\xf4\xfb\x4d\xfb\x4f\x33\x7d\x85\x69\xff\x18\xd3\x7f\x5c\xa2\xf5\xe3\x4c\x7f\xc1\xe4\xdf\x63\x7a\xca\xf4\xdf\xcd\xf4\x1e\xd3\xff\xa3\x4c\x5f\x6b\xda\x1f\x62\xfa\x2d\xd3\xff\x61\xa6\xbf\x6d\xda\x77\x31\xfd\x82\xe9\xff\x41\xa6\x27\x4c\xfe\x0f\x32\xfd\x8b\x26\xff\x0f\x33\xbd\xd2\xb4\xef\x67\x7a\xd8\xb4\x1f\x60\xfa\x9f\xcd\xf1\xf4\x31\xfd\xba\x39\x9e\x27\x98\x1e\x32\xed\x8f\x31\xfd\x3d\x73\xbe\x31\xa6\x5f\x33\xf9\xef\x65\xfa\x1b\xa6\x7d\x9c\xe9\x57\x4d\xff\x8f\x30\xfd\x8a\x39\x9e\x4f\x32\xbd\xa4\x52\xeb\x7c\xdd\xab\x35\xf3\x8e\xaf\xcf\x2d\x66\x9e\xf2\x7d\xe4\xdb\x46\xe7\xeb\x61\xaf\x39\x7e\xbe\xee\x7d\x60\xf2\xcc\xd7\xf3\x5b\xc5\x5a\xe7\xeb\x76\xbd\x39\x5f\xbe\x5e\x9d\x33\xfd\xf0\x7d\xea\x9c\xe9\x87\xef\x47\xdf\x33\xe7\xc5\xf7\xcd\xf7\xc2\x5a\xe7\xeb\xd5\x80\xd1\xf9\xba\xb7\xd1\xf4\xcf\xd7\xe7\x37\xcd\x38\xe1\xeb\xf0\xe3\x26\xcf\xfc\x7e\xe0\xef\xe6\x78\xf8\xbe\xff\x17\x93\x4f\xbe\xef\x7c\xc1\xb4\xe7\xfb\x63\x8d\x69\xcf\xf7\xf1\xaf\x1a\x9d\xef\x47\x9d\xe6\xba\xf0\x7d\xe7\x0f\xe6\xbc\xf8\xfe\xd8\x6b\xb8\x7c\x9f\x2d\x37\xd7\x85\xef\x17\x4f\x9b\xeb\xc2\xef\x13\x9e\x32\xfd\xf3\xfb\x81\x57\x4c\xff\xfc\xbe\xe5\x6d\x93\x7f\xbe\x5f\x3c\x68\x74\xbe\xef\x54\x98\xfe\xf9\xfe\xf8\xa2\xb9\x2e\x7c\x1f\x2c\x31\xfd\xf0\xfb\xab\x03\xe6\x7a\xf1\xfb\xb4\x5a\xff\xa7\xbe\x7d\xbd\x5f\xdf\x77\x92\x7a\xd0\xdf\x45\xfc\x5d\xc4\x3f\x40\xfc\x03\xc4\xef\x24\x7e\x27\xf1\xef\x23\xfe\x7d\xc4\xdf\x49\xfc\x9d\xc4\xdf\x48\xfc\x8d\xc4\xdf\x44\xfc\x4d\xc4\xdf\x4e\xfc\xed\xc4\xdf\x41\xfc\x1d\xc4\x6f\x25\x7e\x2b\xf1\xdb\x88\xdf\x46\xfc\x3a\xe2\xd7\x11\xbf\x9e\xf8\xf5\xc4\x6f\x22\x7e\x13\xf1\xb7\x12\x7f\x2b\xf1\xb7\x11\x7f\x1b\xf1\x9b\x89\xdf\x4c\xfc\xf5\xc4\x5f\x4f\xfc\x0d\xc4\xdf\x40\xfc\xcd\xc4\xdf\x4c\xfc\x2d\xc4\xdf\x42\xfc\x06\xe2\x37\x10\xbf\x91\xf8\x8d\xc4\x7f\x80\xf8\x0f\x10\x7f\x1f\xf1\xf7\x11\x7f\x3f\xf1\xf7\x13\xbf\x83\xf8\x1d\xc4\x8f\x10\x3f\x42\xfc\x16\xe2\xb7\x10\xbf\x9d\xf8\xed\xc4\xdf\x4b\xfc\xbd\xc4\xdf\x4d\xfc\xdd\xc4\xdf\x43\x7c\x5d\x5f\xe5\xd7\x2b\x03\x63\x2e\x57\x2f\xf2\x73\x90\xab\xeb\xdd\xa2\x27\x50\x2f\xd1\xcf\xf0\x81\x7a\x69\xe0\x2b\x25\x5d\x2f\xd3\xcf\xee\x81\xba\x5e\x9b\xcf\x07\xea\x61\xfd\xcc\x1e\xa8\xeb\x55\x6f\x2a\x50\xd7\x9f\xeb\x81\xba\xde\x25\x6f\x06\xea\x73\xf5\x27\x94\xaf\xeb\xd5\xba\x36\x50\x9f\xaf\xaf\x41\xa0\xae\xef\x12\xba\x02\x75\xf3\x4c\x1e\xa8\xeb\xdd\x2a\x19\xa8\xdf\xa5\x9f\xc5\x03\x75\xbd\xab\x9f\x0d\xd4\xcd\x33\x78\xa0\xae\x77\x97\xcb\x81\xfa\x52\xa5\xd4\x6f\x02\x75\xbd\xfb\xdf\x08\xd4\xab\xf5\xfd\x45\xa0\xae\x77\xa1\xaa\xa2\x7c\x7d\x85\x52\xea\x9e\x40\x5d\xdf\xfd\xb4\x06\xea\xfa\xda\x1d\x0c\xd4\xf5\x2e\x7c\x2c\x50\xbf\x5b\xdf\x77\x07\xea\xfa\xee\xe6\x4c\xa0\xae\xc7\xc6\x2b\x81\xba\xde\x4d\x27\x03\xf5\xb5\xfa\x2f\x0d\x19\x6f\x2c\xa3\x06\x53\x5e\x26\x99\x4a\xf4\x7a\xd1\x68\x7f\x6c\xe4\xf8\x90\x17\x1d\x18\xf1\x32\xd1\x78\x7a\x30\x1a\x8b\xc7\xbd\x64\x46\x0d\x16\xb6\x1b\x52\xde\xd0\x9d\x4e\x1a\xa1\x8b\x2d\xef\xa4\x37\x92\x51\xc3\xb1\x64\xba\x11\x77\x1d\x4d\xa6\x12\x4f\xa4\xa3\x29\x6f\x66\xab\x4c\x3c\x19\x3d\xb9\x3d\x1a\x4f\x8c\x8c\x78\xf1\x82\x8d\x22\xb3\x34\x1a\x1e\x3c\x3e\x90\x8a\x49\x6e\x5f\xbc\x3f\x31\x3a\x52\xc8\x1e\x4e\x1c\x17\x6d\xd3\x77\xc1\x9e\x05\xb3\x90\x97\x48\x7a\x92\x35\x9a\x1a\xc8\x78\x05\x92\x91\xf6\x32\xd1\x74\x26\x56\xb0\x4d\x7c\x28\x91\x96\xfc\x42\x5e\x36\x15\xe2\xf9\xc8\x5e\xca\x8b\x05\x2d\x3e\x16\x67\x5e\x6b\x3a\x14\x89\x3b\x73\x24\x22\x13\x3a\x02\x35\x52\x90\x1a\x29\x44\x8d\x88\xd4\x48\x01\xaa\x3f\x1a\x29\x2e\x27\xcf\xe4\x1c\x3e\x75\x78\xa6\xc3\x65\x4e\xc8\x8d\x68\x8a\xb8\xa3\x73\xc6\x4c\x0b\xe8\x90\x62\x26\x06\xa0\x64\x75\x48\x09\x58\x40\x37\x83\xc6\x2c\x18\xe9\xd9\x57\x8c\x59\x57\x8b\x59\x57\x0a\x79\x95\x28\xb0\x42\x14\x58\x1d\xa4\x95\x41\x5c\x15\xa4\x15\x41\x58\x0d\xa4\x95\x60\xb6\x55\xa0\xd0\x0a\x20\xcd\x7e\x71\xe6\x4b\xb3\x5e\x98\xf1\xd2\x0c\x80\xe3\x5f\x1a\xfd\x70\xec\xe3\xe1\x2d\x0e\x6e\x3c\xb4\x69\x27\x52\x1f\xb0\x0b\xbf\x07\x34\x35\x46\x13\xc8\xc8\xea\xd1\x93\x5e\x2a\x3d\x00\x1b\xe8\x6b\x4d\x0f\xc9\x68\xfc\x88\xf2\x32\xd1\x72\xe1\xe9\x53\xe9\xec\x20\x01\x9c\xac\xce\x3b\x0d\xe8\x54\xc4\xeb\xe7\x9d\x91\x86\x96\xcf\xbc\xc9\x57\x4f\xe2\x21\x03\x13\xcd\x18\x45\xb4\xac\xc1\x49\x01\x9d\x8a\xe0\x0a\xa1\xde\x51\xcf\xe6\xfa\xd3\x9e\xf3\x62\x74\x68\x20\xee\x8d\xf8\x6e\x83\xd7\x1f\xed\x4b\xc5\x86\x11\x30\x3b\xb5\xf0\x22\x2a\x2e\xa1\x78\x01\x85\x03\x0e\x74\x2e\xf5\x0d\xbb\x66\x3d\xeb\x61\xa5\xa7\xf7\x8c\x71\x66\x04\xde\x63\x5e\xa6\x5a\x3a\x93\xca\xc4\x7a\x55\x43\xfa\xd4\xb0\xfe\xf9\x9f\x97\xaf\x57\x2a\xf8\x76\x6c\xc2\xff\x0a\xf7\xf5\xca\x99\x3a\xfd\x3f\x49\x43\xfe\xa7\x8c\xe8\xed\x02\xaf\x84\xd4\x9f\xab\x2c\x1c\x4f\xdf\xf1\x85\x49\xfd\xcb\x95\x0a\xbe\x95\x6b\xf2\xbf\x42\xcc\x3d\x07\x8e\xf8\xcf\x71\xb9\xf8\xdc\xfb\xee\x84\xc0\xa7\xef\x95\x25\xfe\xa0\xc0\xef\x06\xfc\x52\xc0\x6f\x0d\x63\x3e\x7d\xbf\x2c\xf1\x23\x61\xcc\x4f\x02\x7e\x39\xe0\xd7\x0a\x7c\xfa\x9e\x59\xe2\xaf\x12\xf8\x13\x80\x5f\x01\xf8\x7f\x2c\xc7\x7c\xfa\xbe\x59\xe2\xff\xae\x1c\xf3\x27\x01\x7f\x0e\xe0\xff\x42\xe0\xd3\xf7\xce\x12\xff\xe7\x02\x7f\x0a\xf0\xe7\x01\x7e\x5f\x05\xe6\x57\x15\xdb\xf1\x7b\x2b\x30\x7f\x1a\xf0\x17\x00\xfe\x23\x02\xff\x8c\x25\xff\x21\x81\x5f\xb5\x9f\xf3\x17\x02\xfe\x3b\x65\x98\x9f\x5b\x28\x66\xe3\xff\xac\x4c\x98\xff\x80\xbf\x08\xf0\xdf\x10\xf8\x63\x96\xfc\xef\x0a\xfc\x6e\xc0\x5f\x02\xf8\xbf\x2e\xc6\xfc\x69\x4b\xfe\x3b\xc5\xc2\xfc\x07\xfc\x65\x80\xff\xa6\xc0\xef\x2f\xb5\xe3\x4f\x0a\xfc\x09\xc0\x5f\x0e\xf8\x8f\x0a\xf3\xef\x9a\x25\xff\x61\x69\xfe\x03\xfe\x4a\xc0\x6f\x13\xf8\x3d\x65\x76\xfc\x16\x69\xfe\x03\x7e\x0d\xe0\xcf\x17\xf8\x53\x96\xfc\x4a\x81\x3f\x0d\xf8\xab\x01\xff\x9f\xc2\xf8\xef\x2a\xb7\xe3\xff\x4d\x18\xff\x55\x1d\x9c\xbf\x06\xf0\x9b\x4b\x30\xff\xb2\x25\xbf\xb1\x44\x98\xff\x80\xbf\x0e\xf0\xab\x05\x7e\x6b\xd8\x8e\xbf\x58\xe0\x77\x03\xfe\xbd\x80\xdf\x22\xdc\x7f\x4c\x5a\xf2\xb7\x0a\xf7\x1f\x49\xc0\xdf\x00\xf8\x35\x02\xbf\xae\xc2\x8e\xbf\x5c\xe0\x4f\x00\x7e\x3d\xe0\xff\x56\xd8\x7f\x5e\xb3\xe4\xbf\x2b\xec\x3f\x93\x80\xbf\x09\xf0\x7f\x22\xf0\xab\x2b\xed\xf8\x3f\x12\xf8\x53\x80\xbf\x05\xf0\xcf\x09\xe3\x6f\xc2\x92\xff\xb2\x30\xfe\xa6\x01\xbf\x11\xf0\xbf\x24\xf0\xc3\x73\xec\xf8\x4f\x0a\xfc\xaa\x03\x9c\xbf\x15\xf0\x07\x84\xfd\x67\xdc\x92\xef\x09\xfb\x4f\x13\xe0\x37\x03\xfe\x41\x81\x7f\xd3\x92\xdf\x25\xf0\xbb\x01\xbf\x05\xf0\xff\x51\x84\xf9\xc9\xb9\x76\xfc\xbf\x16\x09\xf3\x1f\xf0\x77\x00\xfe\xbb\x02\xff\x86\x25\xff\xaa\xc0\x9f\x00\xfc\x36\xc0\x3f\x21\x3c\x3f\x1e\x9b\x67\xc7\xef\x53\xc2\xfc\x07\xfc\x9d\xe8\xfa\x0b\xfc\xf7\x2d\xf9\x5d\x02\x7f\x0a\xf0\x77\x01\xfe\xeb\xc2\xf3\x57\xf7\x7c\x3b\xfe\x37\x84\xe7\xaf\x69\xc0\xdf\x03\xf8\xcf\x0b\xfc\x2b\x96\xfc\xe7\x04\x7e\x55\x27\xe7\xef\x05\xfc\xb9\xc2\xfa\xdb\xbe\xc0\x8e\x1f\x16\xd6\xdf\x26\xc0\xdf\x07\xf8\x1f\x0a\xe7\x7f\xd1\x92\x3f\x2d\x9c\x7f\x37\xe0\x77\x00\xfe\x05\x7f\xfd\xa5\xdf\xc1\x34\xf9\x0d\xaa\xc9\x17\x2e\xf4\xfb\x93\xef\x97\xe2\xf8\xae\x85\x76\xf1\xfa\x39\x1f\xc5\x1f\xbb\xcb\x2e\x7e\x4d\x19\x8e\x1f\x5b\x64\x17\xbf\xae\x08\xc7\x4f\x2c\xb6\x8b\x3f\x21\xf0\xcf\x2f\xb1\x8b\x1f\x13\xf8\x97\x97\xda\xc5\xff\x4a\xc8\xff\xfb\xcb\xec\xe2\x7f\x2a\xe4\x7f\xba\xda\x2e\x7e\x58\xe0\x87\x57\xd8\xc5\x1f\x16\xf8\xb5\x2b\xed\xe2\x6b\x04\x7e\xeb\x2a\xbb\xf8\x8f\x14\x8e\xef\xae\xb1\x8b\x7f\x56\xe0\xf7\xdf\x6d\x17\x9f\x10\xce\x7f\x7c\xb5\x5d\xfc\x0e\x81\x7f\xb6\xd6\x2e\x7e\x89\xc0\x9f\x5c\x63\x17\xdf\x26\x8c\xff\x2b\x6b\xed\xe2\xf7\x08\xe3\xff\xda\x3a\xbb\xf8\x43\xc2\xf9\xdf\xbc\xc7\x2e\x3e\x22\x9c\x7f\xd5\x7a\x4b\xbe\x70\xfe\x75\x1b\xec\xe2\x8f\x0a\xe7\xdf\x5e\x67\x17\xff\xaa\x70\xfe\x3d\xf5\x76\xf1\x2f\x08\xe7\x9f\xdc\x68\x17\x7f\x4b\xd8\x3f\xce\x6c\xb2\x8b\xbf\x2a\xcc\xbf\xd7\x36\xdb\xc5\x7f\x20\xf0\x2f\x6e\xb1\x8b\xbf\x20\xf0\xa7\x1a\xec\xe2\x7f\x20\xf0\x6f\x34\xda\xc5\x7f\x45\xe0\xe7\x7e\x25\x77\xb6\xf8\x3f\x09\xd7\xbf\x7a\x9b\x5d\xfc\xef\x85\xeb\xdf\xd4\x6c\x17\x5f\x2a\x8c\xff\xae\x88\x5d\x7c\xb1\x30\xfe\x8f\xb5\xd8\xc5\x8f\x55\x08\xfb\xff\x76\xdc\x9e\xbe\x3f\x7a\xa9\x1c\xc7\x9f\x16\xe2\x69\xfd\x19\x9f\x4f\xfe\xfb\x7c\x35\xee\xc7\x5f\x26\x06\xbd\x7f\x1b\x17\xee\x1f\x93\xb9\xfb\x37\x3f\x7e\xc4\xff\x9d\x3f\x7a\xff\xf6\x6a\x25\x67\xeb\xd2\xee\xff\x8a\xf3\xb4\xdf\x79\x88\x7c\xff\x94\xfb\x77\x95\xff\x0e\x00\x00\xff\xff\x36\x02\xa9\x67\xf8\x6f\x00\x00")

func trace_eventsBpfBytes() ([]byte, error) {
	return bindataRead(
		_trace_eventsBpf,
		"trace_events.bpf",
	)
}

func trace_eventsBpf() (*asset, error) {
	bytes, err := trace_eventsBpfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "trace_events.bpf", size: 28664, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"trace_events.bpf": trace_eventsBpf,
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
	"trace_events.bpf": &bintree{trace_eventsBpf, map[string]*bintree{}},
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
