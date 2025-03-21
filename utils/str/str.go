package str

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gfile"
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
