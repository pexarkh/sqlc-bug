// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: q.sql

package database

import (
	"context"
)

const UidRangeRight = `-- name: UidRangeRight :many
SELECT uid FROM emails
WHERE expunged = 0 AND user_id = ? AND mailbox_server_id = ?
  AND uid >= ?
  AND uid <= ?
ORDER BY uid ASC
`

type UidRangeRightParams struct {
	UserID          string `db:"user_id" json:"user_id"`
	MailboxServerID string `db:"mailbox_server_id" json:"mailbox_server_id"`
	Uid             int64  `db:"uid" json:"uid"`
	Uid_2           int64  `db:"uid_2" json:"uid_2"`
}

func (q *Queries) UidRangeRight(ctx context.Context, arg *UidRangeRightParams) ([]int64, error) {
	rows, err := q.db.QueryContext(ctx, UidRangeRight,
		arg.UserID,
		arg.MailboxServerID,
		arg.Uid,
		arg.Uid_2,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var uid int64
		if err := rows.Scan(&uid); err != nil {
			return nil, err
		}
		items = append(items, uid)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const UidRangeWrong = `-- name: UidRangeWrong :many
SELECT uid FROM emails
WHERE expunged = 0 AND user_id = ? AND mailbox_server_id = ?
  AND ? <= uid
  AND ? >= uid
ORDER BY uid ASC
`

type UidRangeWrongParams struct {
	UserID          string      `db:"user_id" json:"user_id"`
	MailboxServerID string      `db:"mailbox_server_id" json:"mailbox_server_id"`
	Column3         interface{} `db:"column_3" json:"column_3"`
	Column4         interface{} `db:"column_4" json:"column_4"`
}

func (q *Queries) UidRangeWrong(ctx context.Context, arg *UidRangeWrongParams) ([]int64, error) {
	rows, err := q.db.QueryContext(ctx, UidRangeWrong,
		arg.UserID,
		arg.MailboxServerID,
		arg.Column3,
		arg.Column4,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var uid int64
		if err := rows.Scan(&uid); err != nil {
			return nil, err
		}
		items = append(items, uid)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
