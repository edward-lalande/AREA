CREATE TABLE IF NOT EXISTS "User" (
    id                  SERIAL PRIMARY KEY,
    mail                VARCHAR(255) UNIQUE NOT NULL,
    password            VARCHAR(255) NOT NULL,
    name                VARCHAR(255),
    lastname            VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS "TimeAction" (
    id                  SERIAL PRIMARY KEY,
    user_mail           VARCHAR(255),
    continent           VARCHAR(255),
    city                VARCHAR(255),
    hour                INT,
    minute              INT,
    reaction_service_id INT,
    reaction_id         VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS "DiscordReactions" (
    id                  SERIAL PRIMARY KEY,
    reaction_identifyer VARCHAR(255),
    reaction_type       INT,
    user_token          VARCHAR(255),
    channel_id          VARCHAR(255),
    message             VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS "DiscordActions" (
    id                  SERIAL PRIMARY KEY,
    action_id           INT,
    reaction_identifyer INT,
    user_token          VARCHAR(255)
);
