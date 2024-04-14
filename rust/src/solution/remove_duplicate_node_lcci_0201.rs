use std::collections::HashSet;

use super::config::ListNode;

impl Solution {
    pub fn remove_duplicate_nodes(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        if head.is_none() {
            return head;
        }
        let mut map = HashSet::new();
        let mut head = Box::new(ListNode { val: 0, next: head });
        let mut cur = &mut head;
        while let Some(node) = cur.next.take() {
            if map.contains(&node.val) {
                cur.next = node.next;
            } else {
                map.insert(node.val);
                cur.next = Some(node);
                cur = cur.next.as_mut().unwrap();
            }
        }
        head.next
    }
}
pub struct Solution;

// 总结
// 处理链表时 将数据所有权取出 处理完成之后 再整理所有权 会是一个比较好的习惯 不然每次都在和所有权绕弯子
// HashSet 哈希值集合 相比 `HashMap`不需要再额外的值
