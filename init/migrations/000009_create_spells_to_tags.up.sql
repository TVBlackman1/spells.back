CREATE TABLE spells_to_tags(
    id UUID NOT NULL PRIMARY KEY,
    spell_in_set_id UUID NOT NULL,
    tag_id UUID NOT NULL
);

CREATE TRIGGER update_spells_edit_time BEFORE UPDATE
ON spells_to_tags FOR EACH ROW EXECUTE PROCEDURE 
update_edit_time();