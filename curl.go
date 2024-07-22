package curl

// #cgo CFLAGS: -I${SRCDIR}/include -DCURL_STATICLIB
// #cgo linux,amd64 LDFLAGS: -L${SRCDIR}/linux.amd64
// #cgo LDFLAGS: -lcurl -lcares -lpthread  -lnghttp3  -lnghttp2  -lidn2 -lunistring  -lssh2 -lssh2 -lssl -lcrypto -lssl -lcrypto -lssl -ldl -lcrypto -ldl -lz  -lpsl -lssl -lcrypto -lssl -ldl -pthread -lcrypto -ldl -pthread  -lzstd -lzstd  -lbrotlidec -lbrotlidec -lbrotlicommon -lz
// #include <curl/curl.h>
import "C"
import "unsafe"

type Curl unsafe.Pointer

func CurlEasyInit() Curl {
	return Curl(C.curl_easy_init())
}
