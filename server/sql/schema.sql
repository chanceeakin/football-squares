CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  first_name TEXT,
  last_name TEXT,
  email TEXT UNIQUE NOT NULL
);

CREATE TABLE games (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  title text UNIQUE NOT NULL,
  created_at TIMESTAMP default current_timestamp,
  updated_at TIMESTAMP,
  begun_at TIMESTAMP,
  finished_at TIMESTAMP
);

CREATE TABLE messages (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  message_text TEXT,
  created_at TIMESTAMP default current_timestamp,
  updated_at TIMESTAMP,
  archived BOOLEAN NOT NULL default false,
  user_id uuid REFERENCES users (id),
  game_id uuid REFERENCES games (id)
);


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
