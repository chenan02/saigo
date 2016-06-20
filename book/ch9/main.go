package main

import ("fmt"; "math")

func distance(x1, y1, x2, y2 float64) float64 {
  a := x2 - x1
  b := y2 - y1
  return math.Sqrt(a*a + b*b)
}

type Shape interface {
  area() float64
  perimeter() float64
}

type Circle struct {
  x, y, r float64
}

type Rectangle struct {
  x1, y1, x2, y2 float64
}

func (c *Circle) area() float64 {
  return math.Pi * c.r*c.r
}

func (r *Rectangle) area() float64 {
  h := distance(r.x1, r.y1, r.x1, r.y2)
  l := distance(r.x1, r.y1, r.x2, r.y1)
  return h*l
}

func (c *Circle) perimeter() float64 { 
  return math.Pi * 2 * c.r
}

func (r *Rectangle) perimeter() float64 {
  h := distance(r.x1, r.y1, r.x1, r.y2)
  l := distance(r.x1, r.y1, r.x2, r.y1)
  return 2*h + 2*l
}

func totalAreaAndPerimeter(shapes ...Shape) (float64, float64) {
  area, perimeter := 0.0, 0.0
  for _, shape := range shapes {
    area += shape.area()
    perimeter += shape.perimeter()
  }
  return area, perimeter
}

func main() {
  c := Circle{1, 1, 3}
  r := Rectangle{0, 0, 4, 4}
  totalArea, totalPerimeter := totalAreaAndPerimeter(&c, &r)
  fmt.Println(totalArea, totalPerimeter)
}

//method = acts on a specific type, function does whatever
//anonymous field is more "is" as opposed to "has", so the type containing the field also has allof the fields (android.talk() as opposed to android.person.talk())
