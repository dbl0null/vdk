package format

import (
	"github.com/dbl0null/vdk/av/avutil"
	"github.com/dbl0null/vdk/format/aac"
	"github.com/dbl0null/vdk/format/flv"
	"github.com/dbl0null/vdk/format/mp4"
	"github.com/dbl0null/vdk/format/rtmp"
	"github.com/dbl0null/vdk/format/rtsp"
	"github.com/dbl0null/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
