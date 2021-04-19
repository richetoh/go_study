package unmarshal

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/getveryrichet/go_study/json/common"
)

func ByteToMap() {
	//Decoding json data into Go values
	byt := []byte(`{"num":6.13,"strs":["ab","c"]}`)

	fmt.Println(byt)
	fmt.Println(string(byt))

	//key = string, value= 아무거나 가능하게
	// Json package가 decode된 data를 저장할 곳임
	var dat map[string]interface{}

	//actual decoding
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	//In order to use the values in the decoded map, we'll need to convert them to their appropriate type
	//Here we convert the value in num to the expected float64 type
	num := dat["num"].(float64)
	fmt.Println(num)

	//Accessing nested data requires a series of converions
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)
}

func byteToStruct() {

	//We can also decode JSON into custom data types.
	//This has the advantages of adding additional type-safety to our programs and elimimnating the need for type assertions
	//when accesing the decoded data
	str := `{"page":1, "fruits":["apple", "peach"]}`
	res := common.Book2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

}
func StreamJsonEncoding() {

	//In the example above, we always used bytes and strings as intermediates between the data and JSON representation on standard out
	//We can also stream JSON encodings directoly to os.Writers like os.Stdout or even HTTP response bodies
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lecture": 7}
	enc.Encode(d)

}
