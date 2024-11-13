// Package erk tests
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
	"bytes"
	"fmt"
	"log"
	"testing"
)

func Test(t *testing.T) {
	e, err := New("test_dir", true) // Set to false if input is a file
	if err != nil {
		log.Fatalf("Failed to create Erk instance: %v", err)
	}

	// Build the Merkle tree
	root := e.BuildTree()

	// Get the Merkle root
	merkleRoot := GetMerkleRoot(root)
	expectedMerkleRoot := "2e8acb4bccf85227"
	if fmt.Sprintf("%x", merkleRoot) != expectedMerkleRoot {
		t.Errorf("Expected Merkle Root: %s, got: %x", expectedMerkleRoot, merkleRoot)
	}

	// Print the Merkle tree
	tree := PrintTreeBytes(root, 0)
	expectedTree := "2e8acb4bccf85227\n\ttest_dir/f1.txt: 0a75a91375b27d44\n\ttest_dir/f2.txt: 19a1d238fce6124f\n"

	if !bytes.Equal(tree, []byte(expectedTree)) {
		t.Errorf("Expected Merkle Tree:\n%s\nGot:\n%s", expectedTree, string(tree))
	}
}
