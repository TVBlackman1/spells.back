CREATE TABLE spells(
    id UUID NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    level int NOT NULL,
    classes VARCHAR(255),
    version int DEFAULT 1,
    description text,
    action VARCHAR(255) NOT NULL,
    duration VARCHAR(255) NOT NULL,
    is_verbal BOOLEAN DEFAULT false,
    is_somatic BOOLEAN DEFAULT false,
    is_material BOOLEAN DEFAULT false,
    material_content VARCHAR(255) NOT NULL DEFAULT '',
    magic_school VARCHAR(40),
    distance VARCHAR(40),
    is_ritual BOOLEAN DEFAULT false,
    source_id UUID NOT NULL,
    created timestamp DEFAULT CURRENT_TIMESTAMP,
    edited timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_spells_edit_time BEFORE UPDATE
ON spells FOR EACH ROW EXECUTE PROCEDURE 
update_edit_time();

-- улучшение при лвл апах, классы и архетипы