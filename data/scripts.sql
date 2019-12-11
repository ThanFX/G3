-- SQLite
CREATE TABLE objects (
	id INTEGER NOT NULL,
	"level" INTEGER NOT NULL,
	parent_id INTEGER,
	"name" TEXT,
	is_countable INTEGER NOT NULL,
	is_liquid INTEGER NOT NULL,
	"limit" INTEGER
);

CREATE TABLE masterships (
	id INTEGER NOT NULL,
	name TEXT,
	short_name TEXT,
	min INTEGER,
	max INTEGER
);

select * from map;
select * from masterships;
