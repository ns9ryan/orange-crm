package initialize

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"orange-crm/global"
	"os"
	"strings"
	"time"
)

// 创建一个通用的 LevelEnablerFunc 工厂函数
func createEnabler(minLevel zapcore.Level) zap.LevelEnablerFunc {
	return func(lev zapcore.Level) bool {
		return lev == minLevel
	}
}

// Zap 获取 zap.Logger
func Zap() (logger *zap.Logger) {
	cores := []zapcore.Core{
		// debug日志
		getCore("debug.log", createEnabler(zap.DebugLevel)),
		// info日志
		getCore("info.log", createEnabler(zap.InfoLevel)),
		// warn日志
		getCore("warn.log", createEnabler(zap.WarnLevel)),
		// error日志
		getCore("error.log", createEnabler(zap.ErrorLevel)),
		// dPanic日志
		getCore("dpanic.log", createEnabler(zap.DPanicLevel)),
		// panic日志
		getCore("panic.log", createEnabler(zap.PanicLevel)),
		// fatal日志
		getCore("fatal.log", createEnabler(zap.FatalLevel)),
	}

	return zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())
}

// 获取核心
func getCore(logName string, level zap.LevelEnablerFunc) zapcore.Core {
	writeSyncer := getWriteSyncer(logName)

	fileEncoder := getEncoder(false)
	fileCore := zapcore.NewCore(fileEncoder, writeSyncer, level)

	// 输出控制台
	if global.CONFIG.Zap.LogInConsole {
		consoleEncoder := getEncoder(true)
		consoleCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level)
		// 合并这两个 Core
		return zapcore.NewTee(consoleCore, fileCore)
	}

	return fileCore
}

// 获取编码器
func getEncoder(useCore bool) zapcore.Encoder {
	if global.CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig(useCore))
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig(useCore))
}

// 获取编码器配置
func getEncoderConfig(useCore bool) zapcore.EncoderConfig {
	config := zapcore.EncoderConfig{
		TimeKey:        "time",                         // 设置日志中时间部分的键名。
		LevelKey:       "level",                        // 设置日志中等级部分的键名。
		NameKey:        "name",                         // 设置日志记录器名称的键名。
		CallerKey:      "caller",                       // 设置调用者信息（文件名和行号）的键名。
		FunctionKey:    zapcore.OmitKey,                // 设置调用函数名的键名。
		MessageKey:     "message",                      // 设置日志中消息部分的键名。如果为空字符串，日志消息将被省略。
		StacktraceKey:  "stacktrace",                   // 设置堆栈追踪信息的键名。
		LineEnding:     zapcore.DefaultLineEnding,      // 指定日志条目之间的分隔符（通常是换行符，如 \n）。
		EncodeTime:     customTimeEncoder,              // 定义时间的编码方式，例如 ISO8601 字符串或 Unix 时间戳。
		EncodeCaller:   zapcore.FullCallerEncoder,      // 定义调用者信息的编码方式。
		EncodeDuration: zapcore.SecondsDurationEncoder, // 定义时间间隔（duration）的编码方式。
	}

	// 定义日志等级的编码方式，例如可以用字符串、数字或自定义格式。
	switch global.CONFIG.Zap.EncodeLevel {
	case "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case "LowercaseColorLevelEncoder": // 小写编码器带颜色
		if useCore {
			config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
		} else {
			config.EncodeLevel = zapcore.LowercaseLevelEncoder
		}
	case "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case "CapitalColorLevelEncoder": // 大写编码器带颜色
		if useCore {
			config.EncodeLevel = zapcore.CapitalColorLevelEncoder
		} else {
			config.EncodeLevel = zapcore.CapitalLevelEncoder
		}
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}

	return config
}

// 自定义日志输出时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(global.CONFIG.Zap.Prefix + t.Format("2006-01-02 15:04:05.000"))
}

// 获取写入同步器
func getWriteSyncer(logName string) zapcore.WriteSyncer {
	fileName := strings.Join([]string{global.CONFIG.Zap.Dir, logName}, "/")
	logWriter := &lumberjack.Logger{
		Filename:   fileName, // 日志文件路径
		MaxSize:    2,        // 单个日志文件的最大尺寸（单位：MB）
		MaxBackups: 10,       // 保留旧日志文件的最大数量
		MaxAge:     30,       // 保留旧日志文件的最大天数
		Compress:   true,     // 是否压缩旧日志文件
	}
	return zapcore.AddSync(logWriter)
}
