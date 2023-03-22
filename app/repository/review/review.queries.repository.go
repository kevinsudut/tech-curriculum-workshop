package reviewrepository

const (
	queryInsertReview = `
		INSERT INTO reviews (
			user_id,
			book_id,
			rate,
			"content",
			status
		) VALUES (
			'%d',
			'%d',
			'%d',
			'%s',
			'%d'
		)
		RETURNING id;
	`

	queryUpdateReviewStatus = `
		UPDATE reviews
		SET
			status = '%d'
		WHERE
			id = '%d';
	`

	queryGetReviewByID = `
		SELECT
			id,
			user_id,
			book_id,
			rate,
			content,
			status,
			created_time
		FROM
			reviews
		WHERE
			id = '%d'
	`

	queryGetReviewByBookAndStatus = `
		SELECT
			id,
			user_id,
			book_id,
			rate,
			content,
			status,
			created_time
		FROM
			reviews
		WHERE
			book_id = '%d'
		AND
			status = '%d'
		ORDER BY id;
	`
)
