package main

import "fmt"

const n = 1e4 + 1
const mod = 1e9 + 7

var f = make([]int, n)
var g = make([]int, n)

func qmi(a, k, p int) int {
	res := 1
	for k != 0 {
		if k&1 == 1 {
			res = res * a % p
		}
		k >>= 1
		a = a * a % p
	}
	return res
}

func init() {
	f[0], g[0] = 1, 1
	for i := 1; i < n; i++ {
		f[i] = f[i-1] * i % mod
		g[i] = qmi(f[i], mod-2, mod)
	}
}

func comb(n, k int) int {
	return (f[n] * g[k] % mod) * g[n-k] % mod
}

func countGoodSubsequences(s string) (ans int) {
	cnt := [26]int{}
	mx := 1
	for _, c := range s {
		cnt[c-'a']++
		mx = max(mx, cnt[c-'a'])
	}
	for i := 1; i <= mx; i++ {
		x := 1
		for _, v := range cnt {
			if v >= i {
				x = (x * (comb(v, i) + 1)) % mod
			}
		}
		ans = (ans + x - 1) % mod
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	res := countGoodSubsequences("aabb")
	fmt.Println(res)
}

-----------------------------------------------------------------------------

function firstUniqueIndex(str) {
    const charCount = new Map();

    for (let i = 0; i < str.length; i++) {
        const char = str[i];
        charCount.set(char, (charCount.get(char) || 0) + 1);
    }

    for (let i = 0; i < str.length; i++) {
        if (charCount.get(str[i]) === 1) {
            return i + 1; // 1-based indexing
        }
    }

    return -1;
}

// Example usage:
const inputString = "leetcode";
const result = firstUniqueIndex(inputString);
console.log(result);

------------------------------------------------------------------------------

select 'customer' as category, c.id as id, customer_name as name
from customer c
left join invoice i on c.id = i.customer_id
where i.id is null 

union

select 'product' as category, p.id as id, product_name as name
from product p
left join invoice_item ii on p.id = ii.product_id
where ii.id is null;

-----------------------------------------------------------------------------
https://stackoverflow.com/questions/73431901/group-by-specific-month-in-sql-for-example-group-by-month-november


WITH cte AS (
    SELECT s.name              AS state_name,
           c.name              AS country_name,
           AVG(w.humidity)     AS avg_humidity,
           AVG(w.temperature)  AS avg_temperature
    FROM       country             c
    INNER JOIN state               s 
            ON c.id = s.country_id   
    INNER JOIN state_weather_stats w
            ON s.id = w.state_id
           AND MONTH(w.record_date) = 11
           AND YEAR(w.record_date) = 2019
    GROUP BY s.name,
             c.name
)
SELECT CONCAT_WS(' ', state_name, 
                      country_name, 
                      avg_humidity,
                      CASE WHEN avg_temperature < 15 THEN 'COLD'         
                           WHEN avg_temperature < 30 THEN 'WARM'
                                                     ELSE 'HOT' END) AS temp
FROM cte;
