package log

import (
"io"
"log"
"os"
)

var (
	Debug *log.Logger
	Info *log.Logger
	Warn *log.Logger
	Error *log.Logger
)

func InitLog(debug io.Writer, info io.Writer,  warn io.Writer){
	Debug = log.New(debug,"[Debug]", log.Ldate | log.Ltime | log.Lshortfile)
	Info = log.New(info,"[Info]", log.Ldate | log.Ltime | log.Lshortfile)
	Warn = log.New(warn,"[Warn]", log.Ldate | log.Ltime | log.Lshortfile)
	Error = log.New(os.Stderr,"[Error]", log.Ldate | log.Ltime | log.Lshortfile)
}
