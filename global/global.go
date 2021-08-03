package global

import "awesomeProject/model"

var Gragh *model.Graph
var TeToIndex, StuToIndex, IndexToTe, IndexToStu = make(map[string]int), make(map[string]int), make(map[int]string), make(map[int]string)
var VirtualNode = make(map[string]int)
var InDegree []int
