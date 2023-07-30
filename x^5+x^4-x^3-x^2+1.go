package main

import (
  "image"
  "image/color"
  "image/png"
  "os"
  "math/cmplx"
)

// Returns the valuse of X^4 - 1 for any complex value of X
func f(x complex128) complex128 {
  return cmplx.Pow(x, 5) + cmplx.Pow(x, 4) - cmplx.Pow(x, 3) - cmplx.Pow(x, 2) + 1
}

// Returns the value of the derivative of X^4 - 1 (4X^3) for any complex value of X
func df(x complex128) complex128 {
  return 5 * cmplx.Pow(x, 4) + 4 * cmplx.Pow(x, 3) - 3 * cmplx.Pow(x, 2) - 2 * x
}

// Apply Netwon's Method to the complex function
func newtonsMethod(x complex128) complex128 {
  for i := 0; i < 1000; i++ {
    x -= f(x) / df(x)
  }
  return x
}
// COMPLEX ZEROES
func getColor(p complex128) color.Color {
  d1 := cmplx.Abs(p - -1.44327)
  d2 := cmplx.Abs(p - complex(-0.609585, -0.707177))
  d3 := cmplx.Abs(p - complex(-0.609585, 0.707177))
  d4 := cmplx.Abs(p - complex(0.831219, -0.322384))
  d5 := cmplx.Abs(p - complex(0.831219, 0.322384))

  switch {
  case d1 <= d2 && d1 <= d3 && d1 <= d4 && d1 <= d5:
    return color.RGBA{255, 0, 0, 255} // Red | root d1
  case d2 <= d1 && d2 <= d3 && d2 <= d4 && d2 <= d5:
    return color.RGBA{0, 255, 0, 255} // Green | root d2
  case d3 <= d1 && d3 <= d2 && d3 <= d4 && d3 <= d5:
    return color.RGBA{0, 0, 255, 255} // Blue | root d3
  case d4 <= d1 && d4 <= d2 && d4 <= d3 && d4 <= d5: 
    return color.RGBA{255, 0, 255, 255} // Purple | root d4
  default:
    return color.RGBA{255, 255, 255, 255} // white i think | root d5
  }
}

func NewtonOne() {
  width, height := 1920, 1080
  img := image.NewRGBA(image.Rect(0, 0, width, height))

  for x := 0; x < width; x++ {
    for y := 0; y < height; y++ {
      c := complex(float64(x-width/2), float64(y-height/2))
      val := newtonsMethod(c)
      color := getColor(val)
      img.Set(x, y, color)
    }
  }
  file, err := os.Create("output.png")
  if err != nil {
    panic(err)
  }
  defer file.Close()
  png.Encode(file, img)
}
