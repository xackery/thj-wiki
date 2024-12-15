package generator

import (
	"bytes"
	"fmt"
	"os"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	wZem *os.File
)

type Zem struct {
	Short    string  `db:"short_name"`
	Zem      float32 `db:"zone_exp_multiplier"`
	LongName string  `db:"long_name"`
}

func TestZEM(t *testing.T) {
	err := runZEM()
	if err != nil {
		fmt.Println("Failed:", err)
		os.Exit(1)
	}
}

func runZEM() error {
	var err error
	var db *sqlx.DB

	wZem, err = os.Create("../../content/exploration-and-combat/zem.md")
	if err != nil {
		return fmt.Errorf("os.Create: %w", err)
	}

	context := "ZEM"
	wZem.WriteString(fmt.Sprintf(`---
title: %s
weight: 5
chapter: true
description: %s items
images: [images/%s.png]
---

![%s](images/%s.png)

This page was last updated on %s.

ZEM stands for Zone Experience Modifier. This is a zone-wide buff to experience that improves gains.

Found an error on this page? Check tracked issues and [report new ones in the discord](<https://discord.com/channels/1204418766318862356/1307331696693350501/1307331696693350501>)


How this is calculated:
- The base ZEM of most zones is 2.0
- All ZEMs are subjected by 1.0, making average 1.0
- I multiply the ZEM by 100 to make it a percent, so all zones untouched are at 100%%
- So a zone with 150%% ZEM is 50%% more experience than average

`, context, context, context, context, context, time.Now().Format("2006-01-02")))

	eqDB := os.Getenv("EQ_DB")
	if eqDB == "" {
		return fmt.Errorf("EQ_DB not set")
	}

	noZemBuf := &bytes.Buffer{}
	db, err = sqlx.Open("mysql", eqDB)
	if err != nil {
		return fmt.Errorf("sql.Open: %w", err)
	}
	for i := 0; i < 6; i++ {

		zones, err := parseZones(db, i)
		if err != nil {
			return fmt.Errorf("parseZones: %w", err)
		}
		switch i {
		case 0:
			fmt.Fprintf(wZem, "## Classic\n")
			fmt.Fprintf(noZemBuf, "## Classic No ZEM\n")
			fmt.Printf("Classic Zones: %d\n", len(zones))
		case 1:
			fmt.Fprintf(wZem, "## Kunark\n")
			fmt.Fprintf(noZemBuf, "## Kunark No ZEM\n")
			fmt.Printf("Kunark Zones: %d\n", len(zones))
		case 2:
			fmt.Fprintf(wZem, "## Velious\n")
			fmt.Fprintf(noZemBuf, "## Velious No ZEM\n")
			fmt.Printf("Velious Zones: %d\n", len(zones))
		case 3:
			fmt.Fprintf(wZem, "## Luclin\n")
			fmt.Fprintf(noZemBuf, "## Luclin No ZEM\n")
			fmt.Printf("Luclin Zones: %d\n", len(zones))
		case 4:
			fmt.Fprintf(wZem, "## Planes of Power\n")
			fmt.Fprintf(noZemBuf, "## Planes of Power No ZEM\n")
			fmt.Printf("Planes of Power Zones: %d\n", len(zones))
		case 5:
			fmt.Fprintf(wZem, "## Gates of Discord\n")
			fmt.Fprintf(noZemBuf, "## Gates of Discord No ZEM\n")
			fmt.Printf("Gates of Discord Zones: %d\n", len(zones))
		case 6:
			fmt.Fprintf(wZem, "## Omens of War\n")
			fmt.Fprintf(noZemBuf, "## Omens of War\n")
			fmt.Printf("Omens of War Zones: %d\n", len(zones))
		case 7:
			fmt.Fprintf(wZem, "## Dragons of Norrath\n")
			fmt.Fprintf(noZemBuf, "## Dragons of Norrath No ZEM\n")
			fmt.Printf("Dragons of Norrath Zones: %d\n", len(zones))
		}

		zoneNames := []string{}
		for _, zone := range zones {
			zoneNames = append(zoneNames, zone.Short)
		}

		zems, err := zonesByZoneShort(db, zoneNames)
		if err != nil {
			return fmt.Errorf("zonesByZoneShort: %w", err)
		}

		wZem.WriteString("Zone | ZEM | Name\n")
		wZem.WriteString("---- | --- | ---\n")
		noZemBuf.WriteString("Zone | ZEM | Name\n")
		noZemBuf.WriteString("---- | --- | ---\n")
		for _, zem := range zems {
			zemPercent := (zem.Zem - 1) * 100

			if zemPercent == 100 {
				noZemBuf.WriteString(fmt.Sprintf("%s | %d%% | %s\n", zem.Short, int(zemPercent), zem.LongName))
				continue
			}

			wZem.WriteString(fmt.Sprintf("%s | %d%% | %s\n", zem.Short, int(zemPercent), zem.LongName))
		}
		wZem.WriteString("\n\n")
		noZemBuf.WriteString("\n\n")

	}

	wZem.WriteString(noZemBuf.String())

	return nil
}

func zonesByZoneShort(db *sqlx.DB, zones []string) ([]Zem, error) {

	zems := []Zem{}

	query, args, err := sqlx.In("SELECT short_name, long_name, zone_exp_multiplier FROM zone WHERE short_name IN (?) ORDER BY zone_exp_multiplier DESC, short_name ASC", zones)
	if err != nil {
		return nil, fmt.Errorf("sqlx.In: %w", err)
	}
	query = db.Rebind(query)
	err = db.Select(&zems, query, args...)
	if err != nil {
		return nil, fmt.Errorf("db zone: %w", err)
	}
	return zems, nil
}
