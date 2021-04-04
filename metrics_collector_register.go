package mc

import (
  "github.com/golang/glog"
  
  "github.com/Lunkov/lib-ref"
)

func (m *MetricsCollector) WorkerRegister(worker WorkerInterface) {
  m.modFuncs[worker.GetAPI()] = worker
}

func (m *MetricsCollector) workerInit(code string, initInfo Info) {
  w, ok := m.modFuncs[code]
  if !ok {
    glog.Errorf("ERR: Client API '%s' not found", code)
    return
  }
  w.Init(initInfo)
}

func (m *MetricsCollector) workersClose() {
  for _, w := range m.modFuncs {
    w.Close()
  }
}

func (m *MetricsCollector) workerClose(code string) {
  w, ok := m.modFuncs[code]
  if !ok {
    glog.Errorf("ERR: Client API '%s' not found", code)
    return
  }
  w.Close()
}

func (m *MetricsCollector) workersPublicInfo() []PublicInfo {
  var res []PublicInfo
  for _, w := range m.modFuncs {
    res = append(res, w.GetPublicInfo())
  }
  return res
}

func (m *MetricsCollector) workersResults() (map[string]map[string]interface{}) {
  res := make(map[string]map[string]interface{})
  for _, w := range m.modFuncs {
    res[w.GetAPI()] = w.GetResultArray()
  }
  return res
}

func (m *MetricsCollector) workerRun(code string) {
  w, ok := m.modFuncs[code]
  if !ok {
    glog.Errorf("ERR: Client API '%s' not found\n", code)
    return
  }
  ref.RunMethodIfExists(w, "Start")
  ref.RunMethodIfExists(w, "GetData")
  ref.RunMethodIfExists(w, "Finish")
}

func (m *MetricsCollector) getWorker(code string) *WorkerInterface {
  w, ok := m.modFuncs[code]
  if ok {
    return &w
  }
  return nil
}

func (m *MetricsCollector) workerExists(code string) bool {
  _, ok := m.modFuncs[code]
  return ok
}
