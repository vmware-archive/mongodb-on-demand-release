// Code generated by go-bindata.
// sources:
// om_cluster_docs/3_2_cluster.json
// om_cluster_docs/replica_set.json
// om_cluster_docs/sharded_set.json
// om_cluster_docs/standalone.json
// DO NOT EDIT!

package adapter

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

var _om_cluster_docs3_2_clusterJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe4\x98\x5b\x6f\x9b\x30\x14\xc7\xdf\xf3\x29\x10\xcf\x69\x08\xb4\x49\xb7\xbc\xb5\x59\xa5\x4c\x5a\xd2\xa9\x89\xba\x87\x36\x42\x06\x5c\xb0\x06\x18\xd9\x4e\xa6\x6c\xca\x77\x9f\xcd\xdd\x04\xd8\x45\xcb\x40\x2d\x6f\x70\xce\xf9\xfb\x5c\x7e\x38\x38\x3f\x06\x0a\xbf\x70\xc4\x10\x0e\xe9\x4c\x49\x6e\xc5\xe5\xe0\x6f\xa1\x8f\x81\x73\x0b\x28\x9c\x29\xaa\xb6\x07\x44\xf3\x91\xa5\x05\x38\x74\xb1\x63\x5d\x04\x01\xbd\x00\x3b\x86\x03\x20\x42\xd5\x61\x6d\xe0\x17\x14\xf2\x5b\xae\xab\xce\x67\xcf\xcf\x0d\xa1\x71\xe4\x31\x11\x88\x5d\x3e\x58\x8f\x90\xd0\x24\xa1\xa7\x22\xa3\x10\x04\x22\x93\xcb\xd1\x78\x34\xcd\x82\x94\x1a\xb3\x3e\xba\x4e\xcd\xdb\x44\xd4\x02\xf6\xd7\x5d\x54\xab\xe9\x61\xca\xd2\xc0\xe5\xcd\x7c\xf1\x71\x75\x67\x2e\xee\xd7\x9b\xd5\xcd\xf2\xae\x54\x92\x8f\xdd\xcf\x80\x79\x79\x1b\xb0\xdb\xd0\x06\x2d\x59\xea\x02\xb8\x30\x64\x23\xee\x27\x8b\x3c\x60\x06\x18\x2c\x77\x59\x5c\x14\x7d\x87\x1b\x8f\x40\xea\x61\xdf\x59\xde\xce\x14\x7d\x3c\x1e\x0f\x25\x17\x86\x82\xc2\x65\x41\x78\x0d\xc6\x55\xee\x70\x94\x8a\xe5\x89\x21\x86\x09\x0a\xdd\xff\x52\x70\xb1\x5c\x97\x45\x47\x04\xdb\x90\x52\x28\xd7\x0a\x88\x4b\x0d\x73\x5a\x5d\x3a\x84\xac\xfa\x28\xd6\xc0\x84\x3f\x37\xde\xf1\x3c\x24\xdb\x51\xce\x8a\xc0\xc8\x47\x76\x5c\x7c\x9d\x8a\x30\xaf\x21\x5b\x25\x2d\xb6\xfc\x1d\x54\xdb\xd4\x28\xef\x1d\xef\x5b\x9d\x92\x63\x65\x23\x60\x41\xa4\x39\x80\x01\x4d\xc8\x99\xe3\x76\xc1\x03\x65\x30\xf8\x84\xdd\x5a\x49\x48\x19\x0a\xd3\xdc\xd5\x17\xe4\x43\x75\x78\xda\x87\xfa\x55\x33\x00\xe2\x01\xcb\x19\x0c\x6a\x72\xf9\x6d\xce\xce\x84\xc7\xb0\xba\x2f\xa4\xbd\x2b\x9e\xa7\xcc\x6c\x0e\x91\x30\x27\xd5\x95\xcc\xfb\xe4\xed\xc9\x77\x94\xc2\xc2\xe9\xf7\xd6\xb6\x07\x03\xf0\x98\xf9\x5c\x9e\x6e\x47\x7f\x0b\x9f\xde\x6b\xf8\xf4\x4e\xe0\xd3\x5f\x05\x7c\x7a\xff\xe1\x33\x7a\x0d\x9f\xd1\x09\x7c\xc6\xab\x80\xcf\xe8\x3d\x7c\xfa\x75\x17\xd3\x8d\xab\xa7\xe6\xe5\x1f\x0e\xd9\xf6\x77\x3c\x1f\x92\x76\x77\x9e\xdc\xa9\x7d\x87\x20\x2b\xb6\x9d\x05\xda\x03\x16\xde\xb7\xb2\xe0\x01\xe2\xf0\x2f\xce\x3a\x89\x74\x30\x0f\xd8\x17\xf5\xd8\x38\x7c\x41\x2e\xdd\x93\x7f\xbd\x15\x25\xc2\xe6\x55\x27\x3b\x52\xb6\xf8\x59\x36\x26\xe0\x23\x40\xfb\x85\x6d\xde\xec\xbe\x6f\x61\x46\xeb\xc9\xa1\x3f\xd8\x4e\xba\xc4\x76\xf2\xc6\xb0\x9d\xf4\x1f\xdb\xd6\x33\x47\x7f\xb0\x9d\x76\x89\xed\xf4\x8d\x61\x3b\x3d\x2f\xb6\xe9\x1f\x37\xe9\x99\x82\x9f\x1b\xe4\xbf\x6e\x4c\xe4\x64\x87\x88\x42\x2d\x80\x81\x05\x89\xec\x98\x3b\x57\x2a\x17\xed\x2f\x4e\xfd\xa5\x9a\x95\x9a\x58\xbd\x39\x56\xff\x55\xac\x21\xc7\x02\x62\x21\xfe\x46\xdc\x87\xfe\x61\xa6\x30\xb2\x83\xcd\xd2\x46\x05\xc1\x88\x20\x4c\x10\xe3\x71\xc5\xcf\xc8\x71\x2b\xb5\xab\x78\x17\x4b\x2d\x48\x06\xb6\x86\x64\x2f\x3e\x48\x9f\x24\xcd\xc6\x1f\xd0\x13\xe3\xa4\xcd\x58\x7a\xf3\xb6\xb5\xd4\x9c\x7e\x01\xc7\xa9\x36\xcc\xaa\x3a\x58\x71\x89\xb9\x56\x4e\x8d\x79\xf1\x83\xe3\xe0\x67\x00\x00\x00\xff\xff\xb5\x47\x39\xde\xa8\x16\x00\x00")

func om_cluster_docs3_2_clusterJsonBytes() ([]byte, error) {
	return bindataRead(
		_om_cluster_docs3_2_clusterJson,
		"om_cluster_docs/3_2_cluster.json",
	)
}

func om_cluster_docs3_2_clusterJson() (*asset, error) {
	bytes, err := om_cluster_docs3_2_clusterJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "om_cluster_docs/3_2_cluster.json", size: 5800, mode: os.FileMode(420), modTime: time.Unix(1463071475, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _om_cluster_docsReplica_setJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x58\x4b\x4f\xf3\x38\x14\xdd\xf7\x57\x54\x59\xf7\x99\x79\xea\xdb\x15\x2a\x0d\x8b\x29\x20\xca\xc0\x02\xaa\xca\x49\xdc\xc6\x43\x62\x47\xb6\x0b\xea\xa0\xfc\xf7\xb1\x9d\x34\xb5\x1d\xa7\x0d\x12\x05\x3e\x09\x56\xd4\xbe\x8f\x73\x8f\xcf\x75\xaf\xfb\xda\xe9\x8a\x3f\x8f\x64\x1c\x11\xcc\xbc\x1f\xdd\x62\x41\x2d\x46\xe4\x05\x27\x04\x44\x67\x80\x41\xb1\xe3\x0d\x9f\x01\x1d\x26\x28\x18\xa6\x04\xaf\x49\x14\xf4\xd3\x94\xf5\xc1\x86\x93\x14\x48\x6f\xaf\xe7\x76\xbd\x47\x58\x7c\x94\xb1\xbd\xf3\x1f\x8f\x8f\x0d\xce\xca\x37\x2f\x42\x78\xca\x66\x1a\xdc\x41\xca\x4a\x58\x0f\x1a\x2e\x0c\x52\x85\xe7\x97\x81\x3f\xf8\xa3\x74\x5c\x94\x9e\x01\x08\x9f\x36\x59\x83\x63\x4c\x18\xdf\x39\x17\x30\xfa\x98\x44\xb0\x3f\xd2\xa1\x27\x64\x7d\x0d\x78\xbc\x2f\x98\xac\x1b\x0a\x1e\x16\xc9\xfa\x60\x0d\x31\x1f\x08\x3b\x2b\xcc\x0d\xe1\x80\x43\x83\x53\xb5\xc5\xd0\x7f\xf0\x36\xa6\x90\xc5\x24\x89\x66\x67\xc2\x60\x3c\x1a\x8d\x7a\xa6\x11\x47\xe9\xde\xe8\x82\xca\x52\xfc\x5f\x2b\x93\xbc\xaa\xba\x22\x0c\x71\x42\x11\x5e\x7f\x50\xe9\xfb\x84\x9f\x5b\xbe\x32\xce\x28\x09\x21\x63\xd0\x2e\x1a\xd0\x35\xf3\x97\xbf\xd7\x21\x60\xc8\x6b\x8b\x45\x24\x42\xe5\x8e\xff\xa7\x40\x64\xec\xe6\x16\x3e\x0a\xb3\x04\x85\x85\x74\x9d\x91\xa4\xc1\x1c\xf2\xcb\x92\xf3\x2c\x5c\x2d\xe5\x92\x77\x30\x2a\x13\x94\x0a\x3a\xdd\x11\xa3\xc0\x38\x9c\xe7\x10\x64\x43\xe9\x00\xab\x33\x8a\x00\x07\x47\x12\x6c\x19\x87\xe9\xdf\xe2\xac\xdc\x29\x20\xe3\x08\xef\xaa\xf2\x56\x28\x81\x5e\xcf\xc1\x52\x1d\xc7\x96\x69\x6a\x59\x4a\x69\xed\x50\x29\x65\x98\xa0\x3a\x0e\x78\xad\x15\x7a\x4a\x59\xf5\xea\x97\x4c\x23\x94\x52\x74\xb7\xdb\x4c\xb3\xd3\x0d\x9e\x8b\x4e\xdc\xdf\x54\xda\x9e\xe8\xa4\x78\x1e\xc6\x30\x05\x77\x95\xd5\x6f\xe5\x15\xd8\xfd\x56\xf0\xa1\x04\x3f\x9f\x82\xc7\x5f\x47\xc1\xe3\x6f\x05\x7f\x2b\x58\x81\x7a\x93\x82\xfd\xaf\xa3\x60\xff\x23\x14\xbc\x1b\x2c\x4a\x85\x09\x0d\xd9\xa3\xc5\x12\x45\x86\xa4\xb4\xb8\x29\x4c\x03\x48\x2d\x07\xcd\xc9\xe6\x43\x92\x5e\xff\xa6\xd1\x08\xb1\x59\x2e\xc2\x8c\x5b\x84\x19\x1f\x0f\xe3\x5b\x61\x00\x0d\x10\x87\xf4\x0a\x27\x5b\xb1\xcb\xe9\x06\xb6\xc8\xe3\x5b\x0a\x15\x07\x83\xc4\x70\xc8\x65\x88\x7d\x0f\xe7\x0b\x8b\x5d\x92\x14\x23\xdb\x6e\x81\xc5\x80\x46\x62\xa4\x2c\xd7\x3a\xd5\x31\x99\xef\x12\x39\x86\xfe\xc3\x20\x55\x28\x1a\x9f\x20\x72\xf5\xfa\x45\x9d\xd2\xeb\x6b\x37\x03\x8c\xbd\x10\x1a\x75\xf3\xdc\x78\xa7\x88\xb3\x23\xdb\x54\x0c\xb0\x13\x91\x66\x06\xc3\x18\x60\xc4\x52\x85\xca\xac\x68\x7e\x7e\x33\x99\xf5\xe7\x17\x13\xe3\x0e\x55\x5b\xb3\xab\xcb\xbf\xae\xa6\x67\xfd\xf3\x9b\x3d\xdd\x0b\x2d\xc9\x13\xdc\x96\x28\xc4\x7f\x16\x00\xb1\xa2\x3a\xdc\xe8\xe4\x7f\x49\xc0\x8c\x36\x0e\x09\x5e\xa1\xb2\xb3\x97\x24\x1d\xc8\x88\x7a\x15\x88\x81\x20\x81\xb2\xd6\x15\x48\x98\x76\x62\xde\x46\xf0\xc4\xa6\x30\x81\x5c\x6d\x3f\x2c\xec\xbd\x7b\x80\xcb\x2d\xa3\x28\xe7\xdd\x27\x51\x82\x28\x45\xd8\x75\x21\x55\xa7\x59\xdb\x72\xc7\x6b\x19\xd7\x88\x2f\x0d\xc3\x64\x23\xee\x4f\x3a\x2b\x1e\x20\x9e\xd3\x21\xaf\xad\x2e\x1c\x88\x37\x9a\x86\xec\xf7\x8c\xab\x42\x24\x4c\xdc\x8a\x3a\x74\xdf\x7f\x79\x26\x27\xca\xde\xcd\xa3\x3b\xca\xfb\x41\xa0\x10\x44\x13\xbc\x9d\x8a\x2f\xcd\x40\xfe\xca\xf0\x39\x28\xa4\x10\x14\x0b\x27\x83\x92\x90\x10\x24\x6d\x09\xb9\x17\x37\xe7\x67\x51\x71\x2c\xff\x9b\x1b\x4b\xff\x8d\xe4\x84\x4d\x65\xd4\xd9\xad\xe5\xa9\x20\x35\x1a\x1c\x68\xba\x06\x76\xdb\x50\xdb\xb6\xd7\x5c\xe7\xfa\x1e\x69\x8f\xf7\xd7\xa9\x32\xb7\xec\xa9\xb7\xa7\x3f\xd8\x4a\xad\x74\x7c\x4a\xb2\x1b\x73\xda\x4b\xf5\xbe\xb1\x5a\x41\xe5\x5c\x36\x37\x84\x73\xd4\x90\x43\x8f\x31\xc9\xc8\x60\xfa\xe8\x52\x4c\x60\x9d\xbc\xf3\x7f\x00\x00\x00\xff\xff\x7a\xa9\x97\x3d\xf9\x15\x00\x00")

func om_cluster_docsReplica_setJsonBytes() ([]byte, error) {
	return bindataRead(
		_om_cluster_docsReplica_setJson,
		"om_cluster_docs/replica_set.json",
	)
}

func om_cluster_docsReplica_setJson() (*asset, error) {
	bytes, err := om_cluster_docsReplica_setJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "om_cluster_docs/replica_set.json", size: 5625, mode: os.FileMode(420), modTime: time.Unix(1467922821, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _om_cluster_docsSharded_setJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x9c\x5f\x53\xe3\x36\x10\xc0\xdf\x3b\xd3\xef\x90\xd1\x33\x21\x71\x80\xc2\xf1\x76\x57\x1e\xfa\xd0\x6b\x19\xb8\x96\x87\x3b\xc6\x23\xc7\x22\x51\xcf\xb6\x3c\x92\x09\x93\x32\x7c\xf7\x4a\xb2\x23\xff\x89\xe4\x60\xca\x05\xc9\x98\xb9\x87\x8b\xb3\x2b\xed\xae\x76\x7f\x8e\xd7\xb2\x1f\x7f\xfe\x69\x34\x1a\x01\x92\x66\x98\x24\x0c\x9c\x8f\xf2\x03\xf2\x0f\x84\xe4\x21\x89\x08\x0c\x3f\x41\x86\xf8\x57\x60\xb2\x82\x74\x12\xe1\x60\x12\x93\x64\x41\xc2\x60\x1c\xc7\x6c\x0c\xef\x33\x12\x43\xa1\x0e\x0e\x0c\xba\x37\x38\xe1\x1f\xc5\xe8\xe0\xd7\xf3\x6f\xdf\x0c\xda\xb9\xf2\x53\x31\x08\x90\x52\x17\xc1\xdf\x88\xb2\xc2\xb4\xaf\xe5\xf0\x8f\x20\x81\xb1\xb4\xe9\xe8\x70\x76\x78\x0a\x9e\xf2\xaf\x6e\x37\xca\x01\x9c\x7f\xbf\x4f\xb7\x75\x85\x40\x39\x3e\xce\x08\xc5\xc9\x62\x5b\xac\x1a\x84\x25\x61\xd9\x66\x32\xb6\x84\x34\x44\xe1\x78\x1e\xdd\xb3\x0c\x51\x5f\x7e\xf6\x3d\xdf\x9b\xd6\x7c\x8f\xc8\xe2\x12\x66\xcb\x32\x64\x64\x61\x08\xd9\xa4\xb4\x62\x0c\x17\x28\xc9\x0e\xb9\x6c\x73\xac\x2b\x92\xc1\x0c\xd5\xd7\x46\x7e\xc7\xf0\xbf\xe8\xcb\x92\x22\xb6\x24\x51\xf8\xf9\x13\x97\xf0\xa6\xd3\xe9\x41\x43\x2a\xc3\x71\x29\xf5\x1b\x15\x6e\xce\x8e\x4b\x99\x22\x74\xcd\x08\xa6\x94\xcc\x11\x63\xa8\x11\xf8\xea\xdf\xe3\xf6\x61\x00\xe9\x82\xcd\xfc\x5f\xb6\x8d\x55\x12\x09\xca\xcc\xdf\xe6\x53\x13\x2a\x44\x66\xa7\xdc\x19\xbd\xd8\xd3\x81\x61\x70\x92\x22\x2a\x03\x7b\x49\xc9\x1d\x8e\x78\x58\xc5\x5c\x46\x71\x8a\xd2\x08\xcf\xf3\xf4\x6b\xb7\x49\x48\x5e\xa3\xec\x8f\x6a\x22\xf8\x53\xd0\xd1\x3c\xc6\x97\x9a\x2f\xf3\x8e\xb9\xc2\xa0\x96\x3d\xab\x39\x4c\x27\x42\x13\x4d\xf4\xf9\x37\xe5\xf1\xee\x6a\xc8\x9a\xeb\xc7\xbf\x93\xc5\x2e\x53\x10\xcb\x70\xb2\x09\x10\xe0\x21\x45\xc0\x30\xa6\x94\x4f\xb7\x0d\x5f\x33\x99\xff\x46\xd3\x37\x95\x21\x33\xdf\xe0\xc6\xf6\x61\x9d\x67\x80\xd7\xd5\xf2\x7a\xbe\x44\x31\x2c\x4a\x9a\x9b\x72\xa2\x13\xac\xd6\x74\x3e\xfd\x38\x21\x21\x1a\x7b\xc7\x3a\xef\xda\x6a\x50\xc9\x3c\xab\x16\x95\xf4\x8e\x9a\x6c\xf5\xb2\x1d\x46\x22\x19\x74\x5a\x45\x3d\x7f\x59\xa7\x15\xaf\xb5\x92\x2b\x15\xbb\x02\xae\x75\x99\xa6\x4d\x03\x05\x6c\xa1\x40\x73\xa5\x76\x1b\x62\x0d\x05\x4e\x6d\xa2\xc0\xac\x07\x10\x38\x1d\x20\xf0\x3e\x21\x70\xe6\x2e\x04\xce\x6c\x82\xc0\x51\x0f\x20\x70\x36\x40\x60\xaf\x10\xf0\xec\x80\x80\xe7\x7f\x70\x15\x02\xdc\x74\x9b\x20\xe0\xfe\xe5\x80\xc8\x85\x01\x02\xef\x12\x02\x5e\xf7\xdf\x24\xd6\x50\xc0\x9b\xda\x84\x81\x93\x1e\x60\xa0\xde\xa2\x54\x6a\x03\x07\xfa\xcf\x81\xee\x96\xd8\xc3\x01\xcf\x26\x0e\xe8\x3b\x6b\x8e\x71\xc0\x1b\x38\xb0\x57\x0e\xcc\xec\xe0\xc0\xcc\xf7\xba\x5b\x62\x09\x07\x84\xed\x36\x71\x40\xdf\x5c\x73\x89\x03\x32\x1b\x06\x0e\xbc\x4f\x0e\x1c\x39\xcc\x81\x23\x9b\x38\xa0\xef\xaf\x39\xc6\x01\x7d\xa3\x73\xe0\x40\xff\x39\x70\xec\x30\x07\x8e\x6d\xe2\x80\xbe\xc5\xe6\x18\x07\xf4\xbd\x4e\x87\x39\xe0\x7d\xb0\x94\x03\x73\x92\xdc\xe1\xc5\x15\xeb\x5c\x7e\x62\xa9\x0a\xab\x5a\x66\x2b\x56\xf6\x8a\x44\x95\xd9\xd8\x8a\xee\x9b\x3b\x1b\x37\x5f\x70\x6f\x74\xef\x9c\x29\x6d\xb5\x89\x2b\xfa\x96\x9d\x9d\x5c\xa9\xac\xf6\xc0\x91\x81\x23\xd5\xe9\x5e\x8b\x23\x0e\xf4\x31\x4b\x5b\x6d\xe2\x88\xbe\xe5\x67\x39\x47\x7a\xd7\xa7\x1c\x38\x62\x09\x47\x1c\xe8\x83\x96\xb6\x5a\xc5\x11\x27\x7f\x90\xf4\xae\xd1\xe9\x9d\xfe\x58\x90\xec\x31\xdd\xb5\x0d\x0a\xb9\x16\xcc\x7f\x93\x56\x5f\x61\x83\xe6\x4a\x59\x9b\x1b\xc6\x4a\x71\xe9\x94\x6b\x8a\xff\xb3\xeb\x86\x0d\x75\x33\xb2\xa2\x3d\x56\xac\xdc\x9b\x34\xc7\x5e\xab\x72\x5c\xde\x7d\xbf\x89\xff\x50\x39\x8e\x56\xce\x89\xcb\x95\xe3\xf2\x96\xf5\x4d\xfc\x7f\x70\xe5\x6c\x9e\xfe\x55\x8f\xb9\x16\x17\x4d\xfc\x6a\xa8\x78\xd0\x55\x57\x4c\x3e\x0e\x41\xe5\x49\x0f\xdd\xc4\x31\x8a\x03\x44\xbb\x3d\x2c\xdb\x18\xdf\x14\x6f\x29\x02\x69\x80\x79\x9c\xfe\x4c\xa2\x35\x17\xbd\x83\x11\x43\x6d\xe2\x4b\x1c\x86\x28\x79\x96\x24\xcf\x26\xf3\x7d\x01\xc3\xd3\x84\x4a\x3b\xa5\x98\x50\x9c\x09\x9b\xbc\x36\x39\x16\xc1\x15\xba\x40\x11\x5c\xef\x72\x74\x45\x32\xf9\xcc\xb1\xd7\x8d\x38\xbb\xc3\xdb\x6a\xdf\x9b\x85\x57\xbb\x95\x44\x69\x3b\x14\xde\x99\x95\xe1\xd5\xde\xa1\x57\xda\xfb\x0f\xef\xf6\xe1\xdb\xee\xe7\xf6\x1a\x8e\xf4\x17\x1a\x3d\xc5\x91\xe1\x69\x06\xa5\xed\x50\xbd\xd8\x88\x23\xd3\x2e\x71\xa5\xee\x50\x7c\x6d\xe4\x91\x69\xf7\xad\x52\xef\x03\x90\xf4\xd7\x6f\x3d\x05\x92\x69\x1f\xa5\x52\x77\xa8\x62\x6c\x24\x92\x69\x7f\x9a\x52\x77\x28\xbe\x36\x12\xc9\xb4\xef\x47\xa9\xbb\x4d\x24\x75\x67\xab\x97\x48\x6a\xdf\x72\xa1\xc4\x1d\x2a\x91\xb7\x45\x50\xfb\xad\x67\x25\xee\x50\x3c\xdf\x16\x39\xed\x77\xe0\x94\xb8\x8d\x88\xc9\x3f\xaa\x2e\x51\xe5\x4e\x77\x85\x15\x3a\xf0\x48\xc9\x17\x32\x25\x83\x0b\xa9\x79\xdb\xe6\xdb\xee\x5e\x94\x12\x95\x6c\x7b\xe1\xfb\x49\x5e\xd7\xce\xf6\x82\xaa\xda\xd9\x75\x8f\xcf\xeb\xda\xd9\x9e\xa8\x55\x3b\x8d\x7b\x08\x34\xb9\xd5\xa1\x0f\xaa\x3d\x51\xe5\x75\x74\x8d\xe8\x4a\x36\x89\xb5\xfe\xd4\x84\xae\xf2\xa6\xe6\xce\x33\xe0\x9c\x44\x11\x9a\x6f\xde\xfc\xf8\xd5\x5c\x04\x45\x19\x88\x8e\x36\x38\x2f\x43\x0e\x42\xcc\x60\x10\x21\x1e\xc1\x06\x14\x84\x28\xb9\x7c\x90\xa1\x7d\x7c\x1c\xa5\x90\xb1\x07\x42\xc3\xd1\xd3\x13\x68\x08\xfd\xc5\x84\x53\xc0\xf8\x26\x49\xf0\x1d\xf1\x7a\x17\x83\xf0\xff\x34\xf4\xf9\x11\x79\x0b\xe0\xbc\xd2\xe0\xff\x87\x04\xac\x68\xe2\xfb\xa2\x2d\x5e\x6c\xa5\xc8\x0f\xf9\x24\x3e\x14\xe3\x55\xc6\xb8\xe7\xf3\xb3\x1b\x98\x64\xc2\x8b\x5a\xd1\x36\x53\x0b\x84\x01\x9f\x09\x86\x31\x4e\xb6\xa2\x09\x70\x82\x33\xe9\xaf\xd9\xdd\x42\x92\x92\x48\xc0\x49\xc3\x07\x7d\x2e\xb7\x4d\x5b\x19\x92\xcb\x14\x39\xf4\x39\x7f\xa5\xa4\x26\x41\x9b\xc9\xb9\x95\x48\x32\x1a\xc5\x6a\x34\xdf\x4c\x59\x1b\xaf\x5e\x8e\xae\x46\xea\xa3\x14\xd5\xc4\x49\xa3\xfd\x7f\xa7\xa4\x08\x86\x1f\x93\xf5\x05\xcc\x60\x20\x5e\xa7\xba\x9f\x59\xc5\x7a\x4a\x2f\x5f\x77\xea\x88\xcc\x61\xf4\x0c\x87\x6f\xf8\x79\x7d\x5f\xae\xb6\xce\xd7\x25\xf1\xf3\x77\xc7\xea\x92\x5e\x7d\xb8\x6d\x02\x84\xff\x2c\x41\x39\x41\x0a\x88\x72\x59\xfe\xef\xbf\x00\x00\x00\xff\xff\x90\x44\xa0\x37\x5e\x57\x00\x00")

func om_cluster_docsSharded_setJsonBytes() ([]byte, error) {
	return bindataRead(
		_om_cluster_docsSharded_setJson,
		"om_cluster_docs/sharded_set.json",
	)
}

func om_cluster_docsSharded_setJson() (*asset, error) {
	bytes, err := om_cluster_docsSharded_setJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "om_cluster_docs/sharded_set.json", size: 22366, mode: os.FileMode(420), modTime: time.Unix(1467939366, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _om_cluster_docsStandaloneJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x56\x4d\x4f\xe3\x3c\x10\xbe\xf7\x57\x54\x39\x37\x6d\xe9\xfb\xee\x87\xb8\x15\x2a\x2d\x87\xed\x82\x28\x0b\x07\xa8\x2a\x27\x19\x1a\x2f\x8e\x1d\xd9\x2e\xa8\x8b\xf2\xdf\xd7\x76\x42\x1a\x27\x4e\x09\x12\x5d\x96\x13\x9d\xcf\x67\x26\xcf\x33\xf2\x73\xaf\xaf\xfe\x3c\x96\x4a\xcc\xa8\xf0\x8e\xfb\xb9\xc1\x18\x23\xf6\x44\x09\x43\xd1\x09\x12\xa0\x3c\xde\xe8\x11\xf1\x11\xc1\xc1\x28\x61\x74\xcd\xa2\xc0\x4f\x12\xe1\xa3\x8d\x64\x09\xd2\xd9\xde\xc0\x9d\x7a\x83\xa9\xfa\xa9\x6b\x7b\xa7\xc7\x77\x77\x2d\xc9\x26\x37\xcb\x4b\x78\x26\x66\x16\x5c\x03\x17\x05\xac\xdb\x0a\x2e\x8a\x12\x83\xe7\xbf\xe1\x64\xf8\xa5\x48\x5c\x16\x99\x01\x0a\x1f\x36\x69\x4b\x62\xcc\x84\x7c\x49\xce\x61\xf8\x94\x45\xe0\x8f\xab\xd0\x09\x5b\x5f\x20\x19\xef\x06\x66\xeb\x96\x81\x47\x79\x33\x1f\xad\x81\xca\xa1\x8a\xab\x95\xb9\x64\x12\x49\xb0\x76\x6a\x5c\x02\xff\x86\xab\x98\x83\x88\x19\x89\xe6\x27\x2a\xe0\x68\x3c\x1e\x0f\xec\x20\x89\x93\x5d\xd0\x19\xd7\xa3\x4c\xfe\x2f\x43\xb2\x72\xea\x72\x61\x58\x32\x8e\xe9\xfa\x2f\x8d\xbe\x6b\xf8\xb1\xe3\x9b\xe0\x94\xb3\x10\x84\x80\xfa\xd0\x88\xaf\xc5\x64\xf5\xb9\x09\x81\x82\x6c\x18\xf3\x4a\x8c\x6b\xcf\xe4\xab\x42\x64\x79\xb3\x1a\x3e\xa1\x86\x57\x83\xbb\xab\x44\x81\xb5\xc6\xc7\x10\xa5\x23\x9d\x00\xe5\x36\x23\x24\x91\xb7\xbf\xc1\x56\x48\x48\xbe\xab\xad\xba\x5b\x80\x90\x98\xe6\xd2\x51\x7d\xee\x31\x01\x6f\xe0\x98\xa7\x89\x63\x2b\x2a\xdf\x75\xa5\x49\xf0\x82\xca\x7c\x43\x1b\x54\xcf\x01\xaf\x33\x97\x0e\x49\x80\x41\xf3\x1c\xb4\x42\x29\xe8\x71\xb5\x4d\x2b\x71\xd5\x80\xc7\x5c\x33\xbb\x9b\x52\xf1\x29\xce\xc7\x8b\x30\x86\x04\x5d\x97\x51\x9f\x6c\xfa\x71\x48\x09\x0e\xd1\x02\xa4\x21\x60\x69\x66\x04\x2c\x83\x88\x11\x8f\x94\x62\x0a\x5b\xaf\xac\x6f\x9f\x5d\xad\xb2\x9f\x02\xb8\xc1\xda\x7a\x61\xb5\xf5\xe2\x29\xd2\x41\xcf\xcf\xfd\x14\x09\xf1\xc4\x78\xd4\xcf\x32\xeb\x0c\x2b\x64\x6c\x9b\x28\x7d\x4e\x55\x9b\x39\x84\x31\xa2\x58\x24\x06\x95\xbd\xee\xc5\xe9\xe5\x74\xee\x2f\xce\xa6\xfe\x51\x8d\x47\xde\xfc\xfc\xc7\xb7\xf3\xd9\x89\x7f\x7a\xb9\x23\xc7\xb2\xd2\xe4\x01\xb6\x05\x0a\xf5\x5f\x0d\x80\xb2\x18\x6a\x5a\x14\xfc\xc5\x02\x61\xf1\x2f\x64\xf4\x1e\x17\x94\x5c\xb1\x64\xa8\x2b\x56\xa7\xc0\x02\x05\x04\xf4\xac\xf7\x88\x08\xa8\xb8\x36\x6a\x4f\x62\x06\x04\xa4\x71\xdf\x2e\xeb\xbe\x1b\x44\x0b\x97\x35\x94\x53\xb4\x1a\x25\x8a\x12\x4c\x5d\x4a\x2a\xbf\x66\xc3\xe5\xae\xd7\xb1\xae\x55\x5f\x07\x86\x64\xa3\x84\xcf\xe7\xf9\x7d\xf5\x9c\x09\x59\xc3\xba\x74\x20\xde\x54\x38\x54\x3f\xd7\xae\x09\xb1\x0a\x71\x33\x6a\xdf\xa1\xfa\xe7\x37\x39\x35\xf1\xee\x3d\xba\xab\xbc\x1f\x04\x0e\x28\x9a\xd2\xed\x4c\x5d\xfb\x40\x3f\xa2\x3e\x06\x85\x26\x82\xd9\xc2\xc1\xa0\x10\x16\x22\xd2\x75\x21\x37\x1c\xcb\x8f\x5a\xc5\x6b\xfd\xdf\x2c\xac\xea\x13\xf0\x80\xa2\xb2\xe6\xec\x37\xfa\x94\x90\x5a\x03\xf6\x88\xae\x65\xbb\x5d\x56\xdb\x55\x6b\xae\xef\xfa\x1e\x6d\x5f\xd7\xd7\xa1\x3a\x77\xd4\xd4\xdb\xdb\xef\x95\x52\x27\x1e\x1f\x72\xd9\xad\x3d\xeb\xa6\xa6\x6e\x6a\x52\x30\x3d\x57\xed\x82\x70\x3e\x35\xf4\xa3\xc7\x7a\xc9\xe8\x62\xd5\xa7\x4b\xfe\x30\xeb\x65\xbd\x3f\x01\x00\x00\xff\xff\x4c\x33\xde\x5b\xd8\x0e\x00\x00")

func om_cluster_docsStandaloneJsonBytes() ([]byte, error) {
	return bindataRead(
		_om_cluster_docsStandaloneJson,
		"om_cluster_docs/standalone.json",
	)
}

func om_cluster_docsStandaloneJson() (*asset, error) {
	bytes, err := om_cluster_docsStandaloneJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "om_cluster_docs/standalone.json", size: 3800, mode: os.FileMode(420), modTime: time.Unix(1467931149, 0)}
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
	"om_cluster_docs/3_2_cluster.json": om_cluster_docs3_2_clusterJson,
	"om_cluster_docs/replica_set.json": om_cluster_docsReplica_setJson,
	"om_cluster_docs/sharded_set.json": om_cluster_docsSharded_setJson,
	"om_cluster_docs/standalone.json": om_cluster_docsStandaloneJson,
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
	"om_cluster_docs": &bintree{nil, map[string]*bintree{
		"3_2_cluster.json": &bintree{om_cluster_docs3_2_clusterJson, map[string]*bintree{}},
		"replica_set.json": &bintree{om_cluster_docsReplica_setJson, map[string]*bintree{}},
		"sharded_set.json": &bintree{om_cluster_docsSharded_setJson, map[string]*bintree{}},
		"standalone.json": &bintree{om_cluster_docsStandaloneJson, map[string]*bintree{}},
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
