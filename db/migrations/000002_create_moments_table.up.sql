DO
$$
    BEGIN
        CREATE TYPE moment_type
        AS ENUM ('start','startBreak','stopBreak','end');
    EXCEPTION
        WHEN duplicate_object THEN null;
    END
$$;

CREATE TABLE IF NOT EXISTS moments
(
    id      serial,
    list_id integer     NOT NULL,
    date    date        NOT NULL,
    time    timestamp   NOT NULL,
    type    moment_type NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_lists
        FOREIGN KEY (list_id)
            REFERENCES lists (id)
            ON DELETE CASCADE,
    UNIQUE (date, type)
);