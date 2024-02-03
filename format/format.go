package format

import (
	"github.com/galindocode/vdk/av/avutil"
	"github.com/galindocode/vdk/format/aac"
	"github.com/galindocode/vdk/format/flv"
	"github.com/galindocode/vdk/format/mp4"
	"github.com/galindocode/vdk/format/rtmp"
	"github.com/galindocode/vdk/format/rtsp"
	"github.com/galindocode/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
