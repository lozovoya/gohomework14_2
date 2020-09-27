package app

import (
	"encoding/json"
	"github.com/lozovoya/gohomework14_2/cmd/bank/app/dto"
	"github.com/lozovoya/gohomework14_2/pkg/card"
	"github.com/lozovoya/gohomework14_2/pkg/db"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	mux     *http.ServeMux
	CardSvc *card.Service
	DbSvc   *db.Service
}

func NewServer(mux *http.ServeMux, cardSvc *card.Service, dbSvc *db.Service) *Server {
	return &Server{mux: mux, CardSvc: cardSvc, DbSvc: dbSvc}
}

func (s *Server) Init() {
	s.mux.HandleFunc("/getCards", s.getCards)
	s.mux.HandleFunc("/getTransactions", s.getTransactions)
}

func (s *Server) getCards(w http.ResponseWriter, r *http.Request) {

	cardid, err := strconv.Atoi(r.FormValue("owner_id"))
	if err != nil {
		log.Println(err)
		return
	}

	cards := card.GetCards(s.DbSvc.Ctx, s.DbSvc.Pool, cardid)
	if len(cards) == 0 {
		log.Println("no cards available")

		dtos := dto.MessageDTO{Message: "no cards available"}
		respBody, err := json.Marshal(dtos)
		if err != nil {
			log.Println(err)
			return
		}

		err = s.SendReply(w, respBody)
		if err != nil {
			log.Println(err)
			return
		}
		return
	}

	dtos := make([]dto.CardDTO, len(cards))
	for i, c := range cards {
		dtos[i] = dto.CardDTO{
			Id:      c.Id,
			Number:  c.Number,
			Balance: c.Balance,
			Issuer:  c.Issuer,
			Status:  c.Status,
		}
	}

	respBody, err := json.Marshal(dtos)
	if err != nil {
		log.Println(err)
		return
	}
	err = s.SendReply(w, respBody)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (s *Server) getTransactions(w http.ResponseWriter, r *http.Request) {

	cardid, err := strconv.Atoi(r.FormValue("card_id"))
	if err != nil {
		log.Println(err)
		return
	}

	transactions := card.GetTransactions(s.DbSvc.Ctx, s.DbSvc.Pool, cardid)
	if len(transactions) == 0 {
		log.Println("no transactions available")

		dtos := dto.MessageDTO{
			Message: "no transactions available",
		}
		respBody, err := json.Marshal(dtos)
		if err != nil {
			log.Println(err)
			return
		}

		err = s.SendReply(w, respBody)
		if err != nil {
			log.Println(err)
			return
		}
		return
	}

	dtos := make([]dto.TransactionDTO, len(transactions))
	for i, t := range transactions {
		dtos[i] = dto.TransactionDTO{
			Id:          t.Id,
			Amount:      t.Amount,
			Category:    t.Category,
			Description: t.Description,
			Logo:        t.Logo,
		}
	}

	respBody, err := json.Marshal(dtos)
	if err != nil {
		log.Println(err)
		return
	}
	err = s.SendReply(w, respBody)
	if err != nil {
		log.Println(err)
		return
	}
	return

}

func (s *Server) SendReply(w http.ResponseWriter, respBody []byte) (err error) {

	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(respBody)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
