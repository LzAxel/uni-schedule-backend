CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password_hash bytea NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    role SMALLINT NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE refresh_tokens (
    user_id INT NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    token TEXT NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW()
);


CREATE TABLE schedules (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    slug VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE teachers (
    id SERIAL PRIMARY KEY,
    schedule_id INT NOT NULL REFERENCES schedules(id) ON DELETE CASCADE,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    surname VARCHAR(50),
    UNIQUE (schedule_id, first_name, last_name, surname)
);

CREATE TABLE subjects (
    id SERIAL PRIMARY KEY,
    schedule_id INT NOT NULL REFERENCES schedules(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    UNIQUE (schedule_id, name)
);

CREATE TABLE classes (
    id SERIAL PRIMARY KEY,
    schedule_id INT NOT NULL REFERENCES schedules(id) ON DELETE CASCADE,
    teacher_id INT NOT NULL REFERENCES teachers(id) ON DELETE CASCADE,
    subject_id INT NOT NULL REFERENCES subjects(id) ON DELETE CASCADE,
    class_type VARCHAR(20) NOT NULL CHECK (class_type IN ('lecture', 'practice', 'lab', 'combined')),
    day_of_week VARCHAR(20) NOT NULL CHECK (day_of_week IN ('monday', 'tuesday', 'wednesday', 'thursday', 'friday', 'saturday')),
    class_number SMALLINT NOT NULL CHECK (class_number > 0),
    even_week BOOLEAN
);
