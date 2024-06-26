package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	//ログファイルの読み込み、読み書き、作成、追記
	logfile, err := os.OpenFile(logFile, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln((err))
	}
	//ログの書き込み先を標準出力とログファイルに指定
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	//ログのフォーマットを指定
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//ログの出力先を設定
	log.SetOutput(multiLogFile)
}