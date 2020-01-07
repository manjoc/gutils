package file

import (
	"os"
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
