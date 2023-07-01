CREATE TABLE url_sets(
                     id UUID NOT NULL PRIMARY KEY,
                     url varchar(255),
                     name VARCHAR(255) NOT NULL,
                     created timestamp DEFAULT CURRENT_TIMESTAMP,
                     edited timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_url_sets_edit_time BEFORE UPDATE
    ON url_sets FOR EACH ROW EXECUTE PROCEDURE
    update_edit_time();