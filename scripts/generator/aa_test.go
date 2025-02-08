package generator

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/thj-wiki/scripts/database/dbspell"
	"github.com/xackery/thj-wiki/scripts/dbstr"
	"github.com/xackery/thj-wiki/scripts/spdat"
)

var lastType int

type ability struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Type        int    `db:"type"`
	Classes     int    `db:"classes"`
	FirstRankID int    `db:"first_rank_id"`
	Enabled     bool   `db:"enabled"`
}

type rank struct {
	ID         int `db:"id"`
	Cost       int `db:"cost"`
	LevelReq   int `db:"level_req"`
	NextID     int `db:"next_id"`
	SpellID    int `db:"spell"`
	RecastTime int `db:"recast_time"`
	TitleID    int `db:"title_sid"`
	DescID     int `db:"desc_sid"`
}

type rankEffect struct {
	EffectID int `db:"effect_id"`
	Base1    int `db:"base1"`
	Base2    int `db:"base2"`
}

func TestAAGenerate(t *testing.T) {
	err := runAA()
	if err != nil {
		fmt.Println("Failed:", err)
		os.Exit(1)
	}
}

func runAA() error {
	var err error
	var db *sqlx.DB

	err = dbstr.Load()
	if err != nil {
		return fmt.Errorf("dbstr.Load: %w", err)
	}
	eqDB := os.Getenv("EQ_DB")
	if eqDB == "" {
		return fmt.Errorf("EQ_DB not set")
	}

	db, err = sqlx.Open("mysql", eqDB)
	if err != nil {
		return fmt.Errorf("sql.Open: %w", err)
	}

	err = dbspell.Load(db)
	if err != nil {
		return fmt.Errorf("dbspell.Load: %w", err)
	}

	w, err = os.Create("../../content/classes-and-abilities/aa/_index.en.md")
	if err != nil {
		return fmt.Errorf("os.Create: %w", err)
	}
	defer w.Close()

	w.WriteString(fmt.Sprintf(`---
title: AA
weight: 5
chapter: true
description: AA Breakdown
images: [images/aa.png]
aliases: [/aa]
---

![AAs](images/aa.png)

This page was last updated on %s.

Found an error on this page? Check tracked issues and [report new ones in the discord](<https://discord.com/channels/1204418766318862356/1307325765636980736/1307325765636980736>)

## Filter

Select which classes to include on AA list
{{<rawhtml>}}
<div class="filter-container">

        <label><input type="checkbox" value="ALL" class="filter-checkbox" checked> ALL</label><br>
		<label><input type="checkbox" value="BER" class="filter-checkbox"> BER</label>
		<label><input type="checkbox" value="BRD" class="filter-checkbox"> BRD</label>
		<label><input type="checkbox" value="BST" class="filter-checkbox"> BST</label>
		<label><input type="checkbox" value="CLR" class="filter-checkbox"> CLR</label>
		<label><input type="checkbox" value="DRU" class="filter-checkbox"> DRU</label>
        <label><input type="checkbox" value="ENC" class="filter-checkbox"> ENC</label>
        <label><input type="checkbox" value="MAG" class="filter-checkbox"> MAG</label>
		<label><input type="checkbox" value="MNK" class="filter-checkbox"> MNK</label><br>
		<label><input type="checkbox" value="NEC" class="filter-checkbox"> NEC</label>
		<label><input type="checkbox" value="PAL" class="filter-checkbox"> PAL</label>
        <label><input type="checkbox" value="ROG" class="filter-checkbox"> ROG</label>
		<label><input type="checkbox" value="RNG" class="filter-checkbox"> RNG</label>
		<label><input type="checkbox" value="SHD" class="filter-checkbox"> SHD</label>
        <label><input type="checkbox" value="SHM" class="filter-checkbox"> SHM</label>
        <label><input type="checkbox" value="WAR" class="filter-checkbox"> WAR</label>
		<label><input type="checkbox" value="WIZ" class="filter-checkbox"> WIZ</label>
    </div>
{{</rawhtml>}}
`, time.Now().Format("2006-01-02")))

	abilities, err := abilitiesByName(db, "all")
	if err != nil {
		return fmt.Errorf("abilitiesByName: %w", err)
	}
	// select * from aa_ability ab where ab.first_rank_id in ( select rank_id from aa_rank_effects are where are.effect_id = 375)

	fmt.Printf("%d total abilities found\n", len(abilities))
	for _, ability := range abilities {
		fmt.Println(ability.Name)
		if ability.Type != lastType {
			switch ability.Type {
			case 1:
				printWriterf("## General\n")
			case 2:
				printWriterf("## Archetype\n")
			case 3:
				printWriterf("## Class\n")
			default:
				printWriterf("## Unknown\n")
			}
			lastType = ability.Type
		}
		err = dumpRankData(&ability, ability.ID, db)
		if err != nil {

			fmt.Println("Error dumping ability:", ability.ID, err)
			//			return fmt.Errorf("dumpRankData: %w", err)
		}
	}

	printWriterf(`
{{<rawhtml>}}
<script src="aa.js"></script>
{{</rawhtml>}}
`)

	return nil
}

func analyzeTitle(line string) error {
	records := strings.Split(line, "^")
	if len(records) != 4 {
		return fmt.Errorf("invalid line: %s", line)
	}
	if records[1] != "1" {
		return nil
	}
	title := records[2]
	if strings.Contains(title, "<br>") {
		title = strings.Split(title, "<br>")[0]
	}
	id, err := strconv.Atoi(records[0])
	if err != nil {
		return fmt.Errorf("strconv.Atoi: %w", err)
	}
	titles[id] = title
	return nil
}

func dumpRankData(ability *ability, abilityID int, db *sqlx.DB) error {

	if strings.Contains(ability.Name, "Fundament") {
		return nil
	}
	if strings.Contains(ability.Name, "Unknown AA") {
		return nil
	}
	classes := ClassesFromBitmask(int32(ability.Classes))
	if classes.String() == "" {
		return nil
	}

	ranks, effects, err := ranksAndEffectsByAbilityID(db, ability.FirstRankID)
	if err != nil {
		return fmt.Errorf("ranksAndEffectsByAbilityID: %w", err)
	}

	minRecastTime := -1
	maxRecastTime := -1
	totalCost := 0

	totalRanks := 0
	title := ""
	description := ""
	for _, rank := range ranks {
		if rank.LevelReq > 51 {
			continue
		}
		totalRanks++
		if rank.RecastTime > 0 {
			if minRecastTime == -1 || rank.RecastTime < minRecastTime {
				minRecastTime = rank.RecastTime
			}
			if rank.RecastTime > maxRecastTime {
				maxRecastTime = rank.RecastTime
			}
		}
		totalCost += rank.Cost
		if title == "" {
			title = titles[rank.TitleID]
		}
		if description == "" {
			description = dbstr.AADescription(rank.DescID)
		}
	}
	if totalRanks == 0 {
		return nil
	}
	if description == "" {
		description = "No description available"
	}

	if title == "" {
		title = ability.Name
		//		return fmt.Errorf("no title for ability %d %s", ability.FirstRankID, ability.Name)
	}

	/* if totalCost == 0 {
		return nil
	}
	*/
	printWriterf(`{{<details title="%s (%s)">}}%s`, strings.TrimSpace(title), classes, "\n")
	printWriterf("%s\n\n", description)
	printWriterf("Ability ID: %d has %d ranks and costs %d total", abilityID, len(ranks), totalCost)
	if minRecastTime > 0 || maxRecastTime > 0 {
		if minRecastTime == maxRecastTime {
			printWriterf(" with a recast time of %d seconds", minRecastTime)
		} else {
			printWriterf(" with a recast time of %d to %d seconds", minRecastTime, maxRecastTime)
		}
	}
	printWriterf("\n")
	lastBase := 0
	lastBase2 := 0
	for i, rank := range ranks {
		aaPlural := "AAs"
		if rank.Cost == 1 {
			aaPlural = "AA"
		}
		//printWriterf("Rank %d (%d) costs %d %s (%d total spent) with effect:", i+1, rank.ID, rank.Cost, aaPlural, totalAA)
		printWriterf("- Rank %d (ID %d) costs %d %s with effect:", i+1, rank.ID, rank.Cost, aaPlural)

		if rank.SpellID > 0 {
			dumpSpell(db, rank.SpellID)
			continue
		}

		if len(effects[i]) == 1 {

			effect := effects[i][0]
			baseDelta := effect.Base1 - lastBase
			baseDelta2 := effect.Base2 - lastBase2
			printWriterf(" %s", spdat.EffectName(effect.EffectID))
			if effect.Base1 != 0 {
				if effect.EffectID == 85 {
					printWriterf(" [Spell %d](https://www.thjdi.cc/spell/%d)", effect.Base1, effect.Base1)
				} else {
					printWriterf(" %d", effect.Base1)
				}
			}
			if effect.Base2 != 0 {
				if effect.Base1 == 0 {
					printWriterf(" 0")
				}
				printWriterf(" %d", effect.Base2)
			}

			ratio := float64(baseDelta) / float64(rank.Cost)
			if ratio > 0.1 {
				printWriterf(" (%0.2f ratio)\n", float64(baseDelta)/float64(rank.Cost))
				lastBase = effect.Base1
				lastBase2 = effect.Base2
				continue
			}
			ratio = float64(baseDelta2) / float64(rank.Cost)
			if ratio > 0.1 {
				printWriterf(" (%0.2f ratio)\n", float64(baseDelta2)/float64(rank.Cost))
				lastBase = effect.Base1
				lastBase2 = effect.Base2
				continue
			}
			printWriterf("\n")
			continue
		}
		for j, effect := range effects[i] {
			baseDelta := effect.Base1 - lastBase
			baseDelta2 := effect.Base2 - lastBase2

			printWriterf("\n    - Effect %d: %s", j+1, spdat.EffectName(effect.EffectID))

			if effect.Base1 != 0 {
				if effect.EffectID == 85 {
					printWriterf(" [Spell %d](https://www.thjdi.cc/spell/%d)\n", effect.Base1, effect.Base1)
					spell, ok := dbspell.SpellByID(int32(effect.Base1))
					if !ok {
						return fmt.Errorf("dbspell.SpellByID %d not found", effect.Base1)
					}
					printWriterf("      - %s\n", spell.Name.String)
					printWriterf(spdat.SpDatInfo(spell, 8))
				} else {
					printWriterf(" %d", effect.Base1)
				}
			}
			if effect.Base2 != 0 {
				printWriterf(" %d", effect.Base1)
			}

			ratio := float64(baseDelta) / float64(rank.Cost)
			if ratio > 0.1 {
				printWriterf(" (%0.2f ratio)\n", float64(baseDelta)/float64(rank.Cost))
				lastBase = effect.Base1
				lastBase2 = effect.Base2
				continue
			}
			ratio = float64(baseDelta2) / float64(rank.Cost)
			if ratio > 0.1 {
				printWriterf(" (%0.2f ratio)\n", float64(baseDelta2)/float64(rank.Cost))
				lastBase = effect.Base1
				lastBase2 = effect.Base2
				continue
			}
			printWriterf("\n")
			continue

		}
		printWriterf("\n")
	}

	printWriterf("{{</details>}}\n")
	return nil
}

func abilitiesByName(db *sqlx.DB, name string) ([]ability, error) {
	abilities := []ability{}

	if name == "all" {
		for i := 1; i < 4; i++ {
			tmpAbilities := []ability{}
			err := db.Select(&tmpAbilities, "SELECT id, name, type, classes, first_rank_id, enabled FROM aa_ability WHERE type = ? ORDER BY name ASC", i)
			if err != nil {
				return nil, fmt.Errorf(": %w", err)
			}
			abilities = append(abilities, tmpAbilities...)
		}
		return abilities, nil
	}

	err := db.Select(&abilities, "SELECT id, name, type, classes, first_rank_id, enabled FROM aa_ability WHERE name like ? ORDER BY name ASC", name)
	if err != nil {
		return nil, fmt.Errorf(": %w", err)
	}
	return abilities, nil
}

func ranksAndEffectsByAbilityID(db *sqlx.DB, id int) ([]rank, [][]rankEffect, error) {

	var ranks []rank
	var effects [][]rankEffect
	rank, err := abilityRankById(db, id)
	if err != nil {
		if errors.Unwrap(err) != sql.ErrNoRows {
			return nil, nil, fmt.Errorf("abilityRankById %d: %w", id, err)
		}
		return nil, nil, fmt.Errorf("abilityRankById %d: %w", id, err)
	}

	ranks = append(ranks, *rank)

	rankEffects, err := rankEffectsByRankID(db, id)
	if err != nil {
		return nil, nil, fmt.Errorf("rankEffectsByRankID: %w", err)
	}
	effects = append(effects, rankEffects)

	if rank.NextID == -1 {
		return ranks, effects, nil
	}

	for {
		id = rank.NextID
		rank, err = abilityRankById(db, id)
		if err != nil {
			if errors.Unwrap(err) == sql.ErrNoRows {
				break
			}
			return nil, nil, fmt.Errorf("abilityRankById %d: %w", id, err)
		}
		if rank.LevelReq > 51 {
			break
		}

		ranks = append(ranks, *rank)

		rankEffects, err := rankEffectsByRankID(db, id)
		if err != nil {
			return nil, nil, fmt.Errorf("rankEffectsByRankID: %w", err)
		}
		effects = append(effects, rankEffects)
		if rank.NextID == -1 {
			break
		}
	}

	return ranks, effects, nil
}

func abilityById(db *sqlx.DB, id int) (*ability, error) {
	a := &ability{}
	err := db.Get(a, "SELECT id, name, classes, first_rank_id, type, enabled, recast_time FROM aa_ability WHERE id = ?", id)
	if err != nil {
		return nil, fmt.Errorf(": %w", err)
	}
	return a, nil
}

func abilityRankById(db *sqlx.DB, id int) (*rank, error) {
	r := &rank{}
	err := db.Get(r, "SELECT id, cost, level_req, next_id, spell, recast_time, title_sid, desc_sid FROM aa_ranks WHERE id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("db.Get %d: %w", id, err)
	}
	return r, nil
}

func rankEffectsByRankID(db *sqlx.DB, id int) ([]rankEffect, error) {
	var effects []rankEffect
	err := db.Select(&effects, "SELECT effect_id, base1, base2 FROM aa_rank_effects WHERE rank_id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("db.Select: %w", err)
	}
	return effects, nil
}

var descriptions = make(map[int]string)
var titles = make(map[int]string)
