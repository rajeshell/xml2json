# xml2json
Simple and Generic XML to JSON Convertor


Installation 

go get github.com/rajeshell/xml2json



Example usage: 

package main

import (
	"fmt"

	"github.com/rajeshell/xml2json"
)

var data = []byte(`
   <Processes>
      <Process>
         <CPUTicks>34227</CPUTicks>
         <Flags>48</Flags>
         <MemorySize>2806088</MemorySize>
         <ParentProcess>0</ParentProcess>
         <Priority>20</Priority>
         <ProcessGroup>4</ProcessGroup>
         <ProcessID>4</ProcessID>
         <ProgramName>chrome.exe</ProgramName>
         <State>Blocking</State>
      </Process>
      <Process>
         <CPUTicks>34227</CPUTicks>
         <Flags>48</Flags>
         <MemorySize>2806088</MemorySize>
         <ParentProcess>0</ParentProcess>
         <Priority>20</Priority>
         <ProcessGroup>4</ProcessGroup>
         <ProcessID>4</ProcessID>
         <ProgramName>Kito.exe</ProgramName>
         <State>Blocking</State>
      </Process>
      <Process>
         <CPUTicks>99999</CPUTicks>
         <Flags>48</Flags>
         <MemorySize>2806088</MemorySize>
         <ParentProcess>0</ParentProcess>
         <Priority>20</Priority>
         <ProcessGroup>6</ProcessGroup>
         <ProcessID>6</ProcessID>
         <ProgramName>notepad.exe</ProgramName>
         <State>Blocking</State>
      </Process>

   </Processes>
`)

func main() {

	jsonBytes, err := xml2json.ConvertXML2Json(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonBytes))
}

