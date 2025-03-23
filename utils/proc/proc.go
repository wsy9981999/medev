package proc

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/util/gmode"
	"io"
	"medev/utils/mlog"
)

func Run(ctx context.Context, q any, seq ...bool) error {
	_seq := false
	if len(seq) > 0 {
		_seq = seq[0]
	}

	obj, err := transformToProcessObj(q)
	if err != nil {
		return err
	}
	if _seq {

		for _, process := range obj {
			logCmd(process)
			err := process.Run(ctx)
			if err != nil {
				return err
			}
		}

	} else {
		p := grpool.New()
		var _err error

		for _, process := range obj {
			p.Add(ctx, func(ctx context.Context) {
				logCmd(process)
				if _err != nil {
					return
				}
				_err = process.Run(ctx)
			})
		}
		p.Close()
		return _err
	}

	return nil
}
func RunSingle(ctx context.Context, q string, path ...string) (int, error) {
	cmd := BuildProc(q, func(p *gproc.Process) {
		if len(path) > 0 {
			p.Dir = path[0]
		}
	})

	logCmd(cmd)
	start, err := cmd.Start(ctx)
	if err != nil {
		return 0, err
	}
	err = cmd.Wait()
	if err != nil {
		return 0, err
	}
	return start, nil

}

type Opt func(p *gproc.Process)

func BuildProc(cmd string, opts ...Opt) *gproc.Process {
	processCmd := gproc.NewProcessCmd(cmd)
	processCmd.Stdin = nil
	processCmd.Stdout = nil
	processCmd.Stderr = nil
	processCmd.Dir = gfile.Pwd()
	for _, opt := range opts {
		opt(processCmd)
	}
	logCmd(processCmd)
	return processCmd
}
func WithStdin(r io.Reader) Opt {
	return func(p *gproc.Process) {
		p.Stdin = r
	}
}
func WithStdout(r io.Writer) Opt {
	return func(p *gproc.Process) {
		p.Stdout = r
	}
}
func WithStderr(r io.Writer) Opt {
	return func(p *gproc.Process) {
		p.Stderr = r
	}
}
func WithDir(r string) Opt {
	return func(p *gproc.Process) {
		p.Dir = r
	}
}

func logCmd(p *gproc.Process) {
	if p != nil && !gmode.IsProduct() {
		mlog.Debugf("Run `%s` in `%s`", p.String(), p.Dir)
	}
}
func transformToProcessObj(q any) ([]*gproc.Process, error) {
	switch x := q.(type) {
	case *gproc.Process:
		return []*gproc.Process{x}, nil
	case []*gproc.Process:
		return x, nil

	case string:
		return []*gproc.Process{BuildProc(x)}, nil
	case []string:
		processes := make([]*gproc.Process, len(x))
		for i := range x {
			processes[i] = BuildProc(x[i])
		}
		return processes, nil
	case [][]string:
		processes := make([]*gproc.Process, len(x))
		for i := range x {
			processes[i] = BuildProc(x[i][0], WithDir(x[i][1]))
		}
		return processes, nil
	default:
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "请提供[]*gproc.Process,*gproc.Process,string,[]string的参数")

	}
}
