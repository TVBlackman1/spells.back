CREATE TABLE spells_to_caster_specs(
    id UUID NOT NULL PRIMARY KEY,
    source UUID NOT NULL,
    spell UUID NOT NULL,
    archetypes UUID[] DEFAULT '{}',
    classes UUID[] DEFAULT '{}', -- link to sources_to_classes.id
    created timestamp DEFAULT CURRENT_TIMESTAMP,
    edited timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_spells_to_caster_specs_edit_time BEFORE UPDATE
    ON spells_to_caster_specs FOR EACH ROW EXECUTE PROCEDURE
    update_edit_time();