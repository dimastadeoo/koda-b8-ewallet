package models

import "context"

type UserList struct {
	ID    int64
	Name  string
	Phone string
	Email string
}


func (r *UserRepository) GetAll() ([]UserList, error) {

	rows, err := r.tx.Query(
		context.Background(),
		`
		SELECT
			u.id,
			p.name,
			u.hp_number,
			u.email
		FROM users u
		JOIN profiles p
			ON p.id_user = u.id
		WHERE u.status_account = 'active'
		ORDER BY p.name ASC
		`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []UserList

	for rows.Next() {

		var user UserList

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Phone,
			&user.Email,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, rows.Err()
}