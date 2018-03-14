package main
import(
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"html/template"
	"contracts"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"math/big"
	"strings"
	"time"
	"github.com/ethereum/go-ethereum/rpc"
	"os"
	"io/ioutil"
	"github.com/ethereum/go-ethereum/crypto"
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
	//accounts storage
	accs = make(map[string]*bind.TransactOpts)
	//generate transaction opts. Collection of authorization data required to create a valid Ethereum transaction.
	TransOpt=GenTransactOpts("accounts")
	//creating client that represents a connection to an RPC server.
	rpcClient,err := rpc.Dial("http://127.0.0.1:8545")
	if err!=nil{
		log.Fatalln("Dial 49: ",err)
	}
	//NewClient creates a client that uses the given RPC client.
	conn:=ethclient.NewClient(rpcClient)
	//creating and deploying smart contract
	contractAddr,_, contract,err = contracts.DeploySimpleStorage(TransOpt[0],conn)
	fmt.Println("ContractAddress: ",common.ToHex(contractAddr[:]))
	if err!=nil{
		log.Fatalln("Contract deploy 53: ",err)
	}
	//The approximate time for mining block
	time.Sleep(time.Second*10)

}

func main() {
	//run server
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
	req.ParseForm()
	//check that request comes from /userChange
	if len(req.Form)!=1 && len(req.Form)!=0{
		//checks for nationality is set
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
		//checks for visa is set
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
		//checks for age is set
		if req.Form["age"][0]!=dataToTemplate.Age{
			contract.SetAge(&bind.TransactOpts{
				From:     account.From,
				Signer:   account.Signer,
				GasLimit: uint64(2381623),
				Value:    big.NewInt(0),
			}, req.Form["age"][0])
			fmt.Println("Mining Age\n")
		}
		//checks for speaks is set
		if req.Form["speaks"][0]!=dataToTemplate.Speaks{
			contract.SetSpeaks(&bind.TransactOpts{
				From:     account.From,
				Signer:   account.Signer,
				GasLimit: uint64(2381623),
				Value:    big.NewInt(0),
			}, req.Form["speaks"][0])
			fmt.Println("Mining Speaks\n")
		}
		//checks for medConditions is set
		if !testEq(strings.Split(req.Form["medCondition"][0]," "),dataToTemplate.MedCondition)  {
			contract.SetMedCondition(&bind.TransactOpts{
				From:     account.From,
				Signer:   account.Signer,
				GasLimit: uint64(2381623),
				Value:    big.NewInt(0),
			}, req.Form["medCondition"][0])
			fmt.Println("Mining MedicalConditions\n")
		}
		//checks for medMedication is set
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
		//request comes from /index
		var data string
		if len(req.Form)==0{
			data=dataToTemplate.Name
		} else{
			data = req.Form["name"][0]
		}

		if data=="John Doe"{
			account=TransOpt[0]
		}else{
			inputHex := StringToHex(data)
			account=getAddress(inputHex)
		}
	}

	//Get values from contract
	name, _ := contract.GetName(&bind.CallOpts{
		From:account.From,
	})
	//checks that user is new
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
	//fill in User fields   for paste to template
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



func StringToHex(str string) (string){
	bs := []byte(str)
	//Keccak256 hash
	sha3:= [32]byte(crypto.Keccak256Hash(bs[:]))
	return common.Bytes2Hex(sha3[:])
}

//generate transaction opts for private chain
func GenTransactOpts(src string) ([]*bind.TransactOpts){
	var TransOpt []*bind.TransactOpts
	files, err := ioutil.ReadDir(src)
	if err!=nil{
		log.Fatalln("GenTransactOpts() files: ",err)
	}
	//loads key files
	for _,fInfo:=range files{
		file,err:=os.Open(src+"/"+fInfo.Name())
		if err!=nil{
			log.Fatalln("GenTransactOpts() file: ",err)
		}
		trOpt,err:=bind.NewTransactor(file, "1")
		if err!=nil{
			log.Fatalln("GenTransactOpts() trOpt: ",err)
		}
		TransOpt=append(TransOpt,trOpt)
	}
	return TransOpt
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