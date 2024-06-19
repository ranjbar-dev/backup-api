package loggertool

import (
	"fmt"
	"log"

	"github.com/ranjbar-dev/backup-api/internal/config"
)

var level int

func init() {

	level = config.Single.LogLevel()
}

func Trace(service string, message string, payload any) {

	if level > 0 {
		return
	}

	data := []any{fmt.Sprintf("[INFO] %s - %s", service, message)}

	if payload != nil {
		data = append(data, payload)
	}

	log.Println(data...)
	fmt.Println(data...)
}

func Debug(service string, message string, payload any) {

	if level > 1 {
		return
	}

	data := []any{fmt.Sprintf("[INFO] %s - %s", service, message)}

	if payload != nil {
		data = append(data, payload)
	}

	log.Println(data...)
	fmt.Println(data...)
}

func Info(service string, message string, payload any) {

	if level > 2 {
		return
	}

	data := []any{fmt.Sprintf("[INFO] %s - %s", service, message)}

	if payload != nil {
		data = append(data, payload)
	}

	log.Println(data...)
	fmt.Println(data...)
}

func Warn(service string, message string, payload any) {

	if level > 3 {
		return
	}

	data := []any{fmt.Sprintf("[INFO] %s - %s", service, message)}

	if payload != nil {
		data = append(data, payload)
	}

	log.Println(data...)
	fmt.Println(data...)
}

func Error(service string, message string, payload any) {

	if level > 4 {
		return
	}

	data := []any{fmt.Sprintf("[INFO] %s - %s", service, message)}

	if payload != nil {
		data = append(data, payload)
	}

	log.Println(data...)
	fmt.Println(data...)
}

func Fatal(service string, message string, payload any) {

	if level > 5 {
		return
	}

	data := []any{fmt.Sprintf("[INFO] %s - %s", service, message)}

	if payload != nil {
		data = append(data, payload)
	}

	log.Println(data...)
	fmt.Println(data...)

	panic(payload)
}

func Panic(service string, message string, payload any) {

	if level > 6 {
		return
	}

	data := []any{fmt.Sprintf("[INFO] %s - %s", service, message)}

	if payload != nil {
		data = append(data, payload)
	}

	log.Println(data...)
	fmt.Println(data...)

	panic(payload)
}
