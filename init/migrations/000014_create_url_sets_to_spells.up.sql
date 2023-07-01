CREATE TABLE url_sets_to_spells(
                               id UUID NOT NULL PRIMARY KEY,
                               url_set_id UUID NOT NULL,
                               spell_id UUID NOT NULL
);