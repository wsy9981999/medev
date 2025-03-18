package cmd

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gproc"
	"medev/utils/mlog"
	"medev/utils/proc"
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
	g.Log().Info(ctx, "初始化goframe")
	err = receiver.initGoFrame(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	g.Log().Info(ctx, "初始化前端")
	err = receiver.initFrontend(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	if in.Git {
		g.Log().Info(ctx, "初始化git")

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
	err := run(ctx, []string{
		"git init",
		"git add -A",
		"git commit -m \"Initial Git Commit\"",
	}, gfile.Join(gfile.Pwd(), name))
	if err != nil {
		return err
	}
	return nil
}
func run(ctx context.Context, cmd any, path string) error {
	switch t := cmd.(type) {
	case string:
		return runSlice(ctx, []string{t}, path)
	case []string:
		return runSlice(ctx, t, path)
	default:

	}
	return nil
}
func runSlice(ctx context.Context, cmds []string, path string) error {
	for _, cmd := range cmds {
		g.Log().Debugf(ctx, "run cmd:%s", cmd)
		processCmd := gproc.NewProcessCmd(cmd)
		processCmd.Stdout = nil
		processCmd.Stderr = nil
		processCmd.Stdin = nil
		processCmd.Dir = path
		err := processCmd.Run(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (receiver *cMeDevInit) initGoFrame(ctx context.Context, name string) error {
	if !proc.ProcExistInstance.ExistGfCli() {

		mlog.Fatalf("%+v", gerror.NewCode(gcode.CodeNotFound, "'gf' no found"))
	}
	err := run(ctx, []string{"gf init backend -u"}, gfile.Join(gfile.Pwd(), name))
	if err != nil {
		return err
	}
	err = run(ctx, []string{"go mod tidy"}, gfile.Join(gfile.Pwd(), name, "backend"))
	if err != nil {
		return err
	}
	return nil
}

func (receiver *cMeDevInit) initFrontend(ctx context.Context, name string) error {
	pm := proc.ProcExistInstance.DefaultFrontendPm()
	if pm == "" {
		mlog.Fatalf("%+v", gerror.NewCode(gcode.CodeNotFound, "frontend package manager no found"))
	}
	err := run(ctx, []string{
		fmt.Sprintf("%s create vue --bare --ts --router --jsx --pinia --vitest --eslint  --eslint-with-oxlint   --prettier frontend", pm),
	}, gfile.Join(gfile.Pwd(), name))

	if err != nil {
		return err
	}
	err = run(ctx, []string{
		fmt.Sprintf("%s install", pm), fmt.Sprintf("%s add alova", pm), fmt.Sprintf("%s add -D  @alova/wormhole", pm),
	}, gfile.Join(gfile.Pwd(), name, "frontend"))
	if err != nil {
		return err
	}
	return nil
}
