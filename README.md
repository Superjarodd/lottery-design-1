# "Lottery" Design Patterns for Ethereum and Hyperledger

## Table of Contents
<!-- TOC -->
- [1. Simple Lottery](#1-simple-lottery)
    - [1.1. Ethereum](#11-ethereum)
    - [1.2. Hyperledger](#12-hyperledger)
- [2. Recurring Lottery](#2-recurring-lottery)
    - [2.1. Ethereum](#21-ethereum)
    - [2.2. Hyperledger](#22-hyperledger)
- [3. RNG Lottery](#3-rng-lottery)
    - [3.1. Ethereum](#31-ethereum)
    - [3.2. Hyperledger](#32-hyperledger)
- [4. Powerball](#4-powerball)
    - [4.1. Ethereum](#41-ethereum)
    - [4.2. Hyperledger](#42-hyperledger)
<!-- /TOC -->

## 1. Simple Lottery

Designed by: John Kenneth L. Zarate

This type of lottery works by picking the winner from all participants who bought a lottery ticket. Only one game can be initiated and each participants. Players can buy one ticket per transaction. There is no need to bet a number or a combination of numbers in order to join.

### 1.1. Ethereum

Before deploying the smart contract, a duration (in seconds) must be placed first. Once deployed, anyone can buy lottery ticket worth 0.01 ether. After the duration, buying ticket is probihited. Five minutes after the duration, the account who deployed the smart contract can now draw the winner among the participants. The winning participant can withdraw the jackpot which is the sum of how much tickets all participants bought in the game.

### 1.2. Hyperledger

Same with Ethereum, a duration (in seconds) is placed before deploying the chaincode. But unlike the previous one, participants must receive the chaincode and they need to register themselves into it before joining the game. The reason for this is Hyperledger is a private blockchain platform, meaning chaincode is only deployed within specific "walls" where all participants recognize each other.

Once entered, players enter the ticket price when buying ticket price. The rest of the process remains the same as before: after the duration, deployer draws the winner and the winning participant can claim the jackpot prize.

## 2. Recurring Lottery

Designed by: Jarod Alindogan

### 2.1. Ethereum

### 2.2. Hyperledger

## 3. RNG Lottery

Designed by: Chriszer Tamayo

### 3.1. Ethereum

### 3.2. Hyperledger

## 4. Powerball

Designed by: John Mico Bangit

### 4.1. Ethereum

### 4.2. Hyperledger
