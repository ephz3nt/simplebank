package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"simplebank/util"
	"testing"
	"time"
)

func createRandomTransfer(t *testing.T,fromAccount,toAccount Account)Transfer{
	arg:= CreateTransferParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Amount:        util.RandomMoney(),
	}

	transfer,err:= testQueries.CreateTransfer(context.Background(),arg)
	require.NoError(t,err)
	require.NotEmpty(t,transfer)

	require.Equal(t,arg.FromAccountID,transfer.FromAccountID)
	require.Equal(t,arg.ToAccountID,transfer.ToAccountID)
	require.Equal(t,arg.Amount,transfer.Amount)

	require.NotZero(t,transfer.ID)
	require.NotZero(t,transfer.CreateAt)

	return transfer
}

func TestCreateTransfer(t *testing.T){
	fromAccount,toAccount:=createRandomAccount(t),createRandomAccount(t)
	createRandomTransfer(t,fromAccount,toAccount)
}

func TestGetTransfer(t *testing.T){
	fromAccount,toAccount:=createRandomAccount(t),createRandomAccount(t)
	transfer1:=createRandomTransfer(t,fromAccount,toAccount)
	transfer2,err:=testQueries.GetTransfer(context.Background(),transfer1.ID)
	require.NoError(t,err)
	require.NotEmpty(t,transfer2)

	require.Equal(t,transfer1.ID,transfer2.ID)
	require.Equal(t,transfer1.FromAccountID,transfer2.FromAccountID)
	require.Equal(t,transfer1.ToAccountID,transfer2.ToAccountID)
	require.Equal(t,transfer1.Amount,transfer2.Amount)
	// 检查时间差
	require.WithinDuration(t,transfer1.CreateAt,transfer2.CreateAt,time.Second)
}

func TestListTransfer(t *testing.T){
	fromAccount,toAccount:=createRandomAccount(t),createRandomAccount(t)
	for i:=0;i<10;i++ {
		createRandomTransfer(t,fromAccount,toAccount)
	}

	arg:=ListTransferParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers,err:=testQueries.ListTransfer(context.Background(),arg)
	require.NoError(t,err)
	require.Len(t,transfers,5)

	for _,transfer := range transfers{
		require.NotEmpty(t,transfer)
		require.True(t,transfer.FromAccountID == fromAccount.ID || transfer.ToAccountID == toAccount.ID)
	}
}