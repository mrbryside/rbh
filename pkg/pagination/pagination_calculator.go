package pagination

func CalculateOffset(page int, pageSize int) int {
	return (page - 1) * pageSize
}

func IsNextPage(totalRecords int64, page int, pageSize int) bool {
	return totalRecords > int64(page*pageSize)
}
