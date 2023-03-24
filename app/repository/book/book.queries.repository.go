package bookrepository

/* VULNERABILITY:
1. using %s instead of parameterized queries
*/
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
			id = $1;
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
			title ILIKE '%' || $1 || '%';
	`
)
