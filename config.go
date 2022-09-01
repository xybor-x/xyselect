package xyselect

import (
	"github.com/xybor-x/xyerror"
	"github.com/xybor-x/xylog"
)

var egen = xyerror.Register("xyselect", 200000)

// Errors of package xyselect.
var (
	SelectorError      = egen.NewClass("SelectorError")
	ClosedChannelError = SelectorError.NewClass("ClosedChannelError")
	ExhaustedError     = SelectorError.NewClass("ExhaustedError")
)

var logger = xylog.GetLogger("xybor.xyplatform.xyselect")

func init() {
	logger.AddExtra("module", "xyselect")
}
