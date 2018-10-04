package goclienttest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	gorilla "github.com/gorilla/http"
	"github.com/levigross/grequests"
	"github.com/parnurzeal/gorequest"
	"github.com/valyala/fasthttp"
)

func RunClientTest(totalcount int, concurrent int, url string) {
	loopcount := totalcount / concurrent
	callchan := make(chan int)
	start := time.Now()
	for x := 0; x < concurrent; x++ {
		go CallRoutine(loopcount, url, callchan)
	}

	var totalerrors int
	for y := 0; y < concurrent; y++ {
		totalerrors = totalerrors + <-callchan
	}

	fmt.Printf("Total elapsed time: %s\n", time.Since(start))
	fmt.Printf("Total errors: %d\n", totalerrors)
}

func CallRoutine(count int, url string, callchan chan int) {
	errors := 0
	for x := 0; x < count; x++ {
		if !SendGRequestsCall(url) {
			errors++
		}
	}
	callchan <- errors
}

func SendCall(url string) bool {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil || resp.StatusCode != 200 {
		fmt.Println(err)
		return false
	}

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}

	//fmt.Println(string(b))
	return true
}

func SendFastHttpCall(url string) bool {
	bytes := make([]byte, 800)
	code, _, error := fasthttp.Get(bytes, url)
	if code != 200 {
		fmt.Println(error)
		return false
	}
	return true
}

func SendGorillaCall(url string) bool {
	gorilla.Get(ioutil.Discard, url)
	return true
}

func SendGoRequestCall(url string) bool {
	request := gorequest.New()
	resp, _, errs := request.Get(url).End()
	if resp.StatusCode != 200 {
		return false
	}
	if errs != nil {
		return false
	}
	ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return true
}

func SendGRequestsCall(url string) bool {
	resp, err := grequests.Get(url, nil)
	if resp.StatusCode != 200 {
		return false
	}
	if err != nil {
		return false
	}
	resp.String()
	return true
}
