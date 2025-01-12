package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

// Definir el esquema GraphQL
var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "RootQuery",
				Fields: graphql.Fields{
					"hello": &graphql.Field{
						Type:        graphql.String,
						Description: "Devuelve un saludo",
						Resolve: func(params graphql.ResolveParams) (interface{}, error) {
							return "PROGRAMACION DISTRIBUIDA", nil
						},
					},
				},
			},
		),
	},
)

func main() {
	// Manejar las peticiones GraphQL
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Solo se permiten peticiones POST", http.StatusMethodNotAllowed)
			return
		}

		var request struct {
			Query string `json:"query"`
		}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, "Error al decodificar el JSON", http.StatusBadRequest)
			return
		}

		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: request.Query,
		})

		if len(result.Errors) > 0 {
			log.Printf("Errores en la consulta: %+v\n", result.Errors)
			http.Error(w, "Errores en la consulta", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})

	// Iniciar el servidor
	log.Println("Servidor ejecutándose en http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
