package main
import(
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"ether"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"html/template"
	"contracts"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"math/big"
	"strings"
	"time"
	"github.com/ethereum/go-ethereum/rpc"

)


var(
	TransOpt []*bind.TransactOpts
	tpl *template.Template
	contract  *contracts.SimpleStorage
	contractAddr common.Address
	authMain *bind.TransactOpts
	accs map[string]*bind.TransactOpts
	unicAccs []string
	userNumber int
	account *bind.TransactOpts
	dataToTemplate User
)

type User struct {
	Name string
	Visa string
	Nationality string
	Age string
	Speaks string
	MedCondition []string
	MedMedications []string

}

func init(){
	tpl = template.Must(template.ParseGlob("templates/*"))
	accs = make(map[string]*bind.TransactOpts)
	TransOpt=ether.GenTransactOpts("accounts")
	rpcClient,err := rpc.Dial("http://127.0.0.1:8545")
	if err!=nil{
		log.Fatalln("Dial 49: ",err)
	}
	conn:=ethclient.NewClient(rpcClient)

	contractAddr,_, contract,err = contracts.DeploySimpleStorage(TransOpt[0],conn)
	fmt.Println("ContractAddress: ",common.ToHex(contractAddr[:]))
	if err!=nil{
		//bal,_:=conn.PendingBalanceAt(context.TODO(),TransOpt[0].From)
		//fmt.Printf("Ballance of %v: %v \n",common.ToHex(TransOpt[0].From[:]),bal)
		log.Fatalln("Contract deploy 53: ",err)
	}
	time.Sleep(time.Second*10)
	vis, _ := contract.GetVisa(&bind.CallOpts{
		From:TransOpt[0].From,
	})


	fmt.Println("Visa: ",vis)
	fmt.Println(contract)
}


func main() {
	mux := httprouter.New()
	mux.GET("/", index)
	mux.POST("/userData", userData)
	mux.GET("/userChange", userChange)
	http.ListenAndServe(":8080", mux)

}
func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	req.ParseForm()
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func userData(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	vis, _ := contract.GetVisa(&bind.CallOpts{
		From:TransOpt[0].From,
	})
	fmt.Println("VISA is: ",vis)
	req.ParseForm()
	if len(req.Form)!=1 && len(req.Form)!=0{

		if req.Form["nationality"][0]!=dataToTemplate.Nationality {
			fmt.Println("NAtionality!!!!")
			contract.SetNationality(&bind.TransactOpts{
				From:     account.From,
				Signer:   account.Signer,
				GasLimit: uint64(2381623),
				Value:    big.NewInt(0),
			}, req.Form["nationality"][0])
			fmt.Println("Mining Nationality\n")
		}
		//time.Sleep(time.Second*15)
		if req.Form["visa"][0]!=dataToTemplate.Visa {
			fmt.Println("VISA!!!!")
			contract.SetVisa(&bind.TransactOpts{
				From:     account.From,
				Signer:   account.Signer,
				GasLimit: uint64(2381623),
				Value:    big.NewInt(0),
			}, req.Form["visa"][0])
			fmt.Println("Mining Visa\n")

		}

		if req.Form["age"][0]!=dataToTemplate.Age{
			contract.SetAge(&bind.TransactOpts{
				From:     account.From,
				Signer:   account.Signer,
				GasLimit: uint64(2381623),
				Value:    big.NewInt(0),
			}, req.Form["age"][0])
			fmt.Println("Mining Age\n")
		}
		if req.Form["speaks"][0]!=dataToTemplate.Speaks{
			contract.SetSpeaks(&bind.TransactOpts{
				From:     account.From,
				Signer:   account.Signer,
				GasLimit: uint64(2381623),
				Value:    big.NewInt(0),
			}, req.Form["speaks"][0])
			fmt.Println("Mining Speaks\n")
		}
		if !testEq(strings.Split(req.Form["medCondition"][0]," "),dataToTemplate.MedCondition)  {
			contract.SetMedCondition(&bind.TransactOpts{
				From:     account.From,
				Signer:   account.Signer,
				GasLimit: uint64(2381623),
				Value:    big.NewInt(0),
			}, req.Form["medCondition"][0])
			fmt.Println("Mining MedicalConditions\n")
		}
		if !testEq(strings.Split(req.Form["medMedication"][0]," "),dataToTemplate.MedMedications) {
			contract.SetMedMedication(&bind.TransactOpts{
				From:     account.From,
				Signer:   account.Signer,
				GasLimit: uint64(2381623),
				Value:    big.NewInt(0),
			}, req.Form["medMedication"][0])
			fmt.Println("Mining MedicalMedications\n")
		}

	}else{
		var data string
		if len(req.Form)==0{
			data=dataToTemplate.Name
		} else{
			data = req.Form["name"][0]
		}

		if data=="John Doe"{
			account=TransOpt[0]
		}else{
			inputHex := ether.StringToHex(data)
			account=getAddress(inputHex)

		}


	}



	//Get values from contract
	name, _ := contract.GetName(&bind.CallOpts{
		From:account.From,
	})
	if name ==""{
		contract.SetName(&bind.TransactOpts{
			From:     account.From,
			Signer:   account.Signer,
			GasLimit: uint64(2381623),
			Value:    big.NewInt(0),
		}, req.Form["name"][0])
		fmt.Println("Mining Name\n")
		name = req.Form["name"][0]
	}

	visa, _ := contract.GetVisa(&bind.CallOpts{
		From:account.From,
	})
	nationality, _ := contract.GetNationality(&bind.CallOpts{
		From:account.From,
	})
	age, _ := contract.GetAge(&bind.CallOpts{
		From:account.From,
	})
	speaks, _ := contract.GetSpeaks(&bind.CallOpts{
		From:account.From,
	})
	medConditionStr, _ := contract.GetMedCondition(&bind.CallOpts{
		From:account.From,
	})
	medCondition:=strings.Split(medConditionStr," ")
	medMedicationsStr, _:=contract.GetMedMedication(&bind.CallOpts{
		From:account.From,
	})
	medMedications:=strings.Split(medMedicationsStr," ")







	dataToTemplate=User{name, visa,
		nationality,age,
		speaks,medCondition,
		medMedications,}

	err := tpl.ExecuteTemplate(w, "userData.gohtml", dataToTemplate)
	HandleError(w, err)

}

func userChange(w http.ResponseWriter, req *http.Request, params httprouter.Params){
	err := tpl.ExecuteTemplate(w, "userChange.gohtml", dataToTemplate)
	HandleError(w, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln("Handle func error 212: ",err)
	}
}

//getting address for hex string
func getAddress(str string) (*bind.TransactOpts){
	if _,ok:=accs[str];ok{
		return accs[str]
	}
	userNumber+=1
	accs[str]= TransOpt[userNumber]
	return accs[str]
}



func testEq(a, b []string) bool {

	if a == nil && b == nil {
		return true;
	}

	if a == nil || b == nil {
		return false;
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

//key:="b2cb8c8c402ad1fd6bcdb401db3603094e9c0b3f77e90238271d328aa54ec1f7"
//keyBs:=common.Hex2Bytes(key)
//fmt.Println("keyBs===",keyBs)
//keyEc,err:=crypto.ToECDSA(keyBs)
//fmt.Println("keyEc pub: ",keyEc.X)
//if err!=nil{
//	log.Fatalln(err)
//}
//auth := bind.NewKeyedTransactor(keyEc)
//fmt.Println("auth===",auth.From)



//check for ballance
//addr:=common.HexToAddress("0x919fe30eb4677ab9a34ea08960322d0462a872c9")
//fmt.Println("addr===",addr)
//fmt.Println(conn.PendingBalanceAt(context.TODO(),addr))