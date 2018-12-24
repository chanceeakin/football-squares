CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  first_name TEXT,
  last_name TEXT,
  email TEXT UNIQUE NOT NULL
);

INSERT INTO users (email, first_name, last_name)
VALUES ('Chance@fake.com', 'Chance', 'Eakin');

CREATE TABLE games (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  title text UNIQUE NOT NULL,
  created_at TIMESTAMP default current_timestamp,
  updated_at TIMESTAMP,
  begun_at TIMESTAMP,
  finished_at TIMESTAMP
);

INSERT INTO games (title)
VALUES ('First');

CREATE TABLE messages (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  message_text TEXT,
  created_at TIMESTAMP default current_timestamp,
  updated_at TIMESTAMP,
  archived BOOLEAN NOT NULL default false,
  user_id uuid REFERENCES users (id),
  game_id uuid REFERENCES games (id)
);

INSERT INTO messages (message_text, created_at, user_id, game_id)
VALUES ('hey there', '2018-12-19 22:35:06 -6:00', 'C83F5B51-1031-45F3-ADF9-2CDA751BF8D1', 'C83F5B51-1031-45F3-ADF9-2CDA751BF8D1');

CREATE TABLE games_and_users (
  game_id uuid REFERENCES games (id) ON UPDATE CASCADE ON DELETE CASCADE,
  user_id uuid REFERENCES users (id) ON UPDATE CASCADE,
  CONSTRAINT games_and_user_pkey PRIMARY KEY (game_id, user_id)  -- explicit pk
);

CREATE TABLE bets (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id uuid REFERENCES users (id) ON DELETE CASCADE,
  data TEXT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);
