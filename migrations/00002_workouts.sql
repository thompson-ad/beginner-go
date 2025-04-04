-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS workouts (
    id BIGSERIAL PRIMARY KEY,
    -- userid 
    title VARCHAR(255) NOT NULL,
    description TEXT,
    duration_minutes INTEGER NOT NULL,
    calories_burned INTEGER,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE workouts;
-- +goose StatementEnd