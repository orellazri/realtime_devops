package http

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

func TestSendMessage(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(handlePost))
	defer srv.Close()

	expected := "testing"
	data := url.Values{
		"message": {expected},
	}
	const numIterations int = 100
	var totalTime time.Duration

	for i := 0; i < numIterations; i++ {
		start := time.Now()
		res, err := http.PostForm(srv.URL+"/", data)
		if err != nil {
			t.Fatal(err)
		}
		defer res.Body.Close()
		end := time.Now()
		totalTime += end.Sub(start)

		out, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		actual := string(out)
		if actual != expected {
			t.Errorf("wrong response. expected %v got %v", expected, actual)
		}
	}

	if totalTime/time.Duration(numIterations) < time.Duration(1*time.Microsecond) {
		t.Errorf("write average is too long. expected %v got %v", 1*time.Microsecond, totalTime/time.Duration(numIterations))
	}
}
