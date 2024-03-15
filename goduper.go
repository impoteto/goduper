package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

func calculateHash(file string, full bool, partialSize int64) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var reader io.Reader = f
	if !full {
		// No need for type assertion here, keep it as io.Reader
		reader = io.LimitReader(f, partialSize)
	}

	hash := md5.New()
	if _, err := io.Copy(hash, reader); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil

}
func findDuplicates(directory string, scanType string) {
	var partialSize int64 = 4096
	fileHashes := make(map[string]string)
	fileSizes := make(map[int64][]string)
	var mu sync.Mutex

	var wg sync.WaitGroup
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		wg.Add(1)
		go func() {
			defer wg.Done()

			size := info.Size()
			full := scanType == "full"
			hash, err := calculateHash(path, full, partialSize)
			if err != nil {
				fmt.Println("Error calculating hash:", err)
				return
			}

			mu.Lock()
			defer mu.Unlock()
			key := fmt.Sprintf("%d-%s", size, hash)
			if original, exists := fileHashes[key]; exists {
				fmt.Printf("Duplicate found: '%s' and '%s'\n", original, path)
			} else {
				fileHashes[key] = path
				fileSizes[size] = append(fileSizes[size], path)
			}
		}()

		return nil
	})

	if err != nil {
		fmt.Println("Error walking the file path:", err)
		return
	}

	wg.Wait()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <directory>\n", os.Args[0])
		os.Exit(1)
	}

	directory := os.Args[1]

	fmt.Print("Select scan type [quick/full]: ")
	scanType := ""
	_, err := fmt.Scan(&scanType)
	if err != nil || (scanType != "quick" && scanType != "full") {
		fmt.Println("Invalid option. Please choose 'quick' or 'full'.")
		os.Exit(1)
	}

	findDuplicates(directory, scanType)
}
