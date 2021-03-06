// Code generated by vfsgen; DO NOT EDIT.

package install

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// templates statically implements the virtual filesystem provided to vfsgen.
var templates = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		"/crds.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "crds.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 15124,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\xeb\x6f\xdb\x48\x92\xff\xee\xbf\xa2\xce\xf7\x41\x09\x60\xd1\x93\xe4\x70\x0f\x03\x83\x9b\x39\xd9\x79\xcc\xe4\x61\xc8\x4a\xee\xc3\x60\x90\xb4\xc8\xa2\xd8\xeb\x66\x37\xa7\xbb\x29\xdb\xbb\xd8\xff\x7d\x51\xfd\xa0\x48\x89\xa4\x24\x67\x32\x8b\x01\xd6\x1f\x82\x88\x2c\x16\xab\x7e\xf5\xe8\xaa\xea\xe6\x74\x3a\x3d\x61\x15\xff\x84\xda\x70\x25\x2f\x80\x55\x1c\xef\x2d\x4a\xfa\x65\x92\xdb\xff\x36\x09\x57\xe7\xeb\x67\x4b\xb4\xec\xd9\xc9\x2d\x97\xd9\x05\xcc\x6a\x63\x55\x39\x47\xa3\x6a\x9d\xe2\x25\xe6\x5c\x72\xcb\x95\x3c\x29\xd1\xb2\x8c\x59\x76\x71\x02\x20\x59\x89\x17\x50\xa0\x28\x35\x0a\x64\x06\x4d\x42\x3f\x92\x5c\xd4\xf7\x69\x96\x70\x75\x62\x2a\x4c\x89\x92\x65\x99\x7b\x9c\x89\x6b\xcd\xa5\x45\x3d\x53\xa2\x2e\xa5\xa1\x7b\x53\xf8\xe9\xe6\xc3\xfb\x6b\x66\x8b\x0b\x48\x8c\x65\xb6\x36\x49\xe0\xf7\x9e\x95\x78\x02\x10\xdf\x34\xf7\x57\xdd\x15\xfb\x50\xe1\x05\x18\xab\xb9\x5c\x8d\x33\xb9\x71\xbf\x5a\x6c\x5a\x17\x0e\xe0\x92\x2a\xe9\x65\x37\xbf\xfc\xef\x93\x1f\x12\x7a\xe2\xfb\xef\x4f\x83\x28\xd9\xe9\xd3\x5f\x93\x12\x8d\x61\xab\xb6\xa0\xef\x5a\x57\xc6\x5e\x11\xc1\x4c\x52\x8d\x8c\xde\xb1\xe0\x25\x1a\xcb\xca\xca\x3d\x9a\xa1\x49\x35\xaf\xac\x33\xda\x6c\x9b\x04\xb8\x01\x06\xb6\xf9\xa9\xb1\xd2\x68\x50\x5a\x2e\x57\x60\x0b\x04\x83\x7a\x8d\xda\x51\xc0\x5d\x81\xd2\xf1\x04\xb0\x05\x37\xa0\x96\x7f\xc1\xd4\xc2\x1d\x33\xe0\xde\x8d\x59\x02\x6f\x2c\xb1\x94\xca\xc2\xaa\x66\x9a\x49\x8b\x98\x81\x55\xb0\x24\x56\x16\xb8\x84\x82\x55\x15\x4a\x33\x5d\x62\xae\x34\x82\xd2\x19\xea\xc0\x95\xa5\x5a\x19\x03\x06\x2b\xa6\x99\x45\x50\x15\x6a\x27\xaf\x49\x60\x26\x38\x4a\x6b\xa0\x64\x0f\x8e\x3d\x71\x73\x52\xac\x99\xa8\x31\xbe\xb8\x91\x1f\xb3\xc0\x93\x4b\x98\xbf\x9c\xbd\x78\xf1\xe2\x7f\x20\x57\xba\x04\x26\x33\x22\xe4\x12\x3e\x2e\x66\x49\x0b\xef\x1f\x3b\x58\x67\xcc\xd2\xcf\x95\x56\x75\xe5\xdd\xb3\xe5\x91\xfe\x11\xe7\x78\x00\xde\xd7\x5f\xa3\x28\xdb\xae\x25\xb8\xb1\x3f\x6f\xdf\x79\xcb\x8d\x75\x77\x2b\x51\x6b\x26\xba\x6e\xef\x6e\x98\x42\x69\xfb\x7e\xc3\x7c\x0a\x85\x6e\xfe\x13\x48\xb8\x5c\xd5\x82\xe9\xce\xd3\x27\x00\x26\x55\x24\xb8\x7b\xb8\x62\xa9\x03\xc0\xd4\x4b\x1d\xc2\x2f\x30\xf4\x0e\x79\x01\x7f\xfb\xfb\x09\x10\x74\x3c\x73\x00\xfb\x9b\xaa\x42\xf9\xe3\xf5\x9b\x4f\x2f\x6e\xd2\x02\x4b\x76\x11\x20\xec\xb8\x50\x4b\x9d\xe0\x3c\x0f\x15\x92\x85\x1b\xe8\x81\x39\x22\x08\xa2\x25\xd1\x65\x1c\xae\xde\x67\xc2\x25\x8d\xbf\xd5\x5c\x63\x16\x5f\x34\x85\xe8\xcc\xcd\x05\x0a\xfe\xf0\xa3\xd2\xe4\x0f\x96\x47\x55\x9c\xc7\x6c\x32\x52\x73\x6d\x4b\xe0\x09\x69\xe4\x69\x20\xa3\x1c\x84\xc6\x79\xf6\xda\x5f\xc3\x0c\x8c\xd3\x16\x54\xee\x1d\xaa\x51\xc4\x21\xd3\x62\x0b\x44\xc2\x64\xd0\x21\x81\x1b\x17\x1b\x86\x8c\x56\x8b\x0c\x52\x25\xd7\xa8\x2d\x68\x4c\xd5\x4a\xf2\xbf\x36\x9c\x0d\xc1\x43\xaf\x14\xcc\xa2\xb1\x1d\x8e\x2e\x91\x49\x26\xbc\x1b\x9f\x39\xf7\x24\x17\xd7\xe8\x62\xab\x96\x2d\x6e\x8e\xc4\x24\xf0\x8e\xe2\x86\xcb\x5c\x5d\x40\x61\x6d\x65\x2e\xce\xcf\x57\xdc\xc6\x1c\x9c\xaa\xb2\xac\x25\xb7\x0f\xe7\xa9\x92\x56\xf3\x65\x6d\x95\x36\xe7\x19\xae\x51\x9c\x1b\xbe\x9a\x32\x9d\x16\xdc\x62\x6a\x6b\x8d\xe7\xac\xe2\x53\x27\xb8\xf4\x71\x56\x66\xff\xde\x78\xcc\xa4\x25\xe9\x56\x06\xf2\x7f\xce\xf5\x07\x71\x27\xf7\xf7\x2e\xe2\x1f\xf3\xf2\xef\xa6\x98\xf9\xd5\xcd\x02\xe2\x4b\x9d\x09\xba\x98\xfb\x2c\xd3\x3c\x66\x36\xc0\x13\x50\x5c\xe6\x94\x9f\xc8\x70\xb9\x56\xa5\xe3\x88\x32\xab\x14\x97\xd6\xfd\x48\x5d\xe2\xe8\xb0\x34\xf5\xb2\xe4\xd6\x38\xff\x43\x63\xc9\x3e\x09\xcc\x98\xa4\xbc\xb2\x44\xa8\xab\x2c\x24\x33\x09\x33\x56\xa2\x98\x91\x1b\x7f\x6b\xd8\x09\x61\x33\x25\x48\xf7\x03\xdf\x5e\x40\xbb\x84\x9d\xf8\x02\x88\x6b\xe7\x28\xd1\x6e\x20\xfa\xd8\x4b\x0b\xa6\xdb\x54\x7d\x01\x48\x7f\x8e\xae\x7b\x69\xf0\x4d\x63\x7c\x1a\x5e\xd7\xb5\x10\x37\x98\x6a\xdc\xe1\x0a\x3b\x2b\x5a\x97\x1e\x0a\x25\x32\x1f\xdf\x1a\x73\xd4\x28\x53\x8c\xd1\xc7\x6a\x5b\x10\xde\xe9\x6e\x5c\x37\x70\x79\x26\xb9\xd2\xc0\xd2\x14\x8d\x89\x3e\x1a\x12\x5a\xa5\x0c\xb7\x4a\x3f\x40\xed\xee\xbc\x5e\x2c\xae\x6f\x60\xc9\x0c\x4f\x1d\xf7\xa4\x97\xe9\xfb\x0f\x0b\x78\xf3\xee\xfa\xed\xd5\xbb\xab\xf7\x8b\xab\xcb\x7f\xeb\x21\x1a\x01\x0b\x06\xcc\x13\xff\xa6\x6e\x2d\xea\xb9\x31\x86\x32\xc4\x55\xaf\xf7\xce\x80\xd3\xc5\xbf\x15\xdf\x6f\x96\x57\xdc\xc2\xc7\xf9\x5b\x0a\x7f\x82\x8f\xfe\xeb\xd2\x2b\xba\x3b\x1b\x20\xcf\x00\x93\x55\x42\x2c\x7f\x58\x71\x5b\xd4\xcb\x24\x55\xe5\x85\xd2\xab\x73\xa2\x39\xeb\x15\x8f\x82\xcf\xc7\x5e\xa0\x3f\x6f\xe8\x41\x69\x30\xa6\xf0\x77\x7f\xc0\x7b\x56\x56\x02\x1d\xcb\xe7\xcf\x9f\x3f\x6f\xe8\x12\x8a\xdb\x41\x33\x0c\xa8\x3d\x84\x57\x47\x6f\x5a\x7d\xa3\xd2\xf4\x40\xd4\xda\xf9\x8f\xf3\x6d\xf8\x7c\xc7\x6d\xa1\x6a\xfb\x99\x96\x12\x26\x38\x33\xfd\x6a\x3a\x60\x34\x66\xdc\xc0\x13\x72\xc8\x2f\xb4\xdc\x43\x5d\xad\x34\xcb\x10\x7e\xc9\x05\x5b\x99\x5f\x69\x3d\x5f\x0a\x3c\x77\x74\x5f\x9e\x1e\xad\x54\x45\x35\xe4\x3e\xa5\xa8\xd0\x8c\x4a\xd1\x03\x31\xa0\xbc\x3e\x1a\x05\xb3\x7c\xdd\x84\xd9\xc6\xb8\xbd\x6a\x69\xa5\x8e\xc7\x5e\x63\xbe\x57\xca\x39\xe6\x51\x48\xf2\xb1\xa5\x66\x32\x2d\xe0\x89\xd2\xa0\x6c\x81\x7a\x93\x0f\x9e\x92\xa4\xf5\xa6\x30\xe9\xfe\x5d\x62\xce\x6a\xe1\xd6\x04\x98\x94\xcc\x58\xd4\x13\xe7\x57\x4e\x63\x25\x73\xbe\xaa\x35\x66\x54\x4a\x10\x5d\xf0\xe7\xfc\x11\x2a\x45\x98\x0e\xd0\xac\x52\xfd\xc1\xb4\x95\x96\x42\x34\xc5\xd5\xe9\xb6\x5e\xa2\x96\x68\xd1\x4c\x9d\xa9\x4c\x62\xac\xd2\x6c\x85\xc9\x4a\xa9\x95\x40\x56\x71\x6a\x4d\xca\x5e\x1c\x94\x6e\xf8\x84\x87\x5b\xe1\x74\xb4\xb6\x3e\xb1\xce\x0f\x30\xe3\x4d\xa4\x6c\xa5\xf2\x6e\xe6\xee\xcd\xd2\xbd\x3a\xec\xe6\x1b\x78\xa2\xa8\x9d\x71\x89\xfb\x69\x02\x0b\xb2\xa9\xc6\x8c\x98\x33\x61\xe0\x8e\x0b\x41\x15\x00\xcb\xb2\xa6\x81\xd8\x62\xa9\x28\x6c\x7d\xe2\x7f\xc5\x2d\xd9\x22\x34\x32\xf4\xb2\x92\x6b\xad\x34\x19\xca\x58\xa6\xa9\x86\xf8\xb3\x27\x7b\x73\xcb\xab\x4b\xac\x3e\xba\x92\x68\xbf\xf1\xda\xd4\x1e\x4d\x8b\xf4\x4f\x11\x9b\x39\x0a\x24\xe5\xb8\x82\xae\xa5\x1c\xb2\xdc\xc4\x25\xbb\x0c\xab\x50\x8c\x4d\x22\xca\x5c\x1a\xcb\x84\xa0\xd5\x57\xe9\x90\x0d\xe3\x22\xed\xfc\xf4\x8c\xfe\xdb\xcb\xd3\xa7\xab\x0c\x2b\x94\x19\xca\x94\xa3\x81\xcf\x65\x6d\xec\x67\xb2\x78\xec\x5d\x72\x15\x8a\x49\x92\xb2\x4e\x53\x1c\xb3\xe1\x52\x29\x81\x6c\xb7\x94\x58\xef\xb6\x24\xbd\x70\xc5\xb6\x24\x04\xb6\x65\x7a\x85\x16\xb3\xf6\x6a\x11\x58\x85\xc8\xfe\xaf\xe4\xbb\xe4\xd9\x51\xc1\x97\x2b\x9d\xe2\x47\xbf\x68\x6c\xcb\xd3\x91\xe5\x25\x11\x7a\x93\x95\x4c\xdf\x7a\x10\xda\x7d\x1c\x21\xf2\x65\x3a\x75\x0c\xbf\xc4\x75\xc8\xec\x0a\xb3\x70\xa5\x38\x51\xc5\x3a\x2c\x54\xf6\xde\x92\x74\x51\xab\x7a\x55\x40\x86\x02\x2d\x2d\x5d\x7e\x78\x00\x3c\x07\x89\x98\xed\x02\x3e\x0c\x36\x79\x49\x4f\xff\xb7\xa3\xdc\xe4\xf5\x86\x30\xa2\x1d\x90\xa5\x54\xea\xd4\xa4\xe5\xcb\x19\x20\x81\x37\xb9\x1f\x31\xd4\x55\x25\x38\x66\xbb\xeb\xb3\x6b\xe6\xd4\x1d\x1a\x0b\x9f\x51\xd2\x0a\x1c\x8c\x16\x98\x7e\x6e\x32\x49\xb4\x69\x02\x9f\xa8\xd9\x86\x96\x20\xbb\x6e\xe3\xfa\x3b\x60\x1a\x2f\xe0\x74\xfd\xfc\xf4\x0c\x4e\xd7\x2f\x4e\x27\xbd\x68\xf4\x86\x2b\xca\xba\xdc\xc6\x61\x0a\xeb\xe7\xbb\x97\x5e\x74\x2e\x95\xec\xfe\x35\x37\x7d\x6b\x51\x07\xc5\x77\x0d\x59\xc4\xb0\x64\xf7\xbc\xac\x4b\x60\xa5\xaa\xa5\x25\x28\x35\xae\xb9\x1b\x09\x12\x9e\xb7\x88\x15\xb9\x42\x2f\x7e\x9d\x19\xc1\x0e\xe4\xc0\x6d\x5c\x62\x1d\xab\x67\xdf\xf5\x7b\x05\xf5\xcf\x2b\xec\xbe\xa1\x35\xf8\x1b\xd5\x67\xbe\xa1\xeb\xab\xd9\x16\x03\x42\xee\xe6\xc8\xb6\xd0\xd1\xf0\x2b\x94\x94\xef\x30\x83\xe5\x03\xb0\x3c\xe7\xf7\x31\x4d\xc9\x38\xa1\xd9\x54\x4c\x3e\x42\x7a\x4b\xce\x7e\xb5\x7b\xcc\x4f\x09\xcc\x7e\x72\x2e\xb4\x47\xeb\x86\x6e\x5f\xb8\x3b\x96\x3e\x58\x1c\xfd\xae\x1d\xbd\x02\x8d\xa1\x02\x72\x4d\x22\xf3\x39\x2c\xa4\xee\x0a\x75\xae\x74\x49\x30\x30\x19\x13\xc8\xe1\xb1\xae\x95\x10\x4b\x96\xde\x8e\x2a\x47\x36\x8b\x84\x60\xd0\x5a\x2e\x57\x66\x93\xd4\xfb\xe6\x52\xdd\x37\x1f\xdd\xb0\x66\xdc\x50\xfc\xbf\x56\xea\xb6\x77\xed\xed\x88\x77\xd9\x22\xde\x07\x7e\xa5\x71\xbd\x3d\xb8\x88\x7f\x85\x63\xe0\xe6\x1d\x61\x15\x85\xac\xd6\xd1\xc1\x22\x00\xc7\x2f\x5c\x3e\x97\xed\x55\xe3\xca\x91\x8d\x2a\x40\x90\x47\x39\x7a\x16\x89\x7d\x82\xb8\xf5\x63\xaf\x1c\x47\xae\x59\xa3\xf2\x7c\x8b\x85\x6b\xbf\x9e\x25\xbb\x9f\xa3\xd5\x03\x75\xdb\x76\xf6\x0d\xa4\xc3\xd9\x37\x76\x87\xda\x13\x0e\x16\xc5\x4d\x2d\x16\x46\x97\x25\xbb\xc5\x18\xa6\x4b\xc6\xa9\xbc\x1a\xd6\xa5\x2f\xe5\x36\x56\x2b\x99\x75\x14\xff\xf9\x1f\x3d\x1d\x8f\xc7\xeb\x80\x7e\x27\x00\xbb\xdf\xb2\x91\xe7\xb4\x52\x99\xf9\xd2\xab\x2f\x79\x22\xcf\x81\x51\x8e\x4e\xc9\x6f\x13\x6f\xe9\x90\x8e\x0c\x54\x2a\x23\x5b\x5b\xd7\xe4\x1c\x6d\x41\x82\xfa\x90\x16\xce\xea\x87\xbd\x01\x73\x88\xf9\x58\x6e\x51\x03\xfb\x8a\x20\xb7\xbc\x44\x55\xef\x9f\xe1\x2c\x3c\x5d\x53\x9d\xf2\xd2\xa1\x7e\xc7\x78\xe8\xbd\xe4\x03\x70\x99\xf1\x35\xcf\x6a\x26\xe0\xe7\xa6\xdb\xec\xef\x27\xe3\x5e\x0e\x3c\x11\xfc\x16\xe1\x27\xb5\xf4\x89\xd9\xe5\xb2\xa7\x31\x7f\xed\x57\xeb\xb1\xee\x47\x72\xef\xd5\xf9\xff\x59\x5c\xc2\x87\xdd\xce\x01\x50\x4b\xcb\x05\x30\x21\x7a\x95\xbd\x56\x99\x39\x83\xeb\x4f\x33\x73\xe6\x86\xd5\x3c\x45\x13\x66\xfb\x5c\xba\x98\x95\x75\xb9\x44\x4d\x31\x4b\xb4\x6e\x63\x01\x2e\xb1\x12\xea\xa1\x44\x69\xfb\x27\x42\x37\x96\x59\xcc\x6b\x71\x83\xd6\x4d\x24\xe6\xe8\x5c\xfa\x06\x2d\xd5\x8c\xc0\x25\xb9\x05\xb2\xec\xc1\xed\xf2\x34\x01\x4d\x9a\x8c\x75\xc8\x51\x35\x66\x7c\xc7\x63\x4c\x5e\x8b\xe3\xdc\x8a\xfa\xb9\xd9\xfc\x72\xbc\xf0\xb8\x09\x44\xfb\xf0\x75\xcd\xa1\xf5\x9d\x79\xff\xa8\x56\xe5\xe0\x18\x05\x9f\x09\xfb\x4d\x2f\x62\x73\xe8\x1e\x3a\xbc\xb0\xf0\xd5\x4a\xb3\x6f\x36\x5e\x5f\x74\x69\x41\xad\x51\x6b\x9e\xe1\x56\x03\xb7\x29\xf1\x7c\xe9\xb1\x5b\xda\x75\x0b\xcb\xc5\xa6\x88\x6a\x3d\x8b\xbf\xd5\x4c\x34\x9b\x47\x9b\xeb\xbe\xca\xea\x65\x19\x77\xe6\xe2\xda\x75\x70\xf5\x38\x90\x14\xbe\x51\x42\x38\x30\x19\xb4\xed\xe9\xa2\x27\x26\xc8\xd6\xc6\xf0\xc1\x5d\xc1\x58\x7e\x70\xe5\xed\x4b\x2e\xd0\x4f\x9f\xc6\xfd\xf8\xd3\x16\x71\x6b\x50\x25\x54\xca\x84\x6f\x22\x9a\x71\xa3\x1f\x25\x78\xd2\xdd\xa8\xba\xbc\xba\x9e\x5f\xcd\x7e\x5c\x5c\x5d\x9e\x41\x6d\xd0\x33\x37\x2f\xb5\x2a\x13\xff\xcc\xcf\xf8\xe0\x26\x9b\xd2\x58\x64\x03\xad\x31\xd3\x9a\x6d\x4f\x5b\xb9\xc5\xb2\xa7\xa8\x18\x1d\x3e\x0d\x8f\x9e\x06\x06\x4f\xe3\x63\xa7\xe1\xa1\xd3\xa0\x1b\xae\xf7\xf7\x2f\xa1\x75\xd9\x80\x1e\xda\xe6\xaf\x2b\xf2\xd7\x0d\xee\xfd\x1b\x59\xbf\x17\xc4\xe3\x88\xb9\x7e\x89\x5c\x6b\x60\x52\xda\xd7\xed\xb4\xb7\xb9\x58\xf0\x40\xdf\x77\xe5\xdc\x15\xe9\xb6\x00\xdd\x3a\x99\xb0\xfb\x17\xf6\x95\x47\x2c\x35\x38\xac\x1c\x1f\x57\x92\xdf\x54\xcc\x16\xbd\xb7\xf6\x8d\x2c\x29\x4b\xf8\xe3\x3e\x43\xf7\xb7\xd0\xf8\x10\xc8\xb7\xd7\x97\x59\x04\xf5\x06\x05\xa6\x54\xf5\xb2\xfe\x12\xa5\xfb\xd6\x24\xc0\x6b\x28\x29\xc7\xb3\x01\xcc\x0d\x7a\x99\x6d\x65\x20\x77\xad\x42\x5d\x72\x6b\x07\x86\xc6\xfe\x2f\xec\xfb\xf8\x63\x35\x2e\x3b\x9f\x41\x56\x3b\xbb\x71\xeb\xe7\x7b\x4b\x74\x3d\x1c\x96\x95\xd2\x4c\x73\xf1\x00\xb5\x64\x6b\xc6\x85\xab\x5b\x07\x79\x8f\xd7\x7b\x01\xee\x81\x9d\x9e\x1e\x20\xdb\xfb\x3d\xce\x87\xda\x9b\x3e\xa1\x27\x8a\xbb\x3e\x23\xfa\x76\x77\x89\x06\xb6\x7d\xda\x0a\x0c\x8c\xa0\x21\x6c\xbb\xbc\x63\x95\x4f\x86\x8f\x8b\x0c\xcf\x04\x4a\x56\x75\x62\xe2\xdb\x7a\xff\xc0\xb8\xfe\x10\xef\xbf\xc5\xde\xb6\xa2\x2d\xd8\x20\x60\xb0\x67\xe4\x7f\x30\x83\xde\x8a\xe8\x28\x2e\xfb\x83\x78\x9f\xf7\xe2\xbd\x3f\x22\x73\xe3\x3c\xef\xb1\xf6\x97\x0d\x9f\xe8\xc2\xff\xec\xcc\x58\xeb\xfe\xe6\xe1\x0f\x4b\x8c\x57\x1d\x60\x63\x76\x1c\x89\x68\x66\x1e\x95\x1d\x47\x38\x36\x79\xf3\x51\xd9\x71\x84\xf1\xef\x96\x37\x6b\x7d\x28\xcc\xfd\x5b\xb4\x5b\x5e\xf7\xe8\x1c\xd8\xae\x06\x1f\x97\x00\xc3\xb6\xe9\xbf\x92\xdf\x9f\x23\xf9\xf5\xcd\x2b\xbe\x72\x56\xe1\xe7\x12\x3b\xb2\x7c\xc5\x9c\xa2\x33\x93\xd8\x6d\xf3\x8e\x9d\x51\x1c\x3e\x8f\xe8\xc7\x2d\x9c\x6e\x3d\x19\x40\xac\xd5\x21\xfb\xf3\xdb\x54\x14\x58\xc6\xa5\x09\x4f\xba\xe3\x7d\xd4\x24\xba\x86\x73\x49\xf9\x88\xc9\xf6\x63\x6d\x31\x06\xe2\x63\xf0\x94\x5c\x73\x06\x7c\xd4\xaa\xb3\x86\x6c\x23\x9c\x5a\x1a\xd4\xeb\x90\x50\x43\x5e\x89\x1d\xfe\xc4\xcb\x8e\xbb\xe8\x53\xce\x3c\x83\x82\x99\x56\x19\x76\x57\xf0\xb4\x00\x7f\x64\x05\xb5\xf1\x27\xb3\x51\x42\x8e\x36\x2d\x86\x76\x5e\xff\x88\xf6\xd2\x6c\xce\xd3\x77\x6f\x10\xc7\x23\xbb\x28\xc1\x8c\x5d\x68\x26\x0d\x8f\x67\xdd\x0f\x48\x98\x6f\x77\x1e\x6a\xcf\x39\xfc\xe9\xf8\x54\x69\x8d\xa6\x22\x0b\x0d\x86\x7c\x73\xd4\xd7\xd8\xe8\x54\x69\xc1\xe4\x0a\x9b\x55\xb2\x71\x84\xb1\xb4\x3b\x92\x56\xe2\x18\x23\x63\x16\xa7\x24\xdc\x00\x04\xfe\x14\xc6\x11\xea\x6f\x1e\xd8\xa3\x3a\x58\x35\xa4\xfd\x96\xea\x7e\xeb\xe4\x0f\x55\x3d\x7c\x49\x71\x80\xce\xe1\x0b\x0b\x7f\x46\xb9\xa8\x4b\x26\x5d\x6a\x72\x7b\x5b\x2d\xc2\x18\x71\x19\x5a\xc6\xc5\x50\xeb\x18\x88\x9c\xf2\xb6\xf1\xa3\x33\x48\x55\x59\x09\x2c\xc3\x79\x67\x8d\xcc\x3c\x52\x7b\xff\xe8\x01\x6a\xcd\x1d\xa1\xd7\x6a\xa9\x39\xe6\x50\xb2\xb4\xe0\x12\x37\xda\xe1\x7d\x25\x98\xf4\x69\xae\x6f\x7f\x3e\x08\xe4\x4f\xbc\x79\x8b\x4d\xcc\xb6\x6e\x8f\xd2\x62\x37\x43\x0f\x68\x11\x12\x74\x80\xb5\x11\xe3\x0c\x94\x74\xfe\xf4\x64\xb2\xd0\x35\x4e\xce\x60\xf2\x92\x09\x83\x93\xfe\x91\x39\xc0\xe4\xa3\xbc\x95\xea\x4e\x4e\x7a\x0f\x47\x1e\xe0\x73\x7d\x87\x2a\xfc\xdf\x14\x4e\x49\x86\xd3\xa1\x9b\x4e\xb0\xa1\xbb\x41\xac\x9e\xbb\x4e\xa2\x03\x2a\xbc\x87\x0a\xc7\xe0\xf1\x63\x0f\x9f\xd7\x09\xa6\xf8\xb5\xd1\x30\x52\x73\x25\x04\x66\xff\xc7\xd2\xdb\x6f\x02\x56\x5b\xa0\x01\x92\x28\xe3\xd0\xed\x46\xc0\x0e\x81\x5f\x1e\x31\x7b\xe5\x8f\x5c\xec\x3b\x0c\xf4\x61\x87\xbc\xd9\x50\x55\xc6\x7d\xbe\x81\xd2\xc6\xf3\x1b\x2e\xfe\xc3\x03\x3b\x52\x2d\x1f\x3a\xbb\xa9\xbf\xcf\x2c\xfa\x6b\x4f\xaf\x30\x03\xc8\xdd\x89\xd7\x78\x2e\x85\xea\xb0\xe6\x38\xca\x11\xe7\x49\x5a\x85\xd2\x21\x92\x84\x88\x0d\xb2\x84\x05\x80\x19\x58\xf1\x35\x4a\xc2\xca\x15\xa7\x61\x43\x64\xb0\xf1\x2e\x99\x64\x2b\x7f\x6a\x26\x7c\x8d\x73\xe4\x4e\x46\x3c\x81\xb4\x47\x64\x4f\xd4\x1a\x23\xbf\xe2\x96\x8a\xa5\x82\xd0\x6a\x9d\x0d\xdb\xd4\x4e\x69\xad\x35\x4a\xdb\xd3\x76\x66\xae\x20\x3e\x06\xdb\xb0\xcb\x39\x53\xb5\x1c\x2f\xf2\xe7\x6d\x4a\xf7\x69\x91\x8e\x87\x62\x37\x67\xae\xe2\x81\x17\x66\xa9\x39\x76\x1f\xc9\x65\x3d\x05\x21\xb7\xfe\xcb\xb7\x54\xbb\xf5\x08\xb3\x9d\xcd\x64\xc8\x19\x17\xb5\x46\xd7\x02\xf8\xd3\x3f\x91\x64\x53\x8a\xef\xf0\x6d\x76\x65\x74\x03\x7e\xa8\x79\xbe\x26\x22\xe2\x21\x4a\x58\x3f\xdb\xfc\x0a\x1f\x7b\xfa\x2f\xf6\xdc\x0d\xf0\x9f\x28\x66\x17\x60\x75\xed\xdd\x29\x9c\x72\x0e\x57\x36\x2b\x0e\x4b\x53\xac\x2c\x66\xef\xb7\x3f\xde\x3b\xf5\x49\x3a\x7e\x95\xe7\x7e\xb6\xca\x75\xf8\xe5\xd7\x13\xcf\x15\xb3\x4f\x51\x0e\xba\xf8\x8f\x00\x00\x00\xff\xff\xe1\x4e\x40\xaf\x14\x3b\x00\x00"),
		},
		"/deployment.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "deployment.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 6565,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\xdd\x8f\xe3\xb6\x11\x7f\xdf\xbf\x62\xb0\x79\x28\x50\x58\xf2\x5e\x36\x09\x50\x1d\xee\xe1\xda\x34\xb9\x03\x72\xdb\x45\xf6\x50\xa0\x4f\x31\x4d\x8d\x2d\x76\x29\x52\x25\x47\x76\xd5\x43\xfa\xb7\x17\x43\x7d\x98\xb4\xa5\xfd\x28\x10\x20\xfb\x72\x3e\x91\x9c\x8f\xdf\xcc\x6f\x66\xc8\x2c\xcb\xae\x44\xa3\xfe\x8e\xce\x2b\x6b\x0a\x10\x4d\xe3\xd7\x87\x37\x57\x8f\xca\x94\x05\x7c\x8f\x8d\xb6\x5d\x8d\x86\xae\x6a\x24\x51\x0a\x12\xc5\x15\x80\x11\x35\x16\x50\xa1\xae\x33\xdb\xa0\x13\x64\xdd\xf0\xd5\x37\x42\x62\x01\x5f\xbe\x40\x7e\x37\xfe\x17\x7e\xfd\xf5\xca\x37\x28\xf9\xa4\xc3\x46\x2b\x29\x7c\x01\x6f\xae\x00\x3c\x6a\x94\x64\x1d\xaf\x00\xd4\x82\x64\xf5\x93\xd8\xa2\xf6\xfd\x87\x25\x45\x9e\x9c\x20\xdc\x77\xfd\x2e\xea\x1a\x2c\xe0\x67\x94\x0e\x05\xe1\x15\x00\x61\xdd\x68\x41\x38\x48\x8d\xec\xe6\x3f\x9d\x28\x58\x52\xc1\x7f\xc2\x18\x4b\x82\x94\x35\xd1\xf6\xc6\xd9\x1a\xa9\xc2\xd6\xe7\xca\xae\xbd\x74\x82\xb5\x5f\x93\x6b\xf1\x3a\x6c\x1a\x3d\x0d\xbf\xd1\x1d\x94\xc4\xf7\x52\xda\xd6\xd0\xdd\xa2\xa6\x83\xd5\x6d\x8d\x93\x96\xaf\xc6\x7f\xe1\x1f\xb6\x85\xa3\xd2\x1a\x0c\x62\x09\x54\xa1\x47\xa0\xa3\x1d\x0f\x80\xda\x41\xc7\x5b\x84\x21\x20\x0b\xe8\x49\x6c\xb5\xf2\x15\x1c\x84\x56\xa5\x20\x2c\xe1\xf3\x4f\x0f\x93\x38\x69\x8d\x41\x19\x3c\x02\xb1\x17\xca\x78\x82\xcf\x4a\x6b\x74\xf9\x85\xea\x6c\x40\x86\xc2\x7a\x46\xda\x67\x52\x4c\xab\xc0\xb2\x76\x6a\xff\x49\x34\x45\xf4\x11\xe6\x0e\x65\xfd\xd6\x64\x5b\x89\x3b\xd1\x6a\xfa\x64\x4b\x2c\xe0\xe6\xbb\x9b\x9b\x0b\xfd\x0f\x1c\x50\x02\xbb\x0b\x01\x86\xcd\x63\xbb\x45\x67\x90\x30\x00\x4f\xda\x6f\x2e\xe1\x9a\xb3\x19\x1d\x45\x9a\x7d\x90\x9a\x9a\xdc\x7f\xbb\x7b\xe6\xe4\xb9\xcd\xdf\xcc\xd8\xfc\xb9\x42\xd8\x59\xad\xed\x51\x99\xfd\x10\x24\x50\x1e\x76\xd6\x41\xeb\xf9\x9b\x00\xd9\x7a\xb2\xb5\xf2\x58\xc2\xe6\xd1\xd8\xa3\xf9\xa5\xb2\x9e\xfc\x06\x76\x4a\xe3\x6a\x12\x75\xac\x94\xac\xfa\xe0\x9e\xe2\x6f\xa1\xb4\x63\xcc\xf9\x14\xff\x70\x60\x8f\x06\xf6\x8a\x98\x5b\xd6\x2b\xb2\xae\x03\x27\xa8\x42\x37\x09\xa3\x4a\x98\xc1\x80\x1f\x15\x7d\x68\xb7\x60\x1d\x67\x13\x68\xf5\x88\xf9\x29\xcb\x84\xf6\x76\x52\x55\x73\xce\x82\x3a\x61\xa0\x0c\xd9\x70\x4a\x5a\x43\x42\x19\x74\x2b\xd8\xa2\xb6\xc7\xcb\xe4\x61\x89\xb5\xe8\x7a\x81\x47\x4e\x48\xb2\x4c\x9d\x83\x2a\x11\x84\x81\x8d\xf7\xd5\x2f\x7d\x5a\x9c\x3b\xce\xc5\x46\x59\xc3\xb6\xd6\xd6\x61\x6f\xbb\x35\x08\x9b\x8f\x25\x2f\x51\xf7\x83\xd2\xb8\x79\x1b\x40\xe5\x0c\x16\x46\xe2\x6a\x44\x45\x38\x9c\x24\xf5\x0e\xa7\x42\x06\xf7\x4f\x50\xe5\x70\xf7\xe7\x22\x78\x85\x86\x5c\x07\x8f\xd8\x81\xaf\x6c\xab\x4b\xd8\x9e\x44\x5d\xf7\xb6\x5e\x0f\xc0\xf6\xf2\xae\x4f\x4e\x5c\xb3\xfe\x00\x18\x96\xa0\x0c\xfc\x77\x9d\x7b\x5f\xad\x97\x59\xe5\x7d\x55\x2a\xf7\x52\x3a\xed\x74\xfb\xef\xcc\xfb\xea\x79\x26\x71\x56\x7e\xf9\x92\xb1\x39\xf9\xc3\xc3\x87\x87\x29\xb5\xb9\x08\x4f\xd4\x7a\xf8\x00\x8d\x53\x07\x41\x18\xfc\x25\x0b\x42\x4a\xf4\x9e\xe1\x39\x61\xa3\xd0\xe7\x53\xdd\x1b\x0d\xdf\x2b\xca\x1e\xb1\x9b\xbe\x9f\x53\x0a\x2e\x28\xc5\xbd\x60\xc1\x14\x98\xf7\x20\xa4\x3f\x9a\x09\x50\x87\xa2\xcc\xac\xd1\xdd\x0a\x8e\x08\x47\x6b\xfe\x40\xb0\x45\x10\x5b\x8d\x6c\xbb\xac\x6a\x5b\x06\xaf\x51\xfb\xd8\xd1\xc5\x32\xaa\xfc\x44\xce\x29\x6d\x26\x82\x9e\x51\x89\x2a\x71\xa2\x00\x9f\xf7\x9c\xc0\x0c\x21\x43\xc7\x49\xd8\x63\xf7\x16\x30\xdf\xe7\x2b\x10\x63\x8e\x95\xa1\x77\xf2\xae\x1c\x3e\xee\x26\x11\xa9\xba\x7f\xb6\x9e\x42\x62\xfa\x56\x56\x91\xda\x55\x48\xc9\x01\x99\x84\x2e\x93\x20\xa1\x19\x97\x0e\x1a\xab\x0c\x79\x10\x04\x9b\x35\x92\x5c\x73\xb2\x94\x6b\x4e\x3f\x35\x10\x66\x03\xc2\x83\x18\x2d\x61\x0b\x4e\x65\x66\xe8\x1c\xad\xc7\x33\xa6\x3c\x62\xb7\xba\xac\x3f\x23\x87\xc7\xc2\x33\x09\x4a\x18\x2d\xb6\xf6\x80\x2b\x38\x2a\xaa\x02\xdb\x13\xe6\x0e\x44\x0b\xdd\x9e\x21\x40\x21\xab\x49\x0c\x63\xaa\x4c\x70\xbe\xcf\xa1\xb1\x1e\x60\x09\x15\x3a\x5c\x66\x54\x9a\x98\x2f\xa9\xf5\x81\x55\x7c\xac\x8f\xd4\xd3\xac\xfa\xff\x73\xd2\x94\xaf\x4b\x49\xb2\xa0\xea\xc6\x3a\x02\x61\xba\xa1\x5d\xc0\x07\xd4\x75\xc2\xcc\x53\x1a\x98\x30\x18\x28\x07\xd2\x61\xc0\x59\x68\x4e\xd1\x72\x6d\x1d\x70\x0f\x53\x3b\x25\x99\xe9\x7d\x70\x5a\xd7\x4f\x34\x67\x55\x3f\xc9\xf0\xb8\xfa\xcf\x54\xfd\xa0\xd2\x23\x85\xaf\x9b\x2c\x0b\x13\xcd\x29\x75\xb3\xde\xfa\x28\x51\xdd\xbe\xe5\x28\x2e\x76\x8b\x31\x84\xb1\x7f\x59\x27\x6a\xfd\xda\x60\xa6\x96\xc4\x28\xcd\x57\xca\x7e\x75\x72\x6c\x1a\xc0\xb2\x27\xa6\xc2\xa1\xcf\x3b\x0c\x2c\x36\x16\xae\x0b\x1e\x36\x3d\x5d\x83\xaa\xc5\x1e\xfb\x6e\x9f\x9c\xcc\xe1\x07\xd5\x47\x09\x6a\xee\xda\x0e\x25\x4f\xd4\x27\x79\x0e\x35\x0a\x8f\xdc\x95\x83\x0c\x38\xf4\xe3\x38\x93\xba\x22\x6a\x7c\xb1\x5e\x57\xed\x36\x2f\xad\x7c\x44\x97\x4b\x5b\xaf\x5d\xe0\xb9\x2c\xd7\x89\xa6\x35\x89\xbd\x8f\x04\x73\xa4\x78\xe4\xe6\x29\x9c\xd5\x93\xd8\x27\x34\x82\x5e\x5f\x01\x83\x64\x65\x67\xc5\x16\x6f\xf2\x9b\xfc\x26\x73\xf2\x4f\xe9\xb9\xfb\x56\xeb\x7b\xab\x95\xec\x0a\xf8\xb8\xbb\xb3\x74\xef\xd0\xc7\xae\x71\x22\x44\xc3\xf3\x84\x2b\x51\x13\xb5\x80\x29\x00\xf7\xd6\x51\x01\xb7\x37\xb7\x37\xd3\xaa\x56\x07\x34\xe8\xfd\xbd\xb3\x5b\x8c\x9b\x0c\xcb\xf8\xf1\xbc\xef\x34\x97\x02\xc2\x67\x41\x55\x01\xeb\x0a\x85\xa6\xea\x3f\xd1\x92\x32\x8a\xf9\xf2\x3d\x6a\xd1\x3d\xa0\xb4\xa6\x1c\x6e\x26\xe3\x1f\xa9\x1a\x6d\x4b\xd3\xda\xb7\xd3\x1a\x17\x00\xf5\x3b\xb5\xcc\xdb\xd6\x49\xf4\xb1\x05\x0e\xff\xd5\xa2\x27\x9f\x5a\x25\x9b\xb6\x80\x6f\x6f\xea\xe4\x63\x8d\xb5\x75\x5d\x01\xdf\x7d\xf3\x49\x4d\x0b\x7d\x71\xfa\xc4\x55\x21\x92\xf1\x55\x94\x6b\x1f\x8d\xd4\x6d\x89\x7d\x2d\x1b\xfa\x6a\x5a\x4d\x9e\x18\x7e\xad\x9b\x6b\x72\x2c\x96\xa7\xc3\xb7\xa7\x46\x14\x8d\xa8\x5c\x7d\xfa\x39\x6a\x33\xd6\xce\xd0\x78\xf2\x59\xf3\x66\x27\xaf\xbe\x34\x04\xeb\xee\xfb\x48\x38\x6b\x29\x0c\x6f\xc9\x0e\x0e\xf6\xdf\x8c\xee\x0a\xe0\x8b\xde\x73\x13\xd6\x78\x15\x98\x99\xb3\xb8\xcf\x72\x7d\xe0\x41\xab\x67\x57\x28\xc6\x11\xfe\xcb\x83\x56\x6a\x68\xd2\xea\x67\xa6\x9f\xe7\x83\xf3\xcc\xd0\x93\x0c\x3b\x91\xa8\xd8\xa3\x78\xfc\x59\x0a\xd1\xe0\xc8\x6b\x62\x74\xee\xfb\x79\x90\x66\x7c\x8f\xbb\xec\xf3\xae\x8f\x63\x4f\xd4\x68\x2f\x3a\xec\xa9\xb3\x46\xc2\x5e\xd5\x63\x97\x30\x89\xe4\x6d\x2e\x7a\x5e\x8a\xd3\x6a\xb4\xe2\xf9\x46\x1b\xb7\xda\xb4\xd3\xce\x83\xbc\xd4\x6d\x97\x38\xc1\xda\xd7\x27\xed\xeb\xe4\x9e\x90\x9e\x9f\x45\x1f\x3d\xbe\xf0\xbd\x22\x7e\xa7\x88\x9d\x4b\x5e\x2c\x96\x13\xee\xec\x0e\xbf\x61\x08\x63\xc8\x93\xc7\x89\xcd\xf4\x96\x52\xa2\xd4\xc2\x61\x39\x03\x7c\x74\xe2\x8f\x73\x70\xfb\x97\xe0\x3d\xff\xb4\xf0\x54\x72\x33\xe2\x4f\x95\xa1\xa7\x75\x88\x17\x6a\x38\xdf\xb9\xa0\x44\xb8\x7d\xd2\x3c\x38\xb4\x70\x0c\x75\x62\x7a\xe8\xeb\x23\x03\x52\x18\x9e\x85\x77\xb6\x35\x65\x7c\x83\x9c\x80\x9c\x4e\xbc\xe3\xcb\x61\x7f\x2a\x79\x2e\x9c\xc3\xf1\x2f\xb6\x0e\xa9\x6d\xdb\x90\x3b\x2f\x4b\x9f\x85\x67\xae\xf3\x08\x25\x21\x96\x22\x04\x28\xe3\xde\xfc\xee\x12\xac\xb5\x14\xb9\x4c\xe2\x77\x76\x1e\x0d\xdf\x02\xde\x5d\x84\x28\xd9\xf4\x88\xdd\xac\x82\x35\x69\x9f\xa7\xb5\xef\xdc\xbc\x25\xdb\xc2\xd1\x27\x2d\x3b\xa0\x53\xbb\xee\x19\xcb\x5e\xeb\xfe\x6f\x5f\x6a\x7f\xf3\x4a\x6b\xca\xf1\x67\xa0\xc9\xea\xe9\x1e\xb5\x54\x88\xdf\x1d\xbe\x2e\x5e\x5a\x2b\x57\x87\xdb\x17\xef\x9d\xc6\x0d\xc6\x6d\x78\xa8\xf7\xe9\xb4\xf1\xd7\x90\x72\x65\x8f\xec\x61\xdc\x12\xba\xf8\x40\x46\x12\x6e\x8f\x4c\x90\x6d\x07\x02\x36\xbc\xf1\xe7\xfe\xd6\xb1\x49\x29\xda\x67\x6f\xd9\xfb\x38\x8a\x0a\x3c\x3d\x57\xff\xdc\xcc\x71\x46\xd8\x5e\x30\x08\xe0\x91\x43\x63\x62\xeb\xac\xa9\x91\xa8\x4b\xa3\xdf\x8e\xd7\x39\xcf\xb2\x37\x4b\x76\x1f\xbe\x5e\x1d\x6e\x37\xcb\x91\x5c\x38\x75\x1b\x0f\x15\x23\xfc\xef\xcb\x52\x71\xb6\x09\xfd\xde\xed\xcf\x02\x70\x5a\x3b\x35\x85\x2b\x00\x3e\xea\x84\xd9\xe3\x53\xa7\xb3\xf0\x44\xd6\x7f\x39\x53\x3b\xfc\xfc\x5f\x00\x00\x00\xff\xff\xec\x58\x51\x23\xa5\x19\x00\x00"),
		},
		"/rbac.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "rbac.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 701,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\xbf\x4e\x33\x31\x10\xc4\x7b\x3f\xc5\x76\x91\x3e\xc9\xf7\x29\x1d\x72\x07\x14\x34\x88\xe2\x10\x34\x88\x62\xcf\xb7\x90\x25\x3e\xaf\xb5\xb6\xaf\x20\xca\xbb\xa3\x44\x39\x89\x3f\xc9\x55\xd7\xd9\xa3\xdd\xd5\x6f\x46\x63\xad\x35\x98\xf8\x99\x34\xb3\x44\x07\xe3\xda\x6c\x39\xf6\x0e\x1e\x49\x47\xf6\x74\xed\xbd\xd4\x58\xcc\x40\x05\x7b\x2c\xe8\x0c\x40\xc0\x8e\x42\x3e\xbc\x00\x22\x0e\xe4\x60\x43\x61\xb0\x92\x48\xb1\x88\x9a\x39\x35\x27\xf4\xe4\x60\xb7\x83\xe6\x61\xfa\xc2\x7e\x6f\x7e\x73\x68\x87\xbe\xc1\x5a\x36\xa2\xfc\x89\x85\x25\x36\xdb\xab\xdc\xb0\xfc\x1f\xd7\x1d\x15\x9c\x30\x6f\x43\xcd\x85\xb4\x95\x40\x0b\x30\x6a\x0d\x74\x5c\xb2\x80\x89\xef\x54\x6a\xca\x0e\x5e\x56\xff\x56\xaf\xc7\x4b\x4a\x59\xaa\x7a\xfa\x21\x8e\xa4\xdd\x37\xc1\x42\x94\xd8\x9e\x06\x9f\xda\xfb\xcb\xb3\x0b\x78\xbe\xe1\xd8\x73\x7c\x5f\xc2\xba\x04\x6a\xe9\xed\xb0\x36\x59\x9f\x21\x32\x00\x7f\xf3\x3f\x7f\x38\xd7\xee\x83\x7c\x39\xc5\x7a\xb6\x5c\x97\x49\xe7\x4b\xf3\x15\x00\x00\xff\xff\xb4\x81\x1a\xb8\xbd\x02\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/crds.yaml.tmpl"].(os.FileInfo),
		fs["/deployment.yaml.tmpl"].(os.FileInfo),
		fs["/rbac.yaml.tmpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
