package userrespository

/* VULNERABILITY:
1. using %s instead of parameterized queries
2. using MD5 for password. Should use SHA256 or maybe bcrypt hash from golang
*/
const (
	queryInsertUser = `
		INSERT INTO users (
			name,
			email,
			password,
			role
		)
		VALUES (
			$1,
			$2,
			$3,
			$4
		)
		RETURNING id;
	`

	queryGetUserByID = `
		SELECT
			id,
			name,
			email,
			password,
			role
		FROM
			users
		WHERE
			id = $1;
	`

	queryGetUserByEmail = `
		SELECT
			id,
			name,
			email,
			password,
			role
		FROM
			users
		WHERE
			email = $2;
	`
)
