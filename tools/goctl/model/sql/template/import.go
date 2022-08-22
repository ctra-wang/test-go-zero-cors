package template

const (
	// Imports defines a import template for model in cache case
	Imports = `import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	"github.com/Holyson/test-go-zero-cors/core/stores/builder"
	"github.com/Holyson/test-go-zero-cors/core/stores/cache"
	"github.com/Holyson/test-go-zero-cors/core/stores/sqlc"
	"github.com/Holyson/test-go-zero-cors/core/stores/sqlx"
	"github.com/Holyson/test-go-zero-cors/core/stringx"
)
`
	// ImportsNoCache defines a import template for model in normal case
	ImportsNoCache = `import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	"github.com/Holyson/test-go-zero-cors/core/stores/builder"
	"github.com/Holyson/test-go-zero-cors/core/stores/sqlc"
	"github.com/Holyson/test-go-zero-cors/core/stores/sqlx"
	"github.com/Holyson/test-go-zero-cors/core/stringx"
)
`
)
