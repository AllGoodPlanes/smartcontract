pragma solidity ^0.4.18;

contract Mycontract {
    
   string name;
   uint age;
   
   function setDetails(string _Name, uint _Age) public {
       name = _Name;
       age = _Age;
   }
   
   function getDetails() public constant returns (string, uint) {
       return (name, age);
   }
    
}
