# in brief

These two queries are almost identical except for the order of the comparisons in lines 3 and 4:

```sql
SELECT uid FROM emails
WHERE expunged = 0 AND user_id = ? AND mailbox_server_id = ?
  AND uid >= ?
  AND uid <= ?
ORDER BY uid ASC;
```

```sql
SELECT uid FROM emails
WHERE expunged = 0 AND user_id = ? AND mailbox_server_id = ?
  AND ? <= uid
  AND ? >= uid
ORDER BY uid ASC;
```

However sqlc generates different param structs:

```go
type UidRangeRightParams struct {
	UserID          string `db:"user_id" json:"user_id"`
	MailboxServerID string `db:"mailbox_server_id" json:"mailbox_server_id"`
	Uid             int64  `db:"uid" json:"uid"`
	Uid_2           int64  `db:"uid_2" json:"uid_2"`
}
```

```go
type UidRangeWrongParams struct {
	UserID          string      `db:"user_id" json:"user_id"`
	MailboxServerID string      `db:"mailbox_server_id" json:"mailbox_server_id"`
	Column3         interface{} `db:"column_3" json:"column_3"`
	Column4         interface{} `db:"column_4" json:"column_4"`
}
```

# how to regenerate

```sh
cd database && sqlc generate
```
