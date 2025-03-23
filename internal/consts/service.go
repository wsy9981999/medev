package consts

const Service = `package {{ .Name }}


import(
	"{{ .ServicePkg }}"
	{{  if .Config }} 
	"github.com/gogf/gf/os/gctx"
	
	"github.com/gogf/gf/v2/frame/g"
{{ end }}
)

	{{ if .Config }}

type {{camelLower .Name}}Cfg struct {
	
}
	{{ end }}


func init(){
	service.Register{{ ucfirst .Name }}({{ if .Constructor  }}  news{{ ucfirst .Name }}() {{ else }}  &s{{ucfirst .Name }}{}  {{end}} )
}
type s{{ucfirst .Name }} struct {
	{{ if .Config }}
	cfg *{{camelLower .Name}}Cfg
	{{ end }}

}

{{ if .Constructor }}
func news{{ ucfirst .Name }}() *s{{ucfirst .Name }}{
	i:=&s{{ ucfirst .Name }}{}
	{{ if .Config }} 
		i.cfg=new({{camelLower .Name}}Cfg)
		g.Cfg().MustGet(gctx.New(),"{{.Name}}").Scan(&i.cfg)
	{{ end }}
	return i
}


{{ end }}
func (receiver *s{{ ucfirst .Name }}) Example(){
}

`
