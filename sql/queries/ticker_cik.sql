-- name: InsertTickerCik :many
INSERT INTO ticker_cik_mapping 
(id, ticker, cik)
VALUES (
    $1,
    $2,
    $3
) 
RETURNING *;

-- name: GetCIKByTicker :one
SELECT cik from ticker_cik_mapping
WHERE ticker = $1;

