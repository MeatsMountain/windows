//+build !windows

package ctl

import "errors"

func (s *Server) Start() error {
	return errors.New("not implemented")
}

func isClosedError(err error) bool {
	return false
}
