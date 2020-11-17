package config

import "time"

const (
	TTLForToken           = 12 * time.Hour
	MongoDBName           = "dast"
	MongoTableHierarchies = "hierarchies"
	MongoTableJudgements  = "judgements"
	MongoTableUsers       = "users"
	MongoTableTemplates   = "templates"
	MongoTableTokens      = "tokens"
)
