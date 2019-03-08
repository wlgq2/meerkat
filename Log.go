package meerkat

import (
	"log"
	"sync"
	"io"
	"os"
)

const (
	LevelDebug int = 0
	LevelInfo int = 1
	LevelWarn int = 2
	LevelError int = 3
	LevelFatel int = 4
	LevelSize int = 5
)

type Log struct{
	hanles  [LevelSize]*log.Logger
	CurrentLever int
}
var logObj   *Log
var  logOnce sync.Once

func LogInstance() *Log {
	logOnce.Do(func() {
		logObj = &Log {}
		logObj.CurrentLever = 0
		logObj.DefaultInit()
    })
    return logObj
}

func (obj *Log)Init(out io.Writer, flag int)  {
	obj.hanles[LevelDebug] = log.New(out,"[Debug]",flag)
	obj.hanles[LevelInfo] = log.New(out,"[Info]",flag)
	obj.hanles[LevelWarn] = log.New(out,"[Warn]",flag)
	obj.hanles[LevelError] = log.New(out,"[Error]",flag)
	obj.hanles[LevelFatel] = log.New(out,"[Fatal]",flag)
}

func (obj *Log)DefaultInit()  {
	out,err  := os.Create("meerkat.log")
	//defer out.Close()
	if err == nil{
		obj.hanles[LevelDebug] = log.New(out,"[Debug]",log.Ldate | log.Ltime)
		obj.hanles[LevelInfo] = log.New(out,"[Info]",log.Ldate | log.Ltime)
		obj.hanles[LevelWarn] = log.New(out,"[Warn]",log.Ldate | log.Ltime)
		obj.hanles[LevelError] = log.New(out,"[Error]",log.Ldate | log.Ltime)
		obj.hanles[LevelFatel] = log.New(out,"[Fatal]",log.Ldate | log.Ltime)
	}else {
		out.Close()
	}
}

func (obj *Log) Fatalln(content ...interface{})  {
	obj.Writeln(LevelFatel,content)
}

func (obj *Log) Errorln(content ...interface{})  {
	obj.Writeln(LevelError,content)
}

func (obj *Log)Warnln(content ...interface{})  {
	obj.Writeln(LevelWarn,content)
}

func (obj *Log)Infoln(content ...interface{})  {
	obj.Writeln(LevelInfo,content)
}

func (obj *Log)Debugln(content ...interface{})  {
	obj.Writeln(LevelDebug,content)
}


func (obj *Log) Writeln(level int,content ...interface{})  {
	switch level {
	case  LevelFatel :
		if obj.CurrentLever<=LevelFatel && obj.hanles[LevelFatel] != nil {
			obj.hanles[LevelFatel].Fatalln(content)
		}
	case LevelError:
		fallthrough
	case LevelWarn:
		fallthrough
	case LevelInfo:
		fallthrough
	case LevelDebug:
		if obj.CurrentLever<=level && obj.hanles[level] != nil {
			obj.hanles[level].Println(content)
		}
	default:
	}
}