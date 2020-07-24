package log

import (
	"github.com/rs/zerolog/log"
	"go.uber.org/zap"
)

func init() {
	zap.NewStdLog()
	log.Output()
}
