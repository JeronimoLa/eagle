-- name: InsertF4FilingRecord :many
INSERT INTO f4_filings 
(accession_number, document_type, period_of_report, issuer_cik, issuerTradingSymbol,
rpt_owner_cik, rpt_owner_name, officer_title, created_at, form4_url)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10
) 
RETURNING *;