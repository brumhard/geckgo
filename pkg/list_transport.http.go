package pkg

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// lists
//AddList(ctx context.Context, name string, settings ListSettings) (List, error)
func decodeAddListRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	body := struct {
		Name     string       `json:"name,omitempty"`
		Settings ListSettings `json:"settings,omitempty"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return addListRequest{
		Name:     body.Name,
		Settings: body.Settings,
	}, nil
}

//GetLists(ctx context.Context) ([]List, error)
func decodeGetListsRequest(ctx context.Context, t *http.Request) (interface{}, error) {
	return getListRequest{}, nil
}

//GetList(ctx context.Context, listID string) (List, error)
func decodeGetListRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	listID := mux.Vars(r)["listId"]

	return getListRequest{ListID: listID}, nil
}

//UpdateList(ctx context.Context, listID string, settings ListSettings) (List, error)
func decodeUpdateListRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	listID := mux.Vars(r)["listId"]

	body := struct {
		Settings ListSettings `json:"settings,omitempty"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return updateListRequest{
		ListID:   listID,
		Settings: body.Settings,
	}, nil
}

//DeleteList(ctx context.Context, listID string) error
func decodeDeleteListRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	listID := mux.Vars(r)["listId"]

	return deleteListRequest{ListID: listID}, nil
}