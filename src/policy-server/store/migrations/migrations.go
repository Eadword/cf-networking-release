package migrations

var MigrationsToPerform policyServerMigrations = policyServerMigrations{
	policyServerMigration{
		"1",
		migration_v0001,
	},
	policyServerMigration{
		"2",
		migration_v0002,
	},
	policyServerMigration{
		"3",
		migration_v0003,
	},
	policyServerMigration{
		"4",
		migration_v0004,
	},
}
