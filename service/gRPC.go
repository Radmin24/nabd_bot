package service

import (
	"context"
	"fmt"
	"log"
	"not_a_boring_date_bot/bot"
	pb "not_a_boring_date_bot/grpc"
	"not_a_boring_date_bot/internal/models"
)

type ServergRPC struct {
	pb.UnimplementedGRPCServiceServer
	Bot_g *bot.Bot
}

func (s *ServergRPC) SendMessage(ctx context.Context, req *pb.GRPCMessage) (*pb.Response, error) {

	controllerResp := reqToController(req)

	if s.Bot_g == nil {
		return &pb.Response{
			Status:  "error",
			Message: "Bot object is nil",
		}, fmt.Errorf("bot object is nil")
	}
	err := bot.Sender(controllerResp, s.Bot_g, req.ChatId, ctx)

	if err != nil {
		return &pb.Response{
			Status:  "error",
			Message: err.Error(),
		}, err
	}

	return &pb.Response{
		Status:  "200",
		Message: "Message send",
	}, nil
}

func reqToController(req *pb.GRPCMessage) *models.ControllerResponce {

	log.Println(req.IsKb)
	if req.IsKb {
		return &models.ControllerResponce{
			Answer:    req.Mes,
			Delay:     int(req.Delay),
			Keyboard:  reqToControllerKeyboard(req.Keyboard),
			IsKb:      true,
			IsNextMsg: false,
		}
	} else {
		return &models.ControllerResponce{
			Answer:    req.Mes,
			Delay:     int(req.Delay),
			IsKb:      req.IsKb,
			IsNextMsg: false,
		}
	}
}

func reqToControllerKeyboard(keyboard *pb.Keyboard) models.Keyboard {
	var buttons []models.Button
	if keyboard != nil {
		for _, btn := range keyboard.Buttons {
			buttons = append(buttons, models.Button{
				Caption: btn.Caption,
				Data:    btn.Data,
				Order:   int(btn.Order),
				Row:     int(btn.Row),
			})
		}
	}

	return models.Keyboard{
		Button: buttons,
		Type:   keyboard.GetType(),
	}
}
