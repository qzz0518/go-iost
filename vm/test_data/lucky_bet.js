class Contract {
    constructor() {
        this.maxUserNumber = 10;
        this.userNumber = 0;
        this.totalCoins = 0;
        this.lastLuckyBlock = -1;
        this.round = 0;
        this.tables = [];
        this.clearUserValue()
    }
    clearUserValue() {
        this.tables = [];
    }
    bet(account, luckyNumber, coins) {
        if (coins < 1 || coins > 5) {
            return "bet coins should be >=1 and <= 5"
        }
        if (luckyNumber < 0 && luckyNumber > 9) {
            return "bet coins should be >=1 and <= 5"
        }

        BlockChain.deposit(account, coins);

        if (this.tables[luckyNumber] === undefined) {
            this.tables[luckyNumber] = [];
        }

        this.tables[luckyNumber].push({ account:account, coins : coins})
        this.userNumber ++;
        this.totalCoins += coins;

        if (this.userNumber >= this.maxUserNumber) {
            let bi = JSON.parse(BlockChain.blockInfo());
            let bn = bi.number;
            let ph = bi.parent_hash;
            if ( this.lastLuckyBlock < 0 || bn - this.lastLuckyBlock >= 16 || bn > this.lastLuckyBlock && ph[ph.length-1] % 16 === 0) {
                this.lastLuckyBlock = bn;

                this.getReward(bn);
                this.userNumber = 0;
                this.totalCoins = 0
            }
        }
    }

    getReward() {
        let ln = this.lastLuckyBlock % 10;
        this.round ++;

        let y = new Int64(100);
        let x = new Int64(95);
        let _tc = new Int64(this.totalCoins);

        let tc = _tc.multi(x).div(y);
        let totalVal = 0;
        let kNum = 0;

        _native_log("lucky number is "+ln);

        this.tables[ln].forEach(function (record) {
            totalVal += record.coins;
            kNum ++;
        });

        let result = {
            number: this.lastLuckyBlock,
            user_number: this.userNumber,
            k_number: kNum,
            total_coins : tc,
            records : [],
            rewards : []
        };

        if (kNum > 0) {
            let unit = tc / totalVal;

            let cache = {};

            this.tables[ln].forEach(function (record) {
                if (cache[record.account] === undefined) {
                    cache[record.account] = [record.coins, 1];
                } else {
                    let c = cache[record.account][0];
                    let t = cache[record.account][1];
                    cache[record.account] = [c + record.coins, t+1];
                }
            });

            for (let account in cache) {
                if (cache.hasOwnProperty(account)) {
                    let reward = cache[account][0] * unit;
                    BlockChain.withdraw(account, reward);
                    result.rewards.push({"account":account, "reward": reward, "times": cache[account][1]})
                }
            }
        }

        this.tables.forEach(function (table) {
            table.forEach(function (record) {
                result.records.push(record)
            })
        });

        this["result"+this.round] = JSON.stringify(result)

    }
}

module.exports = Contract;