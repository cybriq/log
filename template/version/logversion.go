package version

import (
	"path/filepath"
	"runtime"

	_l "github.com/cybriq/log"
)

var log *_l.Logger

func init() {
	_, file,_, _ := runtime.Caller(0)
	verPath := filepath.Dir(file)+"/"
	log = _l.Get(_l.Add(verPath))
}
