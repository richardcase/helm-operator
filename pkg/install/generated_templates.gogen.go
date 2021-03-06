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
		"/flux-helm-operator-account.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-helm-operator-account.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 948,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\x39\x6f\xdc\x30\x10\x85\x7b\xfe\x8a\x01\x5c\x38\x09\x2c\x05\xee\x02\x75\xb6\x8b\x14\x09\x52\x28\x47\x13\xa4\x18\x92\x4f\x59\xc6\x5c\x8e\x30\x24\x37\x87\xb0\xff\x3d\x90\xb4\x06\xbc\x8e\xed\x34\xdb\x8d\xe6\xd2\x9b\xc7\xaf\x69\x1a\x73\x46\x9f\x36\xa0\x0c\xdd\x05\x07\x62\xe7\xa4\xa6\x72\x41\x2e\xd6\x5c\xa0\xa4\x12\x91\x2f\x88\x93\x3f\x4a\x91\x0d\xc9\x87\xf4\x9d\x58\x61\xce\x48\x52\xfc\x4d\x09\xf0\xf0\x34\x88\xd2\xbb\x6a\xa1\x09\x05\x99\x7e\x86\xb2\x59\x46\x1a\xcb\x19\x7e\xfe\x03\x72\x26\x27\xa9\xa8\x44\x7a\xd1\x5f\x5f\xdd\xbc\x6c\x0d\x8f\xe1\x0b\x34\x07\x49\x1d\xed\x2e\xcd\x6d\x48\xbe\xa3\x8f\xab\xaa\xab\x55\x94\xd9\xa2\xb0\xe7\xc2\x9d\x21\x8a\x6c\x11\xf3\x1c\x11\x25\xde\xa2\xa3\x21\xd6\x5f\xcd\x06\x71\xdb\xc8\x08\xe5\x22\x6a\x9e\x2e\x4d\x13\x85\x81\xda\x0f\xbc\x45\x1e\xd9\x81\xf6\xfb\x43\xf7\xf2\xd9\xd1\x34\x1d\x57\xa7\x89\x90\xfc\x7e\x6f\x66\xcf\xee\x8b\x55\xcb\xae\xe5\x5a\x36\xa2\xe1\x0f\x97\x20\xa9\xbd\x7d\x93\xdb\x20\xaf\x77\x97\x16\x85\xef\x6e\xb9\x59\xdd\xeb\x25\xe2\x94\x87\x18\xad\x11\xcb\x78\x43\x3c\x86\xb7\x2a\x75\xcc\x1d\x7d\x3d\x7f\x75\xfe\x6d\xd9\xa9\xc8\x52\xd5\xe1\x28\xb9\x83\xda\x7b\x89\x86\x92\xa4\xfe\xd0\xf8\xb9\x7f\xff\x74\xef\x09\xae\xbf\x5e\xc9\x39\xad\x09\x12\xd1\x63\x98\x17\xdc\x99\xf0\x8c\x36\x43\xf4\xef\x9b\x3c\xb3\x3d\x57\xfb\x03\xae\x1c\x5c\x7e\x14\xcd\xff\x08\x7f\x88\xd6\x43\xf6\x1e\xa3\x2d\xe6\x39\xf2\x18\xb8\xc6\xb2\xe2\x37\x53\xfa\x37\x00\x00\xff\xff\xad\xec\xff\x2b\xb4\x03\x00\x00"),
		},
		"/flux-helm-release-crd.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-helm-release-crd.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 7634,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x5b\x6f\xdb\x46\x13\x7d\xd7\xaf\x98\xcf\x5f\x01\xdb\x85\xc4\xdc\x8a\xa2\x15\x10\x24\x41\x8c\x34\xa9\x73\x31\xec\x38\x2f\x86\x0b\x8c\xc8\x21\xb9\xf5\x5e\xd8\xbd\xc8\x56\x8b\xfe\xf7\x62\x97\xa4\x24\x52\x14\x45\x25\xce\x43\x91\xf2\x25\xd1\x5e\x66\xcf\xcc\x9c\x39\x7b\xf1\x64\x32\x19\x61\xc1\x3e\x91\x36\x4c\xc9\x29\x60\xc1\xe8\xce\x92\xf4\xbf\x4c\x74\xf3\x93\x89\x98\x7a\x30\x7f\x34\x23\x8b\x8f\x46\x37\x4c\x26\x53\x78\xe9\x8c\x55\xe2\x9c\x8c\x72\x3a\xa6\x13\x4a\x99\x64\x96\x29\x39\x12\x64\x31\x41\x8b\xd3\x11\x80\x44\x41\x53\xc8\x89\x0b\x4d\x9c\xd0\x90\x89\xfc\x8f\x28\xe5\xee\x2e\x4e\x22\xa6\x46\xa6\xa0\xd8\x8f\xcc\xb4\x72\x45\x39\x74\xad\xb7\xb4\x60\xfc\x00\x80\x72\xdd\xd7\xc4\xc5\x79\x69\x2c\xb4\x72\x66\xec\x69\xbb\xe7\x2d\x33\x36\xf4\x16\xdc\x69\xe4\x4d\x08\xa1\xc3\xe4\x4a\xdb\xf7\x2b\xe3\x13\xc8\xf5\x08\xc0\xc4\xaa\xa0\x29\x84\x8e\x02\x63\x4a\x46\x00\x98\x24\xc1\x33\xe4\x67\x9a\x49\x4b\xfa\xa5\xe2\x4e\xc8\xe5\xc4\x5f\x2f\x3e\xbc\x3f\x43\x9b\x4f\x21\x32\x16\xad\x33\x51\xb5\x92\xb7\x12\xc6\xd4\x81\x58\xc7\x0d\x60\x17\x7e\x29\x63\x35\x93\xd9\x2e\x53\x17\xe1\x57\xc3\x58\xa3\x69\x90\xad\x58\xc9\xd2\x13\x73\xf5\xec\xe8\x79\xe4\xe7\x3c\x7d\x7a\x50\x81\x4a\x0e\x8e\xaf\x23\x41\xc6\x60\xd6\x04\xfd\xae\xd1\xd6\xbf\x50\x9d\xfb\x28\xd6\x84\x7e\xa5\x8f\x4c\x90\xb1\x28\x8a\x86\xc9\x17\x2d\x73\x09\x5a\xdf\x60\xdc\x4c\x57\x7c\xaa\x82\x5b\x02\x9f\xc2\x5f\x7f\x8f\x00\xe6\x35\x3b\xe7\x8f\x56\xbf\x96\x59\x28\x2d\x87\xae\x30\x93\xf4\x9c\x92\x29\x58\xed\xea\xb5\x8c\x55\x1a\x33\x5a\xb6\xcd\x91\xb3\x24\xa0\x2c\x6d\xa8\x82\xe4\x8b\xb3\x37\x9f\x9e\x5c\xc4\x39\x09\x9c\x56\xd3\x0a\xad\x0a\xd2\x96\xd5\x98\x82\xa9\x8a\xb5\xf5\xa7\xe9\x0f\xc7\xb4\x5f\xef\xea\x30\xce\x51\xdb\xc3\xeb\xb5\xde\x2e\x0b\xfe\xf3\xa4\xac\x4b\xae\xd1\x01\x90\x90\x89\x35\x2b\x02\x38\xf8\x98\x53\x20\x77\xed\x33\xd8\x9c\x19\xa8\x98\x01\x16\x75\x46\xd6\x44\xf0\x26\x05\xa9\x2c\x18\x57\x14\x9c\x51\x32\x06\x66\xe1\x96\x71\x0e\x09\xa5\xe8\xb8\x05\xab\x5a\xab\x00\xcc\x1f\x47\xad\xb6\x8d\x04\xaf\xf9\x81\xd6\x92\x96\x53\x38\xf8\x6d\x7e\xf5\x78\xfc\xe4\xfa\xbb\x83\xc6\x88\x35\xda\x0f\xf3\xa7\x76\xc1\xe7\xae\x07\xff\x8c\x20\x23\x49\x1a\x2d\x25\x30\x5b\x00\xa6\x29\xbb\x63\x32\x03\x9b\xd3\x86\x47\xb2\x2e\x5d\xb0\xca\x0f\x80\x9a\x52\xe5\x2a\x9f\xe5\xed\x15\x4e\xfe\x7c\x38\xf9\xf9\xfa\xe8\x6a\x52\xfd\xef\xfb\xba\xe9\xf8\x59\x2b\x08\x65\x3a\x96\x02\xb2\x7f\x20\xc2\xb4\x8e\x68\x78\x5f\x56\xce\xd5\x81\xf1\xad\x66\xa5\x33\xab\x0f\xcd\xa6\xf7\xa5\xe9\xaf\x1f\x02\x26\x48\x39\xdb\xeb\x7a\x70\x9b\x49\x63\x91\x73\x50\x1a\x5c\x91\x69\x4c\xa8\x9e\x0b\x4c\x82\x21\x2f\x58\xa6\x13\xae\xd7\xe1\x8c\x74\xab\x2f\x55\x5a\xa0\x0d\xbd\x3f\xfe\xd0\xa2\xa6\x21\xfb\x09\xb9\x6b\xd7\x60\x0b\xd6\x9b\x74\x19\xf1\x32\xc4\x61\xa2\x97\x0a\x47\x06\x94\x0c\x35\x5b\x83\xed\x44\x36\x53\x8a\x13\xca\x51\x0b\x57\x4c\x97\xe5\xa4\xfd\x96\x0f\x33\xcb\x60\x2d\x23\x94\x6b\xe5\xb2\x1c\x12\xe2\x64\xe9\x81\xa6\xa0\xb5\xc3\xc1\x68\xc5\xf9\x0c\xe3\x9b\x36\x90\x72\x86\x9a\xfd\x4e\xb1\x6d\xf3\x60\x8b\x80\xf9\x8f\x24\xce\xf8\x86\x57\x9d\x9e\x91\x1d\x97\x5e\x15\xa4\x7d\xaa\x96\x50\x8c\xf7\xb3\xa9\x6a\x4a\x2e\x1d\x4e\x91\x71\xa7\xa9\x4d\x84\x7e\x2f\x97\x61\xdf\x0f\x59\x19\xef\x65\xd1\xb8\xc2\x6f\x4d\xdb\x42\x0e\x2c\x05\x49\x94\x84\x53\xc2\x7e\xd0\x6a\x13\x5d\xe8\xfa\x67\x26\xcc\xf8\x80\xbf\x56\xea\xa6\x23\x1b\xbd\x51\xd7\x34\x27\x69\x21\xf7\x53\x21\xd5\x4a\x80\x76\x52\x7a\x21\x4d\x9c\xaf\xfe\x65\x3e\xf6\x06\xb5\xa5\xe2\x37\xf0\xf8\xe3\xc0\x5a\x69\x7b\x89\xbe\x45\x66\x43\xfa\x51\x2e\x80\xc9\x84\xcd\x59\xe2\x90\xc3\xa9\x9b\x91\x96\x64\x7d\xd9\x15\x5e\xfb\x99\x92\xe3\x0e\xfb\x50\x6f\x6f\xc1\xda\x93\x87\x0f\xb7\xe8\x06\xec\xd0\x0e\xe8\xd5\x0f\xff\x79\xa4\xfb\x45\x3c\xf8\xe6\xa4\x65\x3c\x68\xb1\x60\x92\x09\x27\x40\x3a\x31\x23\x0d\x2a\x85\x33\x95\x18\xff\x2f\xc2\x09\x15\x5c\x2d\x04\xc9\x76\xed\x95\x1f\xea\x10\x37\x04\x4d\x98\x2c\xc2\xb9\x88\x60\x46\xa9\xd2\x04\x02\xf5\x4d\xb5\x1b\x2e\xcb\x07\x0d\x18\x17\xc7\x64\x4c\xea\xf8\x5e\xe9\x0c\x52\xf7\x8a\x71\xba\xf0\x14\xb5\xfd\x72\x79\x42\x85\xa6\xd8\xef\xcb\xff\x83\x4b\x43\x95\x4e\xbe\xd2\x4a\x44\x26\x4c\x3f\xa5\xc5\x39\xa5\x41\xe8\x09\xdb\x65\x52\x82\x40\xad\x71\xd1\xea\x61\x96\x44\x07\xbb\x7b\x24\xaa\x79\x06\xf3\x9b\x5d\xe3\x08\x56\x7e\x7d\x3a\x56\x9d\x4d\xb7\x30\x6c\xcd\x67\xbf\xb5\xfb\x9c\x85\x9d\x37\x38\x39\x06\xe1\x8c\xf5\xbb\x31\x93\xcb\x0d\x79\x6d\xaf\xae\xb6\xe2\xf6\xcd\xa5\xdb\xbb\x8e\x8d\x78\x15\xd5\x6e\xcd\xbe\xaf\x08\xf6\x87\x27\x56\x32\x65\xd9\x3b\x2c\xca\x9c\x76\x47\xaa\xd7\x3e\x0c\xcb\xd2\x6e\x28\xd0\x9b\x2d\xe8\xcb\x58\xe9\x85\xc0\xe2\x9e\x92\x06\xfd\x27\xa8\xfa\xbb\xa1\xc5\x40\xb0\xa7\xb4\xa8\x11\x2d\xb1\x7a\x65\xcb\xc8\x86\xc6\xea\x28\xe2\xe5\x7b\xdc\x90\xbe\xb2\x23\x5a\xa0\xe8\x2a\xf8\xc1\x48\x55\x51\xde\x75\x07\xc2\xad\xf5\x6e\xa5\x36\xa0\xc9\x6a\x46\x73\xe4\x75\xcc\x6b\xc8\x8c\x13\x30\x03\x52\x01\x57\x32\x23\x0d\x02\x65\x82\x56\xe9\x36\x79\xdb\x80\xb7\x6d\x3b\x00\xeb\x2a\xf3\x2f\x65\xe4\xbd\x6a\x08\x7c\x45\x3a\x96\x40\xff\xe3\xe2\x36\x2e\xd2\x9d\xbf\x2c\x21\xbf\x08\x67\xc7\xfb\x21\xa4\xd3\xfc\xb3\xf9\xe8\xf4\xd0\xc0\x5d\x9e\xbf\x6d\xc6\xe7\x1b\xcb\x5c\x78\xb2\xf1\x67\x9e\xfb\x49\x5a\x81\x36\xff\xec\xac\xf9\xc9\x03\xa3\xe6\x87\xc2\x2d\xb3\x79\x55\xa0\xe1\x6e\x1a\x7c\x81\xa3\x70\xbc\xcf\x98\x05\x4d\x85\x3a\x86\xdb\x9c\x74\x23\xb9\x3e\x84\x5c\x85\xa3\xdb\xb7\x92\x67\x25\xe9\x43\x47\x7a\x27\xcd\xd7\xbb\xe6\x29\xa7\x23\x8d\xcd\xf1\xeb\x1b\xd0\xce\xc1\x1b\x0a\xb1\x73\xc6\x3a\x33\x5b\x83\xe7\xbb\x5f\x33\x62\x25\xad\xbf\xf1\xa9\xb4\xa7\xae\xb7\x52\x3b\xac\xdd\xb6\xdf\x19\xc4\x26\xe6\x8c\xd9\xc3\x31\x6c\xab\x82\xfe\x0a\xc8\xba\xaf\x58\x2d\xbf\x7e\x61\x36\x68\x16\x45\x59\xe4\xa7\x3c\xcf\x98\xcd\xdd\x2c\x8a\x95\x98\x2a\x9d\x3d\xf0\x9c\xdf\xef\x88\x5d\x7f\xf5\x0d\xd0\x57\xce\xff\xc3\x03\x5c\x42\x29\x93\xe5\xbb\xe3\x87\x17\x17\x1d\x93\xb6\x17\x6c\x03\xf3\x99\x2f\x56\x26\x0d\x4b\xca\x47\xbb\xba\x36\x0d\xf3\x94\xae\x0a\xb4\xde\xe2\xab\x2a\x66\x5d\x97\xd8\x9d\x5e\xe8\x6d\x22\xb6\x11\xc3\x99\x46\x19\xe7\xcd\xad\x5b\xa0\xb1\x9d\x97\xe3\x9d\xeb\x9a\x1b\x56\x9c\x50\x71\x19\xde\x4c\x06\x20\xa8\xc5\x20\x51\x64\x42\xa8\xb5\x93\x70\x98\x50\x71\x58\xbf\xbb\x1c\xa1\x31\x4e\x50\xcd\x2e\x7f\x3b\x5e\xa9\x17\xf2\xf2\x2e\x9c\x3a\x9e\x32\xce\x29\x39\xee\x01\xdd\xad\x09\x4d\xde\xae\xb2\xe1\xe9\x1b\x8e\x82\x63\x38\xac\xde\xdc\xf7\x66\xf2\xca\xda\x80\x50\x54\x8f\xc0\x4b\x36\x5c\x9e\xbf\xfd\x32\xfe\x3a\xcd\x87\xf2\x77\xe0\xb5\x77\x8d\x96\xb2\xeb\xbd\x79\x00\xbc\x79\xf7\xdf\x3a\xfa\x17\xab\x26\x7d\x59\x38\x0c\x89\x39\xe9\xa1\x11\x09\x0b\x9f\x39\xce\xcb\x27\x90\x6e\xbc\xf7\x7a\x1f\x68\xe7\x7f\x86\x86\xc5\x80\xce\xe6\x70\xe4\x21\x33\x51\xf0\x40\xff\x6d\x2c\xdf\x88\xc6\x3f\x01\x00\x00\xff\xff\x11\x60\x36\x8c\xd2\x1d\x00\x00"),
		},
		"/helm-operator-deployment.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "helm-operator-deployment.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 6000,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\x4b\x93\xdb\xb8\x11\xbe\xcf\xaf\xe8\x9a\x39\xf8\x32\x24\xe5\x5a\xef\x1e\xe8\xf2\x21\x59\x67\x6d\x57\xd9\xce\x54\xc6\x95\xaa\x9c\x76\x21\xb0\x25\x22\x02\x01\x06\x68\x4a\x61\x54\x9b\xdf\x9e\x6a\x80\xe2\x43\x14\x35\x33\xce\x21\xe1\x65\x34\x00\xfa\xfd\xf5\x03\x48\x92\xe4\x46\xd4\xea\xaf\xe8\xbc\xb2\x26\x07\x51\xd7\x3e\xdb\xbf\xbe\xd9\x29\x53\xe4\xf0\x1e\x6b\x6d\xdb\x0a\x0d\xdd\x54\x48\xa2\x10\x24\xf2\x1b\x00\x23\x2a\xcc\x61\xa3\x9b\x7f\x26\x25\xea\x2a\xb1\x35\x3a\x41\xd6\x1d\x8f\xa0\x36\x90\x7e\x15\x15\xfa\x5a\x48\x84\xdf\x7f\xef\x4e\x87\x7f\x73\x38\x1e\xa7\xbb\xc7\x23\xa0\x29\xf8\x98\xaf\x51\x32\x6b\x87\xb5\x56\x52\xf8\x1c\x5e\xdf\x00\x78\xd4\x28\xc9\x3a\xde\x01\xa8\x04\xc9\xf2\xb3\x58\xa3\xf6\x71\x61\x59\x13\xa6\x25\x27\x08\xb7\x6d\x3c\x4a\x6d\x8d\x39\xfc\x05\xa5\x43\x41\x78\x03\x40\x58\xd5\x5a\x10\x76\xac\x47\xd6\xf1\xa7\x27\x52\xae\xca\xe1\x4f\x18\x63\x49\x90\xb2\x66\x44\x53\x3b\x5b\x21\x95\xd8\xf8\x54\xd9\xcc\x4b\x27\x58\x85\x5b\x72\x0d\xde\x86\x43\x27\x9b\xc3\x6f\x74\x7b\x25\xf1\x0f\x52\xda\xc6\xd0\xd7\xeb\xe2\xf6\x56\x37\x15\xfa\xbc\xf3\xf7\x9f\x8c\x58\x6b\xfc\xa6\xb4\x46\xf7\xed\xf3\x63\xf4\x7a\xfc\x92\x4e\xf3\xc0\x85\xb4\x4f\xa4\xe8\xf7\x00\xa4\x35\x1b\xb5\xfd\x22\xea\x7c\xb4\x38\x37\x36\xd2\x25\xf1\xf4\xe4\x64\x81\x1b\xd1\x68\xfa\x62\x0b\xcc\x61\xf5\xd3\x6a\x75\x45\x30\x3a\xf2\x23\x62\xcf\xa1\xa0\xa9\xe0\xb8\x16\x8d\xef\x6c\xeb\xad\xfa\x19\x1d\x3d\xf6\xfb\x11\x3b\xd7\x77\x51\x7b\xfe\x75\x66\x07\x3a\x1a\x50\xb7\x6c\xcb\x9b\xd5\x6a\x60\xd1\x9d\xbb\x3b\xfd\x85\xbf\xd9\x06\x0e\x4a\x6b\x30\x88\x05\x50\x89\x1e\x81\x0e\xf6\x14\x18\xd6\xbc\xe5\x23\xc2\x10\x90\x05\xf4\x24\xd6\x5a\xf9\x12\xf6\x42\xab\x42\x10\x16\xf0\xed\xf3\x63\xcf\x4e\x5a\x63\x50\x06\xf8\x80\xd8\x0a\x65\x3c\x41\x34\x6d\x26\x79\x39\xa0\x77\x97\x02\x7a\xf7\xec\x80\xde\x5d\x0d\xe8\x1d\x44\xef\x86\x3c\x82\x5d\xb3\x46\x67\x90\x30\x20\x9b\xb4\x9f\xa9\x37\x77\xfa\x48\xcc\x34\xf4\x77\xff\xcb\xd0\x5f\xb2\xfa\xcd\x60\xf5\xf1\x88\xa6\x18\x1d\xfe\x56\x22\x6c\xac\xd6\xf6\xa0\xcc\xb6\x8b\x36\x28\x0f\x1b\xeb\xa0\xf1\xbc\x26\x40\x36\x9e\x6c\xa5\x3c\x16\xb0\x33\xf6\x60\x7e\x2d\xad\x27\x0f\x1b\xa5\xf1\xbe\x67\x74\x28\x95\x2c\x23\x46\x06\x18\x59\x28\xec\x09\x3a\x4c\xc4\x3f\x1c\xd8\x83\x81\xad\x22\xae\x8c\x16\x9c\xa0\x72\x40\x05\x50\x29\x4c\x27\x78\xab\xa8\x6c\xd6\x60\x1d\xc3\x11\xb4\xda\x61\xca\x30\x7d\xa5\x35\x08\xed\x6d\x2f\xa2\xe2\xfa\x02\x6a\x88\x87\x32\x64\x03\x8d\xb4\x86\x84\x32\xe8\xee\x61\x8d\xda\x1e\xd2\x8b\xb0\xaf\x44\x1b\x19\x1e\x18\xcf\x64\xb9\xcc\xed\x55\x81\x20\x0c\x78\x5f\xfe\x1a\x41\x75\x66\x2e\x77\x10\x65\x0d\xeb\x59\x59\x87\x51\x6f\x6b\x10\x7e\xfb\x54\xf0\x16\xb5\xbf\x28\x8d\xbf\xbd\x0d\x8e\x64\xf8\x0b\x23\xf1\xbe\xf3\xc5\x2b\x87\x3d\xa3\x68\xeb\x94\xc7\x07\x45\x1f\x9b\x75\xf0\x4f\x0a\x5f\xff\x18\x6c\x41\x43\xae\x85\x1d\xb6\xe0\x4b\xdb\xe8\x02\xd6\x03\x8f\xdb\xa8\xe2\x6d\xe7\xcc\xc8\xe8\x76\xd0\xfd\x96\xe5\x06\x37\x61\x01\xca\xc0\xbf\xb3\xd4\xfb\x32\x9b\xbb\xe3\x04\x76\xef\xcb\x42\xb9\x17\xa5\xa1\xf7\xe5\xd3\xe9\x17\x6b\x10\xa7\xc2\xe3\xe3\xc7\x09\xc4\x6f\x86\xb4\x7c\xfc\x18\xcc\x24\x0b\x42\x4a\xf4\x3e\x98\xff\xa1\xc3\x8b\x57\x64\x5d\x3b\x2b\xca\x5b\x45\xc9\x0e\xdb\x97\x55\xe3\xb9\x12\xe3\xc3\x33\xcd\x03\xc8\xd1\xf4\x8e\x74\x28\x8a\xc4\x1a\xdd\xde\xc3\x01\xe1\x60\xcd\x2b\x82\x35\x02\x77\x2e\x56\x5e\x96\x95\x2d\x6e\x5e\x50\x72\x95\xef\xf3\xef\x84\x92\x3e\x05\xfb\x74\xa1\x52\x0c\x40\x67\x42\xcf\x30\x3d\xf9\x8c\xc1\x16\x9d\xf6\x16\x30\xdd\xa6\xf7\x20\x4e\x60\x2a\xc2\xe0\xc3\xa7\x52\xf8\xb4\xe9\x59\x4c\xe4\xfc\xbd\xf1\x14\x00\xe8\x1b\x59\x06\x79\xf7\xc1\xf9\x9d\x2b\x46\xd9\xd0\xd3\x0b\xcd\x6e\x68\xa1\xb6\xca\x90\x07\x41\x90\x21\xc9\x8c\x21\x51\x64\x0c\x32\xd5\xa5\x03\x08\x0f\xe2\x24\x9e\xc5\x0e\x95\xa3\xeb\x29\x8d\xc7\xb3\x3c\xd8\x61\x7b\x1f\x34\x1c\x15\x94\x53\x72\x9e\x2a\x49\xcf\x66\x94\xaa\x62\x6d\xf7\x78\x0f\x07\x45\x25\x7b\x67\x9a\x92\x5d\x26\x85\xd1\x8b\x8d\x46\x21\xcb\x9e\x09\x3b\x51\x99\x60\x74\x04\xcb\x29\xd1\xb1\x80\x12\x1d\x2e\xa7\xcc\x14\x81\xcf\x69\x0a\x21\x6d\x98\x2c\x86\xe6\x7a\xda\x7c\x17\xf8\x96\x6b\xfe\x59\x77\x17\x0e\x03\x76\x02\xf7\x88\xb8\x3e\xd9\x14\xfa\xb4\x15\x95\x3e\x2b\x80\xc2\x14\x5d\x2c\xba\x26\x21\x24\x23\x45\xb9\x30\xde\xb6\x69\xe8\x2c\x5a\x10\xa1\xe3\x7e\xc2\xe1\x43\x2e\x5a\x52\x34\x7e\xa8\x5c\xbd\x40\x0a\x7d\xc8\x55\xe8\x62\x4e\x54\x62\x87\xb1\x86\x33\xdf\x6c\x60\x3c\x58\xbe\x1c\x8b\xb1\xee\x09\xeb\xfe\xd2\xa8\x84\x0e\x3b\xe6\x72\x5d\x44\xd0\x71\xc4\x11\xab\x9a\xda\xf7\xca\xe5\x70\xec\x0b\x5b\xdf\x8b\xfa\x79\x7a\x3e\x59\x9c\x0d\xc5\x5d\xac\x1c\x86\xf8\x18\x0b\xb7\x39\xcf\xf7\x9e\x6e\x41\x55\x62\x8b\xb1\x4b\x4f\x28\x53\xf8\x45\x99\x30\xbf\x41\xc5\xfd\xd6\xa1\xe4\xab\xce\xc0\xcf\xa1\x46\xe1\x91\xbb\x6a\xe0\x01\xfb\x78\x4f\xe2\xcc\x2d\x89\x6a\x9f\x67\x59\xd9\xac\xd3\xc2\xca\x1d\xba\x54\xda\x2a\x73\xd9\x01\xc5\x1e\x0f\xd6\xed\x7c\x36\x91\x96\x91\xd8\xfa\x11\x73\xc6\x04\x5f\x77\xf8\x2a\xc4\x2a\x90\xd8\x4e\xb2\x06\xa2\xcc\x1c\x3a\xee\xca\x86\x42\x21\x8b\x29\xdb\xfc\x75\xba\x4a\x57\x89\x93\x3f\x4e\xe9\x1e\x1a\xad\x1f\xac\x56\xb2\xcd\xe1\xd3\xe6\xab\xa5\x07\x87\x7e\x6c\x5e\x6d\x1d\x8d\xae\x2b\xfd\x64\x49\x54\x4f\xef\x08\x31\x12\x0f\xd6\x51\x0e\x3f\xac\x7e\x18\x86\x7c\xad\xf6\x68\xd0\xfb\x07\x67\xd7\x38\x6e\x1e\xcc\xe3\xc3\x79\x3f\xa9\xe7\x0c\xc2\xb2\xa0\x32\x87\xac\x44\xa1\xa9\xfc\xd7\x68\x4b\x19\x45\x4a\xe8\xf7\xa8\x45\xfb\x88\xd2\x9a\xa2\xbb\x15\x9e\x3e\x52\x15\xda\x86\xfa\xbd\xc1\x01\x8c\x7a\xf5\x7f\xaa\x99\xb7\x8d\x93\xe8\xc7\x1a\x38\xfc\x47\x83\x9e\xfc\x54\x2b\x59\x37\x39\xfc\xb8\xaa\x26\x8b\x15\x56\xd6\xb5\x39\xfc\xf4\xe6\x8b\xea\x37\x62\x61\xfa\xc2\xd5\x61\xc4\xe3\x0e\x3e\x19\xa9\x9b\x02\x63\xb7\xec\x06\xcb\xe9\x1c\xb8\x38\xae\x5a\x37\xef\x5f\xcc\x92\xcb\xda\xdb\x53\x97\x19\x0d\x96\x25\x9e\xda\x71\x81\x52\x0b\x87\x45\xec\x2b\xe9\x88\xf6\xe2\xbc\x14\x6b\x40\xd0\xe6\x21\xfa\xdb\x59\x4b\x61\xe4\x9a\x9c\xe0\x90\xfe\xd9\xe8\x36\x07\xbe\x40\x3f\x31\x17\xc1\xd5\x61\x67\x2a\x6e\xd2\x80\xe7\x13\xc8\x72\xd7\x9a\x6b\x3e\x61\x35\xbf\x65\x3e\x5d\x72\xe7\x2c\xf7\xc2\x75\x2c\x39\xe9\xb3\x0b\x43\xdd\x33\xaa\xec\x0b\xd8\x66\x81\x6e\xf9\x61\xe1\x82\x83\x17\xaf\xf8\x4b\xce\x61\x82\xc9\xb1\x69\x6c\xaf\x8a\x10\xcf\xe4\x7f\x7e\xf2\x4c\xc4\xd5\x30\x2f\x18\xb4\x1c\xed\x89\x41\x73\xb0\x5e\x15\x21\x9e\xc9\xff\xfc\xe4\x2c\x1f\xce\xc0\x26\xdc\x76\x52\x0b\x3e\xda\x03\xa7\xfd\x86\x1b\xde\xe4\x6d\x81\xbd\x9c\x24\x14\x96\x92\xfe\xad\xee\xdd\xe4\x12\x7e\xf6\x66\x77\x79\xb5\x73\xe8\xae\x59\x63\xe2\x5b\x4f\x58\xf5\x5a\x3d\xfd\x50\x75\x07\xbc\x10\xcb\x4d\xe3\xc2\x53\xda\x25\xfd\x4e\x2f\x17\xe8\x28\xe1\xf2\xfc\x6e\xee\xa5\x4c\x8a\x54\x3a\x5a\xa2\xc6\xa0\xc3\xbb\x49\x5c\xce\x8e\xec\xb0\xbd\xc8\x3c\x23\xed\xd3\x71\xfe\x9f\x2b\xb6\xa4\x55\x20\xbc\xa2\xd3\x1e\x9d\xda\xb4\x57\x75\x7a\x9e\xd1\x17\x61\xfd\xb3\xad\x78\x2c\x07\xdb\x84\x8b\xc3\xf2\x7b\xd4\xd3\xef\x50\x11\xc2\xff\x4d\x34\x66\xf4\x97\xe2\x31\x3b\xf4\xec\x88\xcc\xd5\x7b\x76\x4c\x66\xa4\x97\xa2\xf2\x9d\xe6\x0f\xe9\xf9\x9f\x00\x00\x00\xff\xff\x4a\xc7\xa9\xb1\x70\x17\x00\x00"),
		},
		"/tiller-ca-cert-configmap.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "tiller-ca-cert-configmap.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 226,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xc1\x0a\xc2\x30\x0c\x40\xef\xfd\x8a\xfc\x40\x07\x82\xa7\xde\xa4\x78\x53\x2f\x0e\xef\xb1\xcd\xb4\xd8\x66\xa3\x8b\x22\xd4\xfe\xbb\x6c\x4c\xd1\x63\x78\x2f\x2f\x29\x05\x42\x07\xcd\x96\xf1\x1c\xa9\x0d\x31\x52\x6e\x77\x47\xa8\x55\x6b\xad\x70\x08\x27\xca\x63\xe8\xd9\xc0\x63\xa5\x6e\x81\xbd\x01\xdb\x73\x17\x2e\x7b\x1c\x54\x22\x41\x8f\x82\x46\x01\x30\x26\x32\xd0\xc5\xfb\x53\x5f\x29\x26\x2d\x71\xd4\x0e\xb5\x9b\xe5\xe5\xc8\x01\x13\x8d\x03\x3a\x82\x5a\x97\x95\x79\x34\x50\xca\x3f\x2d\x05\x88\xfd\xa4\x7d\xfa\x0e\x1b\x97\xc5\xc0\x4b\x4d\x31\xf6\xc4\x02\x6b\x68\xbe\x1f\xdb\x8d\xa5\x2c\xb6\x67\x99\xc8\x6f\xe1\x1d\x00\x00\xff\xff\x60\xd7\x1c\xac\xe2\x00\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/flux-helm-operator-account.yaml.tmpl"].(os.FileInfo),
		fs["/flux-helm-release-crd.yaml.tmpl"].(os.FileInfo),
		fs["/helm-operator-deployment.yaml.tmpl"].(os.FileInfo),
		fs["/tiller-ca-cert-configmap.yaml.tmpl"].(os.FileInfo),
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
