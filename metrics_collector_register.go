package mc

import (
  "reflect"
  "github.com/golang/glog"
)

var modFuncs = make(map[string]WorkerInterface)

func WorkerRegister(worker WorkerInterface) {
  modFuncs[worker.GetAPI()] = worker
}

func workerInit(code string, initInfo Info) {
  w, ok := modFuncs[code]
  if !ok {
    glog.Errorf("ERR: Client API '%s' not found", code)
    return
  }
  w.Init(initInfo)
}

func workersClose() {
  for _, w := range modFuncs {
    w.Close()
  }
}

func workerClose(code string) {
  w, ok := modFuncs[code]
  if !ok {
    glog.Errorf("ERR: Client API '%s' not found", code)
    return
  }
  w.Close()
}

func workersPublicInfo() []PublicInfo {
  var res []PublicInfo
  for _, w := range modFuncs {
    res = append(res, w.GetPublicInfo())
  }
  return res
}

func workersResults() (map[string]map[string]interface{}) {
  res := make(map[string]map[string]interface{})
  for _, w := range modFuncs {
    res[w.GetAPI()] = w.GetResultArray()
  }
  return res
}

func workerRun(code string) {
  w, ok := modFuncs[code]
  if !ok {
    glog.Errorf("ERR: Client API '%s' not found\n", code)
    return
  }
  runMethodIfExists(code, w, "Start")
  runMethodIfExists(code, w, "GetData")
  runMethodIfExists(code, w, "Finish")
}

func getWorker(code string) *WorkerInterface {
  w, ok := modFuncs[code]
  if ok {
    return &w
  }
  return nil
}

func workerExists(code string) bool {
  _, ok := modFuncs[code]
  return ok
}

func runMethodIfExists(nameInterface string, any WorkerInterface, nameFunc string, args ...interface{}) ([]reflect.Value, bool) {
  v := reflect.ValueOf(any)
	method := v.MethodByName(nameFunc)
	if method.Kind() == reflect.Invalid {
    glog.Warningf("WRN: NOT FOUND runMethodIfExists(%s.%s)\n", nameInterface, nameFunc)
		return []reflect.Value{}, false
	}

	if method.Type().NumIn() != len(args) {
    glog.Errorf("ERR: runMethodIfExists(%s.%s): expected %d args, actually %d.\n", 
      nameInterface,
			nameFunc,
			len(args),
			method.Type().NumIn())
    return []reflect.Value{}, false
	}

	// Create a slice of reflect.Values to pass to the method. Simultaneously
	// check types.
	argVals := make([]reflect.Value, len(args))
	for i, arg := range args {
		argVal := reflect.ValueOf(arg)

		if argVal.Type() != method.Type().In(i) {
      glog.Errorf("ERR: runMethodIfExists(%s): expected arg %d to have type %v.\n", 
        nameFunc,
				i,
				argVal.Type())
		}

		argVals[i] = argVal
	}
  if glog.V(9) {
    glog.Infof("DBG: Call(%s.%s)\n", nameInterface, nameFunc)
  }
	return method.Call(argVals), true
}
