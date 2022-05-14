package cache

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

type diskCache struct {
	key      string
	filename string
}

// DiskCache primary struct with an array of diskCache (key:filename) and the primary directory
type DiskCache struct {
	data      []diskCache
	directory string
}

// New creates a new DiskCache using the tmp directory
func New() *DiskCache {
	return &DiskCache{
		directory: newTmpDir(),
	}
}

func newTmpDir() string {
	dir, err := ioutil.TempDir("dir", "prefix")
	if err != nil {
		log.Fatal(err)
	}
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(dir)
	return dir
}

func hash(key string) string {
	hasher := sha256.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func filename(key string) string {
	return "mantis_dc." + key + ".dat"
}

func (D *DiskCache) newTmpFile(key string) *os.File {
	file, err := ioutil.TempFile(D.directory, filename(key))
	if err != nil {
		log.Fatal(err)
	}
	defer func(name string) {
		_ = os.Remove(name)
	}(file.Name())
	return file
}

// Set some data to a given key
func (D *DiskCache) Set(key string, data any) error {
	key = hash(key)
	file := D.newTmpFile(key)

	marshaledData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	D.data = append(D.data, diskCache{
		key:      key,
		filename: file.Name(),
	})

	_, err = file.Write(marshaledData)
	if err != nil {
		return err
	}
	return nil
}

// Get the data pointed to by a key
func (D *DiskCache) Get(key string) (any, error) {
	key = hash(key)

	for i := range D.data {
		if D.data[i].key == key {
			pattern := filename(key)
			filename := D.directory + string(os.PathSeparator) + pattern

			fileInfo, err := os.Stat(filename)
			if fileInfo.IsDir() && err == nil {
				file, err := os.Open(filename)
				if err != nil {
					return nil, err
				}

				reader := bufio.NewReader(file)
				bytesArr, err := reader.Peek(1024)
				if err != nil {
					return nil, err
				}

				var datum any
				err = json.Unmarshal(bytesArr, datum)
				if err != nil {
					return nil, err
				}

				return datum, nil
			}
		}
	}

	return nil, errors.New("key not found")
}
