package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AleksandrMac/GeekBrains/Go/BackLevel1/http-server/pkg/handlers"
)

type Handler struct{}

func TestDownloadHandlers(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/download/?extension=jpg", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	downloadHandler := &handlers.DownloadHandler{Root: http.Dir("../../data")}
	downloadHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Проверяем тело ответа
	expected := `<table>
	<tr>
		<td>Имя</td>
		<td>Расширение</td>
		<td>Размер в байтах</td>
	</tr>
	<tr>
		<td><a href="file2">file2</a></td>
		<td>.jpg</td>
		<td>17</td>
	</tr>
	<tr>
		<td><a href="file1">file1</a></td>
		<td>.jpg</td>
		<td>19</td>
	</tr>
</table>
`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got \n%v want \n%v",
			rr.Body.String(), expected)
	}
}
