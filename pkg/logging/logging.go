package logging

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

type writeHook struct {
	Writer []io.Writer
	LogLevels []logrus.Level
}

func (w *writeHook) Fire(entry *logrus.Entry) error {

	line,err :=entry.String()
	if err!=nil {
		return err
	}
	for _,v:=range w.Writer{
		v.Write([]byte(line))

	}
	return err
	
}

func (w *writeHook) Levels()  []logrus.Level {
	return w.LogLevels
	
}

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLoger() Logger {
	return Logger{
		Entry: e,
	}
}

func (log *Logger) GetLoggerWithField(k string ,v interface{}) Logger  {
	return Logger{
		Entry: log.WithField(k,v),
	}
}


func init() {

	log := logrus.New()
	log.SetReportCaller(true)
	log.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			filename:=path.Base(f.File)
			return fmt.Sprintf("%s()",f.Function), fmt.Sprintf("%s:%d",filename,f.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	}
	err:=os.MkdirAll("logs",0644)
	if err!=nil {
		panic(err)
	}

	allFile,err:=os.OpenFile("logs/all.log",os.O_CREATE|os.O_APPEND|os.O_WRONLY,0640)
    fmt.Println(allFile)
	if err!=nil {
		panic(err)
	}
 
	log.SetOutput(io.Discard)
	log.AddHook(&writeHook{
		Writer: []io.Writer{allFile,os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	// wrt:=io.MultiWriter(os.Stdout,allFile)
	// log.SetOutput(wrt)
	log.SetLevel(logrus.TraceLevel)

	e = logrus.NewEntry(log)
}