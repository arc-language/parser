
struct Point {
    x: int32
    y: int32
}

struct Rectangle {
    topLeft: Point
    bottomRight: Point
}

func add(a: int32, b: int32) int32 {

    return a + b
}

func multiply(x: int32, y: int32) int32 {
    let result: int32 = x * y
    return result
}

func calculate(n: int32) int32 {

    let x: int32 = 10
    let y: int32 = 20
    
    if n > 0 {
        x = x + n
    } else {
        x = x - n
    }
    
    let sum: int32 = add(x, y)

    return sum
}

func createPoint(x: int32, y: int32) Point {

    let p: Point = Point { x: x, y: y }
    return p
}

func main() int32 {

    const PI: float32 = 3.14159
    let count: int32 = 42
    let name: string = "Arc Language"
    let isReady: bool = true
    
    let result: int32 = calculate(count)
    let product: int32 = multiply(5, 7)
    
    let pt: Point = createPoint(10, 20)
    
    return 0
}