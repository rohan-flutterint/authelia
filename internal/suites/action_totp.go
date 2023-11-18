package suites

import (
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (rs *RodSession) doRegisterTOTP(t *testing.T, page *rod.Page) string {
	rs.doSettingsOpen(t, page)
	rs.doSettingsMenuTwoFactor(t, page)

	require.NoError(t, rs.WaitElementLocatedByID(t, page, "one-time-password-add").Click("left", 1))

	// TODO: dynamically handle OTC and 2FA requests in this test.
	element := rs.WaitElementLocatedByID(t, page, "one-time-code")
	code := doGetOneTimeCodeFromLastMail(t)

	require.NoError(t, element.Type(rs.toInputs(code)...))

	require.NoError(t, rs.WaitElementLocatedByID(t, page, "dialog-verify").Click("left", 1))
	require.NoError(t, rs.WaitElementLocatedByID(t, page, "dialog-next").Click("left", 1))
	require.NoError(t, rs.WaitElementLocatedByID(t, page, "qr-toggle").Click("left", 1))

	secretURLElement := rs.WaitElementLocatedByID(t, page, "secret-url")

	secretURLRaw, err := secretURLElement.Text()
	require.NoError(t, err)

	secretURL, err := url.Parse(secretURLRaw)
	require.NoError(t, err)

	values := secretURL.Query()

	secret := values.Get("secret")

	algorithm := otp.AlgorithmSHA1

	switch strings.ToUpper(values.Get("algorithm")) {
	case "SHA1":
		algorithm = otp.AlgorithmSHA1
	case "SHA256":
		algorithm = otp.AlgorithmSHA256
	case "SHA512":
		algorithm = otp.AlgorithmSHA512
	}

	period, err := strconv.ParseUint(values.Get("period"), 10, 32)
	require.NoError(t, err)

	digits, err := strconv.ParseInt(values.Get("digits"), 10, 32)
	require.NoError(t, err)

	require.NoError(t, rs.WaitElementLocatedByID(t, page, "dialog-next").Click("left", 1))

	passcode, err := totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
		Period:    uint(period),
		Skew:      1,
		Digits:    otp.Digits(digits),
		Algorithm: algorithm,
	})

	require.NoError(t, err)

	rs.doEnterOTP(t, page, passcode)

	require.NoError(t, page.WaitStable(time.Millisecond*100))

	rs.doSettingsMenuClose(t, page)

	return secret
}

func (rs *RodSession) doEnterOTP(t *testing.T, page *rod.Page, code string) {
	inputs := rs.WaitElementsLocatedByID(t, page, "otp-input input")

	for i := 0; i < len(code); i++ {
		err := inputs[i].Type(input.Key(code[i]))
		require.NoError(t, err)
	}
}

func (rs *RodSession) doValidateTOTP(t *testing.T, page *rod.Page, secret string) {
	code, err := totp.GenerateCode(secret, time.Now())
	assert.NoError(t, err)
	rs.doEnterOTP(t, page, code)
}
