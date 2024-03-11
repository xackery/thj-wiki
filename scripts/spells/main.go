package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type zone struct {
	ID        int    `db:"zoneidnumber"`
	ShortName string `db:"short_name"`
	Expansion int    `db:"expansion"`
	LongName  string `db:"long_name"`
}

type spell struct {
	SpellName     string `db:"spell_name"`
	SpellID       int    `db:"spell_id"`
	SpellLevel    int    `db:"spell_level"`
	NpcName       string `db:"npc_name"`
	NpcID         int    `db:"npc_id"`
	MerchantID    int    `db:"merchant_id"`
	ItemID        int    `db:"item_id"`
	ZoneShortName string `db:"zone_short_name"`
	ZoneLongName  string `db:"zone_long_name"`
}

func main() {
	err := run()
	if err != nil {
		fmt.Println("Failed:", err)
		os.Exit(1)
	}
}

func run() error {
	var err error
	var db *sqlx.DB
	db, err = sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", "root", "mariadb", "127.0.0.1", "3306", "peq"))
	if err != nil {
		return fmt.Errorf("sql.Open: %w", err)
	}
	zones, err := zones(context.Background(), db)
	if err != nil {
		return fmt.Errorf("zones: %w", err)
	}

	classes := []string{
		"Warrior",
		"Cleric",
		"Paladin",
		"Ranger",
		"Shadowknight",
		"Druid",
		"Monk",
		"Bard",
		"Rogue",
		"Shaman",
		"Necromancer",
		"Wizard",
		"Magician",
		"Enchanter",
		"Beastlord",
		"Berserker",
	}
	classShortNames := []string{
		"WAR",
		"CLR",
		"PAL",
		"RNG",
		"SHD",
		"DRU",
		"MNK",
		"BRD",
		"ROG",
		"SHM",
		"NEC",
		"WIZ",
		"MAG",
		"ENC",
		"BST",
		"BER",
	}
	for i := 0; i < 16; i++ {

		class := classes[i]
		classShortName := strings.ToLower(classShortNames[i])

		spells, err := spells(context.Background(), db, i+1)
		if err != nil {
			return fmt.Errorf("spells: %w", err)
		}

		out := "---\n"
		context := "Spells"
		if class == "Warrior" || class == "Monk" || class == "Rogue" || class == "Berserker" {
			context = "Disciplines"
		}
		if class == "Bard" {
			context = "Songs"
		}
		out += "title: " + class + " " + context + "\n"
		out += `images: [spells/images/` + classShortName + `.png]` + "\n"
		out += "bookHidden: true\n"
		out += fmt.Sprintf("description: %s %s.\n", class, context)
		out += "---\n"

		out += fmt.Sprintf("![%s %s](images/%s-banner.png)\n\n", class, context, classShortName)

		out += "This was auto generated and likely buggy. Expect issues!\n"
		for _, spell := range spells {
			spell.ZoneLongName = spellZone(strings.ToLower(spell.ZoneShortName), zones)
			if spell.ZoneLongName == "" {
				return fmt.Errorf("zone not found: %s", spell.ZoneShortName)
			}
		}
		var spellBest *spell
		isBazaar := false
		isClassic := false
		isFV := false
		isOT := false
		isPoK := false
		isLuclin := false

		era := ""
		out += "Name|Lvl|Era|Zone|NPC\n"
		out += "---|---|---|---|---\n"
		zones := []string{}
		for _, spell := range spells {
			if spellBest == nil {
				spellBest = spell
			}
			if spell.SpellLevel > 70 {
				continue
			}

			if spellBest.SpellName != spell.SpellName {
				if isBazaar {
					zones = []string{"Bazaar"}
				}

				if spellBest.ZoneShortName != "" && len(zones) == 0 {
					zones = append(zones, spellBest.ZoneLongName)
				}
				if len(zones) == 0 {
					zones = append(zones, "Not Sold")
				}
				isSkipped := false
				if len(zones) == 1 {
					zone := zones[0]
					if zone == "Sunset Home" {
						isSkipped = true
					}
				}
				if !isSkipped {
					if era == "" {
						era = "Unknown"
					}
					out += fmt.Sprintf("%s|%d|%s|%s|%s\n", spellBest.SpellName, spellBest.SpellLevel, era, strings.Join(zones, ", "), strings.ReplaceAll(strings.ReplaceAll(spellBest.NpcName, "_", " "), "#", ""))
				}
				isBazaar = false
				isFV = false
				isOT = false
				isClassic = false
				isPoK = false
				era = ""
				zones = []string{}
				spellBest = spell
			}
			if !isBazaar && spell.ZoneShortName == "bazaar" {
				zones = append(zones, "Bazaar")
				isBazaar = true
				era = "Classic"
				continue
			}
			if isBazaar {
				continue
			}

			if !isClassic && spell.ZoneShortName == "ecommons" {
				isClassic = true
				zones = append(zones, "East Commonlands")
				era = "Classic"
				continue
			}

			if isClassic {
				continue
			}

			if !isFV && spell.ZoneShortName == "firiona" {
				zones = append(zones, "Firiona Vie")
				isFV = true
				era = "Kunark"
				continue
			}

			if !isOT && spell.ZoneShortName == "overthere" {
				zones = append(zones, "Overthere")
				isOT = true
				era = "Kunark"
				continue
			}

			if isFV || isOT {
				continue
			}

			if !isPoK && spell.ZoneShortName == "poknowledge" {
				zones = append(zones, "Plane of Knowledge")
				isPoK = true
				era = "PoP"
				continue
			}

			if !isLuclin && spell.ZoneShortName == "shadowhaven" {
				zones = append(zones, "Shadow Haven")
				isLuclin = true
				era = "Luclin"
				continue
			}

			if isLuclin {
				continue
			}
			if isPoK {
				continue
			}
			if isClassic {
				continue
			}

		}
		err = os.WriteFile(fmt.Sprintf("content/spells/%s.md", classShortName), []byte(out), 0644)
		if err != nil {
			return fmt.Errorf("write file: %w", err)
		}
	}

	return nil
}

func zones(ctx context.Context, db *sqlx.DB) ([]*zone, error) {
	results, err := db.NamedQueryContext(ctx, "SELECT zoneidnumber, short_name, expansion, long_name FROM zone", map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	var zones []*zone
	for results.Next() {
		entry := &zone{}
		err = results.StructScan(&entry)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}
		zones = append(zones, entry)
	}

	return zones, nil
}

func spellZone(shortName string, zones []*zone) string {
	for _, zone := range zones {
		if zone.ShortName != shortName {
			continue
		}
		return zone.LongName
	}
	return ""

}

func spells(ctx context.Context, db *sqlx.DB, classID int) ([]*spell, error) {

	results, err := db.NamedQueryContext(ctx, fmt.Sprintf(`select n.name npc_name, s2.zone zone_short_name, sn.classes%d spell_level, sn.name spell_name, n.id npc_id, m.merchantid merchant_id, i.id item_id, sn.id spell_id from spells_new sn
	inner join items i on i.scrolleffect = sn.id
	inner join merchantlist m on m.item = i.id
	inner join npc_types n on n.merchant_id = m.merchantid
	inner join spawnentry se on se.npcID = n.id
	inner join spawn2 s2 on s2.spawngroupID = se.spawngroupID
	where sn.classes%d < 255 order by sn.classes%d asc, sn.name asc;`, classID, classID, classID), map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	var mobs []*spell
	for results.Next() {
		entry := &spell{}
		err = results.StructScan(&entry)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}
		mobs = append(mobs, entry)
	}

	return mobs, nil
}
