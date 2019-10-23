package select_and_reflection

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("get faster url", func(t *testing.T) {

		slowServer := makeDelayedServer(9 * time.Millisecond)
		fastServer := makeDelayedServer(5 * time.Millisecond)

		// 在某个函数调用前加上 defer 前缀会在 包含它的函数结束时 调用它。
		defer func() {
			slowServer.Close()
			fastServer.Close()
		}()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)
		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		serverA := makeDelayedServer(20 * time.Millisecond)

		defer serverA.Close()

		_, err := ConfigurableRacer(serverA.URL, serverA.URL, 10*time.Millisecond)

		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	// httptest.NewServer 接受一个我们传入的 匿名函数 http.HandlerFunc
	// type HandlerFunc func(ResponseWriter, *Request)
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
