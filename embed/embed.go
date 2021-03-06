package emded

import (
	"io"
	"strings"
	"sync"

	ace "github.com/efjoubert/lnksys/embed/ace"
	anime "github.com/efjoubert/lnksys/embed/anime"
	babel "github.com/efjoubert/lnksys/embed/babel"
	babylon "github.com/efjoubert/lnksys/embed/babylon"
	blockui "github.com/efjoubert/lnksys/embed/blockui"
	bootstrap "github.com/efjoubert/lnksys/embed/bootstrap"
	chart "github.com/efjoubert/lnksys/embed/chart"
	class "github.com/efjoubert/lnksys/embed/class"
	falcor "github.com/efjoubert/lnksys/embed/falcor"
	fontawesome "github.com/efjoubert/lnksys/embed/fontawesome"
	goldenlayout "github.com/efjoubert/lnksys/embed/goldenlayout"
	jquery "github.com/efjoubert/lnksys/embed/jquery"
	datatables "github.com/efjoubert/lnksys/embed/jquery/datatables"
	jqueryui "github.com/efjoubert/lnksys/embed/jquery/jqueryui"
	jspanel "github.com/efjoubert/lnksys/embed/jspanel"
	jss "github.com/efjoubert/lnksys/embed/jss"
	materialdb "github.com/efjoubert/lnksys/embed/materialdb"
	pdf "github.com/efjoubert/lnksys/embed/pdf"
	phaser "github.com/efjoubert/lnksys/embed/phaser"
	pixi "github.com/efjoubert/lnksys/embed/pixi"
	react "github.com/efjoubert/lnksys/embed/react"
	require "github.com/efjoubert/lnksys/embed/require"
	reveal "github.com/efjoubert/lnksys/embed/reveal"
	rxjs "github.com/efjoubert/lnksys/embed/rxjs"
	sip "github.com/efjoubert/lnksys/embed/sip"
	three "github.com/efjoubert/lnksys/embed/three"
	typescript "github.com/efjoubert/lnksys/embed/typescript"
	video "github.com/efjoubert/lnksys/embed/video"
	webaction "github.com/efjoubert/lnksys/embed/webaction"
	iorw "github.com/efjoubert/lnksys/iorw"
)

var cachedResources map[string]*iorw.BufferedRW

func cachedRsrs() map[string]*iorw.BufferedRW {
	if cachedResources == nil {
		cachedResources = map[string]*iorw.BufferedRW{}
	}
	return cachedResources
}

func EmbedFindJS(embedfindjs string) (embedjs io.Reader) {
	embedfindjs = strings.Replace(embedfindjs, "\\", "/", -1)
	if strings.LastIndex(embedfindjs, "/") >= 0 {
		embedfindjs = embedfindjs[strings.LastIndex(embedfindjs, "/")+1:]
	}
	if embedjs = class.ClassFindJS(embedfindjs); embedjs != nil {
		return
	} else if embedjs = react.ReactFindJS(embedfindjs); embedjs != nil {
		return
	} else if embedjs = react.SchedulerFindJS(embedfindjs); embedjs != nil {
		return
	} else if embedjs = jquery.JQueryFindJS(embedfindjs); embedjs != nil {
		return
	} else if embedjs = jqueryui.JQueryUIFindJSCSS(embedfindjs); embedjs != nil {
		return
	} else if embedjs = datatables.DataTablesFindJSCSS(embedfindjs); embedjs != nil {
		return
	} else if embedjs = require.RequireFindJS(embedfindjs); embedjs != nil {
		return
	} else if embedjs = bootstrap.BootstrapFindJSCSS(embedfindjs); embedjs != nil {
		return
	} else if embedjs = fontawesome.FontawesomeFindJSCSS(embedfindjs); embedjs != nil {
		return
	} else if embedjs = materialdb.MaterialDBFindJSCSS(embedfindjs); embedjs != nil {
		return
	} else if embedjs = babel.BabelFindJS(embedfindjs); embedjs != nil {
		return
	} else if embedjs = babylon.BabylonFindJS(embedfindjs); embedjs != nil {
		return
	} else if embedjs = three.ThreeFindJS(embedfindjs); embedjs != nil {
		return
	} else if embedjs = jss.JSSFindJS(embedfindjs); embedjs != nil {
		return
	} else if embedjs = rxjs.RxJSFindJS(embedfindjs); embedjs != nil {
		return
	} else if embedjs = falcor.FalcorFindJS(embedfindjs); embedjs != nil {
		return
	} else if embedjs = video.VideoFindJSCSS(embedfindjs); embedjs != nil {
		return
	} else if embedjs = jspanel.JSPanelFindJSCSS(embedfindjs); embedjs != nil {
		return
	} else if embedjs = ace.AcsJSFindJS(embedfindjs); embedjs != nil {
		return
	} else if embedjs = chart.ChartFindJS(embedfindjs); embedjs != nil {
		return
	} else if embedjs = webaction.WebactionFindJS(embedfindjs); embedjs != nil {
		return
	} else if embedjs = blockui.BlockuiFindJS(embedfindjs); embedjs != nil {
		return
	} else if embedjs = typescript.TypescriptFindJS(embedfindjs); embedjs != nil {
		return
	} else if embedjs = anime.AnimeFindJS(embedfindjs); embedjs != nil {
		return
	} else if embedjs = reveal.RevealFindJS(embedfindjs); embedjs != nil {
		return
	} else if embedjs = pixi.PixiFindJS(embedfindjs); embedjs != nil {
		return
	} else if embedjs = pdf.PdfFindJS(embedfindjs); embedjs != nil {
		return
	} else if embedjs = phaser.PhaserFindJS(embedfindjs); embedjs != nil {
		return
	} else if embedjs = goldenlayout.GoldenLayoutFindJSCSS(embedfindjs); embedjs != nil {
		return
	} else if embedjs = sip.SipFindJS(embedfindjs); embedjs != nil {
		return
	}
	return
}

/*func EmbedFindJS(embedfindjs string) (embedjs io.Reader) {
	if strings.LastIndex(embedfindjs, "/") >= 0 {
		embedfindjs = embedfindjs[strings.LastIndex(embedfindjs, "/")+1:]
	}
	if FindChachedEmbed(embedfindjs); embedjs!=nil {
		return
	} else if embedjs = react.ReactFindJS(embedfindjs); embedjs != nil {
		return CacheEmbedResource(embedfindjs, embedjs)
	} else if embedjs = react.SchedulerFindJS(embedfindjs); embedjs != nil {
		return CacheEmbedResource(embedfindjs, embedjs)
	} else if embedjs = jquery.JQueryFindJS(embedfindjs); embedjs != nil {
		return CacheEmbedResource(embedfindjs, embedjs)
	} else if embedjs = require.RequireFindJS(embedfindjs); embedjs != nil {
		return CacheEmbedResource(embedfindjs, embedjs)
	} else if embedjs = bootstrap.BootstrapFindJSCSS(embedfindjs); embedjs != nil {
		return CacheEmbedResource(embedfindjs, embedjs)
	} else if embedjs = fontawesome.FontawesomeFindJSCSS(embedfindjs); embedjs != nil {
		return CacheEmbedResource(embedfindjs, embedjs)
	} else if embedjs = materialdb.MaterialDBFindJSCSS(embedfindjs); embedjs != nil {
		return CacheEmbedResource(embedfindjs, embedjs)
	} else if embedjs = babel.BabelFindJS(embedfindjs); embedjs != nil {
		return CacheEmbedResource(embedfindjs, embedjs)
	} else if embedjs = babylon.BabylonFindJS(embedfindjs); embedjs != nil {
		return CacheEmbedResource(embedfindjs, embedjs)
	} else if embedjs = three.ThreeFindJS(embedfindjs); embedjs != nil {
		return CacheEmbedResource(embedfindjs, embedjs)
	} else if embedjs = jss.JSSFindJS(embedfindjs); embedjs != nil {
		return CacheEmbedResource(embedfindjs, embedjs)
	} else if embedjs = rxjs.RxJSFindJS(embedfindjs); embedjs != nil {
		return CacheEmbedResource(embedfindjs, embedjs)
	} else if embedjs = falcor.FalcorFindJS(embedfindjs); embedjs != nil {
		return CacheEmbedResource(embedfindjs, embedjs)
	} else if embedjs = video.VideoFindJSCSS(embedfindjs); embedjs != nil {
		return CacheEmbedResource(embedfindjs, embedjs)
	} else if embedjs = jspanel.JSPanelFindJSCSS(embedfindjs); embedjs != nil {
		return CacheEmbedResource(embedfindjs, embedjs)
	} else if embedjs = ace.AcsJSFindJS(embedfindjs); embedjs != nil {
		return CacheEmbedResource(embedfindjs, embedjs)
	} else if embedjs = chart.ChartFindJS(embedfindjs); embedjs != nil {
		return CacheEmbedResource(embedfindjs, embedjs)
	}
	return
}*/

var chcdEmddLck = &sync.Mutex{}

func RemoveCachedEmbedResource(embedfindjs string) (rmvd bool) {
	chcdEmddLck.Lock()
	defer chcdEmddLck.Unlock()
	if cachedResources == nil {
		rmvd = false
	} else {
		if cachedrw, cachedrwok := cachedRsrs()[embedfindjs]; cachedrwok {
			delete(cachedRsrs(), embedfindjs)
			cachedrw.Close()
			cachedResources = nil
			rmvd = true
		} else {
			rmvd = true
		}
	}
	return
}

func CacheEmbedResource(embedfindjs string, embedjs io.Reader) io.Reader {
	var cachedRD *iorw.BufferedRW = iorw.NewBufferedRW(0, nil)
	var buf = make([]byte, 81920)
	io.CopyBuffer(cachedRD, embedjs, buf)
	if RemoveCachedEmbedResource(embedfindjs) {
		func() {
			chcdEmddLck.Lock()
			defer chcdEmddLck.Unlock()

			cachedRsrs()[embedfindjs] = cachedRD
		}()
	} else {
		func() {
			chcdEmddLck.Lock()
			defer chcdEmddLck.Unlock()
			cachedRsrs()[embedfindjs] = cachedRD
		}()
	}
	return FindChachedEmbed(embedfindjs)
}

func FindChachedEmbed(embedfindjs string) (cachedReader *iorw.BufferedRW) {
	func() {
		chcdEmddLck.Lock()
		defer chcdEmddLck.Unlock()
		if strings.LastIndex(embedfindjs, "/") >= 0 {
			embedfindjs = embedfindjs[strings.LastIndex(embedfindjs, "/")+1:]
		}
		if cachedResources == nil {
			return
		} else if cachedrw, cachedrwok := cachedRsrs()[embedfindjs]; cachedrwok {
			cachedReader = iorw.NewBufferedRW(81920, cachedrw)
		}
	}()
	return
}
