CREATE TABLE IF NOT EXISTS "User" (
    id                  SERIAL PRIMARY KEY,
    mail                TEXT UNIQUE NOT NULL,
    password            TEXT NOT NULL,
    name                TEXT,
    lastname            TEXT
);

CREATE TABLE IF NOT EXISTS "TimeAction" (
    id                  SERIAL PRIMARY KEY,
    area_id             TEXT,
    continent           TEXT,
    city                TEXT,
    hour                INT,
    minute              INT
);

CREATE TABLE IF NOT EXISTS "Area" (
    id                  SERIAL PRIMARY KEY,
    user_token          TEXT,
    area_id             TEXT,
    service_action_id   INT,
    service_reaction_id INT
);

CREATE TABLE IF NOT EXISTS "DiscordUser" (
    id                  SERIAL PRIMARY KEY,
    user_token          TEXT,
    access_token        TEXT
);

CREATE TABLE IF NOT EXISTS "DiscordReactions" (
    id                  SERIAL PRIMARY KEY,
    area_id             TEXT,
    reaction_type       INT,
    user_token          TEXT,
    channel_id          TEXT,
    guild_id            TEXT,
    message             TEXT
);

CREATE TABLE IF NOT EXISTS "DiscordAction" (
    id                  SERIAL PRIMARY KEY,
    action_type         INT,
    area_id             TEXT,
    user_token          TEXT
);

CREATE TABLE IF NOT EXISTS "SpotifyActions" (
    id                  SERIAL PRIMARY KEY,
    area_id             TEXT,
    action_type         INT,
    user_token          TEXT,
    user_id             TEXT,    
    is_playing          INT,
    nb_playlists        INT
);

CREATE TABLE IF NOT EXISTS "SpotifyReactions" (
    id                  SERIAL PRIMARY KEY,
    area_id             TEXT,
    reaction_type       INT,
    user_token          TEXT
);

CREATE TABLE IF NOT EXISTS "GitlabActions" (
    id                  SERIAL PRIMARY KEY,
    action_type         INT,
    area_id             TEXT
);