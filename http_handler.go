package main

import (
	pb "github.com/vipos89/oms/common/api"
	"log"
	"net/http"
)

type handler struct {
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *handler {
	return &handler{
		client: client,
	}
}

func (h *handler) registerRotes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customer_id}/orders", h.HandleCreateOrder)

}

func (h *handler) HandleCreateOrder(writer http.ResponseWriter, request *http.Request) {
	log.Println("HandleCreateOrder")
	var items []&pb.ItemsWithQuantities
	err := common.ReadJSON(request, &items)

	if err != nil {
		common.WriteJSONError(writer, http.StatusBadRequest, err.Error())
		return
	}

	h.client.CreateOrder(request.Context(), &pb.CreateOrderRequest{
		CustomerId: request.PathValue("customer_id"),
		Items: items
	})
}

//
//
//func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	w.WriteHeader(http.StatusOK)
//
//}
