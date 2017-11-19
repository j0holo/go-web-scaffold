// Package log is a simple log package with three loggers.
package log

import (
	"log"
	"os"
)

var (
	Info  *log.Logger
	Debug *log.Logger
	Error *log.Logger
)

const (
	flags = log.Ldate | log.Ltime | log.Lshortfile
)

// InitLogger is a small util function that returns a standard golang logger.
// If a path string is empty, the logger will log to os.StdOut. If a logfile
// cannot be created an error will be thrown to the the caller and all
// loggers will be set to nil. Failing fast is better then running without
//reporting any important errors.
func InitLogger(pathInfo string, pathDebug string, pathError string) ([3]*os.File, error) {
	var logErr error

	logFiles := [3]*os.File{}
	loggers := [3]*log.Logger{Info, Debug, Error}
	paths := [3]string{pathInfo, pathDebug, pathError}
	prefixes := [3]string{"INFO:", "DEBUG:", "ERROR:"}

	for i := 0; i < len(loggers); i++ {
		logFile, err := new(loggers[i], paths[i], prefixes[i])
		logFiles[i] = logFile
		if err != nil {
			logErr = err
			break
		}
	}

	if logErr != nil {
		for i := range loggers {
			loggers[i] = nil
			logFiles[i].Close()
		}
		return [3]*os.File{}, logErr
	}

	return [3]*os.File{}, nil
}

func new(logger *log.Logger, path string, prefix string) (*os.File, error) {
	logf, err := openLogFile(path)
	if logf != nil && err == nil {
		logger = log.New(logf, prefix, flags)
	} else {
		if err != nil {
			logger = nil
			return nil, err
		}
		if logf == nil {
			logger = log.New(os.Stderr, prefix, flags)
		}
	}

	return logf, nil
}

func openLogFile(path string) (*os.File, error) {
	if path == "" {
		return nil, nil
	}
	logf, err := os.OpenFile(path, (os.O_APPEND | os.O_WRONLY | os.O_CREATE), 0664)
	if err != nil {
		return nil, err
	}

	return logf, nil
}
