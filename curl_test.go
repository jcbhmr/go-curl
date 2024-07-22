package curl_test

import (
	"testing"

	"github.com/jcbhmr/go-curl/v8"
)

func TestCurlEasyInit(t *testing.T) {
	h := curl.CurlEasyInit()
	t.Logf("%#+v", h)
}
