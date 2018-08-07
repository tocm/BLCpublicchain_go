package cmd

import (
	"fmt"
	"os"
	"flag"
	"log"
	"BLCpublicchain_go/blc"
)

func isValidArgs()  {
	if len(os.Args) < 2 {
		showUsageTips()
		//退出
		os.Exit(1)
	}
}

func showUsageTips()  {
	fmt.Println("======== Usage ==========")
	fmt.Println("\tcreateGenesisBlock ----- 创建创世区块.")
	fmt.Println("\taddTransition -data DATA ----- 新增交易数据.")
	fmt.Println("\tprintchain ----- 是否允许输出区块信息.")


}

type CmdParams struct {
	BlockChain *blc.Blockchain
	CreateGenesisBlockChain bool
	AddTransitionData string
	Printchain bool

}

func InitCmd() *CmdParams  {
	return &CmdParams{blc.Init(),false,"",false}
}

/**
	配置区块链参数
	新增交易data
 */
func (cmdParams *CmdParams)Works()   {

	createGenesisBlockFlagCmd := flag.NewFlagSet("createGenesisBlock", flag.ExitOnError)
	addTransitionFlagCmd := flag.NewFlagSet("addTransition", flag.ExitOnError )
	printChainFlagCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	strFlagTransitionData := addTransitionFlagCmd.String("data", "", "交易数据")

	isValidArgs()

	var err error
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "addTransition":
			err = addTransitionFlagCmd.Parse(os.Args[2:])
			if err != nil {
				log.Panic(err)
			}
			break
		case "printchain":
			err = printChainFlagCmd.Parse(os.Args[2:])
			if err != nil {
				log.Panic(err)
			}
			break

		case "createGenesisBlock":
			err = createGenesisBlockFlagCmd.Parse(os.Args[2:])
			if err != nil {
				log.Panic(err)
			}
			break

		default:
			showUsageTips()
			os.Exit(1)

		}
	}
	if addTransitionFlagCmd.Parsed() {
		//allow add new transition
		if *strFlagTransitionData == "" {
			showUsageTips()
			os.Exit(1)
		}

		fmt.Println("输入参数是：", *strFlagTransitionData)
		cmdParams.AddTransitionData = *strFlagTransitionData

	}

	if printChainFlagCmd.Parsed() {
		//allow to print chain info
		cmdParams.Printchain = true
	}

	if createGenesisBlockFlagCmd.Parsed() {
		cmdParams.CreateGenesisBlockChain = true
	}

}


func SetFlag()  {

	fmt.Println("-----cmd command -----")
	flagStrBLC := flag.String("minining", "", "挖矿")
	flagIntSerialize := flag.Int("serialize", 0, "序列化")
	flagBoolDB := flag.Bool("boltdb",false, "数据库")

	//解析
	flag.Parse()

	fmt.Printf("%s\n", *flagStrBLC)
	fmt.Printf("%d\n", *flagIntSerialize)
	fmt.Printf("%v\n", *flagBoolDB)

}

func GetCommandInputArgs()  {
	fmt.Println("-----get args ------")

	args := os.Args;
	fmt.Printf("%v\n",args)
	for i, v := range args{
		fmt.Printf("index %d, %v\n",i, v)
	}

}