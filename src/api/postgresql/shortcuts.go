package postgresql

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

func In(field string, identifiers []uuid.UUID) (string, []interface{}) {
	queries := make([]interface{}, len(identifiers))
	for i, id := range identifiers {
		queries[i] = id.String()
	}

	return fmt.Sprintf("%s IN (%s)", field, strings.Repeat("?", len(identifiers))), queries
}
