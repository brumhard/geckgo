CREATE TABLE IF NOT EXISTS list_settings
(
    id         serial,
    list_id    integer NOT NULL UNIQUE,
    daily_time integer,
    PRIMARY KEY (id),
    CONSTRAINT fk_lists
        FOREIGN KEY (list_id)
            REFERENCES lists (id)
            ON DELETE CASCADE
);