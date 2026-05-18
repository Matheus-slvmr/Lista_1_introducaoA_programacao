package utils

import (
	"log"
)

// Adicionamos o campo ID na struct
type User struct {
	ID       int    
	Username string
	Email    string
	BornDate string
}

func GetUserByEmail(email string) (*User, error) {
    // Adicionamos o 'id' na query SELECT
	query := `SELECT id, username, email, born_date FROM users WHERE email = $1`
	var user User
	
    // Adicionamos o &user.ID no Scan para ler a informação do banco
	err := DB.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Email, &user.BornDate)
	if err != nil {
		log.Printf("Erro ao buscar usuário no banco de dados: %v", err)
		return nil, err
	}
	return &user, nil
}