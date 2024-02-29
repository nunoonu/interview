# interview

Interview microservice for managing candidates for interviewers.

use Header.Authorization to pass a token, value is Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4iLCJ1c2VySWQiOiIxZmZkOTgwZi0wZjFlLTQzNTYtYTZlYy1iYzUxOTg0MGYxMjkiLCJleHAiOjE3MDkyMjc1OTMsImlhdCI6MTcwOTA1NDc5MywiaXNzIjoiQmlrYXNoIn0.mItrocaqmkkndBQuMhWB8D2nYQ8HhP1oqIjLLsZwndw


# Table creation

Script for creating(DDL) tables and triggers for created_at and updated_at.


## Table people

CREATE TABLE people (
    id VARCHAR(255) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    name varchar(100) not null,
    role varchar(20) NOT NULL,
    created_by int8 NOT NULL,
    created_at timestamptz NOT null,
    updated_at timestamptz NOT null
);

CREATE TRIGGER set_people_insert
BEFORE INSERT on people
FOR EACH ROW
EXECUTE PROCEDURE trigger_insert_timestamp();

CREATE TRIGGER set_people_update
BEFORE UPDATE on people
FOR EACH ROW
EXECUTE PROCEDURE trigger_update_timestamp();



## Table appointment

CREATE TABLE appointment (
    id VARCHAR(255) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    card_name varchar(100) NOT NULL,
    message text NOT NULL,
    is_active bool not null,
    status varchar(15) NOT NULL,
    created_by VARCHAR(255) NOT NULL,
    created_at timestamptz NOT null,
    updated_at timestamptz NOT null
);

CREATE TRIGGER set_appointment_insert
BEFORE INSERT on appointment
FOR EACH ROW
EXECUTE PROCEDURE trigger_insert_timestamp();

CREATE TRIGGER set_appointment_update
BEFORE UPDATE on appointment
FOR EACH ROW
EXECUTE PROCEDURE trigger_update_timestamp();



## Table comment

CREATE TABLE comment (
    id VARCHAR(255) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    appointment_id VARCHAR(255) NOT NULL,
    message text NOT NULL,
    created_by VARCHAR(255) NOT NULL,
    created_at timestamptz NOT null,
    updated_at timestamptz NOT null
);

CREATE TRIGGER set_comment_insert
BEFORE INSERT on comment
FOR EACH ROW
EXECUTE PROCEDURE trigger_insert_timestamp();

CREATE TRIGGER set_comment_update
BEFORE UPDATE on comment
FOR EACH ROW
EXECUTE PROCEDURE trigger_update_timestamp();



## Table history

CREATE TABLE history (
    id VARCHAR(255) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    appointment_id VARCHAR(255) NOT NULL,
    card_name varchar(100) NOT NULL,
    message text NOT NULL,
    status varchar(15) not null,
    created_by VARCHAR(255) NOT NULL,
    created_at timestamptz NOT NULL
);

CREATE TRIGGER set_history_insert
BEFORE INSERT on history
FOR EACH ROW
EXECUTE PROCEDURE trigger_insert_only_created_at_timestamp();


# Store procedure

Store procedure for setting created_at and updated_at

CREATE OR REPLACE FUNCTION trigger_insert_only_created_at_timestamp()
RETURNS TRIGGER AS $$
BEGIN
NEW.created_at = NOW();
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION trigger_insert_timestamp()
RETURNS TRIGGER AS $$
BEGIN
NEW.updated_at = NOW();
NEW.created_at = NOW();
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION trigger_update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
NEW.updated_at = NOW();
RETURN NEW;
END;
$$ LANGUAGE plpgsql;