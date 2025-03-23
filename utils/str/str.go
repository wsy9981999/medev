package str

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gfile"
	"golang.org/x/tools/imports"
	"medev/utils/mlog"
)

const backendDir = "backend"
const frontendDir = "frontend"

func WrapQuote(str string) string {
	return fmt.Sprintf("`%s`", str)
}
func SelectBackend() string {
	if gfile.Exists(backendDir) && gfile.IsDir(backendDir) {
		return gfile.Join(gfile.Pwd(), backendDir)
	}
	return gfile.Pwd()
}
func SelectFrontend() string {
	if gfile.Exists(frontendDir) && gfile.IsDir(frontendDir) {
		return gfile.Join(gfile.Pwd(), frontendDir)
	}
	return gfile.Pwd()
}
func GoFmt(path string) {
	replaceFunc := func(path, content string) string {
		res, err := imports.Process(path, []byte(content), nil)
		if err != nil {
			mlog.Printf(`error format "%s" go files: %v`, path, err)
			return content
		}
		return string(res)
	}

	var err error
	if gfile.IsFile(path) {
		// File format.
		if gfile.ExtName(path) != "go" {
			return
		}
		err = gfile.ReplaceFileFunc(replaceFunc, path)
	} else {
		// Folder format.
		err = gfile.ReplaceDirFunc(replaceFunc, path, "*.go", true)
	}
	if err != nil {
		mlog.Printf(`error format "%s" go files: %v`, path, err)
	}
}
