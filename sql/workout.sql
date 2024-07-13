CREATE TABLE IF NOT EXISTS workout(
 id UUID NOT NULL CONSTRAINT pk_workout_id PRIMARY KEY,
 exercise_id UUID,
 repetitions INT,
 constraint fk_exercise foreign key (exercise_id) references exercise(id)
);
alter table workout owner to postgres;