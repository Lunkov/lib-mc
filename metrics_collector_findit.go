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

var dbHandleRead *sqlx.DB

var districtCode *cache.Cache   // Distrct Code -> UUID
var deviceCode   *cache.Cache   // Device  Code -> UUID
var deviceSN     *cache.Cache   // Device  Serial Number -> UUID
var metricCode   *cache.Cache   // Metric  Code -> UUID

func initCaches() {
  cfg := getConfig()
  
  districtCode = cache.New("memory", 0, "", 0)
  deviceCode   = cache.New("memory", 0, "", 0)
  deviceSN     = cache.New("memory", 0, "", 0)
  metricCode   = cache.New("memory", 0, "", 0)
  
  dbHandleRead = dbConnect(models.ConnectStr(cfg.PostgresRead))
  if dbHandleRead != nil {
    glog.Errorf("ERR: DB: failed to connect database (read)")
    return
  }
}

func closeCaches() {
  if districtCode != nil {
    districtCode.Close()
    districtCode = nil
  }
  
  if deviceCode != nil {
    deviceCode.Close()
    deviceCode = nil
  }
  
  if deviceSN != nil {
    deviceSN.Close()
    deviceSN = nil
  }
  
  if metricCode != nil {
    metricCode.Close()
    metricCode = nil
  }
  
  dbClose(dbHandleRead)
  dbHandleRead = nil
}

func DistrictIDFindByCode(code string) (uuid.UUID, bool) {
  return cacheFindBy(districtCode, "districts", "code", code)
}

func DeviceIDFindByCode(code string) (uuid.UUID, bool) {
  return cacheFindBy(deviceCode, "devices", "code", code)
}

func DeviceIDFindBySN(sn string) (uuid.UUID, bool) {
  return cacheFindBy(deviceSN, "devices", "sn", sn)
}

func MetricIDFindByCode(code string) (uuid.UUID, bool) {
  return cacheFindBy(metricCode, "metrics", "code", code)
}

func cacheFindBy(c *cache.Cache, model string, search string, findit string) (uuid.UUID, bool) {
  uid1, ok := c.GetStr(findit)
  if ok {
    uid, _ := uuid.Parse(uid1)
    return uid, true
  }
  uid2, ok2 := GetOneId(dbHandleRead, "SELECT id FROM $1 WHERE $2=$3;", model, search, findit)
  if !ok2 {
    return uuid.Nil, false
  }
  c.SetStr(findit, uid2)
  uid, _ := uuid.Parse(uid2)
  return uid, true
}

func dbConnect(connectStr string) (*sqlx.DB) {
  var err error
  glog.Infof("LOG: DB: CONNECT: %s \n", connectStr)
  dbHandle, err := sqlx.Connect("postgres", connectStr)
  if err != nil {
    glog.Errorf("ERR: DB: CONNECT ERR: %s \n", err)
  }

  return dbHandle
}

func dbClose(dbHandle *sqlx.DB) {
  if dbHandle != nil {
    dbHandle.Close()
  }
}

func GetOneId(dbHandle *sqlx.DB, sqlStatement string, model string, search string, findit string) (string, bool) {
  if dbHandle == nil {
    glog.Errorf("ERR: DB: dbHandle == nil")
    return "", false
  }
  var err error
  var findId string
  row := dbHandle.QueryRow(sqlStatement, model, search, findit)
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
