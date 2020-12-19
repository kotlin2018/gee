// +build windows

package utils

import (
	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
	"goApp/internal/global"
	"os"
	"path"
	"time"
)

// 使用file-rotatelogs 进行日志分割 适用于windows系统
func GetWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := zaprotatelogs.New(
		path.Join(global.Config.Zap.Director, "%Y-%m-%d.log"),
		zaprotatelogs.WithMaxAge(7*24*time.Hour),
		zaprotatelogs.WithRotationTime(24*time.Hour),
	)
	if global.Config.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}