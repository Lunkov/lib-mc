package mc

import (
  "time"
  "github.com/google/uuid"
)

type NatsInfo struct {
  // SEND TO FOR TESTS
  ReturnArray         bool    `json:"return_array"              yaml:"return_array"`
  // SEND TO FOR PRODUCTION
  Url                 string  `json:"url"                  yaml:"url"`
  SubjectMetric       string  `json:"subject_metric"       yaml:"subject_metric"`
  SubjectDevice       string  `json:"subject_device"       yaml:"subject_device"`
  SubjectDeviceCoord  string  `json:"subject_device_coord" yaml:"subject_device_coord"`
  SubjectStat         string  `json:"subject_stat"         yaml:"subject_stat"`
}

type OrgInfo struct {
  OrgID          string  `json:"org_id"           yaml:"org_id"`
  Org            string  `json:"org"              yaml:"org"`
  Name           string  `json:"name"             yaml:"name"`
  Description    string  `json:"description"      yaml:"description"`
  Icon           string  `json:"icon"             yaml:"icon"`
  Img            string  `json:"img"              yaml:"img"`

  Url_Logo       string  `json:"url_logo"         yaml:"logo"`
  Tel            string  `json:"tel"              yaml:"tel"`
  Email          string  `json:"email"            yaml:"email"`
  URL_dispatcher string  `json:"url_dispatcher"   yaml:"url_dispatcher"`
}

type CronInfo struct {
  EverySeconds   uint64  `json:"every_seconds"    yaml:"every_seconds"`
  EveryMinutes   uint64  `json:"every_minutes"    yaml:"every_minutes"`
}

type StatusInfo struct {
  LastRun        time.Time     `json:"last_run"   yaml:"-"`
  RunTime        time.Duration `json:"run_time"   yaml:"-"`

  LastError      string  `json:"last_error"      yaml:"-"`
  Ok             bool    `json:"ok"              yaml:"-"`
  CntDevices     int64   `json:"cnt_devices"     yaml:"-"`
  CntMetrics     int64   `json:"cnt_metrics"     yaml:"-"`
}

type PublicInfo struct {
  CODE      string       `json:"code"       yaml:"code"`

  Org       OrgInfo      `json:"org"        yaml:"org"`
  Cron      CronInfo     `json:"cron"       yaml:"cron"`
  Status    StatusInfo   `json:"status"     yaml:"status"`
}

type Info struct {
  API       string    `json:"api"        yaml:"api"`
  // FROM
  Url       string    `json:"url"        yaml:"url"`
  UrlLogin  string    `json:"url_login"  yaml:"url_login"`
  UrlLogout string    `json:"url_logout" yaml:"url_logout"`
  UrlDevice string    `json:"url_device" yaml:"url_device"`
  UrlState  string    `json:"url_state"  yaml:"url_state"`

  Port      string    `json:"port"      yaml:"port"`
  Version   string    `json:"version"   yaml:"version"`
  
  Login     string    `json:"login"      yaml:"login"`
  Password  string    `json:"password"   yaml:"password"`
  Token     string    `json:"token"      yaml:"token"`

  Nats      NatsInfo     `json:"nats"       yaml:"nats"`
  Org       OrgInfo      `json:"org"        yaml:"org"`
  Cron      CronInfo     `json:"cron"       yaml:"cron"`
  Status    StatusInfo   `json:"status"     yaml:"status"`

  // WHAT  
  Org_ID             uuid.UUID  `json:"org_id"         yaml:"org_id"`
  Org_CODE           string     `json:"org_code"       yaml:"org_code"`
  District_ID        uuid.UUID  `json:"district_id"    yaml:"district_id"`
  District_CODE      string     `json:"district_code"  yaml:"district_code"`
  
  Device_ID          uuid.UUID  `json:"device_id"              yaml:"device_id"`
  Device_CODE        string     `json:"device_code"            yaml:"device_code"`
  DeviceModel_CODE   string     `json:"default_model_device"   yaml:"default_model_device"`
  Latitude           float64    `json:"latitude"         yaml:"latitude"`
  Longitude          float64    `json:"longitude"        yaml:"longitude"`
  
  ParamsCode         map[string]string      `json:"params_code"  yaml:"params_code"`
  DevicesId          map[string]uuid.UUID   `json:"devices_id"   yaml:"devices_id"`
  TokenId            map[string]uuid.UUID   `json:"tokens_id"    yaml:"tokens_id"`
  Values             map[string]string      `json:"values"       yaml:"values"`
  BuildingsId        map[string]string      `json:"buildings_id" yaml:"buildings_id"`
}

type Result struct {
  API         string       `json:"api"`
  Status      StatusInfo   `json:"status"`
}

type DeviceMetric struct {
  District_ID        uuid.UUID    `json:"district_id"`
  District_CODE      string       `json:"district_code"`
  
  Device_ID          uuid.UUID    `json:"device_id"`
  Device_CODE        string       `json:"device_code"`
  Device_SN          string       `json:"device_sn"`

  Metric_ID          uuid.UUID    `json:"metric_id"`
  Metric_CODE        string       `json:"metric_code"`
  
  DT                 time.Time    `json:"moment"`
  Value              float64      `json:"value"`
}

type DeviceCoord struct {
  Device_ID          uuid.UUID    `json:"device_id"`
  Device_CODE        string       `json:"device_code"`
  Device_SN          string       `json:"device_sn"`

  Latitude           float64      `json:"latitude"`
  Longitude          float64      `json:"longitude"`
}

type DeviceInfo struct {
  Device_ID          uuid.UUID    `json:"device_id"`
  Device_CODE        string       `json:"device_code"`
  Device_SN          string       `json:"device_sn"`

  Name               string     `dv:"name"          json:"name"           yaml:"name"`
  Description        string     `db:"description"   json:"description"    yaml:"description"`

  Org_ID             uuid.UUID  `json:"org_id"         yaml:"org_id"`
  Org_CODE           string     `json:"org_code"       yaml:"org_code"`
  District_ID        uuid.UUID  `json:"district_id"    yaml:"district_id"`
  District_CODE      string     `json:"district_code"  yaml:"district_code"`

  GisLayer_ID        uuid.UUID  `db:"gis_layer_id"  json:"gis_layer_id"   yaml:"gis_layer_id"    gorm:"column:gis_layer_id;type:uuid;`
  GisLayer_CODE      string     `db:""              json:"gis_layer_code" yaml:"gis_layer_code"`
  Latitude           float64      `json:"latitude"`
  Longitude          float64      `json:"longitude"`

  Building_ID        uuid.UUID  `db:"building_id"   json:"building_id"       yaml:"building_id"      gorm:"column:building_id;type:uuid;`
  Building_CODE      string     `db:""              json:"building_code"     yaml:"building_code"`

  DeviceModel_CODE   string     `json:"default_model_device"   yaml:"default_model_device"`

  CoordX             int64      `db:"coord_x"       json:"coord_x"        yaml:"coord_x"`
  CoordY             int64      `db:"coord_y"       json:"coord_y"        yaml:"coord_y"`
  
  NumFloor           int64      `db:"num_floor"     json:"num_floor"      yaml:"num_floor"`
  NumEntrance        int64      `db:"num_entrance"  json:"num_entrance"   yaml:"num_entrance"`
  NumLift            int64      `db:"num_lift"      json:"num_lift"       yaml:"num_lift"`
  NumFlat            int64      `db:"num_flat"      json:"num_flat"       yaml:"num_flat"`
  Roof               bool       `db:"roof"          json:"roof"           yaml:"roof"`
  Attic              bool       `db:"attic"         json:"attic"          yaml:"attic"`
  Basement           bool       `db:"basement"      json:"basement"       yaml:"basement"`
}
