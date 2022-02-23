package main

type IObservable interface {
	subscribe(observer ...IObserver)
	unsubscribe(observer ...IObserver)
	notifyAll()
}

func createNewsObservable(title string) IObservable {
	return news{
		users: make(map[string]IObserver),
		title: title,
	}
}

type news struct {
	title string
	users map[string]IObserver
}

func (n news) subscribe(users ...IObserver) {
	for _, user := range users {
		n.users[user.getID()] = user
	}
}

func (n news) unsubscribe(users ...IObserver) {
	for _, user := range users {
		delete(n.users, user.getID())
	}
}

func (n news) notifyAll() {
	for _, user := range n.users {
		user.update(n.title)
	}
}
