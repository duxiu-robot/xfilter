package xfilter

import (
	"context"

	"github.com/pkg/errors"

	"github.com/duxiu-robot/filter"
	_ "github.com/duxiu-robot/filter/ext/location"
	_ "github.com/duxiu-robot/filter/ext/request"
	_logger "github.com/duxiu-robot/logger"
)

var (
	ErrNilFilter         = errors.New("nil filter")
	ErrInvalidFilterData = errors.New("invalid filter data")

	logger = _logger.Named("xfilter")
)

func Run(ctx context.Context, f filter.Filter, data interface{}) (err error) {
	if f == nil {
		return ErrNilFilter
	}

	if data == nil {
		data = make(map[string]interface{})
	}

	f.Run(ctx, data)

	_resultObserverManager.observe(data)

	return
}
