# "Lottery" Designs for Ethereum and Hyperledger

## UPDATE:

[4:33 PM 2-22-2019] Added simple lottery premature code for Hyperledger Fabric

## Table of Contents
<!-- TOC -->
- [1. Simple Lottery](#1-simple-lottery)
    - [1.1. Ethereum](#11-ethereum)
    - [1.2. Hyperledger](#12-hyperledger)
    - [1.3. App for Hyperledger Fabric](#13-app-for-hyperledger-fabric)
- [2. Recurring Lottery](#2-recurring-lottery)
    - [2.1. Ethereum](#21-ethereum)
    - [2.2. Hyperledger](#22-hyperledger)
- [3. RNG Lottery](#3-rng-lottery)
    - [3.1. Ethereum](#31-ethereum)
    - [3.2. Hyperledger](#32-hyperledger)
- [4. Powerball](#4-powerball)
    - [4.1. Ethereum](#41-ethereum)
    - [4.2. Hyperledger](#42-hyperledger)
- [Reference](#reference)
<!-- /TOC -->

## 1. Simple Lottery

Designed by: John Kenneth L. Zarate

This type of lottery works by picking the winner from all participants who bought a lottery ticket. Only one game can be initiated and each participants. Players can buy one ticket per transaction. There is no need to bet a number or a combination of numbers in order to join.

### 1.1. Ethereum

Before deploying the smart contract, a duration (in seconds) must be placed first. Once deployed, anyone can buy lottery ticket worth 0.01 ether. After the duration, buying ticket is probihited. Five minutes after the duration, the account who deployed the smart contract can now draw the winner among the participants. The winning participant can withdraw the jackpot which is the sum of how much tickets all participants bought in the game.

### 1.2. Hyperledger

Same with Ethereum, a duration (in seconds) is placed before deploying the chaincode. But unlike the previous one, participants must receive the chaincode and they need to register themselves into it before joining the game. The reason for this is Hyperledger is a private blockchain platform, meaning chaincode is only deployed within specific "walls" where all participants recognize each other.

Once entered, players enter the ticket price when buying ticket price. The rest of the process remains the same as before: after the duration, deployer draws the winner and the winning participant can claim the jackpot prize.

### 1.3. App for Hyperledger Fabric

To run the equivalent app written for Hyperledger Fabric:

1. Clone this repo, make sure all prerequisites are performed and you have "fabric-samples" in your machine.

2. Copy the contents of lottery-design/simple-lottery-hyperledger and paste it to your "fabric-samples" folder.

3. In terminal, go to fabric-samples/lottery and run these commands:

```
./startFabric.sh
npm install
node enrollAdmin.js
node registerUser.js
node app.js
```

4. Open Postman. To use the app:

4.1. To view all ticket entries, set method to GET and go to localhost:3000

4.2. To add a ticket entry, set method to POST, go to localhost:3000/ticket.

    Go to Body tab, tick x-www-form-urlencoded, enter "name" under Key and type any name you want under Value

    Click Send.

4.3. To pick a winner, set method to GET and go to localhost:3000/winner

## 2. Recurring Lottery

Designed by: Jarod Alindogan and Kevin Roy Canete

This type of lottery works by playing the game repeatedly. Recurring lottery has no ending. It will just continue every time a winner has declared. players can buy in every round of the lottery. Even if they lose or win. They could still join the round. 

### 2.1. Ethereum

In the Ethereum platform. Players and Deployers will be divided into two (2) diagram. They both have different flow. The deployer starts with deploying contract after deploying the contract. The role of deployer starts. deployer will specify a duration of a round. After the duration of the round. It will conclude and draw a winner. Checking of round winner is an essential so the deployer will be noticed who won the round. Deployer has an option to delete a round for the lottery.

Players process is the second process. When the contract is deployed. All the players or users can start placing their bets. They could buy a lottery ticket for a round to join a game. They have an option to check their balance. If the player won the round. They can withdraw the price or reward.

### 2.2. Hyperledger

In the Hyperledger platform. There are three diagram, Deployer,orderer and player. Deployer will deploy a chaincode and waits for the approval of the orderer to approve the chaincode and it will be endorse. After the approval, the deployer enrolled as an admin and then start the game. Deployer will specify a duration of a round and draw a winner of the round. It has an option the check the round winner and delete a round. 

Players will receive a web app to start the betting. After receiving, players will enroll as a user and then join the game. After joining, players will start placing bets. They also need to buy a ticket for a round of a lottery to join the game. They could check the balance. If they won the round, they could withdraw the price after the round. They have an option to join the game again and place again their bets for another round. 

## 3. RNG Lottery

Designed by: Chriszer Tamayo

The idea (from the book "Building Games with Ethereum Smart Contracts" Chapter 8) behind the RNG lottery is to use a commit-reveal sequence to create a verifiable random number, Every buyer submits a commitment hash when they buy a ticket.

The commitment is generated by hashing together the user's address and a secret number known only to the user. After ticketing period is over, there is a reveal period during which each player must reveal the secret number used to generate their commitment hash.The secret number is hashed on-chain with the player’s address, and this hash must match the commitment submitted with theticket.

Players who don’t reveal their numbers during the reveal period are dropped from the lottery. The secret numbers are hashed together to generate the random seed for picking a winner.

### 3.1. Ethereum

1. The user must register their address along with secret number and will generate a hash to buy lottery ticket.

2. Buy a lottery ticket using the hash (1 ticket only).

3. There is a duration for buying a ticket, after ticket buying period is over, all participants must reveal their secret number during reveal period or else they will be dropped in the list.

4. After the reveal period is over an admin will draw the winner based on random pick in the lottery.

5. Winner claims the prize.

### 3.2. Hyperledger

Just like in Ethereum the duration of ticket buying and reveal must be initialized before deploying the chaincode, the duration will be based on reaching the required blocked in order to measure the current period, the difference only is that in this lottery game, users know each other except for the secret number owned by each of the participants, when the block reach the reveal period, all users must reveal their secret number or else will be removed from the lottery, after that a chaincode will withdraw the winner based on random pick and the rest of the process is same on Ethereum.

## 4. Powerball

Designed by: John Mico Bangit

Powerball is the most popular type of lottery played in the United States. In this section, we are going to write a contract that ports this game onto Ethereum. 

In Powerball, the user picks six numbers per ticket. The first five numbers are standard numbers from 1–69, and the sixth number is a special Powerball number from 1–26 that offers extra rewards. Every three or four days, a drawing is held, and a winning ticket consisting of five standard numbers and a Powerball number is picked. Prizes are paid out based on the number of winning numbers matched on your ticket.

### 4.1. Ethereum

First, the user must check the Max Number and Max Number of Powerball before he/she buy a ticket, after buying ticket worth 0.02 ether the account who deployed the smart contract will draw the numbers from 1-69 on first to fifth Number and then 1-26 on last number. after the draw , the player can now check the winning combination number. The winning participant can withdraw the jackpot which is Prizes are paid out based on the number of winning numbers matched on your ticket.

### 4.2. Hyperledger

There is same process in game but there are some additional process on technical part. First the deployer must deploy the chaincode and then deploy itself as admin, next on Orderer part he/she must approve and deploy the chaincode. On Player part he/she must register first before entering the game. After the technical part the game may now start. The user must check the Max Number and Max Number of Powerball before he/she buy a ticket, after buying ticket worth 0.02 ether the account who deployed the smart contract will draw the numbers from 1-69 on first to fifth Number and then 1-26 on last number. After the draw, the player can now check the winning combination number. The winning participant can withdraw the jackpot. Prizes are paid out based on the number of winning numbers matched on your ticket.

## Reference

Ethereum reference: Iyer, K. & Dannen, C. (2018). _Building Games with Ethereum Smart Contracts_ (pp. 171-209). Brooklyn, New York: Apress
