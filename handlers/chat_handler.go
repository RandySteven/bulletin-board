package handlers

import (
	"context"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"task_mission/entities/dtos/requests"
	"task_mission/enums"
	"task_mission/interfaces/handlers"
	"task_mission/interfaces/usecases"
	"task_mission/utils"
)

type ChatHandler struct {
	usecase  usecases.IChatUseCase
	upgrader websocket.Upgrader
}

func (c *ChatHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		request = &requests.RoomRequest{}
		dataKey = `room`
	)
	if err := utils.BindRequest(r, &request); err != nil {
		utils.ResponseHandler(w, http.StatusBadRequest, "", nil, nil, err)
		return
	}
	result, customErr := c.usecase.CreateRoom(ctx, request)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), "", nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, "", &dataKey, result, customErr)
}

func (c *ChatHandler) GetAllRooms(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		dataKey = `rooms`
	)
	result, customErr := c.usecase.GetAllLoginUserRooms(ctx)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), "", nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, "success get rooms", &dataKey, result, nil)
}

func (c *ChatHandler) SendMessage(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = context.WithValue(r.Context(), enums.RequestID, uuid.New().String())
	)
	conn, err := c.upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}
	defer conn.Close()

	clients := make(map[*websocket.Conn]bool)
	broadcast := make(chan string, 10)

	clients[conn] = true

	for {
		var msg requests.ChatRequest
		err := conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err) {
				delete(clients, conn)
				break
			}
			http.Error(w, "Error reading json", http.StatusInternalServerError)
			return
		}

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			response, customErr := c.usecase.SendChat(ctx, &msg)
			if customErr != nil {
				conn.WriteJSON(map[string]string{"error": customErr.Error()})
				return
			}
			log.Println(response)
		}()

		go func() {
			defer wg.Done()
			broadcast <- msg.Message
			for client := range clients {
				go func(c *websocket.Conn) {
					if err := c.WriteJSON(msg.Message); err != nil {
						c.Close()
						delete(clients, c)
					}
				}(client)
			}
		}()

		wg.Wait()
	}

}

var _ handlers.IChatHandler = &ChatHandler{}

func NewChatHandler(usecase usecases.IChatUseCase) *ChatHandler {
	return &ChatHandler{
		usecase: usecase,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}
