package fields

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

type urlSetToSpellDb string

func (s urlSetToSpellDb) Id() exp.IdentifierExpression {
	return s.T().Col("id")
}

func (s urlSetToSpellDb) UrlSetId() exp.IdentifierExpression {
	return s.T().Col("url_set_id")
}

func (s urlSetToSpellDb) SpellId() exp.IdentifierExpression {
	return s.T().Col("spell_id")
}

func (s urlSetToSpellDb) T() exp.IdentifierExpression {
	tableName := string(s)
	return goqu.T(tableName)
}

func (s urlSetToSpellDb) Aliased(tableName string) urlSetToSpellDb {
	return urlSetToSpellDb(tableName)
}

func UrlSetToSpell() urlSetToSpellDb {
	var s urlSetToSpellDb = "url_sets_to_spells"
	return s
}
