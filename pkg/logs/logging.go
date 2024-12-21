package logging

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path"
	"runtime"
	"strings"
	"template/pkg/env"
	"time"
)

type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (hook *writerHook) Fire(entry *logrus.Entry) (err error) {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range hook.Writer {
		_, err = w.Write([]byte(line))
	}
	return
}

func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}

type sendHook struct {
	client              *http.Client
	logLevels           []logrus.Level
	serviceURL          string
	telegramId          []string
	location            string
	ch                  *amqp.Channel
	SendingAvailability bool
}

func repScore(msg string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(msg, "\u001B[0;37m", ""), "\u001B[0m", ""), "\u001B[0;34m", ""), "\t", ""), "\"", "'"), "\\'", ""), "\n", "")
}

func (hook *sendHook) Fire(entry *logrus.Entry) (err error) {
	message := fmt.Sprintf("Time: %s; %s:%d %s", entry.Time.Format("02-01-2006 15:04.999"), entry.Caller.File, entry.Caller.Line, repScore(entry.Message))

	if hook.SendingAvailability {
		body := fmt.Sprintf("{\"message\":\"%s\", \"timeEvent\":%d,\"receiver_id\":\"%v\",\"level\":\"%s\",\"location\":\"%s\", \"writeDB\":true}", message, time.Now().Unix(), hook.telegramId, entry.Level, hook.location)
		if hook.ch != nil {
			err := hook.ch.PublishWithContext(context.TODO(),
				"logs", // exchange
				"",     // routing key
				false,  // mandatory
				false,  // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(body),
					Headers:     map[string]interface{}{"type": "logger"},
				})
			if err != nil {
				color.Magenta(fmt.Sprintf("Error request %s: %v", hook.serviceURL, err))
				hook.SendingAvailability = false
				go func() {
				reconnect:
					for {
						ch, err := ConnectToRabbit()
						if err != nil {
							color.Red(err.Error())
							time.Sleep(30 * time.Second)
							continue
						}
						hook.SendingAvailability = true
						hook.ch = ch
						break reconnect
					}
				}()
			}
		}
	}
	return err
}

//func (hook *sendHook) Fire(entry *logrus.Entry) error {
//	message := fmt.Sprintf("Time: %s; %s:%d %s", entry.Time.Format("02-01-2006 15:04.999"), entry.Caller.File, entry.Caller.Line, repScore(entry.Message))
//
//	req, err := http.NewRequest("POST", hook.serviceURL,
//		bytes.NewBuffer([]byte(fmt.Sprintf("{\"message\":\"%s\",\"id\":\"%v\",\"level\":\"%s\",\"location\":\"%s\", \"writeDB\":true}",
//			message, hook.telegramId, entry.Level, hook.location))))
//	if err != nil {
//		color.Magenta(fmt.Sprintf("Error new request %s: %v", hook.serviceURL, err))
//		return err
//	}
//	req.Header.Set("Content-Type", "application/json")
//
//	_, err = hook.client.Do(req)
//	if err != nil {
//		color.Magenta(fmt.Sprintf("Error request %s: %v", hook.serviceURL, err))
//	}
//	return err
//}

func (hook *sendHook) Levels() []logrus.Level {
	return hook.logLevels
}

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
	LogLevel string
}

func GetLogger() *Logger {
	return &Logger{Entry: e, LogLevel: strings.ToLower(env.GetEnv("LOG_LEVEL", "ERROR"))}
}

func (l *Logger) GetLoggerWithField(k string, v interface{}) *Logger {
	return &Logger{Entry: l.WithField(k, v)}
}

func ConnectToRabbit() (ch *amqp.Channel, err error) {
	conn, err := amqp.Dial("amqp://ext-user:dYnQ5zsk@77.50.72.170:15476/")
	if err != nil {
		return nil, err
	}

	ch, err = conn.Channel()
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"logs-services", // name
		false,           // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		return nil, err
	}

	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		"logs", // exchange
		false,
		nil)
	if err != nil {
		return nil, err
	}
	return
}

func InitLogger() {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		ForceColors: true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			if env.GetEnvAsBool("LOG_COLOR", false) {
				return fmt.Sprintf("\u001B[0;37m%s()\u001B[0m", frame.Function), fmt.Sprintf("\u001B[0;34m%s:%d\u001B[0m", filename, frame.Line)
			} else {
				return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
			}
		},
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "02-01-2006 15:04.999",
	}

	if env.GetEnvAsBool("LOG_WRITE", true) {
		err := os.MkdirAll("logs", 0644)
		if err != nil {
			color.Red(err.Error())
			panic(err)
		}

		err = os.Chmod("logs", 0777)
		if err != nil {
			color.Red(err.Error())
		}

		allFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
		if err != nil {
			color.Red(err.Error())
		}

		l.SetOutput(io.Discard)

		l.AddHook(&writerHook{
			Writer:    []io.Writer{allFile, os.Stdout},
			LogLevels: logrus.AllLevels,
		})
	}

	if location := env.GetEnv("LOG_LOCATION", ""); location != "" {
		color.Magenta("Start send hook")
		ch, err := ConnectToRabbit()
		if err != nil {
			color.Red(err.Error())
			l.AddHook(&sendHook{
				client:              &http.Client{Timeout: 2 * time.Second},
				logLevels:           []logrus.Level{logrus.FatalLevel, logrus.PanicLevel, logrus.ErrorLevel},
				serviceURL:          env.GetEnv("LOG_SERVICE_URL", ""),
				telegramId:          env.GetEnvAsSlice("LOG_TELEGRAM_RECIPIENT", []string{}, ","),
				location:            location,
				ch:                  nil,
				SendingAvailability: false,
			})

			go func() {
			reconnect:
				for {
					time.Sleep(60 * time.Second)
					ch, err := ConnectToRabbit()
					if err != nil {
						color.Red(err.Error())
						time.Sleep(30 * time.Second)
						continue
					}
					l.AddHook(&sendHook{
						client:              &http.Client{Timeout: 2 * time.Second},
						logLevels:           []logrus.Level{logrus.FatalLevel, logrus.PanicLevel, logrus.ErrorLevel},
						serviceURL:          env.GetEnv("LOG_SERVICE_URL", ""),
						telegramId:          env.GetEnvAsSlice("LOG_TELEGRAM_RECIPIENT", []string{}, ","),
						location:            location,
						ch:                  ch,
						SendingAvailability: true,
					})
					break reconnect
				}
			}()
		} else {
			l.AddHook(&sendHook{
				client:              &http.Client{Timeout: 2 * time.Second},
				logLevels:           []logrus.Level{logrus.FatalLevel, logrus.PanicLevel, logrus.ErrorLevel},
				serviceURL:          env.GetEnv("LOG_SERVICE_URL", ""),
				telegramId:          env.GetEnvAsSlice("LOG_TELEGRAM_RECIPIENT", []string{}, ","),
				location:            location,
				ch:                  ch,
				SendingAvailability: true,
			})
		}
	}

	switch strings.ToLower(env.GetEnv("LOG_LEVEL", "ERROR")) {
	case "info":
		l.SetLevel(logrus.InfoLevel)
	case "debug":
		l.SetLevel(logrus.DebugLevel)
	case "trace":
		l.SetLevel(logrus.TraceLevel)
	case "error":
		l.SetLevel(logrus.ErrorLevel)
	case "panic":
		l.SetLevel(logrus.PanicLevel)
	case "fatal":
		l.SetLevel(logrus.FatalLevel)
	default:
		l.SetLevel(logrus.InfoLevel)
	}

	e = logrus.NewEntry(l)
}
