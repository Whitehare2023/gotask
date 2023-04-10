package blocks

// 生成task.go命令如下：
// abigen -sol Task.sol -pkg blocks -type task -out task.go

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var instance *Task
var ownerAddr string = "ab1e44cf924a3c795e66d15df67ee72ce7aeba5c"
var auth *bind.TransactOpts
var address common.Address
var tx *types.Transaction

const keyinfo = `{"address":"ab1e44cf924a3c795e66d15df67ee72ce7aeba5c","crypto":{"cipher":"aes-128-ctr","ciphertext":"8401367db56a0f38caf4d6c370326ea7fae68ef38fa9bf3b61a6024e535c5f7a","cipherparams":{"iv":"1c25d53571b1eb9078e3f3b15cbed674"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":4096,"p":6,"r":8,"salt":"2943077910a78e07d6893f23738006e8028425b707868bffd528841fd0f7a258"},"mac":"ca59bdd7afa659034b2efe9e7d5581d6c4a4fe634d94f5ea2311205a9465561d"},"id":"9bed8865-b4f4-4d06-91fc-77171277d417","version":3}`

type TaskInfox struct {
	TaskID  string `json:"task_id"`
	Issuer  string `json:"issuer"`
	Worker  string `json:"task_user"`
	Desc    string `json:"task_name"`
	Bonus   int64  `json:"bonus"`
	Comment string `json:"comment"`
	Status  uint8  `json:"task_status"`
}

// 当该包被调用时自动触发执行
func init() {
	// 1.连接到节点
	cli, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Panic("Failed to Dial! ", err)
	}
	defer cli.Close()
	// 2. 身份准备 NewTransactor(keyin io.Reader, passphrase string) (*TransactOpts, error)
	keyin := strings.NewReader(keyinfo)
	chainID, err := cli.ChainID(context.Background()) // 得到chainID
	if err != nil {
		log.Panic("Failed to ChainID! ", err)
	}
	auth, err = bind.NewTransactorWithChainID(keyin, "123", chainID)
	if err != nil {
		log.Panic("Failed to NewTransactor! ", err)
	}
	// 3.部署合约
	// DeployTask(auth *bind.TransactOpts, backend bind.ContractBackend, owner common.Address) (common.Address, *types.Transaction, *Task, error)
	addr, tx, task, err := DeployTask(auth, cli, common.HexToAddress(ownerAddr))
	if err != nil {
		log.Panic("Failed to NewTask", err)
	}
	fmt.Println("addr:", addr.Hex())
	fmt.Println("tx:", tx.Hash())
	instance = task
}

// 注册
func Register(username, passwd string) error {
	_, err := instance.Register(auth, username, passwd)
	if err != nil {
		fmt.Println("Failed to Register! ", err)
		return err
	}
	return nil
}

// 登录
func Login(username, passwd string) (bool, error) {
	return instance.Login(nil, username, passwd)
}

// 发布任务
func Issue(username, passwd, desc string, bonus int64) error {
	_, err := instance.Issue(auth, username, passwd, desc, big.NewInt(bonus))
	if err != nil {
		fmt.Println("Failed to Issue! ", err)
		return err
	}
	return nil
}

// 挖点钱
func Mint(To string, Value int64) error {
	// func (_Task *TaskTransactor) Mint(opts *bind.TransactOpts, _to string, _value *big.Int) (*types.Transaction, error)
	_, err := instance.Mint(auth, To, big.NewInt(Value))
	if err != nil {
		fmt.Println("Failed to Mint!")
		return err
	}
	return nil
}

// 修改任务
// status:1-接受 2-提交 3-确认 4-退回
func Update(username, passwd, common string, status uint8, taskID int64) error {
	// 合约调用
	if status == 1 {
		_, err := instance.Take(auth, username, passwd, big.NewInt(taskID))
		if err != nil {
			fmt.Println("Failed to Take task!", err)
			return err
		}
	} else if status == 2 {
		_, err := instance.Commit(auth, username, passwd, big.NewInt(taskID))
		if err != nil {
			fmt.Println("Failed to Commit!", err)
			return err
		}
	} else if status == 3 || status == 4 {
		if status == 4 {
			status = 1
		}
		_, err := instance.Confirm(auth, username, passwd, big.NewInt(taskID), common, status)
		if err != nil {
			fmt.Println("Failed to Confirm!", err)
			return err
		}
	}
	return nil
}

// 任务查询,列表请求
func Tasklist() ([]TaskInfox, error) {
	var tasks []TaskInfox
	taskbs, err := instance.QryAllTasks(nil)
	if err != nil {
		fmt.Println("Failed to QryAllTasks!", err)
		return nil, err
	}
	var task TaskInfox
	for k, v := range taskbs {
		task.TaskID = fmt.Sprintf("%d", k)
		task.Bonus = v.Bonus.Int64()
		task.Comment = v.Comment
		task.Desc = v.Desc
		task.Issuer = v.Issuer
		task.Status = v.Status
		task.Worker = v.Worker
		tasks = append(tasks, task)
	}
	return tasks, nil
}
