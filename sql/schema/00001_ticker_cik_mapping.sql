-- +goose Up
-- +goose StatementBegin
CREATE TABLE ticker_cik_mapping (
    cik TEXT NOT NULL UNIQUE,
    ticker TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE ticker_cik_mapping;
-- +goose StatementEnd
