package handlers

import (
	"context"
	"net/http"

	"github.com/lozhkindm/celeritas"
)

func (h *Handlers) render(w http.ResponseWriter, r *http.Request, tmpl string, vars, data interface{}) error {
	return h.App.Render.Page(w, r, tmpl, vars, data)
}

func (h *Handlers) sessionPut(ctx context.Context, key string, val interface{}) {
	h.App.Session.Put(ctx, key, val)
}

func (h *Handlers) sessionHas(ctx context.Context, key string) bool {
	return h.App.Session.Exists(ctx, key)
}

func (h *Handlers) sessionGet(ctx context.Context, key string) interface{} {
	return h.App.Session.Get(ctx, key)
}

func (h *Handlers) sessionGetString(ctx context.Context, key string) string {
	return h.App.Session.GetString(ctx, key)
}

func (h *Handlers) sessionRemove(ctx context.Context, key string) {
	h.App.Session.Remove(ctx, key)
}

func (h *Handlers) sessionRenew(ctx context.Context) error {
	return h.App.Session.RenewToken(ctx)
}

func (h *Handlers) sessionDestroy(ctx context.Context) error {
	return h.App.Session.Destroy(ctx)
}

func (h *Handlers) randomString(n int) string {
	return h.App.RandStr(n)
}

func (h *Handlers) encrypt(text string) (string, error) {
	enc := celeritas.Encryption{Key: []byte(h.App.EncryptionKey)}
	encrypted, err := enc.Encrypt(text)
	if err != nil {
		return "", err
	}
	return encrypted, nil
}

func (h *Handlers) decrypt(crypto string) (string, error) {
	enc := celeritas.Encryption{Key: []byte(h.App.EncryptionKey)}
	text, err := enc.Decrypt(crypto)
	if err != nil {
		return "", err
	}
	return text, nil
}
