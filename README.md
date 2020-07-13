# BGW Secure Linear MultiParty Computation
<b> A golang implementation of BGW - mpc</b>

## Brief Introduction
This project is a golang implementation of BGW Secure Linear MultiParty Computation, which supports int and big.Int of golang.

The paper we mainly refer to when implementing this library is 
"Asharov, Gilad , and Y. Lindell . "A Full Proof of the BGW Protocol for Perfectly Secure Multiparty Computation." Journal of Cryptology 30(2015):1-94.". 

## Environment Requirements
- Golang: version 1.9+.

## How to use it
You can simply import our mpc module as a normal Golang package.
```shell
git clone git@github.com:zhengjilai/BGW-SecureLinearMultiPartyComputation.git
mkdir -p $GOPATH/src
cp -r BGW-SecureLinearMultiPartyComputation/loccs.sjtu.edu.cn $GOPATH/src
```

## Structure of this repo
- ```/loccs.sjtu.edu.cn/acrypto/poly``` implements the calculation of polynomial over <i>Zp</i> with single variable 
and a system of solving linear equations over <i>Zp</i>.

- ```/loccs.sjtu.edu.cn/acrypto/secretshare``` implements Shamir's secret sharing scheme over <i>Zp</i>.

- ```/loccs.sjtu.edu.cn/acrypto/mpc``` implements BGW Linear MultiParty Computation, where everyone has an secret <i>x</i><sub>i</sub>, 
and they want to know the output of an linear function <i>f</i>(<i>x</i><sub>1</sub>, <i>x</i><sub>2</sub>, ..., <i>x<sub>n</sub></i>)
while not exposing their own secret (on condition that there are only <i>t</i>&lt;<i>n</i>/2 semi-honest adversaries).

- ```/doc```: Basic documents of this project, including the original paper and our project docs(interfaces, principles and communication analysis).
We also provide an easy explanation of BGW-mpc Multiplication gate, although we have not implemented it.

## Principles and Functions
LinearMultipartyComputation generalize the linear mpc scheme in "Ben-Or M, Goldwasser S, Wigderson A. Completeness theorems for non-cryptographic
fault-tolerant distributed computation. InProceedings of the twentieth annual ACM symposium on Theory of computing 1988 Jan 1 (pp. 1-10). ACM."

In Linear MPC, a given number of participants <i>p</i><sub>1</sub>, <i>p</i><sub>2</sub>, ..., <i>p<sub>n</sub></i>,
each has private data, respectively <i>x</i><sub>1</sub>, <i>x</i><sub>2</sub>, ..., <i>x<sub>n</sub></i>.
Participants want to compute the value of a public function on the private data:
<i>f</i>(<i>x</i><sub>1</sub>, <i>x</i><sub>2</sub>, ..., <i>x<sub>n</sub></i>)
while keeping their own inputs secret, if there are no more than <i>t</i>&lt;<i>n</i>/2 semi-honest
adversaries.

A linear function is in the form <i>f</i>(<i>x</i><sub>1</sub>, <i>x</i><sub>2</sub>, ..., <i>x<sub>n</sub></i>)
= <i>c</i><sub>1</sub><i>x</i><sub>1</sub> + <i>c</i><sub>2</sub><i>x</i><sub>2</sub> + ... +
<i>c<sub>n</sub></i><i>x<sub>n</sub></i>, while <i>c</i><sub>1</sub>, <i>c</i><sub>2</sub>, ..., <i>c<sub>n</sub></i>
are constants.

- Note 1: In the scheme, all element should be in some <i>Zp</i>, i.e. should be non-negative integers.
- Note 2: Each participant has a unique ID, start from 0 to <i>n</i>-1.

## Contributors

- [Haining Lu]()
- [Jilai Zheng](https://github.com/zhengjilai)

Both from Lab of Cryptology and Computer Security, SJTU.
