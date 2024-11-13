 ## erk
This package has the ability to hash a file or files inside a directory and return a merkle root or a detailed tree stating the files and their hashes.

 ## still in development

 ## basic example
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
     tree := erk.PrintTreeBytes(root)
 }
 ```