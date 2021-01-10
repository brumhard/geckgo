CREATE TABLE IF NOT EXISTS list_settings
(
    id         serial,
    list_id    integer NOT NULL,
    daily_time interval,
    PRIMARY KEY (id),
    CONSTRAINT fk_lists
        FOREIGN KEY (list_id)
            REFERENCES lists (id)
            ON DELETE CASCADE
);