CREATE TABLE IF NOT EXISTS location (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(50) NOT NULL,
  longitude FLOAT NOT NULL,
  latitude FLOAT NOT NULL,
  date TIMESTAMP NOT NULL,
  city VARCHAR(50) NOT NULL,
  country VARCHAR(50) NOT NULL,
  image_url VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS "user" (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  first_name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS comment (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id uuid NOT NULL,
  value VARCHAR(255) NOT NULL,
  created_at DATE NOT NULL,

  CONSTRAINT fk_user_id
  FOREIGN KEY(user_id) REFERENCES "user"(id)
);

CREATE TABLE IF NOT EXISTS post (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id uuid NOT NULL,
  location_id uuid NOT NULL,
  created_at DATE NOT NULL,
  likes INTEGER NOT NULL,

  CONSTRAINT fk_user_id
  FOREIGN KEY(user_id) REFERENCES "user"(id),

  CONSTRAINT fk_location_id
  FOREIGN KEY(location_id) REFERENCES location(id)
);

CREATE TABLE IF NOT EXISTS post_comment (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  comment_id UUID NOT NULL,
  post_id UUID NOT NULL,
  
  FOREIGN KEY(comment_id) REFERENCES comment(id),
  FOREIGN KEY(post_id) REFERENCES post(id),

  UNIQUE(comment_id, post_id)
);
