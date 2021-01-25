package notification

import (
	"os"
	"testing"

	"github.com/goharbor/harbor/src/common/dao"
)

func TestMain(m *testing.M) {
	dao.PrepareTestForPostgreSQL()
	os.Exit(m.Run())
}
