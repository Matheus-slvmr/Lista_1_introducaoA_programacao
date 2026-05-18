package handlers

import (
	"log"
	"net/http"
	"servidorHTTP/app/utils"
	"strconv"
	"text/template"
	"time"
)

// Estrutura para enviar os dados para a tela de sucesso
type AppointmentSuccessData struct {
	Hospital  string
	Specialty string
	Date      string
	Time      string
}

// ScheduleHandler processa o formulário de agendamento de consultas
func ScheduleHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	// Obtém os dados do formulário
	userIDStr := request.FormValue("user_id")
	hospital := request.FormValue("hospital")
	specialty := request.FormValue("specialty")
	dateStr := request.FormValue("appointment_date")

	// Converte o ID do usuário para inteiro
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(response, "ID de usuário inválido", http.StatusBadRequest)
		return
	}

	// Converte a string de data para o tipo time.Time do Go
	appointmentDate, err := time.Parse("2006-01-02T15:04", dateStr)
	if err != nil {
		http.Error(response, "Formato de data inválido", http.StatusBadRequest)
		return
	}

	// Salva no banco de dados
	err = utils.InsertAppointment(userID, hospital, specialty, appointmentDate)
	if err != nil {
		http.Error(response, "Erro ao salvar agendamento", http.StatusInternalServerError)
		return
	}

	// Prepara os dados para mostrar na tela de sucesso formatando a data e a hora
	data := AppointmentSuccessData{
		Hospital:  hospital,
		Specialty: specialty,
		Date:      appointmentDate.Format("02/01/2006"), // Formato Brasileiro: Dia/Mês/Ano
		Time:      appointmentDate.Format("15:04"),      // Formato de 24h: Hora:Minuto
	}

	// Carrega o template da tela de sucesso
	tmpl, err := template.ParseFiles("static/appointment_success.html")
	if err != nil {
		log.Printf("Erro ao carregar o template de sucesso: %v", err)
		http.Error(response, "Erro interno do servidor ao carregar a página", http.StatusInternalServerError)
		return
	}

	// Renderiza a página passando os dados formatados
	err = tmpl.Execute(response, data)
	if err != nil {
		log.Printf("Erro ao renderizar o template de sucesso: %v", err)
		http.Error(response, "Erro interno ao exibir a página", http.StatusInternalServerError)
		return
	}
}