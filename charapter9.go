//广义优先搜索算法
package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var tu = map[string][]string{
	"v0": []string{"v1", "v2", "v3"},
	"v1": []string{"v0", "v2", "v3", "v4"},
	"v2": []string{"v0", "v2", "v6"},
	"v3": []string{"v0", "v1", "v5"},
	"v4": []string{"v1"},
	"v5": []string{"v3", "v6"},
	"v6": []string{"v2", "v5"},
}

var maze = [][]int{
	[]int{0, 1, 0, 0, 0},
	[]int{0, 1, 0, 1, 0},
	[]int{0, 0, 0, 0, 0},
	[]int{0, 1, 1, 1, 0},
	[]int{0, 0, 0, 1, 0},
}

func main() {
	exist, path := baseTu("v0", "v6")
	// exist, path := shortRoad([]int{0, 2}, []int{4, 4})
	// exist, path := WordLadder("cog", "cog")
	if exist {
		fmt.Println(path)
	} else {
		fmt.Println("the shortest path does not exists...")
	}
}

/*
广度(宽度)优先搜索，见http://blog.csdn.net/raphealguo/article/details/7523411
个人总结有几点需要注意：
1：灰色的是即将被辐射的元素，灰色的是队列(我这里用数组实现)，每次取队列第一个元素，将此元素扔到黑名单中
2:黑色的是已经辐射(检查)过的元素，每个元素只辐射一次，不能重复辐射。所以需要检查该元素是否被辐射过，为了查找方便，用map来存储黑色元素
3：为了记录最终的辐射路径，这里还是采用map来记录，map的key为某个元素b，如果b元素是从元素a辐射过来的，则map[b] = a,即是从哪个元素辐射过来达到a的那个元素
4：辐射路径的这个map是不能重复改写的，即map中此key只有在不存在的时候才能被写，如果此key已经存在，则不能被改写，否则路径就错了。
**/

func baseTu(vs, vd string) (bool, []string) {
	//将第一个元素入灰色队列
	if vs == vd {
		return true, []string{vs}
	}
	gray := []string{}
	pathRecord := make(map[string]string)
	gray = append(gray, vs)
	for len(gray) > 0 && gray[0] != vd {
		toB := gray[0]
		if toB == vd {
			break
		}
		gray = gray[1:]
		for i := 0; i < len(tu[toB]); i++ {
			if _, ok := pathRecord[tu[toB][i]]; !ok {
				gray = append(gray, tu[toB][i])
				if _, ok2 := pathRecord[tu[toB][i]]; !ok2 {
					pathRecord[tu[toB][i]] = toB //这里记录路径，key为节点，val为从哪个点走到这个节点，要先判断pathRecord里面是否存在此key，存在则不能变，否则路径会被覆盖。。。
				}
			}
		}
	}
	if len(gray) == 0 || gray[0] != vd {
		return false, []string{}
	}
	//从pathRecord查找路径，查找顺序为从vd倒着往vs查找。
	pathArray := []string{vd}
	endPoint := vd
	for pathRecord[endPoint] != vs {

		pathArray = append(pathArray, pathRecord[endPoint])
		endPoint = pathRecord[endPoint]
	}
	pathArray = append(pathArray, vs)
	//上面的路径是从vd到vs来查找的，所以还需要反序一下
	for i := 0; i < len(pathArray)/2; i++ {
		pathArray[i], pathArray[len(pathArray)-1-i] = pathArray[len(pathArray)-1-i], pathArray[i]
	}
	return true, pathArray

}

/*
Given two words (start and end), and a dictionary, find the length of shortest transformation sequence
from start to end, such that:
• Only one letter can be changed at a time
• Each intermediate word must exist in the dictionary
For example, Given:
start = "hit"
end = "cog"
dict = ["hot","dot","dog","lot","log"]
As one shortest transformation is "hit" -> "hot" -> "dot" -> "dog" -> "cog", return its length 5.
Note:
• Return 0 if there is no such transformation sequence.
• All words have the same length.
• All words contain only lowercase alphabetic characters.
**/

var dictArray = []string{"hot", "dot", "dog", "lot", "log"}

var dict = map[string][]string{
	"hot": []string{"dot", "lot"},
	"dot": []string{"hot", "dog", "lot"},
	"dog": []string{"dot", "log"},
	"lot": []string{"hot", "dot", "log"},
	"log": []string{"dog", "lot"},
}

//"hit", "cog"
func WordLadder(vs, vd string) (bool, []string) {
	if vs == vd {
		return true, []string{vs}
	}
	//先建立V和E(点和边)的关系
	addVnode(vs)
	addVnode(vd)
	gray := []string{}
	// black := make(map[string]string)
	pathRecord := map[string]string{
		vs: "",
	}

	gray = append(gray, vs)
	for len(gray) > 0 && gray[0] != vd {
		toB := gray[0]
		for _, v := range dict[toB] {
			if _, ok := pathRecord[v]; !ok {
				pathRecord[v] = toB
				gray = append(gray, v)
			}
		}
		gray = gray[1:]
	}

	if len(gray) == 0 || gray[0] != vd {
		return false, []string{}
	}
	last := vd
	pathArray := []string{vd}
	for pathRecord[last] != vs {
		pathArray = append(pathArray, pathRecord[last])
		last = pathRecord[last]
	}
	pathArray = append(pathArray, vs)
	for i := 0; i < len(pathArray)/2; i++ {
		pathArray[i], pathArray[len(pathArray)-1-i] = pathArray[len(pathArray)-1-i], pathArray[i]
	}
	return true, pathArray
}

func addVnode(vs string) {
	if _, ok := dict[vs]; !ok {
		dict[vs] = []string{}
		for _, v := range dictArray {
			vArr := strings.Split(v, "")
			vsArr := strings.Split(vs, "")
			sort.Strings(vArr)
			sort.Strings(vsArr)
			num := 0
			for i := 0; i < len(vArr); i++ {
				if vArr[i] == vsArr[i] {
					num += 1
				}
			}
			if num == (len(vArr) - 1) {
				dict[v] = append(dict[v], vs)
				dict[vs] = append(dict[vs], v)
			}
		}
	}
}

/*
《迷宫问题》
定义一个二维数组：
int maze[5][5] = {
    0, 1, 0, 0, 0,
    0, 1, 0, 1, 0,
    0, 0, 0, 0, 0,
    0, 1, 1, 1, 0,
    0, 0, 0, 1, 0,
};
它表示一个迷宫，其中的1表示墙壁，0表示可以走的路，只能横着走或竖着走，不能斜着走，要求编程序找出从左上角到右下角的最短路线。
**/
func shortRoad(vs []int, vd []int) (bool, [][]int) {
	if vs[0] == vd[0] && vs[1] == vd[1] {
		return true, [][]int{vs}
	}
	gray := [][]int{}
	pathRecord := make(map[string][]int)

	gray = append(gray, vs)
	var flag bool
	for len(gray) > 0 && !flag {
		if gray[0][0] == vd[0] && gray[0][1] == vd[1] {
			flag = true
		}
		toB := gray[0]
		gray = gray[1:]
		if toB[0] == vd[0] && toB[1] == vd[1] {
			break
		}
		temp := [][]int{}
		if toB[0]-1 >= 0 && toB[0]-1 < len(maze) {
			temp = append(temp, []int{toB[0] - 1, toB[1]})
		}
		if toB[0]+1 >= 0 && toB[0]+1 < len(maze) {
			temp = append(temp, []int{toB[0] + 1, toB[1]})
		}
		if toB[1]-1 >= 0 && toB[1]-1 < len(maze) {
			temp = append(temp, []int{toB[0], toB[1] - 1})
		}
		if toB[1]+1 >= 0 && toB[1]+1 < len(maze) {
			temp = append(temp, []int{toB[0], toB[1] + 1})
		}
		for _, v := range temp {
			if maze[v[0]][v[1]] == 0 || (v[0] == vd[0] && v[1] == vd[1]) {
				if _, ok := pathRecord[strconv.Itoa(v[0])+","+strconv.Itoa(v[1])]; !ok {
					gray = append(gray, v)
					if _, ok2 := pathRecord[strconv.Itoa(v[0])+","+strconv.Itoa(v[1])]; !ok2 {
						pathRecord[strconv.Itoa(v[0])+","+strconv.Itoa(v[1])] = toB
					}
				}
			}

		}

	}
	if (len(gray) == 0 || !(gray[0][0] == vd[0] && gray[0][1] == vd[1])) && !flag {
		return false, [][]int{}
	}
	pathArray := [][]int{vd}
	last := vd

	for !(pathRecord[strconv.Itoa(last[0])+","+strconv.Itoa(last[1])][0] == vs[0] && pathRecord[strconv.Itoa(last[0])+","+strconv.Itoa(last[1])][1] == vs[1]) {
		last = pathRecord[strconv.Itoa(last[0])+","+strconv.Itoa(last[1])]
		pathArray = append(pathArray, last)

	}
	pathArray = append(pathArray, vs)
	for i := 0; i < len(pathArray)/2; i++ {
		pathArray[i], pathArray[len(pathArray)-1-i] = pathArray[len(pathArray)-1-i], pathArray[i]
	}
	return true, pathArray

}
