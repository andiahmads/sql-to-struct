package parser

import (
	"fmt"
	"regexp"
	"strings"
)

var typeMapping = map[string]string{
	"char":     "string",
	"varchar":  "string",
	"text":     "string",
	"longtext": "string",
	"uuid":     "string",
	"enum":     "string",

	"int":      "int",
	"tinyint":  "bool",
	"smallint": "int",
	"bigint":   "int64",

	"datetime":  "string",
	"timestamp": "string",
}

func ConvertSQLToStruct(sql string, structName string) (string, error) {

	// Ambil nama tabel
	re := regexp.MustCompile(`(?i)create\s+table\s+` +
		"`?([a-zA-Z0-9_]+)`?")
	matches := re.FindStringSubmatch(sql)

	if len(matches) < 2 {
		return "", fmt.Errorf("CREATE TABLE Not found")
	}

	tableName := matches[1]

	if structName == "" {
		structName = ToCamel(tableName)
	}

	// Ambil body antara tanda kurung pertama dan terakhir
	start := strings.Index(sql, "(")
	end := strings.LastIndex(sql, ")")

	if start == -1 || end == -1 || end <= start {
		return "", fmt.Errorf("Invalid SQL format")
	}

	body := sql[start+1 : end]

	// Pecah per baris
	lines := strings.Split(body, "\n")

	// Regex kolom
	// columnRegex := regexp.MustCompile("`([^`]+)`\\s+([^\\s,]+)")
	var columnRegex = regexp.MustCompile(`^\s*` +
		"`?" + // optional backtick
		`([a-zA-Z0-9_]+)` + // column name
		"`?" + // optional backtick
		`\s+` +
		`([a-zA-Z0-9()]+)`) // type

	var fields []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		// skip constraint dan key
		if strings.HasPrefix(strings.ToUpper(line), "PRIMARY") ||
			strings.HasPrefix(strings.ToUpper(line), "UNIQUE") ||
			strings.HasPrefix(strings.ToUpper(line), "KEY") ||
			strings.HasPrefix(strings.ToUpper(line), "CONSTRAINT") ||
			strings.HasPrefix(strings.ToUpper(line), ")") {
			continue
		}

		col := columnRegex.FindStringSubmatch(line)
		if len(col) < 3 {
			continue
		}

		colName := col[1]
		colType := strings.ToLower(col[2])

		// ambil base type
		baseType := regexp.MustCompile(`^[a-zA-Z]+`).FindString(colType)

		goType := MapToGoType(baseType)

		fields = append(fields, fmt.Sprintf(
			"%s %s `json:\"%s\"`",
			ToCamel(colName),
			goType,
			colName,
		))
	}

	// Compose struct
	result := "type " + structName + " struct {\n"
	for _, f := range fields {
		result += "    " + f + "\n"
	}
	result += "}"

	return result, nil
}

func MapToGoType(sqlType string) string {
	if goType, ok := typeMapping[sqlType]; ok {
		return goType
	}
	return "string"
}

func ToCamel(input string) string {
	parts := strings.Split(input, "_")
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}
