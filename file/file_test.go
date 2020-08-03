package file

import (
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"strings"
	"testing"
)

const noExistedFile = "/tmp/not_existed_file"

func TestFileExist(t *testing.T) {
	existed, isDir := Exist("./file.go")
	if !existed || isDir {
		t.Errorf("./file.go should exists, but it didn't")
	}

	if !Exists("./file.go") {
		t.Errorf("./file.go should exists, but it didn't")
	}

	existed, isDir = Exist(noExistedFile)
	if existed || isDir {
		t.Errorf("Weird, how could this file exists: %s", noExistedFile)
	}
}

func TestMode(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want os.FileMode
	}{
		{
			name: "file.go",
			args: args{
				path: "./file.go",
			},
			want: 420,
		},
		{
			name: "noExistedFile",
			args: args{
				path: noExistedFile,
			},
			want: 0755,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mode(tt.args.path); got != tt.want {
				t.Errorf("Mode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteStringToFile(t *testing.T) {
	tempDir, err := ioutil.TempDir(os.TempDir(), "WriteStringToFile")
	if err != nil {
		t.Errorf(err.Error())
	}
	defer os.RemoveAll(tempDir)

	filePath := path.Join(tempDir, "a.txt")
	err = WriteStringToFile("abc", filePath, 0777)
	if err != nil {
		t.Errorf(err.Error())
	}

	file, err := os.Open(filePath)
	if err != nil {
		t.Errorf(err.Error())
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		t.Errorf(err.Error())
	}

	content := string(bytes)

	if content != "abc" {
		t.Errorf(err.Error())
	}
}

func TestAppendStringToFile(t *testing.T) {
	tempDir, err := ioutil.TempDir(os.TempDir(), "AppendStringToFile")
	if err != nil {
		t.Errorf(err.Error())
	}
	defer os.RemoveAll(tempDir)

	filePath := path.Join(tempDir, "a.txt")
	err = WriteStringToFile("abc", filePath, 0777)
	if err != nil {
		t.Errorf(err.Error())
	}

	err = AppendStringToFile("def", filePath)
	if err != nil {
		t.Errorf(err.Error())
	}

	file, err := os.Open(filePath)
	if err != nil {
		t.Errorf(err.Error())
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		t.Errorf(err.Error())
	}

	content := string(bytes)

	if content != "abcdef" {
		t.Errorf(err.Error())
	}
}

func TestGetDirList(t *testing.T) {
	tempDir, err := ioutil.TempDir(os.TempDir(), "GetDirList")
	if err != nil {
		t.Errorf(err.Error())
	}
	defer os.RemoveAll(tempDir)

	filePath := path.Join(tempDir, "a.txt")
	err = WriteStringToFile("abc", filePath, 0777)
	if err != nil {
		t.Errorf(err.Error())
	}

	dir1 := path.Join(tempDir, "d1")
	dir2 := path.Join(tempDir, "d2")

	if err = os.MkdirAll(dir1, 0755); err != nil {
		t.Errorf(err.Error())
	}
	if err = os.MkdirAll(dir2, 0755); err != nil {
		t.Errorf(err.Error())
	}

	dirList, err := GetDirList(tempDir)
	if err != nil {
		t.Errorf(err.Error())
	}

	want := []string{"d1", "d2"}
	if !reflect.DeepEqual(dirList, want) {
		t.Errorf("got %v want %v", dirList, want)
	}

	dir3 := path.Join(tempDir, ".d3")
	if err = os.MkdirAll(dir3, 0755); err != nil {
		t.Errorf(err.Error())
	}

	dirList, err = GetDirListWithFilter(tempDir, func(file os.FileInfo) bool {
		if strings.HasPrefix(file.Name(), ".") {
			return false
		}
		return true
	})
	if err != nil {
		t.Errorf(err.Error())
	}
	want = []string{"d1", "d2"}
	if !reflect.DeepEqual(dirList, want) {
		t.Errorf("got %v want %v", dirList, want)
	}
}
