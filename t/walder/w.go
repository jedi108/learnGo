package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"sync/atomic"
)

const (
	worksLimit = 100
)

type wokers struct {
	wg           sync.WaitGroup
	quotaChannel chan struct{}
	numGoRoutine int
	worldCount   int32
}

func NewWorkers(countWorkers int) *wokers {
	wrk := &wokers{}
	wrk.quotaChannel = make(chan struct{}, countWorkers)
	return wrk
}

func (wrk *wokers) Worker(url string) {
	wrk.wg.Add(1)
	go func() {
		defer wrk.wg.Done()
		wrk.quotaChannel <- struct{}{}
		intCounter := int32(СountWorlds(GetUrlContent(url), "Go"))
		atomic.AddInt32(&wrk.worldCount, intCounter)
		fmt.Printf("Count for %s:%v\n", url, intCounter)
		<-wrk.quotaChannel
	}()
}

func (wrk *wokers) Wait() {
	wrk.wg.Wait()
}

func GetUrlContent(url string) string {
	response, err := http.Get(url)
	if err != nil {
		return fmt.Sprintln("error %s", err)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Sprintln("error %s", err)
	}
	return string(contents)
}

func СountWorlds(fromString, findString string) int {
	return strings.Count(string(fromString), findString)
}

func main() {
	wrs := NewWorkers(worksLimit)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		wrs.Worker(scanner.Text())
	}
	wrs.Wait()
	close(wrs.quotaChannel)
	fmt.Printf("Count: %v\n", wrs.worldCount)
}
