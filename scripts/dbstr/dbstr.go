package dbstr

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Load() error {
	data, err := os.ReadFile("dbstr_us.txt")
	if err != nil {
		return fmt.Errorf("dbstr_us.txt: %w", err)
	}

	for _, line := range strings.Split(string(data), "\n") {
		err = analyzeAADesc(line)
		if err != nil {
			return fmt.Errorf("analyzeDesc: %w", err)
		}
		err = analyzeTitle(line)
		if err != nil {
			return fmt.Errorf("analyzeTitle: %w", err)
		}
		err = analyzeSpellDesc(line)
		if err != nil {
			return fmt.Errorf("analyzeSpellDesc: %w", err)
		}

	}

	return nil
}

func AADescription(id int) string {
	return aaDescriptions[id]
}

func SpellDescription(id int) string {
	return spellDescriptions[id]
}

func Title(id int) string {
	return titles[id]
}

func analyzeAADesc(line string) error {
	records := strings.Split(line, "^")
	if len(records) != 4 {
		return fmt.Errorf("invalid line: %s", line)
	}
	if records[1] != "4" {
		return nil
	}
	desc := records[2]
	desc = strings.ReplaceAll(desc, "<br>", "\n")
	id, err := strconv.Atoi(records[0])
	if err != nil {
		return fmt.Errorf("strconv.Atoi: %w", err)
	}
	aaDescriptions[id] = desc
	return nil
}

func analyzeSpellDesc(line string) error {
	records := strings.Split(line, "^")
	if len(records) != 4 {
		return fmt.Errorf("invalid line: %s", line)
	}
	if records[1] != "6" {
		return nil
	}
	desc := records[2]
	desc = strings.ReplaceAll(desc, "<br>", "\n")
	id, err := strconv.Atoi(records[0])
	if err != nil {
		return fmt.Errorf("strconv.Atoi: %w", err)
	}
	spellDescriptions[id] = desc
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

var aaDescriptions = make(map[int]string)
var spellDescriptions = make(map[int]string)
var titles = make(map[int]string)
