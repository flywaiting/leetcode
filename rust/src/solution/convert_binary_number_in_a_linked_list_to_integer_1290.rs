use super::config::ListNode;

impl Solution {
    pub fn get_decimal_value(head: Option<Box<ListNode>>) -> i32 {
        let mut res = 0;
        let mut cur = &head;
        while let Some(a) = cur {
            res *= 2;
            res += a.val;
            cur = &a.next;
        }
        res
    }
}

pub struct Solution;

#[cfg(test)]
mod tests {
    use crate::solution::config::ListNode;

    use super::Solution;

    #[test]
    fn test_1290() {
        assert_eq!(
            Solution::get_decimal_value(ListNode::create(vec![1, 0, 1])),
            5
        );
        assert_eq!(
            Solution::get_decimal_value(ListNode::create(vec![
                1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0
            ])),
            18880
        );
    }
}
