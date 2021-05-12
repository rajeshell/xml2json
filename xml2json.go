package xml2json

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// var data = []byte(`
//    <Processes>
//       <Process>
//          <CPUTicks>34227</CPUTicks>
//          <Flags>48</Flags>
//          <MemorySize>2806088</MemorySize>
//          <ParentProcess>0</ParentProcess>
//          <Priority>20</Priority>
//          <ProcessGroup>4</ProcessGroup>
//          <ProcessID>4</ProcessID>
//          <ProgramName>ttdskbbrowser</ProgramName>
//          <State>Blocking</State>
//       </Process>
//       <Process>
//          <CPUTicks>34227</CPUTicks>
//          <Flags>48</Flags>
//          <MemorySize>2806088</MemorySize>
//          <ParentProcess>0</ParentProcess>
//          <Priority>20</Priority>
//          <ProcessGroup>4</ProcessGroup>
//          <ProcessID>4</ProcessID>
//          <ProgramName>ttdskbbrowser</ProgramName>
//          <State>Blocking</State>
//       </Process>
//       <Process>
//          <CPUTicks>99999</CPUTicks>
//          <Flags>48</Flags>
//          <MemorySize>2806088</MemorySize>
//          <ParentProcess>0</ParentProcess>
//          <Priority>20</Priority>
//          <ProcessGroup>6</ProcessGroup>
//          <ProcessID>6</ProcessID>
//          <ProgramName>ttdskbbrowser</ProgramName>
//          <State>Blocking</State>
//       </Process>

//    </Processes>
// `)

type Node struct {
	XMLName xml.Name
	Content []byte `xml:",innerxml"`
	Nodes   []Node `xml:",any"`
}

func ConvertXML2Json(xmlData []byte) ([]byte, error) {
	buf := bytes.NewBuffer(xmlData)
	dec := xml.NewDecoder(buf)

	var n Node
	err := dec.Decode(&n)
	if err != nil {
		return []byte(""), err
	}

	obj, err := xml2json(n)
	if err != nil {
		fmt.Println(err.Error())
		return []byte(""), err
	}

	byts, err := json.Marshal(obj)

	return byts, err
}

// func main() {
// 	byts, err := ConvertXML2Json(data)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(string(byts))
// }

func xml2json(pn Node) (interface{}, error) {
	obj := make(map[string]interface{})

	if len(pn.Nodes) == 0 {
		return string(pn.Content), nil
	}

	var objArr []interface{}

	ArrayNode, err := IsNodeArray(pn)
	if err != nil {
		return "", err
	}
	// fmt.Println(pn.XMLName.Local, ArrayNode)
	for _, n := range pn.Nodes {

		if ArrayNode {
			ob, _ := xml2json(n)
			objArr = append(objArr, ob)
		} else {

			obj[n.XMLName.Local], _ = xml2json(n)
		}
	}
	if ArrayNode {
		obj[pn.XMLName.Local] = objArr
		// return()
	}
	return (interface{}(obj)), nil

}

func IsNodeArray(n Node) (bool, error) {
	if len(n.Nodes) <= 1 {
		return false, nil
	}
	NodeName := n.Nodes[0].XMLName.Local

	for i, nd := range n.Nodes {
		if nd.XMLName.Local != NodeName {
			return false, nil
		} else {
			if i >= 1 {
				return true, nil
			}
		}
	}
	return true, nil
}
