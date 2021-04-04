package mc

import (
  "encoding/json"
  "gopkg.in/yaml.v2"
  "github.com/golang/glog"
  "github.com/Lunkov/lib-env"
  "github.com/jasonlvhit/gocron"
)

func (m *MetricsCollector) Init(configPath string) bool {
  m.scheduler = gocron.NewScheduler()
  m.LoadConfig(configPath + "/collectors.yaml")
  m.initCaches()
  env.LoadFromFiles(configPath + "/collectors/", "", m.loadYAML)
  <- m.scheduler.Start()
  return true
}

func (m *MetricsCollector) loadYAML(filename string, yamlFile []byte) int {
  var err error
  var mapMod = make(map[string]Info)

  err = yaml.Unmarshal(yamlFile, mapMod)
  if err != nil {
    glog.Errorf("ERR: yamlFile(%s): YAML: %v", filename, err)
  }
  if(len(mapMod) > 0) {
    for key, item := range mapMod {
      if m.workerExists(item.API) {
        if m.Conf.Nats.Url != "" {
          item.Nats = m.Conf.Nats
        }
        m.workerInit(item.API, item)
        
        if item.Cron.EverySeconds > 0 {
          glog.Infof("LOG: CRON: ADD TASK '%s' every %d seconds\n", key, item.Cron.EverySeconds)
          m.scheduler.Every(item.Cron.EverySeconds).Seconds().Lock().Do(m.workerRun, item.API)
        }
        if item.Cron.EveryMinutes > 0 {
          glog.Infof("LOG: CRON: ADD TASK '%s' every %d minutes\n", key, item.Cron.EveryMinutes)
          m.scheduler.Every(item.Cron.EveryMinutes).Minutes().Lock().Do(m.workerRun, item.API)
          go m.workerRun(item.API)
        }
      } else {
        glog.Errorf("ERR: Worker not found for Mod='%s' yamlFile(%s)", item.API, filename)
      }
    }
  }

  return len(mapMod)
}

func (m *MetricsCollector) Close() {
  m.scheduler.Clear()
  m.workersClose()
  m.closeCaches()
}

func (m *MetricsCollector) GetPublicInfo() []PublicInfo {
  return m.workersPublicInfo()
}

func (m *MetricsCollector) GetWorkersResults() map[string]map[string]interface{} {
  return m.workersResults()
}

func (m *MetricsCollector) GetPublicJson() []byte {
  pJson, _ := json.Marshal(m.workersPublicInfo())
  return pJson
}

