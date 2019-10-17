// Code generated by go-bindata.
// sources:
// resources/kotlin.concept.template
// resources/kotlin.contract.template
// resources/kotlin.contractimpl.template
// resources/kotlin.pom.xml
// resources/kotlin.state.template
// DO NOT EDIT!

package contract

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

var _resourcesKotlinConceptTemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x54\x5d\x6f\xd3\x30\x14\x7d\xcf\xaf\xb8\x8d\xa6\x29\xd5\xa6\x08\x5e\x2b\xc2\x28\x65\x42\x08\x54\x90\xb6\x3f\xe0\xa5\xb7\xc3\xc3\xb5\x33\xfb\xa6\x52\x70\xfd\xdf\x91\x9d\xb4\xcd\xe7\x04\x48\xbc\x91\x87\x7c\xf9\xe4\xdc\xe3\x73\xcf\x4d\xc1\xf2\x1f\xec\x11\xc1\xda\x74\x7d\xe7\x5c\x14\xf1\x5d\xa1\x34\x81\x44\x4a\x73\xa5\x37\xcc\x9f\x31\x35\xa8\x39\x13\xfc\x27\x23\xae\x64\xba\xf2\x0b\x77\xc7\x57\x0f\x02\x27\xbe\xe2\x1b\x94\xc4\xa9\x4a\x97\x0f\x86\x34\xcb\xe9\x1b\xd3\x54\x1d\xc1\x4f\x6c\xcf\xd2\x92\xb8\x48\x57\xa5\xd6\x28\xf3\x2a\x8a\xde\x0d\xa9\xad\xdd\xe0\x96\x4b\x84\x78\x49\xa4\x3f\x60\x2e\x62\x70\x6e\xcf\x34\x58\x0b\xe9\x9a\xed\x10\x9c\x5b\xf8\x87\x8b\x03\x7c\x44\xfa\xac\x48\x70\x79\x5f\x15\xe8\x9c\xb5\x7c\x0b\xf8\x0c\xe9\x27\xf3\xb5\xf0\xd2\x99\x00\xd2\x25\x3a\x77\x63\x2d\xca\x8d\x47\x84\x4b\xa4\x0a\x94\x90\x0b\x66\x8c\xb7\x62\xe5\x6f\x9c\x4b\xac\xd5\x4c\x3e\x22\x5c\xf0\x6b\xb8\x60\x44\x1a\x16\x19\xa4\x5e\x07\x7f\x28\x09\x4d\x53\x41\x7a\x04\xbc\xf2\x4f\xf1\x75\x7c\x22\xb5\x96\x70\x57\x08\x46\x1d\xed\x81\xe7\x84\x99\x83\x8d\x00\x00\xc2\x49\xed\x51\x6b\xbe\x41\xd8\x96\x12\x48\xdd\x91\xe6\xf2\x31\x99\xc3\x02\xea\xdb\x1a\x1a\x0e\xbf\xff\x27\xa3\xe4\xa2\x5e\x81\x0c\x62\x1b\x9f\x97\x4f\xc2\x47\x45\xb7\x71\xb5\x41\x01\xd7\x76\x69\xcb\x84\xc1\x36\xd2\x17\x83\xac\xbe\x5c\x05\xb7\xfd\x27\x07\xb8\x57\xb5\xe1\x8d\x0e\xe7\xe0\x0a\xe2\xeb\x8e\x14\xec\x51\xf1\x6d\x62\x6d\x5d\xd1\xb7\xcf\x39\x98\x65\x20\x4b\x21\x82\x9a\xa3\x94\xa5\xd6\xac\x72\x0e\x2e\x2f\xa1\x8b\x9e\xcd\x52\x6e\xd6\x8a\x6e\x77\x05\x55\xc9\xfc\x68\x64\xcb\x9c\x49\xb9\xf0\x1b\x7a\x3b\xee\x74\x32\x32\x46\x9e\x6e\xb4\x2a\xbe\x30\x43\xc9\xeb\xf9\x94\x59\xb1\x6b\xd1\x6b\xa4\x52\xcb\xb0\x14\xd5\xe5\xa2\x7e\xe7\xf1\xb9\x64\xc2\x24\x8a\xbe\xa3\x5e\xc0\x52\x56\x37\xf3\x05\xbc\x57\x4a\x20\x93\x00\xb6\x63\x64\x00\x01\x6f\xa7\x76\x0e\x5d\x2b\xf6\x3e\xf4\x0a\x32\xa8\xb1\xcc\x40\xd2\x02\xf7\xa0\x9e\xeb\xd6\xd7\x3f\x57\xcc\xc2\xc8\x44\x83\x68\x4d\xce\x44\x87\x72\x7a\x04\x3b\xb0\x61\x26\xb2\x3a\x13\xbe\xff\xa4\xd2\xf1\xc5\xae\xfa\xc0\x53\xab\xef\x8b\xf6\x87\x8f\xa1\xaf\xf3\x47\x85\x9a\x64\xce\xe1\x70\x80\x89\xd4\xbe\xac\x70\x28\xb1\x49\x40\x98\xb0\xa1\x40\x3b\xc0\xf7\x67\x34\x0c\xc6\x98\x87\x8d\x8f\xb3\xc1\xb8\xe4\x4a\x12\xe3\xd2\x2c\x85\x48\x82\xd2\xa3\x48\x66\x60\xa5\x84\xc0\xdc\x37\xe6\xcd\x79\xa8\x3b\xbf\xd1\xb5\x6a\x46\xf1\xed\xc8\x66\x5e\xdc\x10\x8c\x8e\xff\x49\x29\x8c\x48\x6d\xb2\xdf\x56\xf9\x97\x55\xbb\x23\xdb\xb6\x78\x94\xee\x85\xe0\xf4\x03\x3d\xb1\xa1\x7f\xd0\xa7\x4e\xa2\xfe\x37\xeb\x2c\x67\xba\x59\x7d\xb2\x31\xf2\x46\x4d\xc3\x72\x5a\x72\x63\x13\x38\xaa\xdc\x35\xbf\x6e\xf7\x2b\x00\x00\xff\xff\xbc\x48\xb9\x08\x3f\x09\x00\x00")

func resourcesKotlinConceptTemplateBytes() ([]byte, error) {
	return bindataRead(
		_resourcesKotlinConceptTemplate,
		"resources/kotlin.concept.template",
	)
}

func resourcesKotlinConceptTemplate() (*asset, error) {
	bytes, err := resourcesKotlinConceptTemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/kotlin.concept.template", size: 2367, mode: os.FileMode(420), modTime: time.Unix(1542209137, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourcesKotlinContractTemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\xc1\x6e\xdb\x38\x10\xbd\xfb\x2b\x06\x42\x0e\x32\x60\x10\x8b\x05\x76\x0f\x5e\x04\x1b\xc3\x41\x5b\xa7\x45\x5a\xc4\x06\x7a\x0c\x68\x6a\x22\x4f\x2a\x91\x02\x39\x72\xe3\xaa\xfc\xf7\x82\x94\xe4\xc8\xb1\x5c\xa0\x39\x84\x34\x35\xf3\xe6\xbd\xc7\x19\x56\x52\x7d\x93\x39\x42\xd3\x88\xfb\xb5\xf7\x93\x09\x95\x95\xb1\x0c\xca\x94\x82\x69\xab\x8c\xc8\xcc\x1e\x59\x52\x21\x94\xd1\x2c\x49\xa3\x15\xca\xd8\x4c\x8a\x65\xf8\xbf\x34\x65\x29\x75\x76\x2b\x59\x7e\x25\xde\x85\xf5\x4f\x21\xde\x15\xe6\xfb\xd2\x68\xb6\x52\x71\x9f\xab\x91\xbb\x10\x65\x2c\xc6\xbc\xf0\xd9\x89\x3e\xf0\xbf\x0b\x91\x6c\xa5\x76\x52\x31\x19\xed\xc4\x27\xcc\x72\xb4\x9b\xd7\xa3\x3e\xe9\x59\xee\xa5\x20\x23\x56\xba\xaa\x79\xcd\x16\x65\x79\x01\xcf\xa1\x25\x59\xd0\x0f\x19\xd2\x5b\xbe\xeb\xfe\x68\x5b\xe0\x09\x60\xcd\x54\x88\x65\x6d\x2d\x6a\x75\xb8\x80\x47\x19\x6a\x26\x3e\x88\xc5\xd6\x45\x21\x5f\xa4\xe5\xc3\x6f\x2d\xb3\x28\x6c\xad\x99\xca\xa0\x8e\xf2\x1c\xad\x58\x6d\xda\xcd\x64\x72\x73\x4e\xa9\xaa\xb7\x05\x29\x50\x85\x74\x2e\xdc\x6b\x6f\xd9\x32\x1c\x78\x0f\x73\x38\xb3\x3d\x9d\xce\xa0\xdf\x43\x33\x01\x80\x40\xa5\x92\x9a\x8c\x06\xb3\x7d\xc6\xe3\x71\xf8\xbb\xb9\xdb\x97\x6b\x96\x4c\xea\x78\xb4\x97\x05\x34\xcd\xd5\xdb\x5a\x8f\xcb\xcf\xf7\x9b\x87\xc5\x72\xf3\xb8\xba\x85\x6b\x48\x42\x48\x68\x33\x31\x12\x9b\x44\x30\x0f\x71\x69\x1a\x2b\x75\x8e\x70\xc5\x2f\x1a\xe6\xd7\x20\xba\x3e\x73\xe0\x7d\x0c\x18\xd1\x1d\x59\x77\xa2\x63\xa2\xb8\x97\x25\x82\xf7\xe9\x11\x8d\x66\x70\x25\x99\x6d\x80\x8c\x11\x0b\x66\x4b\xdb\x9a\xd1\x79\xdf\x34\xf4\x04\x3a\x44\xc1\x5f\xe1\x57\x32\x4b\xc2\x82\x3a\xf3\xbe\x17\x18\x92\x23\xac\xf7\xf3\x58\x26\xa2\xfd\x84\xf7\xc8\x1f\x0d\x17\xa4\x37\x87\x2a\x94\x0c\x0a\x62\xe2\xb4\xb7\x7b\x64\x50\xd2\xe9\xc0\x54\xd2\x34\xf4\xf8\xc4\x84\x8b\x94\x4f\xc2\xab\x9a\x23\x6a\x72\xca\x33\x99\xbd\x21\x3e\xed\xa8\x8d\x27\xab\x96\x67\x32\x8b\xb7\x15\x5d\xec\x2e\xac\x77\xd4\xfb\x64\x7a\xcc\x6d\x51\x7c\x77\x69\xaf\xb8\x66\x8f\xd6\x52\x86\xf0\x54\x6b\xc8\x91\x1f\xd0\x99\xda\x2a\xfc\x20\xdd\x2e\x9d\xce\x61\xcd\x96\x74\x3e\x10\x6c\x91\x6b\xab\x43\x51\x10\xa1\x35\x5b\x53\xc3\x6e\xbd\x93\x7f\xff\xf3\x2f\x1c\x3b\x64\x32\x5a\x60\x30\xe5\x77\xce\xe8\x50\x63\x30\xdd\xe7\x85\x78\x47\x4e\x84\xb1\x8d\xed\x27\x06\x14\x17\xae\xcd\x49\x93\x93\xc7\xe4\xd9\x19\xdd\x09\x1f\xa3\xb0\x47\x4b\x4f\x87\x94\x5f\xe6\x70\xf6\xec\x0c\xef\xb9\x8d\x1b\x7c\x4c\xf9\xe5\x04\x75\x44\x59\x9c\xf5\x54\xcb\x12\x7b\xe3\xa6\xf3\xfe\x09\xf8\xff\x5c\x5a\xd3\x88\x45\x55\x75\x63\xb5\x2a\xab\x42\xbc\x41\xe9\xeb\xf9\x5f\x01\x00\x00\xff\xff\x44\xb5\x7d\x52\xfe\x05\x00\x00")

func resourcesKotlinContractTemplateBytes() ([]byte, error) {
	return bindataRead(
		_resourcesKotlinContractTemplate,
		"resources/kotlin.contract.template",
	)
}

func resourcesKotlinContractTemplate() (*asset, error) {
	bytes, err := resourcesKotlinContractTemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/kotlin.contract.template", size: 1534, mode: os.FileMode(420), modTime: time.Unix(1571329150, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourcesKotlinContractimplTemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x8e\xb1\x6e\xeb\x30\x0c\x45\x77\x7d\x05\x93\xc9\x5e\xf8\x01\x06\x1e\x1e\x8c\x4e\x59\x32\x34\xdd\x8a\x0e\x8c\xca\x0a\x4a\x2d\x89\xa0\xe8\x2c\x86\xfe\xbd\xb0\x9b\xa0\x45\xa7\xa2\x9c\x78\x81\x7b\x70\x8f\x90\x7f\xa7\xc0\xb0\x2c\x78\x3c\xb5\xe6\x5c\x4c\x52\xd4\xc0\x97\x84\x16\xcf\xbe\xe0\x6b\xb9\xb2\x51\x9c\xd0\x17\x65\xd4\x39\x5b\x4c\x8c\xbe\x24\x89\x13\x6b\xc5\x51\xe4\xe1\x16\xfe\x04\xff\x16\x32\x8d\x21\xb0\xe2\xe1\xe9\xf3\x71\xae\x9c\x2f\xec\x6d\x55\x1f\x45\x8e\x94\xb8\xb5\x43\x92\x09\x16\xe7\x00\x00\xae\x34\x01\x89\xc0\x3f\xf8\xa6\x78\xdf\x1e\x45\xba\x9f\xe0\x30\xf8\x89\x6a\xc5\x0b\x5d\x09\x03\xdb\x23\xd7\x32\xab\xe7\xb1\x9e\x4c\x99\x52\xb7\x37\xa5\x5c\xc9\x5b\x2c\xb9\xe2\xa5\x96\xbc\xef\xfb\x6d\xeb\x6d\xce\x10\xd8\x6e\x6a\x5d\xa6\xc4\x03\x9c\x4c\x63\x0e\xfd\x00\x77\xe5\xff\xb0\x6c\xed\xf5\x94\x6d\xd6\xbc\x0a\xe2\x17\x58\xbb\x7e\xb7\x7b\x5e\xe9\x97\xad\xd8\x5c\xfb\x08\x00\x00\xff\xff\xe5\xa1\xe2\xe5\xa1\x01\x00\x00")

func resourcesKotlinContractimplTemplateBytes() ([]byte, error) {
	return bindataRead(
		_resourcesKotlinContractimplTemplate,
		"resources/kotlin.contractimpl.template",
	)
}

func resourcesKotlinContractimplTemplate() (*asset, error) {
	bytes, err := resourcesKotlinContractimplTemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/kotlin.contractimpl.template", size: 417, mode: os.FileMode(420), modTime: time.Unix(1571329150, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourcesKotlinPomXml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x5b\x6f\xdb\xb6\x17\x7f\xb6\x3f\x85\xfe\x41\xf2\xf0\x0f\x26\xd2\x4e\x5a\xac\x08\x38\x01\xc3\x3a\x60\x1d\x7a\x43\x53\x14\x03\xf6\x30\xd0\xd4\x89\x42\x57\x22\x05\x92\x76\xdc\x05\xf9\xee\x03\x6f\xba\xcb\x49\x1e\xe2\x17\x8b\xe7\xc6\xdf\xe1\xb9\x91\xa4\x56\x72\x0b\xcc\x24\x87\xaa\x14\xfa\x97\x93\x5b\x63\xea\x2b\x8c\x2b\xba\x07\x81\x68\x4d\xd9\x2d\x20\xa9\x0a\xfc\xf9\xd3\x07\xfc\x0a\xad\xd0\xea\xc4\x4b\x5e\x1d\x34\x6f\xa4\xef\xee\xee\xd0\xdd\xa5\x93\xbb\x58\xad\xd6\xf8\xaf\x0f\xef\xaf\xd9\x2d\x54\x34\xe5\x42\x1b\x2a\x18\x9c\x24\x07\xcd\xaf\xb4\x23\xbe\x97\x8c\x1a\x2e\xc5\x13\x36\x4b\xe6\x24\x0e\x3a\xf7\xc4\xd4\xc9\xa1\x83\xce\x4f\xb2\x65\x92\x90\x4a\xe6\x50\x7e\x03\xa5\xb9\x14\x99\xe3\x11\xdc\xa3\x59\xa1\x5a\xc9\x1a\x94\xe1\xa0\xed\x32\x49\xc8\x86\x6a\x78\xcb\x55\x86\x0d\x54\x35\xc1\x71\xb9\x5c\x2c\x48\xa1\xe4\xae\x7e\x97\x67\xe1\x9f\xe0\x48\x58\x26\xc9\x82\x50\x65\xf8\x0d\x65\xe6\x5d\x9e\xb5\x9f\x04\x77\xc8\x4e\x6c\x1f\x36\x0f\xff\x04\xef\x5b\x34\x16\xb4\x73\x8f\xc9\xaa\xe6\x25\x28\x64\xa8\x2a\xc0\x64\x6b\xf4\x86\xe0\x69\xd6\xa4\x9a\x96\x3b\xc5\x60\x52\x2d\xb0\xbc\xda\x77\x69\x4a\x2e\x50\x84\xb0\x46\x17\xe8\xe7\x35\xc1\x03\xb2\xf5\x3d\x90\x1a\x33\x25\x15\xc5\x8e\x16\xf0\xad\x55\x6d\xf4\x66\x85\xa6\x0c\xd1\x9a\x1f\xb5\xd1\xe1\x4f\xa9\x6f\xf7\xd5\xd7\xce\x19\xcd\xb3\x97\x04\xf7\x63\xdd\x84\xf3\xf4\x3e\x7c\x3d\xf4\x22\xda\x0d\xe8\xe9\x7d\xbb\x78\x18\xc6\xb4\x09\xe9\xe9\x7d\xf8\x7a\xe8\x45\x95\xe4\x50\x83\xc8\x41\x30\xb7\xf3\x62\xd1\x12\x7e\xf8\x38\xb8\x58\xc4\xad\xa5\x2a\xd0\x16\xcc\x46\x51\x2e\x34\xf2\x0e\xf5\x80\x05\xf9\x0e\x08\x2f\x94\x6a\x93\x97\x7c\x93\x6e\xf3\xef\x6f\x86\x18\x83\x4e\x8b\xb4\x1f\xe2\x87\x61\x1a\xe2\x3e\xc4\x21\xe4\x4e\x31\x08\x30\x88\x49\x95\xd3\x7e\x39\xf4\xea\xc1\xf1\x53\x26\x15\x8c\xea\xa1\x2d\x88\x57\x68\xdd\x43\xb1\x58\x10\xcd\x64\x0d\x59\xad\xe4\x9e\xe7\x90\x13\xec\xd7\x0e\xcf\x4b\x00\xbc\xe1\xc2\x36\xa8\x94\x49\x61\x14\x65\x46\x3f\x07\xed\x33\x10\x31\x59\x21\xc3\x37\x4c\xa2\x5c\xee\xc1\x50\x5e\x1e\x81\x16\x45\xd2\xe0\xc3\x2c\xa2\x15\xba\x40\xeb\xf4\xfa\xe3\xaf\x9f\xaf\xff\xf8\xf4\xf5\x38\xb8\x6e\x52\x4e\xa1\xab\x69\x29\x2b\x2e\x64\x49\x37\x1a\x55\x60\x14\x67\xfa\x08\xc6\x20\x91\x0a\xb8\x4b\x15\x94\x9c\x1d\x81\xf9\xf7\x1a\xad\xd0\xeb\x9f\xfe\xff\x18\xc0\xe4\xec\x0c\x0e\x06\x94\xa0\xe5\xd9\xd9\xb2\x97\x92\xb1\x80\x15\xd4\x52\x73\x23\x55\xdb\xbd\x1b\x52\x53\x59\x44\x0b\x5a\xeb\x5b\x69\x74\xa7\x0e\x40\xd0\x4d\x09\x79\x76\x43\x4b\x0d\x04\xc7\x65\x54\xc1\x23\x1d\xc2\xf3\x8c\x81\x4d\x8b\x92\x60\xde\x4a\x0a\x5a\x41\xf6\x9b\x67\x24\x5f\x9a\xcd\x09\x76\x8c\x28\xb5\x53\x65\x66\xc7\x97\xbe\xc2\xd8\x22\x44\xa3\x21\xe6\x08\x17\x04\x5b\xc9\x50\x81\x23\x57\x5e\xd6\x39\x9f\x5d\x43\xd7\x1c\xf9\x69\x8e\x31\x9e\xc6\xa0\x4b\xf5\xc3\x97\x1c\x52\x97\x0c\xf6\xb6\x1d\xe3\x0e\x0f\xfb\x72\x53\x50\x02\xd5\xa0\x8f\x78\xdd\x59\xc7\xa0\x6f\x76\xbc\x0c\xa3\x34\xe7\x0a\x9c\xc1\xec\xf4\x3e\xcc\xe9\x07\x82\x5b\xaa\x37\xe9\x07\xde\xdb\x09\x59\xac\x15\xc3\x15\xe5\x02\xc7\x2e\x3b\x94\xf5\x16\xe4\xce\xd4\x3b\x33\x69\xc1\x0f\xe1\xa0\x8f\x59\x49\xb5\x73\x68\xa8\x11\x93\xd3\xdb\x6f\x0f\x3e\x52\x3a\xd1\x9b\x72\xaa\x05\xda\x98\x18\xf9\xe9\x0f\xab\x37\xdd\xf1\x60\x43\x62\x1b\x5c\xf9\xd1\x06\x30\x4c\x8c\xd3\xfb\x70\xe7\x43\x9d\xf9\xd6\xa1\xb6\xe3\xa1\x55\xf5\xa6\xea\x72\x57\x70\xa1\x7b\xab\x99\x61\x16\xd2\xdc\xe7\x7c\xd0\x7b\x64\xa0\xf9\xfb\xdc\x96\xaa\xd4\xcb\x3f\x32\xce\x2e\xd0\xe5\xa0\x1d\x07\x01\x26\xc5\x0d\x2f\x76\xca\x5d\x32\x5b\xfa\x23\x51\x9d\x89\x9f\x3f\xd3\x09\x93\x04\x47\xff\x17\x8b\xd1\x51\xcc\xcc\x6a\xef\xe1\x31\xef\x46\xe7\xf8\x84\x4b\x41\xf2\x9c\x21\xdf\x36\x8c\x03\xb0\x9d\x75\x47\xf7\x19\x7d\xe6\x98\xd7\xf6\x0e\x77\xdb\xea\x76\x8f\x81\x50\x21\x69\xa9\x33\xff\xdf\x8a\xbb\x55\xe2\xff\xa7\xb6\xc6\x33\x7b\x77\x18\xee\x3e\x65\x7f\x73\xa1\xb6\x2c\x21\xef\xa8\x12\x99\x51\x3b\x20\x38\x2c\x92\x84\xfc\x2f\x4d\x93\xb7\x5c\xdb\xee\x98\x58\x1a\x17\x85\x4e\xd2\xb4\xaf\x4b\x55\x31\x80\xe6\x89\x59\xba\xa5\x7b\x9a\xd6\x54\xd1\x0a\x0c\x28\x77\x57\x28\x32\x6f\xf5\x77\xd7\x72\x13\x6d\xa7\xa2\x49\xec\x9b\x23\xb9\x91\x2a\xf9\xf3\xfa\x4b\x7a\xb9\x7a\x9d\x50\x21\xa4\x71\x38\xc7\xfb\x61\xbf\xe1\x22\x78\x35\x95\x6e\x49\x2f\xe5\x22\xe1\x65\x4a\x50\xdf\xd2\x1c\x9e\x56\x84\x97\xf6\xfe\x31\x55\x84\xd3\xe9\x35\x97\x58\xa4\xbe\xa5\x1a\xb2\x9a\xb2\xef\xb4\x00\x82\xfd\x72\x39\x91\x4e\xcb\x89\x24\xcb\x1c\xe0\x90\x5a\x83\xa4\x99\xd0\x3a\xd6\x20\xda\xd3\xb8\x8e\xcf\xac\x0e\x8f\x0b\x56\xee\x72\x98\x4a\xdb\xc0\xea\xbc\x2b\xae\x06\xef\x87\x28\x31\xaf\x3b\xbe\x22\x5e\x9d\xcf\xea\x35\x8c\x21\x9a\x36\x62\x63\x17\xc8\x0d\x2f\x6d\xde\x8e\x8c\x79\xfa\x04\xb6\x68\x6b\x06\x5c\xc3\x9e\xea\x21\x73\x87\xd5\xf0\xac\x51\xec\x8c\xe2\x68\xd4\xe6\xbe\xa1\x5c\x80\xf2\x17\x86\xba\xc6\xe7\xe7\xf8\xdc\x56\xff\xcc\xf1\x75\xcc\x45\x23\xff\x3c\x5f\x75\x12\x89\xca\x29\xde\x6a\x29\x70\x0e\x1a\x14\xa7\x25\xff\x17\xd4\x71\xa3\x0d\x63\x7c\xc6\x78\xf2\x90\x23\x79\x98\xa4\x73\x4d\x60\xa6\x45\xf6\xdb\x63\xa4\xc5\xfe\x10\xbf\x2c\x8b\xe0\x70\x9b\x72\x6f\x63\x3b\xf1\xb3\xff\x02\x00\x00\xff\xff\x7d\x10\x8a\xc2\x05\x12\x00\x00")

func resourcesKotlinPomXmlBytes() ([]byte, error) {
	return bindataRead(
		_resourcesKotlinPomXml,
		"resources/kotlin.pom.xml",
	)
}

func resourcesKotlinPomXml() (*asset, error) {
	bytes, err := resourcesKotlinPomXmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/kotlin.pom.xml", size: 4613, mode: os.FileMode(420), modTime: time.Unix(1571332330, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourcesKotlinStateTemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x57\xcd\x72\xdb\x36\x10\x3e\x9b\x4f\xb1\xe6\x64\x3c\x64\xa2\xc1\xa4\x57\x35\x4a\xa2\xd8\x4e\xab\xc4\x75\x32\x91\x9a\x4e\x4f\x19\x98\x02\x65\xd8\x14\xc0\x00\x2b\xc5\x2a\x8c\x77\xef\x00\x24\xc5\x5f\x39\x69\x9a\xde\xca\x83\x45\x02\xfb\xf3\x61\x77\xbf\x5d\x38\xa7\xc9\x2d\x5d\x31\x30\x86\x5c\xce\xad\x0d\x02\xbe\xce\xa5\x42\x10\x0c\x49\x22\xd5\x92\xba\xbf\x8c\xf0\x25\x13\xc8\x71\x47\xa6\x57\x1a\x15\x4d\xf0\x3d\x55\xb8\xab\x84\x6f\xe8\x96\x92\x0d\xf2\x8c\x9c\x6e\x94\x62\x22\xd9\x1d\x30\xa3\x99\xe2\x34\xe3\x7f\x51\xe4\x52\x90\x53\xb7\x31\xaf\x96\xae\x32\x56\x69\x25\x72\x4d\x52\xaa\x91\xa9\xbb\x75\x46\x6e\x68\x72\xab\xa5\x20\x54\x08\x89\x85\xe6\x1b\x2d\xc5\x6c\x25\xa4\x62\x07\x1c\x25\x52\x78\x98\x9a\x3c\x0e\x8c\xe1\x29\xb0\xcf\x40\x66\x7a\x9e\x5c\xb3\xe5\x26\x73\xbe\x00\xd5\x86\x59\x7b\x40\xff\x8c\x65\x0c\xd9\x6b\xa9\xce\xde\x7c\xfc\xed\x80\x4c\x9a\xc9\x2f\x9a\xbc\xce\xe4\x97\x0b\xb9\xe2\xc9\x07\x96\xbe\xa6\x09\x4a\xb5\x0b\x8c\x61\x62\x59\xdb\x76\xc7\x41\x7e\x95\x48\xb2\x94\x5b\x86\x94\x67\x1e\x1f\xe5\x82\xa9\xd2\xe0\xe3\x20\x30\xe6\x11\x17\xf9\x06\x61\x3c\x01\x62\x6d\x60\xcc\x92\xa5\x5c\x30\x08\xe5\x96\x29\xc5\x97\x2c\xb4\x76\x7f\x96\x4b\xba\x66\x10\x66\x5c\x30\xaa\x66\xcb\xd0\xda\x4a\x08\xb6\x34\x33\x86\x65\x9a\x59\xbb\xa5\xaa\x84\x52\x21\xaa\x8d\xe6\x54\x31\x81\xa7\x19\xd5\xba\x69\xd7\xe7\xc4\xaf\x42\x18\x5a\x7b\x5a\xc6\x71\x8e\x14\x59\x61\x16\x06\x24\x0f\x46\xff\xc2\x03\xf4\xda\xa1\xb5\x8d\xaf\x3d\xae\x07\x72\x33\x6a\x2c\xb5\x95\x8a\xc3\xbc\xec\x17\xd0\xcb\x57\x2c\x93\x62\xa5\x17\xb2\x42\x1e\x19\x43\xaa\x77\x8f\xd6\xda\xf1\x38\x71\x2f\x71\x70\xe8\xd4\x81\xdf\x77\x9c\x28\x35\x22\x63\x14\x15\x2b\x06\x8f\xf8\x08\x1e\x51\x44\xe5\xb3\x34\x45\x54\xfc\x6a\x83\x4c\x97\x27\x11\x4e\x02\x9e\xba\xaf\x70\x14\xee\x91\xc2\x96\x2a\x30\xc6\x2b\xfa\xcc\x59\x3b\x06\x63\x0a\x4b\xf7\xf0\x0b\xc3\xb7\x12\x33\x2e\x16\xbb\x9c\xd5\x31\x29\xc4\x67\xfa\x5d\xee\x6a\x9e\x66\x65\x54\x5e\xb4\xa3\x10\x8f\x8d\x41\xb6\xce\x33\x8a\x9d\xac\x42\x51\x4f\xd6\x9a\xe0\x5f\x67\xee\x47\x47\xa4\x89\x79\x5f\xde\x85\x1d\x6b\x7f\x78\xa8\xa0\x8a\x15\x38\x5b\x5f\x89\x16\x98\x8a\xbf\x00\x10\x00\x00\x34\xa9\x05\x39\x55\xc8\x13\x9e\x53\x81\x1a\xc6\x70\xc1\x35\x3e\x6b\x35\xc5\xe7\x5e\xc7\x3d\x2f\x1b\x6d\xaa\x5a\x5b\x31\x8c\x62\x03\xfb\x6f\xf7\xf4\xcc\x4e\x60\xaa\x14\xdd\x0d\xd8\x8e\xe2\x96\xa6\x0b\x0c\xf8\xa8\xbc\x6f\xe8\x5b\xdb\x12\x52\x0c\x37\x4a\xb4\x3c\xec\xf7\x0b\xc9\x87\x3b\xa4\x3f\x4a\xab\x21\x76\xc2\x92\x6e\x04\x08\x76\x87\xa5\x2e\x5b\x4e\x13\xe4\x5b\x8e\xbb\x08\xaf\xb9\xf6\x25\xf4\x81\xa5\x63\xa8\xde\x46\x90\xf6\x9b\xe6\x18\x06\x3a\x69\x3c\x86\x9e\xd1\x17\x60\x82\x66\xe4\x50\xf1\x15\x4c\xa0\x4f\xf3\xd9\x3a\xcf\xc8\x8a\xe1\x42\xf1\xd5\x8a\xa9\x28\x34\x86\xf4\xac\x59\x1b\xd6\x21\xe5\x69\x54\x58\x9b\x80\xd8\x64\x59\x6c\x5a\x71\x9c\xef\x34\xb2\x35\x91\x1b\x24\xb9\xe2\x02\x33\x11\x85\xee\x20\x30\x6c\x17\xb8\x06\x21\x11\x52\xb9\x11\xcb\x91\x7b\xbd\xe6\x62\x05\xba\x92\x0c\xe3\xa1\x2c\x39\xbf\x9d\xec\xb8\xc7\xd3\xd7\x74\xaa\x46\x41\x82\x42\xc1\x04\x0a\x42\x57\x23\x25\xaa\xc7\xb1\x2f\x23\x4d\xa8\x76\xa5\xe4\xb3\x01\x54\x43\xab\xa9\xc7\x23\x80\xb0\xe6\xf5\xa7\xc1\x44\x76\xb0\x3a\xb7\x84\x2e\x97\x7b\x97\xef\x95\xcc\x99\xc2\x5d\x14\x0e\x24\x31\x1c\xcc\xf7\x37\x5b\xac\xaa\x26\x1c\x41\xb3\x9c\xbe\x59\xdf\x35\x11\x9d\xd3\x84\x85\x23\x7f\x50\x77\xcd\xe9\x9c\xa7\x1f\x57\xb9\x5e\x53\xb1\xac\x43\xeb\x3f\xcf\x28\xd2\x3f\x38\x5e\xbb\xdf\x0e\x11\x4b\x05\x92\x6f\xd0\xef\x86\xa8\xa8\xd0\x34\x71\xed\x68\xe6\xfa\x4a\x89\xfe\x2b\x5a\xe5\x42\x89\x74\xb8\x58\x7b\xad\x03\xef\xc4\x7c\x9b\x54\x58\x17\xb5\xe3\x39\x53\x5b\x9e\xb0\xc8\xd5\xd4\xa8\x72\xd6\x46\xd0\x7b\x5a\xbb\x8e\x0c\x84\x8b\xad\xbc\x65\x91\x8b\xf0\xa8\xf4\xd5\xb6\xe1\xdb\xd4\xbb\xab\x1b\x96\xe0\x73\x40\xaa\x6f\x5d\x03\xf3\x09\x59\x31\xdc\x27\x64\xaa\x77\x22\x59\xb8\xdd\xa8\x5d\xae\x64\x31\x9d\xbf\xfd\x34\x3f\xfd\xf5\xfc\xec\xf7\x8b\xf3\xb3\xe9\xe9\x62\xf6\x71\xb6\xf8\x33\xfe\xb9\xe5\xc4\x51\xb3\x30\x5d\x70\x13\xee\xef\x0b\x5f\x84\xeb\xf3\x75\x8e\xbb\x28\x8e\x3b\x04\x39\x3a\x6a\x70\xaa\x6d\xcd\x0e\x11\xea\xe8\xa8\x17\x70\xd0\x14\x26\x10\xf5\xd6\xe3\xc2\xb5\x6b\xe6\x4f\x3b\x40\xa1\xe6\xb2\xa6\x1d\xaf\x0d\x62\x07\xd0\xe8\xbe\xc5\x75\x06\xba\x3d\x15\xe5\x1c\x15\x17\xab\xc8\xcd\xad\xe2\xb5\x81\xd8\xd5\xe9\x8d\x96\x62\x5c\xec\xc0\x04\x42\x13\xd6\xdb\xfb\xf9\x3c\x38\x9b\x9b\x72\x07\x46\x68\x4a\xfd\x1d\xb2\x96\x74\xce\x60\x52\xfc\x3c\xd9\xcf\x64\xb8\x87\x85\x2c\x66\x72\x09\xc4\x5a\x78\x02\xe1\xa8\x85\x85\x75\x6c\xf1\x34\x6a\x0f\x79\x38\x2e\xf2\xea\xe1\x54\x58\x7c\xfb\xb2\x16\x4e\x4e\x3a\x57\x82\xe3\x63\xc2\xf5\xa5\xc4\x32\xf3\xd5\x78\x6f\xe7\xf3\xfb\xf1\x42\x2b\x3e\xad\x3b\xe7\x90\x75\xb2\x54\x32\xbf\xa0\x1a\xa3\x9f\xe2\x43\xe1\x0a\x6d\xc3\x7e\x59\x1f\x6e\x2b\x68\x94\x43\x2b\xf9\xec\xf3\x86\x66\x3a\x92\x78\xcd\xd4\x18\xa6\x62\xf7\x22\x1e\xc3\x2b\x29\x33\x46\x05\x34\x2b\x97\xa7\x85\x90\x9b\x37\x75\x1f\xef\x92\xc1\x37\x0a\x09\x13\x28\x64\xa9\x86\xa8\x21\xdc\x6b\x80\x5c\x9f\x3b\xff\xb5\xc7\x89\xbf\x0d\x04\xbd\xea\x3a\x78\xfb\xeb\xdc\x53\xf6\x17\x8c\xce\x15\xad\x4b\xf2\x4e\x55\x54\x6c\x3f\x39\x01\x94\x64\x78\xb3\xdf\xd1\x4a\xf4\x5d\xd0\x50\x0d\x52\x9e\x46\xff\xc8\x51\x59\x9b\xb1\x6b\x3a\x07\xea\xf6\x61\x84\x7d\x88\x65\x09\x78\x92\xf5\x01\x9a\x9e\x7c\x97\xa6\x9e\x1a\x43\x31\x2c\xe3\x78\xdc\x23\x4c\xf9\x2f\xa7\x9e\x66\x59\xe4\x91\x56\x20\xfd\x8d\x20\xcb\x98\x9f\x19\xcf\x1a\x3c\x69\x5d\xb6\x2f\x65\xc9\xc6\xe7\x03\xa7\x79\xf0\x44\x30\xd8\x01\xf6\x50\x61\x00\x6b\x59\xfc\x4d\x98\xdf\xe9\xb5\x4d\xda\x66\x8c\x07\xcd\x3d\x50\x39\xdd\x8a\x3e\x70\xa0\xff\x20\x51\xad\x92\xfa\x3f\x5b\x0d\x38\x87\xb3\xd5\x35\x36\x64\xbc\x44\x53\x5a\xa9\x87\xf3\x10\x07\x07\x91\xdb\xb2\x7b\xdb\xe0\xef\x00\x00\x00\xff\xff\x25\x95\x06\x23\x47\x13\x00\x00")

func resourcesKotlinStateTemplateBytes() ([]byte, error) {
	return bindataRead(
		_resourcesKotlinStateTemplate,
		"resources/kotlin.state.template",
	)
}

func resourcesKotlinStateTemplate() (*asset, error) {
	bytes, err := resourcesKotlinStateTemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/kotlin.state.template", size: 4935, mode: os.FileMode(420), modTime: time.Unix(1565926679, 0)}
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
	"resources/kotlin.concept.template": resourcesKotlinConceptTemplate,
	"resources/kotlin.contract.template": resourcesKotlinContractTemplate,
	"resources/kotlin.contractimpl.template": resourcesKotlinContractimplTemplate,
	"resources/kotlin.pom.xml": resourcesKotlinPomXml,
	"resources/kotlin.state.template": resourcesKotlinStateTemplate,
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
	"resources": &bintree{nil, map[string]*bintree{
		"kotlin.concept.template": &bintree{resourcesKotlinConceptTemplate, map[string]*bintree{}},
		"kotlin.contract.template": &bintree{resourcesKotlinContractTemplate, map[string]*bintree{}},
		"kotlin.contractimpl.template": &bintree{resourcesKotlinContractimplTemplate, map[string]*bintree{}},
		"kotlin.pom.xml": &bintree{resourcesKotlinPomXml, map[string]*bintree{}},
		"kotlin.state.template": &bintree{resourcesKotlinStateTemplate, map[string]*bintree{}},
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

