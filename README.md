<p align="center">
<img src="./assets/poro.png" width="300" height="300">
</p>
<p>
    <h3 align="center">League of ASCII</h3>
    <p align="center">League of Legends CLI Artist</p>
</p>

## Features
### Negative Effect
If you want to create a negative image, you could do the following below:

```go
package main

import (
	"github.com/bloiseleo/leagueofascii/leagueofascii"
	"github.com/bloiseleo/leagueofascii/leagueofascii/helpers"
)

func main() {
	img, err := helpers.ReadJpg("./assets/poro.jpg")
	if err != nil {
		panic(err)
	}
	err = leagueofascii.GenerateNegativeImage(img, "./results/poro_negated.jpg", leagueofascii.Best_Quality)
	if err != nil {
		panic(err)
	}
}
```
### Graysacle Effect
If you want some shades of gray, you could do the following below:
```go
package main

import (
	"github.com/bloiseleo/leagueofascii/leagueofascii"
	"github.com/bloiseleo/leagueofascii/leagueofascii/helpers"
)

func main() {
	img, err := helpers.ReadJpg("./assets/poro.jpg")
	if err != nil {
		panic(err)
	}
	err = leagueofascii.GrayScale(img, "./results/poro_grayscale.jpg")
	if err != nil {
		panic(err)
	}
}

```
### Image to ASCII
If you want to transform your image to an ASCII art, you could do the following below:
```go
package main

import (
	"github.com/bloiseleo/leagueofascii/leagueofascii"
	"github.com/bloiseleo/leagueofascii/leagueofascii/helpers"
)

func main() {
	img, err := helpers.ReadJpg("./assets/poro_videogame.jpeg")
	if err != nil {
		panic(err)
	}
	art := leagueofascii.CreateAscII(img)
	art.Render()
}

```
Besides, you can also resize the image and create an ASCII art.
```go
package main

import (
	"github.com/bloiseleo/leagueofascii/leagueofascii"
	"github.com/bloiseleo/leagueofascii/leagueofascii/helpers"
)

func main() {
	img, err := helpers.ReadJpg("./assets/poro_videogame.jpeg")
	if err != nil {
		panic(err)
	}
	art := leagueofascii.CreateAscIIAndResize(img, 100, 100)
	art.Render()
}
```
## Useful Links
- [What are Premultiplied Alpha colors?](https://shawnhargreaves.com/blog/premultiplied-alpha.html)
- [How to manipulate an image?](https://medium.com/@shubham0473/from-pixels-to-pictures-a-guide-to-image-manipulation-in-java-3647cac29ca3)
