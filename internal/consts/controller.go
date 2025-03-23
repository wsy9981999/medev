package consts

const Controller = `package {{ .Version }}

import (
	"github.com/gogf/gf/v2/frame/g"
)

type {{ ucfirst .Name }}Req struct{
	g.Meta {{ .Tag }}
}
type {{ ucfirst .Name }}Res struct{
	
}
`
