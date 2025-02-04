CREATE TABLE containers (
    id          SERIAL   PRIMARY KEY,
    ip_address  TEXT     UNIQUE NOT NULL
);

CREATE TABLE pings (
    id                SERIAL     PRIMARY KEY,
    container_id      INTEGER    REFERENCES containers(id) ON DELETE CASCADE,
    ping_time         INTEGER    NOT NULL,
    was_successful    BOOLEAN    NOT NULL,
    date              TIMESTAMP  NOT NULL DEFAULT NOW()
);