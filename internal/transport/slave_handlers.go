package transport

import (
	"encoding/json"
	"net/http"
	"project/internal/logger"
	"project/internal/model"
)

func (t *transport) newSubscribe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	req, err := unmarshalData[model.NewSubscribe](w, r)
	if err != nil {
		logger.GetLogger().Err(err).Msg("failed to unmarshalData")

		return
	}

	if err := t.service.NewSubscribe(r.Context(), req.ChatId, req.UserId); err != nil {
		logger.GetLogger().Err(err).Msg("failed to add new subscribe")

		w.WriteHeader(http.StatusInternalServerError)

		w.Write([]byte(`{"error": "failed to add new subscribe"}`))

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (t *transport) getAllSubsciptions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	req, err := unmarshalData[model.GetAllSubs](w, r)
	if err != nil {
		logger.GetLogger().Err(err).Msg("failed to unmarshalData")

		return
	}

	data, err := t.service.GetAllSubsciptions(r.Context(), req.UserId)
	if err != nil {
		logger.GetLogger().Err(err).Msg("failed to get all subs")

		w.WriteHeader(http.StatusInternalServerError)

		w.Write([]byte(`{"error": "failed to get all subs"}`))

		return
	}

	b, err := json.Marshal(data)
	if err != nil {
		logger.GetLogger().Err(err).Msg("failed to marshal")

		w.WriteHeader(http.StatusInternalServerError)

		w.Write([]byte(`{"error": "failed to get all subs"}`))

		return
	}

	w.WriteHeader(http.StatusOK)

	w.Write(b)
}

func (t *transport) pay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	req, err := unmarshalData[model.NewSubscribe](w, r)
	if err != nil {
		logger.GetLogger().Err(err).Msg("failed to unmarshalData")

		return
	}

	if err := t.service.Pay(r.Context(), req.ChatId, req.UserId); err != nil {
		logger.GetLogger().Err(err).Msg("failed to pay")

		w.WriteHeader(http.StatusInternalServerError)

		w.Write([]byte(`{"error": "failed to pay"}`))

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (t *transport) isSubscribeExists(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	req, err := unmarshalData[model.NewSubscribe](w, r)
	if err != nil {
		logger.GetLogger().Err(err).Msg("failed to unmarshalData")

		return
	}

	data, err := t.service.IsSubscribeExists(r.Context(), req.ChatId, req.UserId)
	if err != nil {
		logger.GetLogger().Err(err).Msg("failed check is subscribe exist")

		w.WriteHeader(http.StatusInternalServerError)

		w.Write([]byte(`{"error": "failed check is subscribe exist"}`))

		return
	}

	b, err := json.Marshal(data)
	if err != nil {
		logger.GetLogger().Err(err).Msg("failed to marshal")

		w.WriteHeader(http.StatusInternalServerError)

		w.Write([]byte(`{"error": "failed check is subscribe exist"}`))

		return
	}

	w.WriteHeader(http.StatusOK)

	w.Write(b)
}

func (t *transport) isPaid(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	req, err := unmarshalData[model.NewSubscribe](w, r)
	if err != nil {
		logger.GetLogger().Err(err).Msg("failed to unmarshalData")

		return
	}

	data, err := t.service.IsPaid(r.Context(), req.ChatId, req.UserId)
	if err != nil {
		logger.GetLogger().Err(err).Msg("failed check paid status")

		w.WriteHeader(http.StatusInternalServerError)

		w.Write([]byte(`{"error": "failed check paid status"}`))

		return
	}

	b, err := json.Marshal(data)
	if err != nil {
		logger.GetLogger().Err(err).Msg("failed to marshal")

		w.WriteHeader(http.StatusInternalServerError)

		w.Write([]byte(`{"error": "failed check paid status"}`))

		return
	}

	w.WriteHeader(http.StatusOK)

	w.Write(b)
}
