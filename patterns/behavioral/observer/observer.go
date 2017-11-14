package main

type Observer interface {
	Update(state string)
}
type Publisher interface {
	Attach(observer Observer)
	Notify()
}

type NewsPublisher struct {
	observers []Observer
	State     string
}

func (newsPublisher *NewsPublisher) Attach(observer Observer) {
	newsPublisher.observers = append(newsPublisher.observers, observer)
}

func (newsPublisher *NewsPublisher) Notify() {
	for _, observer := range newsPublisher.observers {
		observer.Update(newsPublisher.State)
	}
}

type PostObserver struct {
	state string
}

func (postObserver *PostObserver) Update(state string) {
	postObserver.state = state
}

func main() {
	publisher := &NewsPublisher{}
	publisher.Attach(&PostObserver{})
	publisher.Attach(&PostObserver{})
	publisher.State = "news finance"
	publisher.Notify()
}
