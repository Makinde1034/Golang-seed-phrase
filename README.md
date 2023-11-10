# Golang seed phrase genrator

## Sample usage
* Run go get github.com/Makinde1034/Golang-seed-phrase.
* Import github.com/Makinde1034/Golang-seed-phrase in your project.
* Call the GeneratorSeedPhrase function passing a int representing the number of words in seed phrase

```
import (
	"fmt"
	"github.com/Makinde1034/quizzup/models"
	seed "github.com/Makinde1034/Golang-seed-phrase"
)

func main(){
  	seedPhrase := seed.GenerateSeedPhrase(20)
	fmt.Println(seedPhrase)
}
```
### Sample response
```
guava coconut rambutan boysenberry cherry guava avocado pineapple
pomegranate apricot orange date nectarine kiwi banana boysenberry
grapefruit grapefruit blackberry orange
```
