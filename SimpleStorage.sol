pragma solidity ^0.4.17;

contract SimpleStorage{
    
    mapping (address => User) AllUsers;
    struct MainInfo{
            string visa;
            string nationality;
            string age;
            string speaks;
    }
    struct MedInfo{
            string vitalMedicalConditions;
            string medications;
    }
    struct User{
        string      name;
        MainInfo    mainInfo;
        MedInfo     medInfo;
    }
    // init User
    function SimpleStorage(){
        AllUsers[msg.sender]=User("John Doe",MainInfo("1111-1111-1111-1111","American","30","english"),MedInfo("Diabetes Asthma","Zyrtec Aspirin"));
    }
    
    // setters
    function setName(string _value){
        AllUsers[msg.sender].name=_value;
    }
    function setVisa(string _value){
        AllUsers[msg.sender].mainInfo.visa=_value;
    } 
    function setNationality(string _value){
        AllUsers[msg.sender].mainInfo.nationality=_value;
    } 
    function setAge(string _value){
        AllUsers[msg.sender].mainInfo.age=_value;
    } 
    function setSpeaks(string _value){
        AllUsers[msg.sender].mainInfo.speaks=_value;
    } 
    function setMedCondition(string _value){
        AllUsers[msg.sender].medInfo.vitalMedicalConditions=_value;
    }
    function setMedMedication(string _value){
        AllUsers[msg.sender].medInfo.medications=_value;
    }
    
    // getters
    function getName() constant returns(string){
        return AllUsers[msg.sender].name;
    }

    function getVisa() constant returns (string){
        return AllUsers[msg.sender].mainInfo.visa;
    }
    function getNationality() constant returns (string){
        return AllUsers[msg.sender].mainInfo.nationality;
    }
    function getAge() constant returns (string){
        return AllUsers[msg.sender].mainInfo.age;
    }
    function getSpeaks() constant returns (string){
        return AllUsers[msg.sender].mainInfo.speaks;
    } 
    function getMedCondition() constant returns (string){
        return AllUsers[msg.sender].medInfo.vitalMedicalConditions;
    }
    function getMedMedication() constant returns (string){
        return AllUsers[msg.sender].medInfo.medications;
    }
}