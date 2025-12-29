-- name: InsertTickerCik :exec
INSERT INTO ticker_cik_mapping 
(ticker, cik)
VALUES (
    ?,
    ?
);

-- name: GetCIKByTicker :one
SELECT cik from ticker_cik_mapping
WHERE ticker = ?;

