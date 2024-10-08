package configs

import (
	"fmt"
	"io"
	"os"
	"time"

	"aichat/configs/initstruct"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger 全局变量，导出给调用者使用
var Logger *zap.Logger

// Init 初始化日志相关目录
func LoggerInit(config initstruct.LoggerConfig) error {
	// 创建日志根目录
	if _, err := os.Stat(config.BaseDirectoryName); os.IsNotExist(err) {
		err := os.MkdirAll(config.BaseDirectoryName, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error creating directory, err: %v", err)
		}
	}

	// 创建日志子目录
	if err := os.MkdirAll(fmt.Sprintf("%s/%s", config.BaseDirectoryName, config.InfoDirectoryName), os.ModePerm); err != nil {
		return fmt.Errorf("error creating info directory, err: %v", err)
	}
	if err := os.MkdirAll(fmt.Sprintf("%s/%s", config.BaseDirectoryName, config.WarnDirectoryName), os.ModePerm); err != nil {
		return fmt.Errorf("error creating warn directory, err: %v", err)
	}
	if err := os.MkdirAll(fmt.Sprintf("%s/%s", config.BaseDirectoryName, config.ErrorDirectoryName), os.ModePerm); err != nil {
		return fmt.Errorf("error creating err directory, err: %v", err)
	}

	// 自定义初始化zap库
	initLogger(config)

	return nil
}

// getWriter 获取wirter文件写入
func getWriter(logBasePath, logLevelPath, LogFileName string, config initstruct.LoggerConfig) io.Writer {
	return &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s/%s", logBasePath, logLevelPath, LogFileName),
		MaxBackups: config.LogFileMaxBackups,
		MaxSize:    config.LogFileMaxSize,
		MaxAge:     config.LogFileMaxAge,
		Compress:   config.LogFileCompress,
	}
}

// initLog 初始化日志
func initLogger(c initstruct.LoggerConfig) {
	// 获取io.Writer实现
	infoWriter := getWriter(c.BaseDirectoryName, c.InfoDirectoryName, c.InfoFileName, c)
	warnWriter := getWriter(c.BaseDirectoryName, c.WarnDirectoryName, c.WarnFileName, c)
	errWriter := getWriter(c.BaseDirectoryName, c.ErrorDirectoryName, c.ErrorFileName, c)

	// 获取日志默认配置
	encoderConfig := zap.NewProductionEncoderConfig()

	// 自定义日志输出格式
	// 修改TimeKey
	encoderConfig.TimeKey = "time"
	// 修改MessageKey
	encoderConfig.MessageKey = "message"
	// 时间格式符合人类观看
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	// 日志等级大写INFO
	encoderConfig.EncodeLevel = func(lvl zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(lvl.CapitalString())
	}
	// 日志打印时所处代码位置
	encoderConfig.EncodeCaller = func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(caller.TrimmedPath())
	}
	// 加载自定义配置为json格式
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	// 自定义日志级别 info
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zap.InfoLevel && lvl >= zap.DebugLevel
	})
	// 自定义日志级别 warn
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zap.WarnLevel && lvl > zap.InfoLevel
	})
	// 自定义日志级别 err
	errLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zap.FatalLevel && lvl > zap.WarnLevel
	})

	// 日志文件输出位置
	var core zapcore.Core
	if c.LogPrintTag {
		//同时在文件和终端输出日志
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel), // info级别日志
			zapcore.NewCore(encoder, zapcore.AddSync(warnWriter), warnLevel), // warn级别日志
			zapcore.NewCore(encoder, zapcore.AddSync(errWriter), errLevel),   // error级别日志
			zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), zap.DebugLevel),
		)
	} else {
		//只在文件输出日志
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel), // info级别日志
			zapcore.NewCore(encoder, zapcore.AddSync(warnWriter), warnLevel), // warn级别日志
			zapcore.NewCore(encoder, zapcore.AddSync(errWriter), errLevel),   // error级别日志
		)
	}

	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}
