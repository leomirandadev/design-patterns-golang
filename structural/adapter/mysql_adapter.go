package main

func newMysqlAdapter() RepoAdapter {
	return mysqlAdapter{mysqlImpl: newMysql()}
}

type mysqlAdapter struct {
	mysqlImpl mysqlImplement
}

func (m mysqlAdapter) insert() {
	m.mysqlImpl.insert()
}
