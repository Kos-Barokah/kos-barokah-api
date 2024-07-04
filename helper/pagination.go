package helper

import (
	"math"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func PaginateQuery(qry *gorm.DB, page, pageSize uint) (*gorm.DB, uint, uint, error) {
	var totalPage int64
	var totalItems int64

	if page > 0 && pageSize > 0 {
		offset := (page - 1) * pageSize

		if err := qry.Count(&totalItems).Error; err != nil {
			return nil, 0, 0, err
		}

		totalPage = int64(math.Ceil(float64(totalItems) / float64(pageSize)))

		qry = qry.Offset(int(offset)).Limit(int(pageSize))
	}

	convertPage := uint(totalPage)
	convertItems := uint(totalItems)

	return qry, convertPage, convertItems, nil
}

func GetPaginationQuery(c echo.Context) (*string, *uint, *uint, error) {
	search := c.QueryParam("search")

	page, _ := strconv.Atoi(c.QueryParam("page"))

	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))

	if search == "" {
		search = ""
	}

	if page <= 0 {
		page = 0
	}

	if pageSize <= 0 {
		pageSize = 0
	}

	convertPage := uint(page)
	convertPageSize := uint(pageSize)

	return &search, &convertPage, &convertPageSize, nil
}
