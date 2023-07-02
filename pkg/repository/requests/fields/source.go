package fields

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

type sourceDb string

func (s sourceDb) Id() exp.IdentifierExpression {
	return s.T().Col("id")
}

func (s sourceDb) Name() exp.IdentifierExpression {
	return s.T().Col("name")
}

func (s sourceDb) Description() exp.IdentifierExpression {
	return s.T().Col("description")
}

func (s sourceDb) IsOfficial() exp.IdentifierExpression {
	return s.T().Col("is_official")
}

func (s sourceDb) Author() exp.IdentifierExpression {
	return s.T().Col("author")
}

func (s sourceDb) T() exp.IdentifierExpression {
	tableName := string(s)
	return goqu.T(tableName)
}

func (s sourceDb) Aliased(tableName string) sourceDb {
	return sourceDb(tableName)
}

func Source() sourceDb {
	var s sourceDb = "sources"
	return s
}
