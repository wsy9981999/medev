package cmd

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"medev/internal/consts"
	"medev/utils/mlog"
	"medev/utils/proc"
	"medev/utils/str"
)

var Controller = cMeDevController{}

type cMeDevController struct {
	g.Meta `name:"controller"`
}

type cMeDevControllerInput struct {
	g.Meta  `name:"controller"`
	Module  string `arg:"true" name:"module" required:"true"`
	Name    string `arg:"true" name:"name" required:"true"`
	Version string `name:"version" short:"v"  d:"v1"`
	Method  string `name:"method" short:"m"  d:"get"`
	Path    string `name:"path"  short:"p" `
	Force   bool   `name:"force" short:"f" orphan:"true"`
}
type cMeDevControllerOutput struct {
}

// /api/模块/版本/定义文件.go
func (cmd *cMeDevController) Index(ctx context.Context, in cMeDevControllerInput) (out *cMeDevControllerOutput, err error) {
	in.Version = gstr.ToLower(in.Version)

	content, err := g.View().ParseContent(ctx, consts.Controller, g.MapStrAny{
		"Name":    in.Name,
		"Version": in.Version,
		"Tag":     buildTag(in),
	})
	if err != nil {
		return nil, err
	}
	apiFile := gfile.Join(str.SelectBackend(), "api", in.Module, in.Version, gstr.CaseSnake(in.Name)+".go")
	if gfile.Exists(apiFile) {
		if in.Force {
			err = gfile.PutContents(apiFile, content)
			if err != nil {
				return
			}
		} else {
			g.Log().Info(ctx, "api file already exists")
			return

		}
	}
	err = gfile.PutContents(apiFile, content)
	if err != nil {
		return
	}
	if !proc.ProcExistInstance.ExistGfCli() {

		mlog.Fatalf("%+v", gerror.NewCode(gcode.CodeNotFound, "'gf' no found"))
	}
	err = proc.BuildProc("gf gen ctrl", proc.WithDir(str.SelectBackend())).Run(ctx)

	return
}

func buildTag(in cMeDevControllerInput) string {

	return str.WrapQuote("path:\"/" + in.Version + "/" + gstr.CaseKebab(in.Module+"/"+in.Name) + "\" method:\"" + gstr.ToUpper(in.Method) + "\"")
}
