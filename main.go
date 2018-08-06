package main

import (
		"github.com/LukeJoeDavis/moql/discovery"
		"github.com/LukeJoeDavis/moql/generate"
	"fmt"
	)

func main() {
	tables := discover.GetTables()

	inserts := make([]string, 0)
	for _, table := range tables{
		func (tempTable string){
			fmt.Println(tempTable + " starting")
			discoveredTable := discover.GetColumns(tempTable)
			discoveredTable.Name = tempTable
			inserts = append(inserts, generate.CreateInserts(discoveredTable)...)
			fmt.Println(tempTable + " complete")
		}(table)
	}
}
