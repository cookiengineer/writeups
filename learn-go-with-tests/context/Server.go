package context

import "log"
import "net/http"

func Server(store Store) http.HandlerFunc {

	return func(response http.ResponseWriter, request *http.Request) {

		data, err := store.Fetch(request.Context())

		if err != nil {
			log.Fatal(err)
		}

		response.Write([]byte(data))

	}

}
