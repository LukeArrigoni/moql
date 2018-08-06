package generate

import (
	"github.com/LukeJoeDavis/moql/domain"
	"math/rand"
	"time"
	"strings"
	)

func CreateInserts(table domain.DiscoveredTable) []string {

	insertStatements := make([]string, 0)

	insertStatements = append(insertStatements, CreateStatement(table, "min"))
	insertStatements = append(insertStatements, CreateStatement(table, "max"))

	return insertStatements
}

func CreateStatement(table domain.DiscoveredTable, valueType string) string {

	columnStatment := "INSERT INTO " + table.Name + "("
	valueStatement := ") VALUES ("

	for i, column := range table.Columns {
		columnName, value := CreateValue(column, valueType)
		columnStatment += columnName
		valueStatement += "\"" + value + "\""

		if i<len(table.Columns)-1{
			columnStatment += ","
			valueStatement += ","
		}
	}

	return columnStatment + valueStatement + ");"
}

func CreateValue(column domain.DiscoveredColumn, valueType string) (string, string) {

	valueToInsert := ""
	if strings.Contains(column.Type, "enum"){
		return column.Field, column.Enums[0]
	}

	switch valueType {
	case "min":
		valueToInsert = GenerateMinValue(column)
		break
	case "max":
		valueToInsert = GenerateMaxValue(column)
		break
	}

	return column.Field, valueToInsert
}

func GenerateMaxValue(column domain.DiscoveredColumn) string {
	if strings.Contains(column.Type, "char"){
		return RandStringRunes(column.MaxCharacterLimit)
	} else if strings.Contains(column.Type, "tinytext") {
		return RandStringRunes(255)
	} else if strings.Contains(column.Type, "mediumtext") {
		return RandStringRunes(16777215)
	} else if strings.Contains(column.Type, "mediumblob") {
		return RandStringRunes(16777215)
	}  else if strings.Contains(column.Type, "longtext") {
		return RandStringRunes(4294967295)
	}  else if strings.Contains(column.Type, "longblob") {
		return RandStringRunes(4294967295)
	} else if strings.Contains(column.Type, "test") || strings.Contains(column.Type, "blob") {
		return RandStringRunes(65535)
	} else if strings.Contains(column.Type, "tinyint") {
		return "127"
	} else if strings.Contains(column.Type, "smallint") {
		return "32767"
	} else if strings.Contains(column.Type, "mediumint") {
		return "8388607"
	} else if strings.Contains(column.Type, "bigint") {
		return "9223372036854775807"
	} else if strings.Contains(column.Type, "int") {
		return "2147483647"
	} else if strings.Contains(column.Type, "datetime") {
		return "9999-12-31 23:59:59"
	} else if strings.Contains(column.Type, "date") {
		return "9999-12-31"
	} else if strings.Contains(column.Type, "timestamp") {
		return "2038-01-09 03:14:07"
	} else if strings.Contains(column.Type, "time") {
		return "838:59:59"
	} else if strings.Contains(column.Type, "year") {
		return "2069"
	} else if strings.Contains(column.Type, "bit") {
		return "true"
	} else if strings.Contains(column.Type, "float") || strings.Contains(column.Type, "double") || strings.Contains(column.Type, "decimal"){
		return RandNumberSize(column.MaxDigits-column.MaxDecimals) + "." + RandNumberSize(column.MaxDecimals)
	} else {
		return ""
	}
}

func GenerateMinValue(column domain.DiscoveredColumn) string {
	if strings.Contains(column.Type, "char"){
		return RandStringRunes(0)
	} else if strings.Contains(column.Type, "tinytext") {
		return RandStringRunes(0)
	} else if strings.Contains(column.Type, "mediumtext") {
		return RandStringRunes(0)
	} else if strings.Contains(column.Type, "mediumblob") {
		return RandStringRunes(0)
	}  else if strings.Contains(column.Type, "longtext") {
		return RandStringRunes(0)
	}  else if strings.Contains(column.Type, "longblob") {
		return RandStringRunes(0)
	} else if strings.Contains(column.Type, "text") || strings.Contains(column.Type, "blob") {
		return RandStringRunes(0)
	} else if strings.Contains(column.Type, "tinyint") {
		return "-128"
	} else if strings.Contains(column.Type, "smallint") {
		return "-32768"
	} else if strings.Contains(column.Type, "mediumint") {
		return "-8388608"
	} else if strings.Contains(column.Type, "bigint") {
		return "-9223372036854775808"
	} else if strings.Contains(column.Type, "int") {
		return "-2147483648"
	} else if strings.Contains(column.Type, "datetime") {
		return "1000-01-01 00:00:00"
	} else if strings.Contains(column.Type, "date") {
		return "9999-12-31"
	} else if strings.Contains(column.Type, "timestamp") {
		return "1970-01-01 00:00:00"
	} else if strings.Contains(column.Type, "time") {
		return "-838:59:59"
	} else if strings.Contains(column.Type, "year") {
		return "1901"
	} else if strings.Contains(column.Type, "bit") {
		return "false"
	} else if strings.Contains(column.Type, "float") || strings.Contains(column.Type, "double") || strings.Contains(column.Type, "decimal"){
		return "0"
	} else {
		return ""
	}

}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numberRunes = []rune("1234567890")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandNumberSize(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = numberRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}


func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}