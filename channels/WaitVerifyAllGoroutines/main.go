/*

Проверка прекращения работы worker’ов

В этом примере горутина запускается, ждет передачи данных от канала die (или его закрытия).
В случае, когда придет сигнал, горутина выполнит завершающие действия и пошлет сигнал в функцию main
(через этот же канал die) о том, что она завершена.

*/

package main

import "time"

//import "fmt"

func worker(die chan bool) {
	for {
		select {
		// ... выполняем что-нибудь в других case
		//case ch := <-die:
		//	fmt.Println("gorun", ch)
		case <-die:
			// ... выполняем необходимые действия перед завершением.
			die <- true
			return
		}
	}
}

func main() {
	die := make(chan bool)
	go worker(die)
	die <- true

	// Ждем, пока все горутины закончат выполняться
	<-die

	time.Sleep(time.Second * 2)
}
