//go:build !testing
// +build !testing

package gorocksdb

// #cgo LDFLAGS: -lrocksdb -pthread -lstdc++ -ldl -lm -llz4 -lsnappy
import "C"
