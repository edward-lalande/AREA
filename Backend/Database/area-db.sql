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
    action_type         INT,
    city                TEXT,
    continent           TEXT,
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
    channel_id          TEXT,
    message_id          TEXT,
    area_id             TEXT,
    user_token          TEXT
);

CREATE TABLE IF NOT EXISTS "GithubActions" (
    id                  SERIAL PRIMARY KEY,
    area_id             TEXT,
    action_type         INT,
    user_token          TEXT,
    pusher              TEXT,
    value               TEXT,
    number              INT
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

CREATE TABLE IF NOT EXISTS "GoogleActions" (
    id                  SERIAL PRIMARY KEY,
    user_token          TEXT,
    area_id             TEXT,
    action_type         INT,
    nb_message          INT,
    nb_events           INT
);

CREATE TABLE IF NOT EXISTS "GoogleReactions" (
    id                  SERIAL PRIMARY KEY,
    user_token          TEXT,
    area_id             TEXT,
    reaction_type       INT,
    summary             TEXT,
    description         TEXT,
    start_time          TEXT,
    end_time            TEXT,
    attendees           TEXT,
    recipient           TEXT,
    subject             TEXT,
    message             TEXT
);

CREATE TABLE IF NOT EXISTS "GitlabReactions" (
    id                  SERIAL PRIMARY KEY,
    user_token          TEXT,
    reaction_type       INT,
    area_id             TEXT,
    project_id          TEXT,
    body                TEXT
);

CREATE TABLE IF NOT EXISTS "MeteoActions" (
    id                  SERIAL PRIMARY KEY,
    area_id             TEXT,
    action_type         INT,
    latitude            TEXT,
    longitude           TEXT,
    value               INT
);
