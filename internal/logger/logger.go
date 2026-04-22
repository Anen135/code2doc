package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func initLogger() error {
	if err := os.MkdirAll("log", 0755); err != nil {
		return err
	}

	logFiles = make(map[Level]*os.File, 4)
	levelsUsed = make(map[Level]bool, 4)
	logChan = make(chan entry, 100)
	closed = false

	wg.Add(1)
	go writeLoop()
	logChan <- entry{level: INFO, message: formatLine(INFO, "Logger initialized")}
	return nil
}

func writeLoop() {
	defer wg.Done()

	for e := range logChan {
		file := logFiles[e.level]

		if file == nil && !levelsUsed[e.level] {
			dir := filepath.Join("log", e.level.String())
			if err := os.MkdirAll(dir, 0755); err != nil {
				continue
			}

			timestamp := time.Now().Format("2006-01-02_15-04-05")
			filename := filepath.Join(dir, fmt.Sprintf("code2doc_%s.log", timestamp))

			f, err := os.Create(filename)
			if err != nil {
				continue
			}

			logFiles[e.level] = f
			levelsUsed[e.level] = true
			file = f
		}

		if file != nil {
			_, _ = file.WriteString(e.message)
		}
	}

	for _, file := range logFiles {
		if file != nil {
			_ = file.Sync()
		}
	}
}

func formatLine(level Level, message string) string {
	pc, file, line, ok := runtime.Caller(3)
	funcName := "unknown"
	if ok {
		if fn := runtime.FuncForPC(pc); fn != nil {
			funcName = fn.Name()
		}
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")

	if ok && file != "" {
		return fmt.Sprintf("[%s] [%s] [%s:%d] (%s) %s\n", timestamp, level, file, line, funcName, message)
	}
	return fmt.Sprintf("[%s] [%s] %s\n", timestamp, level, message)
}

func Log(level Level, message string) {
	once.Do(func() {
		if err := initLogger(); err != nil {
			fmt.Fprintf(os.Stderr, "CRITICAL: Failed to initialize logger: %v\n", err)
		}
	})
	line := formatLine(level, message)

	mu.RLock()
	defer mu.RUnlock()

	if closed || logChan == nil {
		fmt.Fprintf(os.Stderr, "CRITICAL (logger is closed): %s\n", line)
		return
	}

	logChan <- entry{level: level, message: line}
}

func Done() {
	mu.Lock()
	if closed || logChan == nil {
		mu.Unlock()
		return
	}
	closed = true
	logChan <- entry{
		level:   INFO,
		message: formatLine(INFO, "Logger closed"),
	}

	close(logChan)
	mu.Unlock()

	wg.Wait()

	for level, file := range logFiles {
		if file != nil {
			_ = file.Sync()
			_ = file.Close()
		} else if !levelsUsed[level] {
			dir := filepath.Join("log", level.String())
			_ = os.Remove(dir)
		}
	}
}
