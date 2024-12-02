CREATE TABLE "Reactions" (
    id                  SERIAL NOT NULL PRIMARY KEY,
    service_id          INT,
    action_id           INT,
    reaction_identifyer INT,
    user_email          VARCHAR(255),
    user_message        VARCHAR(255)
);

CREATE TABLE "Actions" (
    id                  SERIAL NOT NULL PRIMARY KEY,
    action_id           INT,
    reaction_identifyer INT,
    user_email          VARCHAR(255)
);
