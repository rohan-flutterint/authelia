package suites

import (
	"testing"
	"time"

	"github.com/go-rod/rod"
	"github.com/stretchr/testify/require"
)

func (rs *RodSession) verifyIsSecondFactorPage(t *testing.T, page *rod.Page) {
	require.NoError(t, page.WaitStable(time.Millisecond*50))

	rs.WaitElementLocatedByID(t, page, "second-factor-stage")
}
