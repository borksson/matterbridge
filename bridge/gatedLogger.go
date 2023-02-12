package bridge

import (
	"github.com/sirupsen/logrus"
)

type GatedLogger struct {
	logrus.Entry
}