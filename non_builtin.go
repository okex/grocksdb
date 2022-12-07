//go:build !testing
// +build !testing

package gorocksdb

// #cgo LDFLAGS: -lrocksdb -pthread -lstdc++ -ldl -lm -lzstd -llz4 -lz -lsnappy
import "C"
