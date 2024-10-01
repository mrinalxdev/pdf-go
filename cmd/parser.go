package cmd

import (
	"strings"
	"unicode"
)

type DocumentStructure struct {
	Title    string
	MetaData map[string]string
	Sections []Section
}

type Section struct {
	Header     string
	Paragraphs []string
}

func ParseText(text string) DocumentStructure {
	lines := strings.Split(text, "\n")
	doc := DocumentStructure{
		MetaData: make(map[string]string),
	}

	var currentSection *Section
	inMetaData := true

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		if trimmedLine == "" {
			continue
		}

		if inMetaData {
			if strings.Contains(trimmedLine, ":") {
				parts := strings.SplitN(trimmedLine, ":", 2)
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				doc.MetaData[key] = value
				continue
			} else {
				inMetaData = false // coz the first non-Metadata line is assumed to be the title
				doc.Title = trimmedLine
				continue
			}
		}

		if isHeader(trimmedLine) {
			if currentSection != nil {
				doc.Sections = append(doc.Sections, *currentSection)
			}
			currentSection = &Section{
				Header: trimmedLine,
			}
		} else {
			if currentSection == nil {
				currentSection = &Section{}
			}

			if len(currentSection.Paragraphs) > 0 && !strings.HasSuffix(currentSection.Paragraphs[len(currentSection.Paragraphs)-1], ".") {
				currentSection.Paragraphs[len(currentSection.Paragraphs)-1] += " " + trimmedLine
			} else {
				currentSection.Paragraphs = append(currentSection.Paragraphs, trimmedLine)
			}
		}

	}

	if currentSection != nil {
		doc.Sections = append(doc.Sections, *currentSection)
	}

	return doc
}

func isHeader(line string) bool {
	if len(line) < 4 {
		return false
	}

	allCaps := true
	for _, r := range line {
		if !unicode.IsUpper(r) && !unicode.IsSpace(r) && !unicode.IsPunct(r) {
			allCaps = false
			break
		}
	}

	return allCaps
}
