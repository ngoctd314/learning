package gsql

import (
	"strings"

	"gorm.io/gorm"
)

type Cond struct {
	selectedField []string
	limit         *int
	offset        *int
	query         string
	params        []any
	operator      string
}

func (c *Cond) applyTx(tx *gorm.DB) *gorm.DB {
	if c.selectedField != nil {
		tx = tx.Select(c.selectedField)
	}
	if notEmptyString(c.query) && c.params != nil {
		tx = tx.Where(c.query, c.params...)
	}
	if c.limit != nil {
		tx = tx.Limit(*c.limit)
	}
	if c.offset != nil {
		tx = tx.Offset(*c.offset)
	}

	return tx
}

func (c *Cond) SetSelectedField(selectedField ...string) *Cond {
	if selectedField != nil {
		c.selectedField = selectedField
	}
	return c
}

func (c *Cond) SetLimit(limit int) {
	c.limit = &limit
}

func (c *Cond) SetOffset(offset int) {
	c.offset = &offset
}

func (condition *Cond) not() *Cond {
	sb := strings.Builder{}
	sb.WriteString("NOT(")
	sb.WriteString(condition.query)
	sb.WriteString(")")

	condition.query = sb.String()
	condition.operator = ""
	return condition
}

func (condition *Cond) appendCondition(appendOperator string, conditions ...*Cond) *Cond {
	// wrap previous condition in () if needed
	builder := strings.Builder{}
	if condition.operator != appendOperator && !isEmptyString(condition.operator) &&
		!isEmptyString(condition.query) {
		condition.query = wrapQuery(condition.query)
	}
	builder.WriteString(condition.query)
	condition.operator = appendOperator
	for _, c := range conditions {
		if isEmptyString(c.query) {
			continue
		}
		if c.operator != appendOperator && !isEmptyString(c.operator) {
			c.query = wrapQuery(c.query)
		}
		if builder.Len() > 0 {
			builder.WriteString(appendOperator)
		}
		builder.WriteString(c.query)
		condition.params = append(condition.params, c.params...)
	}
	condition.query = builder.String()
	return condition
}

func wrapQuery(query string) string {
	sb := strings.Builder{}
	sb.WriteString("(")
	sb.WriteString(query)
	sb.WriteString(")")

	return sb.String()
}

// Equal represents "field = value".
func Equal(field string, value interface{}) *Cond {
	sb := strings.Builder{}
	sb.WriteString(field)
	sb.WriteString(" = ?")

	return &Cond{
		query:  sb.String(),
		params: []any{value},
	}
}

// NotEqual represents "field <> value".
func NotEqual(field string, value any) *Cond {
	sb := strings.Builder{}
	sb.WriteString(field)
	sb.WriteString(" <> ?")

	return &Cond{
		query:  sb.String(),
		params: []any{value},
	}
}

// GreaterThan represents "field > value".
func GreaterThan(field string, value interface{}) *Cond {
	sb := strings.Builder{}
	sb.WriteString(field)
	sb.WriteString(" > ?")

	return &Cond{
		query:  sb.String(),
		params: []any{value},
	}
}

// GreaterEqualThan represents "field >= value".
func GreaterEqualThan(field string, value interface{}) *Cond {
	sb := strings.Builder{}
	sb.WriteString(field)
	sb.WriteString(" >= ?")

	return &Cond{
		query:  sb.String(),
		params: []any{value},
	}
}

// LessThan represents "field < value".
func LessThan(field string, value interface{}) *Cond {
	sb := strings.Builder{}
	sb.WriteString(field)
	sb.WriteString(" < ?")

	return &Cond{
		query:  sb.String(),
		params: []any{value},
	}
}

// LessEqualThan represents "field <= value".
func LessEqualThan(field string, value interface{}) *Cond {
	sb := strings.Builder{}
	sb.WriteString(field)
	sb.WriteString(" <= ?")

	return &Cond{
		query:  sb.String(),
		params: []interface{}{value},
	}
}

// In represents "field IN (value...)".
func In(field string, values ...interface{}) *Cond {
	sb := strings.Builder{}
	sb.WriteString(field)
	sb.WriteString(" IN (?)")

	return &Cond{
		query:  sb.String(),
		params: []any{values},
	}
}

// NotIn represents "field NOT IN (value...)".
func NotIn(field string, values ...interface{}) *Cond {
	sb := strings.Builder{}
	sb.WriteString(field)
	sb.WriteString(" NOT IN (?)")

	return &Cond{
		query:  sb.String(),
		params: []any{values},
	}
}

// Like represents "field LIKE value".
func Like(field string, value string) *Cond {
	sb := strings.Builder{}
	sb.WriteString(field)
	sb.WriteString(" LIKE ?")

	return &Cond{
		query:  sb.String(),
		params: []any{value},
	}
}

// NotLike represents "field NOT LIKE value".
func NotLike(field string, value string) *Cond {
	sb := strings.Builder{}
	sb.WriteString(field)
	sb.WriteString(" NOT LIKE ?")

	return &Cond{
		query:  sb.String(),
		params: []any{value},
	}
}

// IsNull represents "field IS NULL".
func IsNull(field string, value string) *Cond {
	sb := strings.Builder{}
	sb.WriteString(field)
	sb.WriteString(" IS NULL")

	return &Cond{
		query:  sb.String(),
		params: []any{value},
	}
}

// IsNotNull represents "field IS NOT NULL".
func IsNotNull(field string, value string) *Cond {
	sb := strings.Builder{}
	sb.WriteString(field)
	sb.WriteString(" IS NOT NULL")

	return &Cond{
		query:  sb.String(),
		params: []any{value},
	}
}

// Between represents "field BETWEEN lower AND upper".
func Between(field string, lower, upper string) *Cond {
	sb := strings.Builder{}
	sb.WriteString(field)
	sb.WriteString(" BETWEEN ? AND ?")

	return &Cond{
		query:  sb.String(),
		params: []any{lower, upper},
	}
}

// NotBetween represents "field NOT BETWEEN lower AND upper".
func NotBetween(field string, lower, upper string) *Cond {
	sb := strings.Builder{}
	sb.WriteString(field)
	sb.WriteString("NOT BETWEEN ? AND ?")

	return &Cond{
		query:  sb.String(),
		params: []any{lower, upper},
	}
}

// And will Join simple a slice of condition into a condition with AND WiseFunc for where statement
func And(conditions ...*Cond) *Cond {
	object := &Cond{}
	return object.appendCondition(" AND ", conditions...)
}

// Raw represents and raw query
func Raw(query string, params []interface{}) *Cond {
	object := &Cond{
		query:  query,
		params: params,
	}
	return object
}

// Or will Join simple a slice of condition into a condition with OR WiseFunc for where statement
func Or(conditions ...*Cond) *Cond {
	object := &Cond{}
	return object.appendCondition(" OR ", conditions...)
}

func Not(condition *Cond) *Cond {
	object := &Cond{
		query:    condition.query,
		params:   condition.params,
		operator: condition.operator,
	}
	return object.not()
}

func isEmptyString(s string) bool {
	return strings.TrimSpace(s) == ""
}
