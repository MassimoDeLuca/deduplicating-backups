package main

import (
	"io"
	"os"
	"path/filepath"

	"github.com/cespare/xxhash"
)

func ProcessFile(sourcePath, destinationFolder string) error {
	// fileSQL := `INSERT INTO files (path) VALUES (?);`
	// _, err := db.Exec(fileSQL, sourcePath)
	// if err != nil {
	// 	fmt.Println("Error inserting file data:", err)
	// 	return err
	// }

	// Open the source file for reading
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// Determine the filename for the destination file
	destinationFilePath := filepath.Join(destinationFolder, filepath.Base(sourcePath))

	// Create the destination file for writing
	destinationFile, err := os.Create(destinationFilePath)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	// create the buffer (1 MB in this case)
	bufferSize := 1024 * 1024
	buffer := make([]byte, bufferSize)

	// Loop to read and write in chunks
	for {
		// Read a chunk from the source file
		bytesRead, err := sourceFile.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		}

		xxhash.Sum64(buffer[:bytesRead])
		// hashes = append(hashes, hash)
		// fmt.Printf("Hash: %d\n", hash)

		// Write the chunk to the destination file
		_, err = destinationFile.Write(buffer[:bytesRead])
		if err != nil {
			return err
		}

		// Reached EOF
		if bytesRead == 0 {
			break
		}
	}

	return nil
}
