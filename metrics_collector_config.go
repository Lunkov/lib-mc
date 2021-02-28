package mc

import (
  "io/ioutil"
  "path/filepath"

  "gopkg.in/yaml.v2"
  
  "github.com/golang/glog"
  
  "github.com/Lunkov/lib-model"
)

type ConfigInfo struct {
  ConfigPath      string

  Nats            NatsInfo                `yaml:"nats"`
  PostgresRead    models.PostgreSQLInfo   `yaml:"postgres_read"`
}

var globConf = ConfigInfo{}

func setConfig(conf ConfigInfo) {
  globConf = conf
}

func getConfig() *ConfigInfo {
  return &globConf
}

func loadConfig(filename string) ConfigInfo {
  var err error
  var conf = ConfigInfo{}
    
  yamlFile, err := ioutil.ReadFile(filename)
  if err != nil {
    glog.Errorf("ERR: yamlFile(%s)  #%v ", filename, err)
    return conf
  }

  err = yaml.Unmarshal(yamlFile, &conf)
  if err != nil {
    glog.Errorf("ERR: yamlFile(%s): YAML: %v", filename, err)
  }
  
  if conf.ConfigPath == "" {
    conf.ConfigPath = filepath.Dir(filename)
  }
  return conf
}

