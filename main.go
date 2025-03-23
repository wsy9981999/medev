package main

import (
	"medev/internal/cmd"
	_ "medev/internal/packed"
	"medev/utils/mlog"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
)

func main() {
	g.View().BindFuncMap(g.MapStrAny{
		"ucfirst":    gstr.UcFirst,
		"camelLower": gstr.CaseCamelLower,
		"lower":      gstr.ToLower,
	})
	ctx := gctx.New()
	object, err := gcmd.NewFromObject(cmd.MeDev)
	if err != nil {
		mlog.Fatalf("%+v", err)
	}
	err = object.AddObject(cmd.MeDevInit, cmd.Controller, cmd.Service)
	if err != nil {
		mlog.Fatalf("%+v", err)

	}
	object.Run(ctx)
}
