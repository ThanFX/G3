package models

var eventChan = make(chan string, 100)
var events = make([]string, 0)

func GetEvents() []string {
	evArr := make([]string, len(events), cap(events))
	copy(evArr, events)
	events = nil
	return evArr
}

func NewEvent(e string) {
	eventChan <- e
}

func EventLoop() {
	for {
		s := <-eventChan
		events = append(events, s)
	}
}
