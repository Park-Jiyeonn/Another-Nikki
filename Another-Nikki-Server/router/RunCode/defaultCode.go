package RunCode

import (
	"Another-Nikki/util"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var rx *rand.Rand

func DefaultCode(c *gin.Context) {
	time.Sleep(time.Millisecond * 100)
	language := c.Query("lang")
	if rx == nil {
		rx = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	if language == "c++" {
		util.SendResp(c, 200, gin.H{
			"code": `#include <bits/stdc++.h>
using namespace std;
#define int long long
const int N = 1e6 + 5, mod = 1e9 + 7;
void solve()
{
	int a, b;
	cin >> a >> b;
	cout << a + b << "\n";
}
signed main()
{
	ios::sync_with_stdio(0);
	cin.tie(0);
	int tt = 1;
	// cin >> tt;
	while (tt--) solve();
	return 0;
}`,
			"input": strconv.Itoa(rx.Intn(10)) + " " + strconv.Itoa(rx.Intn(10)),
		}, "success")
		return
	} else if language == "python" {
		util.SendResp(c, 200, gin.H{
			"code": `s = input().split()
print(int(s[0]) + int(s[1]))`,
			"input": strconv.Itoa(rx.Intn(10)) + " " + strconv.Itoa(rx.Intn(10)),
		}, "success")
		return
	} else if language == "java" {
		util.SendResp(c, 200, gin.H{
			"code": `import java.io.*;
import java.util.*;
public class Main {
    public static void main(String args[]) throws Exception {
        Scanner cin=new Scanner(System.in);
        int a = cin.nextInt(), b = cin.nextInt();
        System.out.println(a+b);
    }
}`,
			"input": strconv.Itoa(rx.Intn(10)) + " " + strconv.Itoa(rx.Intn(10)),
		}, "success")
		return
	} else if language == "go" {
		util.SendResp(c, 200, gin.H{
			"code": `package main

import "fmt"

func main() {
    var a, b int
    fmt.Scanf("%d%d", &a, &b)
    fmt.Println(a+b)
}`,
			"input": strconv.Itoa(rx.Intn(10)) + " " + strconv.Itoa(rx.Intn(10)),
		}, "success")
		return
	}

	util.SendResp(c, 404, "暂不支持这种语言呢～", "error")
}
