package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	name     string
	file     bool
	size     int
	parent   *Node
	children map[string]Node
}

func newNode(name string, options ...func(*Node)) *Node {
	n := Node{name: name, children: make(map[string]Node)}
	for _, option := range options {
		option(&n)
	}
	return &n
}

func nodeIsFile(node *Node) {
	node.file = true
}

func nodeSize(size int) func(*Node) {
	return func(n *Node) {
		n.size = size
	}
}

func nodeParent(parent Node) func(*Node) {
	return func(n *Node) {
		n.parent = &parent
	}
}

func (n *Node) String() string {
	var node strings.Builder
	var nodeParent string
	nodeName := n.name
	nodeIsFile := strconv.FormatBool(n.file)
	nodeSize := n.size
	if n.parent != nil {
		nodeParent = (*n.parent).name
	}
	node.WriteString("{")
	node.WriteString(fmt.Sprintf("Name: %s | File: %s | Size: %d | Parent: %s | Children: [",
		nodeName, nodeIsFile, nodeSize, nodeParent))
	for k := range n.children {
		node.WriteString(k + ", ")
	}
	node.WriteString("]}")
	return node.String()
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	input := readInputFile(file)
	file.Close()
	filesystem := parseInput(input)
	fmt.Printf("Filesystem root: %s\n", filesystem.String())
}

func readInputFile(F *os.File) []byte {
	var input bytes.Buffer
	scanner := bufio.NewScanner(F)
	for scanner.Scan() {
		input.Write(scanner.Bytes())
		input.Write([]byte{'\n'})
	}
	return input.Bytes()
}

func parseInput(input []byte) Node {
	// inputLines := bytes.Split(input, []byte{'\n'})
	// for _, line := range inputLines {
	// 	if line[0] == '$' {
	// 		// TODO: Parse line as command
	// 	} else {
	// 		// TODO: Parse line as command output
	// 	}
	//
	return *newNode("/")
}
