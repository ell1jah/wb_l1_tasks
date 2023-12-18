package main

import (
	"io"
	"os"
)

// интерфейс логгера
type LoggerI interface {
	Info(string)
}

// тип использующий логгер
type Worker struct {
	Logger LoggerI
}

func (w *Worker) Work() {
	w.Logger.Info("some work is done\n")
}

// логер с другим интерфейсом
type ActualLogger struct{}

func (al *ActualLogger) Info(w io.Writer, text string) {
	w.Write([]byte(text))
}

// адаптер логера под интерфейс LoggerI
type LoggerAdapter struct {
	logger *ActualLogger
	w      io.Writer
}

func NewLoggerAdapter(logger *ActualLogger, w io.Writer) *LoggerAdapter {
	return &LoggerAdapter{
		logger: logger,
		w:      w,
	}
}

func (la *LoggerAdapter) Info(text string) {
	la.logger.Info(la.w, text)
}

func main() {
	al := &ActualLogger{}
	la := NewLoggerAdapter(al, os.Stdout)

	worker := Worker{la}
	worker.Work()
}
