package itemquest

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var (
	regNumbers = regexp.MustCompile("(?m)([0-9]+)")
)

type ItemEntry struct {
	ID      int
	NpcName string
}

func LoadZone(shortname string) ([]*ItemEntry, error) {
	itemQuests := []*ItemEntry{}
	err := filepath.Walk("../../../thj/quests/"+shortname+"/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("walk %q: %w", path, err)
		}

		// ignore directories
		if info.IsDir() {
			return nil
		}
		// only parse .pl and .lua
		language := strings.ToLower(filepath.Ext(path))
		if language != ".lua" && language != ".pl" {
			return nil
		}
		if language == ".pl" {
			luaPath := strings.Replace(path, ".pl", ".lua", 1)
			if _, err := os.Stat(luaPath); err == nil {
				return nil
			}
		}
		npcName := strings.Replace(filepath.Base(path), language, "", 1)
		zoneName := strings.Replace(path, "../../../thj/quests/", "", 1)
		zoneIndex := strings.Index(zoneName, "/")
		if zoneIndex != -1 {
			zoneName = zoneName[0:zoneIndex]
		}
		if !strings.EqualFold(zoneName, shortname) {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		r := bufio.NewReader(f)

		lineNumber := 0
		itemsCommentLineNumber := -1
		lineType := 0
		items := []string{}
		isCommentBlock := false
		for {
			lineNumber++
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			}

			if line[len(line)-1] == '\n' {
				line = line[0 : len(line)-1]
			}

			if strings.Contains(line, "--[[") {
				isCommentBlock = true
			}
			if isCommentBlock && strings.Contains(line, "]]---") {
				isCommentBlock = false
			}

			if !isCommentBlock && lineType == 0 && itemsLineType(line, language) > 0 {
				itemsCommentLineNumber = lineNumber
				lineType = itemsLineType(line, language)
			}

			if !isCommentBlock && itemsCommentLineNumber == -1 && firstCharacter(line) != commentMarker(language) {
				itemsCommentLineNumber = lineNumber
			}

			possibleItems := findItems(line, language)
			for _, newItem := range possibleItems {
				isNew := true
				for _, oldItem := range items {
					if oldItem != newItem {
						continue
					}
					isNew = false
					break
				}
				if !isNew {
					continue
				}
				items = append(items, newItem)
			}
		}
		if len(items) == 0 {
			return nil
		}
		for _, item := range items {
			itemID, err := strconv.Atoi(item)
			if err != nil {
				return fmt.Errorf("strconv.Atoi: %w", err)
			}

			itemQuests = append(itemQuests, &ItemEntry{
				ID:      itemID,
				NpcName: npcName,
			})
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return itemQuests, nil
}

// itemsLineType has 3 possible values: 0 (none), 1: (items: generated line), 2: (!items: manually edited line to skip)
func itemsLineType(line string, language string) int {
	if !isComment(line, language) {
		return 0
	}
	commentMark := commentMarker(language)
	idx := strings.Index(line, string(commentMark))
	if idx == -1 {
		return 0
	}
	line = strings.ReplaceAll(line, string(commentMark), "")
	line = strings.TrimSpace(line)
	idx = strings.Index(line, "items:")
	if idx == 0 {
		return 1
	}
	idx = strings.Index(line, "!items:")
	if idx == 0 {
		return 2
	}
	return 0
}

func isComment(line string, language string) bool {
	if firstCharacter(line) == commentMarker(language) {
		return true
	}
	return false
}

// commentMarker returns either // or -- based on the language provided
func commentMarker(language string) string {
	if language == ".lua" {
		return "-"
	}
	if language == ".pl" {
		return "#"
	}
	return "-"
}

// firstCharacter returns the first valid character detected on a line
func firstCharacter(line string) string {
	line = strings.TrimSpace(line)
	if len(line) == 0 {
		return ""
	}
	return line[0:1]
}

func findItems(line string, language string) []string {

	idx := -1
	if language == ".lua" {
		idx = strings.Index(strings.ToLower(line), "summonitem(")
		if idx != -1 {
			idx += 11
			return findItemIDs(idx, line)
		}
		idx = strings.Index(strings.ToLower(line), "check_turn_in(")
		if idx != -1 {
			idx += 14
			return findItemIDs(idx, line)
		}

	}
	if language == ".pl" {
		idx = strings.Index(strings.ToLower(line), "summonitem(")
		if idx != -1 {
			idx += 11
			return findItemIDs(idx, line)
		}
		idx = strings.Index(strings.ToLower(line), "check_handin(")
		if idx != -1 {
			idx += 14
			return findItemIDs(idx, line)
		}
	}
	return []string{}
}

func findItemIDs(index int, line string) []string {
	items := []string{}
	itemMatches := regNumbers.FindAllStringSubmatch(line[index:], -1)
	for _, groups := range itemMatches {
		for _, match := range groups {
			id, err := strconv.Atoi(match)
			if err != nil {
				continue
			}
			if id < 1000 {
				continue
			}
			items = append(items, match)
		}
	}
	return items
}
