func displayTable(orders [][]string) [][]string {

	tableFood := make(map[int]map[string]int)
	tablesSet := make(map[int]struct{})
	foodsSet := make(map[string]struct{})

	for _, order := range orders {
		tableNum, _ := strconv.Atoi(order[1])
		food := order[2]

		if _, ok := tableFood[tableNum]; !ok {
			tableFood[tableNum] = make(map[string]int)
		} 
		tableFood[tableNum][food]++
		

		tablesSet[tableNum] = struct{}{}
		foodsSet[food] = struct{}{}
	}

	tables := make([]int, 0, len(tablesSet))
	for table := range tablesSet {
		tables = append(tables, table)
	}
	sort.Ints(tables)

	foods := make([]string, 0, len(foodsSet))
	for food := range foodsSet {
		foods = append(foods, food)
	}
	sort.Strings(foods)

	result := make([][]string, 0, len(tables))

	header := []string{"Table"}
	header = append(header, foods...)
	result = append(result, header)

	for _, table := range tables {
		row := []string{strconv.Itoa(table)}

		for _, food := range foods {
			count := tableFood[table][food]
			row = append(row, strconv.Itoa(count))
		}

		result = append(result, row)
	}

	return result
}