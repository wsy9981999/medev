package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

type cMeDev struct {
	g.Meta `name:"medev"`
}

type cMeDevInput struct {
	g.Meta `name:"medev"`
}
type cMeDevOutput struct {
}

var MeDev = cMeDev{}

func (receiver *cMeDev) Index(ctx context.Context, in cMeDevInput) (out *cMeDevOutput, err error) {
	gcmd.CommandFromCtx(ctx).Print()
	return
}
