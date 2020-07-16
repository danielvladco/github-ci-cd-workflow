package httpserver

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"testing"
)

func Test_HttpHandler(t *testing.T)  {

	l, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()

	go http.Serve(l, HttpHandler())
	t.Run("Test hello response", func(t *testing.T) {
		res, err := http.DefaultClient.Get(fmt.Sprintf("http://%s/hello", l.Addr()))
		if err != nil {
			t.Fatal(err)
		}
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		defer res.Body.Close()

		if string(b) != `{"message": "helloworld""}` {
			t.Errorf("extepcted: %q; received: %q", `{"message": "helloworld""}`, string(b))
		}
	})
}
