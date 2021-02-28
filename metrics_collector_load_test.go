package mc

import (
  "testing"
  "github.com/stretchr/testify/assert"
  
  "time"
  "flag"
  "github.com/golang/glog"
)


func TestCheckLoadClients(t *testing.T) {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", ".")
	flag.Set("v", "9")
	flag.Parse()

	glog.Info("Logging configured")
  
  //assert.Equal(t, "worker.api.test", wt.API)
  
  WorkerRegister(NewWorker())

  go Init("./etc.tests/")
  time.Sleep(2 * time.Second)

  res := GetPublicInfo()
  res_need := []PublicInfo([]PublicInfo{PublicInfo{CODE:"worker.api.test", Org:OrgInfo{OrgID:"", Org:"", Name:"", Description:"", Icon:"", Img:"", Url_Logo:"", Tel:"", Email:"", URL_dispatcher:""}, Cron:CronInfo{EverySeconds:0x0, EveryMinutes:0x11}, Status:StatusInfo{LastError:"", Ok:true, CntDevices:1, CntMetrics:5}}})
  res[0].Status.RunTime = 0
  res[0].Status.LastRun = res_need[0].Status.LastRun
  assert.Equal(t, res_need, res)

  defer Close()
}
