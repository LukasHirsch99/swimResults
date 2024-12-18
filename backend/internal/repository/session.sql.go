// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: session.sql

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createSession = `-- name: CreateSession :one
INSERT INTO session (meetid, displaynr, day, warmupstart, sessionstart)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, meetid, day, warmupstart, sessionstart, displaynr
`

type CreateSessionParams struct {
	Meetid       int32
	Displaynr    int32
	Day          pgtype.Date
	Warmupstart  pgtype.Time
	Sessionstart pgtype.Time
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.db.QueryRow(ctx, createSession,
		arg.Meetid,
		arg.Displaynr,
		arg.Day,
		arg.Warmupstart,
		arg.Sessionstart,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Meetid,
		&i.Day,
		&i.Warmupstart,
		&i.Sessionstart,
		&i.Displaynr,
	)
	return i, err
}

const deleteSessionsForMeet = `-- name: DeleteSessionsForMeet :exec
DELETE FROM session WHERE meetid = $1
`

func (q *Queries) DeleteSessionsForMeet(ctx context.Context, meetid int32) error {
	_, err := q.db.Exec(ctx, deleteSessionsForMeet, meetid)
	return err
}

const getSessionByPk = `-- name: GetSessionByPk :one
SELECT id, meetid, day, warmupstart, sessionstart, displaynr FROM session WHERE meetid = $1 AND displaynr = $2
`

type GetSessionByPkParams struct {
	Meetid    int32
	Displaynr int32
}

func (q *Queries) GetSessionByPk(ctx context.Context, arg GetSessionByPkParams) (Session, error) {
	row := q.db.QueryRow(ctx, getSessionByPk, arg.Meetid, arg.Displaynr)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Meetid,
		&i.Day,
		&i.Warmupstart,
		&i.Sessionstart,
		&i.Displaynr,
	)
	return i, err
}

const getSessionCntForMeet = `-- name: GetSessionCntForMeet :one
SELECT count(*) FROM session WHERE meetid = $1
`

func (q *Queries) GetSessionCntForMeet(ctx context.Context, meetid int32) (int64, error) {
	row := q.db.QueryRow(ctx, getSessionCntForMeet, meetid)
	var count int64
	err := row.Scan(&count)
	return count, err
}
