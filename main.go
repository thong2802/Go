package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    fmt.Println("Input your name")

    var reader *bufio.Reader= bufio.NewReader(os.Stdin)
    //	name, err := reader.ReadString('\n')
    // ReadString tra ve 2 gia tri String va err, neu khong biet tra ve err ta viet nhu sau
    name, _ := reader.ReadString('\n')

    fmt.Println("Hi: " + name)


}
/*
Biên dịch là giống như bạn soạn thảo ra bản dịch trước.
Thông dịch là người khác nói gì thì bạn dịch ra ngay lập tức.
Tương tự trong lập trình theo mình thì:
Python là ngôn ngữ thông dịch vì khi bạn gõ code xong trong Terminal hay trong IDE thì Python báo lỗi ngay (nếu có lỗi).
Còn các ngôn ngữ khác như Pascal, C hay C++ thì bạn phải compile trước, sau đó nếu có lỗi thì IDE mới báo.
*/