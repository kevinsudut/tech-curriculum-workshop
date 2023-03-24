package reviewrepository

/* VULNERABILITY:
1. using %s instead of parameterized queries
*/
const (
	queryInsertReview = `
		INSERT INTO reviews (
			user_id,
			book_id,
			rate,
			"content",
			status
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5
		)
		RETURNING id;
	`

	queryUpdateReviewStatus = `
		UPDATE reviews
		SET
			status = $1
		WHERE
			id = $2;
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
			id = $1
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
			book_id = $1
		AND
			status = $2
		ORDER BY id;
	`
)
