use super::config::ListNode;

pub struct Solution;

impl Solution {
    pub fn remove_nth_from_end(head: Option<Box<ListNode>>, n: i32) -> Option<Box<ListNode>> {
        let mut head = head;
        let mut size = 0;
        let mut p = &head;
        while let Some(node) = p {
            p = &node.next;
            size += 1;
        }

        size -= n;
        if size == 0 {
            return head.unwrap().next;
        }

        let mut cur = head.as_mut().unwrap();
        for _ in 1..size {
            cur = cur.next.as_mut().unwrap();
        }
        cur.next = cur.next.take().unwrap().next;
        head
    }
}

// 提升
// 利用快慢指针
// 我想到的是 用快慢指针一分为二 然后再计算是从中间开始还是从头开始

// 升级版快慢指针
// 先让快指针走 n 步 再两个指针一起走 此时就能定位到需要移除的节点前面
// n + m = size
