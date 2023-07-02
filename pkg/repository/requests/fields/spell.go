package fields

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

type spellDb string

func (s spellDb) Id() exp.IdentifierExpression {
	return s.T().Col("id")
}

func (s spellDb) Name() exp.IdentifierExpression {
	return s.T().Col("name")
}

func (s spellDb) Level() exp.IdentifierExpression {
	return s.T().Col("level")
}

func (s spellDb) Description() exp.IdentifierExpression {
	return s.T().Col("description")
}

func (s spellDb) CastingTime() exp.IdentifierExpression {
	return s.T().Col("casting_time")
}

func (s spellDb) Duration() exp.IdentifierExpression {
	return s.T().Col("duration")
}

func (s spellDb) IsVerbal() exp.IdentifierExpression {
	return s.T().Col("is_verbal")
}

func (s spellDb) IsSomatic() exp.IdentifierExpression {
	return s.T().Col("is_somatic")
}

func (s spellDb) ISMaterial() exp.IdentifierExpression {
	return s.T().Col("is_material")
}

func (s spellDb) MaterialContent() exp.IdentifierExpression {
	return s.T().Col("material_content")
}

func (s spellDb) MagicalSchool() exp.IdentifierExpression {
	return s.T().Col("magical_school")
}

func (s spellDb) Distance() exp.IdentifierExpression {
	return s.T().Col("distance")
}

func (s spellDb) IsRitual() exp.IdentifierExpression {
	return s.T().Col("is_ritual")
}

func (s spellDb) SourceId() exp.IdentifierExpression {
	return s.T().Col("source_id")
}

func (s spellDb) T() exp.IdentifierExpression {
	tableName := string(s)
	return goqu.T(tableName)
}

func (s spellDb) Aliased(tableName string) spellDb {
	return spellDb(tableName)
}

func Spell() spellDb {
	var s spellDb = "spells"
	return s
}
