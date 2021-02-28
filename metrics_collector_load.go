package mc

import (
  "encoding/json"
  "gopkg.in/yaml.v2"
  "github.com/golang/glog"
  "github.com/Lunkov/lib-env"
  "github.com/jasonlvhit/gocron"
)

var scheduler *gocron.Scheduler

func Init(configPath string) bool {
  scheduler = gocron.NewScheduler()
  setConfig(loadConfig(configPath + "/collectors.yaml"))
  initCaches()
  env.LoadFromYMLFiles(configPath + "/collectors/", loadYAML)
  <- scheduler.Start()
  return true
}

func loadYAML(filename string, yamlFile []byte) int {
  var err error
  var mapMod = make(map[string]Info)

  err = yaml.Unmarshal(yamlFile, mapMod)
  if err != nil {
    glog.Errorf("ERR: yamlFile(%s): YAML: %v", filename, err)
  }
  config := getConfig()
  if(len(mapMod) > 0) {
    for key, item := range mapMod {
      if workerExists(item.API) {
        if config.Nats.Url != "" {
          item.Nats = config.Nats
        }
        workerInit(item.API, item)
        
        if item.Cron.EverySeconds > 0 {
          glog.Infof("LOG: CRON: ADD TASK '%s' every %d seconds\n", key, item.Cron.EverySeconds)
          scheduler.Every(item.Cron.EverySeconds).Seconds().Lock().Do(workerRun, item.API)
        }
        if item.Cron.EveryMinutes > 0 {
          glog.Infof("LOG: CRON: ADD TASK '%s' every %d minutes\n", key, item.Cron.EveryMinutes)
          scheduler.Every(item.Cron.EveryMinutes).Minutes().Lock().Do(workerRun, item.API)
          go workerRun(item.API)
        }
      } else {
        glog.Errorf("ERR: Worker not found for Mod='%s' yamlFile(%s)", item.API, filename)
      }
    }
  }

  return len(mapMod)
}

func Close() {
  scheduler.Clear()
  workersClose()
  closeCaches()
}

func GetPublicInfo() []PublicInfo {
  return workersPublicInfo()
}

func GetPublicJson() []byte {
  pJson, _ := json.Marshal(workersPublicInfo())
  return pJson
}

