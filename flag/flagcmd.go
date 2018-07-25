package flag

import (
	"fmt"
	"os"
	"flag"
	"log"
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
	fmt.Println("\taddTransition -data DATA ----- 新增交易数据.")
	fmt.Println("\tprintchain ----- 是否允许输出区块信息.")

}

type FlagBlcParams struct {
	AddTransitionData string
	Printchain bool

}

/**
	配置区块链参数
	新增交易data
 */
func GetFlagForBlockChain() *FlagBlcParams  {
	//添加交易
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

		default:
			showUsageTips()
			os.Exit(1)

		}
	}

	flagBlcParams := new(FlagBlcParams)

	if addTransitionFlagCmd.Parsed() {
		//allow add new transition
		if *strFlagTransitionData == "" {
			showUsageTips()
			os.Exit(1)
		}
		fmt.Println("输入参数是：", *strFlagTransitionData)
		flagBlcParams.AddTransitionData = *strFlagTransitionData

	}

	if printChainFlagCmd.Parsed() {
		//allow to print chain info
		flagBlcParams.Printchain = true
	}

	return flagBlcParams

}


func SetFlag()  {

	fmt.Println("-----flag command -----")
	flagStrBLC := flag.String("minining", "", "挖矿")
	flagIntSerialize := flag.Int("serialize", 0, "序列化")
	flagBoolDB := flag.Bool("boltdb",false, "数据库")

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