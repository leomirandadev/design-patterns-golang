package main

func newMongoAdapter() RepoAdapter {
	return mongoAdapter{mongoImpl: newMongo()}
}

type mongoAdapter struct {
	mongoImpl mongoImplement
}

func (m mongoAdapter) insert() {
	m.mongoImpl.insertOne()
}
