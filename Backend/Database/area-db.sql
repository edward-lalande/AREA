CREATE TABLE IF NOT EXISTS "User" (
    id                  SERIAL PRIMARY KEY,
    mail                VARCHAR(255) UNIQUE NOT NULL,
    password            VARCHAR(255) NOT NULL,
    name                VARCHAR(255),
    lastname            VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS "TimeAction" (
    id                  SERIAL PRIMARY KEY,
    area_id             VARCHAR(255),
    continent           VARCHAR(255),
    city                VARCHAR(255),
    hour                INT,
    minute              INT
);

CREATE TABLE IF NOT EXISTS "Area" (
    id                  SERIAL PRIMARY KEY,
    user_token          VARCHAR(255),
    area_id             VARCHAR(255),
    service_action_id   INT,
    service_reaction_id INT
);

CREATE TABLE IF NOT EXISTS "DiscordUser" (
    id                  SERIAL PRIMARY KEY,
    user_token          VARCHAR(255),
    access_token        VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS "DiscordReactions" (
    id                  SERIAL PRIMARY KEY,
    area_id             VARCHAR(255),
    reaction_type       INT,
    user_token          VARCHAR(255),
    channel_id          VARCHAR(255),
    guild_id            VARCHAR(255),
    message             VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS "DiscordActions" (
    id                  SERIAL PRIMARY KEY,
    action_id           INT,
    action_type         INT,
    area_id             VARCHAR(255),
    user_token          VARCHAR(255)
);
