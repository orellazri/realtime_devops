package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"time"
)

type Connection struct {
	srv *httptest.Server
}

func NewConnection() *Connection {
	srv := httptest.NewServer(http.HandlerFunc(handlePost))
	return &Connection{srv}
}

func (conn *Connection) Close() {
	conn.srv.Close()
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
		return
	}

	w.Write([]byte(r.FormValue("message")))
}

func (conn *Connection) BenchmarkWrite(numIterations int) (time.Duration, error) {
	var totalTime time.Duration
	for i := 0; i < numIterations; i++ {
		start := time.Now()
		res, err := http.PostForm(conn.srv.URL+"/", url.Values{
			"message": {"testing"},
		})
		if err != nil {
			return time.Duration(0), err
		}
		res.Body.Close()
		end := time.Now()
		totalTime += end.Sub(start)
	}

	return totalTime, nil
}

func Benchmark(numIterations int) (time.Duration, error) {
	srv := httptest.NewServer(http.HandlerFunc(handlePost))
	defer srv.Close()

	expected := "testing"
	data := url.Values{
		"message": {expected},
	}
	var totalTime time.Duration

	for i := 0; i < numIterations; i++ {
		start := time.Now()
		res, err := http.PostForm(srv.URL+"/", data)
		if err != nil {
			return time.Duration(0), err
		}
		defer res.Body.Close()
		end := time.Now()
		totalTime += end.Sub(start)

		out, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return time.Duration(0), err
		}

		actual := string(out)
		if actual != expected {
			return time.Duration(0), fmt.Errorf("wrong response. expected %v got %v", expected, actual)
		}
	}

	return totalTime, nil
}
