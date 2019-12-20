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

CREATE TABLE persons (
	id INTEGER NOT NULL,
	name TEXT,
	age INTEGER,
	is_male INTEGER,
	chunk_id TEXT,
	day_action TEXT
);

CREATE TABLE person_masterships (
	person_id INTEGER NOT NULL,
	mastery_id INTEGER NOT NULL,
	skill TEXT
);

CREATE TABLE person_inventory (
	id TEXT NOT NULL,
	person_id INTEGER NOT NULL,
	item_id INTEGER NOT NULL,
	weight TEXT,
	quality TEXT,
	creation_date INTEGER,
	exp_date INTEGER,
	is_deleted INTEGER
);

CREATE TABLE params (
	"key" TEXT,
	value TEXT
);


select * from map;
select * from masterships;
