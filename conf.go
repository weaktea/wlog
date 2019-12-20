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

type WlogConfig struct {
	ToStderr         bool
	AlsoToStderr     bool
	CStderrThreshold string `xml:",comment"`
	StderrThreshold  severity
	TraceLocation    traceLocation
	CLogThreshold    string `xml:",comment"`
	LogThreshold     severity
	CFlushInterval   string `xml:",comment"`
	FlushInterval    int64
	LogDir           string
	CMaxSize         string `xml:",comment"`
	MaxSize          uint64
}

var config WlogConfig

func LoadXmlConfig() {
	config.StderrThreshold = ERROR
	config.LogThreshold = INFO
	config.LogDir = "./log/"
	config.MaxSize = 1024 * 1024 * 512
	file, err := os.Open(CONFIG_FILE_PATH)
	if err != nil {
		fmt.Println(CONFIG_FILE_PATH, " not exist!")
		GenConfigFile()
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

func GenConfigFile() {
	SERVERITY_COMMENT := "0:DEBUG 1:INFO 2:WARNING 3:ERROR 4:FATAL"
	config.CStderrThreshold = SERVERITY_COMMENT
	config.CLogThreshold = SERVERITY_COMMENT
	config.CFlushInterval = "ms"
	config.CMaxSize = "bytes"
	buff, err := xml.MarshalIndent(&config, "", "    ")
	if err != nil {
		fmt.Println("Marshal failed")
		return
	}

	err = createDir(CONFIG_FILE_DIR)
	if err != nil {
		fmt.Println("createDir failed ", err.Error())
		return
	}
	file, err := os.Create(CONFIG_FILE_PATH)
	if err != nil {
		fmt.Println(CONFIG_FILE_PATH, " Create failed!")
		return
	}
	defer file.Close()
	_, err = file.Write(buff)
	if err != nil {
		fmt.Println(CONFIG_FILE_PATH, " Write failed!")
	}
}
