CREATE TABLE users
(
    id            BIGSERIAL PRIMARY KEY,
    username      TEXT        NOT NULL UNIQUE,
    password_hash bytea       NOT NULL,
    role          SMALLINT    NOT NULL,
    created_at    TIMESTAMPTZ NOT NULL
);

CREATE TABLE teachers
(
    id         BIGSERIAL PRIMARY KEY,
    short_name TEXT NOT NULL UNIQUE,
    full_name  TEXT NOT NULL
);

CREATE TABLE lessons
(
    id          BIGSERIAL PRIMARY KEY,
    name        TEXT                            NOT NULL,
    location    TEXT                            NOT NULL,
    teacher_id  BIGINT REFERENCES teachers (id) NOT NULL,
    lesson_type SMALLINT                        NOT NULL
);

CREATE TABLE schedules
(
    id         BIGSERIAL PRIMARY KEY,
    creator_id BIGINT REFERENCES users (id) NOT NULL,
    name       TEXT                         NOT NULL,
    slug       TEXT                         NOT NULL UNIQUE
);

CREATE TABLE schedule_slots
(
    id                  BIGSERIAL PRIMARY KEY,
    schedule_id         BIGINT REFERENCES schedules (id) NOT NULL,
    weekday             SMALLINT                         NOT NULL,
    number              INT                              NOT NULL,
    is_alternating      BOOLEAN                          NOT NULL,
    even_week_lesson_id BIGINT REFERENCES lessons (id),
    odd_week_lesson_id  BIGINT REFERENCES lessons (id)
);