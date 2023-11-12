// Copyright (c) 2023 Luka Ivanović
// This code is licensed under MIT licence (see LICENCE for details)

package main

import (
	"bufio"
	"crypto/md5"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var dir string
	var verbose bool
	var empty bool
	flag.StringVar(&dir, "directory", "./", "directory to scan")
	flag.BoolVar(&verbose, "verbose", false, "output additional information")
	flag.BoolVar(&empty, "empty", false, "scan for empty files")
	flag.Parse()

	index := make(map[string][]string)
	duplicates := make([]string, 0)
	empties := make([]string, 0)

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if empty {
			info, err := d.Info()
			if err != nil {
				return err
			}
			if info.Size() == 0 {
				empties = append(empties, path)
			}
		}
		if d.IsDir() {
			return nil
		}
		hash := CalculateFileHash(path)
		if files, present := index[hash]; present && (len(files) == 1) {
			duplicates = append(duplicates, hash)
		}
		index[hash] = append(index[hash], path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	if verbose {
		fmt.Println("Number of duplicates:", len(duplicates))
		if empty {
			fmt.Println("Number of empty files:", len(empties))
		}
		fmt.Println("---")
	}
	for _, hash := range duplicates {
		if len(index[hash]) > 1 {
			paths := index[hash]
			fmt.Println(strings.Join(paths, ", "))
		}
	}
	if empty {
		fmt.Println("---")
		for _, path := range empties {
			fmt.Println(path)
		}
	}
}

func CalculateFileHash(path string) string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	buffer := make([]byte, 4096)
	m := md5.New()
	for {
		n, err := r.Read(buffer)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		m.Write(buffer[:n])
	}
	return base64.RawStdEncoding.EncodeToString(m.Sum(nil))
}
