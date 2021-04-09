package log

import (
	"gBlog/common/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func Init() {
	hook := lumberjack.Logger{
		Filename:   config.GetLOGConfig().Path, // 日志文件路径
		MaxSize:    128,                        // megabytes
		MaxBackups: 30,                         // 最多保留300个备份
		MaxAge:     7,                          // days
		Compress:   true,                       // 是否压缩 disabled by default
	}

	w := zapcore.AddSync(&hook)

	// 设置日志级别,debug可以打印出info,debug,warn；info级别可以打印warn，info；warn只能打印warn
	// debug->info->warn->error
	var level zapcore.Level
	switch config.GetLOGConfig().Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	// 时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		w,
		level,
	)

	Log = zap.New(core)
	Log.Info("DefaultLogger init success")
}

func GetLog() *zap.Logger {
	return Log
}
