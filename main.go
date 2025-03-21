package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
	"medev/internal/cmd"
	"medev/utils/mlog"
)

func main() {
	g.View().BindFuncMap(g.MapStrAny{
		"ucfist": gstr.UcFirst,
	})
	ctx := gctx.New()
	object, err := gcmd.NewFromObject(cmd.MeDev)
	if err != nil {
		mlog.Fatalf("%+v", err)
	}
	err = object.AddObject(cmd.MeDevInit, cmd.Controller)
	if err != nil {
		mlog.Fatalf("%+v", err)

	}
	object.Run(ctx)
}
