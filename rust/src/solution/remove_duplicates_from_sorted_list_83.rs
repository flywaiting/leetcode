use super::config::ListNode;

impl Solution {
    pub fn delete_duplicates(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut head = head; // cur 需要赋予可变权
        let mut cur = &mut head;
        while let Some(node) = cur {
            // 引用下一个节点值 保留其所有权
            while node.next.is_some() && node.val == node.next.as_ref().unwrap().val {
                node.next = node.next.take().unwrap().next;
            }
            cur = &mut node.next;
        }
        head
    }
}

pub struct Solution;
