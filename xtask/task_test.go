package xtask_test

import (
	"compress/gzip"
	"context"
	"crypto/md5"
	"fmt"
	"github.com/pubgo/x/xerror"
	"github.com/pubgo/x/xtask"
	"github.com/pubgo/x/xtry"
	"github.com/storyicon/graphquery"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

var _try = xtry.Try

type result struct {
	path string
	sum  [md5.Size]byte
}

func MD5All(ctx context.Context, root string) (map[string][md5.Size]byte, error) {
	g, ctx := xtask.WithContext(ctx)
	paths := make(chan string, 10)

	g.Go(func() {
		defer close(paths)
		xerror.Panic(filepath.Walk(root, func(path string, info os.FileInfo, _err error) (err error) {
			defer xerror.RespErr(&err)

			xerror.Panic(_err)

			if !info.Mode().IsRegular() {
				return
			}

			select {
			case paths <- path:
			case <-ctx.Done():
				return ctx.Err()
			}
			return
		}))
	})

	// Start a fixed number of goroutines to read and digest files.
	c := make(chan result, 10)
	const numDigesters = 20
	for i := 0; i < numDigesters; i++ {
		g.Go(func() {
			for path := range paths {
				data, err := ioutil.ReadFile(path)
				xerror.Panic(err)

				select {
				case c <- result{path, md5.Sum(data)}:
				case <-ctx.Done():
					xerror.Panic(ctx.Err())
				}
			}
		})
	}
	go func() {
		g.Wait()
		close(c)
	}()

	m := make(map[string][md5.Size]byte)
	for r := range c {
		m[r.path] = r.sum
	}

	return m, g.Wait()
}
func TestPipelines(t *testing.T) {
	defer xerror.Debug()

	m, err := MD5All(context.Background(), ".")
	xerror.Panic(err)

	for k, sum := range m {
		fmt.Printf("%s:\t%x\n", k, sum)
	}
}

func xTestJustError(t *testing.T) {
	defer xerror.Assert()

	var g xtask.Group
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}

	for _, url := range urls {
		url := url

		g.Go(func() {
			resp, err := http.Get(url)
			xerror.PanicM(err, "http get error")
			defer resp.Body.Close()
		})
	}
	// Wait for all HTTP fetches to complete.
	xerror.PanicM(g.Wait(), "Successfully fetched all URLs.")
}

func TestTasks(t *testing.T) {
	defer xerror.Assert()

	type DomainInfo struct {
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	}

	//zerolog.SetGlobalLevel(zerolog.WarnLevel)

	xtask.NewAsyncTask(10, time.Second+time.Millisecond*10, func(domain string) {

		url := "https://www.iplocation.net/"
		payload := strings.NewReader("query=" + domain + "&submit=IP%2BLookup")
		req, _ := http.NewRequest("POST", url, payload)
		req.Header.Add("referer", "https://www.iplocation.net/")
		req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36")
		req.Header.Add("origin", "https://www.iplocation.net")
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("cookie", "visid_incap_877543=3e3EmxD2TqqkK2d84rE5yzVQQV0AAAAAQUIPAAAAAADYhpD96RZx4pHCFt2J0B7b; incap_ses_962_877543=jbNdEGayAmnZ0ZTCsbZZDTZQQV0AAAAA/0noF3FJFKN2u3ETSimn3Q==; _ga=GA1.2.1953547087.1564561476; _gid=GA1.2.157467278.1564561476; ci_session=fb1e89fb6aece39e158979d5c9a7863b2148a437, visid_incap_877543=3e3EmxD2TqqkK2d84rE5yzVQQV0AAAAAQUIPAAAAAADYhpD96RZx4pHCFt2J0B7b; incap_ses_962_877543=jbNdEGayAmnZ0ZTCsbZZDTZQQV0AAAAA/0noF3FJFKN2u3ETSimn3Q==; _ga=GA1.2.1953547087.1564561476; _gid=GA1.2.157467278.1564561476; ci_session=fb1e89fb6aece39e158979d5c9a7863b2148a437; visid_incap_877543=EffHUN7oRlC9KiaAV9tOPgFSQV0AAAAAQUIPAAAAAABN0VoQt3L0ENAv4wD57rIO; incap_ses_962_877543=oiSKen5wSX4H25XCsbZZDQFSQV0AAAAAyVccT7DEqURWJ1TtCtFZsw==")
		req.Header.Add("Accept", "*/*")
		req.Header.Add("Cache-Control", "no-cache")
		req.Header.Add("Postman-Token", "509c4555-5327-4068-bc53-6e7312bf409e,0ad8ac07-f642-424f-a990-8c217ed6b5c5")
		req.Header.Add("Host", "www.iplocation.net")
		req.Header.Add("accept-encoding", "gzip, deflate")
		req.Header.Add("content-length", "40")
		req.Header.Add("Connection", "keep-alive")
		req.Header.Add("cache-control", "no-cache")
		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, err := gzip.NewReader(res.Body)
		xerror.Panic(err)
		body1, _ := ioutil.ReadAll(body)
		if res.StatusCode != http.StatusOK {
			panic("状态码不正确")
		}

		expr := "{ latitude `css(\"#wrapper > section > div > div > div.col.col_8_of_12 > div:nth-child(11) > div > table > tbody:nth-child(4) > tr > td:nth-child(3)\")` longitude `css(\"#wrapper > section > div > div > div.col.col_8_of_12 > div:nth-child(11) > div > table > tbody:nth-child(4) > tr > td:nth-child(4)\")` }"
		response := graphquery.ParseFromString(string(body1), expr)

		var di = &DomainInfo{}
		xerror.Panic(response.Decode(di))
		fmt.Println(fmt.Sprintf("%s,%s,%s", domain, di.Latitude, di.Longitude))
		//w.Write([]string{ips[0], ips[1], di.Latitude, di.Longitude})
	})

	handle2 := xtask.NewAsyncTask(10, time.Second+time.Millisecond*10, func(i int) {
		defer xerror.Resp(func(err xerror.IErr) {
			err.P()
		})

		xerror.PanicT(i == 29, "_handle1 90999 error")
	})

	handle1 := xtask.NewAsyncTask(10, time.Second+time.Millisecond*10, func(i int) {
		defer xerror.Resp(func(err xerror.IErr) {
			err.P()
		})

		handle2.Do(i)
		xerror.PanicT(i == 29, "_handle1 90999 error")
	})

	for i := 0; i < 100; i++ {
		handle1.Do(i)
	}

	xtask.WaitAndStop(handle1, handle2)
}

func TestErrLog(t *testing.T) {
	defer xerror.Assert()

	var _task = xtask.NewAsyncTask(500, time.Second+time.Millisecond*10, func(i int) {
		defer xerror.Resp(func(err xerror.IErr) {
			err.P()
		})

		xerror.PanicT(i == 100, "90999 error")
	})

	for i := 0; i < 100000; i++ {
		go _task.Do(i)
	}
	_task.Wait()
}

func parserArticleWithReadability(i int) {
	errChan := make(chan bool)
	go func() {
		time.Sleep(time.Second * 4)
		errChan <- true
	}()

	for {
		select {
		case <-time.After(3 * time.Second):
			xerror.PanicM("readbility timeout", "等待 %d", i)
		case <-errChan:
			return
		}
	}

}

func TestW(t *testing.T) {
	defer xerror.Assert()

	var _task = xtask.NewAsyncTask(10000, time.Second*2, func(i int) {
		xerror.ErrHandle(_try(func() {})(func() {
			parserArticleWithReadability(i)
			fmt.Println("ok", i)
		}), func(err xerror.IErr) {
			xerror.PanicM(err, "testW")
		})
	})
	for i := 0; i < 1000000; i++ {
		_task.Do(i)
	}
	_task.Wait()
}

func TestUrl(t *testing.T) {
	defer xerror.Assert()

	client := &http.Client{Transport: &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    3 * time.Second,
		DisableCompression: true,
	}}
	client.Timeout = 5 * time.Second

	var _task = xtask.NewAsyncTask(200, time.Second*2, func(c *http.Client, i int) {
		xerror.Panic(xtry.Retry(3, func(i int) {
			req, err := http.NewRequest(http.MethodGet, "https://www.yuanben.io", nil)
			xerror.Panic(err)

			resp, err := c.Do(req)
			xerror.Panic(err)
			xerror.PanicT(resp.StatusCode != http.StatusOK, "状态不正确%d", resp.StatusCode)
			fmt.Println("try: ", i, "ok")
		}))
	})
	for i := 0; i < 3000; i++ {
		_task.Do(client, i)
	}
	_task.Wait()
}
