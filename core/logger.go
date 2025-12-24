package core

import (
	"github.com/mattn/go-colorable"
	log "github.com/sirupsen/logrus"
)

func Init(level log.Level) {
	log.SetOutput(colorable.NewColorableStdout())
	log.SetLevel(level)

	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true, // always color
		FullTimestamp:  true,
		TimestampFormat: "2006-01-02 15:04:05",
		PadLevelText:   true,
	})
}
