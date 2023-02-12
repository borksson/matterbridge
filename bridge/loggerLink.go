package bridge

import (
)

type LoggerLink struct {
	gl *GatedLogger
}

func NewLoggerLink(gl *GatedLogger) *LoggerLink {
	l := &LoggerLink{
		gl: gl,
	}
	gl.AddLink(l)
	return l
}

func (l *LoggerLink) Update(args...interface{}) {
}