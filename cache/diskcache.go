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

type DiskCache struct {
	data      []diskCache
	directory string
}

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
		err := os.RemoveAll(path)
		if err != nil {

		}
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
		err := os.Remove(name)
		if err != nil {

		}
	}(file.Name())
	return file
}

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

func (D *DiskCache) Get(key string) (any, error) {
	key = hash(key)

	for i := range D.data {
		if D.data[i].key == key {
			pattern := filename(key)
			filename := D.directory + string(os.PathSeparator) + pattern

			fileInfo, err := os.Stat(filename)
			if fileInfo.IsDir() && (err == nil || !os.IsNotExist(err) || os.IsExist(err)) {
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
			} else {
				return nil, errors.New("cache file " + filename + " not found")
			}
		}
	}

	return nil, errors.New("key not found")
}
