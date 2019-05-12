# Go Patterns
http://tmrts.com/go-patterns/ の写経

## Creation Patterns
### Builder
- オブジェクトの生成過程を抽象化する
```
type Builder interface {
	Paint(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Build() Interface
}
```
- いろいろなオブジェクトを生成する場合に便利
- パラメータが増えてくると苦しい

### Factory Pattern
- 動的に(この例だとswitch文で)オブジェクト生成し、生成するオブジェクトごとにファクトリ(newHogeHoge)を準備する
```
	switch t {
	case MemoryStorage:
		return newMemoryStorage()
	case DiskStorage:
		return newDiskStorage()
	default:
		return newTempStorage()

```
- interfaceを満たすstructをいくつか用意し、それに対応するファクトリを用意しておく


### Object Pool
- オブジェクトの初期化に時間がかかるオブジェクトはあらかじめ生成して、chanのキューにいれておく
- 使い終わったら戻す

### Singleton
- Global変数で定義してsync.Doで初期化関数を実行