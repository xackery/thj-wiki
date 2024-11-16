package generator

import (
	"fmt"
	"os"
	"sort"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/thj-wiki/scripts/dbstr"
)

func TestAugGenerate(t *testing.T) {
	err := runAug()
	if err != nil {
		fmt.Println("Failed:", err)
		os.Exit(1)
	}
}

func runAug() error {
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

	w, err = os.Create("../../content/equipment-guide/augs/_index.en.md")
	if err != nil {
		return fmt.Errorf("os.Create: %w", err)
	}
	defer w.Close()

	w.WriteString(fmt.Sprintf(`---
title: Augs
weight: 5
chapter: true
description: Aug List
images: [images/aug.png]
---

![Augs](images/aug.png)

This page was last updated on %s.

Found an error on this page? Check tracked issues and [report new ones in the discord](<https://discord.com/channels/1204418766318862356/1307325765636980736/1307325765636980736>)

`, time.Now().Format("2006-01-02")))

	trackedAugs := []string{}

	for i := 0; i < 8; i++ {
		zones, err := parseZones(db, i)
		if err != nil {
			return fmt.Errorf("parseZones: %w", err)
		}
		if i > 2 {
			break
		}

		switch i {
		case 0:
			w.WriteString("## Classic\n")
			fmt.Printf("Classic Zones: %d\n", len(zones))

		case 1:
			w.WriteString("## Kunark\n")
			fmt.Printf("Kunark Zones: %d\n", len(zones))
		case 2:
			w.WriteString("## Velious\n")
			fmt.Printf("Velious Zones: %d\n", len(zones))
		case 3:
			w.WriteString("## Luclin\n")
			fmt.Printf("Luclin Zones: %d\n", len(zones))
		case 4:
			w.WriteString("## Planes of Power\n")
			fmt.Printf("Planes of Power Zones: %d\n", len(zones))
		case 5:
			w.WriteString("## Gates of Discord\n")
			fmt.Printf("Gates of Discord Zones: %d\n", len(zones))
		case 6:
			w.WriteString("## Omens of War\n")
			fmt.Printf("Omens of War Zones: %d\n", len(zones))
		case 7:
			w.WriteString("## Dragons of Norrath\n")
			fmt.Printf("Dragons of Norrath Zones: %d\n", len(zones))
		}
		fmt.Println("Zones:", len(zones))
		w.WriteString("Zone|NPC|Item|Stats\n")
		w.WriteString("---|---|---|---\n")

		for _, zone := range zones {

			type baseAugItem struct {
				Zone     string `db:"zone"`
				NPCName  string `db:"npc_name"`
				NPCID    int    `db:"npc_id"`
				ItemID   int    `db:"item_id"`
				ItemName string `db:"item_name"`
			}

			baseAugItems := []baseAugItem{}
			err := db.Select(&baseAugItems, `SELECT
    s2.zone,
    n.name AS npc_name,
    i.id AS item_id,
	i.name AS item_name,
	n.id as npc_id
FROM
    items i
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
    i.augtype > 0
    AND s2.zone = ?
GROUP BY i.name
`, zone.Short)
			if err != nil {
				return fmt.Errorf("baseAugItems: %w", err)
			}

			type augItem struct {
				Zone     string `db:"zone"`
				ItemID   int    `db:"item_id"`
				NPCID    int    `db:"npc_id"`
				NPCName  string `db:"npc_name"`
				ItemName string `db:"item_name"`
				Classes  int    `db:"classes"`
				AC       int    `db:"ac"`
				HP       int    `db:"hp"`
				Mana     int    `db:"mana"`
				End      int    `db:"endur"`
				Str      int    `db:"astr"`
				Sta      int    `db:"asta"`
				Agi      int    `db:"aagi"`
				Atk      int    `db:"attack"`
			}

			items := []augItem{}

			for _, baseAugItem := range baseAugItems {

				isUnique := true
				for _, trackedAug := range trackedAugs {
					if baseAugItem.Zone+baseAugItem.ItemName == trackedAug {
						isUnique = false
						break
					}
				}
				if !isUnique {
					continue
				}
				trackedAugs = append(trackedAugs, baseAugItem.Zone+baseAugItem.ItemName)
				item := augItem{}
				err = db.Get(&item, `SELECT
				i.name item_name,
				i.classes,
				i.ac,
				i.hp,
				i.mana,
				i.endur,
				i.astr,
				i.asta,
				i.aagi,
				i.attack
				FROM items i
				WHERE i.id = ?
				`, baseAugItems[0].ItemID+2000000)
				if err != nil {
					fmt.Println("Error getting item", baseAugItem.ItemID+2000000, err)
					continue
				}
				item.NPCName = baseAugItem.NPCName
				item.Zone = baseAugItem.Zone
				item.ItemID = baseAugItem.ItemID + 2000000
				item.NPCID = baseAugItem.NPCID
				items = append(items, item)
			}

			sort.Slice(items, func(i, j int) bool {
				return (items[i].Zone < items[j].Zone && items[i].ItemName < items[j].ItemName)
			})

			for _, item := range items {

				printWriterf(`%s|{{<npc id=%d name="%s">}}|{{<item id=%d name="%s">}}`, zone.Short, item.NPCID, item.NPCName, item.ItemID, item.ItemName)
				printWriterf(" (%s)|", ClassesFromBitmask(int32(item.Classes)))
				if item.AC > 0 {
					printWriterf(" AC %d", item.AC)
				}
				if item.HP > 0 {
					printWriterf(" HP %d", item.HP)
				}
				if item.Mana > 0 {
					printWriterf(" Mana %d", item.Mana)
				}
				if item.End > 0 {
					printWriterf(" End %d", item.End)
				}
				if item.Str > 0 {
					printWriterf(" Str %d", item.Str)
				}
				if item.Sta > 0 {
					printWriterf(" Sta %d", item.Sta)
				}
				if item.Agi > 0 {
					printWriterf(" Agi %d", item.Agi)
				}
				if item.Atk > 0 {
					printWriterf(" Atk %d", item.Atk)
				}

				printWriterf(" (TODO)")

				printWriterf("\n")
			}
		}
	}

	return nil
}
