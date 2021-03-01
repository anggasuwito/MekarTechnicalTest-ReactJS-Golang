package constants

const (
	CHECK_ADMIN_USERNAME = `SELECT 
								a.id,
								a.username,
								a.password 
							FROM 
								admin AS a 
							WHERE 
								a.username = ? 
									AND 
								a.status = ?`
	GET_ALL_WORKS = `SELECT
						p.id,
						p.nama
					FROM
						pekerjaan AS p
					WHERE
						p.status = ?`
	GET_ALL_STUDIES = `SELECT
						p.id,
						p.nama
					FROM
						pendidikan AS p
					WHERE
						p.status = ?`
	GET_ALL_USERS = `SELECT 
						u.id,
						u.nama,
						u.tanggal_lahir,
						u.no_ktp,
						pk.id,
						pk.nama,
						pd.id,
						pd.nama
					FROM 
						user AS u
					JOIN
						pekerjaan AS pk
							ON
						pk.id = u.id_pekerjaan
					JOIN
						pendidikan AS pd
							ON
						pd.id = u.id_pendidikan
					WHERE 
						u.status = ?`
	GET_USER_BY_ID = `SELECT 
						u.id,
						u.nama,
						u.tanggal_lahir,
						u.no_ktp,
						pk.id,
						pk.nama,
						pd.id,
						pd.nama
					FROM 
						user AS u
					JOIN
						pekerjaan AS pk
							ON
						pk.id = u.id_pekerjaan
					JOIN
						pendidikan AS pd
							ON
						pd.id = u.id_pendidikan
					WHERE 
						u.id = ? 
							AND 
						u.status = ?`
	CREATE_NEW_USER = `INSERT INTO 
							user (
								nama,
								tanggal_lahir,
								no_ktp,
								id_pekerjaan,
								id_pendidikan) 
						VALUES 
							(?,?,?,?,?)`
	UPDATE_USER = `UPDATE 
						user
					SET 
						nama = ?,
						tanggal_lahir = ?,
						no_ktp = ?,
						id_pekerjaan = ?,
						id_pendidikan = ?
					WHERE
						id = ?
							AND
						status = ?`
	DELETE_USER = `UPDATE 
						user
					SET 
						status = ?
					WHERE
						id = ?
							AND
						status = ?`
)
