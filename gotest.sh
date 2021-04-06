#!/usr/bin/env bash
function go_test_category ()
{
	mkdir -p gotest-result
    tempfile=./gotest-result/profile-go-$1.out
    resultfile=./gotest-result/result-go-$1.txt
    set -e
    echo "" > $resultfile
    for d in $(go list ./leetcode-go/$1/... | grep -v vendor); do
        echo $d
        go test -coverprofile=$tempfile -covermode=atomic $d
        if [ -f $tempfile ]; then
            cat $tempfile >> $resultfile
            rm $tempfile
        fi
    done
}

#------------------------------------menu-----------------------------------------------
function operation_menu ()
{
	cat << EOF
`echo -e "\033[34m--------------------------------------------------- \033[0m"`
`echo -e "\033[34m|************ leetcode category menu ************| \033[0m"`
`echo -e "\033[34m--------------------------------------------------- \033[0m"`
`echo -e "\033[34m  1  数组(array) \033[0m"`
`echo -e "\033[34m  2  动态规划(dynamic-programming) \033[0m"`
`echo -e "\033[34m  3  字符串(string) \033[0m"`
`echo -e "\033[34m  4  数学(math) \033[0m"`
`echo -e "\033[34m  5  贪心算法(greedy) \033[0m"`
`echo -e "\033[34m  6  深度优先搜索(depth-first-search) \033[0m"`
`echo -e "\033[34m  7  树(tree) \033[0m"`
`echo -e "\033[34m  8  哈希表(hash-table) \033[0m"`
`echo -e "\033[34m  9 二分查找(binary-search) \033[0m"`
`echo -e "\033[34m  10 广度优先搜索(breadth-first-search) \033[0m"`
`echo -e "\033[34m  11 排序(sort) \033[0m"`
`echo -e "\033[34m  12 双指针(two-pointers) \033[0m"`
`echo -e "\033[34m  13 回溯算法(backtracking) \033[0m"`
`echo -e "\033[34m  14 设计(design) \033[0m"`
`echo -e "\033[34m  15 栈(stack) \033[0m"`
`echo -e "\033[34m  16 位运算(bit-manipulation) \033[0m"`
`echo -e "\033[34m  17 图(graph) \033[0m"`
`echo -e "\033[34m  18 堆(heap) \033[0m"`
`echo -e "\033[34m  19 链表(linked-list) \033[0m"`
`echo -e "\033[34m  20 递归(recursion) \033[0m"`
`echo -e "\033[34m  21 并查集(union-find) \033[0m"`
`echo -e "\033[34m  22 sliding-window \033[0m"`
`echo -e "\033[34m  23 字典树(trie) \033[0m"`
`echo -e "\033[34m  24 分治算法(divide-and-conquer) \033[0m"`
`echo -e "\033[34m  25 线段树(segment-tree) \033[0m"`
`echo -e "\033[34m  26 ordered-map \033[0m"`
`echo -e "\033[34m  27 队列(queue) \033[0m"`
`echo -e "\033[34m  28 几何(geometry) \033[0m"`
`echo -e "\033[34m  29 line-sweep \033[0m"`
`echo -e "\033[34m  30 极小化极大(minimax) \033[0m"`
`echo -e "\033[34m  31 树状数组(binary-indexed-tree) \033[0m"`
`echo -e "\033[34m  32 脑筋急转弯(brainteaser) \033[0m"`
`echo -e "\033[34m  33 拓扑排序(topological-sort) \033[0m"`
`echo -e "\033[34m  34 dequeue \033[0m"`
`echo -e "\033[34m  35 random \033[0m"`
`echo -e "\033[34m  36 二叉搜索树(binary-search-tree) \033[0m"`
`echo -e "\033[34m  37 rolling-hash \033[0m"`
`echo -e "\033[34m  38 suffix array\033[0m"`
`echo -e "\033[34m  39 拒绝采样(rejection-sampling) \033[0m"`
`echo -e "\033[34m  40 蓄水池抽样(reservoir-sampling) \033[0m"`
`echo -e "\033[34m  41 meet-in-the-middle \033[0m"`
`echo -e "\033[34m  42 记忆化(memoization) \033[0m"`
`echo -e "\033[34m  43 oop \033[0m"`
EOF
read -p "input your select num：" num1
select_menu $num1
}

function select_menu ()
{
    category="array"
    case $1 in
        1)
            category="array"
            ;;
        2)
            category="dynamic-programming"
            ;;
        3)
            category="string"
            ;;
        4)
            category="math"
            ;;
        5)
            category="greedy"
            ;;
        6)
            category="depth-first-search"
            ;;
        7)
            category="tree"
            ;;
        8)
            category="hash-table"
            ;;
        9)
            category="binary-search"
            ;;
        10)
            category="breadth-first-search"
            ;;
        11)
            category="sort"
            ;;
        12)
            category="two-pointers"
            ;;
        13)
            category="backtracking"
            ;;
        14)
            category="design"
            ;;
        15)
            category="stack"
            ;;
        16)
            category="bit-manipulation"
        	;;
        17)
            category="graph"
            ;;
        18)
            category="heap"
            ;;
        19)
            category="linked-list"
            ;;
        20)
            category="recursion"
            ;;
        21)
            category="union-find"
            ;;
        22)
            category="sliding-window"
            ;;
        23)
            category="trie"
            ;;
        24)
            category="divide-and-conquer"
         	;;
        25)
            category="segment-tree"
            ;;
        26)
            category="ordered-map"
            ;;
        27)
            category="queue"
            ;;
        28)
            category="geometry"
            ;;
        29)
            category="line-sweep"
            ;;
        30)
            category="minimax"
            ;;
        31)
            category="binary-indexed-tree"
            ;;
        32)
            category="brainteaser"
            ;;
        33)
            category="topological-sort"
            ;;
        34)
            category="dequeue"
         	;;
        35)
            category="random"
            ;;
        36)
            category="binary-search-tree"
            ;;
        37)
            category="rolling-hash"
            ;;
        38)
            category="suffix array"
            ;;
        39)
            category="rejection-sampling"
            ;;
        40)
            category="reservoir-sampling"
            ;;
        41)
            category="meet-in-the-middle"
            ;;
        42)
            category="memoization"
            ;;
        43)
            category="oop"
            ;;
        *)
            echo "输入错误，请重新选择菜单!!!"
            operation_menu
            ;;
    esac
    go_test_category $category
	operation_menu
}

operation_menu
