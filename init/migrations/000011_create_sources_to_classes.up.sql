CREATE TABLE sources_to_classes(
    id UUID NOT NULL PRIMARY KEY,
    source UUID NOT NULL,
    name varchar(30) NOT NULL,
    created timestamp DEFAULT CURRENT_TIMESTAMP,
    edited timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_sources_to_classes_edit_time BEFORE UPDATE
    ON sources_to_classes FOR EACH ROW EXECUTE PROCEDURE
    update_edit_time();