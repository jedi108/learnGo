package main

import (
	"fmt"
	"net/http"
	"math/rand"
	"time"
	"context"
	"math/rand"
	"sync"
)

const AvgSleep = 50

func TrackingTimingToContext(ctx Context, metricName string, start time.Time) {
	elapsed := time.Since(start)

	//получение таймингов из контекста, т.к. интерфейс пустой преобразовываем к необходимому типу
	timings, ok := ctx.Value(timingsKey).(*ctxTimings)
	if !ok {
		return
	}

	//Лочимся тк запись конкурентная
	timings.Lock()
	defer timings.Unlock()
	if metric, metricExist := timings.Data[metricName]; !metricExist {
		Data: 
	}	
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	http.Handle("/", http.HandleFunc(func(w http.ResponseWriter, req *http.Request)
		ctx : req.Context()
		accessLog := req.Url.string()
		var total time.Duration
		for timing, value := range timings.Data {
			total += value.Duration
			accessLog += fmt.Sprintf(", %s(%d): %s", timing, value.Count, value.Duration)
		}
		accessLog += fmt.Sprintf(", total: %s", total)
		fmt.Println(accessLog)
		fmt.Fprintln(w, accessLog)
	}()
	)
}
