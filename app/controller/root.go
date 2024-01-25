package controller

import (
	"strconv"
	"strings"

	"github.com/CSCI-X050-A7/backend/app/schema"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

var db *gorm.DB

func NewDB(newDB *gorm.DB) {
	db = newDB
}

func ListObjs[Schema any](
	statement *gorm.DB,
	pagination schema.Pagination,
) (schemaValue []Schema, count int64, err error) {
	if err := statement.Count(&count).Error; err != nil {
		return nil, 0, errors.Wrapf(err, "failed get %T count", schemaValue)
	}
	err = statement.Scopes(Paginate(pagination)).Scan(&schemaValue).Error
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}
	return schemaValue, count, nil
}

func GetPagination(c *fiber.Ctx) schema.Pagination {
	ps := c.Query("limit")
	pn := c.Query("offset")
	limit, offset := 10, 0

	if ps != "" {
		psInt, err := strconv.Atoi(ps)
		if err != nil {
			logrus.Error(err)
		} else {
			limit = psInt
		}
	}

	if pn != "" {
		pnInt, err := strconv.Atoi(pn)
		if err != nil {
			logrus.Error(err)
		} else {
			offset = pnInt
		}
	}

	return schema.Pagination{Ordering: nil, Limit: &limit, Offset: &offset}
}

func Paginate(pagination schema.Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pagination.Offset == nil || *pagination.Offset < 0 {
			pagination.Offset = schema.Pointer(0)
		}
		if pagination.Limit == nil || *pagination.Limit > 100 {
			pagination.Limit = schema.Pointer(100)
		}
		if pagination.Ordering == nil {
			pagination.Ordering = schema.Pointer("-created_at")
		}
		statement := db.Offset(*pagination.Offset).Limit(*pagination.Limit)
		for _, item := range strings.Split(*pagination.Ordering, ",") {
			name := item
			desc := false
			if strings.HasPrefix(name, "-") {
				name = strings.TrimLeft(item, "-")
				desc = true
			}
			if name != "" {
				statement = statement.Order(clause.OrderByColumn{
					Column: clause.Column{Name: name},
					Desc:   desc,
				})
			}
		}
		return statement
	}
}
