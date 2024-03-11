/*
 * @Author:       Kit-Hung
 * @Date:         2024/3/1 11:28
 * @Description： 数据交换示例
 */
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(nums)

	swap(0, 2, nums)
	fmt.Println(nums)
}

// 交换切片中下标为 i 、 j 的数
// 注意：i 和 j 不能相等
func swap(i, j int, nums []int) {
	nums[i] = nums[i] ^ nums[j]
	nums[j] = nums[i] ^ nums[j]
	nums[i] = nums[i] ^ nums[j]
}
