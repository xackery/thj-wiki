package main

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type item struct {
	ItemID     int     `db:"item_id"`
	ItemName   string  `db:"item_name"`
	ItemSlots  int     `db:"item_slots"`
	NpcName    string  `db:"npc_name"`
	NpcLevel   int     `db:"npc_level"`
	DropChance float32 `db:"chance"`
	Awakened   *itemAwakened
}

type itemAwakened struct {
	Name     string `db:"name"`
	HealAmt  int    `db:"healamt"`
	SpellAmt int    `db:"spelldmg"`
	HP       int    `db:"hp"`
	Mana     int    `db:"mana"`
}

func (e *itemAwakened) String() string {
	return fmt.Sprintf("%s: %d hp, %d mana, %d heal, %d spell", strings.TrimSuffix(e.Name, " (Awakened)"), e.HP, e.Mana, e.HealAmt, e.SpellAmt)
}

type byHealAmt []*item

func (a byHealAmt) Len() int           { return len(a) }
func (a byHealAmt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byHealAmt) Less(i, j int) bool { return a[i].Awakened.HealAmt > a[j].Awakened.HealAmt }

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
	if len(os.Args) < 2 {
		return fmt.Errorf("missing zone")
	}

	zone := os.Args[1]
	items, err := items(context.Background(), db, zone)
	if err != nil {
		return fmt.Errorf("items: %w", err)
	}
	finals := []*item{}

	lastNpc := ""
	for _, item := range items {
		if item.ItemSlots == 0 {
			continue
		}
		if lastNpc == item.NpcName {
			continue
		}
		awakened, err := awakened(context.Background(), db, item.ItemID)
		if err != nil {
			//return fmt.Errorf("awakened: %w", err)
			continue
		}
		lastNpc = item.NpcName

		item.Awakened = awakened
		if awakened.HealAmt < 5 && awakened.SpellAmt < 5 {
			continue
		}

		finals = append(finals, item)

	}

	sort.Sort(byHealAmt(finals))

	lastNPC := ""
	lastItem := ""
	for _, item := range finals {
		if lastNPC == item.NpcName && lastItem == item.ItemName {
			continue
		}
		lastNPC = item.NpcName
		lastItem = item.ItemName

		fmt.Println(item.NpcName, item.NpcLevel, item.DropChance, item.Awakened)
	}

	return nil
}

func items(ctx context.Context, db *sqlx.DB, zone string) ([]*item, error) {
	results, err := db.NamedQueryContext(ctx, `select i.id item_id, i.name item_name, i.slots item_slots, n.name npc_name, n.level npc_level, lde.chance FROM items i
	inner join lootdrop_entries lde on lde.item_id = i.id
	inner join loottable_entries lte on lte.lootdrop_id = lde.lootdrop_id
	inner join npc_types n on n.loottable_id = lte.loottable_id
	inner join spawnentry se on se.npcID = n.id
	inner join spawn2 s2 on s2.spawngroupID = se.spawngroupID
	where s2.zone = :zone ORDER BY n.id`, map[string]interface{}{"zone": zone})
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	var items []*item
	for results.Next() {
		entry := &item{}
		err = results.StructScan(&entry)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}
		items = append(items, entry)
	}

	return items, nil
}

func awakened(ctx context.Context, db *sqlx.DB, itemID int) (*itemAwakened, error) {
	results, err := db.NamedQueryContext(ctx, `select name, healamt, spelldmg, hp, mana from items i where id = :id`, map[string]interface{}{"id": itemID + 2000000})
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	var item *itemAwakened
	for results.Next() {
		entry := &itemAwakened{}
		err = results.StructScan(&entry)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}
		item = entry
	}
	if item == nil {
		return nil, fmt.Errorf("no item for %d", itemID+2000000)
	}

	return item, nil
}
