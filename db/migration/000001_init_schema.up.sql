CREATE TABLE "location" (
	id uuid PRIMARY KEY,
    name VARCHAR ( 50 ) NOT NULL,
	longitude FLOAT NOT NULL,
	latitude FLOAT NOT NULL,
	date TIMESTAMP NOT NULL,
	city VARCHAR ( 50 ),
	country VARCHAR ( 50 )
);
