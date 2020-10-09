package app

import (
	"context"
	"encoding/json"
	"github.com/lozovoya/gohomework14_2/cmd/bank/app/dto"
	"github.com/lozovoya/gohomework14_2/pkg/card"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	mux     *http.ServeMux
	CardSvc *card.Service
}

func NewServer(mux *http.ServeMux, cardSvc *card.Service) *Server {
	return &Server{mux: mux, CardSvc: cardSvc}
}

func (s *Server) Init() {
	s.mux.HandleFunc("/getCards", s.getCards)
	s.mux.HandleFunc("/getTransactions", s.getTransactions)
	s.mux.HandleFunc("/getMonMostFreq", s.getMonMostFreq)
	s.mux.HandleFunc("/getMonMostValue", s.getMonMostValue)
}

func (s *Server) getCards(w http.ResponseWriter, r *http.Request) {

	cardid, err := strconv.Atoi(r.FormValue("owner_id"))
	if err != nil {
		log.Println(err)
		return
	}

	cards := s.CardSvc.GetCards(context.Background(), cardid)
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

	transactions := s.CardSvc.GetTransactions(context.Background(), cardid)
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

func (s *Server) getMonMostFreq(w http.ResponseWriter, r *http.Request) {

	cardid, err := strconv.Atoi(r.FormValue("card_id"))
	if err != nil {
		log.Println(err)
		return
	}

	catid, count := s.CardSvc.GetMonMostFreq(context.Background(), cardid)

	dtos := dto.MonMostDTO{
		CatId: catid,
		Count: count,
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

func (s *Server) getMonMostValue(w http.ResponseWriter, r *http.Request) {

	cardid, err := strconv.Atoi(r.FormValue("card_id"))
	if err != nil {
		log.Println(err)
		return
	}

	catid, count := s.CardSvc.GetMonMostValue(context.Background(), cardid)

	dtos := dto.MonMostDTO{
		CatId: catid,
		Count: count,
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
