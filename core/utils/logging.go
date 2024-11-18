package utils

import (
	"os"
	"sync"

	"log/slog"
)

var (
	slogger *slog.Logger
	once    sync.Once
)

// GetLogger はプロジェクト全体で使用する統一されたロガーを返します。
// ロガーがまだ初期化されていない場合は、初期化を行います。
func GetLogger() *slog.Logger {
	once.Do(func() {
		opts := slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug, // 必要に応じてログレベルを変更
		}
		handler := slog.NewTextHandler(os.Stdout, &opts)
		slogger = slog.New(handler)
	})
	return slogger
}
