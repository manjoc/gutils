package file

import (
	"io/ioutil"
	"os"

	"github.com/CatchZeng/gutils/convert"
)

// Exists reports whether the file or directory exists.
func Exists(path string) (existed bool) {
	existed, _ = Exist(path)
	return
}

// Exist reports whether the file or directory exists.
func Exist(path string) (existed bool, isDir bool) {
	info, err := os.Stat(path)
	if err != nil {
		return !os.IsNotExist(err), false
	}
	return true, info.IsDir()
}

// Mode get file mode
func Mode(path string) os.FileMode {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return 0755
	}
	return fileInfo.Mode()
}

// WriteStringToFile write string to file
func WriteStringToFile(content, path string, mode os.FileMode) (err error) {
	bytes := convert.StringToBytes(content)
	return ioutil.WriteFile(path, bytes, mode)
}