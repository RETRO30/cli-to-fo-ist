package builder

import (
	"cli-tool/src/database"
)

type Builder struct {
	Title string
	menu
	database database.Database
	