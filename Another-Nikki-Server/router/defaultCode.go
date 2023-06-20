package router

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var rx *rand.Rand

func DefaultCode(c *gin.Context) {
	language := c.Query("lang")
	if rx == nil {
		rx = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	if language == "c++" {
		c.JSON(http.StatusOK, gin.H{
			"code":
`#include <bits/stdc++.h>
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
		})
		return
	} else if language == "python" {
		c.JSON(http.StatusOK, gin.H{
			"code":
`s = input().split()
print(int(s[0]) + int(s[1]))`,
			"input": strconv.Itoa(rx.Intn(10)) + " " + strconv.Itoa(rx.Intn(10)),
		})
		return
	} else if language == "java" {
		c.JSON(http.StatusOK, gin.H{
			"code":
`import java.io.*;
import java.util.*;
public class Main {
    public static void main(String args[]) throws Exception {
        Scanner cin=new Scanner(System.in);
        int a = cin.nextInt(), b = cin.nextInt();
        System.out.println(a+b);
    }
}`,
			"input": strconv.Itoa(rx.Intn(10)) + " " + strconv.Itoa(rx.Intn(10)),
		})
		return
	} else if language == "go" {
		c.JSON(http.StatusOK, gin.H{
			"code":
`package main

import "fmt"

func main() {
    var a, b int
    fmt.Scanf("%d%d", &a, &b)
    fmt.Println(a+b)
}`,
		"input": strconv.Itoa(rx.Intn(10)) + " " + strconv.Itoa(rx.Intn(10)),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":"error",
		"input": "暂不支持这种语言呢～",
	})
}