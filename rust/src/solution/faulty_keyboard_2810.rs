impl Solution {
    pub fn final_string(s: String) -> String {
        s.chars().fold(String::new(), |mut s, c| {
            if c == 'i' {
                s.chars().rev().collect()
            } else {
                s.push(c);
                s
            }
        })
    }

    pub fn final_string_2(s: String) -> String {
        let mut list = Vec::<char>::new();
        let mut rev = false;
        for c in s.chars() {
            if c == 'i' {
                rev = !rev;
            } else if rev {
                list.insert(0, c);
            } else {
                list.push(c);
            }
        }

        if rev {
            list.reverse();
        }
        list.iter().collect()
    }
}

pub struct Solution;

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn test_2810() {
        assert_eq!(Solution::final_string("string".to_string()), "rtsng");
        assert_eq!(Solution::final_string("poiinter".to_string()), "ponter");
        assert_eq!(Solution::final_string_2("string".to_string()), "rtsng");
        assert_eq!(Solution::final_string_2("poiinter".to_string()), "ponter");
    }
}

// 解题思路
// 直觉直接用数组收集字符 遇到特定字符对数组进行翻转

// 学习
// 双端队列
// 一种巧妙的思路 通过标记决定字符收集的位置 最后再同意处理翻转问题 避免在处理过程中频繁的翻转数组
