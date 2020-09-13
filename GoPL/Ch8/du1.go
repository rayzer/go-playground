package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dir: %v", err)
		return nil
	}
	return entries
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	filesizes := make(chan int64)
	go func() {
		for _, dir := range roots {
			walkDir(dir, filesizes)
		}
		close(filesizes)
	}()

	var nfiles, nbytes int64
	for size := range filesizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles int64, nbytes int64) {
	fmt.Printf("%d files %1.f GB\n", nfiles, float64(nbytes)/1e9)
}
