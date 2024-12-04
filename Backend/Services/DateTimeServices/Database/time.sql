CREATE TABLE "Action" (
    id           SERIAL NOT NULL PRIMARY KEY,
    user_mail    VARCHAR(32),
    continent    VARCHAR(32),
    city         VARCHAR(32),
    hour         INT,
    minute       INT
);
