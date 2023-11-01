package transport

// #cgo CFLAGS: -Iinternal/transport
// #cgo LDFLAGS: internal/transport/cow_say.a
// #include <cow_say.h>
import "C"

import (
	"encoding/json"
	"log"
	"net/http"
	"unsafe"
)

type buyCandyResponse struct {
	Change  int64  `json:"change"`
	Message string `json:"thanks"`
}

type buyCandyErrorResponse struct {
	Error string `json:"error"`
}

const thank = "Thank you!"

func newResponseSuccess(w http.ResponseWriter, r *http.Request, statusCode int, change int64) {
	log.Printf("[%s] %s - Response: SOLD", r.Method, r.URL.Path)

	mes := C.CString(thank)
	defer C.free(unsafe.Pointer(mes))

	cowSay := C.ask_cow(mes)
	defer C.free(unsafe.Pointer(cowSay))

	resp, _ := json.Marshal(buyCandyResponse{
		Change:  change,
		Message: C.GoString(cowSay),
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resp)
}

func newResponseError(w http.ResponseWriter, r *http.Request, statusCode int, err string) {
	log.Printf("[%s] %s - Response: Error: %s", r.Method, r.URL.Path, err)

	resp, _ := json.Marshal(buyCandyErrorResponse{
		Error: err,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resp)
}
