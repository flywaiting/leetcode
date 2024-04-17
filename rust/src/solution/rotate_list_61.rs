use super::config::ListNode;

pub struct Solution;

impl Solution {
    pub fn rotate_right(head: Option<Box<ListNode>>, k: i32) -> Option<Box<ListNode>> {
        let mut head = Box::new(ListNode {
            val: -1,
            next: head,
        });
        let mut size = 0;
        let mut p = &head;
        while let Some(node) = p.next.as_ref() {
            size += 1;
            p = node
        }

        if size <= 1 || k % size == 0 {
            return head.next;
        }

        size = size - (k % size);
        let mut cur = head.next.as_mut();
        for _ in 1..size {
            cur = cur.unwrap().next.as_mut();
        }
        let mut first = cur.unwrap().next.take();
        let mut last = &mut first;
        while last.is_some() && last.as_ref().unwrap().next.is_some() {
            last = &mut last.as_mut().unwrap().next;
        }
        last.as_mut().unwrap().next = head.next;

        first
    }
}

// 总结
// 思路很简单 计数获得总长度 取余得到最终分割点 拼接

// `rust`的问题在于 所有权限制第一次计数时 不能保留末尾节点以供拼接时使用 因为在分割阶段必须再次拿到可变引用 与前面的可变引用冲突
// 分割时 拿到新的链表头 然后定位到末尾节点 拼接到原始头部即可w
