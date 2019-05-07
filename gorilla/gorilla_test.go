package gorilla_test

import (
	"testing"

	"github.com/gorilla/sessions"
	"github.com/itech-eng/session/gorilla"
	"github.com/itech-eng/session/test"
)

func TestAll(t *testing.T) {
	engine := sessions.NewCookieStore([]byte("something-very-secret"))
	manager := gorilla.New("_session", engine)
	test.TestAll(manager, t)
}
