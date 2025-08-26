package health

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/nandoalvarado022/comments-play/internal/config"
)

type Status struct {
	Service string
	Status  string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	db, err := config.InitDB()
	dbStatus := "Activa ✅"
	if err != nil {
		dbStatus = fmt.Sprintf("Error de conexión ❌: %v", err)
	} else {
		defer db.Close()
	}

	statuses := []Status{
		{Service: "Base de datos", Status: dbStatus},
	}

	tmpl := `
<!DOCTYPE html>
<html>
<head>
    <style>
        table {
            width: 80%;
            margin: 20px auto;
            border-collapse: collapse;
        }
        th, td {
            padding: 12px;
            text-align: left;
            border: 1px solid #ddd;
        }
        th {
            background-color: #f2f2f2;
        }
        tr:nth-child(even) {
            background-color: #f9f9f9;
        }
        tr:hover {
            background-color: #f5f5f5;
        }
    </style>
</head>
<body>
    <h1 style="text-align: center;">Estado del Sistema</h1>
    <table>
        <tr>
            <th>Servicio</th>
            <th>Estado</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.Service}}</td>
            <td>{{.Status}}</td>
        </tr>
        {{end}}
    </table>
</body>
</html>`

	t, err := template.New("health").Parse(tmpl)
	if err != nil {
		http.Error(w, "Error al procesar la plantilla", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = t.Execute(w, statuses)
	if err != nil {
		http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
		return
	}
}
