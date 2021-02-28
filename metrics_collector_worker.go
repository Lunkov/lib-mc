package mc

import (
  "time"
  "github.com/golang/glog"
  "github.com/Lunkov/lib-mq"
)

type WorkerInterface interface {
  GetAPI() string
  Init(conn Info)
  Close()
  
  GetData()
  
  Start()
  Finish()
  
  SendMetrics(metrics *[]DeviceMetric)
  SendDeviceCoord(deviceCoord *DeviceCoord)
  SendDeviceNew(device *DeviceInfo)
  
  GetResultArray() map[string]interface{}
  GetPublicInfo() PublicInfo
}

type WorkerInfo struct {
  API               string
  ClientData        Info
  Nats              mq.NatsInfo
  resArray          map[string]interface{}
}

func (w *WorkerInfo) GetAPI() string {
  return w.API
}

func (w *WorkerInfo) Init(conn Info) {
  w.ClientData = conn
  if w.ClientData.Nats.ReturnArray {
    w.resArray = make(map[string]interface{})
  }
  if w.ClientData.Nats.Url != "" {
    w.Nats.NatsInit(w.ClientData.Nats.Url)
  }
  glog.Infof("LOG: Init Client API '%s'", w.API)
}

func (w *WorkerInfo) Close() {
  if w.ClientData.Nats.Url != "" {
    w.Nats.NatsClose()
  }
}

func (w *WorkerInfo) SendMetrics(metrics *[]DeviceMetric) {
  if w.ClientData.Nats.Url != "" {
    w.Nats.NatsSendMsg(w.ClientData.Nats.SubjectMetric, metrics)
  }
  if w.ClientData.Nats.ReturnArray {
    w.resArray[w.ClientData.Nats.SubjectMetric] = (*metrics)
  }
  glog.Infof("LOG: Close Client API '%s'", w.API)
}

func (w *WorkerInfo) SendDeviceCoord(deviceCoord *DeviceCoord) {
  if w.ClientData.Nats.Url != "" {
    w.Nats.NatsSendMsg(w.ClientData.Nats.SubjectDeviceCoord, deviceCoord)
  }
  if w.ClientData.Nats.ReturnArray {
    w.resArray[w.ClientData.Nats.SubjectDeviceCoord] = (*deviceCoord)
  }
  if glog.V(9) {
    glog.Infof("DBG: SendDeviceCoord(%s).(SN=%s)", w.API, deviceCoord.Device_SN)
  }
}

func (w *WorkerInfo) SendDeviceNew(device *DeviceInfo) {
  if w.ClientData.Nats.Url != "" {
    w.Nats.NatsSendMsg(w.ClientData.Nats.SubjectDevice, device)
  }
  if w.ClientData.Nats.ReturnArray {
    w.resArray[w.ClientData.Nats.SubjectDevice] = (*device)
  }
  if glog.V(9) {
    glog.Infof("DBG: SendDeviceNew(%s).(SN=%s)", w.API, device.Device_SN)
  }
}

func (w *WorkerInfo) Start() {
  w.ClientData.Status.LastRun = time.Now()
  w.ClientData.Status.Ok = false
  w.ClientData.Status.LastError = ""
  w.ClientData.Status.CntDevices = 0
  w.ClientData.Status.CntMetrics = 0
  if glog.V(9) {
    glog.Infof("DBG: Start(%s)", w.API)
  }
}

func (w *WorkerInfo) Finish() {
  var result Result
  w.ClientData.Status.RunTime = time.Since(w.ClientData.Status.LastRun)
  result.API = w.API
  if w.ClientData.Status.CntDevices == 0 || w.ClientData.Status.CntMetrics == 0 {
    w.ClientData.Status.Ok = false
  } 
  result.Status = w.ClientData.Status
  
  if w.ClientData.Nats.Url != "" {
    w.Nats.NatsSendMsg(w.ClientData.Nats.SubjectStat, result)
  }
  if w.ClientData.Nats.ReturnArray {
    w.resArray[w.ClientData.Nats.SubjectStat] = result
  }
  if glog.V(9) {
    glog.Infof("DBG: Finish(%s)", w.API)
  }
}

func (w *WorkerInfo) GetResultArray() map[string]interface{} {
  return w.resArray;
}

func (w *WorkerInfo) GetPublicInfo() PublicInfo {
  var res PublicInfo
  res.CODE = w.API
  res.Org  = w.ClientData.Org
  res.Cron = w.ClientData.Cron
  res.Status = w.ClientData.Status
  return res;
}
