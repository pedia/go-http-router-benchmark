Go HTTP Router Benchmark
========================

This benchmark suite aims to compare the performance of HTTP request routers for [Go](https://golang.org) by implementing the routing structure of some real world APIs.
Some of the APIs are slightly adapted, since they can not be implemented 1:1 in some of the routers.

Of course the tested routers can be used for any kind of HTTP request → handler function routing, not only (REST) APIs.


#### Tested routers & frameworks:

 * [Beego](http://beego.me/)
 * [go-json-rest](https://github.com/ant0ine/go-json-rest)
 * [chi](https://github.com/go-chi/chi)
 * [Denco](https://github.com/naoina/denco)
 * [Gocraft Web](https://github.com/gocraft/web)
 * [Goji](https://github.com/zenazn/goji/)
 * [Gorilla Mux](http://www.gorillatoolkit.org/pkg/mux)
 * [http.ServeMux](http://golang.org/pkg/net/http/#ServeMux)
 * [HttpRouter](https://github.com/julienschmidt/httprouter)
 * [HttpTreeMux](https://github.com/dimfeld/httptreemux)
 * [Kocha-urlrouter](https://github.com/naoina/kocha-urlrouter)
 * [Martini](https://github.com/go-martini/martini)
 * [Pat](https://github.com/bmizerany/pat)
 * [Possum](https://github.com/mikespook/possum)
 * [R2router](https://github.com/vanng822/r2router)
 * [TigerTonic](https://github.com/rcrowley/go-tigertonic)
 * [Traffic](https://github.com/pilu/traffic)


## Motivation

Go is a great language for web applications. Since the [default *request multiplexer*](http://golang.org/pkg/net/http/#ServeMux) of Go's net/http package is very simple and limited, an accordingly high number of HTTP request routers exist.

Unfortunately, most of the (early) routers use pretty bad routing algorithms. Moreover, many of them are very wasteful with memory allocations, which can become a problem in a language with Garbage Collection like Go, since every (heap) allocation results in more work for the Garbage Collector.

Lately more and more bloated frameworks pop up, outdoing one another in the number of features. This benchmark tries to measure their overhead.

Beware that we are comparing apples to oranges here, we compare feature-rich frameworks to packages with simple routing functionality only. But since we are only interested in decent request routing, I think this is not entirely unfair. The frameworks are configured to do as little additional work as possible.

If you care about performance, this benchmark can maybe help you find the right router, which scales with your application.

## Results

```text
go test -benchmem -run=^$ -bench ^BenchmarkHandle$ bench

goos: darwin
goarch: amd64
pkg: bench
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkHandle/Chi404-12         	 1759804	       682.1 ns/op	    1153 B/op	       9 allocs/op
BenchmarkHandle/ServeMux404-12    	 1863427	       644.0 ns/op	     848 B/op	       7 allocs/op
BenchmarkHandle/Denco404-12       	 2107784	       577.5 ns/op	     976 B/op	       8 allocs/op
BenchmarkHandle/httprouter404-12  	 1891880	       632.2 ns/op	     848 B/op	       7 allocs/op
BenchmarkHandle/Gorilla404-12     	  616036	      1904 ns/op	     896 B/op	       8 allocs/op
BenchmarkHandle/treemux404-12     	 2476545	       481.1 ns/op	     848 B/op	       7 allocs/op
BenchmarkHandle/treectxmux404-12  	 2489038	       482.8 ns/op	     848 B/op	       7 allocs/op
BenchmarkHandle/ChiPost-12        	 4703595	       254.9 ns/op	     304 B/op	       2 allocs/op
BenchmarkHandle/ServeMuxPost-12   	15012880	        76.64 ns/op	       0 B/op	       0 allocs/op
BenchmarkHandle/DencoPost-12      	92945304	        12.92 ns/op	       0 B/op	       0 allocs/op
BenchmarkHandle/httprouterPost-12 	 1211866	       989.5 ns/op	    1296 B/op	      12 allocs/op
BenchmarkHandle/GorillaPost-12    	 2162367	       554.4 ns/op	     720 B/op	       7 allocs/op
BenchmarkHandle/treemuxPost-12    	33001948	        33.66 ns/op	       0 B/op	       0 allocs/op
BenchmarkHandle/treectxmuxPost-12 	33824539	        35.60 ns/op	       0 B/op	       0 allocs/op
BenchmarkHandle/ChiGet-12         	 4748202	       252.0 ns/op	     304 B/op	       2 allocs/op
BenchmarkHandle/ServeMuxGet-12    	14324800	        80.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkHandle/DencoGet-12       	69499999	        17.19 ns/op	       0 B/op	       0 allocs/op
BenchmarkHandle/httprouterGet-12  	46995631	        25.68 ns/op	       0 B/op	       0 allocs/op
BenchmarkHandle/GorillaGet-12     	 2154531	       555.3 ns/op	     720 B/op	       7 allocs/op
BenchmarkHandle/treemuxGet-12     	40016260	        28.85 ns/op	       0 B/op	       0 allocs/op
BenchmarkHandle/treectxmuxGet-12  	39406746	        30.31 ns/op	       0 B/op	       0 allocs/op
BenchmarkHandle/ChiGet2Arg-12     	 3644430	       331.0 ns/op	     304 B/op	       2 allocs/op
BenchmarkHandle/ServeMuxGet2Arg-12         	 1601828	       768.1 ns/op	     944 B/op	       9 allocs/op
BenchmarkHandle/DencoGet2Arg-12            	 6305853	       185.3 ns/op	     128 B/op	       1 allocs/op
BenchmarkHandle/httprouterGet2Arg-12       	11636846	       104.0 ns/op	      64 B/op	       1 allocs/op
BenchmarkHandle/GorillaGet2Arg-12          	  163254	      7331 ns/op	     897 B/op	       8 allocs/op
BenchmarkHandle/treemuxGet2Arg-12          	 3036121	       394.6 ns/op	     384 B/op	       4 allocs/op
BenchmarkHandle/treectxmuxGet2Arg-12       	 2304343	       522.6 ns/op	     688 B/op	       6 allocs/op
PASS
```