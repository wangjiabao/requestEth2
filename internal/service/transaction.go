package service

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	pb "requestEth/api/requestEth/v1"
)

type TransactionService struct {
	pb.UnimplementedTransactionServer
}

func NewTransactionService() *TransactionService {
	return &TransactionService{}
}

func (s *TransactionService) UsdtBalance(ctx context.Context, req *pb.UsdtBalanceRequest) (*pb.UsdtBalanceReply, error) {
	//client, err := ethclient.Dial("https://data-seed-prebsc-1-s3.binance.org:8545/")
	client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")
	if err != nil {
		return &pb.UsdtBalanceReply{
			Balance: "",
			Err:     err.Error(),
		}, nil
	}

	tokenAddress := common.HexToAddress("0x55d398326f99059fF775485246999027B3197955")
	instance, err := NewUsdt(tokenAddress, client)
	if err != nil {
		return &pb.UsdtBalanceReply{
			Balance: "",
			Err:     err.Error(),
		}, nil
	}

	address := common.HexToAddress(req.Address)
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		return &pb.UsdtBalanceReply{
			Balance: "",
			Err:     err.Error(),
		}, nil
	}

	return &pb.UsdtBalanceReply{
		Balance: bal.String(),
		Err:     "",
	}, nil
}