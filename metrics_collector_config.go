package mc

import (
  "io/ioutil"
  "path/filepath"

  "gopkg.in/yaml.v2"
  
  "github.com/golang/glog"
  
  "github.com/Lunkov/lib-model"
  "github.com/Lunkov/lib-cache"
)

type ConfigInfo struct {
  ConfigPath         string

  Nats               NatsInfo                `yaml:"nats"`
  
  cacheDictricCodes  cache.CacheConfig       `yaml:"cache_dicstrict_codes"`
  cacheDeviceCodes   cache.CacheConfig       `yaml:"cache_device_codes"`
  cacheDeviceSN      cache.CacheConfig       `yaml:"cache_device_sn"`
  cacheMetricsCodes  cache.CacheConfig       `yaml:"cache_metrics_codes"`
  
  PostgresRead       models.PostgreSQLInfo   `yaml:"postgres_read"`
}

func (m *MetricsCollector) setConfig(conf ConfigInfo) {
  m.Conf = conf
}

func (m *MetricsCollector) GetConfig() *ConfigInfo {
  return &m.Conf
}

func (m *MetricsCollector) LoadConfig(filename string) {
  var err error
  var conf = ConfigInfo{}

  yamlFile, err := ioutil.ReadFile(filename)
  if err != nil {
    glog.Errorf("ERR: yamlFile(%s)  #%v ", filename, err)
    return
  }

  err = yaml.Unmarshal(yamlFile, &conf)
  if err != nil {
    glog.Errorf("ERR: yamlFile(%s): YAML: %v", filename, err)
  }
  
  if conf.ConfigPath == "" {
    conf.ConfigPath = filepath.Dir(filename)
  }
  m.Conf = conf
}

