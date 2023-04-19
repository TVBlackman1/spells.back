CREATE TABLE sets(
    id UUID NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    user_id UUID NOT NULL,
    description text,
    set_to_sources UUID NOT NULL,
    created timestamp DEFAULT CURRENT_TIMESTAMP,
    edited timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_spells_edit_time BEFORE UPDATE
ON sets FOR EACH ROW EXECUTE PROCEDURE
update_edit_time();