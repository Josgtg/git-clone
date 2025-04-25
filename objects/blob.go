package objects

import (
	"hash/fnv"
	"log"
	"os"
)

// Object that points to a file
type Blob struct {
	Path string
}

func (blob Blob) Name() string {
	return "Blob"
}

func (blob Blob) Hash() int {
	hasher := fnv.New128a()
	hash, err := hasher.Write([]byte(blob.Path))
	if err != nil {
		log.Printf("could not generate hash for blob pointing to \"%s\":\n%s\nreturning empty hash\n", blob.Path, err)
	}

	return hash
}

func (blob Blob) String() string {
	content, err := os.ReadFile(blob.Path)
	if err != nil {
		log.Printf("could not read blob pointing to \"%s\":\n%s\nreturning empty string\n", blob.Path, err)
	}

	return string(content)
}
