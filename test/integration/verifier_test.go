package integration

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/iost-official/go-iost/common"
	"github.com/iost-official/go-iost/core/contract"
	"github.com/iost-official/go-iost/core/tx"
	"github.com/iost-official/go-iost/ilog"
	. "github.com/iost-official/go-iost/verifier"
	"github.com/iost-official/go-iost/vm/database"
	"github.com/iost-official/go-iost/vm/host"
	"github.com/iost-official/go-iost/vm/native"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTransfer(t *testing.T) {
	ilog.Stop()

	Convey("test transfer success case", t, func() {
		s := NewSimulator()
		defer s.Clear()
		acc := prepareAuth(t, s)

		s.SetGas(acc.ID, 100000)
		createAccountsWithResource(s)
		createToken(t, s, acc)

		Reset(func() {
			s.Visitor.SetTokenBalanceFixed("iost", acc0.ID, "1000")
			s.Visitor.SetTokenBalanceFixed("iost", acc1.ID, "0")
			s.SetGas(acc.ID, 100000)
			s.SetRAM(acc.ID, 10000)
		})

		Convey("test transfer success case", func() {
			r, err := s.Call("token.iost", "transfer", fmt.Sprintf(`["iost","%v","%v","%v",""]`, acc0.ID, acc1.ID, 0.0001), acc.ID, acc.KeyPair)
			So(err, ShouldBeNil)
			So(r.Status.Message, ShouldEqual, "")
			So(s.Visitor.TokenBalance("iost", acc0.ID), ShouldEqual, int64(99999990000))
			So(s.Visitor.TokenBalance("iost", acc1.ID), ShouldEqual, int64(10000))
			So(r.GasUsage, ShouldEqual, 713400)
		})

		Convey("test of token memo", func() {
			r, err := s.Call("token.iost", "transfer", fmt.Sprintf(`["iost","%v","%v","%v","memo"]`, acc0.ID, acc1.ID, 0.0001), acc.ID, acc.KeyPair)
			So(err, ShouldBeNil)
			So(r.Status.Code, ShouldEqual, tx.Success)

			memo := "test"
			for i := 0; i < 10; i++ {
				memo = memo + memo
			}
			r, err = s.Call("token.iost", "transfer", fmt.Sprintf(`["iost","%v","%v","%v","%v"]`, acc0.ID, acc1.ID, 0.0001, memo), acc.ID, acc.KeyPair)
			So(err, ShouldBeNil)
			So(r.Status.Code, ShouldEqual, tx.ErrorRuntime)
			So(r.Status.Message, ShouldContainSubstring, "memo too large")
		})
	})
}

func TestSetCode(t *testing.T) {
	ilog.SetLevel(ilog.LevelInfo)
	Convey("set code", t, func() {
		s := NewSimulator()
		defer s.Clear()
		acc := prepareAuth(t, s)
		s.SetAccount(acc.ToAccount())
		s.SetGas(acc.ID, 10000000)
		s.SetRAM(acc.ID, 3000)

		c, err := s.Compile("hw", "test_data/helloworld", "test_data/helloworld")
		So(err, ShouldBeNil)
		So(len(c.Encode()), ShouldEqual, 146)
		cname, r, err := s.DeployContract(c, acc.ID, acc.KeyPair)
		s.Visitor.Commit()
		So(err, ShouldBeNil)
		So(r.Status.Code, ShouldEqual, tx.Success)
		So(cname, ShouldEqual, "ContractAhFA9ToFpBVg6hFgyRRf37XYh4w3e3a7TZUxGTSFdawA")
		So(r.GasUsage, ShouldEqual, 3593500)
		So(s.Visitor.TokenBalance("ram", acc.ID), ShouldEqual, int64(2694))

		r, err = s.Call(cname, "hello", "[]", acc.ID, acc.KeyPair)
		So(err, ShouldBeNil)
		So(r.Status.Message, ShouldEqual, "")
	})
}

func TestStringGas(t *testing.T) {
	ilog.SetLevel(ilog.LevelInfo)
	Convey("string op gas", t, func() {
		s := NewSimulator()
		defer s.Clear()
		acc := prepareAuth(t, s)
		s.SetAccount(acc.ToAccount())
		s.SetGas(acc.ID, 10000000)
		s.SetRAM(acc.ID, 3000)

		c, err := s.Compile("so", "test_data/stringop", "test_data/stringop")
		So(err, ShouldBeNil)
		So(c, ShouldNotBeNil)
		cname, r, err := s.DeployContract(c, acc.ID, acc.KeyPair)
		So(err, ShouldBeNil)
		So(r.Status.Code, ShouldEqual, tx.Success)

		r, err = s.Call(cname, "f1", "[]", acc.ID, acc.KeyPair)
		So(err, ShouldBeNil)
		So(r.Status.Code, ShouldEqual, 0)
		gas2 := r.GasUsage

		r, err = s.Call(cname, "f2", "[]", acc.ID, acc.KeyPair)
		So(err, ShouldBeNil)
		So(r.Status.Code, ShouldEqual, 0)
		So(r.GasUsage-gas2, ShouldBeBetweenOrEqual, 1300, 1500)

		r, err = s.Call(cname, "f3", "[]", acc.ID, acc.KeyPair)
		So(err, ShouldBeNil)
		So(r.Status.Code, ShouldEqual, 0)
		So(r.GasUsage-gas2, ShouldEqual, 1400)

		r, err = s.Call(cname, "f4", "[]", acc.ID, acc.KeyPair)
		So(err, ShouldBeNil)
		So(r.Status.Code, ShouldEqual, 0)
		So(r.GasUsage-gas2, ShouldBeGreaterThan, 1400)
	})
}

func TestJS_Database(t *testing.T) {
	//ilog.Stop()
	ilog.SetLevel(ilog.LevelInfo)
	Convey("test of s database", t, func() {
		s := NewSimulator()
		defer s.Clear()

		c, err := s.Compile("datatbase", "test_data/database", "test_data/database")
		So(err, ShouldBeNil)

		acc := prepareAuth(t, s)
		s.SetGas(acc.ID, 1000000)
		s.SetRAM(acc.ID, 10000)

		cname, _, err := s.DeployContract(c, acc.ID, acc.KeyPair)
		So(err, ShouldBeNil)

		So(s.Visitor.Contract(cname), ShouldNotBeNil)
		So(database.Unmarshal(s.Visitor.Get(cname+"-"+"num")), ShouldEqual, "9")
		So(database.Unmarshal(s.Visitor.Get(cname+"-"+"string")), ShouldEqual, "hello")
		So(database.Unmarshal(s.Visitor.Get(cname+"-"+"bool")), ShouldEqual, "true")
		So(database.Unmarshal(s.Visitor.Get(cname+"-"+"array")), ShouldEqual, "[1,2,3]")
		So(database.Unmarshal(s.Visitor.Get(cname+"-"+"obj")), ShouldEqual, `{"foo":"bar"}`)

		s.SetGas(acc.ID, 1000000)
		r, err := s.Call(cname, "read", `[]`, acc.ID, acc.KeyPair)

		So(err, ShouldBeNil)
		So(r.Status.Message, ShouldEqual, "")
		So(len(r.Returns), ShouldEqual, 1)
		So(r.Returns[0], ShouldEqual, `["true"]`)
	})

}

func TestAmountLimit(t *testing.T) {
	ilog.Stop()
	Convey("test of amount limit", t, func() {
		s := NewSimulator()
		defer s.Clear()
		createAccountsWithResource(s)

		createToken(t, s, acc0)

		ca, err := s.Compile("Contracttransfer", "./test_data/transfer", "./test_data/transfer.js")
		So(err, ShouldBeNil)
		So(ca, ShouldNotBeNil)
		s.SetContract(ca)

		ca, err = s.Compile("Contracttransfer1", "./test_data/transfer1", "./test_data/transfer1.js")
		So(err, ShouldBeNil)
		So(ca, ShouldNotBeNil)
		contractTransfer1, _, err := s.DeployContract(ca, acc0.ID, acc0.KeyPair)
		So(err, ShouldBeNil)

		s.SetRAM(acc0.ID, 10000)

		Reset(func() {
			s.Visitor.SetTokenBalanceFixed("iost", acc0.ID, "1000")
			s.Visitor.SetTokenBalanceFixed("iost", acc1.ID, "0")
			s.SetGas(acc0.ID, 100000)
			s.SetRAM(acc0.ID, 10000)
		})

		Convey("test of amount limit", func() {
			r, err := s.Call("Contracttransfer", "transfer", fmt.Sprintf(`["%v", "%v", "%v"]`, acc0.ID, acc1.ID, "10"), acc0.ID, acc0.KeyPair)
			s.Visitor.Commit()

			So(err, ShouldBeNil)
			So(r.Status.Code, ShouldEqual, tx.Success)
			balance0 := common.Fixed{Value: s.Visitor.TokenBalance("iost", acc0.ID), Decimal: s.Visitor.Decimal("iost")}
			balance1 := common.Fixed{Value: s.Visitor.TokenBalance("iost", acc1.ID), Decimal: s.Visitor.Decimal("iost")}
			So(balance0.ToString(), ShouldEqual, "990")
			So(balance1.ToString(), ShouldEqual, "10")

			r, err = s.Call("Contracttransfer", "transfer1", fmt.Sprintf(`["%v", "%v", "%v"]`, acc0.ID, acc1.ID, "10"), acc0.ID, acc0.KeyPair)
			s.Visitor.Commit()

			So(err, ShouldBeNil)
			So(r.Status.Message, ShouldEqual, "")
			balance0 = common.Fixed{Value: s.Visitor.TokenBalance("iost", acc0.ID), Decimal: s.Visitor.Decimal("iost")}
			balance1 = common.Fixed{Value: s.Visitor.TokenBalance("iost", acc1.ID), Decimal: s.Visitor.Decimal("iost")}
			So(balance0.ToString(), ShouldEqual, "980")
			So(balance1.ToString(), ShouldEqual, "20")

			r, err = s.Call("Contracttransfer", "transfer2", fmt.Sprintf(`["%v", "%v", "%v"]`, acc0.ID, acc1.ID, "9"), acc0.ID, acc0.KeyPair)
			s.Visitor.Commit()

			So(err, ShouldBeNil)
			So(r.Status.Message, ShouldEqual, "")
			balance0 = common.Fixed{Value: s.Visitor.TokenBalance("iost", acc0.ID), Decimal: s.Visitor.Decimal("iost")}
			balance1 = common.Fixed{Value: s.Visitor.TokenBalance("iost", acc1.ID), Decimal: s.Visitor.Decimal("iost")}
			So(balance0.ToString(), ShouldEqual, "971")
			So(balance1.ToString(), ShouldEqual, "29")

			r, err = s.Call("Contracttransfer", "transfer2", fmt.Sprintf(`["%v", "%v", "%v"]`, acc0.ID, acc1.ID, "9.9"), acc0.ID, acc0.KeyPair)
			s.Visitor.Commit()
			So(err, ShouldBeNil)
			So(r.Status.Message, ShouldContainSubstring, "exceed amountLimit in abi")

			r, err = s.Call("Contracttransfer", "transfer3", fmt.Sprintf(`["%v", "%v", "%v"]`, acc0.ID, acc1.ID, "10.1"), acc0.ID, acc0.KeyPair)
			s.Visitor.Commit()

			So(err, ShouldBeNil)
			So(r.Status.Message, ShouldEqual, "")
			balance0 = common.Fixed{Value: s.Visitor.TokenBalance("iost", acc0.ID), Decimal: s.Visitor.Decimal("iost")}
			balance1 = common.Fixed{Value: s.Visitor.TokenBalance("iost", acc1.ID), Decimal: s.Visitor.Decimal("iost")}
			So(balance0.ToString(), ShouldEqual, "960.9")
			So(balance1.ToString(), ShouldEqual, "39.1")

			r, err = s.Call("Contracttransfer", "transfer4", fmt.Sprintf(`["%v", "%v", "%v"]`, acc0.ID, acc1.ID, "10"), acc0.ID, acc0.KeyPair)
			s.Visitor.Commit()

			So(err, ShouldBeNil)
			So(r.Status.Message, ShouldContainSubstring, "exceed amountLimit in abi")
		})

		Convey("test amount limit transfer to self", func() {
			r, err := s.Call("Contracttransfer", "transfer", fmt.Sprintf(`["%v", "%v", "%v"]`, acc0.ID, acc0.ID, "1000"), acc0.ID, acc0.KeyPair)
			s.Visitor.Commit()

			So(err, ShouldBeNil)
			So(r.Status.Message, ShouldEqual, "")
			balance0 := common.Fixed{Value: s.Visitor.TokenBalance("iost", acc0.ID), Decimal: s.Visitor.Decimal("iost")}
			So(balance0.ToString(), ShouldEqual, "1000")

			r, err = s.Call("Contracttransfer", "transferFreeze", fmt.Sprintf(`["%v", "%v", "%v"]`, acc0.ID, acc0.ID, "1000"), acc0.ID, acc0.KeyPair)
			s.Visitor.Commit()

			So(err, ShouldBeNil)
			So(r.Status.Message, ShouldEqual, "")
			balance0 = common.Fixed{Value: s.Visitor.TokenBalance("iost", acc0.ID), Decimal: s.Visitor.Decimal("iost")}
			So(balance0.ToString(), ShouldEqual, "0")
		})

		Convey("test amount limit transferFreeze", func() {
			r, err := s.Call("Contracttransfer", "transferFreeze", fmt.Sprintf(`["%v", "%v", "%v"]`, acc0.ID, acc1.ID, "110"), acc0.ID, acc0.KeyPair)
			s.Visitor.Commit()

			So(err, ShouldBeNil)
			So(r.Status.Message, ShouldContainSubstring, "exceed amountLimit in abi")
			So(r.Status.Code, ShouldEqual, tx.ErrorRuntime)

			r, err = s.Call("Contracttransfer", "transferFreeze", fmt.Sprintf(`["%v", "%v", "%v"]`, acc0.ID, acc1.ID, "100"), acc0.ID, acc0.KeyPair)
			s.Visitor.Commit()

			So(err, ShouldBeNil)
			So(r.Status.Message, ShouldEqual, "")
			balance0 := common.Fixed{Value: s.Visitor.TokenBalance("iost", acc0.ID), Decimal: s.Visitor.Decimal("iost")}
			So(balance0.ToString(), ShouldEqual, "900")
		})

		Convey("test amount limit destroy", func() {
			r, err := s.Call("Contracttransfer", "destroy", fmt.Sprintf(`["%v", "%v"]`, acc0.ID, "110"), acc0.ID, acc0.KeyPair)
			s.Visitor.Commit()

			So(err, ShouldBeNil)
			So(r.Status.Message, ShouldContainSubstring, "exceed amountLimit in abi")
			So(r.Status.Code, ShouldEqual, tx.ErrorRuntime)

			r, err = s.Call("Contracttransfer", "destroy", fmt.Sprintf(`["%v", "%v"]`, acc0.ID, "100"), acc0.ID, acc0.KeyPair)
			s.Visitor.Commit()

			So(err, ShouldBeNil)
			So(r.Status.Message, ShouldEqual, "")
			balance0 := common.Fixed{Value: s.Visitor.TokenBalance("iost", acc0.ID), Decimal: s.Visitor.Decimal("iost")}
			So(balance0.ToString(), ShouldEqual, "900")
		})

		Convey("test out of amount limit, use signers ID", func() {
			r, err := s.Call("Contracttransfer", "transfer", fmt.Sprintf(`["%v", "%v", "%v"]`, acc0.ID, acc1.ID, "200"), acc0.ID, acc0.KeyPair)
			s.Visitor.Commit()

			So(err, ShouldBeNil)
			So(r.Status.Message, ShouldContainSubstring, "exceed amountLimit in abi")
			So(r.Status.Code, ShouldEqual, tx.ErrorRuntime)
		})

		Convey("test out of amount limit", func() {
			r, err := s.Call("Contracttransfer", "transfer", fmt.Sprintf(`["%v", "%v", "%v"]`, acc0.ID, acc1.ID, "110"), acc0.ID, acc0.KeyPair)
			s.Visitor.Commit()

			So(err, ShouldBeNil)
			So(r.Status.Message, ShouldContainSubstring, "exceed amountLimit in abi")
			So(r.Status.Code, ShouldEqual, tx.ErrorRuntime)
		})

		Convey("test amount limit two level invocation", func() {
			r, err := s.Call(contractTransfer1, "transfer", fmt.Sprintf(`["%v", "%v", "%v"]`, acc0.ID, acc1.ID, "120"), acc0.ID, acc0.KeyPair)
			s.Visitor.Commit()

			So(err, ShouldBeNil)
			So(r.Status.Message, ShouldEqual, "")
			balance0 := common.Fixed{Value: s.Visitor.TokenBalance("iost", acc0.ID), Decimal: s.Visitor.Decimal("iost")}
			balance1 := common.Fixed{Value: s.Visitor.TokenBalance("iost", acc1.ID), Decimal: s.Visitor.Decimal("iost")}
			So(balance0.ToString(), ShouldEqual, "880")
			So(balance1.ToString(), ShouldEqual, "120")
		})

		Convey("test amount limit transfer from multi signers", func() {
			s.SetAccount(acc2.ToAccount())
			s.Visitor.SetTokenBalanceFixed("iost", acc2.ID, "1000")

			trx := tx.NewTx([]*tx.Action{{
				Contract:   "Contracttransfer",
				ActionName: "transfermulti",
				Data:       fmt.Sprintf(`["%v", "%v", "%v", "%v"]`, acc0.ID, acc2.ID, acc1.ID, "60"),
			}}, []string{acc2.ID + "@active"}, s.GasLimit, 100, s.Head.Time+10000000, 0)
			trx.Time = s.Head.Time
			trx.AmountLimit = append(trx.AmountLimit, &contract.Amount{Token: "*", Val: "unlimited"})
			sign, err := tx.SignTxContent(trx, acc2.ID, acc2.KeyPair)
			if err != nil {
				t.Fatal(err)
			}
			trx.Signs = append(trx.Signs, sign)

			r, err := s.CallTx(trx, acc0.ID, acc0.KeyPair)
			s.Visitor.Commit()
			So(err, ShouldBeNil)
			So(r.Status.Message, ShouldEqual, "")

			balance0 := common.Fixed{Value: s.Visitor.TokenBalance("iost", acc0.ID), Decimal: s.Visitor.Decimal("iost")}
			balance1 := common.Fixed{Value: s.Visitor.TokenBalance("iost", acc1.ID), Decimal: s.Visitor.Decimal("iost")}
			balance2 := common.Fixed{Value: s.Visitor.TokenBalance("iost", acc2.ID), Decimal: s.Visitor.Decimal("iost")}
			So(balance0.ToString(), ShouldEqual, "940")
			So(balance1.ToString(), ShouldEqual, "120")
			So(balance2.ToString(), ShouldEqual, "940")

			trx = tx.NewTx([]*tx.Action{{
				Contract:   "Contracttransfer",
				ActionName: "transfermulti",
				Data:       fmt.Sprintf(`["%v", "%v", "%v", "%v"]`, acc0.ID, acc2.ID, acc1.ID, "61"),
			}}, []string{acc2.ID + "@active"}, s.GasLimit, 100, s.Head.Time+10000000, 0)
			trx.Time = s.Head.Time
			trx.AmountLimit = append(trx.AmountLimit, &contract.Amount{Token: "*", Val: "unlimited"})
			sign, err = tx.SignTxContent(trx, acc2.ID, acc2.KeyPair)
			if err != nil {
				t.Fatal(err)
			}
			trx.Signs = append(trx.Signs, sign)

			r, err = s.CallTx(trx, acc0.ID, acc0.KeyPair)
			s.Visitor.Commit()
			So(err, ShouldBeNil)
			So(r.Status.Message, ShouldContainSubstring, "exceed amountLimit in abi. need 122")
		})

		Convey("test invalid amount limit", func() {
			ca, err = s.Compile("Contracttransfer2", "./test_data/transfer2", "./test_data/transfer2.js")
			So(err, ShouldBeNil)
			So(ca, ShouldNotBeNil)
			_, _, err := s.DeployContract(ca, acc0.ID, acc0.KeyPair)
			So(err.Error(), ShouldContainSubstring, "abnormal char in amount")
		})

	})
}

func TestTxAmountLimit(t *testing.T) {
	ilog.Stop()
	Convey("test of tx amount limit", t, func() {
		s := NewSimulator()
		defer s.Clear()
		createAccountsWithResource(s)

		createToken(t, s, acc0)
		s.SetRAM(acc0.ID, 10000)

		Reset(func() {
			s.Visitor.SetTokenBalanceFixed("iost", acc0.ID, "1000")
			s.Visitor.SetTokenBalanceFixed("iost", acc1.ID, "0")
			s.SetGas(acc0.ID, 100000)
			s.SetRAM(acc0.ID, 10000)
		})

		Convey("test of tx amount limit", func() {
			trx := tx.NewTx([]*tx.Action{{
				Contract:   "token.iost",
				ActionName: "transfer",
				Data:       fmt.Sprintf(`["iost", "%v", "%v", "%v", ""]`, acc0.ID, acc1.ID, "10"),
			}}, nil, 10000000, 100, 10000000, 0)
			trx.AmountLimit = append(trx.AmountLimit, &contract.Amount{Token: "iost", Val: "100"})
			r, err := s.CallTx(trx, acc0.ID, acc0.KeyPair)
			s.Visitor.Commit()

			So(err, ShouldBeNil)
			So(r.Status.Code, ShouldEqual, tx.Success)
			balance0 := common.Fixed{Value: s.Visitor.TokenBalance("iost", acc0.ID), Decimal: s.Visitor.Decimal("iost")}
			balance2 := common.Fixed{Value: s.Visitor.TokenBalance("iost", acc1.ID), Decimal: s.Visitor.Decimal("iost")}
			So(balance0.ToString(), ShouldEqual, "990")
			So(balance2.ToString(), ShouldEqual, "10")
		})

		Convey("test out of amount limit", func() {
			trx := tx.NewTx([]*tx.Action{{
				Contract:   "token.iost",
				ActionName: "transfer",
				Data:       fmt.Sprintf(`["iost", "%v", "%v", "%v", ""]`, acc0.ID, acc1.ID, "110"),
			}}, nil, 10000000, 100, 10000000, 0)
			trx.AmountLimit = append(trx.AmountLimit, &contract.Amount{Token: "iost", Val: "100"})
			r, err := s.CallTx(trx, acc0.ID, acc0.KeyPair)
			s.Visitor.Commit()

			So(err, ShouldBeNil)
			So(r.Status.Code, ShouldEqual, tx.ErrorRuntime)
			So(r.Status.Message, ShouldContainSubstring, "exceed amountLimit in tx")
			// todo
			// balance2 := common.Fixed{Value: s.Visitor.TokenBalance("iost", acc1.ID), Decimal: s.Visitor.Decimal("iost")}
			// So(balance2.ToString(), ShouldEqual, "0")
		})

		Convey("test invalid amount limit", func() {
			trx := tx.NewTx([]*tx.Action{{
				Contract:   "token.iost",
				ActionName: "transfer",
				Data:       fmt.Sprintf(`["iost", "%v", "%v", "%v", ""]`, acc0.ID, acc1.ID, "110"),
			}}, nil, 10000000, 100, 10000000, 0)
			trx.AmountLimit = append(trx.AmountLimit, &contract.Amount{Token: "iost1", Val: "100"})

			_, err := s.CallTx(trx, acc0.ID, acc0.KeyPair)
			s.Visitor.Commit()
			So(err.Error(), ShouldContainSubstring, "token not exists in amountLimit")
		})
	})
}

func TestTokenMemo(t *testing.T) {
	ilog.Stop()
	Convey("test of token memo", t, func() {
		s := NewSimulator()
		defer s.Clear()
		createAccountsWithResource(s)

		createToken(t, s, acc0)
		s.SetRAM(acc0.ID, 10000)

		Reset(func() {
			s.Visitor.SetTokenBalanceFixed("iost", acc0.ID, "1000")
			s.Visitor.SetTokenBalanceFixed("iost", acc1.ID, "0")
			s.SetGas(acc0.ID, 100000)
			s.SetRAM(acc0.ID, 10000)
		})

	})
}

func TestNativeVM_GasLimit(t *testing.T) {
	ilog.Stop()
	Convey("test native vm gas limit", t, func() {
		s := NewSimulator()
		defer s.Clear()
		createAccountsWithResource(s)

		createToken(t, s, acc0)
		s.SetGas(acc0.ID, 100000)

		tx0 := tx.NewTx([]*tx.Action{{
			Contract:   "token.iost",
			ActionName: "transfer",
			Data:       fmt.Sprintf(`["iost", "%v", "%v", "%v", ""]`, acc0.ID, acc1.ID, "10"),
		}}, nil, 350000, 100, 10000000, 0)

		r, err := s.CallTx(tx0, acc0.ID, acc0.KeyPair)
		t.Log(err, r, r.Status)
		s.Visitor.Commit()
		So(err, ShouldBeNil)
		So(r.Status.Message, ShouldContainSubstring, "out of gas")
		So(r.Status.Code, ShouldEqual, tx.ErrorRuntime)
	})
}

func TestNativeVM_GasPledgeShortCut(t *testing.T) {
	ilog.Stop()
	Convey("test one can pledge for gas without initial gas", t, func() {
		s := NewSimulator()
		defer s.Clear()
		createAccountsWithResource(s)
		s.SetContract(native.GasABI())

		err := createToken(t, s, acc0)
		So(err, ShouldBeNil)
		var pledgeAmount int64 = 100
		var initialBalance int64 = 1000
		var expectedGasAfterPlegde = pledgeAmount * int64(native.GasImmediateReward.ToFloat())
		pledgeAction := &tx.Action{
			Contract:   "gas.iost",
			ActionName: "pledge",
			Data:       fmt.Sprintf(`["%v", "%v", "%v"]`, acc0.ID, acc0.ID, pledgeAmount),
		}
		var txGasLimit int64 = 100000
		Convey("normal case", func() {
			s.SetGas(acc0.ID, 0)
			tx0 := tx.NewTx([]*tx.Action{pledgeAction}, nil, txGasLimit*100, 100, 10000000, 0)
			tx0.AmountLimit = append(tx0.AmountLimit, &contract.Amount{Token: "iost", Val: "unlimited"})
			r, err := s.CallTx(tx0, acc0.ID, acc0.KeyPair)
			s.Visitor.Commit()
			So(err, ShouldBeNil)
			So(r.Status.Message, ShouldEqual, "")
			txGasUsage := r.GasUsage / 100
			So(s.GetGas(acc0.ID), ShouldEqual, expectedGasAfterPlegde-txGasUsage)
			So(s.Visitor.TokenBalanceFixed("iost", acc0.ID).ToString(), ShouldEqual, strconv.Itoa(int(initialBalance-pledgeAmount)))
		})
		SkipConvey("vm can kill tx if gas limit is not enough(TODO it is not possible in current code)", func() {
			s.SetGas(acc0.ID, 0)
			anotherAction := &tx.Action{
				Contract:   "token.iost",
				ActionName: "transfer",
				Data:       fmt.Sprintf(`["iost", "%v", "%v", "%v", ""]`, acc0.ID, acc1.ID, 5),
			}
			// the first action can run succ
			tx0 := tx.NewTx([]*tx.Action{pledgeAction, anotherAction}, nil, txGasLimit*100, 100, 10000000, 0)
			r, err := s.CallTx(tx0, acc0.ID, acc0.KeyPair)
			//txGasUsage := r.GasUsage/100
			s.Visitor.Commit()
			So(err, ShouldBeNil)
			So(r.Status.Message, ShouldContainSubstring, "out of gas")
			So(r.Status.Code, ShouldEqual, tx.ErrorRuntime)
			So(s.GetGas(acc0.ID), ShouldEqual, expectedGasAfterPlegde-txGasLimit)
			So(s.Visitor.TokenBalanceFixed("iost", acc0.ID).ToString(), ShouldEqual, strconv.Itoa(int(initialBalance-pledgeAmount)))
		})
	})
}

func TestDomain(t *testing.T) {
	Convey("test of domain", t, func() {
		s := NewSimulator()
		defer s.Clear()

		c, err := s.Compile("datatbase", "test_data/database", "test_data/database")
		So(err, ShouldBeNil)

		acc := prepareAuth(t, s)
		s.SetGas(acc.ID, 1e8)
		s.SetRAM(acc.ID, 10000)

		cname, _, err := s.DeployContract(c, acc.ID, acc.KeyPair)
		So(err, ShouldBeNil)
		s.Visitor.SetContract(native.ABI("domain.iost", native.DomainABIs))
		r1, err := s.Call("domain.iost", "Link", fmt.Sprintf(`["abc_0_de.io","%v"]`, cname), acc.ID, acc.KeyPair)
		So(err, ShouldBeNil)
		So(r1.Status.Message, ShouldEqual, "")
		r2, err := s.Call("abc_0_de.io", "read", "[]", acc.ID, acc.KeyPair)
		So(err, ShouldBeNil)
		So(r2.Status.Message, ShouldEqual, "")

		r1, err = s.Call("domain.iost", "Link", fmt.Sprintf(`["abcde#A","%v"]`, cname), acc.ID, acc.KeyPair)
		So(err, ShouldBeNil)
		So(r1.Status.Message, ShouldContainSubstring, "url contains invalid character")
	})
}

func TestAuthority(t *testing.T) {
	ilog.SetLevel(ilog.LevelInfo)
	s := NewSimulator()
	defer s.Clear()
	Convey("test of Auth", t, func() {

		ca, err := s.Compile("auth.iost", "../../contract/account", "../../contract/account.js")
		So(err, ShouldBeNil)
		s.Visitor.SetContract(ca)
		s.Visitor.SetContract(native.GasABI())

		acc := prepareAuth(t, s)
		s.SetGas(acc.ID, 1e8)
		s.SetRAM(acc.ID, 1000)
		s.SetRAM("myidid", 1000)

		r, err := s.Call("auth.iost", "SignUp", array2json([]interface{}{"myidid", acc.KeyPair.ID, acc.KeyPair.ID}), acc.ID, acc.KeyPair)
		So(err, ShouldBeNil)
		So(r.Status.Message, ShouldEqual, "")
		So(database.Unmarshal(s.Visitor.MGet("auth.iost-auth", "myidid")), ShouldStartWith, `{"id":"myidid",`)

		r, err = s.Call("auth.iost", "AddPermission", array2json([]interface{}{"myidid", "perm1", 1}), acc.ID, acc.KeyPair)
		So(err, ShouldBeNil)
		So(r.Status.Message, ShouldEqual, "")
		So(database.Unmarshal(s.Visitor.MGet("auth.iost-auth", "myidid")), ShouldContainSubstring, `"perm1":{"name":"perm1","groups":[],"items":[],"threshold":1}`)

		r, err = s.Call("auth.iost", "SignUp", array2json([]interface{}{"invalid#id", acc.ID, acc.ID}), acc.ID, acc.KeyPair)
		So(err, ShouldBeNil)
		So(r.Status.Message, ShouldContainSubstring, "id contains invalid character")

		Convey("referrer can be updated 1 time per 30 days", func() {
			accNew, _ := host.ReadAuth(s.Visitor, "myidid")
			So(accNew.Referrer, ShouldEqual, acc.ID)
			So(accNew.ReferrerUpdateTime, ShouldEqual, s.Head.Time)
			s.SetGas("myidid", 10000000)
			r, err = s.Call("auth.iost", "UpdateReferrer", array2json([]interface{}{"myidid", "hahaha"}), "myidid", acc.KeyPair)
			So(err, ShouldBeNil)
			So(r.Status.Message, ShouldContainSubstring, "referrer can only be updated one time per 30 days")
			s.Head.Time += 30 * 24 * 3600 * 1e9
			r, err = s.Call("auth.iost", "UpdateReferrer", array2json([]interface{}{"myidid", "hahaha"}), "myidid", acc.KeyPair)
			So(err, ShouldBeNil)
			So(r.Status.Message, ShouldEqual, "")
			accNew, _ = host.ReadAuth(s.Visitor, "myidid")
			So(accNew.Referrer, ShouldEqual, "hahaha")
		})
	})

}

func TestGasLimit2(t *testing.T) {
	ilog.Stop()
	Convey("test of gas limit 2", t, func() {
		s := NewSimulator()
		defer s.Clear()
		createAccountsWithResource(s)

		createToken(t, s, acc0)

		ca, err := s.Compile("Contracttransfer", "./test_data/transfer", "./test_data/transfer.js")
		So(err, ShouldBeNil)
		So(ca, ShouldNotBeNil)
		s.SetContract(ca)

		Convey("test of amount limit", func() {
			s.Visitor.SetTokenBalanceFixed("iost", acc0.ID, "1000")
			s.Visitor.SetTokenBalanceFixed("iost", acc1.ID, "0")
			s.SetGas(acc0.ID, 2000000)
			s.SetRAM(acc0.ID, 10000)

			acts := make([]*tx.Action, 0)
			for i := 0; i < 2; i++ {
				acts = append(acts, tx.NewAction("Contracttransfer", "transfer", fmt.Sprintf(`["%v", "%v", "%v"]`, acc0.ID, acc1.ID, "10")))
			}
			trx := tx.NewTx(acts, nil, 10000000, 100, s.Head.Time, 0)
			trx.AmountLimit = append(trx.AmountLimit, &contract.Amount{Token: "*", Val: "unlimited"})

			r, err := s.CallTx(trx, acc0.ID, acc0.KeyPair)
			s.Visitor.Commit()

			So(err, ShouldBeNil)
			So(r.Status.Message, ShouldEqual, "")
			So(r.GasUsage, ShouldEqual, int64(7250200))
			balance0 := common.Fixed{Value: s.Visitor.TokenBalance("iost", acc0.ID), Decimal: s.Visitor.Decimal("iost")}
			balance2 := common.Fixed{Value: s.Visitor.TokenBalance("iost", acc1.ID), Decimal: s.Visitor.Decimal("iost")}
			So(balance0.ToString(), ShouldEqual, "980")
			So(balance2.ToString(), ShouldEqual, "20")

			// out of gas
			s.Visitor.SetTokenBalanceFixed("iost", acc0.ID, "1000")
			s.Visitor.SetTokenBalanceFixed("iost", acc1.ID, "0")
			s.SetGas(acc0.ID, 2000000)
			s.SetRAM(acc0.ID, 10000)
			acts = []*tx.Action{}
			for i := 0; i < 4; i++ {
				acts = append(acts, tx.NewAction("Contracttransfer", "transfer", fmt.Sprintf(`["%v", "%v", "%v"]`, acc0.ID, acc1.ID, "10")))
			}
			trx = tx.NewTx(acts, nil, 2000000, 100, s.Head.Time, 0)
			trx.AmountLimit = append(trx.AmountLimit, &contract.Amount{Token: "*", Val: "unlimited"})

			r, err = s.CallTx(trx, acc0.ID, acc0.KeyPair)
			s.Visitor.Commit()

			So(err, ShouldBeNil)
			So(r.Status.Code, ShouldEqual, tx.ErrorRuntime)
			So(r.Status.Message, ShouldEqual, "out of gas")
			So(r.GasUsage, ShouldEqual, int64(2000000))
		})
	})
}
