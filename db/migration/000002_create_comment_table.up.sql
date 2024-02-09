CREATE TABLE "comment" (
    id uuid PRIMARY KEY,
    post_id uuid,
    value VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    user_id uuid,
    CONSTRAINT fk_location
        FOREIGN KEY(post_id)
            REFERENCES location (id)
);
