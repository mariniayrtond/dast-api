package config

import "time"

const (
	TTLForToken = 12 * time.Hour
	// TODO: esto probablemente tenga que ser una struct
	MongoDBUserName       = "amarini"
	MongoDBPass           = "9W00BNdHyOpN97xv"
	MongoDBName           = "dast"
	MongoTableHierarchies = "hierarchies"
	MongoTableJudgements  = "judgements"
	MongoTableUsers       = "users"
	MongoTableTemplates   = "templates"
	MongoTableTokens      = "tokens"
)
