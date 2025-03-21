package consts

const Controller = `package {{ .Version }}

import (
	"github.com/gogf/gf/v2/frame/g"
)

type {{ ucfist .Name }}Req struct{
	g.Meta {{ .Tag }}
}
type {{ ucfist .Name }}Res struct{
	
}
`
