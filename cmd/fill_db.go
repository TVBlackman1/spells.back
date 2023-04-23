package main

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
	"log"
	"os"
	"spells.tvblackman1.ru/lib/config"
	"spells.tvblackman1.ru/lib/importer"
	"spells.tvblackman1.ru/lib/postgres"
	"spells.tvblackman1.ru/pkg/domain/dto"
	"spells.tvblackman1.ru/pkg/domain/usecases"
	"spells.tvblackman1.ru/pkg/repository"
	"strings"
)

func main() {
	config, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal(err)
	}
	dbConfig := postgres.PostgresConfig{
		Dbname:   config.POSTGRES_DBNAME,
		Host:     config.POSTGRES_HOST,
		Port:     config.POSTGRES_PORT,
		User:     config.POSTGRES_USER,
		Password: config.POSTGRES_PASS,
	}
	postgresDb, err := postgres.Connect(dbConfig)
	defer postgresDb.Close()

	if err != nil {
		log.Fatal(err)
	}
	repo, err := repository.NewRepository(postgresDb)
	if err != nil {
		log.Fatalf("repo err: %s", err.Error())
	}
	useCases := usecases.NewUseCases(repo)
	file, err := os.Open("./init/seeds/dnd-spells.json")
	if err != nil {
		log.Fatal(err)
	}
	data := importer.GetSpellsData(file)
	fmt.Println(len(data.AllSpells))
	//for _, school := range data.SchoolList {
	//	fmt.Println(school.Text.Ru.Title)
	//}
	for ind, spell := range data.AllSpells {
		ruSpell := spell.Ru
		spellLevel, err := ruSpell.Level.Int64()
		if err != nil {
			fmt.Errorf("not valid spell level: %s of spell %s", ruSpell.Level.String(), ruSpell.Name)
		}
		err = useCases.Spell.CreateSpell(dto.UserId(uuid.New()), dto.CreateSpellDto{
			Name:                 ruSpell.Name,
			Level:                int(spellLevel),
			Classes:              []string{},
			Description:          strings.Replace(ruSpell.Text, "'", "", -1),
			CastingTime:          ruSpell.CastingTime,
			Duration:             ruSpell.Duration,
			IsVerbal:             slices.Contains(strings.Split(ruSpell.Components, ", "), "В"),
			IsSomatic:            slices.Contains(strings.Split(ruSpell.Components, ", "), "С"),
			HasMaterialComponent: slices.Contains(strings.Split(ruSpell.Components, ", "), "М"),
			MaterialComponent:    ruSpell.Materials,
			MagicalSchool:        ruSpell.School,
			Distance:             ruSpell.Range,
			IsRitual:             len(ruSpell.Ritual) > 0,
			SourceId:             dto.SourceId{},
		})
		if err != nil {
			file, _err := os.Create("output.txt")
			fmt.Fprintf(file, "%+v", dto.CreateSpellDto{
				Name:                 ruSpell.Name,
				Level:                int(spellLevel),
				Classes:              []string{"Warrior", "Druid"},
				Description:          strings.Replace(ruSpell.Text, "'", "", -1),
				CastingTime:          ruSpell.CastingTime,
				Duration:             ruSpell.Duration,
				IsVerbal:             slices.Contains(strings.Split(ruSpell.Components, ", "), "В"),
				IsSomatic:            slices.Contains(strings.Split(ruSpell.Components, ", "), "С"),
				HasMaterialComponent: slices.Contains(strings.Split(ruSpell.Components, ", "), "М"),
				MaterialComponent:    ruSpell.Materials,
				MagicalSchool:        ruSpell.School,
				Distance:             ruSpell.Range,
				IsRitual:             len(ruSpell.Ritual) > 0,
				SourceId:             dto.SourceId{},
			})
			if _err != nil {

			}
			log.Fatal(err)
		}
		if ind%20 == 0 {
			fmt.Printf("%d/%d\n", ind, len(data.AllSpells))
		}
	}
	//err = useCases.Spell.CreateSpell(dto.UserId(uuid.New()), dto.CreateSpellDto{
	//	Name:                 "first",
	//	Level:                0,
	//	Classes:              []string{"Warrior", "Druid"},
	//	Description:          "About this",
	//	CastingTime:          "1 action",
	//	Duration:             "30 sec",
	//	IsVerbal:             false,
	//	IsSomatic:            true,
	//	HasMaterialComponent: false,
	//	MaterialComponent:    "asd",
	//	MagicalSchool:        "illusion",
	//	Distance:             "30 feet",
	//	IsRitual:             true,
	//	SourceId:             dto.SourceId{},
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}

}
