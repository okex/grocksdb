package gorocksdb

import "testing"

func TestOptEnv(t *testing.T) {
	opt := NewDefaultEnvOptions()
	defer opt.Destroy()
}
