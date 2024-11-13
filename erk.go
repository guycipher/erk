// Package erk
// BSD 3-Clause License
// Copyright (c) 2024, Alex Gaetano Padula
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
// list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
// this list of conditions and the following disclaimer in the documentation
// and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
package erk

import (
	"fmt"
	"github.com/cespare/xxhash/v2"
	"os"
	"path/filepath"
)

// MNode represents a node in a Merkle tree
type MNode struct {
	left  *MNode // Left child
	right *MNode // Right child
	hash  []byte // Hash of the node
	path  string // Path of the node (file path)
}

// FileData represents the file path and its contents
type FileData struct {
	path    string // File path
	content []byte // File content
}

// Erk represents the main Erk struct
type Erk struct {
	input     string // input to scan
	isDir     bool   // is input a directory
	filesData []*FileData
}

// hashFileData hashes provided file data using xxhash
func hashFileData(data []byte) []byte {
	h := xxhash.New()
	h.Write(data)
	return h.Sum(nil)
}

// New initiates a new Erk instance
func New(input string, isDir bool) (*Erk, error) {
	e := &Erk{
		input: input,
		isDir: isDir,
	}

	if isDir {
		files, err := readFilesInDirectory(input)
		if err != nil {
			return nil, err
		}
		e.filesData = files
	} else {
		file, err := readFile(input)
		if err != nil {
			return nil, err
		}
		e.filesData = append(e.filesData, file)

	}

	return e, nil
}

func readFilesInDirectory(dir string) ([]*FileData, error) {
	var files []*FileData

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			files = append(files, &FileData{path: path, content: content})
		}
		return nil
	})

	return files, err
}

// readFile reads a file and returns its path and content
func readFile(path string) (*FileData, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return &FileData{path: path, content: content}, nil
}

// BuildTree builds a Merkle tree from a list of FileData
func (e *Erk) BuildTree() *MNode {
	var nodes []MNode

	// Create leaf nodes
	for _, file := range e.filesData {
		node := MNode{
			hash: hashFileData(file.content),
			path: file.path,
		}
		nodes = append(nodes, node)
	}

	// Build the tree
	for len(nodes) > 1 {
		var newLevel []MNode

		for i := 0; i < len(nodes); i += 2 {
			if i+1 == len(nodes) {
				newLevel = append(newLevel, nodes[i])
			} else {
				left := nodes[i]
				right := nodes[i+1]
				combinedHash := append(left.hash, right.hash...)
				newNode := MNode{
					left:  &left,
					right: &right,
					hash:  hashFileData(combinedHash),
				}
				newLevel = append(newLevel, newNode)
			}
		}
		nodes = newLevel
	}

	return &nodes[0]
}

// GetMerkleRoot returns the root hash of the Merkle tree
func GetMerkleRoot(root *MNode) []byte {
	return root.hash
}

// PrintTree prints the Merkle tree
func PrintTree(node *MNode, level int) {
	if node == nil {
		return
	}
	if node.path != "" {
		fmt.Printf("%s%s: %x\n", string(' '+level*2), node.path, node.hash)
	} else {
		fmt.Printf("%s%x\n", string(' '+level*2), node.hash)
	}
	PrintTree(node.left, level+1)
	PrintTree(node.right, level+1)
}
