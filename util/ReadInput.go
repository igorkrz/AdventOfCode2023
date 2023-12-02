package util

import (
	"os"
	// "path"
	"path/filepath"
	"runtime"
	"strings"
)

// ReadFile is a wrapper over io/ioutil.ReadFile but also determines the dynamic
// absolute path to the file.
//
// Deprecated in favor of go:embed, refer to scripts/skeleton/tmpls
func ReadFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	content := make([]byte, stat.Size())
	_, err = file.Read(content)
	if err != nil {
		panic(err)
	}

	strContent := string(content)
	return strings.TrimRight(strContent, "\n")
}

// Dirname is a port of __dirname in node
func Dirname() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("getting calling function")
	}
	return filepath.Dir(filename)
}
