package log
import (
	"os"
	"github.com/apex/log"
	"github.com/apex/log/handlers/json"
	"github.com/apex/log/handlers/text"
	"github.com/pkg/errors"
)
// Logger provides a leveled-logging interface
type Logger interface {
	log.Interface
}
// GetLogger create and return Logger instance
func GetLogger(format, level string) (Logger, error) {
	var handler log.Handler
	switch format {
	case "json":
		handler = json.New(os.Stdout)
	case "text":
		handler = text.New(os.Stdout)
	default:
		return nil, errors.Errorf("log: invalid format (%s)", format)
	}
	lvl, err := log.ParseLevel(level)
	if err != nil {
		return nil, errors.Wrapf(err, "log: invalid level (%s)", level)
	}
	logger := &log.Logger{
		Handler: handler,
		Level:   lvl,
	}
	return logger, nil
}
// Fields represents a map of entry level data used for structured logging
type Fields map[string]interface{}
// Fields implements log.Fielder
func (f Fields) Fields() log.Fields {
	return log.Fields(f)
}
