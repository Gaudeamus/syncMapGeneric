# syncMapGeneric

It's a proposal based on original sync.Map implementation with the generic approach.

Benchmarks blow show significant improvement in comparison with classic sync.Map with underlying map[any]any storage

None: code comments inside syncmap.go are left from original sync package, just for reference  
## Benchmark with different map implementations
```
goos: darwin
goarch: arm64
pkg: github.com/Gaudeamus/syncMapGeneric
cpu: Apple M1 Pro

BenchmarkLoadMostlyHits
BenchmarkLoadMostlyHits/*syncmap.DeepCopyMap[int,int]
BenchmarkLoadMostlyHits/*syncmap.DeepCopyMap[int,int]-10         	676598252	         1.641 ns/op	       0 B/op	       0 allocs/op
BenchmarkLoadMostlyHits/*syncmap.RWMutexMap[int,int]
BenchmarkLoadMostlyHits/*syncmap.RWMutexMap[int,int]-10          	 7315011	       161.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkLoadMostlyHits/*syncmap.SmartMutexMap[int,int]
BenchmarkLoadMostlyHits/*syncmap.SmartMutexMap[int,int]-10       	 9109879	       162.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkLoadMostlyHits/*syncmap.SyncMapClassic[int,int]
BenchmarkLoadMostlyHits/*syncmap.SyncMapClassic[int,int]-10      	353132542	         3.564 ns/op	       0 B/op	       0 allocs/op
BenchmarkLoadMostlyHits/*syncmap.SyncMap[int,int]
BenchmarkLoadMostlyHits/*syncmap.SyncMap[int,int]-10             	1000000000	         0.9794 ns/op	       0 B/op	       0 allocs/op

BenchmarkLoadMostlyMisses
BenchmarkLoadMostlyMisses/*syncmap.DeepCopyMap[int,int]
BenchmarkLoadMostlyMisses/*syncmap.DeepCopyMap[int,int]-10       	1000000000	         1.173 ns/op	       0 B/op	       0 allocs/op
BenchmarkLoadMostlyMisses/*syncmap.RWMutexMap[int,int]
BenchmarkLoadMostlyMisses/*syncmap.RWMutexMap[int,int]-10        	 7807576	       160.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkLoadMostlyMisses/*syncmap.SmartMutexMap[int,int]
BenchmarkLoadMostlyMisses/*syncmap.SmartMutexMap[int,int]-10     	 7729086	       166.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkLoadMostlyMisses/*syncmap.SyncMapClassic[int,int]
BenchmarkLoadMostlyMisses/*syncmap.SyncMapClassic[int,int]-10    	1000000000	         2.338 ns/op	       0 B/op	       0 allocs/op
BenchmarkLoadMostlyMisses/*syncmap.SyncMap[int,int]
BenchmarkLoadMostlyMisses/*syncmap.SyncMap[int,int]-10           	1000000000	         1.090 ns/op	       0 B/op	       0 allocs/op

BenchmarkLoadOrStoreBalanced
BenchmarkLoadOrStoreBalanced/*syncmap.DeepCopyMap[int,int]
BenchmarkLoadOrStoreBalanced/*syncmap.DeepCopyMap[int,int]                --- SKIP: map_bench_test.go:108: DeepCopyMap has quadratic running time.
BenchmarkLoadOrStoreBalanced/*syncmap.RWMutexMap[int,int]
BenchmarkLoadOrStoreBalanced/*syncmap.RWMutexMap[int,int]-10     	 4357566	       322.1 ns/op	      86 B/op	       1 allocs/op
BenchmarkLoadOrStoreBalanced/*syncmap.SmartMutexMap[int,int]
BenchmarkLoadOrStoreBalanced/*syncmap.SmartMutexMap[int,int]-10  	 4022876	       319.2 ns/op	      92 B/op	       1 allocs/op
BenchmarkLoadOrStoreBalanced/*syncmap.SyncMapClassic[int,int]
BenchmarkLoadOrStoreBalanced/*syncmap.SyncMapClassic[int,int]-10 	 4695938	       281.7 ns/op	      76 B/op	       2 allocs/op
BenchmarkLoadOrStoreBalanced/*syncmap.SyncMap[int,int]
BenchmarkLoadOrStoreBalanced/*syncmap.SyncMap[int,int]-10        	 5907718	       206.3 ns/op	      37 B/op	       1 allocs/op

BenchmarkLoadOrStoreUnique
BenchmarkLoadOrStoreUnique/*syncmap.DeepCopyMap[int,int]
BenchmarkLoadOrStoreUnique/*syncmap.DeepCopyMap[int,int]                  --- SKIP: map_bench_test.go:140: DeepCopyMap has quadratic running time.
BenchmarkLoadOrStoreUnique/*syncmap.RWMutexMap[int,int]
BenchmarkLoadOrStoreUnique/*syncmap.RWMutexMap[int,int]-10       	 2417982	       483.4 ns/op	     150 B/op	       2 allocs/op
BenchmarkLoadOrStoreUnique/*syncmap.SmartMutexMap[int,int]
BenchmarkLoadOrStoreUnique/*syncmap.SmartMutexMap[int,int]-10    	 2694669	       435.0 ns/op	     137 B/op	       2 allocs/op
BenchmarkLoadOrStoreUnique/*syncmap.SyncMapClassic[int,int]
BenchmarkLoadOrStoreUnique/*syncmap.SyncMapClassic[int,int]-10   	 2313794	       581.3 ns/op	     147 B/op	       4 allocs/op
BenchmarkLoadOrStoreUnique/*syncmap.SyncMap[int,int]
BenchmarkLoadOrStoreUnique/*syncmap.SyncMap[int,int]-10          	 3235099	       400.2 ns/op	      71 B/op	       2 allocs/op

BenchmarkLoadOrStoreCollision
BenchmarkLoadOrStoreCollision/*syncmap.DeepCopyMap[int,int]
BenchmarkLoadOrStoreCollision/*syncmap.DeepCopyMap[int,int]-10   	644831940	         1.821 ns/op	       0 B/op	       0 allocs/op
BenchmarkLoadOrStoreCollision/*syncmap.RWMutexMap[int,int]
BenchmarkLoadOrStoreCollision/*syncmap.RWMutexMap[int,int]-10    	 9639828	       126.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkLoadOrStoreCollision/*syncmap.SmartMutexMap[int,int]
BenchmarkLoadOrStoreCollision/*syncmap.SmartMutexMap[int,int]-10 	 7323920	       162.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkLoadOrStoreCollision/*syncmap.SyncMapClassic[int,int]
BenchmarkLoadOrStoreCollision/*syncmap.SyncMapClassic[int,int]-10         	647271153	         3.645 ns/op	       0 B/op	       0 allocs/op
BenchmarkLoadOrStoreCollision/*syncmap.SyncMap[int,int]
BenchmarkLoadOrStoreCollision/*syncmap.SyncMap[int,int]-10                	1000000000	         0.7040 ns/op	       0 B/op	       0 allocs/op

BenchmarkLoadAndDeleteBalanced
BenchmarkLoadAndDeleteBalanced/*syncmap.DeepCopyMap[int,int]
BenchmarkLoadAndDeleteBalanced/*syncmap.DeepCopyMap[int,int]                      --- SKIP: map_bench_test.go:172: DeepCopyMap has quadratic running time.
BenchmarkLoadAndDeleteBalanced/*syncmap.RWMutexMap[int,int]
BenchmarkLoadAndDeleteBalanced/*syncmap.RWMutexMap[int,int]-10            	10117982	       119.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkLoadAndDeleteBalanced/*syncmap.SmartMutexMap[int,int]
BenchmarkLoadAndDeleteBalanced/*syncmap.SmartMutexMap[int,int]-10         	 7671374	       134.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkLoadAndDeleteBalanced/*syncmap.SyncMapClassic[int,int]
BenchmarkLoadAndDeleteBalanced/*syncmap.SyncMapClassic[int,int]-10        	420256539	         5.273 ns/op	       0 B/op	       0 allocs/op
BenchmarkLoadAndDeleteBalanced/*syncmap.SyncMap[int,int]
BenchmarkLoadAndDeleteBalanced/*syncmap.SyncMap[int,int]-10               	881752946	         1.451 ns/op	       0 B/op	       0 allocs/op

BenchmarkLoadAndDeleteUnique
BenchmarkLoadAndDeleteUnique/*syncmap.DeepCopyMap[int,int]
BenchmarkLoadAndDeleteUnique/*syncmap.DeepCopyMap[int,int]                        --- SKIP: map_bench_test.go:200: DeepCopyMap has quadratic running time.
BenchmarkLoadAndDeleteUnique/*syncmap.RWMutexMap[int,int]
BenchmarkLoadAndDeleteUnique/*syncmap.RWMutexMap[int,int]-10              	10072063	       119.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkLoadAndDeleteUnique/*syncmap.SmartMutexMap[int,int]
BenchmarkLoadAndDeleteUnique/*syncmap.SmartMutexMap[int,int]-10           	 8064949	       160.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkLoadAndDeleteUnique/*syncmap.SyncMapClassic[int,int]
BenchmarkLoadAndDeleteUnique/*syncmap.SyncMapClassic[int,int]-10          	1000000000	         1.913 ns/op	       0 B/op	       0 allocs/op
BenchmarkLoadAndDeleteUnique/*syncmap.SyncMap[int,int]
BenchmarkLoadAndDeleteUnique/*syncmap.SyncMap[int,int]-10                 	1000000000	         0.5024 ns/op	       0 B/op	       0 allocs/op

BenchmarkLoadAndDeleteCollision
BenchmarkLoadAndDeleteCollision/*syncmap.DeepCopyMap[int,int]
BenchmarkLoadAndDeleteCollision/*syncmap.DeepCopyMap[int,int]-10          	 6182064	       216.8 ns/op	     100 B/op	       1 allocs/op
BenchmarkLoadAndDeleteCollision/*syncmap.RWMutexMap[int,int]
BenchmarkLoadAndDeleteCollision/*syncmap.RWMutexMap[int,int]-10           	 8927059	       131.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkLoadAndDeleteCollision/*syncmap.SmartMutexMap[int,int]
BenchmarkLoadAndDeleteCollision/*syncmap.SmartMutexMap[int,int]-10        	13885507	        83.92 ns/op	       0 B/op	       0 allocs/op
BenchmarkLoadAndDeleteCollision/*syncmap.SyncMapClassic[int,int]
BenchmarkLoadAndDeleteCollision/*syncmap.SyncMapClassic[int,int]-10       	406676811	         4.317 ns/op	       0 B/op	       0 allocs/op
BenchmarkLoadAndDeleteCollision/*syncmap.SyncMap[int,int]
BenchmarkLoadAndDeleteCollision/*syncmap.SyncMap[int,int]-10              	1000000000	         1.322 ns/op	       0 B/op	       0 allocs/op

BenchmarkRange
BenchmarkRange/*syncmap.DeepCopyMap[int,int]
BenchmarkRange/*syncmap.DeepCopyMap[int,int]-10                           	  964442	      1222 ns/op	       0 B/op	       0 allocs/op
BenchmarkRange/*syncmap.RWMutexMap[int,int]
BenchmarkRange/*syncmap.RWMutexMap[int,int]-10                            	   11365	    144210 ns/op	   18432 B/op	       1 allocs/op
BenchmarkRange/*syncmap.SmartMutexMap[int,int]
BenchmarkRange/*syncmap.SmartMutexMap[int,int]-10                         	    9024	    142348 ns/op	   18432 B/op	       1 allocs/op
BenchmarkRange/*syncmap.SyncMapClassic[int,int]
BenchmarkRange/*syncmap.SyncMapClassic[int,int]-10                        	  931092	      1645 ns/op	       0 B/op	       0 allocs/op
BenchmarkRange/*syncmap.SyncMap[int,int]
BenchmarkRange/*syncmap.SyncMap[int,int]-10                               	  892432	      1237 ns/op	       0 B/op	       0 allocs/op

BenchmarkAdversarialAlloc
BenchmarkAdversarialAlloc/*syncmap.DeepCopyMap[int,int]
BenchmarkAdversarialAlloc/*syncmap.DeepCopyMap[int,int]-10                	 5536736	       222.7 ns/op	     666 B/op	       0 allocs/op
BenchmarkAdversarialAlloc/*syncmap.RWMutexMap[int,int]
BenchmarkAdversarialAlloc/*syncmap.RWMutexMap[int,int]-10                 	14426847	       132.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkAdversarialAlloc/*syncmap.SmartMutexMap[int,int]
BenchmarkAdversarialAlloc/*syncmap.SmartMutexMap[int,int]-10              	12134408	       125.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkAdversarialAlloc/*syncmap.SyncMapClassic[int,int]
BenchmarkAdversarialAlloc/*syncmap.SyncMapClassic[int,int]-10             	 6550972	       195.6 ns/op	      39 B/op	       0 allocs/op
BenchmarkAdversarialAlloc/*syncmap.SyncMap[int,int]
BenchmarkAdversarialAlloc/*syncmap.SyncMap[int,int]-10                    	 9515954	       169.1 ns/op	      27 B/op	       0 allocs/op

BenchmarkAdversarialDelete
BenchmarkAdversarialDelete/*syncmap.DeepCopyMap[int,int]
BenchmarkAdversarialDelete/*syncmap.DeepCopyMap[int,int]-10               	21570382	        57.16 ns/op	     166 B/op	       0 allocs/op
BenchmarkAdversarialDelete/*syncmap.RWMutexMap[int,int]
BenchmarkAdversarialDelete/*syncmap.RWMutexMap[int,int]-10                	10725729	       104.3 ns/op	      17 B/op	       0 allocs/op
BenchmarkAdversarialDelete/*syncmap.SmartMutexMap[int,int]
BenchmarkAdversarialDelete/*syncmap.SmartMutexMap[int,int]-10             	11519174	       112.0 ns/op	      17 B/op	       0 allocs/op
BenchmarkAdversarialDelete/*syncmap.SyncMapClassic[int,int]
BenchmarkAdversarialDelete/*syncmap.SyncMapClassic[int,int]-10            	32158432	        38.65 ns/op	      11 B/op	       0 allocs/op
BenchmarkAdversarialDelete/*syncmap.SyncMap[int,int]
BenchmarkAdversarialDelete/*syncmap.SyncMap[int,int]-10                   	39474223	        33.54 ns/op	       7 B/op	       0 allocs/op

BenchmarkDeleteCollision
BenchmarkDeleteCollision/*syncmap.DeepCopyMap[int,int]
BenchmarkDeleteCollision/*syncmap.DeepCopyMap[int,int]-10                 	 6897621	       159.5 ns/op	      48 B/op	       1 allocs/op
BenchmarkDeleteCollision/*syncmap.RWMutexMap[int,int]
BenchmarkDeleteCollision/*syncmap.RWMutexMap[int,int]-10                  	10799455	       113.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkDeleteCollision/*syncmap.SmartMutexMap[int,int]
BenchmarkDeleteCollision/*syncmap.SmartMutexMap[int,int]-10               	10238922	       114.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkDeleteCollision/*syncmap.SyncMapClassic[int,int]
BenchmarkDeleteCollision/*syncmap.SyncMapClassic[int,int]-10              	704181100	         1.790 ns/op	       0 B/op	       0 allocs/op
BenchmarkDeleteCollision/*syncmap.SyncMap[int,int]
BenchmarkDeleteCollision/*syncmap.SyncMap[int,int]-10                     	1000000000	         0.5685 ns/op	       0 B/op	       0 allocs/op

BenchmarkSwapCollision
BenchmarkSwapCollision/*syncmap.DeepCopyMap[int,int]
BenchmarkSwapCollision/*syncmap.DeepCopyMap[int,int]-10                   	 4324386	       266.8 ns/op	     336 B/op	       2 allocs/op
BenchmarkSwapCollision/*syncmap.RWMutexMap[int,int]
BenchmarkSwapCollision/*syncmap.RWMutexMap[int,int]-10                    	 8623508	       144.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkSwapCollision/*syncmap.SmartMutexMap[int,int]
BenchmarkSwapCollision/*syncmap.SmartMutexMap[int,int]-10                 	 8608393	       139.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkSwapCollision/*syncmap.SyncMapClassic[int,int]
BenchmarkSwapCollision/*syncmap.SyncMapClassic[int,int]-10                	 7252850	       164.1 ns/op	      16 B/op	       1 allocs/op
BenchmarkSwapCollision/*syncmap.SyncMap[int,int]
BenchmarkSwapCollision/*syncmap.SyncMap[int,int]-10                       	12322694	       114.8 ns/op	       8 B/op	       1 allocs/op

BenchmarkSwapMostlyHits
BenchmarkSwapMostlyHits/*syncmap.DeepCopyMap[int,int]
BenchmarkSwapMostlyHits/*syncmap.DeepCopyMap[int,int]-10                  	   58872	     17226 ns/op	   82087 B/op	       4 allocs/op
BenchmarkSwapMostlyHits/*syncmap.RWMutexMap[int,int]
BenchmarkSwapMostlyHits/*syncmap.RWMutexMap[int,int]-10                   	 4812505	       250.5 ns/op	      12 B/op	       1 allocs/op
BenchmarkSwapMostlyHits/*syncmap.SmartMutexMap[int,int]
BenchmarkSwapMostlyHits/*syncmap.SmartMutexMap[int,int]-10                	 6848562	       183.9 ns/op	      12 B/op	       1 allocs/op
BenchmarkSwapMostlyHits/*syncmap.SyncMapClassic[int,int]
BenchmarkSwapMostlyHits/*syncmap.SyncMapClassic[int,int]-10               	51780493	        25.83 ns/op	      28 B/op	       2 allocs/op
BenchmarkSwapMostlyHits/*syncmap.SyncMap[int,int]
BenchmarkSwapMostlyHits/*syncmap.SyncMap[int,int]-10                      	100000000	        12.86 ns/op	       8 B/op	       1 allocs/op

BenchmarkSwapMostlyMisses
BenchmarkSwapMostlyMisses/*syncmap.DeepCopyMap[int,int]
BenchmarkSwapMostlyMisses/*syncmap.DeepCopyMap[int,int]-10                	 2043745	       592.2 ns/op	     836 B/op	       6 allocs/op
BenchmarkSwapMostlyMisses/*syncmap.RWMutexMap[int,int]
BenchmarkSwapMostlyMisses/*syncmap.RWMutexMap[int,int]-10                 	 3645537	       337.6 ns/op	      15 B/op	       1 allocs/op
BenchmarkSwapMostlyMisses/*syncmap.SmartMutexMap[int,int]
BenchmarkSwapMostlyMisses/*syncmap.SmartMutexMap[int,int]-10              	 4405324	       272.8 ns/op	      15 B/op	       1 allocs/op
BenchmarkSwapMostlyMisses/*syncmap.SyncMapClassic[int,int]
BenchmarkSwapMostlyMisses/*syncmap.SyncMapClassic[int,int]-10             	 1950470	       551.1 ns/op	     116 B/op	       5 allocs/op
BenchmarkSwapMostlyMisses/*syncmap.SyncMap[int,int]
BenchmarkSwapMostlyMisses/*syncmap.SyncMap[int,int]-10                    	 2726409	       481.2 ns/op	      75 B/op	       3 allocs/op

BenchmarkCompareAndSwapCollision
BenchmarkCompareAndSwapCollision/*syncmap.DeepCopyMap[int,int]
BenchmarkCompareAndSwapCollision/*syncmap.DeepCopyMap[int,int]-10         	333212389	         3.944 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareAndSwapCollision/*syncmap.RWMutexMap[int,int]
BenchmarkCompareAndSwapCollision/*syncmap.RWMutexMap[int,int]-10          	 8484439	       154.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareAndSwapCollision/*syncmap.SmartMutexMap[int,int]
BenchmarkCompareAndSwapCollision/*syncmap.SmartMutexMap[int,int]-10       	 5353370	       221.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareAndSwapCollision/*syncmap.SyncMapClassic[int,int]
BenchmarkCompareAndSwapCollision/*syncmap.SyncMapClassic[int,int]-10      	119984504	        12.08 ns/op	       1 B/op	       0 allocs/op
BenchmarkCompareAndSwapCollision/*syncmap.SyncMap[int,int]
BenchmarkCompareAndSwapCollision/*syncmap.SyncMap[int,int]-10             	186388846	         6.082 ns/op	       0 B/op	       0 allocs/op

BenchmarkCompareAndSwapNoExistingKey
BenchmarkCompareAndSwapNoExistingKey/*syncmap.DeepCopyMap[int,int]
BenchmarkCompareAndSwapNoExistingKey/*syncmap.DeepCopyMap[int,int]-10     	270082104	         4.397 ns/op	       8 B/op	       1 allocs/op
BenchmarkCompareAndSwapNoExistingKey/*syncmap.RWMutexMap[int,int]
BenchmarkCompareAndSwapNoExistingKey/*syncmap.RWMutexMap[int,int]-10      	 9463696	       124.3 ns/op	       8 B/op	       0 allocs/op
BenchmarkCompareAndSwapNoExistingKey/*syncmap.SmartMutexMap[int,int]
BenchmarkCompareAndSwapNoExistingKey/*syncmap.SmartMutexMap[int,int]-10   	 7988172	       147.5 ns/op	       7 B/op	       0 allocs/op
BenchmarkCompareAndSwapNoExistingKey/*syncmap.SyncMapClassic[int,int]
BenchmarkCompareAndSwapNoExistingKey/*syncmap.SyncMapClassic[int,int]-10  	516478348	         2.227 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareAndSwapNoExistingKey/*syncmap.SyncMap[int,int]
BenchmarkCompareAndSwapNoExistingKey/*syncmap.SyncMap[int,int]-10         	1000000000	         0.5190 ns/op	       0 B/op	       0 allocs/op

BenchmarkCompareAndSwapValueNotEqual
BenchmarkCompareAndSwapValueNotEqual/*syncmap.DeepCopyMap[int,int]
BenchmarkCompareAndSwapValueNotEqual/*syncmap.DeepCopyMap[int,int]-10     	584937136	         2.092 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareAndSwapValueNotEqual/*syncmap.RWMutexMap[int,int]
BenchmarkCompareAndSwapValueNotEqual/*syncmap.RWMutexMap[int,int]-10      	 8972211	       129.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareAndSwapValueNotEqual/*syncmap.SmartMutexMap[int,int]
BenchmarkCompareAndSwapValueNotEqual/*syncmap.SmartMutexMap[int,int]-10   	 7976872	       153.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareAndSwapValueNotEqual/*syncmap.SyncMapClassic[int,int]
BenchmarkCompareAndSwapValueNotEqual/*syncmap.SyncMapClassic[int,int]-10  	582701534	         3.752 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareAndSwapValueNotEqual/*syncmap.SyncMap[int,int]
BenchmarkCompareAndSwapValueNotEqual/*syncmap.SyncMap[int,int]-10         	1000000000	         0.9539 ns/op	       0 B/op	       0 allocs/op

BenchmarkCompareAndSwapMostlyHits
BenchmarkCompareAndSwapMostlyHits/*syncmap.DeepCopyMap[int,int]
BenchmarkCompareAndSwapMostlyHits/*syncmap.DeepCopyMap[int,int]                   --- SKIP: map_bench_test.go:430: DeepCopyMap has quadratic running time.
BenchmarkCompareAndSwapMostlyHits/*syncmap.RWMutexMap[int,int]
BenchmarkCompareAndSwapMostlyHits/*syncmap.RWMutexMap[int,int]-10         	 4415949	       230.6 ns/op	      12 B/op	       1 allocs/op
BenchmarkCompareAndSwapMostlyHits/*syncmap.SmartMutexMap[int,int]
BenchmarkCompareAndSwapMostlyHits/*syncmap.SmartMutexMap[int,int]-10      	 6857208	       176.2 ns/op	      12 B/op	       1 allocs/op
BenchmarkCompareAndSwapMostlyHits/*syncmap.SyncMapClassic[int,int]
BenchmarkCompareAndSwapMostlyHits/*syncmap.SyncMapClassic[int,int]-10     	74525092	        20.28 ns/op	      21 B/op	       1 allocs/op
BenchmarkCompareAndSwapMostlyHits/*syncmap.SyncMap[int,int]
BenchmarkCompareAndSwapMostlyHits/*syncmap.SyncMap[int,int]-10            	127443468	         9.230 ns/op	       8 B/op	       1 allocs/op

BenchmarkCompareAndSwapMostlyMisses
BenchmarkCompareAndSwapMostlyMisses/*syncmap.DeepCopyMap[int,int]
BenchmarkCompareAndSwapMostlyMisses/*syncmap.DeepCopyMap[int,int]-10      	123405423	        12.49 ns/op	      16 B/op	       1 allocs/op
BenchmarkCompareAndSwapMostlyMisses/*syncmap.RWMutexMap[int,int]
BenchmarkCompareAndSwapMostlyMisses/*syncmap.RWMutexMap[int,int]-10       	 8151286	       137.5 ns/op	      15 B/op	       1 allocs/op
BenchmarkCompareAndSwapMostlyMisses/*syncmap.SmartMutexMap[int,int]
BenchmarkCompareAndSwapMostlyMisses/*syncmap.SmartMutexMap[int,int]-10    	 9268621	       136.8 ns/op	      15 B/op	       1 allocs/op
BenchmarkCompareAndSwapMostlyMisses/*syncmap.SyncMapClassic[int,int]
BenchmarkCompareAndSwapMostlyMisses/*syncmap.SyncMapClassic[int,int]-10   	192419268	         6.102 ns/op	       8 B/op	       1 allocs/op
BenchmarkCompareAndSwapMostlyMisses/*syncmap.SyncMap[int,int]
BenchmarkCompareAndSwapMostlyMisses/*syncmap.SyncMap[int,int]-10          	1000000000	         1.068 ns/op	       0 B/op	       0 allocs/op

BenchmarkCompareAndDeleteCollision
BenchmarkCompareAndDeleteCollision/*syncmap.DeepCopyMap[int,int]
BenchmarkCompareAndDeleteCollision/*syncmap.DeepCopyMap[int,int]-10       	711520803	         1.722 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareAndDeleteCollision/*syncmap.RWMutexMap[int,int]
BenchmarkCompareAndDeleteCollision/*syncmap.RWMutexMap[int,int]-10        	 9117541	       133.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareAndDeleteCollision/*syncmap.SmartMutexMap[int,int]
BenchmarkCompareAndDeleteCollision/*syncmap.SmartMutexMap[int,int]-10     	 5827983	       211.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareAndDeleteCollision/*syncmap.SyncMapClassic[int,int]
BenchmarkCompareAndDeleteCollision/*syncmap.SyncMapClassic[int,int]-10    	114252405	        10.91 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareAndDeleteCollision/*syncmap.SyncMap[int,int]
BenchmarkCompareAndDeleteCollision/*syncmap.SyncMap[int,int]-10           	204869083	         5.652 ns/op	       0 B/op	       0 allocs/op

BenchmarkCompareAndDeleteMostlyHits
BenchmarkCompareAndDeleteMostlyHits/*syncmap.DeepCopyMap[int,int]
BenchmarkCompareAndDeleteMostlyHits/*syncmap.DeepCopyMap[int,int]                 --- SKIP: map_bench_test.go:502: DeepCopyMap has quadratic running time.
BenchmarkCompareAndDeleteMostlyHits/*syncmap.RWMutexMap[int,int]
BenchmarkCompareAndDeleteMostlyHits/*syncmap.RWMutexMap[int,int]-10       	 2724265	       375.2 ns/op	      11 B/op	       1 allocs/op
BenchmarkCompareAndDeleteMostlyHits/*syncmap.SmartMutexMap[int,int]
BenchmarkCompareAndDeleteMostlyHits/*syncmap.SmartMutexMap[int,int]-10    	 3793737	       321.2 ns/op	      11 B/op	       1 allocs/op
BenchmarkCompareAndDeleteMostlyHits/*syncmap.SyncMapClassic[int,int]
BenchmarkCompareAndDeleteMostlyHits/*syncmap.SyncMapClassic[int,int]-10   	49992188	        23.49 ns/op	      27 B/op	       2 allocs/op
BenchmarkCompareAndDeleteMostlyHits/*syncmap.SyncMap[int,int]
BenchmarkCompareAndDeleteMostlyHits/*syncmap.SyncMap[int,int]-10          	123700382	         9.654 ns/op	       7 B/op	       0 allocs/op

BenchmarkCompareAndDeleteMostlyMisses
BenchmarkCompareAndDeleteMostlyMisses/*syncmap.DeepCopyMap[int,int]
BenchmarkCompareAndDeleteMostlyMisses/*syncmap.DeepCopyMap[int,int]-10    	731954838	         1.536 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareAndDeleteMostlyMisses/*syncmap.RWMutexMap[int,int]
BenchmarkCompareAndDeleteMostlyMisses/*syncmap.RWMutexMap[int,int]-10     	 9616260	       125.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareAndDeleteMostlyMisses/*syncmap.SmartMutexMap[int,int]
BenchmarkCompareAndDeleteMostlyMisses/*syncmap.SmartMutexMap[int,int]-10  	11705132	        96.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareAndDeleteMostlyMisses/*syncmap.SyncMapClassic[int,int]
BenchmarkCompareAndDeleteMostlyMisses/*syncmap.SyncMapClassic[int,int]-10 	980330413	         2.207 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompareAndDeleteMostlyMisses/*syncmap.SyncMap[int,int]
BenchmarkCompareAndDeleteMostlyMisses/*syncmap.SyncMap[int,int]-10        	984738602	         1.272 ns/op	       0 B/op	       0 allocs/op
PASS
```

## Side to side comparison with classic sync.Map

```
goos: darwin
goarch: arm64
pkg: github.com/Gaudeamus/syncMapGeneric
cpu: Apple M1 Pro
                                                          │   classic     │                generic                │
                                                          │   sync.Map    │                syncMap                │
------------------------------------------------------------------------------------------------------------------│                                                    
                                                          │    sec/op     │    sec/op      vs base                │
LoadMostlyHits/*syncmap.SyncMap[int,int]-10                  3.460n ±  6%    1.025n ± 16%  -70.38% (p=0.000 n=10)
LoadMostlyMisses/*syncmap.SyncMap[int,int]-10                2.264n ±  3%    1.228n ± 11%  -45.77% (p=0.000 n=10)
LoadOrStoreBalanced/*syncmap.SyncMap[int,int]-10             310.1n ±  2%    206.5n ± 29%  -33.40% (p=0.000 n=10)
LoadOrStoreUnique/*syncmap.SyncMap[int,int]-10               562.4n ±  7%    412.5n ±  6%  -26.66% (p=0.000 n=10)
LoadOrStoreCollision/*syncmap.SyncMap[int,int]-10           2.9290n ± 34%   0.7798n ± 22%  -73.38% (p=0.000 n=10)
LoadAndDeleteBalanced/*syncmap.SyncMap[int,int]-10           4.242n ± 35%    1.468n ±  5%  -65.38% (p=0.000 n=10)
LoadAndDeleteUnique/*syncmap.SyncMap[int,int]-10            1.8440n ± 41%   0.5278n ±  1%  -71.37% (p=0.000 n=10)
LoadAndDeleteCollision/*syncmap.SyncMap[int,int]-10          3.856n ± 26%    1.060n ± 13%  -72.51% (p=0.000 n=10)
Range/*syncmap.SyncMap[int,int]-10                           1.331µ ± 23%    1.280µ ±  9%   -3.83% (p=0.025 n=10)
AdversarialAlloc/*syncmap.SyncMap[int,int]-10                199.2n ±  8%    156.9n ±  9%  -21.23% (p=0.001 n=10)
AdversarialDelete/*syncmap.SyncMap[int,int]-10               38.56n ±  9%    34.69n ± 10%  -10.05% (p=0.004 n=10)
DeleteCollision/*syncmap.SyncMap[int,int]-10                1.4440n ± 41%   0.5640n ±  8%  -60.94% (p=0.000 n=10)
SwapCollision/*syncmap.SyncMap[int,int]-10                   165.7n ±  2%    121.8n ±  2%  -26.50% (p=0.000 n=10)
SwapMostlyHits/*syncmap.SyncMap[int,int]-10                  28.73n ± 33%    13.44n ±  7%  -53.24% (p=0.000 n=10)
SwapMostlyMisses/*syncmap.SyncMap[int,int]-10                601.4n ±  7%    520.1n ±  3%  -13.52% (p=0.000 n=10)
CompareAndSwapCollision/*syncmap.SyncMap[int,int]-10        13.885n ± 21%    6.327n ± 10%  -54.43% (p=0.000 n=10)
CompareAndSwapNoExistingKey/*syncmap.SyncMap[int,int]-10    2.3880n ±  3%   0.5291n ±  1%  -77.84% (p=0.000 n=10)
CompareAndSwapValueNotEqual/*syncmap.SyncMap[int,int]-10    3.1185n ± 32%   0.9878n ± 16%  -68.32% (p=0.000 n=10)
CompareAndSwapMostlyHits/*syncmap.SyncMap[int,int]-10        20.18n ±  7%    10.39n ± 14%  -48.51% (p=0.000 n=10)
CompareAndSwapMostlyMisses/*syncmap.SyncMap[int,int]-10      6.510n ± 47%    1.131n ± 27%  -82.62% (p=0.000 n=10)
CompareAndDeleteCollision/*syncmap.SyncMap[int,int]-10      12.320n ± 17%    5.797n ± 10%  -52.95% (p=0.000 n=10)
CompareAndDeleteMostlyHits/*syncmap.SyncMap[int,int]-10      28.18n ± 15%    10.50n ± 18%  -62.73% (p=0.000 n=10)
CompareAndDeleteMostlyMisses/*syncmap.SyncMap[int,int]-10    2.272n ± 39%    1.389n ±  5%  -38.90% (p=0.007 n=10)
geomean                                                      17.76n          8.085n        -54.47%

                                                          │   classic      │                  generic                 │
                                                          │   sync.Map     │                  syncMap                 │
----------------------------------------------------------------------------------------------------------------------│
                                                          │      B/op      │    B/op      vs base                     │
LoadMostlyHits/*syncmap.SyncMap[int,int]-10                  0.000 ±  0%     0.000 ±  0%         ~ (p=1.000 n=10) ¹
LoadMostlyMisses/*syncmap.SyncMap[int,int]-10                0.000 ±  0%     0.000 ±  0%         ~ (p=1.000 n=10) ¹
LoadOrStoreBalanced/*syncmap.SyncMap[int,int]-10             82.00 ±  2%     37.50 ±  7%   -54.27% (p=0.000 n=10)
LoadOrStoreUnique/*syncmap.SyncMap[int,int]-10              155.50 ±  9%     73.00 ±  3%   -53.05% (p=0.000 n=10)
LoadOrStoreCollision/*syncmap.SyncMap[int,int]-10            0.000 ±  0%     0.000 ±  0%         ~ (p=1.000 n=10) ¹
LoadAndDeleteBalanced/*syncmap.SyncMap[int,int]-10           0.000 ±  0%     0.000 ±  0%         ~ (p=1.000 n=10) ¹
LoadAndDeleteUnique/*syncmap.SyncMap[int,int]-10             0.000 ±  0%     0.000 ±  0%         ~ (p=1.000 n=10) ¹
LoadAndDeleteCollision/*syncmap.SyncMap[int,int]-10          0.000 ±  0%     0.000 ±  0%         ~ (p=1.000 n=10) ¹
Range/*syncmap.SyncMap[int,int]-10                           0.000 ±  0%     0.000 ±  0%         ~ (p=1.000 n=10) ¹
AdversarialAlloc/*syncmap.SyncMap[int,int]-10                39.50 ±  1%     27.00 ±  4%   -31.65% (p=0.000 n=10)
AdversarialDelete/*syncmap.SyncMap[int,int]-10              10.000 ± 10%     9.000 ± 11%   -10.00% (p=0.002 n=10)
DeleteCollision/*syncmap.SyncMap[int,int]-10                 0.000 ±  0%     0.000 ±  0%         ~ (p=1.000 n=10) ¹
SwapCollision/*syncmap.SyncMap[int,int]-10                  16.000 ±  0%     8.000 ±  0%   -50.00% (p=0.000 n=10)
SwapMostlyHits/*syncmap.SyncMap[int,int]-10                 28.000 ±  0%     8.000 ±  0%   -71.43% (p=0.000 n=10)
SwapMostlyMisses/*syncmap.SyncMap[int,int]-10               118.50 ±  2%     74.00 ±  3%   -37.55% (p=0.000 n=10)
CompareAndSwapCollision/*syncmap.SyncMap[int,int]-10         1.500 ± 33%     0.000 ±  0%  -100.00% (p=0.000 n=10)
CompareAndSwapNoExistingKey/*syncmap.SyncMap[int,int]-10     0.000 ±  0%     0.000 ±  0%         ~ (p=1.000 n=10) ¹
CompareAndSwapValueNotEqual/*syncmap.SyncMap[int,int]-10     0.000 ±  0%     0.000 ±  0%         ~ (p=1.000 n=10) ¹
CompareAndSwapMostlyHits/*syncmap.SyncMap[int,int]-10       21.000 ±  0%     8.000 ±  0%   -61.90% (p=0.000 n=10)
CompareAndSwapMostlyMisses/*syncmap.SyncMap[int,int]-10      8.000 ±  0%     0.000 ±  0%  -100.00% (p=0.000 n=10)
CompareAndDeleteCollision/*syncmap.SyncMap[int,int]-10       0.000 ±  0%     0.000 ±  0%         ~ (p=1.000 n=10) ¹
CompareAndDeleteMostlyHits/*syncmap.SyncMap[int,int]-10     27.000 ±  0%     7.000 ±  0%   -74.07% (p=0.000 n=10)
CompareAndDeleteMostlyMisses/*syncmap.SyncMap[int,int]-10    0.000 ±  0%     0.000 ±  0%         ~ (p=1.000 n=10) ¹
geomean                                                                  ²                ?                       ² ³
¹ all samples are equal
² summaries must be >0 to compute geomean
³ ratios must be >0 to compute geomean

                                                          │   classic    │                  generic                │
                                                          │   sync.Map   │                  syncMap                │
-------------------------------------------------------------------------------------------------------------------│      
                                                          │  allocs/op   │ allocs/op   vs base                     │
LoadMostlyHits/*syncmap.SyncMap[int,int]-10                 0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
LoadMostlyMisses/*syncmap.SyncMap[int,int]-10               0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
LoadOrStoreBalanced/*syncmap.SyncMap[int,int]-10            2.000 ± 0%     1.000 ± 0%   -50.00% (p=0.000 n=10)
LoadOrStoreUnique/*syncmap.SyncMap[int,int]-10              4.000 ± 0%     2.000 ± 0%   -50.00% (p=0.000 n=10)
LoadOrStoreCollision/*syncmap.SyncMap[int,int]-10           0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
LoadAndDeleteBalanced/*syncmap.SyncMap[int,int]-10          0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
LoadAndDeleteUnique/*syncmap.SyncMap[int,int]-10            0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
LoadAndDeleteCollision/*syncmap.SyncMap[int,int]-10         0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
Range/*syncmap.SyncMap[int,int]-10                          0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AdversarialAlloc/*syncmap.SyncMap[int,int]-10               0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
AdversarialDelete/*syncmap.SyncMap[int,int]-10              0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
DeleteCollision/*syncmap.SyncMap[int,int]-10                0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
SwapCollision/*syncmap.SyncMap[int,int]-10                  1.000 ± 0%     1.000 ± 0%         ~ (p=1.000 n=10) ¹
SwapMostlyHits/*syncmap.SyncMap[int,int]-10                 2.000 ± 0%     1.000 ± 0%   -50.00% (p=0.000 n=10)
SwapMostlyMisses/*syncmap.SyncMap[int,int]-10               5.000 ± 0%     3.000 ± 0%   -40.00% (p=0.000 n=10)
CompareAndSwapCollision/*syncmap.SyncMap[int,int]-10        0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
CompareAndSwapNoExistingKey/*syncmap.SyncMap[int,int]-10    0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
CompareAndSwapValueNotEqual/*syncmap.SyncMap[int,int]-10    0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
CompareAndSwapMostlyHits/*syncmap.SyncMap[int,int]-10       1.000 ± 0%     1.000 ± 0%         ~ (p=1.000 n=10) ¹
CompareAndSwapMostlyMisses/*syncmap.SyncMap[int,int]-10     1.000 ±  ?     0.000 ± 0%  -100.00% (p=0.003 n=10)
CompareAndDeleteCollision/*syncmap.SyncMap[int,int]-10      0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
CompareAndDeleteMostlyHits/*syncmap.SyncMap[int,int]-10     2.000 ± 0%     0.000 ± 0%  -100.00% (p=0.000 n=10)
CompareAndDeleteMostlyMisses/*syncmap.SyncMap[int,int]-10   0.000 ± 0%     0.000 ± 0%         ~ (p=1.000 n=10) ¹
geomean                                                                ²               ?                       ² ³
¹ all samples are equal
² summaries must be >0 to compute geomean
³ ratios must be >0 to compute geomean
```
