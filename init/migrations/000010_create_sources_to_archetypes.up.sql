CREATE TABLE sources_to_archetypes(
    id UUID NOT NULL PRIMARY KEY,
    source UUID NOT NULL,
    name varchar(30) NOT NULL,
    base_class_id UUID NOT NULL -- link to sources_to_classes.id
    created timestamp DEFAULT CURRENT_TIMESTAMP,
    edited timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_sources_to_archetypes_edit_time BEFORE UPDATE
    ON sources_to_archetypes FOR EACH ROW EXECUTE PROCEDURE
    update_edit_time();