package mlog

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

var ctx = context.TODO()

func Fatal(q ...any) {
	g.Log().Fatal(ctx, q...)
}
func Fatalf(t string, q ...any) {
	g.Log().Fatalf(ctx, t, q...)
}
func Info(q ...any) {
	g.Log().Info(ctx, q...)

}
