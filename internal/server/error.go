package server

import "errors"

var (
	ErrCertificatePathEmpty             = errors.New("server: certificate path is empty")
	ErrAlreadyDeclareOnShutdownFunction = func(name string) error { return errors.New("server: already declare shutdown '" + name + "'") }
	ErrRunningServer                    = errors.New("server: server already running")
)
