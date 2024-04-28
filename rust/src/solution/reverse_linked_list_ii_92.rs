use super::config::ListNode;

pub struct Solution;

impl Solution {
    pub fn reverse_between(
        head: Option<Box<ListNode>>,
        left: i32,
        right: i32,
    ) -> Option<Box<ListNode>> {
        if left >= right {
            return head;
        }

        let mut head = head;
        let mut res = Box::new(ListNode::new(0));
        let mut cur = &mut res;
        let mut p = None;
        let mut i = 0;
        while let Some(mut node) = head {
            i += 1;
            head = node.next.take();
            if i < left {
                cur = cur.next.insert(node);
            } else {
                node.next = p;
                p = Some(node);
            }
            if i == right {
                break;
            }
        }

        cur.next = p;
        while let Some(ref mut node) = cur.next {
            cur = node;
        }
        cur.next = head;
        res.next
    }
}

// 解题思路
// 思路很简单 `rust`操作还是一样很蛋疼
// 分3个部分进行处理 然后拼接起来
// 因为引用相关限制 需要重新遍历一次反转部分的链表 以到达反转之后的尾部 然后拼接剩余链表元素
