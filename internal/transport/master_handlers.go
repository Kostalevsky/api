package transport

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"project/internal/logger"
	"project/internal/model"
)

func unmarshalData[T any](w http.ResponseWriter, r *http.Request) (T, error) {
	var res T

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		w.Write([]byte(`{"error": "failed to read request"}`))

		return res, fmt.Errorf("failed to read all %w", err)
	}

	defer r.Body.Close()

	if err := json.Unmarshal(body, &res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		w.Write([]byte(`{"error": "failed to unmarshal"}`))

		return res, fmt.Errorf("failed to unmasral. %w", err)
	}

	return res, nil
}

func (t *transport) addNewChat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	req, err := unmarshalData[model.AddNewChat](w, r)
	if err != nil {
		logger.GetLogger().Err(err).Msg("failed to unmarshalData")

		return
	}

	if err := t.service.AddNewChat(r.Context(), req.ChatId, req.OwnerId, req.Name, req.Description, req.Price); err != nil {
		logger.GetLogger().Err(err).Msg("failed to add new chat")

		w.WriteHeader(http.StatusInternalServerError)

		w.Write([]byte(`{"error": "failed to add new chat"}`))

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (t *transport) getChatsInfoByOwnerId(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	req, err := unmarshalData[model.Owner](w, r)
	if err != nil {
		logger.GetLogger().Err(err).Msg("failed to unmarshalData")

		return
	}

	data, err := t.service.GetChatsInfoByOwnerId(r.Context(), req.OwnerId)
	if err != nil {
		logger.GetLogger().Err(err).Msg("failed to get chats by chat id")

		w.WriteHeader(http.StatusInternalServerError)

		w.Write([]byte(`{"error": "failed to get chats by chat id"}`))

		return
	}

	b, err := json.Marshal(data)
	if err != nil {
		logger.GetLogger().Err(err).Msg("failed to marshal")

		w.WriteHeader(http.StatusInternalServerError)

		w.Write([]byte(`{"error": "failed to get chats by chat id"}`))

		return
	}

	w.WriteHeader(http.StatusOK)

	w.Write(b)
}

func (t *transport) disableChat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	req, err := unmarshalData[model.Chat](w, r)
	if err != nil {
		logger.GetLogger().Err(err).Msg("failed to unmarshalData")

		return
	}

	if err := t.service.DisableChat(r.Context(), req.ChatId); err != nil {
		logger.GetLogger().Err(err).Msg("failed to disable chat")

		w.WriteHeader(http.StatusInternalServerError)

		w.Write([]byte(`{"error": "failed to disable chat"}`))

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (t *transport) changeDescription(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	req, err := unmarshalData[model.ChangeDescription](w, r)
	if err != nil {
		logger.GetLogger().Err(err).Msg("failed to unmarshalData")

		return
	}

	if err := t.service.ChangeDescription(r.Context(), req.ChatId, req.Description); err != nil {
		logger.GetLogger().Err(err).Msg("failed to change description")

		w.WriteHeader(http.StatusInternalServerError)

		w.Write([]byte(`{"error": "failed to change description"}`))

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (t *transport) changePrice(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	req, err := unmarshalData[model.ChangePrice](w, r)
	if err != nil {
		logger.GetLogger().Err(err).Msg("failed to unmarshalData")

		return
	}

	if err := t.service.ChangePrice(r.Context(), req.ChatId, req.Price); err != nil {
		logger.GetLogger().Err(err).Msg("failed to change price")

		w.WriteHeader(http.StatusInternalServerError)

		w.Write([]byte(`{"error": "failed to change price"}`))

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (t *transport) getAllSlaves(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	req, err := unmarshalData[model.Chat](w, r)
	if err != nil {
		logger.GetLogger().Err(err).Msg("failed to unmarshalData")

		return
	}

	slaves, err := t.service.GetAllSlaves(r.Context(), req.ChatId)
	if err != nil {
		logger.GetLogger().Err(err).Msg("failed to get all slaves")

		w.WriteHeader(http.StatusInternalServerError)

		w.Write([]byte(`{"error": "failed to get all slaves"}`))

		return
	}

	b, err := json.Marshal(slaves)
	if err != nil {
		logger.GetLogger().Err(err).Msg("failed to marshal")

		w.WriteHeader(http.StatusInternalServerError)

		w.Write([]byte(`{"error": "failed to get all slaves"}`))

		return
	}

	w.WriteHeader(http.StatusOK)

	w.Write(b)
}
