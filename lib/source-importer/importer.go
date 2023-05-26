package source_importer

type SpellCondition int32
type SpellFeature int32

const (
	Default SpellCondition = iota
	No
	MinimalLevel
	Custom
)

const (
	NAtDay SpellFeature = iota
	NotWasteSpellSlot
)
