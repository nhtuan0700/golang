# Term
- Command
- Module: module path
- Package
- Functions
- Blank Identify
- Map
- Workspace
- Array / Slice / Slice header
- String


## Commands
- `go mod init`
- `go mod tidy`
- `go get <.|package_name>`
- `go mod edit -replace <path1>=<path2>`
- `go test -v`
- `go run`
- `go build`
- `go install`
- `go work init`
- `go work use`


## Module
Việc đặt tên cho module không ảnh hưởng đến code, nhưng chúng ta cần phải đặt tên đúng ý nghĩa.

Phân biệt `go mod tidy` & `go get .`
- `go mod tidy`: đồng bộ các package được chỉ định trong source (thêm/xoá để đồng bộ)
- `go get .`: fetch and install package được chỉ định trong source


## Package
- Trong 1 thư mục chỉ nên chứa 1 package
- Việc đặt tên package cho thư mục con thì lưu ý phải đặt tên khớp với thư mục con hiện tại đó nếu không go sẽ không thể tìm thấy package


## Blank Identify: 
- for loop, used imports and variables


## Array / Slice / Slice Header: 
### Array
- Hầu như khi code sẽ không thấy nó xuất hiện, bởi vì kích thước cố định của nó
- Mục đích chính là nơi để lưu trữ 1 slice

### Slice: 
Mô tả 1 phần liền kề trong mảng
- `Slice không phải array - Slice mô tả 1 phần của mảng`
- Mỗi phần tử trong slice sẽ tham chiếu đến phần tử tương ứng của mảng
- Mặc dù slice "chứa" 1 con trỏ, nhưng bản thân nó vẫn là 1 giá trị (slice header)

### Slice header:
Key: len, pointer, capacity
- có thể coi nó là header (giá trị) của slice (len, pointer, capacity), và chứa 1 con trỏ đến các phần tử array tương ứng
- `nội dung` bên trong slice có thể thay đổi, nhưng header thì không
- header có thể được copy nhưng array mà nó trỏ tới thì luôn được shared


## String
### What is a string
là 1 read-only slice of byte, bên trong slice đó có thể chứa bất kỳ byte nào,  `indexing a string yields its bytes`, 
- Khi chúng ta thêm 1 ký tự vào 1 chuỗi, có nghĩa là chúng ta sẽ thêm byte đại diện của nó vào chuỗi

### Raw string
là chuỗi được bọc bởi dấu `, nó chỉ có thể chứa text bằng chữ (literal text). 

### Regular string/string: 
nó sẽ được bọc bởi double quote ", ngoài việc nó chứa text bằng chữ (literal string), nó có thể chứa các escape value, byte tuỳ ý

### Literal string: 
nó luôn chứa ký tự UTF-8 hợp lệ, không chứa escape value

### UTF-8: 
source trong Go được defined là UTF-8 text. Điều này cho thấy khi chúng ta viết 1 đoạn text `⌘` trong source code, text editor dùng để lập trình sẽ đặt UTF-8 encoding của ký tự `⌘` vào source text. Khi chúng ta in hexadecimal bytes, chúng ta chỉ cần trích xuất dữ liệu và editor đã đặt ở trong file
- Tóm lại, Go source code là UTF-8, vì vậy source code cho string literal là UTF-8 text. Nếu bên trong string literal không chứa escape value thì chuỗi đó cũng sẽ giữ chính xác source text bên trong dấu ngoặc kép ". Raw string luôn chứa nội dung biểu thị hợp lệ của UTF-8, tương tự nếu literal string không chứa escape thì nó cũng sẽ luôn chứa nội dung biểu thị hợp lệ của UTF-8

### Code points, characters, runes:
#### Code point: 
Unicode standard sử dụng 1 thuật ngữ "code point" để chỉ mục được biểu thị bởi 1 giá trị. Code point U+2318 biểu thị `⌘`

### Character
- Example: "A", "à": đây là character. 
- "à" có code point (U+00E0) nhưng nó lại có cách biểu diễn khác. Ví dụ chúng ta sẽ kết hợp U+0061 với U0300 để tạo thành chữ "à". Vì vậy 1 ký tự có thể được biểu diễn bằng 1 số sequences code point khác nhau, do đó sequences UTF-8 bytes khác nhau

#### Rune:
là 1 short term thay thế cho Unicode code point, 1 alias cho kiểu `int32`, program có thể clear khi 1 gía trị integer đại diện 1 code point. `Character constant` được gọi là `rune constant` trong Go. Type và value của expression '⌘' là kiểu `int` và value `0x2318`
