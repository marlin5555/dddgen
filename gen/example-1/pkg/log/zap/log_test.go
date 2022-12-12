package zap_test

import (
	"context"
	"testing"

	"github.com/marlin5555/dddgen/pkg/log/zap"
)

// TestGrpcLog test for grpc log
func TestGrpcLog(t *testing.T) {
	log := zap.InitLog("", "info")
	log.Infof("infof, add sth")
	log.Info("info")
	log.InfoContextf(context.TODO(), "infoContext")
	log.Error("error")
	log.Errorf("errorf")
	log.ErrorContextf(context.TODO(), "errorcontext")
	log.Warnf("warnf")
	log.Warn("warn")
	log.WarnContextf(context.TODO(), "warnContext")
	log.Debugf("debugf")
	log.Debug("debug")
	log.DebugContextf(context.TODO(), "debugcontext")
}
