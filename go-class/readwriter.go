package io

// You have to satisfy Reader and Writer interfaces to be a ReadWriter
type ReadWriter interface {
	Reader
	Writer
}
