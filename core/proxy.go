package core
import(
	"fmt"
	"net/http"
	"io"
)

func Proxy(w http.ResponseWriter, r *http.Request) {

    resp, err := http.DefaultTransport.RoundTrip(r)
    if err != nil {
        http.Error(w, err.Error(), 502)
        return
    }
    defer resp.Body.Close()

    for k, vv := range resp.Header {
        for _, v := range vv {
            w.Header().Add(k, v)
        }
    }

    w.WriteHeader(resp.StatusCode)
    io.Copy(w, resp.Body)

	fmt.Printf("[%s] %s\n", r.Method, r.URL)

}
