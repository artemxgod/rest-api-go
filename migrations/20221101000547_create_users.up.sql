CREATE TABLE "users" (
  id serial NOT NULL PRIMARY KEY,
  email character varying(300) NOT NULL unique,
  encrypted_password character varying(100) NOT NULL
);
