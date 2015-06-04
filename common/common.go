package common

import (
    "strconv"
    
    defaultLog "log"
)

type LogWrap struct {    
    WorkerId int
}

func (l *LogWrap) Printf( format string, v ...interface{}) {
    defaultLog.Printf( "Worker " + strconv.Itoa( l.WorkerId ) + ": " +  format, v... )
}

func (l *LogWrap) Fatalf( format string, v ...interface{}) {
    defaultLog.Fatalf( "Worker " + strconv.Itoa( l.WorkerId ) + ": " +  format, v... )
}

func (l *LogWrap) Panicf( format string, v ...interface{}) {
    defaultLog.Panicf( "Worker " + strconv.Itoa( l.WorkerId ) + ": " +  format, v... )
}