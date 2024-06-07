func twoSum(nums []int, target int) []int {


var out []int
out = make([]int, 2)
for i := 0; i < len(nums); i++ {

    for j:=i+1; j <len(nums); j++{
        sum:=nums[i]+nums[j]
        if sum==target{
            out[0]=i
            out[1]=j
            
        }
    }

   }

return out

}
