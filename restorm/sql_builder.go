package restorm

import (
	"fmt"
	"strings"
)

type sqlBuilder struct {
	hint       string
	table      string
	forceIndex string
	fields     string
	where      string
	groupBy    string
	orderBy    string
	limit      string
	offset     string
	args       []interface{}
}

func (s *sqlBuilder) limitFormat() string {
	if s.limit != "" {
		return "LIMIT " + s.limit
	}
	return ""
}

func (s *sqlBuilder) offsetFormat() string {
	if s.offset != "" {
		return "OFFSET " + s.offset
	}
	return ""
}

func (s *sqlBuilder) orderFormat() string {
	if s.orderBy != "" {
		return "ORDER BY " + s.orderBy
	}
	return ""
}

func (s *sqlBuilder) _table() string {
	table := "`" + s.table + "`"
	if s.forceIndex != "" {
		table += fmt.Sprintf(" force index(%s)", s.forceIndex)
	}
	return table
}

func (s *sqlBuilder) groupFormat() string {
	if s.groupBy != "" {
		return "GROUP BY " + s.groupBy
	}
	return ""
}

//queryString Assemble the query statement
func (s *sqlBuilder) queryString() string {
	if s.fields == "" {
		s.fields = "*"
	}
	return fmt.Sprintf("%sSELECT %s FROM %s %s %s %s %s %s;", s.hint, s.fields, s._table(), s.where, s.groupFormat(), s.orderFormat(), s.limitFormat(), s.offsetFormat())
}

//countString Assemble the count statement
func (s *sqlBuilder) countString() string {
	return fmt.Sprintf("%sSELECT count(*) FROM %s %s;", s.hint, s._table(), s.where)
}

//insertString Assemble the insert statement
func (s *sqlBuilder) insertString(params map[string]interface{}) string {
	var cols, vls []string
	for k, v := range params {
		cols = append(cols, fmt.Sprintf("`%s`", k))
		vls = append(vls, "?")
		s.args = append(s.args, v)
	}
	return fmt.Sprintf("INSERT INTO %s (%s) VALUES(%s);", s._table(), strings.Join(cols, ","), strings.Join(vls, ","))
}

//updateString Assemble the update statement
func (s *sqlBuilder) updateString(params map[string]interface{}) string {
	var updateFields []string
	args := make([]interface{}, 0)

	for k := range params {
		updateFields = append(updateFields, fmt.Sprintf("%s=?", fmt.Sprintf("`%s`", k)))
		args = append(args, params[k])
	}
	s.args = append(args, s.args...)
	return fmt.Sprintf("UPDATE %s SET %s %s;", s._table(), strings.Join(updateFields, ","), s.where)
}

//deleteString Assemble the delete statement
func (s *sqlBuilder) deleteString() string {
	return fmt.Sprintf("DELETE FROM %s %s;", s._table(), s.where)
}

func (s *sqlBuilder) Where(filter ...interface{}) {
	str := ""
	var args []interface{}
	_l := len(filter)
	if _l > 0 {
		str = filter[0].(string)
		if _l > 1 {
			args = filter[1:]
		}
	}

	if str == "" {
		return
	}

	if s.where != "" {
		s.where = fmt.Sprintf("%s AND (%s)", s.where, str)
	} else {
		s.where = fmt.Sprintf("WHERE (%s)", str)
	}

	if args == nil || len(args) == 0 {
		return
	}

	if s.args == nil {
		s.args = args
	} else {
		s.args = append(s.args, args...)
	}
}
