## CalcHashの性能
CalcPassHash

どうやってもOpenSSLを呼び出したときにハッシュが合わない。参考値。

CalcPassHash1から4

文字列かByteスライスか。普通に文字列結合すれば良いようなきがする。

```
name             time/op
CalcPassHash-4   17.8ms ± 2%
CalcPassHash1-4  2.39µs ± 1%
CalcPassHash2-4  2.83µs ± 3%
CalcPassHash3-4  2.40µs ± 1%
CalcPassHash4-4  2.26µs ± 1%
```
