package biz

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/moyrne/delay-record/pkg/pingclient/internal/repo"
	"github.com/moyrne/delay-record/pkg/pingserver/params"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log = newLogger()

func newLogger() *zap.Logger {
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "time",
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "logger",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), os.Stdout, zap.DebugLevel)
	return zap.New(core)
}

type Ping struct {
	Host   string
	Client *http.Client
	Data   repo.PingRepo
}

func NewPingClient(host string, data repo.PingRepo) *Ping {
	return &Ping{
		Host:   host,
		Client: &http.Client{Timeout: time.Second * 3},
		Data:   data,
	}
}

func (p *Ping) Ping() {
	resp := p.ping()
	log.Info("ping", zap.String("host", p.Host), zap.Error(resp.Error))

	if err := p.Data.InsertRecord(&repo.Ping{
		StartTime: resp.StartTime,
		Start:     resp.StartTime.Unix(),
		EndTime:   resp.EndTime,
		End:       resp.EndTime.Unix(),
		Delay:     resp.EndTime.Sub(resp.StartTime).Microseconds(),
		Error:     resp.Error,
	}); err != nil {
		log.Error("ping record insert", zap.Error(err))
		return
	}
}

type pingResponse struct {
	StartTime time.Time
	EndTime   time.Time
	Error     error
}

func (p *Ping) ping() pingResponse {
	startTime := time.Now()
	resp, err := p.Client.Get(p.Host)
	if err != nil {
		return pingResponse{StartTime: startTime, EndTime: time.Now(), Error: errors.Wrap(err, p.Host)}
	}
	pingResp := pingResponse{StartTime: startTime, EndTime: time.Now()}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		pingResp.Error = errors.WithStack(err)
		return pingResp
	}

	var ping params.PingResponse
	if err := json.Unmarshal(data, &ping); err != nil {
		pingResp.Error = errors.WithStack(err)
		return pingResp
	}

	return pingResp
}
