package userrespository

const (
	queryInsertUser = `
		INSERT INTO users (
			name,
			email,
			password,
			role
		)
		VALUES (
			'%s',
			'%s',
			MD5('%s'),
			'%s'
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
			id = '%d';
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
			email = '%s';
	`
)
