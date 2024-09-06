package logger

import (
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type ZapOption struct {
	LogPath         string
	ApplicationName string
	Debug           bool
}

func NewOption(conf *viper.Viper) ZapOption {
	return ZapOption{
		LogPath:         conf.GetString("log.log_path"),
		ApplicationName: conf.GetString("log.application_name"),
		Debug:           conf.GetBool("log.debug"),
	}
}

func NewZapLogger(zapOption ZapOption) *zap.Logger {
	hook := lumberjack.Logger{
		Filename:   zapOption.LogPath, // 日誌檔案路徑
		MaxSize:    128,               // 每個日誌檔案儲存的大小 單位:M
		MaxAge:     7,                 // 檔案最多儲存多少天
		MaxBackups: 30,                // 日誌檔案最多儲存多少個備份
		Compress:   false,             // 是否壓縮
	}
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "file",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 短路徑編碼器
		EncodeName:     zapcore.FullNameEncoder,
	}
	// 設定日誌級別
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)
	var writes = []zapcore.WriteSyncer{zapcore.AddSync(&hook)}
	// 如果是開發環境，同時在控制檯上也輸出
	if zapOption.Debug {
		writes = append(writes, zapcore.AddSync(os.Stdout))
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(writes...),
		atomicLevel,
	)

	// 開啟開發模式，堆疊跟蹤
	caller := zap.AddCaller()
	// 跳過一層調用棧
	callerSkip := zap.AddCallerSkip(1)
	// 開啟檔案及行號
	development := zap.Development()

	// 設定初始化欄位
	field := zap.Fields(zap.String("ApplicationName", zapOption.ApplicationName))

	// 構造日誌
	zapLogger := zap.New(core, caller, callerSkip, development, field)
	zapLogger.Info("log 初始化成功")
	return zapLogger
}
