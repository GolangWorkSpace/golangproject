package main

import (
	//log "github.com/brasbug/log4go"
	log "Project/log4go"
)


func SetLog() {

	//if err := log.SetupLogWithConf("log.json"); err != nil {
	//
	//	panic(err)
	//}

	w := log.NewFileWriter()
	w.SetPathPattern("./log/log-%Y%M%D.log")
	c := log.NewConsoleWriter()
	c.SetColor(true)
	log.Register(w)
	log.Register(c)
	log.SetLevel(log.DEBUG)
	log.SetLayout("2006-01-02 15:04:05")


}

func main() {

	 SetLog()

	defer log.Close()

	var name = "skoo"
	log.Debug("log4go by %s", name)
	log.Info("log4go by %s", name)
	log.Warn("log4go by %s", name)
	log.Error("log4go by %s", name)


	log.Fatal("log4go by %s", name)
}
