package format

import (
	"github.com/tsdrm/go-trans"
	"github.com/tsdrm/go-trans/format/flv"
)

func Init() {
	var formats = go_trans.GetFormats()
	for _, format := range formats {
		go_trans.RegisterPlugin(format, getTransPlugin(format))
	}
}

func getTransPlugin(format string) func() go_trans.TransPlugin {
	switch format {
	case "flv":
		return func() go_trans.TransPlugin { return &flv.Flv{} }
	}
	return nil
}
