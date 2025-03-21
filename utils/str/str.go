package str

import "fmt"

func WrapQuote(str string) string {
	return fmt.Sprintf("`%s`", str)
}
