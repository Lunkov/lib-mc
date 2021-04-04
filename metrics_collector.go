package mc

import (
  _ "github.com/lib/pq"
  "github.com/jmoiron/sqlx"
  
  "github.com/jasonlvhit/gocron"
  
  "github.com/Lunkov/lib-cache"
)

type MetricsCollector struct {
  Conf       ConfigInfo
  
  dbHandleRead   *sqlx.DB

  districtCode   cache.ICache   // Distrct Code -> UUID
  deviceCode     cache.ICache   // Device  Code -> UUID
  deviceSN       cache.ICache   // Device  Serial Number -> UUID
  metricCode     cache.ICache   // Metric  Code -> UUID

  scheduler     *gocron.Scheduler

  modFuncs       map[string]WorkerInterface
}

func New() *MetricsCollector {
  return &MetricsCollector{modFuncs: make(map[string]WorkerInterface)}
}
