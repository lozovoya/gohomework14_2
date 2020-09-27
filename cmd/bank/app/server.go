package app

import (
	"github.com/lozovoya/gohomework14_2/pkg/card"
	"github.com/lozovoya/gohomework14_2/pkg/db"
	"log"
	"net/http"
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

	log.Println("get cards")
	cards := card.GetCards(s.DbSvc.Ctx, s.DbSvc.Pool, 3)
	log.Println(cards)

	//cards := card.GetCards(ctx)

	//cards := s.cardSvc.AllCards()
	//if len(cards) == 0 {
	//	log.Println("no cards available")
	//	err := s.SendReply(w, cards, "no cards available")
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	return
	//}
	//
	//err := s.SendReply(w, cards, "")
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
}

func (s *Server) getTransactions(w http.ResponseWriter, r *http.Request) {

	log.Println("get transactions")
	transactions := card.GetTransactions(s.DbSvc.Ctx, s.DbSvc.Pool, 2)
	log.Println(transactions)

}

func (s *Server) SendReply(w http.ResponseWriter, message string) (err error) {

	var respBody []byte

	//if len(cards) != 0 {
	//	dtos := make([]*dto.CardDTO, len(cards))
	//	for i, c := range cards {
	//		dtos[i] = &dto.CardDTO{
	//			Id:       c.Id,
	//			Number:   c.Number,
	//			Issuer:   c.Issuer,
	//			HolderId: c.HolderId,
	//			Type:     c.Type,
	//		}
	//	}
	//
	//	respBody, err = json.Marshal(dtos)
	//	if err != nil {
	//		log.Println(err)
	//		return err
	//	}
	//} else {
	//	var dtos = &dto.MessageDTO{
	//		Message: message,
	//	}
	//	respBody, err = json.Marshal(dtos)
	//	if err != nil {
	//		log.Println(err)
	//		return err
	//	}
	//}

	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(respBody)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
