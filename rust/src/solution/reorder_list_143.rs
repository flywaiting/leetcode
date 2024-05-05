use super::config::ListNode;

pub struct Solution;

impl Solution {
    pub fn reorder_list(head: &mut Option<Box<ListNode>>) {
        if head.is_none() {
            return;
        }

        let mut slow = head.as_ref();
        let mut fast = head.as_ref().unwrap().next.as_ref();

        let mut cnt = 1;
        while fast.is_some() && fast.as_ref().unwrap().next.is_some() {
            cnt += 1;
            slow = slow.as_ref().unwrap().next.as_ref();
            fast = fast.as_ref().unwrap().next.as_ref().unwrap().next.as_ref();
        }
        // if cnt < 2 {
        //     return;
        // }

        let mut p = head.take(); // 1. 去除引用 以便后续节点整理
        let mut cur = p.as_mut();
        for _ in 1..cnt {
            cur = cur.unwrap().next.as_mut();
        }

        let mut tail = cur.unwrap().next.take(); // 同上 反转剩余链表 引用数据无法操作
        let mut last = None;
        while let Some(mut node) = tail {
            tail = node.next.take();
            node.next = last.take();
            last = Some(node);
        }

        let mut cur = p.as_mut().unwrap();
        while let Some(mut node) = last {
            last = node.next.take();
            node.next = cur.next.take();
            cur.next = Some(node); // insert 操作在这里会截断 next，尚不清楚具体发生了什么

            // 需要判定后续节点是否为空 当然 如果直接使用 Option 就不用做判空处理
            if cur.next.as_ref().unwrap().next.is_none() {
                break;
            } else {
                cur = cur.next.as_mut().unwrap().next.as_mut().unwrap();
            }
        }

        *head = p;
    }
}

// 解题思路
// 整体思路不复杂 用双指针找到中点 - 分割反转后半部分链表 - 拼接
// `rust`需要对具体发生什么有细节了解 同时 所有权管理也限制了方便操作。。。
