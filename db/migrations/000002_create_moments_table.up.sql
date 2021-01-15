CREATE TYPE moment_type AS ENUM ('start','startBreak','stopBreak','end');
CREATE TABLE IF NOT EXISTS moments
(
    id      serial,
    list_id integer     NOT NULL,
    date    date        NOT NULL,
    time    time        NOT NULL,
    type    moment_type NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_lists
        FOREIGN KEY (list_id)
            REFERENCES lists (id)
            ON DELETE CASCADE
);