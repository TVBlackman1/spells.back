package requests

import (
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"spells.tvblackman1.ru/lib/tribool"
	"spells.tvblackman1.ru/pkg/domain/dto"
	"spells.tvblackman1.ru/pkg/repository/requests/fields"
)

func SelectSpellsWithSourceName(searchParams dto.SearchSpellDto) *goqu.SelectDataset {
	sel := []any{
		fields.Spell().Id(),
		fields.Spell().Name().As("spells_name"),
		fields.Spell().Level(),
		fields.Spell().Description(),
		fields.Spell().CastingTime(),
		fields.Spell().Duration(),
		fields.Spell().IsVerbal(),
		fields.Spell().IsSomatic(),
		fields.Spell().ISMaterial(),
		fields.Spell().MaterialContent(),
		fields.Spell().MagicalSchool(),
		fields.Spell().Distance(),
		fields.Spell().IsRitual(),
		fields.Spell().SourceId(),
		fields.Source().Name().As("sources_name"),
	}

	request := goqu.Dialect("postgres").From(fields.Spell().T()).
		LeftJoin(fields.Source().T(), goqu.On(fields.Source().Id().Eq(fields.Spell().SourceId()))).
		Select(sel...)

	if len(searchParams.Sources) > 0 {
		sources := make([]string, len(searchParams.Sources), len(searchParams.Sources))
		for i := range sources {
			sources[i] = uuid.UUID(searchParams.Sources[i]).String()
		}
		request = request.Where(fields.Source().Id().In(sources))
	}
	// TODO Parse to all uuid invokes
	if uuid.UUID(searchParams.Id) != uuid.Nil {
		stringId := uuid.UUID(searchParams.Id).String()
		request = request.Where(fields.Spell().Id().Eq(stringId))
	} else if len(searchParams.EqualsName) > 0 {
		request = request.Where(fields.Spell().Name().Eq(searchParams.EqualsName))
	} else if len(searchParams.LikeName) > 0 {
		like := fmt.Sprintf("%%%s%%", searchParams.LikeName)
		request = request.Where(fields.Spell().Name().ILike(like))
	}
	if searchParams.IsVerbal != tribool.Unset {
		request = request.Where(fields.Spell().IsVerbal().Eq(searchParams.IsVerbal == tribool.True))
	}
	if searchParams.IsSomatic != tribool.Unset {
		request = request.Where(fields.Spell().IsSomatic().Eq(searchParams.IsSomatic == tribool.True))
	}
	if searchParams.HasMaterialComponent != tribool.Unset {
		request = request.Where(fields.Spell().ISMaterial().Eq(searchParams.HasMaterialComponent == tribool.True))
	}
	if searchParams.IsRitual != tribool.Unset {
		request = request.Where(fields.Spell().IsRitual().Eq(searchParams.IsRitual == tribool.True))
	}
	if len(searchParams.Levels) > 0 {
		request = request.Where(fields.Spell().Level().In(searchParams.Levels))
	}
	if len(searchParams.MagicalSchools) > 0 {
		request = request.Where(fields.Spell().MagicalSchool().In(searchParams.MagicalSchools))
	}
	return request
}

func CountRows(request *goqu.SelectDataset) *goqu.SelectDataset {
	return request.Select(goqu.L("COUNT(*)"))
}
