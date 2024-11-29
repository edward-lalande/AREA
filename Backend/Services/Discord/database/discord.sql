CREATE TABLE "Reaction" (
    id                  SERIAL NOT NULL PRIMARY KEY,
    action_id           INT,
    reaction_identifyer INT,
    user_email          VARCHAR(32),
    user_message        VARCHAR(32)
);

CREATE TABLE "Actions" (
    id                  SERIAL NOT NULL PRIMARY KEY,
    action_id           INT,
    reaction_identifyer INT,
    user_email          VARCHAR(32)
);
