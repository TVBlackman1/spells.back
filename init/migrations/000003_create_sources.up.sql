CREATE TABLE sources(
    id UUID NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    version_number int DEFAULT 1,
    description text DEFAULT '',
    is_official BOOLEAN DEFAULT false,
    author VARCHAR(255),
    uploaded_by UUID,
    created timestamp DEFAULT CURRENT_TIMESTAMP,
    edited timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_spells_edit_time BEFORE UPDATE
ON sources FOR EACH ROW EXECUTE PROCEDURE 
update_edit_time();