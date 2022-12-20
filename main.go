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

const CMD_MARKER = "$"

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
	// Remove the last newline added since this will create an extra blank line
	// during the parsing stage if left in
	input.UnreadByte()
	return input.Bytes()
}

func parseInput(input []byte) *Node {
	inputLines := bytes.Split(input, []byte{'\n'})
	var filesystem *Node
	filesystem = newNode("/")
	for _, line := range inputLines[:len(inputLines)-1] {
		command, args := parseLine(line)
		command(filesystem, args...)
	}
	cd(filesystem, "/")
	return filesystem
}

func parseLine(line []byte) (func(*Node, ...string), []string) {
	var command func(*Node, ...string)
	var args []string
	lineSplit := strings.Split(string(line), " ")
	if lineSplit[0] == CMD_MARKER {
		if lineSplit[1] == "cd" {
			command = cd
		} else if lineSplit[1] == "ls" {
			command = ls
		}
		args = lineSplit[2:]
	} else {
		command = createFile
		args = lineSplit
	}
	return command, args
}

func ls(node *Node, args ...string) {
	// Function is programatically unnecessary since ls
	// in input does nothing, but included for completeness
	return
}

func cd(node *Node, args ...string) {
	nodeName := node.name
	if args[0] == "/" {
		for nodeName != "/" {
			*node = *node.parent
			nodeName = node.name
		}
	} else if args[0] == ".." {
		*node = *node.parent
	} else {
		_, ok := node.children[args[0]]
		if ok {
			*node = node.children[args[0]]
		}
	}
}

func createFile(node *Node, args ...string) {
	var newFile *Node
	fname := args[1]
	if args[0] == "dir" {
		newFile = newNode(fname, nodeParent(*node))
		node.children[fname] = *newFile
	} else {
		size, _ := strconv.Atoi(args[0])
		newFile = newNode(fname, nodeIsFile, nodeSize(size), nodeParent(*node))
	}
	node.children[fname] = *newFile
}
