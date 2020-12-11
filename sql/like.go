package sql

import (
	"strings"
)

// WildcardPrefix inserts wildcard as prefix.
func WildcardPrefix(val string) string {
	return "%" + EscapeWildcard(val)
}

// WildcardSuffix inserts wildcard as suffix.
func WildcardSuffix(val string) string {
	return EscapeWildcard(val) + "%"
}

// WildcardSurround surrounds the value with wildcards.
func WildcardSurround(val string) string {
	return "%" + EscapeWildcard(val) + "%"
}

// EscapeWildcard escapes all wildcards in value to avoid injection.
func EscapeWildcard(val string) (ret string) {
	ret = val
	ret = strings.ReplaceAll(ret, "%", "\\%")
	ret = strings.ReplaceAll(ret, "_", "\\_")
	return
}
