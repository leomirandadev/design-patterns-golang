package main

func main() {
	user1 := newUser("123", "Leo")
	user2 := newUser("323", "Nicole")

	newsObservable := createNewsObservable("Noticia 1")
	newsObservable.subscribe(user1, user2)

	newsObservable.notifyAll()
}
