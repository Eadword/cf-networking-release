package migrations

var migration_v0004 = map[string][]string{
	"mysql": {
		`CREATE TABLE IF NOT EXISTS ip_ranges (
		id int NOT NULL AUTO_INCREMENT,
		group_id int REFERENCES groups(id),
		start_ip varchar(255) DEFAULT '',
		end_ip varchar(255) DEFAULT '',
		UNIQUE (group_id)
		);`,
	},

	"postgres": {
		`CREATE TABLE IF NOT EXISTS ip_ranges (
		id SERIAL PRIMARY KEY,
		group_id int REFERENCES groups(id),
		start_ip text DEFAULT '',
		end_ip text DEFAULT '',
		UNIQUE (group_id)
		);`,
	},
}