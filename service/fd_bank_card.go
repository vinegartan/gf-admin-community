// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/model"
	"github.com/SupenBysz/gf-admin-community/model/entity"
)

type (
	IFdBankCard interface {
		CreateBankCard(ctx context.Context, info model.BankCardRegister) (*entity.FdBankCard, error)
		GetBankCardById(ctx context.Context, id int64) (*entity.FdBankCard, error)
		GetBankCardByCardNumber(ctx context.Context, cardNumber string) (*entity.FdBankCard, error)
		UpdateBankCardState(ctx context.Context, bankCardId int64, state int) (bool, error)
	}
)

var (
	localFdBankCard IFdBankCard
)

func FdBankCard() IFdBankCard {
	if localFdBankCard == nil {
		panic("implement not found for interface IFdBankCard, forgot register?")
	}
	return localFdBankCard
}

func RegisterFdBankCard(i IFdBankCard) {
	localFdBankCard = i
}
