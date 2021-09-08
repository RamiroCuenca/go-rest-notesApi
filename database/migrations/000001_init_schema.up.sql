CREATE TABLE IF NOT EXISTS notes (
    id SERIAL,
    owner_name VARCHAR(55) NOT NULL,
    title VARCHAR(100) NOT NULL,
    details VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    CONSTRAINT notes_id_pk PRIMARY KEY (id)
);