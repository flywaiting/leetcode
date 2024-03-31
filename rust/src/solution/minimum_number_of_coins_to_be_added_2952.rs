// 解题思路
// 递归 假设 【1, x-1] 已经实现 然后扩展到 [1, 2x-1]
// 集合的扩展特性是解题关键 虽然有个大概思路 但是没能整理成相关逻辑 最终还是看了题解
impl Solution {
    pub fn minimum_added_coins(coins: Vec<i32>, target: i32) -> i32 {
        let mut coins = coins;
        coins.sort_unstable();
        let mut ans = 0;
        let mut x = 1;
        let mut i = 0;
        while x <= target {
            if i < coins.len() && coins[i] <= x {
                x += coins[i];
                i += 1;
            } else {
                ans += 1;
                x *= 2;
            }
        }
        ans
    }
}

pub struct Solution;

#[cfg(test)]
mod tests {
    use super::Solution;
    #[test]
    fn test_2952() {
        assert_eq!(Solution::minimum_added_coins(vec![1, 4, 10], 19), 2);
        assert_eq!(
            Solution::minimum_added_coins(vec![1, 4, 10, 5, 7, 19], 19),
            1
        );
        assert_eq!(Solution::minimum_added_coins(vec![1, 1, 1], 20), 3);
    }
}
