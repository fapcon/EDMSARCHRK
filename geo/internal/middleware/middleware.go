package middlware

import (
	"encoding/json"
	"geo/internal/models"
	"gitlab.com/ptflp/gopubsub/queue"
	"golang.org/x/time/rate"
	"log"
	"net/http"
	"time"
)

type Limit struct {
	Req    int
	Queuer queue.MessageQueuer
}

func NewLimit(req int, queuer queue.MessageQueuer) *Limit {
	return &Limit{Req: req, Queuer: queuer}
}
func (l *Limit) Limiter(next http.Handler) http.Handler {
	limiter := rate.NewLimiter(rate.Every(time.Minute/time.Duration(l.Req)), l.Req)
	data := &models.Data{
		Email: "qwer",
		Phone: "1234",
	}
	jdata, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			err = l.Queuer.Publish("ratelimiter", jdata)
			if err != nil {
				log.Fatal(err)
			}
			http.Error(w, "Too many requests, please try again later.", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
