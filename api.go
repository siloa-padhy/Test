package main

import (
	"net/http"
	"context"
	"time"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"codetest/utils"
    )
	type HelloHandler struct{}

	var data []byte
	
	var a utils.SampleInput
	var startTime time.Time=time.Now()
	var start time.Time =startTime.Add(1*time.Second)
	    // start:=startTime.Add(1*time.Second)
	var maxtime time.Time=startTime.Add(60 * time.Second)
	var collection []utils.SampleInput
	func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		// var body []byte
		duration := time.Now().Sub(startTime)
		fmt.Println("duration",duration.Seconds())
		fmt.Println("durationmax",maxtime.Second())
		// fmt.Println(maxtime.Second())
		w.Header().Set("Content-Type", "application/json")
		reqBody, _ := ioutil.ReadAll(r.Body)
	    json.Unmarshal(reqBody,&a)
		if int(duration.Seconds())<maxtime.Second(){
			collection=append(collection,a)
		}else{
			fmt.Println("Time out")
			// collection = nil
		}
		fmt.Fprintf(w, "","Trxn Success",collection)
	}
	type WorldHandler struct{}
	func (g *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		// var body []byte
		var b utils.SampleResponse
		w.Header().Set("Content-Type", "application/json")
		// reqBody, _ := ioutil.ReadAll(r.Body)
        sum,Avg,min,max,count:=CalculateTrxn()
		b.Sum=sum
		b.Avg=float64(Avg)
		b.Max=float64(max)
		b.Min=float64(min)
		b.Count=count
		fmt.Fprintf(w, "","Trxn Success",b)
	}
	func CalculateTrxn()(float64,int,int,int,int){
		count:=1
		sum:=0.0
		a:=[]int{}
		for i,_:=range collection{
			sum=sum+collection[i].Amount
			count=count+i
            // if max<collection[i].Amount{
			// 	max=collection[i+1].Amount
				
			// }
			a=append(a, int(collection[i].Amount))
			
		}
		Avg:=int(sum)/count
        min,max:=findMinAndMax(a)
	 return sum,Avg,min,max,count
	}
    func findMinAndMax(a []int) (min int, max int) {
		min = a[0]
		max = a[0]
		for _, value := range a {
			if value < min {
				min = value
			}
			if value > max {
				max = value
			}
		}
		return min, max
	}
	type MyHandler struct{}	
	func (p *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		// var body []byte
		collection=nil
		startTime=time.Now()
		fmt.Fprintf(w, "","Trxndelete Success")
	}
	
func main() {
	fmt.Println("Server Starting...")
	trxn := HelloHandler{}
	trxncal := WorldHandler{}
	delete :=MyHandler{}
	// world := WorldHandler{}
	http.HandleFunc("/transaction", processTimeout(trxn.ServeHTTP, 2*time.Second))
	http.HandleFunc("/response", processTimeout(trxncal.ServeHTTP,10*time.Second))
	http.HandleFunc("/delete", processTimeout(delete.ServeHTTP,5*time.Second))
	// http.Handle("/hello", &hello)
	// http.Handle("/world", &world)
	http.ListenAndServe(":8080", nil)
}

func processTimeout(h http.HandlerFunc, duration time.Duration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), duration)
		defer cancel()

		r = r.WithContext(ctx)

		processDone := make(chan bool)
		go func() {
			h(w, r)
				fmt.Println("Waiting....")
		
			// processDone <- true
		}()

		select {
		case <-ctx.Done():
			w.Write([]byte("Hello world"))
			break
		case <-processDone:
			w.Write([]byte(`{"error": "process timeout"}`))
		
		}
	}
}
// package main

// import (
//     "fmt"
//     "github.com/gorilla/mux"
//     "net/http"
//     "time"
// )

// func rootHandler(w http.ResponseWriter, r *http.Request) {
//     time.Sleep(5 * time.Second)
//     fmt.Fprintf(w, "Hello!")
// 	w.Write([]byte("Hellooooooo!"))
// }

// func main() {
//     mux := mux.NewRouter()
//     mux.HandleFunc("/", rootHandler)

//     muxWithMiddlewares := http.TimeoutHandler(mux, time.Second*3, "Timeout!")

//     http.ListenAndServe(":8080", muxWithMiddlewares)
// }