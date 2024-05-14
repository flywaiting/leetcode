use super::config::ListNode;

pub struct Solution;

impl Solution {
    pub fn next_larger_nodes(head: Option<Box<ListNode>>) -> Vec<i32> {
        let mut list = Vec::new();
        let mut p = head.as_ref();
        while let Some(node) = p {
            p = node.next.as_ref();
            list.push(node.val);
        }
        let mut max = i32::MIN;
        for i in (0..list.len()).rev() {
            if list[i] >= max {
                max = list[i];
                list[i] = 0;
            } else if i > 0 && list[i] > list[i - 1] {
                (list[i], max) = (max, list[i]);
            } else {
                list[i] = max;
            }
        }
        list
    }
}
