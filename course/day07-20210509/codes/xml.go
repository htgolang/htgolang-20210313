package main

// encoding/xml
// xml
//
import (
	"encoding/xml"
	"fmt"
)

type CDATA struct {
	Content string `xml:",cdata"`
}

type Scores struct {
	Score []float64 `xml:"score"`
}
type Name struct {
	Type  string `xml:"type,attr"`
	Value CDATA
}

type User struct {
	XMLName  xml.Name `xml:"User"`
	Id       int      `xml:"ID,attr"`
	Name     Name     `xml:"name"`
	Password string   `xml:"-"`
	Scores   Scores   `xml:"scores"`
}

func main() {
	var u = User{Id: 100, Name: Name{"nickname", CDATA{"kk"}}, Scores: Scores{Score: []float64{1, 43, 444}}}
	b, err := xml.MarshalIndent(u, "", "\t")
	fmt.Println(err, string(b))

	// users := []User{
	// 	{Id: 100, Name: "kk", Scores: []float64{1, 43, 444}},
	// 	{Id: 101, Name: "kk1", Scores: []float64{1, 43, 444}},
	// 	{Id: 102, Name: "kk2", Scores: []float64{1, 43, 444}},
	// 	{Id: 103, Name: "kk3", Scores: []float64{1, 43, 444}},
	// }

	// b, err = xml.MarshalIndent(users, "", "\t")
	// fmt.Println(err, string(b))

	// txt := `
	// <User>
	//     <Id>101</Id>
	//     <Name>kk1</Name>
	//     <Scores>1</Scores>
	//     <Scores>43</Scores>
	//     <Scores>444</Scores>
	// </User>
	// `
	// var us User
	// err = xml.Unmarshal([]byte(txt), &us)
	// fmt.Println(err, us)

}
