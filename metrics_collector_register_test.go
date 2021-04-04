package mc

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "flag"
  "github.com/golang/glog"
)

func TestCheckRegisterWorker(t *testing.T) {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", ".")
	flag.Set("v", "9")
	flag.Parse()

	glog.Info("Logging configured")
  
  m := New()
  
  m.WorkerRegister(NewWorker())
  
  m.workerInit("worker.api.test", Info{Nats:NatsInfo{ReturnArray: true, SubjectStat:"stats", SubjectMetric: "metrics"}})
  m.workerRun("worker.api.test")
  defer m.workerClose("worker.api.test")
  
  resP := m.workersPublicInfo()
  resP_need := []PublicInfo{{CODE:"worker.api.test", Org:OrgInfo{OrgID:"", Org:"", Name:"", Description:"", Icon:"", Img:"", Url_Logo:"", Tel:"", Email:"", URL_dispatcher:""}, Cron:CronInfo{EverySeconds:0x0, EveryMinutes:0x0}, Status:StatusInfo{RunTime:56631, LastError:"", Ok:true, CntDevices:1, CntMetrics:5}}}
  resP_need[0].Status.LastRun = resP[0].Status.LastRun
  resP_need[0].Status.RunTime = resP[0].Status.RunTime
  assert.Equal(t, resP_need, resP)
}
