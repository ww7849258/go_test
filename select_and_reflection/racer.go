package select_and_reflection

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecoundTimeout = 10 * time.Second

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, tenSecoundTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	// select 则允许你同时在 多个 channel 等待。第一个发送值的 channel「胜出」，case 中的代码会被执行。
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
		// 使用 select 时，time.After 是一个很好用的函数。
		// 当你监听的 channel 永远不会返回一个值时你可以潜在地编写永远阻塞的代码，尽管在我们的案例中它没有发生。
		// time.After 会在你定义的时间过后发送一个信号给 channel 并返回一个 chan 类型（就像 ping 那样）。
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}

func measureTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
