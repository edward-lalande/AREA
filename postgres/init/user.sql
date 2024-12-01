CREATE TABLE "User" (
    id           SERIAL NOT NULL PRIMARY KEY,
    mail         VARCHAR(32),
    password     VARCHAR(32),
    name         VARCHAR(32),
    lastname     VARCHAR(32)
);
