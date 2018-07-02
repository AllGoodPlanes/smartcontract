package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>

    <link rel="stylesheet" type="text/css" href="main.css">

    <script src="./node_modules/web3/dist/web3.min.js"></script>

</head>
<body>
    <div class="container">

        <h1>Member Details</h1>

        <h2 id="details"></h2>

        <label for="name" class="col-lg-2 control-label">Member Name</label>
        <input id="name" type="text">

        <label for="name" class="col-lg-2 control-label">Member Age</label>
        <input id="age" type="text">

        <button id="button">Add Details</button>


    </div>

    <script src="https://code.jquery.com/jquery-3.2.1.slim.min.js"></script>

        <script>

        if (typeof web3 !== 'undefined') {
            web3 = new Web3(web3.currentProvider);
        } else {
            // set the provider you want from Web3.providers
            web3 = new Web3(new Web3.providers.HttpProvider("http://localhost:8545"));
        }

web3.eth.defaultAccount = web3.eth.accounts[0];

var MycontractContract = web3.eth.contract([
	{
		"constant": false,
		"inputs": [
			{
				"name": "_Name",
				"type": "string"
			},
			{
				"name": "_Age",
				"type": "uint256"
			}
		],
		"name": "setDetails",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "getDetails",
		"outputs": [
			{
				"name": "",
				"type": "string"
			},
			{
				"name": "",
				"type": "uint256"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	}
]);

 var Mycontract = MycontractContract.at('0xda79e9b6f6cc07ee1372e599cce7085523c06aa7');
        console.log(Mycontract);

Mycontract.getDetails(function(error, result){
            if(!error)
                {
                    $("#details").html(result[0]+' ('+result[1]+' years old)');
                    console.log(result);
                }
            else
                console.error(error);
        });

        $("#button").click(function() {
		Mycontract.setDetails.sendTransaction($("#name").val(), $("#age").val(),{from:web3.eth.accounts[0],gas:40000}, function(error, result){
			    if(!error){
				    console.log(result);
	    } else {
	    console.error(error);
    }
    });
        });

    </script>

</body>
</html>`,
	)
}
func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
