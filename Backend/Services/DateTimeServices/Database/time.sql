CREATE TABLE "Action" (
    id                  SERIAL NOT NULL PRIMARY KEY,
    user_mail           VARCHAR(255),
    continent           VARCHAR(255),
    city                VARCHAR(255),
    hour                INT,
    minute              INT
    reaction_service_id INT,
);
