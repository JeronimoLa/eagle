-- +goose Up
-- +goose StatementBegin
CREATE TABLE ticker_cik_mapping (
    id UUID PRIMARY KEY,
    ticker TEXT NOT NULL,
    cik TEXT NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE ticker_cik_mapping;
-- +goose StatementEnd
