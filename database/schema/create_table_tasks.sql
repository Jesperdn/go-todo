
-- create table
CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY ,
    name TEXT NOT NULL ,
    completed BOOLEAN
);
