CREATE TABLE "Action" (
    id           SERIAL NOT NULL PRIMARY KEY,
    time         VARCHAR(32),
    continent    VARCHAR(32),
    city         VARCHAR(32),
    user_mail    VARCHAR(32)
);
