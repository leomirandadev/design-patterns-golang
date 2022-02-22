package main

func main() {
	repo := newMongoAdapter()
	repo.insert()

	repo1 := newMysqlAdapter()
	repo1.insert()
}
