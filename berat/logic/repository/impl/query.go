package impl

const (
	insertBeratQuery = `
		INSERT INTO 
			berat(
				tanggal,
				min,
				max
			)
		VALUES
			($1,$2,$3)
	`

	GetAllBeratQuery = `
		SELECT
			id,
			tanggal,
			max,
			min
		FROM
			berat
		ORDER BY 
			tanggal DESC
	`
	GetBeratByIdQuery = `
		SELECT
			id,
			tanggal,
			max,
			min
		FROM
			berat
		WHERE
			id = $1
	`
)
