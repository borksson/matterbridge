package bridge

import (
	"github.com/sirupsen/logrus"
)

type GatedLogger struct {
	logrus.Entry
	links []*LoggerLink
}		

func (l *GatedLogger) AddLink(link *LoggerLink) {
	if l.links == nil {
		l.links = make([]*LoggerLink, 0)
	}
	l.links = append(l.links, link)
}

func (l *GatedLogger) UpdateLinks(args ...interface{}) {
	for _, link := range l.links {
		link.Update(args...)
	}
}

func (l *GatedLogger) Info(args ...interface{}) {
	// Intercept log messages and gate them based on the log level.
	l.UpdateLinks(args...)
	l.Entry.Info(args...)
}