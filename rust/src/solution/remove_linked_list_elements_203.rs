use super::config::ListNode;

impl Solution {
    pub fn remove_elements(head: Option<Box<ListNode>>, val: i32) -> Option<Box<ListNode>> {
        let mut head = Box::new(ListNode { val: 0, next: head });
        let mut cur = &mut head;
        // 对比 cur.next.as_mut() 无需再额外创建 option 简洁, 直接
        while let Some(ref mut a) = cur.next {
            if a.val == val {
                cur.next = a.next.take(); // 转移所有权
            } else {
                // next 所有权并为转移 可以直接使用
                cur = cur.next.as_mut().unwrap(); // 不能直接赋值, cur, a 都是可变引用, 可能引起悬垂指针或者数据竞争
            }
        }
        head.next
    }
}
pub struct Solution;
