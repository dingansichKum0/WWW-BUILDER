+++
title = "web3.0"
description = "web3.0与区块链"
date = 2022-07-22
lastmod = 2022-07-25T17:53:59+08:00
tags = ["分享"]
categories = ["balabala"]
draft = false
+++

## 历史回顾 {#历史回顾}


### web1.0 读 {#web1-dot-0-读}

用户是单纯的内容消费者，内容由网站提供，网站让你看什么，你就看什么。用户没有产出内容，就像看报纸一样。
例如：新闻门户网站。


### web2.0 读+写 {#web2-dot-0-读-plus-写}

用户是内容的生产者，网站只是一个向用户提供服务的平台。此时用户虽然产出内容但内容本质还是归平台所有。2.0的互联网只是娱乐、工作、学习的工具。
例如：抖音、微信、b站。

以抖音为例，在web2.0时代用户创造小视频，抖音平台通过视频获得收益，并向用户发放一定（保证用户不跑路情况下的最少）的收益。
在web2.0时代平台为王，用户要像产出获得收入只能向平台打工，或者自己成为平台。

所以web2.0的互联网创业途径也千篇一律，烧钱，熬死对手，一家独大，成为平台。


## web3.0 读+写+拥有 {#web3-dot-0-读-plus-写-plus-拥有}

在web3.0的时代用户将完全生活与互联网中，是生活的一部分。
互联网不仅提供服务，还是一个生活空间，人们的一部分生活可以在网上完成，在互联网中娱乐、工作、学习、消费、交际，创造价值，价值被所有人认可，拥有属于自己而非平台的资产。

以抖音为例，在web3.0时代用户可以向抖音授权自己创造的视频nft，视频收益会按照智能合约公开透明智能执行，发放给用户收益。


### 特征 {#特征}


#### 数字资产 {#数字资产}

用户产出内容完全归用户所有，属于用户的数字资产。比如你是up主，创作的视频都是你的资产，你可以授权给a站，b站，c站。

-   在b站上授权你的视频创造的收入直接打进你的数字钱包
-   去美团点外卖也直接在钱包扣除。
-   买入的nft（游戏装备、电影、音乐等等）也都存入钱包，并可以交易。


#### 只有一个账号（钱包） {#只有一个账号-钱包}

不再像当前在各个平台都有一个账户。web3.0时代只有一个账户（钱包），你的资产，货币都在上面。


#### 安全可信 {#安全可信}

没有人能关停它、串改它。大家都能认可的信用背书，使数字资产的价值受到保障。


#### 去中心化 {#去中心化}

去中心不属于某一巨头。天然的基础设施。


### 区块链是web3.0天然的基础设施 {#区块链是web3-dot-0天然的基础设施}

区块链是一种防篡改、可追溯、共享的分布式账本技术。
web3.0是区块链的应用。

现在主要存在一些基于以太坊开发的web3.0应用。


#### 特点 {#特点}

<!--list-separator-->

-  公开透明

    中心化的情况下，因为银行保证了你账本的真实性，所以无需公开透明；但没有银行时，只有允许账本公开透明，才能使得卖家知道买家账上是否有足够的钱来支付款项

<!--list-separator-->

-  可供所有人查询并对账

    中心化的情况下，银行保证了你账本是真实的；但没有银行时，只有允许其他人都可以查询且对账，才能确保买家账本的真实性，否则买家可能虚报自己的账本

<!--list-separator-->

-  记账权力是公平且安全的

    记账权利是指谁可以有权利修改账本，中心化的情况下，银行只有在你存款了后才会在你的账本上记账，只有银行才有记账权利，但在没有银行时，谁有权利记账呢？
    必须保证修改账本的权利是公平且安全的，否则任何人都可以随意篡改自己的账本，例如自己给自己账本上+1亿元

<!--list-separator-->

-  记录了所有人已确认交易记录的

    已确认交易是为了避免“双重交易”，如果某笔交易没有被确认却被记账了，会导致账本余额数量对不上。例如A账本有1.5万元，需要给B支付1万元，指令已发出但后面交易取消，
    如果此时依然记账，A账本就只有0.5万元，是错误的。中心化的情况下，取消交易后银行会调整账本，但没有银行时，谁有权利调整呢？


#### 智能合约 {#智能合约}

传统合约的数字化版本。目的是提供优于传统合约的安全方法，并减少与合约相关的其他交易成本。

本质就是一个存放在区块链上读写区块的程序。


#### 以太坊与比特币区别 {#以太坊与比特币区别}

<!--list-separator-->

-  比特币

    比特币是价值本身，可以理解为黄金。比特币诞生之初就是货币，为支付存在的。

<!--list-separator-->

-  以太坊

    是一个有偿使用（花费gas）分布式计算的操作系统，可以在上面给各类应用提供计算。


### 开发 {#开发}

web3.0 app是通过智能合约读写区块链。


#### 岗位 {#岗位}

前端开发、前端对应区块链sdk开发、智能合约开发、区块链开发。


#### 前端 {#前端}

前端开发与2.0无异，唯一的区别是与后端交互的变化。3.0时代前端可以通过对应区块链前端的sdk直接与该区块链的智能合约通信。


#### 后端 {#后端}

3.0时代的后端程序不再限于java/golang编写的后端服务，可以是一个个智能合约（Solidity和Vyper编写）。


## 元宇宙 {#元宇宙}

Web3.0是基础设施，元宇宙是上层建筑。
在元宇宙中，AR/VR解决元宇宙前端的技术需要，而Web3.0在后端提供强有力的技术支撑。

例如一些科幻电影《头号玩家》《失控玩家》


## 现状 {#现状}

目前公链根本达不到大吞吐的性能要求，大规模web3.0应用还不现实。

web3.0的定义目前还不完善，法律监管，落地应用都还处于探索阶段。完全实现web3.0将会是触碰多方利益的大洗牌。


## 未来 {#未来}

web3.0虽然还很遥远甚至会有阻碍，但它指明了发展方向。web3.0并不是新型概念，正是由于当下世界处于滞胀没有增长点，急需一场革新来刺激需求，web3.0最近才会如火如荼。
web1.0推动来互联网的发展，2.0推动了移动互联网的发展，3.0或许就是推动AR/VR元宇宙发展的动力。


## [demo](https://github.com/zakudriver/web3.0-example) {#demo}