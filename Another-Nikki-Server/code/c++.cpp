#include<bits/stdc++.h>
using namespace std;
#define int long long
const int N = 1e6 + 5, mod = 1e9 + 7;
int a[N];
void solve()
{
    cout << "Hello Docker, you can make a online judge !!!\n";
    int n = 1e6;
    a[0] = 1;
    for (int i = 1; i <= n; i++) {
        a[i] = a[i - 1] * i % mod;
    }
}
signed main()
{
    auto start = clock(); 

    ios::sync_with_stdio(0);
    cin.tie(0);
    int tt = 1;
    // cin >> tt;
    while(tt--) solve();

    auto end = clock();
    cout << double(end - start) << " ms" << "\n";
    
    return 0;
}