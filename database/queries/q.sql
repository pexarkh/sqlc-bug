-- name: UidRangeRight :many
SELECT uid FROM emails
WHERE expunged = 0 AND user_id = ? AND mailbox_server_id = ?
  AND uid >= ?
  AND uid <= ?
ORDER BY uid ASC;

-- name: UidRangeWrong :many
SELECT uid FROM emails
WHERE expunged = 0 AND user_id = ? AND mailbox_server_id = ?
  AND ? <= uid
  AND ? >= uid
ORDER BY uid ASC;
