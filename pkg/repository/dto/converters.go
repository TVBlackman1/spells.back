package dbdto

import (
	"spells.tvblackman1.ru/pkg/domain/dto"
)

func DbSpellToSpellDto(spellDb SpellDb) dto.SpellDto {
	res := dto.SpellDto{
		Id:                   dto.SpellId(spellDb.Id),
		Name:                 spellDb.Name,
		Level:                spellDb.Level,
		Description:          spellDb.Description,
		CastingTime:          spellDb.CastingTime,
		Duration:             spellDb.Duration,
		IsVerbal:             spellDb.IsVerbal,
		IsSomatic:            spellDb.IsSomatic,
		HasMaterialComponent: spellDb.HasMaterial,
		MagicalSchool:        spellDb.MagicalSchool,
		Distance:             spellDb.Distance,
		IsRitual:             spellDb.IsRitual,
		SourceId:             dto.SourceId(spellDb.SourceId),
		SourceName:           spellDb.SourceName,
	}
	if spellDb.MaterialContent.Valid {
		res.MaterialComponent = spellDb.MaterialContent.String
	}
	return res
}

func DbSpellMarkedToSpellMarkedDto(spellMarkedDb SpellMarkedDb) dto.SpellMarkedDto {
	common := spellMarkedDb.SpellDb
	spellDtoPart := DbSpellToSpellDto(common)
	return dto.SpellMarkedDto{
		SpellDto: spellDtoPart,
		InSet:    spellMarkedDb.InSet,
	}
}
