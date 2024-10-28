CREATE TABLE IF NOT EXISTS users
(
    id            BIGSERIAL PRIMARY KEY,
    username      TEXT        NOT NULL UNIQUE,
    password_hash TEXT        NOT NULL,
    role          SMALLINT    NOT NULL,
    created_at    TIMESTAMPTZ NOT NULL
);
CREATE TABLE schedules
(
    id      SERIAL PRIMARY KEY,
    user_id INTEGER      NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    slug    VARCHAR(100) NOT NULL UNIQUE
);
CREATE TABLE teachers
(
    id         SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name  VARCHAR(50) NOT NULL,
    surname    VARCHAR(50) NOT NULL,
    schedule_id INTEGER     NOT NULL REFERENCES schedules (id) ON DELETE CASCADE
);
CREATE TABLE subjects
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    schedule_id INTEGER     NOT NULL REFERENCES schedules (id) ON DELETE CASCADE
);
CREATE TABLE classes
(
    id         SERIAL PRIMARY KEY,
    schedule_id INTEGER     NOT NULL REFERENCES schedules (id) ON DELETE CASCADE,
    subject_id INTEGER     NOT NULL REFERENCES subjects (id) ON DELETE CASCADE,
    teacher_id INTEGER     NOT NULL REFERENCES teachers (id) ON DELETE SET NULL,
    class_type VARCHAR(20) NOT NULL CHECK (class_type IN ('lecture', 'practice', 'lab', 'combined'))
);

CREATE TABLE schedule_entries
(
    id            SERIAL PRIMARY KEY,
    day           VARCHAR(10) NOT NULL CHECK (day IN
                                              ('monday', 'tuesday', 'wednesday', 'thursday', 'friday', 'saturday')),
    class_number  INTEGER     NOT NULL,
    even_class_id INTEGER     REFERENCES classes (id) ON DELETE SET NULL,
    odd_class_id  INTEGER     REFERENCES classes (id) ON DELETE SET NULL,
    is_static     BOOLEAN     NOT NULL,
    schedule_id   INTEGER     NOT NULL REFERENCES schedules (id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS refresh_tokens
(
    user_id    BIGINT REFERENCES users (id) NOT NULL UNIQUE,
    token      TEXT                         NOT NULL,
    updated_at TIMESTAMPTZ                  NOT NULL
);