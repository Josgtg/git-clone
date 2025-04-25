package objects

import (
	"fmt"
	"hash/fnv"
	"log"
	"strings"
)

// Object that points to either blobs or another trees, building the working tree of the repository
type Tree struct {
	objects []Object
}

func (tree Tree) Name() string {
	return "Tree"
}

func (tree Tree) Hash() int {
	hasher := fnv.New128a()
	var hash int
	var err error
	for _, o := range tree.objects {
		hash, err = hasher.Write(fmt.Appendf([]byte{}, "%d%d", hash, o.Hash()))
		if err != nil {
			log.Printf("could not generate hash for object in tree, returning empty hash")
			return 0
		}
	}

	return hash
}

func (tree Tree) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Tree: %d", tree.Hash()))
	for _, o := range tree.objects {
		builder.WriteString(fmt.Sprintf("\n%d\t%s", o.Hash(), o.Name()))
	}
	return builder.String()
}

