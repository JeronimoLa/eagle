-- name: InsertF4FilingRecord :exec
INSERT INTO f4_filings 
(accession_number, document_type, period_of_report, issuer_cik, issuerTradingSymbol,
rpt_owner_cik, rpt_owner_name, officer_title, created_at, form4_url)
VALUES (
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?
);