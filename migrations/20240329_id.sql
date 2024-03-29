-- +goose Up
CREATE TABLE id (
                         id uuid NOT NULL,
                         uid text,
                         len integer DEFAULT 0,
                         allow boolean DEFAULT true,
                         create_at timestamp DEFAULT CURRENT_TIMESTAMP,
                         update_at timestamp DEFAULT CURRENT_TIMESTAMP,
                         PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE ID;