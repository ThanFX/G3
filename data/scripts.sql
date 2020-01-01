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
	id SERIAL PRIMARY KEY,
	name TEXT,
	short_name TEXT,
	min INTEGER,
	max INTEGER
);

CREATE TABLE mastery_items (
	id SERIAL PRIMARY KEY,
	mastery TEXT,
	category TEXT,
	ingredient TEXT,
	name TEXT,
	rarity INTEGER,
	areas TEXT,
	area_size INTEGER,
	min INTEGER,
	max INTEGER,
	is_countable BOOLEAN,
	is_liquid BOOLEAN,
	"limit" INTEGER
);

CREATE TABLE "map" (
	id TEXT NOT NULL UNIQUE,
	chunk TEXT
);


CREATE TABLE persons (
	id SERIAL PRIMARY KEY,
	name TEXT,
	age INTEGER,
	is_male BOOLEAN,
	chunk_id TEXT,
	day_action TEXT
);

CREATE TABLE person_masterships (
	person_id INTEGER NOT NULL,
	mastery_id INTEGER NOT NULL,
	skill TEXT
);
ALTER TABLE public.person_masterships ADD CONSTRAINT person_masterships_fk FOREIGN KEY (person_id) REFERENCES public.persons(id);
ALTER TABLE public.person_masterships ADD CONSTRAINT person_masterships_fk_1 FOREIGN KEY (mastery_id) REFERENCES public.masterships(id);


CREATE TABLE person_inventory (
	id TEXT NOT NULL UNIQUE,
	person_id INTEGER NOT NULL,
	item_id INTEGER NOT NULL,
	weight TEXT,
	quality TEXT,
	creation_date INTEGER,
	exp_date INTEGER,
	is_deleted BOOLEAN
);
ALTER TABLE public.person_inventory ADD CONSTRAINT person_inventory_fk FOREIGN KEY (person_id) REFERENCES public.persons(id);
ALTER TABLE public.person_inventory ADD CONSTRAINT person_inventory_fk_1 FOREIGN KEY (item_id) REFERENCES public.mastery_items(id);


CREATE TABLE params (
	"key" TEXT UNIQUE,
	value TEXT
);


select * from map;
select * from masterships;

UPDATE params SET value=9842 WHERE key='date';
