package mc

import (
  "github.com/google/uuid"
  "github.com/golang/glog"
  "database/sql"
  _ "github.com/lib/pq"
  "github.com/jmoiron/sqlx"

  
  "github.com/Lunkov/lib-cache"
  "github.com/Lunkov/lib-model"
)

func (m *MetricsCollector) initCaches() {
  m.districtCode = cache.NewConfig(&m.Conf.cacheDictricCodes)
  m.deviceCode   = cache.NewConfig(&m.Conf.cacheDeviceCodes)
  m.deviceSN     = cache.NewConfig(&m.Conf.cacheDeviceSN)
  m.metricCode   = cache.NewConfig(&m.Conf.cacheMetricsCodes)
  
  if !m.dbConnect(models.ConnectStr(m.Conf.PostgresRead)) {
    glog.Errorf("ERR: DB: failed to connect database (read)")
  }
}

func (m *MetricsCollector) closeCaches() {
  if m.districtCode != nil {
    m.districtCode.Close()
    m.districtCode = nil
  }
  
  if m.deviceCode != nil {
    m.deviceCode.Close()
    m.deviceCode = nil
  }
  
  if m.deviceSN != nil {
    m.deviceSN.Close()
    m.deviceSN = nil
  }
  
  if m.metricCode != nil {
    m.metricCode.Close()
    m.metricCode = nil
  }
  
  m.dbClose()
  m.dbHandleRead = nil
}

func (m *MetricsCollector) DistrictIDFindByCode(code string) (uuid.UUID, bool) {
  return m.cacheFindBy(m.districtCode, "districts", "code", code)
}

func (m *MetricsCollector) DeviceIDFindByCode(code string) (uuid.UUID, bool) {
  return m.cacheFindBy(m.deviceCode, "devices", "code", code)
}

func (m *MetricsCollector) DeviceIDFindBySN(sn string) (uuid.UUID, bool) {
  return m.cacheFindBy(m.deviceSN, "devices", "sn", sn)
}

func (m *MetricsCollector) MetricIDFindByCode(code string) (uuid.UUID, bool) {
  return m.cacheFindBy(m.metricCode, "metrics", "code", code)
}

func (m *MetricsCollector) cacheFindBy(c cache.ICache, model string, search string, findit string) (uuid.UUID, bool) {
  var s uuid.UUID
  uid1, ok := c.Get(findit, &s)
  if ok {
    uid, ok2 := uid1.(uuid.UUID)
    return uid, ok2
  }
  uid2, ok2 := m.GetOneId("SELECT id FROM $1 WHERE $2=$3;", model, search, findit)
  if !ok2 {
    return uuid.Nil, false
  }
  c.Set(findit, uid2)
  uid, _ := uuid.Parse(uid2)
  return uid, true
}

func (m *MetricsCollector) dbConnect(connectStr string) bool {
  var err error
  m.dbHandleRead, err = sqlx.Connect("postgres", connectStr)
  if err != nil {
    glog.Errorf("ERR: DB: CONNECT ERR: %s \n", err)
  }

  return err == nil
}

func (m *MetricsCollector) dbClose() {
  if m.dbHandleRead != nil {
    m.dbHandleRead.Close()
  }
}

func (m *MetricsCollector) GetOneId(sqlStatement string, model string, search string, findit string) (string, bool) {
  if m.dbHandleRead == nil {
    glog.Errorf("ERR: DB: dbHandle == nil")
    return "", false
  }
  var err error
  var findId string
  row := m.dbHandleRead.QueryRow(sqlStatement, model, search, findit)
  ok := false
  switch err = row.Scan(&findId); err {
  case sql.ErrNoRows:
    glog.Errorf("ERR: DB: No rows were returned: %s", sqlStatement)
  case nil:
    ok = true
    break
  default:
    glog.Errorf("ERR: DB: %s: %v", sqlStatement, err)
  }
  return findId, ok
}
