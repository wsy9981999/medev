package proc

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/gogf/gf/v2/os/gstructs"
	"reflect"
	"sync"
)

var ProcExistInstance = NewProcExist()

type ProcExist struct {
	Git  bool `json:"git,omitempty" name:"git" cmd:"git -v"`
	Gf   bool `json:"gf,omitempty" name:"gf" cmd:"gf -v"`
	Bun  bool `json:"bun,omitempty" name:"bun" cmd:"bun -v"`
	Pnpm bool `json:"pnpm,omitempty" name:"pnpm" cmd:"pnpm -v"`
	Npm  bool `json:"npm,omitempty"  name:"npm" cmd:"npm -v"`
	Yarn bool `json:"yarn,omitempty"  name:"yarn" cmd:"yarn -v"`
}

func NewProcExist() *ProcExist {
	p := &ProcExist{}
	field, err := gstructs.TagMapName(p, []string{"cmd"})
	if err != nil {
		return nil
	}
	pp := reflect.ValueOf(p).Elem()
	wg := new(sync.WaitGroup)
	wg.Add(len(field))
	for cmd := range field {
		//fmt.Printf("out:%p,%p\n", &x, &f)
		processCmd := gproc.NewProcessCmd(cmd)
		processCmd.Stdout = nil
		processCmd.Stderr = nil
		processCmd.Stdin = nil
		tf := pp.FieldByName(field[cmd])
		go func(tf reflect.Value, q *gproc.Process) {
			//fmt.Printf("in:%p\n", &f)
			defer wg.Done()

			err := processCmd.Run(gctx.New())
			if err != nil {
				return
			}
			tf.SetBool(processCmd.ProcessState.ExitCode() == 0)
		}(tf, processCmd)
	}
	wg.Wait()

	return p
}

func (receiver *ProcExist) DefaultFrontendPm() string {
	f := []string{"bun", "pnpm", "yarn", "npm"}

	field, err := gstructs.TagMapField(receiver, []string{"name"})
	if err != nil {
		return ""
	}
	for _, f2 := range f {
		if field[f2].Value.Bool() {
			return f2
		}
	}

	return ""
}
func (receiver *ProcExist) ExistGfCli() bool {
	return receiver.Gf
}
func (receiver *ProcExist) ExistGitCli() bool {
	return receiver.Git
}
