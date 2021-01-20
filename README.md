# go-doxx

Basic Doxx Go Script for pulling information through CPF 

🚧 **go-doxx is under development** 🚧
## Examples 

```shell
 $ doxx -cpf 453.178.287-91
```
## Contents

* [Features](#features)
* [Installation](#installation)

## Features

#### Pull cpf data:
* full name
* full address
* mother's name
* etc


### Installation

doxx is installed by running one of the following commands in your terminal. You can install this via the command-line with either `curl`, `wget` or another similar tool.

| Method    | Command                                                                                           |
|:----------|:--------------------------------------------------------------------------------------------------|
| **curl**  | `sh -c "$(curl -fsSL https://raw.githubusercontent.com/TheG0ds/go-doxx/master/install.sh)"` |
| **wget**  | `sh -c "$(wget -O- https://raw.githubusercontent.com/TheG0ds/go-doxx/master/install.sh)"`   |
| **fetch** | `sh -c "$(fetch -o - https://raw.githubusercontent.com/TheG0ds/go-doxx/master/install.sh)"` |

#### Manual inspection

It's a good idea to inspect the install script from projects you don't yet know. You can do
that by downloading the install script first, looking through it so everything looks normal,
then running it:

```shell
wget https://raw.githubusercontent.com/TheG0ds/go-doxx/master/install.sh
sh install.sh
```


