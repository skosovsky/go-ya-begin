CREATE TABLE subjects (
    id    INTEGER
        PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(128) DEFAULT '' NOT NULL
);
CREATE INDEX subjects_title
    ON subjects (title);

CREATE TABLE teachers (
    id           INTEGER PRIMARY KEY AUTOINCREMENT,
    full_name    VARCHAR(256) DEFAULT '' NOT NULL,
    phone_number CHAR(11)     DEFAULT '' NOT NULL
);
CREATE INDEX teachers_full_name
    ON teachers (full_name);

CREATE TABLE subjectteacher (
    teacher_id INTEGER DEFAULT 0 NOT NULL,
    subject_id INTEGER DEFAULT 0 NOT NULL
);
CREATE INDEX subjectteacher_subject
    ON subjectteacher (subject_id);
CREATE INDEX subjectteacher_teacher
    ON subjectteacher (teacher_id);

CREATE TABLE classes (
    id            INTEGER PRIMARY KEY AUTOINCREMENT,
    serial_year   INTEGER DEFAULT 0   NOT NULL,
    serial_letter CHAR(1) DEFAULT ' ' NOT NULL
);

CREATE TABLE students (
    id                  INTEGER PRIMARY KEY AUTOINCREMENT,
    full_name           VARCHAR(256) DEFAULT '' NOT NULL,
    parent_phone_number CHAR(11)     DEFAULT '' NOT NULL,
    date_birth          CHAR(8)      DEFAULT '' NOT NULL,
    address_home        VARCHAR(512) DEFAULT '' NOT NULL,
    class_id            INTEGER      DEFAULT 0  NOT NULL
);
CREATE INDEX students_full_name
    ON students (full_name);
CREATE INDEX students_class
    ON students (class_id);

CREATE TABLE classteacher (
    class_id   INTEGER DEFAULT 0 NOT NULL,
    teacher_id INTEGER DEFAULT 0 NOT NULL
);
CREATE INDEX classteacher_class
    ON classteacher (class_id);
CREATE INDEX classteacher_teacher
    ON classteacher (teacher_id);
