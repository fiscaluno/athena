package institution

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/fiscaluno/athena/pkg/entity"
)

var mockInstitution entity.Institution

func init() {

	mockInstitution = entity.Institution{
		ID:            entity.NewID(),
		Name:          "Name",
		ImageURL:      "ImageURL",
		AverageRating: 5.0,
		RatedByCount:  1,
		Website:       "Website",
		Cnpj:          "Cnpj",
		Address:       "Address",
		City:          "City",
		Province:      "Province",
		Emails:        []string{"Emails"},
		Phones:        []string{"Phones"},
	}

}

func mockingServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		switch r.URL.Path {

		case "/institution":

			mockEntity := mockInstitution

			mockEntities := []*entity.Institution{&mockEntity}

			resp, _ := json.Marshal(mockEntities)
			fmt.Fprintln(w, string(resp))

		}
	}))
}

func TestAPIRepo_FindAll(t *testing.T) {

	mockURI := mockingServer().URL
	mockPath := "/institution"

	mockEntity := mockInstitution

	mockEntities := []*entity.Institution{&mockEntity}

	type fields struct {
		uri  string
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*entity.Institution
		wantErr bool
	}{
		{
			"OK",
			fields{
				mockURI,
				mockPath,
			},
			mockEntities,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &APIRepo{
				uri:  tt.fields.uri,
				path: tt.fields.path,
			}
			got, err := r.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("APIRepo.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("APIRepo.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
