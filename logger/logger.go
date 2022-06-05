package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Logger struct {
	file *os.File
}

// Trace — вывод всего подряд. На тот случай, если Debug не позволяет локализовать ошибку. В нем полезно отмечать вызовы разнообразных блокирующих и асинхронных операций.
// Debug — журналирование моментов вызова «крупных» операций. Старт/остановка потока, запрос пользователя и т.п.
// Info — разовые операции, которые повторяются крайне редко, но не регулярно. (загрузка конфига, плагина, запуск бэкапа)
// Warning — неожиданные параметры вызова, странный формат запроса, использование дефолтных значений в замен не корректных. Вообще все, что может свидетельствовать о не штатном использовании.
// Error — повод для внимания разработчиков. Тут интересно окружение конкретного места ошибки.
// Fatal — тут и так понятно. Выводим все до чего дотянуться можем, так как дальше приложение работать не будет.

func NewLogger() *Logger {
	if file, e := os.OpenFile(time.Now().Format("server_log_2006_01_02")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0660); e == nil {
		l := &Logger{file: file}
		l.Info("logger start")
		go l.fileChange((86400 - (time.Now().Unix() % 86400) - 3600*5) + 300)
		return l
	}
	log.Println("log file not created")
	os.Exit(9999)
	return nil
}

func (l *Logger) fileChange(t int64) {
	time.Sleep(time.Second * time.Duration(t))
	if file, e := os.OpenFile(time.Now().Format("server_log_2006_01_02")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0660); e == nil {
		l.file.Close()
		l.file = file
	}
	go l.fileChange(86400)
}

func (l *Logger) Trace(v interface{}) {
	if _, e := l.file.WriteString(fmt.Sprintf("%s [TRACE] %v\n", time.Now().Format("2006-01-02 15-04-05"), v)); e != nil {
		log.Println(e)
	}
}
func (l *Logger) Debug(v interface{}) {
	if _, e := l.file.WriteString(fmt.Sprintf("%s [DEBUG] %v\n", time.Now().Format("2006-01-02 15-04-05"), v)); e != nil {
		log.Println(e)
	}
}
func (l *Logger) Info(v interface{}) {
	if _, e := l.file.WriteString(fmt.Sprintf("%s [ INFO] %v\n", time.Now().Format("2006-01-02 15-04-05"), v)); e != nil {
		log.Println(e)
	}
}
func (l *Logger) Warn(v interface{}) {
	if _, e := l.file.WriteString(fmt.Sprintf("%s [ WARN] %v\n", time.Now().Format("2006-01-02 15-04-05"), v)); e != nil {
		log.Println(e)
	}
}
func (l *Logger) Error(v interface{}) {
	if _, e := l.file.WriteString(fmt.Sprintf("%s [ERROR] %v\n", time.Now().Format("2006-01-02 15-04-05"), v)); e != nil {
		log.Println(e)
	}
}
func (l *Logger) Fatal(v interface{}) {
	if _, e := l.file.WriteString(fmt.Sprintf("%s [FATAL] %v\n", time.Now().Format("2006-01-02 15-04-05"), v)); e != nil {
		log.Println(e)
	}
	os.Exit(1)
}
func (l *Logger) Panic(v interface{}) {
	if _, e := l.file.WriteString(fmt.Sprintf("%s [PANIC] %v\n", time.Now().Format("2006-01-02 15-04-05"), v)); e != nil {
		log.Println(e)
	}
	os.Exit(1)
}
