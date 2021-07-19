package rest

import (
	helpers "github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/bools"
	"regexp"
	"strconv"
	"strings"
)



type Args struct {
	Sort   			string
	Order  			string
	Offset 			string
	Limit  			string
	Search 			string
	WithChildren 	string
}

var ArgsType = struct {
	Sort   				string
	Order   			string
	Offset   			string
	Limit   			string
	Search   			string
	WithChildren 		string
}{
	Sort:   			"Sort",
	Order:   			"Order",
	Offset:   			"Offset",
	Limit:   			"Limit",
	Search:   			"Search",
	WithChildren:   	"WithChildren",
}

var ArgsDefault = struct {
	Sort   			string
	Order   		string
	Offset   		string
	Limit   		string
	Search   		string
	WithChildren   	string
}{
	Sort:   			"ID",
	Order:   			"DESC",
	Offset:   			"0",
	Limit:   			"25",
	Search:   			"",
	WithChildren:   	"false",
}


// Offset returns the starting number of result for pagination
func Offset(offset string) int {
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		offsetInt = 0
	}
	return offsetInt
}

// Limit returns the number of result for pagination
func Limit(limit string) int {
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 25
	}
	return limitInt
}

// SortOrder returns the string for sorting and ordering data
func SortOrder(table, sort, order string) string {
	return table + "." + ToSnakeCase(sort) + " " + ToSnakeCase(order)
}


func WithChildren(value string) (bool, error)  {
	if value == "" {
		return false, nil
	} else  {
		c, err := helpers.Boolean(value)
		return c, err
	}
}


// ToSnakeCase changes string to database table
func ToSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")

	return strings.ToLower(snake)
}
