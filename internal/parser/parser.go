package parser

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/cahfofpai/birthday.md/internal/models"
)

// LineType represents the type of a line in the birthday.md file
type LineType int

const (
	LineTypeInvalid LineType = iota
	LineTypeBirthday
	LineTypeHeading
	LineTypeComment
	LineTypeBlank
)

// Parser is responsible for parsing the birthday.md file
type Parser struct {
	filePath string
	errors   []string
}

// NewParser creates a new Parser instance
func NewParser(filePath string) *Parser {
	return &Parser{
		filePath: filePath,
		errors:   make([]string, 0),
	}
}

// Parse parses the birthday.md file and returns a slice of Birthday instances
func (p *Parser) Parse() ([]*models.Birthday, error) {
	file, err := os.Open(p.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	birthdays := make([]*models.Birthday, 0)
	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		lineType := p.getLineType(line)

		switch lineType {
		case LineTypeBirthday:
			birthday, err := p.parseBirthdayLine(line)
			if err != nil {
				p.errors = append(p.errors, fmt.Sprintf("Line %d: %s", lineNumber, err.Error()))
				continue
			}
			birthdays = append(birthdays, birthday)
		case LineTypeInvalid:
			p.errors = append(p.errors, fmt.Sprintf("Line %d: Invalid line format: %s", lineNumber, line))
		// The other oineTypes are ignored
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return birthdays, nil
}

// GetErrors returns the errors encountered during parsing
func (p *Parser) GetErrors() []string {
	return p.errors
}

// getLineType determines the type of a line
func (p *Parser) getLineType(line string) LineType {
	line = strings.TrimSpace(line)

	if line == "" {
		return LineTypeBlank
	}

	if strings.HasPrefix(line, "#") {
		return LineTypeHeading
	}

	if strings.HasPrefix(line, "<!--") && strings.HasSuffix(line, "-->") {
		return LineTypeComment
	}

	// Check if the line matches the birthday format
	birthdayRegexWithYear := regexp.MustCompile(`^\d{1,2}\.\d{1,2}\.\d{4}\s+.+$`)
	birthdayRegexWithoutYear := regexp.MustCompile(`^\d{1,2}\.\d{1,2}\.\s+.+$`)

	if birthdayRegexWithYear.MatchString(line) || birthdayRegexWithoutYear.MatchString(line) {
		return LineTypeBirthday
	}

	return LineTypeInvalid
}

// parseBirthdayLine parses a birthday line and returns a Birthday instance
func (p *Parser) parseBirthdayLine(line string) (*models.Birthday, error) {
	line = strings.TrimSpace(line)

	// Check if the line has a year
	hasYear := true
	birthdayRegexWithoutYear := regexp.MustCompile(`^\d{1,2}\.\d{1,2}\.\s+.+$`)
	if birthdayRegexWithoutYear.MatchString(line) {
		hasYear = false
	}

	// Split the line into date and name parts
	parts := strings.SplitN(line, " ", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid birthday format: %s", line)
	}

	date := parts[0]
	name := strings.TrimSpace(parts[1])

	// Parse the date
	dateParts := strings.Split(date, ".")
	if (hasYear && len(dateParts) != 3) || (!hasYear && len(dateParts) != 2) {
		return nil, fmt.Errorf("invalid date format: %s", date)
	}

	day, err := strconv.Atoi(dateParts[0])
	if err != nil {
		return nil, fmt.Errorf("invalid day: %s", dateParts[0])
	}

	month, err := strconv.Atoi(dateParts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid month: %s", dateParts[1])
	}

	var year int
	if hasYear {
		year, err = strconv.Atoi(dateParts[2])
		if err != nil {
			return nil, fmt.Errorf("invalid year: %s", dateParts[2])
		}
	} else {
		year = 0
	}

	// Validate day and month
	if day < 1 || day > 31 {
		return nil, fmt.Errorf("invalid day: %d", day)
	}

	if month < 1 || month > 12 {
		return nil, fmt.Errorf("invalid month: %d", month)
	}

	return models.NewBirthday(name, day, month, year, hasYear), nil
}