-- SQLite
select * from fishes;

CREATE TABLE objects (
	id INTEGER NOT NULL,
	"level" INTEGER NOT NULL,
	parent_id INTEGER,
	"name" TEXT,
	is_countable INTEGER NOT NULL,
	is_liquid INTEGER NOT NULL,
	"limit" INTEGER
);