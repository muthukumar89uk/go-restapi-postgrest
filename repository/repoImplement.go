package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	c "myProject/postgrest/helpers"
	"myProject/postgrest/models"
	"net/http"

	"github.com/spf13/viper"
)

type ToDoRepositoryImp struct{}

func NewToDoRepository() ToDoRepositoryInterface {
	return &ToDoRepositoryImp{}
}

func (t *ToDoRepositoryImp) GetAllTasks(ctx context.Context) (resp []models.Todo, err error) {
	response, err := makeRequestToPostgrest(http.MethodGet, viper.GetString(c.PostgrestBaseURL), nil)

	err = handleResponse(response, err, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (t *ToDoRepositoryImp) GetTaskById(ctx context.Context, id string) (resp []models.Todo, err error) {
	url := fmt.Sprintf("%s?id=eq.%s", viper.GetString(c.PostgrestBaseURL), id)

	response, err := makeRequestToPostgrest(http.MethodGet, url, nil)

	err = handleResponse(response, err, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (t *ToDoRepositoryImp) CreateTask(ctx context.Context, task models.Todo) (err error) {
	newTask, err := json.Marshal(task)
	if err != nil {
		return err
	}

	response, err := makeRequestToPostgrest(http.MethodPost, viper.GetString(c.PostgrestBaseURL), bytes.NewReader(newTask))

	err = handleResponse(response, err, nil)
	if err != nil {
		return err
	}

	return nil
}

func (t *ToDoRepositoryImp) UpdateTask(ctx context.Context, task models.TodoUpdate, id string) (err error) {
	newTask, err := json.Marshal(task)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s?id=eq.%s", viper.GetString(c.PostgrestBaseURL), id)

	response, err := makeRequestToPostgrest(http.MethodPatch, url, bytes.NewReader(newTask))

	err = handleResponse(response, err, nil)
	if err != nil {
		return err
	}

	return nil
}

func (t *ToDoRepositoryImp) DeleteTask(ctx context.Context, id string) (err error) {
	url := fmt.Sprintf("%s?id=eq.%s", viper.GetString(c.PostgrestBaseURL), id)

	response, err := makeRequestToPostgrest(http.MethodDelete, url, nil)

	err = handleResponse(response, err, nil)
	if err != nil {
		return err
	}

	return nil
}

func makeRequestToPostgrest(method string, url string, body io.Reader) (*http.Response, error) {
	if method == http.MethodGet || method == http.MethodDelete {
		request, err := http.NewRequest(method, url, nil)
		if err != nil {
			return nil, err
		}

		request.Header.Set(viper.GetString(c.Authorization), viper.GetString(c.Token))

		resp, err := http.DefaultClient.Do(request)
		if err != nil {
			return nil, err
		}

		return resp, err
	} else {
		request, err := http.NewRequest(method, url, body)
		if err != nil {
			return nil, err
		}

		request.Header.Set(viper.GetString(c.ContentType), viper.GetString(c.ApplicationJSON))

		request.Header.Set(viper.GetString(c.Authorization), viper.GetString(c.Token))

		resp, err := http.DefaultClient.Do(request)
		if err != nil {
			return nil, err
		}

		return resp, err
	}
}

func handleResponse(resp *http.Response, err error, data interface{}) error {
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		var errorResp models.PostgrestErrorResponse

		body, _ := io.ReadAll(resp.Body)

		if err := json.Unmarshal(body, &errorResp); err == nil {
			return fmt.Errorf("postgrest API error: %s - %s", errorResp.Code, errorResp.Message)
		}

		return fmt.Errorf("postgerst API error: %s", resp.Status)
	}

	if data != nil {
		err = json.NewDecoder(resp.Body).Decode(data)
		if err != nil {
			return err
		}
	}

	return nil
}
