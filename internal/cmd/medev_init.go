package cmd

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"medev/utils/mlog"
	"medev/utils/proc"
	"medev/utils/str"
)

var MeDevInit = cMeDevInit{}

type cMeDevInit struct {
	g.Meta `name:"init"`
}

type cMeDevInitInput struct {
	g.Meta `name:"init"`

	Name string `name:"name" arg:"true" v:"required"`
	Git  bool   `name:"git" orphan:"true" short:"g"`
}
type cMeDevInitOutput struct {
}

func (receiver *cMeDevInit) Init(ctx context.Context, in cMeDevInitInput) (out *cMeDevInitOutput, err error) {
	err = gfile.Mkdir(in.Name)
	if err != nil {
		return nil, err
	}
	mlog.Info("初始化goframe")
	err = receiver.initGoFrame(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	mlog.Info("初始化前端")
	err = receiver.initFrontend(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	if in.Git {
		mlog.Info("初始化git")

		err = receiver.initGit(ctx, in.Name)
		if err != nil {
			return nil, err
		}
	}

	return

}

func (receiver *cMeDevInit) initGit(ctx context.Context, name string) error {
	if !proc.ProcExistInstance.ExistGitCli() {
		mlog.Fatalf("%+v", gerror.NewCode(gcode.CodeNotFound, "'git' no found"))

	}
	q := gfile.Join(gfile.Pwd(), name)
	return proc.Run(ctx, [][]string{
		{"git init", q},
		{"git add -A", q},
		{"git commit -m \"Initial Git Commit\"", q},
	}, true)
}

func (receiver *cMeDevInit) initGoFrame(ctx context.Context, name string) error {
	if !proc.ProcExistInstance.ExistGfCli() {

		mlog.Fatalf("%+v", gerror.NewCode(gcode.CodeNotFound, "'gf' no found"))
	}
	base := gfile.Join(gfile.Pwd(), name)

	return proc.Run(ctx, [][]string{
		{"gf init backend -u", base},
		{"go mod tidy", str.SelectBackend()},
	}, true)
}

func (receiver *cMeDevInit) initFrontend(ctx context.Context, name string) error {
	pm := proc.ProcExistInstance.DefaultFrontendPm()
	if pm == "" {
		mlog.Fatalf("%+v", gerror.NewCode(gcode.CodeNotFound, "frontend package manager no found"))
	}
	base := gfile.Join(gfile.Pwd(), name)
	return proc.Run(ctx, [][]string{
		{
			fmt.Sprintf("%s create vue --bare --ts --router --jsx --pinia --vitest --eslint  --eslint-with-oxlint   --prettier frontend", pm),
			base,
		},
		{
			fmt.Sprintf("%s install", pm),
			str.SelectFrontend(),
		},
		{
			fmt.Sprintf("%s add alova", pm),
			str.SelectFrontend(),
		},
		{
			fmt.Sprintf("%s add -D  @alova/wormhole", pm),
			str.SelectFrontend(),
		},
		{
			fmt.Sprintf("%s run alova init", pm),
			str.SelectFrontend(),
		},
	}, true)
}
