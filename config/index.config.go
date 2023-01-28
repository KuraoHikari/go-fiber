package config

import (
	"path/filepath"
	"runtime"
)

var (
	//Get current file full from runtime

	_, b, _, _ = runtime.Caller(0)

	ProjectRootPath = filepath.Join(filepath.Dir(b), "../")
)
