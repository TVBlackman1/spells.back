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
	if len(searchParams.EqualsName) > 0 {
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

//func main() {
//	allSpellsRequest := selectSpellsWithSourceName(dto.SearchSpellDto{})
//	//sqlRequest, _, _ := allSpellsRequest.ToSQL()
//
//	usedSpellsRequest := goqu.Dialect("postgres").
//		From(fields.UrlSetToSpell().T()).
//		Where(fields.UrlSetToSpell().UrlSetId().Eq("02677b6a-0e34-455c-a5da-1316e681eb44")).
//		LeftJoin(fields.Spell().T(), goqu.On(fields.UrlSetToSpell().SpellId().Eq(fields.Spell().Id())))
//	//sqlRequest2, _, _ := usedSpellsRequest.ToSQL()
//	//fmt.Println(sqlRequest2)
//	allSpellsWithMark := allSpellsRequest.LeftJoin(usedSpellsRequest.As("used"), goqu.On(
//		fields.UrlSetToSpell().Aliased("used").SpellId().
//			Eq(
//				fields.Spell().Id())))
//	allSpellsWithMark = allSpellsWithMark.SelectAppend(
//		fields.UrlSetToSpell().Aliased("used").UrlSetId(),
//	)
//	readySql, _, _ := allSpellsWithMark.ToSQL()
//	fmt.Println(readySql)
//
//	//allSpellsRequest.LeftJoin()
//	//Where(fields.UrlSetToSpell().)
//	//request.LeftJoin()
//
//	//fmt.Println(sqlRequest)
//}
