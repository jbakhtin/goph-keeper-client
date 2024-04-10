package query

const (
	CreateSession = `
		INSERT INTO sessions (access_token, refresh_token)
		VALUES($1, $2)
		RETURNING id, access_token, refresh_token, created_at, updated_at
	`

	UpdateSessionByID = `
		UPDATE sessions
		SET access_token = $1, refresh_token = $2
		WHERE sessions.id = $1
		RETURNING id, access_token, refresh_token, created_at, updated_at
	`

	GetSessionByID = `
		SELECT id, user_id, refresh_token, finger_print, expire_at, created_at, closed_at, updated_at FROM sessions
		WHERE sessions.id = $1 AND sessions.closed_at is NULL LIMIT 1
	`

	SearchSessionsTemp = `SELECT id, user_id, refresh_token, finger_print, expire_at, created_at, closed_at, updated_at FROM sessions`
)
