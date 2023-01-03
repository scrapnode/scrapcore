package database

type Configs struct {
	Dsn        string   `json:"dsn"`
	MigrateDir string   `json:"migrate_dir"`
	SeedFiles  []string `json:"seed_files"`
}
