package endpoint

import "net/http"

func putToQueue(queueSvc queueSvc) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		// TODO: parse req and call queueSvc.Put(...)
	}
}
