package lnks

import (
	"strings"

	"github.com/efjoubert/lnksys/env"
	"github.com/efjoubert/lnksys/network"
	"github.com/efjoubert/lnksys/service"
)

//LnkService LnkService
type LnkService struct {
	*service.Service
	brkrfnc func(exenme string,exealias string,args ...string)
}

//NewLnkService NewLnkService
func NewLnkService(brokerfunc ...interface{}) (lnksrvs *LnkService, err error) {
	lnksrvs = &LnkService{}
	var srv, svrerr = service.NewService("LnkService", "LnkService", "LnkService", func(srvs *service.Service, args ...string) {
		lnksrvs.startLnkService(args...)
	}, func(srvs *service.Service, args ...string) {
		lnksrvs.runLnkService(args...)
	}, func(srvs *service.Service, args ...string) {
		lnksrvs.stopLnkService(args...)
	})
	if len(brokerfunc)==1 {
		if brfnc,brfcnok:= brokerfunc[0].(func(exenme string,exealias string,args ...string)); brfcnok {
			lnksrvs.brkrfnc=brfnc
		}
	}
	if svrerr == nil {
		lnksrvs.Service = srv
	} else {
		err = svrerr
		lnksrvs = nil
	}
	return
}

func (lnksrvs *LnkService) startLnkService(args ...string) {
	network.MapRoots("/", strings.Replace(lnksrvs.ServiceExeFolder(), "\\", "/", -1), "resources/", "./resources", "apps/", "./apps")
	network.DefaultServeHttp(nil, "GET", "/@"+lnksrvs.ServiceName()+".conf@.js", nil)
}

func (lnksrvs *LnkService) runLnkService(args ...string) {
	if lnksrvs.IsConsole() {
		var d = make(chan bool, 1)
		env.WrapupCall(func() {
			d <- true
		})
		var running = true
		for running {
			select {
			case e := <-d:
				if e {
					running = false
					break
				}
			}
		}
	} else if lnksrvs.IsBroker() {
		if lnksrvs.brkrfnc!=nil {
			lnksrvs.brkrfnc(lnksrvs.ServiceExeName(),lnksrvs.ServiceName(),args...)
		}
	}
}

func (lnksrvs *LnkService) stopLnkService(args ...string) {
	if lnksrvs.IsService() {
		env.ShutdownEnvironment()
	}
}
