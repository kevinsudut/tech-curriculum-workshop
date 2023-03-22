package bookrepository

const (
	queryGetAllBook = `
		SELECT
			id,
			title,
			subtitle,
			description,
			isbn,
			author,
			published,
			publisher,
			pages
		FROM
			books
		ORDER BY id;
	`

	queryGetBookByID = `
		SELECT
			id,
			title,
			subtitle,
			description,
			isbn,
			author,
			published,
			publisher,
			pages
		FROM
			books
		WHERE
			id = '%d';
	`

	queryGetBookByTitle = `
		SELECT
			id,
			title,
			subtitle,
			description,
			isbn,
			author,
			published,
			publisher,
			pages
		FROM
			books
		WHERE
			title ILIKE '%%%s%%';
	`
)
