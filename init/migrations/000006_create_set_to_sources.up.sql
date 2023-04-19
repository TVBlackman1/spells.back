CREATE TABLE set_to_sources(
    id UUID NOT NULL PRIMARY KEY,
    set_id UUID NOT NULL,
    source_id UUID NOT NULL
);

CREATE TRIGGER update_spells_edit_time BEFORE UPDATE
ON set_to_sources FOR EACH ROW EXECUTE PROCEDURE 
update_edit_time();