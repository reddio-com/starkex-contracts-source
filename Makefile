default: all
prepare-abi:
	rm -rf tmp
	mkdir -p tmp
	cd tmp && git clone https://github.com/starkware-libs/starkex-contracts.git
	mkdir -p abi
abi: prepare-abi
	cd tmp/starkex-contracts && git checkout StarkExchange-v4.5
	cd tmp && solc-0.6.12 --abi starkex-contracts/scalable-dex/contracts/src/components/TokenRegister.sol  --optimize --overwrite --bin-runtime --base-path ./ --allow-paths `pwd`/starkex-contracts/scalable-dex/contracts/src --output-dir  abi \
	&& cp abi/TokenRegister.abi ../abi
	cd tmp && solc-0.6.12 --abi starkex-contracts/scalable-dex/contracts/src/components/Users.sol  --optimize --overwrite --bin-runtime --base-path ./ --allow-paths `pwd`/starkex-contracts/scalable-dex/contracts/src --output-dir  abi \
	&& cp abi/Users.abi ../abi
	cd tmp && solc-0.6.12 --abi starkex-contracts/scalable-dex/contracts/src/starkex/components/VaultDepositWithdrawal.sol  --optimize --overwrite --bin-runtime --base-path ./ --allow-paths `pwd`/starkex-contracts/scalable-dex/contracts/src --output-dir  abi \
	&& cp abi/VaultDepositWithdrawal.abi ../abi
	cd tmp && solc-0.6.12 --abi starkex-contracts/scalable-dex/contracts/src/interactions/Deposits.sol  --optimize --overwrite --bin-runtime --base-path ./ --allow-paths `pwd`/starkex-contracts/scalable-dex/contracts/src --output-dir  abi \
	&& cp abi/Deposits.abi ../abi
	cd tmp && solc-0.6.12 --abi starkex-contracts/scalable-dex/contracts/src/interactions/Withdrawals.sol  --optimize --overwrite --bin-runtime --base-path ./ --allow-paths `pwd`/starkex-contracts/scalable-dex/contracts/src --output-dir  abi \
	&& cp abi/Withdrawals.abi ../abi
	cd tmp && solc-0.6.12 --abi starkex-contracts/scalable-dex/contracts/src/interactions/AcceptModifications.sol  --optimize --overwrite --bin-runtime --base-path ./ --allow-paths `pwd`/starkex-contracts/scalable-dex/contracts/src --output-dir  abi \
	&& cp abi/AcceptModifications.abi ../abi
	cd tmp && solc-0.6.12 --abi starkex-contracts/scalable-dex/contracts/src/starkex/interactions/UpdateState.sol  --optimize --overwrite --bin-runtime --base-path ./ --allow-paths `pwd`/starkex-contracts/scalable-dex/contracts/src/starkex,`pwd`/starkex-contracts/scalable-dex/contracts/src --output-dir  abi \
	&& cp abi/UpdateState.abi ../abi
	cd tmp && solc-0.6.12 --abi starkex-contracts/scalable-dex/contracts/src/starkex/interactions/FullWithdrawals.sol --optimize --overwrite --bin-runtime --base-path ./ --allow-paths `pwd`/starkex-contracts/scalable-dex/contracts/src/starkex,`pwd`/starkex-contracts/scalable-dex/contracts/src --output-dir  abi \
	&& cp abi/FullWithdrawals.abi ../abi
	cd tmp && solc-0.6.12 --abi starkex-contracts/scalable-dex/contracts/src/starkex/interactions/StarkExForcedActionState.sol --optimize --overwrite --bin-runtime --base-path ./ --allow-paths `pwd`/starkex-contracts/scalable-dex/contracts/src/starkex,`pwd`/starkex-contracts/scalable-dex/contracts/src --output-dir  abi \
	&& cp abi/StarkExForcedActionState.abi ../abi
	rm -rf tmp
prepare-go-souce-dir:
	mkdir -p source/tokenRegister
	mkdir -p source/users
	mkdir -p source/vaultDepositWithdrawal
	mkdir -p source/deposits
	mkdir -p source/withdrawals
	mkdir -p source/acceptModifications
	mkdir -p source/updateState
	mkdir -p source/fullWithdrawals
	mkdir -p source/starkExForcedActionState
go-source: prepare-go-souce-dir
	abigen --abi=abi/TokenRegister.abi --pkg=tokenRegister --out=source/tokenRegister/tokenRegister.go
	abigen --abi=abi/Users.abi --pkg=users --out=source/users/users.go
	abigen --abi=abi/VaultDepositWithdrawal.abi --pkg=vaultDepositWithdrawal --out=source/vaultDepositWithdrawal/VaultDepositWithdrawal.go
	abigen --abi=abi/Deposits.abi --pkg=deposits --out=source/deposits/deposits.go
	abigen --abi=abi/Withdrawals.abi --pkg=withdrawals --out=source/withdrawals/withdrawals.go
	abigen --abi=abi/AcceptModifications.abi --pkg=acceptModifications --out=source/acceptModifications/AcceptModifications.go
	abigen --abi=abi/UpdateState.abi --pkg=updateState --out=source/updateState/updateState.go
	abigen --abi=abi/FullWithdrawals.abi --pkg=fullWithdrawals --out=source/fullWithdrawals/FullWithdrawals.go
	abigen --abi=abi/StarkExForcedActionState.abi --pkg=starkExForcedActionState --out=source/starkExForcedActionState/StarkExForcedActionState.go
all: abi go-source