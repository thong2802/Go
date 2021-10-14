package main

import "fmt"

func main() {
	// 1. Arrays


	// Array (mảng) trong Go tương tự các ngôn ngữ khác, tuy nhiên nó có kích thước cố định (fixed size) và các phần tử bên trong phải cùng loại dữ liệu
	//// Khởi tạo một mảng gồm 2 string
	var a [2] string
	// Gán giá trị cho các phần tử trong mảng
	a[0] = "Hi Go"
	a[1] = "Backend"
	
	// In kết quả ra console
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// Khởi tạo một mảng gồm 6 số int và gán luôn giá trị cho nó
	arrays :=[6]int{1,2,3,4,5,6}
	fmt.Println(arrays)

	//// Khởi tạo mảng nhưng không ghi rõ kích thước (thay bằng dấu ba chấm), 
	// trình biên dịch sẽ tự hiểu dựa vào số phần tử đã khai báo
	numbers :=[...]int{44,22,55}
	fmt.Println(numbers)

	//Không giống đa số các ngôn ngữ khác, Array trong Go không phải là dạng tham chiếu (reference types) mà là dạng tham trị (value types). 
	//Khi gán giá trị nó cho một biến mới thì nó sẽ tạo ra một bản copy của Array cũ, 
	//và mọi thay đổi ở Array mới không ảnh hưởng gì đến Array cũ:

	d :=[...]int{1, 2, 3, 4, 5}
	e := d  //// e là một array mới có giá trị giống d

	e[0] = 9

	fmt.Println("d is", d)	// In ra 1 2 3 4 5
	fmt.Println("e is", e)	// In ra 9 2 3 4 5 

	// 2. Slices
	//Slice là một tham chiếu đến Array, nó mô tả một phần (hoặc toàn bộ) Array. 
	//Nó có kích thước động nên thường được sử dụng nhiều hơn Array.

	// Slice có thể tạo ra từ một Array bằng cách cung cấp 2 chỉ số (low và high) xác định vị trí phần tử trong Array. 
	//Ví dụ:

	fmt.Println("===========================================")
	// Khởi tạo Array primes:
	primes := []int{2, 3, 5, 7, 11, 13}

	// Khởi tạo Slice s bằng cách cắt từ phần tử ở vị trí 1 (low) đến phần tử ở vị trí 3 (high - 1) của Array primes
	var s []int = primes[1:4] 

	// In ra giá trị của Slice s
	fmt.Println(s)   // Giá trị của s là [3, 5, 7]

	//Một Slice sẽ có 2 thuộc tính là length (len) và capacity (cap). 
	//Length là số phần tử chứa trong Slice, còn capacity là số phần tử chứa trong Array mà Slice tham chiếu đến (bắt đầu tính từ phần tử đầu tiên của Slice). 
	//Để lấy ra length của Slice ta dùng hàm len(), còn để lấy ra capacity thì ta dùng hàm cap(). Ví dụ:


	s1 := []int{2, 3, 5, 7, 11, 13}

	fmt.Println(s1[0:0])	//s[], len(s) = 0, cap(s) = 6
	fmt.Println(s1[0:4])	//s[2,3,5,7], len(a) = 4, cap(s) = 6
	fmt.Println(s1[2:4])	//s[5,7], len(a) = 2, cap(s) = 4, cap được tính từ vị trí số 2 trở đi
	fmt.Println(s1[2:])		//s[5,7,11,13], len(a) = 4, cap(s) = 6

	// Khi tạo Slice ta có thể bỏ qua các chỉ số low và high, khi đó Go sẽ tự sử dụng giá trị mặc định: 0 cho low và length của Slice cho high. Ví dụ:
	s4 := []int{2, 3, 5, 7, 11, 13}

	s4 = s4[:0]   // s = [0:0]
	s4 = s4[:4]   // s = [0:4]
	s4 = s4[2:]   // s = [2:len(s)] => s = [2:4]
	s4 = s4[:4]   // s = [0:4]

	//Ngoài việc tạo Slice như trên, chúng ta có thể tạo theo các cách sau:

	//Khai báo như một mảng nhưng không chỉ ra kích thước mảng:
//	q := []int{2, 3, 5, 7, 11, 13}

	//Sử dụng hàm make() với công thức: func make([]T, len, cap) []T

//	t := make([]int , 0)		// len(a)=5
//	h := make([]int, 0, 5)	// len(b)=0, cap(b)=5

	// Do Slice chỉ là tham chiếu đến Array, do đó thay đổi giá trị của Slice,
	// sẽ làm thay đổi giá trị của Array mà nó tham chiếu đến. Nếu có nhiều Slice cùng tham chiếu đến một Array thì khi thay đổi giá trị một Slice có thể làm thay đổi giá trị các Slice khác. Ví dụ:

	fmt.Println("===================hi=======================")

	numbers1 := [4]int{1, 2, 3, 4}
	a3 := numbers1[0:2]   // a = [1, 2]
	b3 := numbers1[1:3]   // b = [2, 3]

	b3[0] = 5   // Thay đổi giá trị phần tử đầu tiên của Slice b

	fmt.Println(a3, b3)     // a = [1, 5], b = [5, 3]
	fmt.Println(numbers1)   // numbers = [1, 5, 3, 4]
	

	// 3. Append
	








}