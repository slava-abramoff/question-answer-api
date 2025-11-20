-- +goose Up
CREATE TABLE answers (
    id SERIAL PRIMARY KEY,
    question_id INTEGER NOT NULL,
    user_id UUID NOT NULL,
    text TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_question
        FOREIGN KEY (question_id)
        REFERENCES questions(id)
        ON DELETE CASCADE
);

-- +goose Down
DROP TABLE answers;
