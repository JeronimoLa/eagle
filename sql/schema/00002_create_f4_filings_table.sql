-- +goose Up
-- +goose StatementBegin
CREATE TABLE f4_filings (
    accession_number TEXT UNIQUE NOT NULL,
    document_type TEXT NOT NULL,
    period_of_report TEXT NOT NULL,
    issuer_cik TEXT NOT NULL,
    issuerTradingSymbol TEXT NOT NULL,
    rpt_owner_cik TEXT NOT NULL,
    rpt_owner_name TEXT NOT NULL,
    officer_title TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    form4_url TEXT NOT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE f4_filings;
-- +goose StatementEnd
