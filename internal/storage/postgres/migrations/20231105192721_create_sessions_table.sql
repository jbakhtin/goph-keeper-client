-- +goose Up
-- +goose StatementBegin
CREATE TABLE sessions (
    id bigserial NOT NULL PRIMARY KEY,
    access_token text,
    refresh_token text,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now()
);

CREATE TRIGGER set_timestamp_trigger_sessions
BEFORE UPDATE ON sessions
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE sessions;
-- +goose StatementEnd
