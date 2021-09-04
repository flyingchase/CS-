package main

//func main() {
//	server:=http.Server {
//	    Addr : "localhost:8080",
//	}
//	http.HandleFunc("/header",func(w http.ResponseWriter,r *http.Request){
//		fmt.Fprintln(w,r.Header)
//		fmt.Fprintln(w,r.Header["Accept-Encoding"])
//		fmt.Fprintln(w,r.Header.Get("Accept-Encoding"))
//
//	})
//	server.ListenAndServe()
//}

func twosum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}
