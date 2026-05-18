package utils

import (
	"log"
	"time"
)

// InsertAppointment salva um novo agendamento de consulta no banco de dados
func InsertAppointment(userID int, hospital, specialty string, date time.Time) error {
	query := `INSERT INTO appointments (user_id, hospital_name, specialty, appointment_date) VALUES ($1, $2, $3, $4)`
	
	_, err := DB.Exec(query, userID, hospital, specialty, date)
	if err != nil {
		log.Printf("Erro ao inserir agendamento no banco de dados: %v", err)
		return err
	}
	
	log.Println("Agendamento inserido com sucesso!")
	return nil
}