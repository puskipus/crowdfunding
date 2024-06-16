package payment

import (
	"crowdfunding/user"
	"strconv"

	"github.com/veritrans/go-midtrans"
)

type service struct {
}

type Service interface {
	GetPaymentResponse(transaction Transaction, user user.User) (midtrans.SnapResponse, error)
}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentResponse(transaction Transaction, user user.User) (midtrans.SnapResponse, error) {

	MIDTRANS_SERVER_KEY := "SB-Mid-server-2mVj7oljyBIsq6rsUAVSLuiB"
	MIDTRANS_CLIENT_KEY := "SB-Mid-client-W0Z6ft5ya6OrCaFV"

	midclient := midtrans.NewClient()
	midclient.ServerKey = MIDTRANS_SERVER_KEY
	midclient.ClientKey = MIDTRANS_CLIENT_KEY
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return snapTokenResp, err
	}

	return snapTokenResp, nil
}
