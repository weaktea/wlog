package wlog

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	CONFIG_FILE_DIR   = "./conf/"
	CONFIG_FILE_NAME  = "wlog.xml"
	CONFIG_FILE_PATH  = CONFIG_FILE_DIR + CONFIG_FILE_NAME
)

type ConfigParam struct {
	ToStderr         bool
	AlsoToStderr     bool
	StderrThreshold  severity
	TraceLocation    traceLocation
	LogThreshold     severity
	FlushInterval    int64
	LogDir           string
	MaxSize          uint64
}

var config ConfigParam

func LoadXmlConfig() {
	config.StderrThreshold = ERROR
	config.LogThreshold = INFO
	config.LogDir = "./log/"
	config.MaxSize = 1024 * 1024 * 512
	file, err := os.Open(CONFIG_FILE_PATH)
	if err != nil {
		fmt.Println(CONFIG_FILE_PATH, " not exist!")
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(CONFIG_FILE_PATH, " read failed!")
		return
	}
	err = xml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println(CONFIG_FILE_PATH, " Unmarshal failed!")
		return
	}
}
