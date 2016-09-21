package main

import (
	"bytes"
	"fmt"
)

var style string = `
<style type="text/css"> 
  body,table{ 
    font-size:12px; 
  }
  table{
    table-layout:fixed; 
    empty-cells:show; 
    border-collapse: collapse; 
    margin:0 auto; 
  }
  td{
    height:30px; 
  }
  h1,h2,h3{
    font-size:12px; 
    margin:0; 
    padding:0; 
  }
  .table{
    border:1px solid #cad9ea; 
    color:#666; 
  }
  .table th {
    background-repeat:repeat-x; 
    height:30px; 
  }
  .table td,.table th{
    border:1px solid #cad9ea;
    padding:0 1em 0;
  }
  .table tr.alter{
    background-color:#f5fafe;
  }
</style>
`

type tableGen struct {
	buf *bytes.Buffer
}

func NewTableGen() *tableGen {
	var buf bytes.Buffer
	fmt.Fprintln(&buf, style)
	fmt.Fprintln(&buf, "<table class=\"table\"")
	return &tableGen{
		buf: &buf,
	}
}

func (t *tableGen) AddHeader(elems ...string) {
	fmt.Fprintln(t.buf, "  <tr>")
	defer fmt.Fprintln(t.buf, "  </tr>")
	for _, elem := range elems {
		fmt.Fprintln(t.buf, "    <th>", elem, "</th>")
	}
}

func (t *tableGen) AddBody(elems ...string) {
	fmt.Fprintln(t.buf, "  <tr>")
	defer fmt.Fprintln(t.buf, "  </tr>")
	for _, elem := range elems {
		fmt.Fprintln(t.buf, "    <td>", elem, "</td>")
	}
}

func (t *tableGen) Gen() []byte {
	fmt.Fprintln(t.buf, "</table>")
	return t.buf.Bytes()
}
