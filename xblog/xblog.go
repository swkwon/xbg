package xblog

import "errors"

// log level error, warning, info, debug
// error: An unusual situation, fail function, self recovery failure, need someone.
// warning: Self-recoverable, Need to fix it soon.
// info: Normal situation, Information needed for development and operation.
// debug: Information needed for debugging.
// Println: plain text. no timestamp, no log level.

// log format: json, log, text
type (
	ConsoleConfig struct {
		Enable bool   `json:"enable"`
		Format string `json:"format"`
	}

	FileConfig struct {
		Format     string `json:"format"`
		Rotation   string `json:"rotation"`
		RotateSize int    `json:"rotate_size"`
		Prefix     string `json:"prefix"`
		OutDir     string `json:"out_dir"`
	}

	NetworkConfig struct {
		Format          string `json:"format"`
		Host            string `json:"host"`
		HTTPMethod      string `json:"http_method"`
		HTTPContentType string `json:"http_content_type"`
	}

	CustomFileConfig struct {
		Key string `json:"key"`
		FileConfig
	}

	CustomNetworkConfig struct {
		Key string `json:"key"`
		NetworkConfig
	}

	Config struct {
		Console *ConsoleConfig `json:"console"`
	}

	XBLog interface {
		Error()
		Warning()
		Info()
		Debug()
	}

	xblogImpl struct {
	}

	LogData map[string]interface{}
)

func (l *xblogImpl) Error(args ...interface{}) {

}

func (l *xblogImpl) Warning(args ...interface{}) {

}

func (l *xblogImpl) Info(args ...interface{}) {

}

func (l *xblogImpl) Debug(args ...interface{}) {

}

func New(confing *Config) (XBLog, error) {
	return nil, errors.New("not implement")
}
