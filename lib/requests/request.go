package requests

import (
	"fmt"
	"strings"
)

type Request struct {
	tableName             string
	conditions            strings.Builder
	selects               strings.Builder
	leftJoins             strings.Builder
	request               strings.Builder
	alreadyAddedCondition bool // TODO refactor: remove property
	limit                 int
	offset                int
	orderBy               strings.Builder
}

func NewRequest(tableName string) *Request {
	return &Request{
		tableName:             tableName,
		conditions:            strings.Builder{},
		selects:               strings.Builder{},
		leftJoins:             strings.Builder{},
		request:               strings.Builder{},
		alreadyAddedCondition: false,
		limit:                 0,
		offset:                0,
		orderBy:               strings.Builder{},
	}
}

func (req *Request) Where(condition string) *Request {
	req.addPreCondition()
	req.conditions.WriteRune('(')
	req.conditions.WriteString(condition)
	req.conditions.WriteRune(')')
	return req
}

func (req *Request) Select(selects string) *Request {
	selects = strings.TrimSpace(selects)
	selects = strings.Trim(selects, ",")

	if req.selects.Len() > 0 {
		req.selects.WriteString(", ")
	}
	req.selects.WriteString(selects)
	return req
}

func (req *Request) LeftJoin(tableName string, mapping string) *Request {
	if req.leftJoins.Len() > 0 {
		req.leftJoins.WriteRune(' ')
	}
	currentLeftJoin := fmt.Sprintf("left join %s on %s", tableName, mapping)
	req.leftJoins.WriteString(currentLeftJoin)
	return req
}

func (req *Request) Limit(limit int) *Request {
	req.limit = limit
	return req
}

func (req *Request) Offset(offset int) *Request {
	req.offset = offset
	return req
}

// TODO add tests for library for examples: snippets of code
// OrderBy("table.field")
// OrderBy("table.field desc)
// OrderBy("table.field desc, table.field2)
// OrderBy("table.field").OrderBy("table.field desc")
func (req *Request) OrderBy(orderBy string) *Request {
	orderBy = strings.TrimSpace(orderBy)
	orderBy = strings.Trim(orderBy, ",")

	if req.orderBy.Len() > 0 {
		req.orderBy.WriteString(", ")
	}
	req.orderBy.WriteString(orderBy)
	return req
}

func (req *Request) addPreCondition() {
	if !req.alreadyAddedCondition {
		req.alreadyAddedCondition = true
	} else {
		req.conditions.WriteString(" and ")
	}
}

// TODO refactor method: use :=
func (req *Request) Clone() *Request {
	return &Request{
		tableName:             req.tableName,
		conditions:            req.conditions,
		selects:               req.selects,
		leftJoins:             req.leftJoins,
		request:               req.request,
		alreadyAddedCondition: req.alreadyAddedCondition,
		limit:                 req.limit,
		offset:                req.offset,
		orderBy:               req.orderBy,
	}
}

func (req *Request) String() string {
	req.request.Reset()
	req.request.WriteString("select ")
	req.request.WriteString(req.selects.String())
	req.request.WriteString(" from ")
	req.request.WriteString(req.tableName)
	if req.leftJoins.Len() > 0 {
		req.request.WriteRune(' ')
		req.request.WriteString(req.leftJoins.String())
	}
	if req.conditions.Len() > 0 {
		req.request.WriteString(" where ")
		req.request.WriteString(req.conditions.String())
	}
	if req.orderBy.Len() > 0 {
		req.request.WriteString(" order by ")
		req.request.WriteString(req.orderBy.String())
	}
	if req.offset > 0 {
		req.request.WriteString(fmt.Sprintf(" offset %d ", req.offset))
	}
	if req.limit > 0 {
		req.request.WriteString(fmt.Sprintf(" limit %d", req.limit))
	}
	return req.request.String()
}
