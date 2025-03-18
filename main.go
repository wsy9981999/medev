package main

import (
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	"medev/internal/cmd"
	"medev/utils/mlog"
)

func main() {
	ctx := gctx.New()
	object, err := gcmd.NewFromObject(cmd.MeDev)
	if err != nil {
		mlog.Fatalf("%+v", err)
	}
	err = object.AddObject(cmd.MeDevInit)
	if err != nil {
		mlog.Fatalf("%+v", err)

	}
	object.Run(ctx)
}
