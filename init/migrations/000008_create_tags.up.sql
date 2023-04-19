CREATE TABLE tags(
    id UUID NOT NULL PRIMARY KEY,
    name VARCHAR(50),
    is_global BOOLEAN DEFAULT false,
    set UUID,
    description text,
    color VARCHAR(7),
    created timestamp DEFAULT CURRENT_TIMESTAMP,
    edited timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_spells_edit_time BEFORE UPDATE
ON tags FOR EACH ROW EXECUTE PROCEDURE
update_edit_time();