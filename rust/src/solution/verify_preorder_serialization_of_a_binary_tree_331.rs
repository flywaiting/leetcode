impl Solution {
    pub fn is_valid_serialization(preorder: String) -> bool {
        let mut is_pre = true;
        let mut node = 1;
        for c in preorder.chars() {
            if node == 0 {
                return false;
            }
            if c == ',' {
                is_pre = true;
                continue;
            }

            if c == '#' {
                node -= 1;
                continue;
            }

            if is_pre {
                node += 1;
                is_pre = false;
            }
        }

        node == 0
    }
}

pub struct Solution;

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn test_331() {
        assert!(Solution::is_valid_serialization(
            "9,3,4,#,#,1,#,#,2,#,6,#,#".to_string()
        ));
        assert!(!Solution::is_valid_serialization("1,#".to_string()));
        assert!(!Solution::is_valid_serialization("9,#,#,1".to_string()));
    }
}

// 解题思路
// 一开始寻找规律 发现就是个验证入栈、出栈的过程
// 其中的问题点有 堆栈管理 计数引起的移动问题
// 而 rust 本身集合就比较难搞 至少当前对我而言 还是磕磕绊绊的在学习

// 学习提升
// 观摩他人的解题 发现大多数都是对字符串进行切分 然后当一个个的数据进行处理 这当然是比较简单、符合直觉的 但是需要而外的开销 这并不是必要的
// 其实 堆栈的计数再深入观察一下 会发现 根本不需要用到堆栈 这就是个简单的计数问题
// 首先 计数引起的移动问题 确保计数不需要而外的特殊处理 压缩成一个数完全可以满足要求
// 其次 管理一个数 避免了堆栈的增减相关操作 都不需要去处理这些东西
