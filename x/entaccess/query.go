package entaccess

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	iam "go.deployport.com/iam-sdk/client"
)

// ApplyQueryAssert returns a function that can be used to filter SQL queries based on IAM assertions.
func ApplyQueryAssert(entries []*iam.InternalQueryAssertionEntry) func(sel *sql.Selector) {
	return func(sel *sql.Selector) {
		sel.Where(sql.P().Append(func(builder *sql.Builder) {
			for _, entry := range entries {
				if snippet := entry.Snippet; snippet != nil {
					builder.WriteString(*snippet)
				} else if v := entry.Value; v != nil {
					if b := v.B; b != nil {
						builder.Arg(*b)
					} else if s := v.S; s != nil {
						builder.Arg(*s)
					} else {
						builder.AddError(fmt.Errorf("unsupported value type in assertion entry: %#v", v))
					}
				}
			}
		}))
	}
}
