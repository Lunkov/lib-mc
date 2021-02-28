package mc

import (
  "testing"
  "github.com/stretchr/testify/assert"
  
  "time"
  "flag"
  "github.com/golang/glog"
)

type TestWorkerInfo struct {
  WorkerInfo
}

func NewWorker() *TestWorkerInfo {
  w := new(TestWorkerInfo)
  w.API = "worker.api.test"
  return w
}


func (w *TestWorkerInfo) GetData() {
  glog.Infof("LOG: Worker GetData '%s'", w.GetAPI())  

  var dm []DeviceMetric
  
  layout := "2006-01-02T15:04:05.000Z"
  str := "2014-11-12T11:45:26.371Z"
  dt, _ := time.Parse(layout, str)

  dm = append(dm, DeviceMetric{Device_CODE: "C.1", Metric_CODE: "Weather.Station.Air.Temperature", DT: dt, Value: float64(14)})
  dm = append(dm, DeviceMetric{Device_CODE: "C.1", Metric_CODE: "Weather.Station.Air.Wind.Speed", DT: dt, Value: float64(2)})
  dm = append(dm, DeviceMetric{Device_CODE: "C.1", Metric_CODE: "Weather.Station.Air.Wind.Direction", DT: dt, Value: float64(134)})
  dm = append(dm, DeviceMetric{Device_CODE: "C.1", Metric_CODE: "Weather.Station.Air.Humidity", DT: dt, Value: float64(41)})
  dm = append(dm, DeviceMetric{Device_CODE: "C.1", Metric_CODE: "Weather.Station.Air.Pressure", DT: dt, Value: float64(1020)})

  w.ClientData.Status.CntDevices = 1
  w.ClientData.Status.CntMetrics = 5
  w.ClientData.Status.Ok = true
  
  w.SendMetrics(&dm)
}

func TestCheckWorker(t *testing.T) {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", ".")
	flag.Set("v", "9")
	flag.Parse()

	glog.Info("Logging configured")
  
  wt := NewWorker()
  wt.Init(Info{Nats:NatsInfo{ReturnArray: true, SubjectStat:"stats", SubjectMetric: "metrics"}})
  assert.Equal(t, "worker.api.test", wt.API)
  assert.Equal(t, "worker.api.test", wt.GetAPI())
  
  defer wt.Close()

  wt.Start()
  //wt.GetData()
  runMethodIfExists(wt.API, wt, "GetData")
  wt.Finish()
  
  res := wt.GetResultArray()
  res_need := map[string]interface{}{"stats":Result{API:"worker.api.test", Status:StatusInfo{LastError:"", Ok:true, CntDevices:1, CntMetrics:5}}}
  tr_need, _ := (res_need["stats"]).(Result)
  tr, _ := (res["stats"]).(Result)
  tr.Status.LastRun = tr_need.Status.LastRun
  tr.Status.RunTime = tr_need.Status.RunTime
  assert.Equal(t, tr_need, tr)

  resP := wt.GetPublicInfo()
  resP_need := PublicInfo{CODE:"worker.api.test", Org:OrgInfo{OrgID:"", Org:"", Name:"", Description:"", Icon:"", Img:"", Url_Logo:"", Tel:"", Email:"", URL_dispatcher:""}, Cron:CronInfo{EverySeconds:0x0, EveryMinutes:0x0}, Status:StatusInfo{RunTime:56631, LastError:"", Ok:true, CntDevices:1, CntMetrics:5}}
  resP_need.Status.LastRun = resP.Status.LastRun
  resP_need.Status.RunTime = resP.Status.RunTime
  assert.Equal(t, resP_need, resP)

}
