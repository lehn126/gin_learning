package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Check file or directory is exist
func IsExist(path string) bool {
	if path == "" {
		return false
	}
	_, e := os.Stat(path)
	if e == nil {
		return true
	}
	if !os.IsNotExist(e) {
		return true
	}
	return false
}

// Check given path is a directory
func IsDir(path string) bool {
	stat, e := os.Stat(path)
	if e != nil {
		return false
	}
	return stat.IsDir()
}

// Check given path is a file (not directory)
func IsFile(path string) bool {
	stat, e := os.Stat(path)
	if e != nil {
		return false
	}
	return !stat.IsDir()
}

// Get the parent directory of a file or directory
func GetParentDir(path string) string {
	return filepath.Dir(path)
}

// Read file and return the pointer of it's content
func ReadFile(path string) *string {
	constant, e := os.ReadFile(path)
	if e != nil {
		panic(fmt.Sprintf("meet error when open file %s, error: %s", path, e))
	}
	rt := string(constant)
	return &rt
}

// Read file with io buffer and return the pointer of it's content
func ReadFileWithBuffer(path string) *string {
	sbuilder := strings.Builder{}
	f, e := os.Open(path)
	if e != nil {
		panic(fmt.Sprintf("meet error when open file %s, error: %s", path, e))
	}
	defer f.Close() // Close file after after all of the processing

	reader := bufio.NewReader(f)
	buffer := make([]byte, 256)
	for {
		len, e := reader.Read(buffer)
		if e != nil {
			if e == io.EOF {
				sbuilder.Write(buffer[:len])
			} else {
				panic(fmt.Sprintf("meet error when read file %s, error: %s", path, e))
			}
			break
		} else {
			sbuilder.Write(buffer[:len])
		}
	}

	rt := sbuilder.String()
	return &rt
}

// Read each lines in a file and return the pointer of line array
func ReadFileToLines(path string) *[]string {
	lines := []string{}
	f, e := os.Open(path)
	if e != nil {
		panic(fmt.Sprintf("meet error when open file %s, error: %s", path, e))
	}
	defer f.Close() // Close file after after all of the processing

	scanner := bufio.NewScanner(f)
	// Use "\n" as splitter
	// This line can by removed because "\n" is the default splitter
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return &lines
}
