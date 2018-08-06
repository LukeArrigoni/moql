package discover

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
		"github.com/LukeJoeDavis/moql/domain"
	"regexp"
	"strings"
	"strconv"
	"github.com/LukeJoeDavis/moql/data"
	"log"
	)

func GetTables() []string {

	db := data.GetDataConnection()
	defer db.Close()

	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		panic(err.Error())
	}

	tables := make([]string, 0)

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		tables = append(tables, name)
	}
	return tables
}

func GetColumns(tableName string) domain.DiscoveredTable {

	db := data.GetDataConnection()
	defer db.Close()

	rows, err := db.Query("DESCRIBE " + tableName)
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	table := domain.DiscoveredTable{}

	// Fetch rows
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var value string
		column := new(domain.DiscoveredColumn)
		findLimit := regexp.MustCompile("\\(([^()]+)\\)")


		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}

			switch columns[i] {
			case "Field":
				column.Field = value
				break
			case "Type":
				column.Type = value
				limit := findLimit.FindAllString(value, 1)
				if limit != nil {
					limitation := ParseLimitation(limit)

					if strings.Contains(value, "enum"){
						column.Enums = ParseEnums(limitation)
					} else if strings.Contains(value, "double"){
						doubleLimitations := strings.Split(limitation, ",")
						column.MaxDigits, err = strconv.Atoi(doubleLimitations[0])
						if err != nil {
							panic(err.Error())
						}

						column.MaxDecimals, err = strconv.Atoi(doubleLimitations[1])
						if err != nil {
							panic(err.Error())
						}
					} else if strings.Contains(value, "bit"){
						//do nothing
					} else {
						column.MaxCharacterLimit, err = strconv.Atoi(limitation)
						if err != nil {
							panic(err.Error())
						}
					}
				}

				break
			case "Null":
				column.Null = false
				if value == "YES" {
					column.Null = true
				}
				break
			case "Key":
				column.Key = value
				break
			case "Default":
				column.Default = value
				break
			case "Extra":
				column.Extra = value
				break
			}
		}
		table.Columns = append(table.Columns, *column)

	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
	return table
}

func ParseLimitation(limit []string) string {
	return strings.Replace(strings.Replace(limit[0], "(", "", -1), ")", "", -1)
}

func ParseEnums(enums string) []string {
	return strings.Split(strings.Replace(enums, "'", "", -1), ",")
}
