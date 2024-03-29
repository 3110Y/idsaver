package httpcontroller

import (
	"encoding/json"
	"github.com/3110Y/idsaver/internal/application/dto"
	"github.com/3110Y/idsaver/internal/application/service"
	"io"
	"log"
	"net/http"
)

type IdController struct {
	IdService *service.IdService
}

func NewIdController(idService *service.IdService) *IdController {
	return &IdController{IdService: idService}
}

func (i *IdController) PostId(writer http.ResponseWriter, request *http.Request) {
	dtoId := dto.IdDTO{}
	reqBody, err := io.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	json.Unmarshal(reqBody, &dtoId)
	add, err := i.IdService.Add(request.Context(), dtoId)
	if err != nil {
		log.Fatal(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	marshal, err := json.Marshal(add)
	if err != nil {
		log.Fatal(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	writer.Write(marshal)

}

func (i *IdController) GetId(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	id := ctx.Value("article").(string)
	item, err := i.IdService.Item(ctx, id)
	if err != nil {
		log.Fatal(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	marshal, err := json.Marshal(item)
	if err != nil {
		log.Fatal(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	writer.Write(marshal)
}
