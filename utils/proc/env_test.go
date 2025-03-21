package proc

import (
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

func TestNewProcExist(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		NewProcExist()
	})
}
