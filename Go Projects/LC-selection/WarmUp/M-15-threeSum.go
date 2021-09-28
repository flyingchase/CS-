package WarmUp

import "sort"

/*
public List<List<Integer>> threeSum(int[] nums) {
        List<List<Integer>> res = new ArrayList<>();
        Arrays.sort(nums);
        if (nums.length <= 2) {
            return res;
        }
        for (int i = 0; i < nums.length - 2; i++) {
            if (i>0&&nums[i]==nums[i-1]) {
                continue;
            }
            int l = i + 1, r = nums.length - 1;
            while (l < r) {
                int sum = nums[i] + nums[l] + nums[r];
                if (sum == 0) {
                    res.add(new ArrayList<>(Arrays.asList(nums[i], nums[l], nums[r])));
                    while (l < r && nums[l] == nums[l + 1]) {
                        l++;
                    }
                    while (l < r && nums[r] == nums[r - 1]) {
                        r--;
                    }
                    l++;
                    r--;
                } else if (sum > 0) {
                    r--;
                } else {
                    l++;
                }
            }
        }
        return res;
    }
*/
// 四数求和
//class Solution {
//public List<List<Integer>> fourSum(int[] nums, int target) {
//        List<List<Integer>> res = new ArrayList<>();
//        Arrays.sort(nums);
//        if (nums == null || nums.length < 4) {
//            return res;
//        }
//
//        for (int i = 0; i < nums.length - 3; i++) {
//              if (i>0&&nums[i]==nums[i-1]){
//                continue;
//            }
//            for (int j = i + 1; j < nums.length - 2; j++) {
//                if (j>i+1&&nums[j]==nums[j-1]) {
//                    continue;
//                }
//                int l = j + 1, r = nums.length - 1;
//
//                while (l < r) {
//                    int sum = nums[i] + nums[j] + nums[l] + nums[r];
//                    if (sum == target) {
//                        res.add(Arrays.asList(nums[i], nums[j], nums[l], nums[r]));
//                        while (l < r && nums[l] == nums[l + 1]) {
//                            l++;
//                        }
//                        while (l < r && nums[r] == nums[r - 1]) {
//                            r--;
//                        }
//                        l++;
//                        r--;
//                    } else if (sum>target) {
//                        r--;
//                    }else {
//                        l++;
//                    }
//                }
//            }
//        }
//        return res;
//    }
//}
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) <= 2 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				lists:=make([]int,0)
				lists=append(lists,nums[i],nums[l],nums[r])
				res = append(res, lists)
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}else if sum>0 {
				r--
			}else {
			    l++
			}

		}

	}
	return res
}
