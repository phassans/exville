CREATE TABLE IF NOT EXISTS viraagh_user
(
  user_id       SERIAL,
  fl_name       TEXT      NOT NULL,
  username      TEXT      NOT NULL,
  password      TEXT      NOT NULL,
  linkedIn_URL  TEXT      NOT NULL,
  insert_time   TIMESTAMP NOT NULL,
  PRIMARY KEY   (user_id)
);

CREATE TABLE IF NOT EXISTS school
(
  school_id       SERIAL,
  school          TEXT      NOT NULL,
  degree          TEXT      NULL,
  field_of_study  TEXT      NULL,
  insert_time     TIMESTAMP NOT NULL,
  PRIMARY KEY     (school_id)
);

CREATE TABLE IF NOT EXISTS company
(
  company_id    SERIAL,
  company       TEXT      NOT NULL,
  location      TEXT      NULL,
  insert_time   TIMESTAMP NOT NULL,
  PRIMARY KEY   (company_id)
);

CREATE TABLE IF NOT EXISTS user_to_school
(
  user_id       NUMERIC   NOT NULL,
  school_id     NUMERIC   NOT NULL,
  from_year     INTEGER   NULL,
  to_year       INTEGER   NULL,
  insert_time   TIMESTAMP NOT NULL,
  PRIMARY KEY   (user_id,school_id)
);

CREATE TABLE IF NOT EXISTS user_to_company
(
  user_id       NUMERIC   NOT NULL,
  company_id    NUMERIC   NOT NULL,
  title         TEXT      NULL,
  from_year     INTEGER   NULL,
  to_year       INTEGER   NULL,
  insert_time   TIMESTAMP NOT NULL,
  PRIMARY KEY   (user_id,company_id)
);