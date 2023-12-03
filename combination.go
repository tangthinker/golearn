package main

import (
	"encoding/json"
	"fmt"
)

type image struct {
	iColor
	iLocation
}

type iColor struct {
	red   int
	green int
	blue  int
}

type iLocation struct {
	x int
	y int
}

func (color iColor) printColor() {
	fmt.Println("the image point is red:", color.red, "green:", color.green,
		"blue:", color.blue)
}

func (location iLocation) printLocation() {
	fmt.Println("the image point is in x:", location.x, "y:", location.y)
}

func newImage(color iColor, location iLocation) image {
	return image{color, location}
}

func newColor(red, green, blue int) iColor {
	return iColor{red, green, blue}
}

func newLocation(x, y int) iLocation {
	return iLocation{x, y}
}

func (ima image) printImage() {
	fmt.Println("the image point in x:", ima.x, "y:", ima.y, ",and colored red:", ima.red,
		"green:", ima.green, "blue:", ima.blue, ". It is a beautiful pointer!")
}

//提供String()接口供Println方法调用
func (ima image) String() string {
	return fmt.Sprintf("[location:[x:%v,y:%v],color:[red:%v,green:%v,blue:%v]]",
		ima.x, ima.y, ima.red, ima.green, ima.blue)
}

func (ima image) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, 10)
	b = append(b, 'r', 'g', 'b', 'x', 'y')
	return b, nil
}

func main() {
	ima := newImage(newColor(255, 255, 255), newLocation(222, 333))
	ima.printLocation()
	ima.printColor()
	ima.printImage()
	fmt.Println(ima)
	resultByte, err := json.Marshal(ima)
	fmt.Println("result:", string(resultByte), "error:", err)

}
