CREATE TABLE users(
    id UUID NOT NULL PRIMARY KEY,
    login VARCHAR(50) NOT NULL,
    hash_password VARCHAR(255) NOT NULL,
    email VARCHAR(50) NOT NULL,
    created timestamp DEFAULT CURRENT_TIMESTAMP,
    edited timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_spells_edit_time BEFORE UPDATE
ON users FOR EACH ROW EXECUTE PROCEDURE 
update_edit_time();