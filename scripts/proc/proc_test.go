package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/thj-wiki/scripts/dbstr"
	"github.com/xackery/thj-wiki/scripts/itemquest"
	"github.com/xackery/thj-wiki/scripts/spdat"
)

type zone struct {
	ID        int    `db:"id"`
	Short     string `db:"short_name"`
	Long      string `db:"long_name"`
	Expansion int    `db:"expansion"`
}

type item struct {
	ItemID           int    `db:"item_id"`
	IsQuest          bool   `db:"is_quest"`
	Zone             string `db:"zone"`
	NPCName          string `db:"npc_name"`
	ItemName         string `db:"item_name"`
	ProcEffect       int    `db:"proceffect"`
	ProcRate         int    `db:"procrate"`
	EffectID1        int    `db:"effectid1"`
	EffectBaseValue1 int    `db:"effect_base_value1"`
	SpellName        string `db:"spell_name"`
	SpellID          int    `db:"spell_id"`
	Classes          int    `db:"classes"`
}

var w *os.File

func TestProcGenerate(t *testing.T) {
	err := run()
	if err != nil {
		t.Fatalf("Failed: %v", err)
	}
}

func run() error {
	var err error
	var db *sqlx.DB

	w, err = os.Create("../../content/procs/_index.en.md")
	if err != nil {
		return fmt.Errorf("os.Create: %w", err)
	}
	defer w.Close()

	w.WriteString(`---
title: Proc Weapons
weight: 5
chapter: true
description: Quick Reference list of Proc Weapons in The Heroes' Journey
images: [images/proc.png]
---


![Procs](images/proc.png)


Stats are based on regular form. Legendary links are in clicked details.

## Filter

Select which classes to include on Proc Weapon list
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
`)

	printWriterf(`
{{<rawhtml>}}
<script src="procs.js"></script>
{{</rawhtml>}}
`)
	eqDB := os.Getenv("EQ_DB")
	if eqDB == "" {
		return fmt.Errorf("EQ_DB not set")
	}

	db, err = sqlx.Open("mysql", eqDB)
	if err != nil {
		return fmt.Errorf("sql.Open: %w", err)
	}

	for i := 0; i < 8; i++ {
		zones, err := parseZones(db, i)
		if err != nil {
			return fmt.Errorf("parseZones: %w", err)
		}

		switch i {
		case 0:
			w.WriteString("# Classic\n")
			fmt.Printf("Classic Zones: %d\n", len(zones))
		case 1:
			w.WriteString("# Kunark\n")
			fmt.Printf("Kunark Zones: %d\n", len(zones))
		case 2:
			w.WriteString("# Velious\n")
			fmt.Printf("Velious Zones: %d\n", len(zones))
		case 3:
			w.WriteString("# Luclin\n")
			fmt.Printf("Luclin Zones: %d\n", len(zones))
		case 4:
			w.WriteString("# Planes of Power\n")
			fmt.Printf("Planes of Power Zones: %d\n", len(zones))
		case 5:
			w.WriteString("# Gates of Discord\n")
			fmt.Printf("Gates of Discord Zones: %d\n", len(zones))
		case 6:
			w.WriteString("# Omens of War\n")
			fmt.Printf("Omens of War Zones: %d\n", len(zones))
		case 7:
			w.WriteString("# Dragons of Norrath\n")
			fmt.Printf("Dragons of Norrath Zones: %d\n", len(zones))
		}

		fmt.Println("Zones:", len(zones))
		for _, zone := range zones {
			start := time.Now()
			items, err := itemsByZone(db, zone.Short)
			if err != nil {
				return fmt.Errorf("dropsByNPCID: %w", err)
			}

			questItems, err := itemquest.LoadZone(zone.Short)
			if err != nil {
				fmt.Printf("Error loading zone %s: %v\n", zone.Short, err)
			}

			fmt.Printf("Quest Items: %d\n", len(questItems))
			for _, questItem := range questItems {
				item, err := questItemConvert(db, questItem, zone.Short)
				if err != nil {
					if errors.Unwrap(err) == sql.ErrNoRows {
						continue
					}
					return fmt.Errorf("questItemConvert: %w", err)
				}

				items = append(items, *item)
			}

			//if len(items) == 0 {
			//printWriterf("- %s: No items with procs\n", zone.Short)
			//}

			for _, item := range items {
				questStr := ""
				if item.IsQuest {
					questStr = " (Quest)"
				}

				if strings.HasPrefix(item.ItemName, "Summoned: ") {
					continue
				}
				printWriterf(`{{<details title="%s: %s: %s mod %d (%s)">}}`,
					zone.Short,
					item.ItemName, item.SpellName,
					item.ProcRate, ClassesFromBitmask(int32(item.Classes)))
				printWriterf("\n")

				printWriterf("[%s](https://retributioneq.com/allaclone/?a=item&id=%d) is obtained from %s%s in %s and has a combat proc of [%s](https://retributioneq.com/allaclone/?a=spell&id=%d) with a mod of %d.\n- Proc mod and stats of Legendary not displayed on this page. Link: [%s (Legendary)](https://retributioneq.com/allaclone/?a=item&id=%d)\n\n",
					item.ItemName, item.ItemID,
					CleanName(item.NPCName), questStr, zone.Short,
					item.SpellName, item.SpellID, item.ProcRate, item.ItemName, item.ItemID+2000000)

				dumpSpell(db, item.SpellID)
				printWriterf(`{{</details>}}`)
				printWriterf("\n")
			}

			sort.Slice(items, func(i, j int) bool {
				return items[i].ItemName < items[j].ItemName
			})
			fmt.Printf("Zone %s took %0.2f seconds\n", zone.Short, time.Since(start).Seconds())
		}
		return nil
	}

	return nil
}

func parseZones(db *sqlx.DB, expansion int) ([]zone, error) {

	allZones := []zone{}
	zones, err := zonesByExpansionID(db, expansion)
	if err != nil {
		return nil, fmt.Errorf("zonesByExpansionID: %w", err)
	}
	if expansion == 0 {
		zones = append(zones, zone{Short: "cabwest"})
		zones = append(zones, zone{Short: "cabeast"})
		zones = append(zones, zone{Short: "fieldofbone"})
		zones = append(zones, zone{Short: "kurns"})
		zones = append(zones, zone{Short: "warslikswood"})
		zones = append(zones, zone{Short: "swampofnohope"})
		zones = append(zones, zone{Short: "lakeofillomen"})
		zones = append(zones, zone{Short: "sharvahl"})
		zones = append(zones, zone{Short: "paladul"})
		zones = append(zones, zone{Short: "hollowshade"})
		zones = append(zones, zone{Short: "grimling"})
	}
	if expansion == 1 {
		zones = append(zones, zone{Short: "veksar"})
		zones = append(zones, zone{Short: "chardokb"})
		zones = append(zones, zone{Short: "gunthak"})
		zones = append(zones, zone{Short: "hatesfury"})
		zones = append(zones, zone{Short: "nadox"})
	}
	allZones = append(allZones, zones...)

	sort.Slice(allZones, func(i, j int) bool {
		return allZones[i].Short < allZones[j].Short
	})

	return allZones, nil
}

func zonesByExpansionID(db *sqlx.DB, expansionID int) ([]zone, error) {
	zones := []zone{}
	err := db.Select(&zones, "SELECT id, short_name, long_name FROM zone WHERE expansion = ?", expansionID)
	if err != nil {
		return nil, fmt.Errorf("db.Select: %w", err)
	}
	return zones, nil
}

func itemsByZone(db *sqlx.DB, zone string) ([]item, error) {
	items := []item{}
	err := db.Select(&items, `SELECT
    s2.zone,
    n.name AS npc_name,
    i.name AS item_name,
    i.proceffect,
    i.procrate,
    sn.effectid1,
    sn.effect_base_value1,
    sn.name AS spell_name,
	sn.id AS spell_id,
    i.classes
FROM
    items i
JOIN
    spells_new sn ON sn.id = i.proceffect
JOIN
    lootdrop_entries lde ON i.id = lde.item_id
JOIN
    loottable_entries lte ON lde.lootdrop_id = lte.lootdrop_id
JOIN
    npc_types n ON lte.loottable_id = n.loottable_id
JOIN
    spawnentry se ON n.id = se.npcid
JOIN
    spawngroup sg ON se.spawngroupid = sg.id
JOIN
    spawn2 s2 ON sg.id = s2.spawngroupid
WHERE
    i.proceffect > 0
    AND s2.zone = ?
GROUP BY item_name
`, zone)
	if err != nil {
		return nil, fmt.Errorf("items: %w", err)
	}
	return items, nil
}

type (
	Class   int
	Classes []Class
)

const (
	ClassWarrior Class = 1 << iota
	ClassCleric
	ClassPaladin
	ClassRanger
	ClassShadowKnight
	ClassDruid
	ClassMonk
	ClassBard
	ClassRogue
	ClassShaman
	ClassNecromancer
	ClassWizard
	ClassMagician
	ClassEnchanter
	ClassBeastlord
	ClassBerserker
)

var classToString = map[Class]string{
	ClassWarrior:      "Warrior",
	ClassCleric:       "Cleric",
	ClassPaladin:      "Paladin",
	ClassRanger:       "Ranger",
	ClassShadowKnight: "Shadow Knight",
	ClassDruid:        "Druid",
	ClassMonk:         "Monk",
	ClassBard:         "Bard",
	ClassRogue:        "Rogue",
	ClassShaman:       "Shaman",
	ClassNecromancer:  "Necromancer",
	ClassWizard:       "Wizard",
	ClassMagician:     "Magician",
	ClassEnchanter:    "Enchanter",
	ClassBeastlord:    "Beastlord",
	ClassBerserker:    "Berserker",
}

var classToShortString = map[Class]string{
	ClassWarrior:      "WAR",
	ClassCleric:       "CLR",
	ClassPaladin:      "PAL",
	ClassRanger:       "RNG",
	ClassShadowKnight: "SHD",
	ClassDruid:        "DRU",
	ClassMonk:         "MNK",
	ClassBard:         "BRD",
	ClassRogue:        "ROG",
	ClassShaman:       "SHM",
	ClassNecromancer:  "NEC",
	ClassWizard:       "WIZ",
	ClassMagician:     "MAG",
	ClassEnchanter:    "ENC",
	ClassBeastlord:    "BST",
	ClassBerserker:    "BER",
}

func (c Class) String() string {
	return classToString[c]
}

func (c Class) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func (c Classes) String() string {
	str := &strings.Builder{}

	if len(c) == 16 {
		return "ALL"
	}

	for i, class := range c {
		str.WriteString(classToShortString[class])
		if i != len(c)-1 {
			str.WriteString(" ")
		}
	}
	return str.String()
}

func (c Classes) MarshalJSON() ([]byte, error) {
	return []byte(`"` + c.String() + `"`), nil
}

func ClassesFromBitmask(mask int32) Classes {
	var classes Classes
	var i int32

	for i = 1; i <= mask; i <<= 1 {
		if i&mask != 0 {
			classes = append(classes, Class(i))
		}
	}
	return classes
}

type ItemInfo struct {
	ZoneShort  string
	NPCID      int
	ItemName   string
	AC         int
	HP         int   `db:"hp"`
	Mana       int   `db:"mana"`
	Int        int   `db:"int"`
	Wis        int   `db:"wis"`
	Type       int   `db:"item_type"`
	SpellID    int   `db:"spell_id"`
	SpellSPA   []int `db:"spell_spas"`
	SpellValue []int `db:"spell_values"`
	Classes    Classes
}

func dumpSpell(db *sqlx.DB, id int) {
	type Spell struct {
		ID                  int    `db:"id"`
		Name                string `db:"name"`
		Range               int    `db:"range"`
		Aoerange            int    `db:"aoerange"`
		Pushback            int    `db:"pushback"`
		Pushup              int    `db:"pushup"`
		CastTime            int    `db:"cast_time"`
		RecoveryTime        int    `db:"recovery_time"`
		RecastTime          int    `db:"recast_time"`
		BuffDurationFormula int    `db:"buffdurationformula"`
		BuffDuration        int    `db:"buffduration"`
		AEDuration          int    `db:"aeduration"`
		Mana                int    `db:"mana"`
		EffectBaseValue1    int    `db:"effect_base_value1"`
		EffectBaseValue2    int    `db:"effect_base_value2"`
		EffectBaseValue3    int    `db:"effect_base_value3"`
		EffectBaseValue4    int    `db:"effect_base_value4"`
		EffectBaseValue5    int    `db:"effect_base_value5"`
		EffectBaseValue6    int    `db:"effect_base_value6"`
		EffectBaseValue7    int    `db:"effect_base_value7"`
		EffectBaseValue8    int    `db:"effect_base_value8"`
		EffectBaseValue9    int    `db:"effect_base_value9"`
		EffectBaseValue10   int    `db:"effect_base_value10"`
		EffectBaseValue11   int    `db:"effect_base_value11"`
		EffectBaseValue12   int    `db:"effect_base_value12"`
		EffectLimitValue1   int    `db:"effect_limit_value1"`
		EffectLimitValue2   int    `db:"effect_limit_value2"`
		EffectLimitValue3   int    `db:"effect_limit_value3"`
		EffectLimitValue4   int    `db:"effect_limit_value4"`
		EffectLimitValue5   int    `db:"effect_limit_value5"`
		EffectLimitValue6   int    `db:"effect_limit_value6"`
		EffectLimitValue7   int    `db:"effect_limit_value7"`
		EffectLimitValue8   int    `db:"effect_limit_value8"`
		EffectLimitValue9   int    `db:"effect_limit_value9"`
		EffectLimitValue10  int    `db:"effect_limit_value10"`
		EffectLimitValue11  int    `db:"effect_limit_value11"`
		EffectLimitValue12  int    `db:"effect_limit_value12"`
		DescNum             int    `db:"descnum"`
		EffectDescNum       int    `db:"effectdescnum"`
		EffectDescNum2      int    `db:"effectdescnum2"`
		Max1                int    `db:"max1"`
		Max2                int    `db:"max2"`
		Max3                int    `db:"max3"`
		Max4                int    `db:"max4"`
		Max5                int    `db:"max5"`
		Max6                int    `db:"max6"`
		Max7                int    `db:"max7"`
		Max8                int    `db:"max8"`
		Max9                int    `db:"max9"`
		Max10               int    `db:"max10"`
		Max11               int    `db:"max11"`
		Max12               int    `db:"max12"`
		Icon                int    `db:"icon"`
		Memicon             int    `db:"memicon"`
		Components1         int    `db:"components1"`
		Components2         int    `db:"components2"`
		Components3         int    `db:"components3"`
		Components4         int    `db:"components4"`
		ComponentCounts1    int    `db:"component_counts1"`
		ComponentCounts2    int    `db:"component_counts2"`
		ComponentCounts3    int    `db:"component_counts3"`
		ComponentCounts4    int    `db:"component_counts4"`
		Noexpendreagent1    int    `db:"noexpendreagent1"`
		Noexpendreagent2    int    `db:"noexpendreagent2"`
		Noexpendreagent3    int    `db:"noexpendreagent3"`
		Noexpendreagent4    int    `db:"noexpendreagent4"`
		Formula1            int    `db:"formula1"`
		Formula2            int    `db:"formula2"`
		Formula3            int    `db:"formula3"`
		Formula4            int    `db:"formula4"`
		Formula5            int    `db:"formula5"`
		Formula6            int    `db:"formula6"`
		Formula7            int    `db:"formula7"`
		Formula8            int    `db:"formula8"`
		Formula9            int    `db:"formula9"`
		Formula10           int    `db:"formula10"`
		Formula11           int    `db:"formula11"`
		Formula12           int    `db:"formula12"`
		Lighttype           int    `db:"lighttype"`
		Goodeffect          int    `db:"goodeffect"`
		Activated           int    `db:"activated"`
		Resisttype          int    `db:"resisttype"`
		Resistdiff          int    `db:"resistdiff"`
		Effectid1           int    `db:"effectid1"`
		Effectid2           int    `db:"effectid2"`
		Effectid3           int    `db:"effectid3"`
		Effectid4           int    `db:"effectid4"`
		Effectid5           int    `db:"effectid5"`
		Effectid6           int    `db:"effectid6"`
		Effectid7           int    `db:"effectid7"`
		Effectid8           int    `db:"effectid8"`
		Effectid9           int    `db:"effectid9"`
		Effectid10          int    `db:"effectid10"`
		Effectid11          int    `db:"effectid11"`
		Effectid12          int    `db:"effectid12"`
		Targettype          int    `db:"targettype"`
		Basediff            int    `db:"basediff"`
		Skill               int    `db:"skill"`
		Zonetype            int    `db:"zonetype"`
		Environmenttype     int    `db:"environmenttype"`
		Timeofday           int    `db:"timeofday"`
		Classes1            int    `db:"classes1"`
		Classes2            int    `db:"classes2"`
		Classes3            int    `db:"classes3"`
		Classes4            int    `db:"classes4"`
		Classes5            int    `db:"classes5"`
		Classes6            int    `db:"classes6"`
		Classes7            int    `db:"classes7"`
		Classes8            int    `db:"classes8"`
		Classes9            int    `db:"classes9"`
		Classes10           int    `db:"classes10"`
		Classes11           int    `db:"classes11"`
		Classes12           int    `db:"classes12"`
		NpcNoLos            int    `db:"npc_no_los"`
		Field160            int    `db:"field160"`
		Numhitstype         int    `db:"numhitstype"`
		Numhits             int    `db:"numhits"`
		Aemaxtargets        int    `db:"aemaxtargets"`
		Maxtargets          int    `db:"maxtargets"`
		ShortBuffBox        int    `db:"short_buff_box"`
		RecourseLink        int    `db:"recourselink"`
	}

	spell := Spell{}
	err := db.Get(&spell, "SELECT id, name, `range`, aoerange, pushback, pushup, cast_time, recovery_time, recast_time, buffdurationformula, buffduration, aeduration, mana, effect_base_value1, effect_base_value2, effect_base_value3, effect_base_value4, effect_base_value5, effect_base_value6, effect_base_value7, effect_base_value8, effect_base_value9, effect_base_value10, effect_base_value11, effect_base_value12, effect_limit_value1, effect_limit_value2, effect_limit_value3, effect_limit_value4, effect_limit_value5, effect_limit_value6, effect_limit_value7, effect_limit_value8, effect_limit_value9, effect_limit_value10, effect_limit_value11, effect_limit_value12, max1, max2, max3, max4, max5, max6, max7, max8, max9, max10, max11, max12, icon, memicon, components1, components2, components3, components4, component_counts1, component_counts2, component_counts3, component_counts4, noexpendreagent1, noexpendreagent2, noexpendreagent3, noexpendreagent4, formula1, formula2, formula3, formula4, formula5, formula6, formula7, formula8, formula9, formula10, formula11, formula12, lighttype, goodeffect, activated, resisttype, resistdiff, effectid1, effectid2, effectid3, effectid4, effectid5, effectid6, effectid7, effectid8, effectid9, effectid10, effectid11, effectid12, targettype, basediff, skill, zonetype, environmenttype, timeofday, recourselink, short_buff_box, descnum, effectdescnum, effectdescnum2, npc_no_los, field160, numhitstype, numhits, aemaxtargets, maxtargets FROM spells_new WHERE id = ?", id)
	if err != nil {
		return
	}

	if spell.BuffDuration > 0 {
		duration := spdat.DurationCalc(spell.BuffDurationFormula, 70)
		if duration > spell.BuffDuration {
			duration = spell.BuffDuration
		}
		printWriterf("- Lasts %ds\n", duration*6)
	}

	if spell.DescNum > 0 {
		desc := dbstr.SpellDescription(spell.DescNum)
		if desc != "" {
			printWriterf("- Description: %s", desc)
		}
	}

	if spell.Resisttype != 0 {
		printWriterf("- Resist Type: %s (%d modifier)\n", spdat.ResistName(spell.Resisttype), spell.Resistdiff)
	}

	printWriterf("\n\n")
	dumpSpellEffect(1, spell.Effectid1, spell.EffectBaseValue1, spell.EffectLimitValue1, spell.Max1)
	dumpSpellEffect(2, spell.Effectid2, spell.EffectBaseValue2, spell.EffectLimitValue2, spell.Max2)
	dumpSpellEffect(3, spell.Effectid3, spell.EffectBaseValue3, spell.EffectLimitValue3, spell.Max3)
	dumpSpellEffect(4, spell.Effectid4, spell.EffectBaseValue4, spell.EffectLimitValue4, spell.Max4)
	dumpSpellEffect(5, spell.Effectid5, spell.EffectBaseValue5, spell.EffectLimitValue5, spell.Max5)
	dumpSpellEffect(6, spell.Effectid6, spell.EffectBaseValue6, spell.EffectLimitValue6, spell.Max6)
	dumpSpellEffect(7, spell.Effectid7, spell.EffectBaseValue7, spell.EffectLimitValue7, spell.Max7)
	dumpSpellEffect(8, spell.Effectid8, spell.EffectBaseValue8, spell.EffectLimitValue8, spell.Max8)
	dumpSpellEffect(9, spell.Effectid9, spell.EffectBaseValue9, spell.EffectLimitValue9, spell.Max9)
	dumpSpellEffect(10, spell.Effectid10, spell.EffectBaseValue10, spell.EffectLimitValue10, spell.Max10)
	dumpSpellEffect(11, spell.Effectid11, spell.EffectBaseValue11, spell.EffectLimitValue11, spell.Max11)
	dumpSpellEffect(12, spell.Effectid12, spell.EffectBaseValue12, spell.EffectLimitValue12, spell.Max12)

	if spell.RecourseLink != 0 {
		printWriterf("- Recourse:\n")
		dumpSpell(db, spell.RecourseLink)
	}
}

func dumpSpellEffect(index int, id, base int, limit int, max int) {
	if id == 254 {
		return
	}
	if id == 10 && base == 0 && limit == 0 {
		return
	}
	printWriterf("- Effect %d:", index)
	printWriterf(" %s", spdat.EffectName(id))
	if base != 0 {
		switch id {
		case 340:
			printWriterf(" %d%% chance", base)
		case 85:
			printWriterf(" [Spell %d](https://retributioneq.com/allaclone/?a=spell&id=%d)", base, base)
		default:
			printWriterf(" base %d", base)
		}
	}
	if limit != 0 {
		switch id {
		case 340, 374:
			printWriterf(" [Spell %d](https://retributioneq.com/allaclone/?a=spell&id=%d)", limit, limit)
		default:
			printWriterf(" limit %d", limit)
		}
	}
	if max != 0 {
		printWriterf(" max %d", max)
	}

	printWriterf("\n")
}

func printWriterf(format string, args ...interface{}) {
	w.WriteString(fmt.Sprintf(format, args...))
}

func questItemConvert(db *sqlx.DB, questItem *itemquest.ItemEntry, shortname string) (*item, error) {
	finalItem := &item{}
	err := db.Get(finalItem, `SELECT
	i.id AS item_id,
    i.name AS item_name,
    i.proceffect,
    i.procrate,
    sn.effectid1,
    sn.effect_base_value1,
    sn.name AS spell_name,
	sn.id AS spell_id,
    i.classes
FROM
    items i
JOIN
    spells_new sn ON sn.id = i.proceffect
WHERE
	i.id = ?
`, questItem.ID)
	if err != nil {
		return nil, fmt.Errorf("query %d: %w", questItem.ID, err)
	}
	finalItem.NPCName = questItem.NpcName
	finalItem.IsQuest = true
	finalItem.Zone = shortname
	return finalItem, nil
}

func CleanName(in string) string {
	out := in
	out = strings.ReplaceAll(out, "_", " ")
	out = strings.ReplaceAll(out, "-", "`")
	out = strings.ReplaceAll(out, "#", "")
	out = strings.ReplaceAll(out, "!", "")
	out = strings.ReplaceAll(out, "~", "")
	return out
}
