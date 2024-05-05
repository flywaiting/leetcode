use std::collections::HashMap;

use super::config::ListNode;

pub struct Solution;

impl Solution {
    pub fn remove_zero_sum_sublists(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut dummy = Some(Box::new(ListNode { val: 0, next: head }));
        let mut map = HashMap::new();

        let mut cnt = 0;
        let mut sum = 0;
        let mut p = dummy.as_ref();
        while let Some(node) = p {
            sum += node.val;
            p = node.next.as_ref();
            map.insert(sum, cnt);
            cnt += 1;
        }

        cnt = 0;
        sum = 0;
        let mut p = &mut dummy;
        while let Some(ref mut node) = p {
            sum += node.val;
            let v = map.get(&sum).unwrap();
            while v > &cnt {
                node.next = node.next.as_mut().unwrap().next.take();
                cnt += 1;
            }
            p = &mut node.next;
            cnt += 1;
        }
        // None
        dummy.unwrap().next
    }
}

// 学习
// 前缀和
// 当前缀和结果重复出现时 说明两个相同结果之间的 开闭 区间和为 0

// 自我提升
// 为了不新建节点 用计数的办法来确定 需要跳过的节点
