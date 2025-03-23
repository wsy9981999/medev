package cmd

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"medev/internal/consts"
	"medev/utils/mlog"
	"medev/utils/proc"
	"medev/utils/str"
)

var Service = cService{}

type cService struct {
	g.Meta `name:"service"`
}

type cServiceInput struct {
	g.Meta      `name:"service"`
	Name        string `name:"name" arg:"true" v:"required"`
	Constructor bool   `name:"con" short:"c" orphan:"true"`
	Config      bool   `name:"cfg" short:"C" orphan:"true"`
	Force       bool   `name:"force" short:"f" orphan:"true"`
}
type cServiceOutput struct {
}

func (receiver *cService) Index(ctx context.Context, in cServiceInput) (out *cServiceOutput, err error) {
	contentPath := gfile.Join(str.SelectBackend(), "internal", "logic", in.Name, gstr.CaseSnake(in.Name)+".go")
	if gfile.Exists(contentPath) {
		if !in.Force {
			mlog.Info(ctx, "logic file already exists")
			return
		}
	}
	if !proc.ProcExistInstance.ExistGfCli() {
		mlog.Fatalf("%+v", gerror.NewCode(gcode.CodeNotFound, "'gf' no found"))
	}
	if in.Config {
		in.Constructor = true
	}
	pkg := GetImportPath("internal/service")
	content, err := g.View().ParseContent(ctx, consts.Service, g.MapStrAny{
		"ServicePkg":  pkg,
		"Name":        in.Name,
		"Config":      in.Config,
		"Constructor": in.Constructor,
	})
	if err != nil {
		return nil, err
	}
	err = gfile.PutContents(contentPath, content)
	if err != nil {
		return
	}
	str.GoFmt(contentPath)
	err = proc.BuildProc("gf gen service").Run(ctx)

	return
}

func GetImportPath(dirPath string) string {
	// If `filePath` does not exist, create it firstly to find the import path.
	var realPath = gfile.RealPath(dirPath)
	if realPath == "" {
		_ = gfile.Mkdir(dirPath)
		realPath = gfile.RealPath(dirPath)
	}

	var (
		newDir     = gfile.Dir(realPath)
		oldDir     string
		suffix     = gfile.Basename(dirPath)
		goModName  = "go.mod"
		goModPath  string
		importPath string
	)
	for {
		goModPath = gfile.Join(newDir, goModName)
		if gfile.Exists(goModPath) {
			match, _ := gregex.MatchString(`^module\s+(.+)\s*`, gfile.GetContents(goModPath))
			importPath = gstr.Trim(match[1]) + "/" + suffix
			importPath = gstr.Replace(importPath, `\`, `/`)
			importPath = gstr.TrimRight(importPath, `/`)
			return importPath
		}
		oldDir = newDir
		newDir = gfile.Dir(oldDir)
		if newDir == oldDir {
			return ""
		}
		suffix = gfile.Basename(oldDir) + "/" + suffix
	}
}
