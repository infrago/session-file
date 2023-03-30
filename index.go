package session_file

import (
	"github.com/infrago/session"
)

func Driver() session.Driver {
	return &fileDriver{}
}

func init() {
	session.Register("file", Driver())
}
