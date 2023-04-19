CREATE TABLE spells_in_set(
    id UUID NOT NULL PRIMARY KEY,
    original_spell UUID NOT NULL,
    set UUID NOT NULL,
    tags UUID NOT NULL,
    master_comment text,
    visual_style_comment text,
    created timestamp DEFAULT CURRENT_TIMESTAMP,
    edited timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_spells_edit_time BEFORE UPDATE
ON spells_in_set FOR EACH ROW EXECUTE PROCEDURE
update_edit_time();