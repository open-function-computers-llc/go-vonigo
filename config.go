package vonigo

import (
	"github.com/Sirupsen/logrus"
)

// Config This is where we will add anything we need to make this library work
// the way we want it to
type Config struct {
	SecurityToken string
	Logger        *logrus.Logger
}
