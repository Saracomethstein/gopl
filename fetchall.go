package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

/*
Fetchall выполняет параллельную выбрку URL и сообщает
о затраченном времени и размере ответа для каждого из них.
*/
func fetchall() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetchUp(url, ch)
	}

	var info string
	dt := time.Now()
	info = fmt.Sprintln("Report from: ", dt.Format(time.RFC822))
	reportWriter(info)

	for range os.Args[1:] {
		reportWriter(<-ch)
	}
	info = fmt.Sprintf("%.2fs elapsed\n\n", time.Since(start).Seconds())
	reportWriter(info)
}

func fetchUp(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs\t%7d\t%s\n", secs, nbytes, url)
}

func reportWriter(str string) (l int) {
	file, err := os.OpenFile("report.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		fmt.Println(err)
		return
	}

	l, err = file.WriteString(str)

	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}

	err = file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	return l
}
