CREATE TABLE IF NOT EXISTS exercise (
  id UUID NOT NULL CONSTRAINT pk_exercise_id PRIMARY KEY,
  exercise VARCHAR NOT NULL CONSTRAINT uq_exercise_name UNIQUE,
  description VARCHAR,
  animation bytea
);
alter table exercise owner to postgres;