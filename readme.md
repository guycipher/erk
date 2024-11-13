 ## erk
 erk provides functionality to compute the Merkle root and generate a detailed Merkle tree for files within a directory or a single file. It uses the xxHash algorithm(provided by github.com/cespare/xxhash/v2) for hashing file contents.

 ### features
 - hash individual files or all files within a directory.
 - construct a Merkle tree from the hashed file data.
 - retrieve the Merkle root of the constructed tree.
 - print the Merkle tree in a human-readable format.

 ### basic example
 ```go
 package main

 import (
     "fmt"
     "log"
     "github.com/guycipher/erk"
 )

 func main() {
     // Create a new Erk instance
     e, err := erk.New("path/to/directory/or/file", true) // Set to false if input is a file
     if err != nil {
        log.Fatalf("Failed to create Erk instance: %v", err)
     }

     // Build the Merkle tree
     root := e.BuildTree()

     // Get the Merkle root
     merkleRoot := erk.GetMerkleRoot(root)
     fmt.Printf("Merkle Root: %x\n", merkleRoot)

     // Print the Merkle tree
     tree := erk.PrintTreeBytes(root, 0)
 }
 ```

 ### example tree output
 ```
 2e8acb4bccf85227
     test_dir/f1.txt: 0a75a91375b27d44
     test_dir/f2.txt: 19a1d238fce6124f
 ```