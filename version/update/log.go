package main

import (
	_l "github.com/cybriq/log"

	"github.com/cybriq/log/version"
)

var log = _l.Get(_l.Add(version.PathBase))
