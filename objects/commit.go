package objects

import (
	"fmt"
	"hash/fnv"
	"log"
	"strconv"
	"strings"
	"time"
)

type Commit struct {
	Date time.Time

	Author string
	Committer string

	// Root tree of the repository
	Tree *Tree
	Parent *Commit
	
	Message string
}

func (commit Commit) Name() string {
	return "Commit"
}

func (commit Commit) Hash() int {
	commitBuilder := strings.Builder{}
	commitBuilder.WriteString(commit.Date.String())
	commitBuilder.WriteString(commit.Author)
	commitBuilder.WriteString(commit.Committer)
	commitBuilder.WriteString(strconv.Itoa(commit.Tree.Hash()))
	commitBuilder.WriteString(strconv.Itoa(commit.Parent.Hash()))
	commitBuilder.WriteString(commit.Message)

	hasher := fnv.New128a()
	hash, err := hasher.Write([]byte(commitBuilder.String()))
	if err != nil {
		log.Printf("could not hash commit with message \"%s\":\n%s\nreturning empty hash\n", commit.Message, err)
	}

	return hash
}

func (commit Commit) String() string {
	commitBuilder := strings.Builder{}
	commitBuilder.WriteString(fmt.Sprintf("Commit: %d\n\n", commit.Hash())) 
	commitBuilder.WriteString(fmt.Sprintf("Date: %s\n", commit.Date.String())) 
	commitBuilder.WriteString(fmt.Sprintf("Author: %s\n", commit.Author)) 
	commitBuilder.WriteString(fmt.Sprintf("Committer: %s\n", commit.Committer)) 
	commitBuilder.WriteString(fmt.Sprintf("Tree: %d\n", commit.Tree.Hash())) 
	commitBuilder.WriteString(fmt.Sprintf("Parent: %d\n", commit.Parent.Hash())) 
	commitBuilder.WriteString(fmt.Sprintf("Message: %s", commit.Message)) 

	return commitBuilder.String()
}

