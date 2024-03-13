package main

import (
	"context"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type zone struct {
	ID        int    `db:"zoneidnumber"`
	ShortName string `db:"short_name"`
	Expansion int    `db:"expansion"`
}

type mob struct {
	zone   *zone
	ID     int    `db:"id"`
	Name   string `db:"name"`
	Level  int    `db:"level"`
	Chance int    `db:"chance"`
}

type item struct {
	mob         *mob
	ID          int     `db:"id"`
	Name        string  `db:"name"`
	Chance      float32 `db:"chance"`
	ClickEffect int     `db:"clickeffect"`
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

	sympHeal := make(map[int][]*item)
	sympStrike := make(map[int][]*item)

	for _, zone := range zones {

		if zone.ID != 31 {
			continue
		}
		mobs, err := mobs(context.Background(), db, zone.ID)
		if err != nil {
			return fmt.Errorf("mobs %d: %w", zone.ID, err)
		}

		fmt.Println(zone)

		for _, mob := range mobs {
			mob.zone = zone
			items, err := items(context.Background(), db, mob.ID)
			if err != nil {
				return fmt.Errorf("items %d: %w", mob.ID, err)
			}
			fmt.Println(mob)
			for _, item := range items {
				item.mob = mob
				switch item.ClickEffect {
				case 24434:
					sympHeal[1] = append(sympHeal[1], item)
				case 24435:
					sympHeal[2] = append(sympHeal[2], item)
				case 24436:
					sympHeal[3] = append(sympHeal[3], item)
				case 24437:
					sympHeal[4] = append(sympHeal[4], item)
				case 24438:
					sympHeal[5] = append(sympHeal[5], item)
				case 24439:
					sympHeal[6] = append(sympHeal[6], item)
				case 24440:
					sympHeal[7] = append(sympHeal[7], item)
				case 24441:
					sympHeal[8] = append(sympHeal[8], item)
				case 24442:
					sympHeal[9] = append(sympHeal[9], item)
				case 24443:
					sympHeal[10] = append(sympHeal[10], item)
				case 33361:
					sympHeal[11] = append(sympHeal[11], item)
				case 33362:
					sympHeal[12] = append(sympHeal[12], item)
				case 33363:
					sympHeal[13] = append(sympHeal[13], item)
				case 33364:
					sympHeal[14] = append(sympHeal[14], item)
				case 33365:
					sympHeal[15] = append(sympHeal[15], item)
				case 24356:
					sympStrike[1] = append(sympStrike[1], item)
				case 24357:
					sympStrike[2] = append(sympStrike[2], item)
				case 24358:
					sympStrike[3] = append(sympStrike[3], item)
				case 24359:
					sympStrike[4] = append(sympStrike[4], item)
				case 24360:
					sympStrike[5] = append(sympStrike[5], item)
				case 24361:
					sympStrike[6] = append(sympStrike[6], item)
				case 24362:
					sympStrike[7] = append(sympStrike[7], item)
				case 24363:
					sympStrike[8] = append(sympStrike[8], item)
				case 24364:
					sympStrike[9] = append(sympStrike[9], item)
				case 24365:
					sympStrike[10] = append(sympStrike[10], item)
				case 33331:
					sympStrike[11] = append(sympStrike[11], item)
				case 33332:
					sympStrike[12] = append(sympStrike[12], item)
				case 33333:
					sympStrike[13] = append(sympStrike[13], item)
				case 33334:
					sympStrike[14] = append(sympStrike[14], item)
				case 33335:
					sympStrike[15] = append(sympStrike[15], item)
				}
				//fmt.Printf("%s %s %s %0.2f\n", zone.ShortName, mob.Name, item.Name, item.Chance)
			}
		}
	}

	for level, items := range sympHeal {
		for _, item := range items {
			fmt.Printf("%d %s %s %s %0.2f\n", level, item.mob.zone.ShortName, item.mob.Name, item.Name, item.Chance)
		}
	}
	for level, items := range sympStrike {
		for _, item := range items {
			fmt.Printf("%d %s %s %s %0.2f\n", level, item.mob.zone.ShortName, item.mob.Name, item.Name, item.Chance)
		}
	}

	return nil
}

func zones(ctx context.Context, db *sqlx.DB) ([]*zone, error) {
	results, err := db.NamedQueryContext(ctx, "SELECT zoneidnumber, short_name, expansion FROM zone WHERE expansion < 2", map[string]interface{}{})
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

func mobs(ctx context.Context, db *sqlx.DB, zoneID int) ([]*mob, error) {

	results, err := db.NamedQueryContext(ctx, "select n.name, n.id, se.chance from npc_types n INNER JOIN spawnentry se ON se.npcid = n.id WHERE n.id > :id1 and n.id < :id2;", map[string]interface{}{"id1": zoneID*1000 - 1, "id2": zoneID*1000 + 1000})
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	var mobs []*mob
	for results.Next() {
		entry := &mob{}
		err = results.StructScan(&entry)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}
		mobs = append(mobs, entry)
	}

	return mobs, nil
}

func items(ctx context.Context, db *sqlx.DB, mobID int) ([]*item, error) {
	results, err := db.NamedQueryContext(ctx, "select i.id, i.name, lde.chance FROM items i INNER JOIN lootdrop_entries lde ON lde.item_id = i.id INNER JOIN loottable_entries lte ON lte.lootdrop_id = lde.lootdrop_id INNER JOIN npc_types n ON n.loottable_id = lte.loottable_id WHERE n.id = :id;", map[string]interface{}{"id": mobID})
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
		resultsRC, err := db.NamedQueryContext(ctx, "select i.id, i.name FROM items i WHERE i.id = (:id+1000000);", map[string]interface{}{"id": entry.ID})
		if err != nil {
			return nil, fmt.Errorf("queryRC: %w", err)
		}
		for resultsRC.Next() {
			entryRC := &item{}
			err = resultsRC.StructScan(&entryRC)
			if err != nil {
				return nil, fmt.Errorf("scanRC: %w", err)
			}
			entryRC.Chance = entry.Chance
			entryRC.mob = entry.mob
			items = append(items, entryRC)
		}
		resultsApoc, err := db.NamedQueryContext(ctx, "select i.id, i.name FROM items i WHERE i.id = (:id+2000000);", map[string]interface{}{"id": entry.ID})
		if err != nil {
			return nil, fmt.Errorf("queryApoc: %w", err)
		}
		for resultsApoc.Next() {
			entryApoc := &item{}
			err = resultsApoc.StructScan(&entryApoc)
			if err != nil {
				return nil, fmt.Errorf("scanApoc: %w", err)
			}
			entryApoc.Chance = entry.Chance
			entryApoc.mob = entry.mob
			items = append(items, entryApoc)
		}

	}

	return items, nil
}
