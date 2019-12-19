package wlog

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

var (
	MaxSize  uint64 = 1024 * 1024 * 1800 // MaxSize is the maximum size of a log file in bytes.
	logDir   string
	pid      = os.Getpid()
	program  = filepath.Base(os.Args[0])
	host     = "unknownhost"

	userName = "unknownuser"
)

func init() {
	logDir = *flag.String("log_dir", "./log/", "If non-empty, write log files in this directory")
	h, err := os.Hostname()
	if err == nil {
		host = shortHostname(h)
	}

	current, err := user.Current()
	if err == nil {
		userName = current.Username
	}

	// Sanitize userName since it may contain filepath separators on Windows.
	userName = strings.Replace(userName, `\`, "_", -1)
}

// shortHostname returns its argument, truncating at the first period.
// For instance, given "www.google.com" it returns "www".
func shortHostname(hostname string) string {
	if i := strings.Index(hostname, "."); i >= 0 {
		return hostname[:i]
	}
	return hostname
}

// logName returns a new log file name containing tag, with start time t, and
// the name for the symlink for tag.
func logName(t time.Time) (name string) {
	name = fmt.Sprintf("%s.%04d%02d%02d-%02d%02d%02d.%d.log",
		program,
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		pid)
	return name
}

// create creates a new log file and returns the file and its filename, which
// contains tag ("INFO", "FATAL", etc.) and t.  If the file is created
// successfully, create also attempts to update the symlink for that tag, ignoring
// errors.
func create(t time.Time) (f *os.File, filename string, err error) {
	if len(logDir) == 0 {
		return nil, "", errors.New("log: no log dir")
	}
	s, err := os.Stat(logDir)
	if err != nil{
		if !os.IsExist(err){//不存在
			os.Mkdir(logDir, os.ModePerm)
		}
	}else{
		if !s.IsDir(){//是个文件
			return nil, "", errors.New("logDir is a file")
		}
	}

	name := logName(t)

	fname := filepath.Join(logDir, name)
	f, err = os.Create(fname)
	if err == nil {
		return f, fname, nil
	}

	return nil, "", fmt.Errorf("log: cannot create log: %v", err)
}
